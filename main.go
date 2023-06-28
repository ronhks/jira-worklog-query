package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.GET("/", landing)
	router.GET("/dashboard", dashboard)
	router.POST("/login", login)
	err := router.Run("localhost:8080")
	if err != nil {
		ErrorLogger.Println(err)
	}
}

func landing(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func dashboard(context *gin.Context) {
	context.HTML(http.StatusOK, "dashboard.html", nil)
}

func login(context *gin.Context) {
	loggedIn := rand.Int()%2 == 0
	if loggedIn {
		context.Redirect(http.StatusFound, "/dashboard")
	} else {
		context.Redirect(http.StatusFound, "/")
	}
}
