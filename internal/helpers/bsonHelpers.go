package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FieldFromM(m primitive.M, field string) string {

	result, ok := m[field]

	if ok {
		return result.(string)
	}

	return ""
}

func FieldFromSliceOfM(m []primitive.M, field string) []string {
	var result []string

	if len(m) == 0 {
		return result
	}

	for i := range m {

		row, ok := m[i][field]

		if ok {
			result = append(result, row.(string))
		}

	}

	return result
}
