package optimizer_curve_test

import (
	"log"
	"os"
	"power-optimizer/src/optimizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

var curveData = make(map[string]interface{})
var curveImportsData []map[string]interface{}

func TestMain(m *testing.M) {
	// Before tests
	initData()

	// Run tests
	exitVal := m.Run()

	// After tests

	// Exit tests
	os.Exit(exitVal)
}

/*
	Reading the curve json and curve imports json. They contain the data which will be used for tests
*/
func initData() {
	curveJson, curveErr := os.Open("./../data/curve.json")
	defer curveJson.Close()

	if curveErr != nil {
		log.Fatal(curveErr)
	}

	curveImportsJson, curveImportsErr := os.Open("./../data/curve-imports.json")
	defer curveImportsJson.Close()

	if curveImportsErr != nil {
		log.Fatal(curveImportsErr)
	}

	curve, parseCurveErr := optimizer.ParseBodyMap(curveJson)

	if parseCurveErr != nil {
		log.Fatal(parseCurveErr)
	}

	curveImports, parseCurveImportsErr := optimizer.ParseBodyArray(curveImportsJson)

	if parseCurveImportsErr != nil {
		log.Fatal(parseCurveImportsErr)
	}

	curveData = curve
	curveImportsData = curveImports
}

func TestCalculateCurrentAnnualContracted(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualContracted(splitByPeriods, powerInf.ContractedPowers, powerInf.PowerPrices, powerInf.TariffPeriods)

	assert.Equal(t, 1542.180870, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["1"], 6))
	assert.Equal(t, 1401.389100, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["2"], 6))
	assert.Equal(t, 789.634170, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["3"], 6))
	assert.Equal(t, 1130.546130, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["4"], 6))
	assert.Equal(t, 248.721060, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["5"], 6))
	assert.Equal(t, 147.965790, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["6"], 6))

	assert.Equal(t, 1492.433100, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["1"], 6))
	assert.Equal(t, 1356.183000, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["2"], 6))
	assert.Equal(t, 764.162100, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["3"], 6))
	assert.Equal(t, 1094.076900, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["4"], 6))
	assert.Equal(t, 240.697800, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["5"], 6))
	assert.Equal(t, 143.192700, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["6"], 6))

	assert.Equal(t, 1392.937560, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["1"], 6))
	assert.Equal(t, 1265.770800, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["2"], 6))
	assert.Equal(t, 713.217960, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["3"], 6))
	assert.Equal(t, 1021.138440, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["4"], 6))
	assert.Equal(t, 224.651280, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["5"], 6))
	assert.Equal(t, 133.646520, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["6"], 6))

	assert.Equal(t, 61937.404800, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualExcessByCurve(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualExcessByCurve(splitByPeriods, powerInf.ExcessPrices, powerInf.ContractedPowers, powerInf.Curve, powerInf.Kps, powerInf.TariffPeriods)

	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["1"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["2"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["3"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["4"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["5"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-05"].(map[string]float64)["6"], 6))

	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["1"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["2"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["3"], 6))
	assert.Equal(t, 2345.860499, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["4"], 6))
	assert.Equal(t, 152.669814, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["5"], 6))
	assert.Equal(t, 105.165664, optimizer.RoundNumber(response["2022-04"].(map[string]float64)["6"], 6))

	assert.Equal(t, 2603.695978, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualExcess(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualExcess(splitByPeriods, powerInf)

	assert.Equal(t, 2603.695978, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualPower(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	splitByPeriods := true
	totalImport, annualContracted, annualExcess := optimizer.CalculateAnnualPower(splitByPeriods, powerInf)

	assert.Equal(t, 61937.404800, optimizer.RoundNumber(annualContracted["annual"].(float64), 6))
	assert.Equal(t, 2603.695978, optimizer.RoundNumber(annualExcess["annual"].(float64), 6))
	assert.Equal(t, 64541.100000, totalImport)
}

func TestGetPowerCombinations(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	response := optimizer.GetPowerCombinations(powerInf.ContractedPowers, powerInf.CalcType)

	assert.Equal(t, 100, len(response))

	assert.Equal(t, 811.00, response[0][0])
	assert.Equal(t, 811.00, response[0][1])
	assert.Equal(t, 811.00, response[0][2])
	assert.Equal(t, 1471.00, response[0][3])
	assert.Equal(t, 1471.00, response[0][4])
	assert.Equal(t, 1471.00, response[0][5])

	assert.Equal(t, 811.00, response[2][0])
	assert.Equal(t, 811.00, response[2][1])
	assert.Equal(t, 811.00, response[2][2])
	assert.Equal(t, 1470.00, response[2][3])
	assert.Equal(t, 1470.00, response[2][4])
	assert.Equal(t, 1471.00, response[2][5])

	assert.Equal(t, 810.00, response[23][0])
	assert.Equal(t, 810.00, response[23][1])
	assert.Equal(t, 811.00, response[23][2])
	assert.Equal(t, 1470.00, response[23][3])
	assert.Equal(t, 1470.00, response[23][4])
	assert.Equal(t, 1470.00, response[23][5])

	assert.Equal(t, 809.00, response[61][0])
	assert.Equal(t, 810.00, response[61][1])
	assert.Equal(t, 810.00, response[61][2])
	assert.Equal(t, 1470.00, response[61][3])
	assert.Equal(t, 1471.00, response[61][4])
	assert.Equal(t, 1471.00, response[61][5])
}

func TestSelectBestImport(t *testing.T) {
	/*
		This test was created with a source info different than the current. I decided to keep this info, because curve.json and curve-imports.json are
		independent.
	*/
	newImport, selectedPowers := optimizer.SelectBestImport(curveImportsData, 153585.364665)

	assert.Equal(t, 2297.00, selectedPowers[0])
	assert.Equal(t, 2297.00, selectedPowers[1])
	assert.Equal(t, 2297.00, selectedPowers[2])
	assert.Equal(t, 2297.00, selectedPowers[3])
	assert.Equal(t, 2297.00, selectedPowers[4])
	assert.Equal(t, 2948.00, selectedPowers[5])

	assert.Equal(t, 153518.87, newImport)
}
