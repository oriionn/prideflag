package data

import (
	"slices"

	"prideflag.fun/src/utils"
)

func GetChoices(flagsExists []string) ([4]Data, error) {
	data := [4]Data{}
	names := []string{}
	length := int64(len(DATASET))

	for i := range data {
		choice, err := GetChoice(names, length, flagsExists, 10)
		if err != nil {
			return data, err
		}

		data[i] = choice
		names = append(names, choice.Name)
	}

	return data, nil
}

func GetChoice(names []string, length int64, flagsExists []string, maxAttempts int) (Data, error) {
	n, err := utils.RandomInt(length)
	if err != nil {
		return Data{}, err
	}
	choice := DATASET[n]

	if maxAttempts <= 0 {
		return choice, nil
	}

	if slices.Contains(names, choice.Name) || slices.Contains(flagsExists, choice.File) {
		return GetChoice(names, length, flagsExists, maxAttempts - 1)
	}

	return choice, nil
}
