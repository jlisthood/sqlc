// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: batch.go

package querytest

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4"
)

const getValues = `-- name: GetValues :batchmany
SELECT a, b
FROM myschema.foo
WHERE b = $1
`

type GetValuesBatchResults struct {
	br  pgx.BatchResults
	ind int
}

func (q *Queries) GetValues(ctx context.Context, b []sql.NullInt32) *GetValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range b {
		vals := []interface{}{
			a,
		}
		batch.Queue(getValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &GetValuesBatchResults{br, 0}
}

func (b *GetValuesBatchResults) Query(f func(int, []MyschemaFoo, error)) {
	for {
		rows, err := b.br.Query()
		if err != nil && (err.Error() == "no result" || err.Error() == "batch already closed") {
			break
		}
		defer rows.Close()
		var items []MyschemaFoo
		for rows.Next() {
			var i MyschemaFoo
			if err := rows.Scan(&i.A, &i.B); err != nil {
				break
			}
			items = append(items, i)
		}

		if f != nil {
			f(b.ind, items, rows.Err())
		}
		b.ind++
	}
}

func (b *GetValuesBatchResults) Close() error {
	return b.br.Close()
}

const insertValues = `-- name: InsertValues :batchone
INSERT INTO myschema.foo (a, b)
VALUES ($1, $2)
RETURNING a
`

type InsertValuesBatchResults struct {
	br  pgx.BatchResults
	ind int
}

type InsertValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) InsertValues(ctx context.Context, arg []InsertValuesParams) *InsertValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.A,
			a.B,
		}
		batch.Queue(insertValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertValuesBatchResults{br, 0}
}

func (b *InsertValuesBatchResults) QueryRow(f func(int, sql.NullString, error)) {
	for {
		row := b.br.QueryRow()
		var a sql.NullString
		err := row.Scan(&a)
		if err != nil && (err.Error() == "no result" || err.Error() == "batch already closed") {
			break
		}
		if f != nil {
			f(b.ind, a, err)
		}
		b.ind++
	}
}

func (b *InsertValuesBatchResults) Close() error {
	return b.br.Close()
}

const updateValues = `-- name: UpdateValues :batchexec
UPDATE myschema.foo SET a = $1, b = $2
`

type UpdateValuesBatchResults struct {
	br  pgx.BatchResults
	ind int
}

type UpdateValuesParams struct {
	A sql.NullString
	B sql.NullInt32
}

func (q *Queries) UpdateValues(ctx context.Context, arg []UpdateValuesParams) *UpdateValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.A,
			a.B,
		}
		batch.Queue(updateValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &UpdateValuesBatchResults{br, 0}
}

func (b *UpdateValuesBatchResults) Exec(f func(int, error)) {
	for {
		_, err := b.br.Exec()
		if err != nil && (err.Error() == "no result" || err.Error() == "batch already closed") {
			break
		}
		if f != nil {
			f(b.ind, err)
		}
		b.ind++
	}
}

func (b *UpdateValuesBatchResults) Close() error {
	return b.br.Close()
}
