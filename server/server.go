package server

import (
	"secretemail/controller"
	"secretemail/db"
	"secretemail/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type server struct {
	router        *gin.Engine
	msgController *controller.MessagesController
	keyController *controller.KeysController
}

func New() (*server, error) {
	db, err := db.Initialize()
	if err != nil {
		return nil, err
	}

	keysService := services.NewKeysService(db)
	msgService := services.NewMessagesService(db)

	keysController := controller.NewKeysController(keysService)
	messagessController := controller.NewMessagesController(msgService)
	return &server{
		router:        gin.Default(),
		keyController: keysController,
		msgController: messagessController,
	}, nil

}

func (s *server) AttachEndpoints() *server {

	r := s.router

	r.POST("/generate_keys", s.keyController.HandleGenerateKeys)
	r.POST("/send_message", s.msgController.HandleSendMessage)
	r.GET("/messages", s.msgController.HandleGetMessages)

	// r.POST("/bid", s.controller.Bid)

	return s
}

func (s *server) Run() error {
	return s.router.Run(viper.GetString("port"))
}
