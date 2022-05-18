package optimizer

import "time"

func SelectBestImport(imports []map[string]interface{}, currentImport float64) (float64, [6]float64) {
	var selectedPowers [6]float64
	var importValue float64

	for _, importInfo := range imports {
		importValue = importInfo["import"].(float64)

		if importValue < currentImport {
			currentImport = importValue

			selectedPowers[0] = importInfo["1"].(float64)
			selectedPowers[1] = importInfo["2"].(float64)
			selectedPowers[2] = importInfo["3"].(float64)
			selectedPowers[3] = importInfo["4"].(float64)
			selectedPowers[4] = importInfo["5"].(float64)
			selectedPowers[5] = importInfo["6"].(float64)
		}
	}

	return currentImport, selectedPowers
}

// SelectExcessPrice Selecting the excess price which will apply to the priceDate
func SelectExcessPrice(excessPrices []map[string]interface{}, priceDate string) float64 {
	excessPrice := 0.0
	sampleDate := "2006-01-02"

	priceDateTime, _ := time.Parse(sampleDate, priceDate)

	for _, excessPriceData := range excessPrices {
		startDateTime, _ := time.Parse(sampleDate, excessPriceData["init"].(string))
		endDateTime, _ := time.Parse(sampleDate, excessPriceData["end"].(string))

		if isDateBetween(priceDateTime, startDateTime, endDateTime) {
			excessPrice = excessPriceData["price"].(float64)
			break
		}
	}

	return excessPrice
}

// SelectKp Selecting the kp which will be applied for the day
func SelectKp(kps []map[string]interface{}, day string) map[string]interface{} {
	sampleDate := "2006-01-02"
	dayTime, _ := time.Parse(sampleDate, day)

	for _, kpsData := range kps {
		startDateTime, _ := time.Parse(sampleDate, kpsData["init"].(string))
		endDateTime, _ := time.Parse(sampleDate, kpsData["end"].(string))

		if isDateBetween(dayTime, startDateTime, endDateTime) {
			return kpsData
		}
	}

	return make(map[string]interface{})
}
