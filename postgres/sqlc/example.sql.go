// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: example.sql

package sqlc

import (
	"context"
)

const sampleQuery = `-- name: SampleQuery :one
SELECT
    1::int
`

func (q *Queries) SampleQuery(ctx context.Context) (int32, error) {
	row := q.db.QueryRow(ctx, sampleQuery)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}
