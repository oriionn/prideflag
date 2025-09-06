package data

import (
	"slices"

	"prideflag.fun/src/utils"
)

func GetChoices() ([4]Data, error) {
	data := [4]Data{}
	names := []string{}
	length := int64(len(DATASET))

	for i := range data {
		choice, err := GetChoice(names, length)
		if err != nil {
			return data, err
		}

		data[i] = choice
		names = append(names, choice.Name)
	}

	return data, nil
}

func GetChoice(names []string, length int64) (Data, error) {
	n, err := utils.RandomInt(length)
	if err != nil {
		return Data{}, err
	}
	choice := DATASET[n]

	if slices.Contains(names, choice.Name) {
		return GetChoice(names, length)
	}

	return choice, nil
}
