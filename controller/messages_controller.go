package controller

import (
	"secretemail/services"

	"github.com/gin-gonic/gin"
)

type MessagesController struct {
	MsgService *services.MessagesService
}

func NewMessagesController(msgService *services.MessagesService) *MessagesController {
	return &MessagesController{MsgService: msgService}
}

func (MessagesController) HandleSendMessage(ginCtx *gin.Context) {

}

func (MessagesController) HandleGetMessages(ginCtx *gin.Context) {

}
