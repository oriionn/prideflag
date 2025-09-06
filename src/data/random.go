package data

import (
	"slices"

	"prideflag.fun/src/utils"
)

func GetChoices(flagsExists []string) ([4]Data, int, error) {
	data := [4]Data{}
	names := []string{}
	length := int64(len(DATASET))

	trueChoice, err := utils.RandomInt(4)
	if err != nil {
		return data, 0, err
	}

	for i := range data {
		choice, err := GetChoice(names, length, flagsExists, 40, i == int(trueChoice))
		if err != nil {
			return data, 0, err
		}

		data[i] = choice
		names = append(names, choice.Name)
	}

	return data, int(trueChoice), nil
}

func GetChoice(names []string, length int64, flagsExists []string, maxAttempts int, isTrueChoice bool) (Data, error) {
	n, err := utils.RandomInt(length)
	if err != nil {
		return Data{}, err
	}
	choice := DATASET[n]

	if maxAttempts <= 0 {
		return choice, nil
	}

	if slices.Contains(names, choice.Name) || (slices.Contains(flagsExists, choice.File) && isTrueChoice) {
		return GetChoice(names, length, flagsExists, maxAttempts - 1, isTrueChoice)
	}

	return choice, nil
}
