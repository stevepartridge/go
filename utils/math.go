package utils

import (
	"math"
	"strconv"
)

type _math struct{}

var Math = _math{}

func (_ _math) Round(x float64) float64 {
	v, frac := math.Modf(x)
	if x > 0.0 {
		if frac > 0.5 || (frac == 0.5 && uint64(v)%2 != 0) {
			v += 1.0
		}
	} else {
		if frac < -0.5 || (frac == -0.5 && uint64(v)%2 != 0) {
			v -= 1.0
		}
	}

	return v
}

func (_ _math) RoundFloat64(num float64, decimals int) float64 {
	str := strconv.FormatFloat(num, 'f', decimals, 64)
	num, _ = strconv.ParseFloat(str, 64)
	return num
}

func (_ _math) RoundFloat64ToString(num float64, decimals int) string {
	flt := Math.RoundFloat64(num, decimals)
	return strconv.FormatFloat(flt, 'f', decimals, 64)
}

func (_ _math) RoundFloat64ToInt(num float64) int {
	str := Math.RoundFloat64ToString(num*100.0, 2)
	flt, _ := strconv.ParseFloat(str, 64)
	return int(flt)
}

func (_ _math) IntCentsToFloat64(cents int, decimals int) float64 {
	return Math.RoundFloat64(float64(cents)/100.0, decimals)
}

func (_ _math) IntCentsToFloatString(cents int, decimals int) string {
	flt := Math.IntCentsToFloat64(cents, decimals)
	return strconv.FormatFloat(flt, 'f', decimals, 64)
}
