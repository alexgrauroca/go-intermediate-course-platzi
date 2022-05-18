package optimizer_maximeter_test

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
var maximeterImportsData []map[string]interface{}

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

	maximeterImportsJson, maximeterImportsErr := os.Open("./../data/maximeter-imports.json")
	defer maximeterImportsJson.Close()

	if maximeterImportsErr != nil {
		log.Fatal(maximeterImportsErr)
	}

	maximeter, parseMaximeterErr := optimizer.ParseBodyMap(maximeterJson)

	if parseMaximeterErr != nil {
		log.Fatal(parseMaximeterErr)
	}

	maximeterImports, parseMaximeterImportsErr := optimizer.ParseBodyArray(maximeterImportsJson)

	if parseMaximeterImportsErr != nil {
		log.Fatal(parseMaximeterImportsErr)
	}

	maximeterData = maximeter
	maximeterImportsData = maximeterImports
}

func TestCalculateCurrentAnnualContracted(t *testing.T) {
	powerInf := optimizer.InitPowerInf(maximeterData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualContracted(splitByPeriods, powerInf.ContractedPowers, powerInf.PowerPrices, powerInf.TariffPeriods)

	assert.Equal(t, 64.428000, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["1"], 6))
	assert.Equal(t, 56.638500, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["2"], 6))
	assert.Equal(t, 28.789500, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["3"], 6))
	assert.Equal(t, 25.093500, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["4"], 6))
	assert.Equal(t, 18.079500, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["5"], 6))
	assert.Equal(t, 10.837500, optimizer.RoundNumber(response["2021-06"].(map[string]float64)["6"], 6))

	assert.Equal(t, 66.575600, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["1"], 6))
	assert.Equal(t, 58.526450, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["2"], 6))
	assert.Equal(t, 29.749150, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["3"], 6))
	assert.Equal(t, 25.929950, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["4"], 6))
	assert.Equal(t, 18.682150, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["5"], 6))
	assert.Equal(t, 11.198750, optimizer.RoundNumber(response["2021-07"].(map[string]float64)["6"], 6))

	assert.Equal(t, 51.152640, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["1"], 6))
	assert.Equal(t, 46.960200, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["2"], 6))
	assert.Equal(t, 22.761200, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["3"], 6))
	assert.Equal(t, 19.363400, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["4"], 6))
	assert.Equal(t, 12.920600, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["5"], 6))
	assert.Equal(t, 8.254400, optimizer.RoundNumber(response["2022-02"].(map[string]float64)["6"], 6))

	assert.Equal(t, 1471.924410, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualExcessByMaximeters(t *testing.T) {
	powerInf := optimizer.InitPowerInf(maximeterData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualExcessByMaximeters(splitByPeriods, powerInf.ExcessPrices, powerInf.ContractedPowers, powerInf.Maximeters)

	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["1"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["2"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["3"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["4"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["5"], 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(response["2021-06-30"].(map[string]float64)["6"], 6))

	assert.Equal(t, 0.0, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualExcess(t *testing.T) {
	powerInf := optimizer.InitPowerInf(maximeterData)
	splitByPeriods := true
	response := optimizer.CalculateAnnualExcess(splitByPeriods, powerInf)

	assert.Equal(t, 0.0, optimizer.RoundNumber(response["annual"].(float64), 6))
}

func TestCalculateCurrentAnnualPower(t *testing.T) {
	powerInf := optimizer.InitPowerInf(maximeterData)
	splitByPeriods := true
	totalImport, annualContracted, annualExcess := optimizer.CalculateAnnualPower(splitByPeriods, powerInf)

	assert.Equal(t, 1471.924410, optimizer.RoundNumber(annualContracted["annual"].(float64), 6))
	assert.Equal(t, 0.0, optimizer.RoundNumber(annualExcess["annual"].(float64), 6))
	assert.Equal(t, 1471.92, totalImport)
}

func TestGetPowerCombinations(t *testing.T) {
	powerInf := optimizer.InitPowerInf(maximeterData)
	response := optimizer.GetPowerCombinations(powerInf.ContractedPowers, powerInf.CalcType)

	assert.Equal(t, 63, len(response))

	assert.Equal(t, 41.00, response[0][0])
	assert.Equal(t, 51.00, response[0][1])
	assert.Equal(t, 51.00, response[0][2])
	assert.Equal(t, 51.00, response[0][3])
	assert.Equal(t, 51.00, response[0][4])
	assert.Equal(t, 51.00, response[0][5])

	assert.Equal(t, 41.00, response[2][0])
	assert.Equal(t, 50.00, response[2][1])
	assert.Equal(t, 50.00, response[2][2])
	assert.Equal(t, 51.00, response[2][3])
	assert.Equal(t, 51.00, response[2][4])
	assert.Equal(t, 51.00, response[2][5])

	assert.Equal(t, 40.00, response[23][0])
	assert.Equal(t, 50.00, response[23][1])
	assert.Equal(t, 50.00, response[23][2])
	assert.Equal(t, 51.00, response[23][3])
	assert.Equal(t, 51.00, response[23][4])
	assert.Equal(t, 51.00, response[23][5])

	assert.Equal(t, 39.00, response[61][0])
	assert.Equal(t, 49.00, response[61][1])
	assert.Equal(t, 49.00, response[61][2])
	assert.Equal(t, 49.00, response[61][3])
	assert.Equal(t, 49.00, response[61][4])
	assert.Equal(t, 50.00, response[61][5])
}

func TestSelectBestImport(t *testing.T) {
	/*
		This test was created with a source info different than the current. I decided to keep this info, because maximeter.json and maximeter-imports.json are
		independent.
	*/
	newImport, selectedPowers := optimizer.SelectBestImport(maximeterImportsData, 153585.364665)

	assert.Equal(t, 39.00, selectedPowers[0])
	assert.Equal(t, 49.00, selectedPowers[1])
	assert.Equal(t, 49.00, selectedPowers[2])
	assert.Equal(t, 49.00, selectedPowers[3])
	assert.Equal(t, 49.00, selectedPowers[4])
	assert.Equal(t, 49.00, selectedPowers[5])

	assert.Equal(t, 1440.15, newImport)
}
