package optimizer_helpers_test

/*
I had to split the tests because with the maximeters process, if I run all the tests in a single package, they fails and I don't know why,
it feels like Go is saving in memory the vars and it doesn't work well with gorutines, but the curve tests works... :/
*/
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

func TestSelectExcessPrice(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	response := optimizer.SelectExcessPrice(powerInf.ExcessPrices, "2021-06-30")

	assert.Equal(t, 1.4064, optimizer.RoundNumber(response, 6))
}

func TestSelectKps(t *testing.T) {
	powerInf := optimizer.InitPowerInf(curveData)
	response := optimizer.SelectKp(powerInf.Kps, "2022-01-01")

	assert.Equal(t, 1.0, optimizer.RoundNumber(response["1"].(float64), 6))
	assert.Equal(t, 1.0, optimizer.RoundNumber(response["2"].(float64), 6))
	assert.Equal(t, 0.542746, optimizer.RoundNumber(response["3"].(float64), 6))
	assert.Equal(t, 0.41026, optimizer.RoundNumber(response["4"].(float64), 6))
	assert.Equal(t, 0.026371, optimizer.RoundNumber(response["5"].(float64), 6))
	assert.Equal(t, 0.026371, optimizer.RoundNumber(response["6"].(float64), 6))
	assert.Equal(t, "2022-04-30", response["end"].(string))
	assert.Equal(t, "2021-05-01", response["init"].(string))
}
