package willyoumakeit

func ZeroFuel(distanceToPump, mpg, fuelLeft int) bool {
	return mpg * fuelLeft >= distanceToPump
}