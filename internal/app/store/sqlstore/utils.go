package sqlstore

import "database/sql"

func deleteByIdTableWithIdField(id int, table, field string, s *sql.DB) error {
	query := "DELETE FROM " + table + " WHERE " + field + "=$1"
	if err := s.QueryRow(
		query,
		id,
	).Err(); err != nil {
		return err
	}
	return nil
}
