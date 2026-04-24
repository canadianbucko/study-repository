package sql

import (
	"context"
	"yayca/dto"
)

var ctx context.Context

func ShoveThatIntoSQL(dto dto.EmployeeDTO) error {

	ctx = context.Background()

	conn, err := CreateConnection(ctx)
	if err != nil {
		return err
	}

	sqlQuery := `
	INSERT INTO employees (id, full_name, position)
	VALUES ($1, $2, $3);
	`

	_, err = conn.Exec(ctx, sqlQuery, dto.ID, dto.FullName, dto.Position)

	if err != nil {
		return err
	}
	return nil

}
