package webapi

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"messenger_v2/webapi/models"
)

type Webapi interface {
	registerHandler() gin.HandlerFunc
	loginHandler() gin.HandlerFunc
	createChat() gin.HandlerFunc
	run()
}
type webapi struct {
	log *log.Logger
	db  *mongo.Database
}

func (w webapi) registerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.User
		c.BindJSON(&user)
		c.JSON(200, &user)
	}
}

func (w webapi) loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.UserLogin
		c.BindJSON(&user)
		c.JSON(200, &user)
	}
}

func (w webapi) createChat() gin.HandlerFunc {
	return func(c *gin.Context) {
		var create *models.CreateChat
		c.BindJSON(&create)
		c.JSON(200, &create)
	}
}

func (w webapi) run() {
	router := gin.Default()

	auth := router.Group("")
	auth.POST("sign-up", w.registerHandler())
	auth.POST("sign-in", w.loginHandler())

	chat := router.Group("/chat")
	chat.POST("/add-chat", nil)

	message := router.Group("/message")
	message.POST("/add-message", nil)

	group := router.Group("/group")
	group.POST("/add-chat", nil)

	err := router.Run("8080")
	if err != nil {
		log.Fatal(err)
	}
}
func mneLen(db *mongo.Database, logger *log.Logger) Webapi {
	return webapi{
		log: logger,
		db:  db,
	}
}
