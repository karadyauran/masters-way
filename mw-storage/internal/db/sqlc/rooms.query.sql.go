// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: rooms.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (created_at, name, type)
VALUES ($1, $2, $3)
RETURNING uuid, name, type
`

type CreateRoomParams struct {
	CreatedAt pgtype.Timestamp `json:"created_at"`
	Name      pgtype.Text      `json:"name"`
	Type      RoomType         `json:"type"`
}

type CreateRoomRow struct {
	Uuid pgtype.UUID `json:"uuid"`
	Name pgtype.Text `json:"name"`
	Type RoomType    `json:"type"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (CreateRoomRow, error) {
	row := q.db.QueryRow(ctx, createRoom, arg.CreatedAt, arg.Name, arg.Type)
	var i CreateRoomRow
	err := row.Scan(&i.Uuid, &i.Name, &i.Type)
	return i, err
}

const getIsPrivateRoomAlreadyExists = `-- name: GetIsPrivateRoomAlreadyExists :one
SELECT EXISTS (
    SELECT 1
    FROM (
        SELECT DISTINCT room_uuid
        FROM users_rooms
        WHERE users_rooms.user_uuid = $1
    ) AS user1_rooms
    JOIN (
        SELECT DISTINCT room_uuid
        FROM users_rooms
        WHERE users_rooms.user_uuid = $2
    ) AS user2_rooms ON user1_rooms.room_uuid = user2_rooms.room_uuid
    JOIN rooms ON rooms.uuid = user1_rooms.room_uuid
    WHERE rooms.type = 'private'
) AS is_private_room_already_exists
`

type GetIsPrivateRoomAlreadyExistsParams struct {
	User1 pgtype.UUID `json:"user_1"`
	User2 pgtype.UUID `json:"user_2"`
}

func (q *Queries) GetIsPrivateRoomAlreadyExists(ctx context.Context, arg GetIsPrivateRoomAlreadyExistsParams) (bool, error) {
	row := q.db.QueryRow(ctx, getIsPrivateRoomAlreadyExists, arg.User1, arg.User2)
	var is_private_room_already_exists bool
	err := row.Scan(&is_private_room_already_exists)
	return is_private_room_already_exists, err
}

const getRoomByUUID = `-- name: GetRoomByUUID :one
SELECT
    rooms.uuid,
    rooms.name,
    rooms.type,
    users_rooms.is_room_blocked,
    ARRAY(
        SELECT
            users_rooms.user_uuid
        FROM users_rooms
        WHERE users_rooms.room_uuid = rooms.uuid
    )::UUID[] AS user_uuids,
    ARRAY(
            SELECT
                users_rooms.user_role
            FROM users_rooms
            WHERE users_rooms.room_uuid = rooms.uuid
    )::VARCHAR[] AS user_roles
FROM rooms
JOIN users_rooms ON rooms.uuid = users_rooms.room_uuid
WHERE rooms.uuid = $1 AND users_rooms.user_uuid = $2
`

type GetRoomByUUIDParams struct {
	RoomUuid pgtype.UUID `json:"room_uuid"`
	UserUuid pgtype.UUID `json:"user_uuid"`
}

type GetRoomByUUIDRow struct {
	Uuid          pgtype.UUID   `json:"uuid"`
	Name          pgtype.Text   `json:"name"`
	Type          RoomType      `json:"type"`
	IsRoomBlocked bool          `json:"is_room_blocked"`
	UserUuids     []pgtype.UUID `json:"user_uuids"`
	UserRoles     []string      `json:"user_roles"`
}

func (q *Queries) GetRoomByUUID(ctx context.Context, arg GetRoomByUUIDParams) (GetRoomByUUIDRow, error) {
	row := q.db.QueryRow(ctx, getRoomByUUID, arg.RoomUuid, arg.UserUuid)
	var i GetRoomByUUIDRow
	err := row.Scan(
		&i.Uuid,
		&i.Name,
		&i.Type,
		&i.IsRoomBlocked,
		&i.UserUuids,
		&i.UserRoles,
	)
	return i, err
}

const getRoomsByUserUUID = `-- name: GetRoomsByUserUUID :many
SELECT
    rooms.uuid,
    rooms.name,
    rooms.type,
    users_rooms.is_room_blocked,
    ARRAY(
        SELECT
            users_rooms.user_uuid
        FROM users_rooms
        WHERE users_rooms.room_uuid = rooms.uuid
    )::UUID[] AS user_uuids,
    ARRAY(
            SELECT
                users_rooms.user_role
            FROM users_rooms
            WHERE users_rooms.room_uuid = rooms.uuid
    )::VARCHAR[] AS user_roles
FROM rooms
JOIN users_rooms ON rooms.uuid = users_rooms.room_uuid
WHERE users_rooms.user_uuid = $1 AND rooms.type = $2
`

type GetRoomsByUserUUIDParams struct {
	UserUuid pgtype.UUID `json:"user_uuid"`
	RoomType RoomType    `json:"room_type"`
}

type GetRoomsByUserUUIDRow struct {
	Uuid          pgtype.UUID   `json:"uuid"`
	Name          pgtype.Text   `json:"name"`
	Type          RoomType      `json:"type"`
	IsRoomBlocked bool          `json:"is_room_blocked"`
	UserUuids     []pgtype.UUID `json:"user_uuids"`
	UserRoles     []string      `json:"user_roles"`
}

func (q *Queries) GetRoomsByUserUUID(ctx context.Context, arg GetRoomsByUserUUIDParams) ([]GetRoomsByUserUUIDRow, error) {
	rows, err := q.db.Query(ctx, getRoomsByUserUUID, arg.UserUuid, arg.RoomType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRoomsByUserUUIDRow{}
	for rows.Next() {
		var i GetRoomsByUserUUIDRow
		if err := rows.Scan(
			&i.Uuid,
			&i.Name,
			&i.Type,
			&i.IsRoomBlocked,
			&i.UserUuids,
			&i.UserRoles,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
