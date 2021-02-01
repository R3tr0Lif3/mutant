package app

import (
	"errors"
	"strings"

	"mutant/db"

	"mutant/models"

	"mutant/utils"
)

//=============================================================================

// Is Mutant

//=============================================================================

// @Diego Guerrero

//=============================================================================

/**

 * The principal router contains all logic here.

 *

 * @CharacterValidation

 * @Dynamic Array

 * @Find Mutan

 * @Call Database

 */

//IsMutantLogic validate is mutant
func IsMutantLogic(mutant *models.MutantModel) (bool, error) {

	if !utils.CharacterValidation(mutant.Dna) {

		return false, errors.New("invalid character(s) found please check only characters allowed (A,T,C,G) ")

	}

	exist, err := db.DNAExist(mutant.Dna)
	if err != nil {
		return false, err
	}

	if exist.Verified == true {

		return bool(exist.IsMutannt), nil

	}

	// It has the value of the highest chain of the DNA sequence to build the dynamic array

	rows := 0

	columns := 0

	input := mutant.Dna

	for i := 0; i < len(input); i++ {

		if len(input[i]) > columns {

			columns = len(input[i])

		}

	}

	rows = len(input)

	if rows < columns {

		rows = columns

	} else {

		columns = rows

	}

	var matrix = make([][]string, rows)

	for i := range matrix {

		matrix[i] = make([]string, columns)

	}

	// fill dynamic matrix nxn with the DNA sequence

	for r := 0; r < len(input); r++ {

		array := strings.Split(input[r], "")

		for c := 0; c < len(array); c++ {

			matrix[r][c] = array[c]

		}

	}

	//call all functions to traverse the array and find the 4-letter sequence

	totalLR := goLeftRight(matrix)

	totalTB := goTopButtom(matrix)

	totalAB := goDiagonalAB(matrix)

	totalBA := goDiagonalBA(matrix)

	totalXY := goDiagonalXY(matrix)

	totalYX := goDiagonalYX(matrix)

	if totalLR+totalTB+totalAB+totalBA+totalXY+totalYX >= 2 {

		db.SaveDNA(mutant.Dna, true)
		return true, nil

	}

	db.SaveDNA(mutant.Dna, false)
	return false, nil

}

func goLeftRight(array [][]string) (total int) {

	var values []string

	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array); j++ {

			if array[i][j] != "" {

				values = append(values, array[i][j])

			}

		}

		values = append(values, "")

	}

	return countDnaSecuence(values)

}

func goTopButtom(array [][]string) (total int) {

	var values []string

	for j := 0; j < len(array); j++ {

		for i := 0; i < len(array); i++ {

			if array[i][j] != "" {

				values = append(values, array[i][j])

			}

		}

		values = append(values, "")

	}

	return countDnaSecuence(values)

}

func goDiagonalAB(array [][]string) (total int) {

	var values []string

	loop := len(array)

	for k := 1; k < len(array); k++ {

		loop--

		j := loop

		for i := 0; i <= loop; i++ {

			if array[i][j] != "" {

				values = append(values, array[i][j])

			}

			j--

		}

		values = append(values, "")

	}

	return countDnaSecuence(values)

}

func goDiagonalBA(array [][]string) (total int) {

	var diagonalValues []string

	loop := len(array)

	for k := 0; k < len(array); k++ {

		loop--

		i := len(array) - 1

		for j := k; j < len(array); j++ {

			if array[i][j] != "" {

				diagonalValues = append(diagonalValues, array[i][j])

			}

			i--

		}

		diagonalValues = append(diagonalValues, "")

	}

	return countDnaSecuence(diagonalValues)

}

func goDiagonalXY(array [][]string) (total int) {

	var diagonalValues []string

	loop := len(array)

	for k := 0; k < len(array); k++ {

		loop--

		i := len(array) - 1

		j := loop

		for l := 0; l <= loop; l++ {

			if array[i][j] != "" {

				diagonalValues = append(diagonalValues, array[i][j])

			}

			i--

			j--

		}

		diagonalValues = append(diagonalValues, "")

	}

	return countDnaSecuence(diagonalValues)

}

func goDiagonalYX(array [][]string) (total int) {

	var diagonalValues []string

	loop := len(array) - 1

	for k := 0; k < len(array); k++ {

		loop--

		j := len(array) - 1

		i := loop

		for l := 0; l <= loop; l++ {

			if array[i][j] != "" {

				diagonalValues = append(diagonalValues, array[i][j])

			}

			i--

			j--

		}

		diagonalValues = append(diagonalValues, "")

	}

	return countDnaSecuence(diagonalValues)

}

func countDnaSecuence(dnaSection []string) (totalFound int) {

	count := 0

	totalMutant := 0

	for left := 1; left < len(dnaSection); left++ {

		if dnaSection[left-1] == dnaSection[left] && dnaSection[left] != "" && dnaSection[left-1] != "" {

			count++

		} else {

			if count >= 3 {

				totalMutant++

				count = 0

			}

			count = 0

		}

	}

	return totalMutant

}
