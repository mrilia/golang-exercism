package lasagna
import "slices"
// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg int) int{
    if avg == 0{
        avg = 2
    }

    return len(layers) * avg
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodleQty := 0
    sauceQty := 0.0
    for i:=0; i<len(layers); i++{
        switch layers[i]{
            case "noodles":
				noodleQty += 50
            case "sauce":
            	sauceQty += 0.2
        }
    }

    return noodleQty, sauceQty
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(hisIngredients []string, myIngredients []string) {
    myIngredients[len(myIngredients)-1] = hisIngredients[len(hisIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64{
    scaledAmounts := slices.Clone(amounts)
    if portions == 2{
        return scaledAmounts
    }

    scaleFactor := float64(portions) / 2
    for i:=0; i<len(amounts); i++{
        scaledAmounts[i] *= scaleFactor
    }

    return scaledAmounts
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
