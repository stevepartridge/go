package db

import (
	"errors"
	"github.com/stevepartridge/go/utils"
	"strconv"
	"strings"
)

func MysqlIntSliceConditions(operator, field string, ints []int) string {
	intsSql := ""
	if len(ints) > 0 {

		intsSql = " " + operator + " " + field + " "

		if len(ints) == 1 {
			intsSql = intsSql + " = " + strconv.Itoa(ints[0]) + " "
		}

		if len(ints) > 1 {
			intsSql = intsSql + ` IN (` + utils.IntSliceToString(ints, ",") + `) `
		}
	}
	return intsSql
}

// Returns query, error, exacts, wildcards
func MysqlBuildConditionalQueryFromStringSlice(table, column string, conditions ...string) (string, error, []interface{}, []interface{}) {

	query := ""
	operator := "OR"

	wildcards := make([]interface{}, 0)
	exacts := make([]interface{}, 0)
	nots := make([]interface{}, 0)

	if len(conditions) == 0 {
		return query, errors.New("Conditions is length 0"), exacts, wildcards
	}

	if conditions[0] == "OR" || conditions[0] == "AND" {
		operator = conditions[0]
		conditions = conditions[1:]
	}

	for _, condition := range conditions {
		condition = strings.TrimSpace(condition)
		if strings.Contains(condition, "*") {
			condition = strings.Replace(condition, "*", "%", -1)
			wildcards = append(wildcards, condition)
		} else {
			if strings.Index(condition, "!") == 0 {
				nots = append(nots, condition[1:])
			} else {
				exacts = append(exacts, condition)
			}
		}
	}

	if len(wildcards) > 0 {

		for i, _ := range wildcards {
			query = query + ` ` + table + `.` + column
			if strings.Index(wildcards[i].(string), "!") == 0 {
				wildcards[i] = wildcards[i].(string)[1:]
				query = query + ` NOT`
			}
			query = query + ` LIKE ? `
			if i < len(wildcards)-1 {
				if strings.Index(wildcards[i].(string), "&") == 0 {
					query = query + ` AND `
					wildcards[i] = wildcards[i].(string)[1:]
				} else {
					query = query + ` ` + operator + ` `
				}
			}
		}
		query = query + ` `

		if len(exacts) > 0 || len(nots) > 0 {
			query = query + ` ` + operator + ` `
		}
	}

	if len(exacts) > 0 {

		query = query + ` ` + table + `.` + column + ` IN (`
		for i, _ := range exacts {
			query = query + `?`
			if i < len(exacts)-1 {
				query = query + `,`
			}
		}
		query = query + `) `

		if len(nots) > 0 {
			query = query + ` ` + operator + ` `
		}
	}

	if len(nots) > 0 {

		query = query + ` ` + table + `.` + column + ` NOT IN (`
		for i, _ := range nots {
			query = query + `?`
			exacts = append(exacts, nots[i])
			if i < len(nots)-1 {
				query = query + `,`
			}
		}
		query = query + `) `
	}

	return query, nil, exacts, wildcards

}
