// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: liquidity_params.sql

package sqlc

import (
	"context"
)

const fetchLiquidityParams = `-- name: FetchLiquidityParams :one
SELECT params FROM liquidity_params WHERE id = 1
`

func (q *Queries) FetchLiquidityParams(ctx context.Context) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, fetchLiquidityParams)
	var params []byte
	err := row.Scan(&params)
	return params, err
}

const upsertLiquidityParams = `-- name: UpsertLiquidityParams :exec
INSERT INTO liquidity_params (
    id, params
) VALUES (
    1, $1
) ON CONFLICT (id) DO UPDATE SET
    params = excluded.params
`

func (q *Queries) UpsertLiquidityParams(ctx context.Context, params []byte) error {
	_, err := q.db.ExecContext(ctx, upsertLiquidityParams, params)
	return err
}