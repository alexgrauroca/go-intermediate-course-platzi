package optimizer_maximeter_calculate_contracted_powers_test

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

var maximeterData = make(map[string]interface{})

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
	Reading the maximeter json and maximeter imports json. They contain the data which will be used for tests
*/
func initData() {
	maximeterJson, maximeterErr := os.Open("./../data/maximeter.json")
	defer maximeterJson.Close()

	if maximeterErr != nil {
		log.Fatal(maximeterErr)
	}

	maximeter, parseMaximeterErr := optimizer.ParseBodyMap(maximeterJson)

	if parseMaximeterErr != nil {
		log.Fatal(parseMaximeterErr)
	}

	maximeterData = maximeter
}

func TestCalculateOptimizedContractedPowers(t *testing.T) {
	response, httpStatus := optimizer.CalculateOptimizedContractedPowers(maximeterData)

	assert.Equal(t, 200, httpStatus)

	assert.Equal(t, 1056.79, response["newImport"].(float64))
	assert.Equal(t, 1471.92, response["currentImport"].(float64))

	assert.Equal(t, 1471.924410, optimizer.RoundNumber(response["currentAnnualContracted"].(map[string]interface{})["annual"].(float64), 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["currentAnnualExcess"].(map[string]interface{})["annual"].(float64), 6))

	assert.Equal(t, 40.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["1"].(float64), 6))
	assert.Equal(t, 50.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["2"].(float64), 6))
	assert.Equal(t, 50.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["3"].(float64), 6))
	assert.Equal(t, 50.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["4"].(float64), 6))
	assert.Equal(t, 50.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["5"].(float64), 6))
	assert.Equal(t, 50.00, optimizer.RoundNumber(response["currentContractedPowers"].(map[string]interface{})["6"].(float64), 6))

	assert.Equal(t, 1000.534177, optimizer.RoundNumber(response["newAnnualContracted"].(map[string]interface{})["annual"].(float64), 6))
	assert.Equal(t, 56.256000, optimizer.RoundNumber(response["newAnnualExcess"].(map[string]interface{})["annual"].(float64), 6))

	assert.Equal(t, 29.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["1"].(float64), 6))
	assert.Equal(t, 29.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["2"].(float64), 6))
	assert.Equal(t, 36.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["3"].(float64), 6))
	assert.Equal(t, 36.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["4"].(float64), 6))
	assert.Equal(t, 36.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["5"].(float64), 6))
	assert.Equal(t, 36.00, optimizer.RoundNumber(response["newPowers"].(map[string]interface{})["6"].(float64), 6))

	assert.Equal(t, true, response["optimized"].(bool))
}
