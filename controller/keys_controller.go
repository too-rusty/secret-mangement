package controller

import (
	"secretemail/services"

	"github.com/gin-gonic/gin"
)

type KeysController struct {
	KeyService *services.KeysService
}

func NewKeysController(keyService *services.KeysService) *KeysController {
	return &KeysController{KeyService: keyService}
}

func (k KeysController) HandleGenerateKeys(ginCtx *gin.Context) {}
