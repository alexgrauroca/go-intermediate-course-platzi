package optimizer

func GetPowerCombinations(currentPowers [6]float64, calcType string) [][6]float64 {
	var combinations [][6]float64
	var powers [6]float64

	combinePowers(1, currentPowers, &powers, &combinations, calcType)

	return combinations
}

func combinePowers(period int, currentPowers [6]float64, powers *[6]float64, combinations *[][6]float64, calcType string) {
	cpPeriod := period - 1
	prevKey := period - 2
	prevPower := 0.0
	newPower := currentPowers[cpPeriod] + 1

	if prevKey >= 0 {
		prevPower = powers[prevKey]
	}

	// +P
	if newPower >= prevPower {
		powers[period-1] = newPower

		if period < 6 {
			combinePowers(period+1, currentPowers, powers, combinations, calcType)
		} else {
			var combination [6]float64

			for key, value := range powers {
				combination[key] = value
			}

			*combinations = append(*combinations, combination)
		}
	}

	newPower = currentPowers[cpPeriod]

	// P
	if newPower >= prevPower {
		powers[period-1] = newPower

		if period < 6 {
			combinePowers(period+1, currentPowers, powers, combinations, calcType)
		} else {
			var combination [6]float64

			for key, value := range powers {
				combination[key] = value
			}

			*combinations = append(*combinations, combination)
		}
	}

	newPower = currentPowers[cpPeriod] - 1

	// -P
	if newPower > 0 && newPower >= prevPower {
		powers[period-1] = newPower

		if period < 6 {
			combinePowers(period+1, currentPowers, powers, combinations, calcType)
		} else {
			var combination [6]float64

			for key, value := range powers {
				combination[key] = value
			}

			*combinations = append(*combinations, combination)
		}
	}
}
