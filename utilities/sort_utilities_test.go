package utilities

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	elements := []int{9, 8, 7, 6, 5}

	BubbleSort(elements)

	fmt.Println("********** Worst case **********")
	assert.NotNil(t, elements)              // elements is not nil when calling this func
	assert.EqualValues(t, 5, len(elements)) // should have 5 elements length

	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	elements := []int{5, 6, 7, 8, 9}

	BubbleSort(elements)

	fmt.Println("********** Best case **********")
	assert.NotNil(t, elements)              // elements is not nil when calling this func
	assert.EqualValues(t, 5, len(elements)) // should have 5 elements length

	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}

// Nil check
func TestBubbleSortNilSlice(t *testing.T) {
	// Execution
	BubbleSort(nil)
}

// creating the slice elements depending on the argument passed into this function
func getElements(number int) []int {
	result := make([]int, number)
	i := 0

	for j := number - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elements := getElements(5)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements)) // should have 5 elements length

	assert.EqualValues(t, 4, elements[0])
	assert.EqualValues(t, 3, elements[1])
	assert.EqualValues(t, 2, elements[2])
	assert.EqualValues(t, 1, elements[3])
	assert.EqualValues(t, 0, elements[4])
}

func BenchMarkBabbleSort10(b *testing.B) {
	elements := getElements(10)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

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
