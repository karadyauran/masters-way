// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: job_dones.query.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createJobDone = `-- name: CreateJobDone :one
INSERT INTO job_dones(
    created_at,
    updated_at,
    description,
    time,
    owner_uuid,
    day_report_uuid
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING uuid, created_at, updated_at, description, time, owner_uuid, day_report_uuid
`

type CreateJobDoneParams struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Description   string    `json:"description"`
	Time          int32     `json:"time"`
	OwnerUuid     uuid.UUID `json:"owner_uuid"`
	DayReportUuid uuid.UUID `json:"day_report_uuid"`
}

func (q *Queries) CreateJobDone(ctx context.Context, arg CreateJobDoneParams) (JobDone, error) {
	row := q.queryRow(ctx, q.createJobDoneStmt, createJobDone,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Description,
		arg.Time,
		arg.OwnerUuid,
		arg.DayReportUuid,
	)
	var i JobDone
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.Time,
		&i.OwnerUuid,
		&i.DayReportUuid,
	)
	return i, err
}

const deleteJobDone = `-- name: DeleteJobDone :exec
DELETE FROM job_dones
WHERE uuid = $1
`

func (q *Queries) DeleteJobDone(ctx context.Context, argUuid uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteJobDoneStmt, deleteJobDone, argUuid)
	return err
}

const getListJobsDoneByDayReportId = `-- name: GetListJobsDoneByDayReportId :many
SELECT uuid, created_at, updated_at, description, time, owner_uuid, day_report_uuid FROM job_dones
WHERE job_dones.day_report_uuid = $1
ORDER BY created_at
`

func (q *Queries) GetListJobsDoneByDayReportId(ctx context.Context, dayReportUuid uuid.UUID) ([]JobDone, error) {
	rows, err := q.query(ctx, q.getListJobsDoneByDayReportIdStmt, getListJobsDoneByDayReportId, dayReportUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []JobDone{}
	for rows.Next() {
		var i JobDone
		if err := rows.Scan(
			&i.Uuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.Time,
			&i.OwnerUuid,
			&i.DayReportUuid,
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

const updateJobDone = `-- name: UpdateJobDone :one
UPDATE job_dones
SET
updated_at = coalesce($1, updated_at),
description = coalesce($2, description),
time = coalesce($3, time)
WHERE uuid = $4
RETURNING uuid, created_at, updated_at, description, time, owner_uuid, day_report_uuid
`

type UpdateJobDoneParams struct {
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	Description sql.NullString `json:"description"`
	Time        sql.NullInt32  `json:"time"`
	Uuid        uuid.UUID      `json:"uuid"`
}

func (q *Queries) UpdateJobDone(ctx context.Context, arg UpdateJobDoneParams) (JobDone, error) {
	row := q.queryRow(ctx, q.updateJobDoneStmt, updateJobDone,
		arg.UpdatedAt,
		arg.Description,
		arg.Time,
		arg.Uuid,
	)
	var i JobDone
	err := row.Scan(
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.Time,
		&i.OwnerUuid,
		&i.DayReportUuid,
	)
	return i, err
}
