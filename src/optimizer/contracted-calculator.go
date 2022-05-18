package optimizer

func CalculateAnnualContracted(splitByPeriods bool, contractedPowers [6]float64, powerPrices map[string][6]float64, tariffPeriods map[string]map[string]interface{}) map[string]interface{} {
	annualImport := 0.0
	splitedInfo := make(map[string]interface{})
	periodsCache := []string{"0", "1", "2", "3", "4", "5", "6"}
	monthKey := ""

	for day := range tariffPeriods {
		if splitByPeriods {
			monthKey, _ = dateToYearMonth(day)

			if splitedInfo[monthKey] == nil {
				defaultInfo := make(map[string]float64)
				defaultInfo["1"] = 0.0
				defaultInfo["2"] = 0.0
				defaultInfo["3"] = 0.0
				defaultInfo["4"] = 0.0
				defaultInfo["5"] = 0.0
				defaultInfo["6"] = 0.0

				splitedInfo[monthKey] = make(map[string]float64)
				splitedInfo[monthKey] = defaultInfo
			}
		}

		powerPricesDay := powerPrices[day]

		for period := 1; period <= 6; period++ {
			strPeriod := periodsCache[period]
			periodKey := period - 1
			value := (contractedPowers[periodKey] * powerPricesDay[periodKey])
			annualImport += value
			globalCounter++

			if splitByPeriods {
				splitedInfo[monthKey].(map[string]float64)[strPeriod] += value
			}
		}
	}

	splitedInfo["annual"] = annualImport

	return splitedInfo
}
