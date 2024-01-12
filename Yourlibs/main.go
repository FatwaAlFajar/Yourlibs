package main

import (
	"github.com/FatwaAlFajar/Yourlibs/controllers/auth"
	"github.com/FatwaAlFajar/Yourlibs/controllers/books"
	"github.com/FatwaAlFajar/Yourlibs/controllers/borrowing"
	"github.com/FatwaAlFajar/Yourlibs/controllers/profile"
	"github.com/FatwaAlFajar/Yourlibs/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// initialize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	r := gin.Default()

	// connect to the database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES
	r.POST("/anggota/tambah", profile.CreateMember)
	r.POST("/user/register", auth.RegisterUser)

	// r.POST("/admin/register", auth.RegisterAdmin)
	r.POST("/login", auth.Login)

	// profile
	r.GET("/profile/all", profile.GetAllProfile)
	// r.GET("/profile/:id", profile.GetProfile)
	// r.PUT("/profile/ubah/:id", profile.UpdateProfile)

	// book
	r.POST("/book/tambah", books.CreateBook)
	r.GET("/books/all", books.GetAllBooks)

	// borrow
	r.POST("/borrow/tambah", borrowing.CreateBorrow)
	r.GET("borrow/all", borrowing.GetAllBorrowing)

	// run the server
	r.Run(":9888")
}
