// Package dog build functionality between human attributes and dogs
package dog

import (
	"errors"
	"fmt"
)

const oneHumanYearEquivalentInDogYear = 7

// Convert get human years and convert it dog year
// humanYear int
// Return Error if humanYear less than or equal 0
func Convert(humanYear int) (int, error) {
	if humanYear <= 0 {
		errorMsg := fmt.Sprintf("Human year must be great 0 but you enter %d", humanYear)
		return 0, errors.New(errorMsg)
	}
	dogYear := humanYear * oneHumanYearEquivalentInDogYear
	return dogYear, nil
}
