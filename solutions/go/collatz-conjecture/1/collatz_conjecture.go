package collatzconjecture
import "errors"
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("zero is an error") 
	}
	if n == 1 {
		return 0, nil
	}

	if n%2 == 0 {
		steps, err := CollatzConjecture(n/2)
		if err != nil {
			return 0, err
		}
		return 1 + steps, nil
	}
    
	steps, err := CollatzConjecture(n*3+1)
	if err != nil {
		return 0, err
	}
	return 1 + steps, nil
}
