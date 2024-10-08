package services

import (
	openapiMWChatWebSocket "mw-chat-bff/apiAutogenerated/mw-chat-websocket"
	"mw-chat-bff/internal/schemas"

	"github.com/gin-gonic/gin"
	lop "github.com/samber/lo/parallel"
)

type ChatWebSocketService struct {
	chatWebSocketAPI *openapiMWChatWebSocket.APIClient
}

func NewChatWebSocketService(mwChatWebSocketAPI *openapiMWChatWebSocket.APIClient) *ChatWebSocketService {
	return &ChatWebSocketService{mwChatWebSocketAPI}
}

func (cs *ChatWebSocketService) SendMessage(ctx *gin.Context, roomID string, messageResponse *schemas.SendMessagePayload) error {
	request := openapiMWChatWebSocket.SchemasSendMessagePayload{
		Message: openapiMWChatWebSocket.SchemasMessageResponse{
			MessageId:      messageResponse.Message.MessageID,
			Message:        messageResponse.Message.Message,
			MessageReaders: []openapiMWChatWebSocket.SchemasMessageReader{},
			OwnerId:        messageResponse.Message.OwnerID,
			OwnerImageUrl:  messageResponse.Message.OwnerImageURL,
			OwnerName:      messageResponse.Message.OwnerName,
			RoomId:         roomID,
		},
		Users: messageResponse.UserIDs,
	}
	_, err := cs.chatWebSocketAPI.SocketAPI.SendMessageEvent(ctx).Request(request).Execute()
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatWebSocketService) SendRoom(ctx *gin.Context, populatedRoom *schemas.RoomPopulatedResponse) error {
	openapiUsers := lop.Map(populatedRoom.Users, func(user schemas.UserResponse, _ int) openapiMWChatWebSocket.SchemasUserResponse {
		return openapiMWChatWebSocket.SchemasUserResponse{
			ImageUrl: user.ImageURL,
			Name:     user.Name,
			Role:     user.Role,
			UserId:   user.UserID,
		}
	})

	openapiMessages := lop.Map(populatedRoom.Messages, func(message schemas.MessageResponse, _ int) openapiMWChatWebSocket.SchemasMessageResponse {
		return openapiMWChatWebSocket.SchemasMessageResponse{
			MessageId:      message.MessageID,
			Message:        message.Message,
			MessageReaders: []openapiMWChatWebSocket.SchemasMessageReader{},
			OwnerId:        message.OwnerID,
			OwnerImageUrl:  message.OwnerImageURL,
			OwnerName:      message.OwnerName,
			RoomId:         populatedRoom.RoomID,
		}
	})

	request := openapiMWChatWebSocket.SchemasRoomPopulatedResponse{
		ImageUrl: populatedRoom.ImageURL,
		Messages: openapiMessages,
		Name:     populatedRoom.Name,
		RoomId:   populatedRoom.RoomID,
		RoomType: populatedRoom.RoomType,
		Users:    openapiUsers,
	}

	_, err := cs.chatWebSocketAPI.SocketAPI.SendRoomEvent(ctx).Request(request).Execute()
	if err != nil {
		return err
	}

	return nil
}
