// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: way_collections.query.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createWayCollection = `-- name: CreateWayCollection :one
INSERT INTO way_collections(
    owner_uuid,
    created_at,
    updated_at,
    name
) VALUES (
    $1, $2, $3, $4
) RETURNING uuid, owner_uuid, created_at, updated_at, name
`

type CreateWayCollectionParams struct {
	OwnerUuid uuid.UUID `json:"owner_uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (q *Queries) CreateWayCollection(ctx context.Context, arg CreateWayCollectionParams) (WayCollection, error) {
	row := q.queryRow(ctx, q.createWayCollectionStmt, createWayCollection,
		arg.OwnerUuid,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i WayCollection
	err := row.Scan(
		&i.Uuid,
		&i.OwnerUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const deleteWayCollection = `-- name: DeleteWayCollection :exec
DELETE FROM way_collections
WHERE uuid = $1
`

func (q *Queries) DeleteWayCollection(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteWayCollectionStmt, deleteWayCollection, argUuid)
	return err
}

const getListWayCollectionsByUserId = `-- name: GetListWayCollectionsByUserId :many
SELECT uuid, owner_uuid, created_at, updated_at, name FROM way_collections
WHERE way_collections.owner_uuid = $1
ORDER BY created_at
`

func (q *Queries) GetListWayCollectionsByUserId(ctx context.Context, ownerUuid uuid.UUID) ([]WayCollection, error) {
	rows, err := q.query(ctx, q.getListWayCollectionsByUserIdStmt, getListWayCollectionsByUserId, ownerUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []WayCollection{}
	for rows.Next() {
		var i WayCollection
		if err := rows.Scan(
			&i.Uuid,
			&i.OwnerUuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
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

const updateWayCollection = `-- name: UpdateWayCollection :one
UPDATE way_collections
SET
name = coalesce($1, name),
updated_at = coalesce($2, updated_at)
WHERE uuid = $3
RETURNING uuid, owner_uuid, created_at, updated_at, name
`

type UpdateWayCollectionParams struct {
	Name      sql.NullString `json:"name"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	Uuid      uuid.UUID      `json:"uuid"`
}

func (q *Queries) UpdateWayCollection(ctx context.Context, arg UpdateWayCollectionParams) (WayCollection, error) {
	row := q.queryRow(ctx, q.updateWayCollectionStmt, updateWayCollection, arg.Name, arg.UpdatedAt, arg.Uuid)
	var i WayCollection
	err := row.Scan(
		&i.Uuid,
		&i.OwnerUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
