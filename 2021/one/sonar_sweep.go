package one

type (
	sonarSweep struct {
		Measurements []int
	}
)

func NewSonarSweep(puzzleInput []int) *sonarSweep {
	return &sonarSweep{Measurements: puzzleInput}
}

func (ss *sonarSweep) NumberOfMeasurementIncreased(measurements []int) (increasedNTime int) {
	if len(measurements) == 0 {
		return 0
	}

	prevMeasurement := measurements[0]
	for _, measurement := range measurements {
		if measurement > prevMeasurement {
			increasedNTime++
		}
		prevMeasurement = measurement
	}

	return
}

func (ss *sonarSweep) ThreeWindowMeasurements() (increasedNTime int) {
	threeMeasurementsResult := []int{}
	lastThreeIndex := 0

	for lastThreeIndex+2 < len(ss.Measurements) {
		firstIndex := lastThreeIndex
		secondIndex := lastThreeIndex + 1
		thirdIndex := lastThreeIndex + 2

		// Sum of three-measurements windows
		sum := ss.Measurements[firstIndex] + ss.Measurements[secondIndex] + ss.Measurements[thirdIndex]
		threeMeasurementsResult = append(threeMeasurementsResult, sum)

		lastThreeIndex++
	}

	return ss.NumberOfMeasurementIncreased(threeMeasurementsResult)
}
