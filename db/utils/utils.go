package utils

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

type Utils struct{}

func (u *Utils) StringToFloatSlice(s string, a []float64) []float64 {
	r := strings.Trim(s, "{}")
	if a == nil {
		a = make([]float64, 0, 10)
	}
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.ParseFloat(t, 64)
		a = append(a, i)
	}
	return a
}

func (u *Utils) StringToIntSlice(s string) []int {
	r := strings.Trim(s, "{}")
	a := make([]int, 0, 10)
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.Atoi(t)
		a = append(a, i)
	}
	return a
}

func (u *Utils) StringToStringSlice(s string) []string {
	r := strings.Trim(s, "{}")
	a := make([]string, 0, 10)
	for _, s := range strings.Split(r, ",") {
		a = append(a, s)
	}
	return a
}

func (u *Utils) StringSliceToString(sl []string) string {
	r := "{"
	total := len(sl)
	for i, s := range sl {
		r = r + s
		if i < (total - 1) {
			r = r + ","
		}
	}
	return r + "}"
}

func (u *Utils) ResultInterfaceToTime(result interface{}) time.Time {
	if result != nil {
		return result.(time.Time)
	}
	return time.Time{}
}

func (u *Utils) NullStringToTime(str sql.NullString) time.Time {
	if str.Valid {
		if str.String != "0000-00-00 00:00:00" {
			t, _ := time.Parse("2006-01-02 15:04:05", str.String)
			return t
		}
	}
	return time.Time{}
}

func (u *Utils) NullStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""

}

func (u *Utils) NullFloatToFloat(flt sql.NullFloat64) float64 {
	if flt.Valid {
		return flt.Float64
	}
	return 0

}

func (u *Utils) NullInt64ToInt64(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}
	return 0
}

func (u *Utils) NullBoolToBool(b sql.NullBool) bool {
	if b.Valid {
		return b.Bool
	}
	return false
}
