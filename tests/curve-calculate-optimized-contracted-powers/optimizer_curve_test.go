package optimizer_curve_calculate_contracted_powers_test

import (
	"log"
	"os"
	"power-optimizer/src/optimizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

var curveData = make(map[string]interface{})

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

	curve, parseCurveErr := optimizer.ParseBodyMap(curveJson)

	if parseCurveErr != nil {
		log.Fatal(parseCurveErr)
	}

	curveData = curve
}

func TestCalculateOptimizedContractedPowers(t *testing.T) {
	response, httpStatus := optimizer.CalculateOptimizedContractedPowers(curveData)

	assert.Equal(t, 200, httpStatus)

	assert.Equal(t, 63997.93, response["newImport"].(float64))
	assert.Equal(t, 64541.1, response["currentImport"].(float64))

	assert.Equal(t, 61937.40480, optimizer.RoundNumber(response["currentAnnualContracted"].(map[string]interface{})["annual"].(float64), 6))
	assert.Equal(t, 2603.695978, optimizer.RoundNumber(response["currentAnnualExcess"].(map[string]interface{})["annual"].(float64), 6))

	assert.Equal(t, 810.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[0].(float64), 6))
	assert.Equal(t, 810.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[1].(float64), 6))
	assert.Equal(t, 810.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[2].(float64), 6))
	assert.Equal(t, 1470.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[3].(float64), 6))
	assert.Equal(t, 1470.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[4].(float64), 6))
	assert.Equal(t, 1470.00, optimizer.RoundNumber(response["currentContractedPowers"].([]interface{})[5].(float64), 6))

	assert.Equal(t, 61333.58165, optimizer.RoundNumber(response["newAnnualContracted"].(map[string]interface{})["annual"].(float64), 6))
	assert.Equal(t, 2664.347956, optimizer.RoundNumber(response["newAnnualExcess"].(map[string]interface{})["annual"].(float64), 6))

	assert.Equal(t, 800.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[0], 6))
	assert.Equal(t, 800.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[1], 6))
	assert.Equal(t, 800.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[2], 6))
	assert.Equal(t, 1465.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[3], 6))
	assert.Equal(t, 1465.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[4], 6))
	assert.Equal(t, 1465.00, optimizer.RoundNumber(response["newPowers"].([6]float64)[5], 6))

	assert.Equal(t, true, response["optimized"].(bool))
}
