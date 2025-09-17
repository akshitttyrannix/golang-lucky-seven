package helper

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FloatToDecimal128(val float64) (primitive.Decimal128, error) {
	strVal := strconv.FormatFloat(val, 'f', -1, 64) // precise string
	return primitive.ParseDecimal128(strVal)
}

func StringToDecimal128(val string) (primitive.Decimal128, error) {
	return primitive.ParseDecimal128(val)
}
