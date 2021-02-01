package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChars(t *testing.T) {
	input := []string{"AGCGA", "AGC", "TTA", "AGA", "CCC", "TC"}
	var result = CharacterValidation(input)
	assert.Equal(t, true, result)

}
