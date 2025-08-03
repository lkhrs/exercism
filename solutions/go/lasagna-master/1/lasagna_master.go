package lasagna

// Takes layers and prep time per layer and returns a total prep time
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Takes layers and returns total amounts of noodles and sauce
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// Takes two lists of ingredients and replaces the last ingredient on the second list with the last ingredient from the first list
func AddSecretIngredient(list, recipe []string) {
	recipe[len(recipe)-1] = list[len(list)-1]
}

// Takes a two-portion recipe and scales it for a given amount of portions
func ScaleRecipe(base []float64, portions int) (amounts []float64) {
	for i := 0; i < len(base); i++ {
		amounts = append(amounts, base[i]*0.5*float64(portions))
	}
	return
}
