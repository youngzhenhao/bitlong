// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: invoice_events.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const deleteInvoiceEvents = `-- name: DeleteInvoiceEvents :exec
DELETE
FROM invoice_events
WHERE invoice_id = $1
`

func (q *Queries) DeleteInvoiceEvents(ctx context.Context, invoiceID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInvoiceEvents, invoiceID)
	return err
}

const insertInvoiceEvent = `-- name: InsertInvoiceEvent :exec
INSERT INTO invoice_events (
    created_at, invoice_id, htlc_id, set_id, event_type, event_metadata
) VALUES (
    $1, $2, $3, $4, $5, $6
)
`

type InsertInvoiceEventParams struct {
	CreatedAt     time.Time
	InvoiceID     int32
	HtlcID        sql.NullInt64
	SetID         []byte
	EventType     int32
	EventMetadata []byte
}

func (q *Queries) InsertInvoiceEvent(ctx context.Context, arg InsertInvoiceEventParams) error {
	_, err := q.db.ExecContext(ctx, insertInvoiceEvent,
		arg.CreatedAt,
		arg.InvoiceID,
		arg.HtlcID,
		arg.SetID,
		arg.EventType,
		arg.EventMetadata,
	)
	return err
}

const selectInvoiceEvents = `-- name: SelectInvoiceEvents :many
SELECT id, created_at, invoice_id, htlc_id, set_id, event_type, event_metadata
FROM invoice_events
WHERE (
    invoice_id = $1 OR 
    $1 IS NULL
) AND (
    htlc_id = $2 OR 
    $2 IS NULL
) AND (
    set_id = $3 OR 
    $3 IS NULL
) AND (
    event_type = $4 OR 
    $4 IS NULL
) AND (
    created_at >= $5 OR 
    $5 IS NULL
) AND (
    created_at <= $6 OR 
    $6 IS NULL
) 
LIMIT $8 OFFSET $7
`

type SelectInvoiceEventsParams struct {
	InvoiceID     sql.NullInt32
	HtlcID        sql.NullInt64
	SetID         []byte
	EventType     sql.NullInt32
	CreatedAfter  sql.NullTime
	CreatedBefore sql.NullTime
	NumOffset     int32
	NumLimit      int32
}

func (q *Queries) SelectInvoiceEvents(ctx context.Context, arg SelectInvoiceEventsParams) ([]InvoiceEvent, error) {
	rows, err := q.db.QueryContext(ctx, selectInvoiceEvents,
		arg.InvoiceID,
		arg.HtlcID,
		arg.SetID,
		arg.EventType,
		arg.CreatedAfter,
		arg.CreatedBefore,
		arg.NumOffset,
		arg.NumLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InvoiceEvent
	for rows.Next() {
		var i InvoiceEvent
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.InvoiceID,
			&i.HtlcID,
			&i.SetID,
			&i.EventType,
			&i.EventMetadata,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
