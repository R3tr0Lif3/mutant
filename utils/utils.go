package utils

import "strings"

//CharacterValidation check is the dna secuence contains characters allows
func CharacterValidation(dna []string) (isOk bool) {
	joinString := ""
	for _, d := range dna {
		joinString += d
	}
	values := strings.Split(joinString, "")
	for _, v := range values {
		switch v {
		case "A":
		case "T":
		case "C":
		case "G":
		default:
			return false

		}
	}
	return true
}
