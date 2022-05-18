package optimizer

import (
	"math"
	"strconv"
	"time"
)

func CalculateAnnualExcess(splitByPeriods bool, powerInf powerInfo) map[string]interface{} {
	switch powerInf.CalcType {
	case "maximeters":
		return CalculateAnnualExcessByMaximeters(splitByPeriods, powerInf.ExcessPrices, powerInf.ContractedPowers, powerInf.Maximeters)
	case "curve":
		return CalculateAnnualExcessByCurve(splitByPeriods, powerInf.ExcessPrices, powerInf.ContractedPowers, powerInf.Curve, powerInf.Kps, powerInf.TariffPeriods)
	}

	return nil
}

func CalculateAnnualExcessByMaximeters(splitByPeriods bool, excessPrices []map[string]interface{}, contractedPowers [6]float64, maximeters []map[string]interface{}) map[string]interface{} {
	annualImport := 0.0
	splitedInfo := make(map[string]interface{})
	periodsCache := []string{"0", "1", "2", "3", "4", "5", "6"}

	for _, maximeter := range maximeters {
		excessPrice := SelectExcessPrice(excessPrices, maximeter["end"].(string))
		monthKey := ""

		if splitByPeriods {
			monthKey = maximeter["end"].(string)

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

		for period := 1; period <= 6; period++ {
			strPeriod := periodsCache[period]
			excessPower := (maximeter[strPeriod].(float64) - contractedPowers[period-1])
			globalCounter++

			if excessPower > 0 {
				value := (excessPrice * excessPower)
				annualImport += value
				globalCounter++

				if splitByPeriods {
					splitedInfo[monthKey].(map[string]float64)[strPeriod] = value
				}
			}
		}
	}

	splitedInfo["annual"] = annualImport

	return splitedInfo
}

func CalculateAnnualExcessByCurve(splitByPeriods bool, excessPrices []map[string]interface{}, contractedPowers [6]float64, curve map[string][25][4]float64, kps []map[string]interface{}, tariffPeriods map[string]map[string]interface{}) map[string]interface{} {
	annualImport := 0.0
	monthlyPeriods := make(map[string]map[string]interface{})
	splitedInfo := make(map[string]interface{})
	layout1 := "2006-01-02"
	lastDay := ""

	periodsCache := []string{"0", "1", "2", "3", "4", "5", "6"}
	//quartersCache := []string{"q0", "q1", "q2", "q3", "q4"}

	for day, tarifPeriodData := range tariffPeriods {
		curveDay, curveDayOk := curve[day]
		lastDay = day

		if !curveDayOk {
			continue
		}

		monthKey, tmpMonthKey := dateToYearMonth(day)
		tmpMontlyPeriod := monthlyPeriods[monthKey]

		// First I need to register the last date of the month with data, this will be used to select prices and kp
		if tmpMontlyPeriod == nil {
			tmpMontlyPeriod = make(map[string]interface{}, 8)
			tmpMontlyPeriod["1"] = 0.0
			tmpMontlyPeriod["2"] = 0.0
			tmpMontlyPeriod["3"] = 0.0
			tmpMontlyPeriod["4"] = 0.0
			tmpMontlyPeriod["5"] = 0.0
			tmpMontlyPeriod["6"] = 0.0
			tmpMontlyPeriod["end"] = day
			tmpMontlyPeriod["key"] = monthKey
		} else {
			endTime, _ := time.Parse(layout1, tmpMontlyPeriod["end"].(string))

			if endTime.Before(tmpMonthKey) {
				tmpMontlyPeriod["end"] = day
			}
		}

		if splitByPeriods && splitedInfo[monthKey] == nil {
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

		for hour, intPeriod := range tarifPeriodData {
			strPeriod := intPeriod.(string)
			atoiPeriod, _ := strconv.Atoi(strPeriod)
			inHour, _ := strconv.Atoi(hour)
			period := atoiPeriod - 1

			if period < 0 || period > 5 {
				continue
			}

			contractedPower := contractedPowers[period]

			if contractedPower <= 0.0 {
				continue
			}

			curveDayHour := curveDay[inHour-1]
			tmpMontlyPeriodValue := tmpMontlyPeriod[strPeriod].(float64)

			calculateQuarter(&curveDayHour[0], &contractedPower, &tmpMontlyPeriodValue)
			calculateQuarter(&curveDayHour[1], &contractedPower, &tmpMontlyPeriodValue)
			calculateQuarter(&curveDayHour[2], &contractedPower, &tmpMontlyPeriodValue)
			calculateQuarter(&curveDayHour[3], &contractedPower, &tmpMontlyPeriodValue)

			tmpMontlyPeriod[strPeriod] = tmpMontlyPeriodValue
		}

		monthlyPeriods[monthKey] = tmpMontlyPeriod
	}

	/*
		If I only have 1 kp or 1 excessPrice, I'm getting them out of the loop. This is for performance reasons
	*/
	kp := make(map[string]interface{})
	selectKp := true

	if len(kps) == 1 {
		kp = SelectKp(kps, lastDay)
		selectKp = false
	}

	excessPrice := 0.0
	selectExcessPrice := true

	if len(excessPrices) == 1 {
		excessPrice = SelectExcessPrice(excessPrices, lastDay)
		selectExcessPrice = false
	}

	for _, monthlyInfo := range monthlyPeriods {
		day := monthlyInfo["end"].(string)

		if selectKp {
			kp = SelectKp(kps, day)
		}

		if selectExcessPrice {
			excessPrice = SelectExcessPrice(excessPrices, day)
		}

		for period := 1; period <= 6; period++ {
			strPeriod := periodsCache[period]
			value := 0.0

			if monthlyInfo[strPeriod].(float64) > 0.0 {
				value = kp[strPeriod].(float64) * excessPrice * math.Sqrt(monthlyInfo[strPeriod].(float64))
				annualImport += value
				globalCounter++
			}

			if splitByPeriods {
				splitedInfo[monthlyInfo["key"].(string)].(map[string]float64)[strPeriod] = value + splitedInfo[monthlyInfo["key"].(string)].(map[string]float64)[strPeriod]
			}
		}
	}

	splitedInfo["annual"] = annualImport

	return splitedInfo
}

func calculateQuarter(quarter *float64, power *float64, tmpMontlyPeriodValue *float64) {
	tmpExcess := *quarter - *power
	globalCounter++

	if tmpExcess > 0 {
		*tmpMontlyPeriodValue = *tmpMontlyPeriodValue + math.Pow(tmpExcess, 2)
		globalCounter++
	}
}
