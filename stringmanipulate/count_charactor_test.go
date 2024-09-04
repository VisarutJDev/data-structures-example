package stringmanipulate

import (
	"fmt"
	"testing"
)

func TestCountFrequency(t *testing.T) {
	str := "hello world"
	freq := CountFrequency(str)
	for ch, count := range freq {
		fmt.Printf("Character '%c' occurs %d times\n", ch, count)
	}
}
