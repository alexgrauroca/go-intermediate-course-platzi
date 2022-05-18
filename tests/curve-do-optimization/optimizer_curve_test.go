package optimizer_curve_do_optimization_test

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

func TestDoOptimization(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)

	newImport, selectedPowers := optimizer.DoOptimization(powerInf, 153585.364665)

	assert.Equal(t, 800.00, selectedPowers[0])
	assert.Equal(t, 800.00, selectedPowers[1])
	assert.Equal(t, 800.00, selectedPowers[2])
	assert.Equal(t, 1465.00, selectedPowers[3])
	assert.Equal(t, 1465.00, selectedPowers[4])
	assert.Equal(t, 1465.00, selectedPowers[5])

	assert.Equal(t, 63997.93, newImport)
}
