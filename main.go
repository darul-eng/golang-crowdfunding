package main

import (
	"github.com/gin-gonic/gin"
	"golang-vue-nuxtjs-crowdfunding/handler"
	"golang-vue-nuxtjs-crowdfunding/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("users", userHandler.RegisterUser)
	api.POST("sessions", userHandler.Login)

	router.Run()

	//userInput := user.RegisterUserInput{
	//	Name:       "Tes Simpan dari service",
	//	Occupation: "softwate engineer",
	//	Email:      "contoh@gmail.com",
	//	Password:   "darul",
	//}
	//
	//userService.RegisterUser(userInput)

	//
	//fmt.Println("Connection to database success")
	//
	//var users []user.User
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println("===================")
	//}

	//router := gin.Default()
	//router.GET("/handler", handler)
	//router.Run()
}

//func handler(ctx *gin.Context) {
//	dsn := "root:@tcp(127.0.0.1:3306)/golang_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//
//	var users []user.User
//	db.Find(&users)
//
//	ctx.JSON(http.StatusOK, users)
//}
