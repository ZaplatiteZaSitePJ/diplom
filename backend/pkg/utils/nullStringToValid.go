package utils

import "database/sql"

func NullStringToValid(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}

	return ns.String
}

func NullStringSliceToValid(ns []sql.NullString) []string {
	strs := make([]string, 0, 5)

	for _, n := range ns {
		if !n.Valid {
			break
		} else {
			strs = append(strs, n.String)
		}
	}

	return strs
}