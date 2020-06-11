package utilities

// BubbleSort -
func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		// sort in a descending way from ascending way
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
	//return elements
}
