package app

import (
	"mutant/models"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

//=============================================================================

// Cover Test

//=============================================================================

// @Diego Guerrero

//=============================================================================

/**

*Run in console the command:

go test -v ./app/... -cover

*/

func TestGoLeftRight(t *testing.T) {
	input := [][]string{
		{"A", "A", "A", "A"},
		{"_", "_", "_", "W"},
		{"_", "_", "_", "T"},
		{"T", "T", "T", "T"},
	}
	var result = goLeftRight(input)
	assert.Equal(t, 2, result)

}

func TestGoTopButton(t *testing.T) {
	input := [][]string{
		{"A", "B", "C", "D"},
		{"A", "D", "X", "Y"},
		{"A", "X", "Z", "O"},
		{"A", "R", "T", "T"},
	}
	var result = goTopButtom(input)
	assert.Equal(t, 1, result)

}

func TestGoAB(t *testing.T) {
	input := [][]string{
		{"U", "R", "T", "B", "T", "A"},
		{"A", "B", "B", "Y", "A", "A"},
		{"T", "B", "A", "A", "T", "A"},
		{"B", "A", "A", "O", "T", "A"},
		{"R", "R", "T", "T", "T", "A"},
		{"U", "R", "T", "T", "T", "A"},
	}
	var result = goDiagonalAB(input)
	assert.Equal(t, 2, result)

}
func TestGoBA(t *testing.T) {
	input := [][]string{
		{"A", "R", "T", "B", "T", "A"},
		{"A", "E", "B", "Y", "A", "T"},
		{"T", "B", "E", "A", "T", "T"},
		{"B", "A", "A", "T", "T", "A"},
		{"R", "R", "T", "T", "Z", "A"},
		{"U", "R", "T", "X", "Y", "A"},
	}
	var result = goDiagonalBA(input)
	assert.Equal(t, 3, result)

}

func TestGoXY(t *testing.T) {
	input := [][]string{
		{"T", "R", "T", "B", "T", "A"},
		{"T", "T", "B", "Y", "A", "T"},
		{"T", "T", "T", "A", "T", "T"},
		{"B", "T", "T", "T", "T", "A"},
		{"R", "R", "T", "T", "Z", "A"},
		{"U", "R", "T", "T", "Y", "A"},
	}
	var result = goDiagonalXY(input)
	assert.Equal(t, 3, result)

}
func TestGoYX(t *testing.T) {
	input := [][]string{
		{"T", "R", "T", "B", "T", "A"},
		{"T", "W", "T", "T", "A", "T"},
		{"T", "E", "T", "T", "T", "T"},
		{"B", "W", "R", "T", "T", "T"},
		{"R", "R", "W", "T", "Z", "T"},
		{"U", "R", "T", "T", "Y", "A"},
	}
	var result = goDiagonalYX(input)
	assert.Equal(t, 2, result)

}

func TestIsMutantLogic(t *testing.T) {
	mutant := models.MutantModel{Dna: []string{"AGCGA", "AGC", "TTA", "AGA", "CCC", "TC"}}
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {

		t.Error(err)

	}

	defer conn.Close()

	_, err = conn.Do("DEL", mutant.Dna)

	if err != nil {

		t.Error(err)

	}
	result, err := IsMutantLogic(&mutant)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, false, result)

}
