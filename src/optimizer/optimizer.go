package optimizer

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var globalCounter int

func CalculateOptimizedContractedPowers(info map[string]interface{}) (map[string]interface{}, int) {
	globalCounter = 0

	response := make(map[string]interface{})
	/*
		To reduce the amount of duplicated info and operations, I make here the transformation of all interface info to the correct
		types and then setting the values to nil, in order to free memory
	*/
	currentContractedPowers := info["contractedPowers"]
	powerInf := InitPowerInf(info)
	httpStatus := 404

	currentImport, currentAnnualContracted, currentAnnualExcess := CalculateAnnualPower(true, powerInf)
	newImport, selectedPowers := DoOptimization(powerInf, currentImport)

	response["currentImport"] = currentImport
	response["currentAnnualContracted"] = currentAnnualContracted
	response["currentAnnualExcess"] = currentAnnualExcess
	response["currentContractedPowers"] = currentContractedPowers
	response["optimized"] = false

	if currentImport != newImport {
		response["optimized"] = true
		response["newImport"] = newImport

		powerInf.ContractedPowers = selectedPowers
		_, newAnnualContracted, newAnnualExcess := CalculateAnnualPower(true, powerInf)

		response["newPowers"] = selectedPowers
		response["newAnnualContracted"] = newAnnualContracted
		response["newAnnualExcess"] = newAnnualExcess

		httpStatus = 200
	}

	return response, httpStatus
}

func DoOptimization(powerInf powerInfo, currentImport float64) (float64, [6]float64) {
	var imports []map[string]interface{}
	var cleanImports []map[string]interface{}

	importsBackup := make(map[string]map[string]interface{})

	var wg sync.WaitGroup
	var newImport float64

	contractedPowers := powerInf.ContractedPowers
	selectedPowers := contractedPowers

	fmt.Println("currentImport", currentImport)

	count := 0

	for {
		powerCombinations := GetPowerCombinations(contractedPowers, powerInf.CalcType)
		wg.Add(len(powerCombinations))
		imports = cleanImports

		for _, combination := range powerCombinations {
			/*
				I'm saving all the time the previous calculated imports. If the power combination is in the backup, I use it instead of calculate
				again the same values. I only save 1 iteration for prevent an extremly RAM usage of the process.

				I also need to execute a go function, because it expects the total of power combinations threads.
			*/
			if len(importsBackup) > 0 {
				importsKey := getImportsBackupKey(combination)

				if importsBackup[importsKey] != nil && importsBackup[importsKey]["import"] != nil {
					imports = append(imports, importsBackup[importsKey])
					go func(wg *sync.WaitGroup) {
						defer wg.Done()
					}(&wg)
					continue
				}
			}

			powerInf.ContractedPowers = combination

			go func(wg *sync.WaitGroup, imports *[]map[string]interface{}, powerInf powerInfo) {
				defer wg.Done()

				tmpContractedPowers := powerInf.ContractedPowers

				newImport, _, _ := CalculateAnnualPower(false, powerInf)
				tmpImport := make(map[string]interface{})

				tmpImport["1"] = tmpContractedPowers[0]
				tmpImport["2"] = tmpContractedPowers[1]
				tmpImport["3"] = tmpContractedPowers[2]
				tmpImport["4"] = tmpContractedPowers[3]
				tmpImport["5"] = tmpContractedPowers[4]
				tmpImport["6"] = tmpContractedPowers[5]

				tmpImport["import"] = newImport

				*imports = append(*imports, tmpImport)
			}(&wg, &imports, powerInf)
		}
		wg.Wait()
		time.Sleep(20 * time.Millisecond)

		importsBackup = parseToImportsBackup(imports)
		newImport, selectedPowers = SelectBestImport(imports, currentImport)
		count++

		if newImport >= currentImport {
			break
		}

		currentImport = newImport
		contractedPowers = selectedPowers
	}

	fmt.Println("Selected import:", currentImport)
	fmt.Println("Selected powers:", contractedPowers)
	fmt.Println("Total iterations:", count)
	fmt.Println("Total operations:", globalCounter)

	return currentImport, contractedPowers
}

func getImportsBackupKey(powers [6]float64) string {
	var builder strings.Builder

	for period := 0; period < 6; period++ {
		builder.WriteString(strconv.FormatFloat(powers[period], 'f', -1, 64))

		if period < 6 {
			builder.WriteString("-")
		}
	}

	return builder.String()
}

func CalculateAnnualPower(splitByPeriods bool, powerInf powerInfo) (float64, map[string]interface{}, map[string]interface{}) {
	annualContracted := CalculateAnnualContracted(splitByPeriods, powerInf.ContractedPowers, powerInf.PowerPrices, powerInf.TariffPeriods)
	annualExcess := CalculateAnnualExcess(splitByPeriods, powerInf)

	fixedImport := RoundNumber(annualContracted["annual"].(float64), 6)
	excessImport := RoundNumber(annualExcess["annual"].(float64), 6)
	totalImport := RoundNumber(fixedImport+excessImport, 2)

	return totalImport, annualContracted, annualExcess
}
