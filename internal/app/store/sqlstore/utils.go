package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
)

func deleteByFields[T string](
	fields map[string]T,
	table string,
	s *sql.DB,
) error {
	query := "DELETE FROM " + table + " WHERE "

	for field, value := range fields {
		str := fmt.Sprintf("%v", value)
		query += field + "=" + `'` + str + `'` + " AND "
	}
	query = strings.TrimSuffix(query, " AND ")
	if err := s.QueryRow(
		query,
	).Err(); err != nil {
		return err
	}
	return nil
}
