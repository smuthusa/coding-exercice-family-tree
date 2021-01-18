package util

import (
	"fmt"
	"geektrust/errors"
	"geektrust/model"
	"strings"
)

const MALE = "male"
const FEMALE = "female"

var LineBreak = "\n"

type Logger func(string)

func Println(str string) {
	fmt.Println(str)
}
func NoOpPrint(str string) {
}
func GetGenderEnumFromString(gender string) (model.Gender, error) {
	genderLowercase := strings.ToLower(gender)
	if genderLowercase == MALE {
		return model.Male, nil
	} else if genderLowercase == FEMALE {
		return model.Female, nil
	}
	return -1, errors.InvalidGender
}