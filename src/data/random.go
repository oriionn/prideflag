package data

import "prideflag.fun/src/utils"

func GetChoices() ([4]Data, error) {
	data := [4]Data{}
	length := int64(len(DATASET))

	for i := range data {
		n, err := utils.RandomInt(length)
		if err != nil {
			return data, err
		}
		data[i] = DATASET[n]
	}

	return data, nil
}
