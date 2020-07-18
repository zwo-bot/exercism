// Package proverb doing rhyme stuff
package proverb

import (
	"fmt"
)

// Proverb doing some rhyme stuff
func Proverb(rhyme []string) []string {
	var result []string
	for i, v := range rhyme {
		if len(rhyme) > 1 && i < len(rhyme)-1 {
			result = append(result, fmt.Sprintf("For want of a %v the %v was lost.", v, rhyme[i+1]))
		}
		if i == len(rhyme)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
		}
	}
	return result
}
