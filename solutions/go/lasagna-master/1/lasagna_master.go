package lasagnamaster

func PreparationTime(layers []string, m int) int {
	if m == 0 {
		m = 2
	}
	return m * len(layers)
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := range layers {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, sum int) []float64 {
	qt := make([]float64, len(quantities))
	copy(qt, quantities)
	for i := range quantities {
		qt[i] = qt[i] * float64(sum) / 2
	}
	return qt
}