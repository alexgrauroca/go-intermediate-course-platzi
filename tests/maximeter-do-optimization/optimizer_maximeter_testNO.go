package optimizer_maximeter_do_optimization_test

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

func TestDoOptimization(t *testing.T) {
	powerInf1 := optimizer.InitPowerInf(maximeterData)
	newImport, selectedPowers := optimizer.DoOptimization(powerInf1, 153585.364665)

	assert.Equal(t, 29.00, selectedPowers[0])
	assert.Equal(t, 29.00, selectedPowers[1])
	assert.Equal(t, 36.00, selectedPowers[2])
	assert.Equal(t, 36.00, selectedPowers[3])
	assert.Equal(t, 36.00, selectedPowers[4])
	assert.Equal(t, 36.00, selectedPowers[5])

	assert.Equal(t, 1056.79, newImport)
}
