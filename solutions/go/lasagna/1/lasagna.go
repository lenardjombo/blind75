package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    remainingOvenTime := OvenTime - actualMinutesInOven
	return remainingOvenTime
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    // Assuming each layer takes 2 minutes to prepare
    TimeTakenToPrepareOneLayer := 2
    preparationTime := numberOfLayers * TimeTakenToPrepareOneLayer
	return preparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    elapsedTime := PreparationTime(numberOfLayers) + actualMinutesInOven 
    return elapsedTime
}
