// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddUserToRoom(ctx context.Context, arg AddUserToRoomParams) (AddUserToRoomRow, error)
	CreateMessage(ctx context.Context, arg CreateMessageParams) (CreateMessageRow, error)
	CreateMessageStatus(ctx context.Context, arg CreateMessageStatusParams) error
	CreateRoom(ctx context.Context, arg CreateRoomParams) (CreateRoomRow, error)
	GetChatPreview(ctx context.Context, receiverUuid pgtype.UUID) (int64, error)
	GetIsPrivateRoomAlreadyExists(ctx context.Context, arg GetIsPrivateRoomAlreadyExistsParams) (bool, error)
	GetMessagesByRoomUUID(ctx context.Context, roomUuid pgtype.UUID) ([]GetMessagesByRoomUUIDRow, error)
	GetRoomByUUID(ctx context.Context, arg GetRoomByUUIDParams) (GetRoomByUUIDRow, error)
	GetRoomsByUserUUID(ctx context.Context, arg GetRoomsByUserUUIDParams) ([]GetRoomsByUserUUIDRow, error)
	GetUsersUUIDsInRoom(ctx context.Context, roomUuid pgtype.UUID) ([]pgtype.UUID, error)
	RegenerateDbData(ctx context.Context) error
	RemoveEverything(ctx context.Context) error
	SetAllRoomMessagesAsRead(ctx context.Context, arg SetAllRoomMessagesAsReadParams) error
	UpdateMessageStatus(ctx context.Context, arg UpdateMessageStatusParams) error
}

var _ Querier = (*Queries)(nil)
