package optimizer

type powerInfo struct {
	ContractedPowers [6]float64
	PowerPrices      map[string][6]float64
	TariffPeriods    map[string]map[string]interface{}
	Maximeters       []map[string]interface{}
	Curve            map[string][25][4]float64
	ExcessPrices     []map[string]interface{}
	CalcType         string
	Tariff           string
	Kps              []map[string]interface{}
}

func InitPowerInf(info map[string]interface{}) powerInfo {
	var powerInf powerInfo

	powerInf.ContractedPowers = parsePeriodsSlice(info["contractedPowers"])
	powerInf.PowerPrices = parsePeriodsMap(info["powerPrices"])
	powerInf.TariffPeriods = parseNested(info["tariffPeriods"])

	/*
		I saved the value multiplied by 2, because at all the operations I have to do this operation, so doing this operation I can
		optimize the process, only if the calculation method is by maximeters
	*/
	powerInf.ExcessPrices = parseNestedSlice(info["excessPrices"])

	for day, powerPriceInfo := range powerInf.PowerPrices {
		var tmpInfo [6]float64

		for strPeriod, powerPrice := range powerPriceInfo {
			tmpInfo[strPeriod] = RoundNumber(powerPrice, 6)
		}

		powerInf.PowerPrices[day] = tmpInfo
	}

	if info["curve"] == nil {
		powerInf.Maximeters = parseNestedSlice(info["maximeters"])
		powerInf.CalcType = "maximeters"

		for key, excessPrice := range powerInf.ExcessPrices {
			excessPrice["price"] = RoundNumber(excessPrice["price"].(float64)*2, 6)
			powerInf.ExcessPrices[key] = excessPrice
		}
	} else {
		tmpCurve := parsePeriodsCurve(info["curve"])
		powerInf.Kps = parseNestedSlice(info["kps"])
		powerInf.CalcType = "curve"
		curve := make(map[string][25][4]float64)

		for day, curveDay := range tmpCurve {
			var curveDayInfo [25][4]float64

			for hour, curveHour := range curveDay {
				for quarter, cuveQuarter := range curveHour {
					curveDayInfo[hour][quarter] = cuveQuarter
				}
			}

			curve[day] = curveDayInfo
		}

		powerInf.Curve = curve

		for key, excessPrice := range powerInf.ExcessPrices {
			excessPrice["price"] = RoundNumber(excessPrice["price"].(float64), 6)
			powerInf.ExcessPrices[key] = excessPrice
		}
	}

	return powerInf
}
