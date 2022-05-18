package optimizer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func parseToImportsBackup(imports []map[string]interface{}) map[string]map[string]interface{} {
	importsBackup := make(map[string]map[string]interface{})

	for _, importInfo := range imports {
		var powers [6]float64

		powers[0] = importInfo["1"].(float64)
		powers[1] = importInfo["2"].(float64)
		powers[2] = importInfo["3"].(float64)
		powers[3] = importInfo["4"].(float64)
		powers[4] = importInfo["5"].(float64)
		powers[5] = importInfo["6"].(float64)

		importKey := getImportsBackupKey(powers)
		importsBackup[importKey] = importInfo
	}

	return importsBackup
}

func parseNested(f interface{}) map[string]map[string]interface{} {
	elements := make(map[string]map[string]interface{})
	tmpElements := f.(map[string]interface{})

	for key, value := range tmpElements {
		elements[key] = value.(map[string]interface{})
	}

	return elements
}

func parseNestedSlice(f interface{}) []map[string]interface{} {
	var elements []map[string]interface{}
	tmpElements := f.([]interface{})

	for _, value := range tmpElements {
		elements = append(elements, value.(map[string]interface{}))
	}

	return elements
}

func parsePeriodsSlice(f interface{}) [6]float64 {
	var elements [6]float64
	tmpElements := f.([]interface{})

	for key, value := range tmpElements {
		elements[key] = value.(float64)
	}

	return elements
}

func parsePeriodsMap(f interface{}) map[string][6]float64 {
	elements := make(map[string][6]float64)
	mapElements := f.(map[string]interface{})

	for mapKey, mapInterface := range mapElements {
		tmpElements := mapInterface.([]interface{})
		var periods [6]float64

		for key, value := range tmpElements {
			periods[key] = value.(float64)
		}

		elements[mapKey] = periods
	}

	return elements
}

func parsePeriodsCurve(f interface{}) map[string][25][4]float64 {
	elements := make(map[string][25][4]float64)
	mapElements := f.(map[string]interface{})

	for mapKey, mapInterface := range mapElements {
		hoursElements := mapInterface.([]interface{})
		var periods [25][4]float64

		for hourKey, hourInterface := range hoursElements {
			tmpElements := hourInterface.([]interface{})

			for key, value := range tmpElements {
				periods[hourKey][key] = value.(float64)
			}
		}

		elements[mapKey] = periods
	}

	return elements
}

func ParseBodyMap(jsonFile *os.File) (map[string]interface{}, error) {
	var body map[string]interface{}

	tmpBody, ioErr := ioutil.ReadAll(jsonFile)

	if ioErr != nil {
		return nil, ioErr
	}

	if tmpBody == nil {
		return nil, errors.New("Body required2")
	}

	err := json.Unmarshal([]byte(tmpBody), &body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func ParseBodyArray(jsonFile *os.File) ([]map[string]interface{}, error) {
	var body []map[string]interface{}

	tmpBody, ioErr := ioutil.ReadAll(jsonFile)

	if ioErr != nil {
		return nil, ioErr
	}

	if tmpBody == nil {
		return nil, errors.New("Body required")
	}

	err := json.Unmarshal([]byte(tmpBody), &body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
