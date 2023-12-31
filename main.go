package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type AuthBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	godotenv.Load(".env")

	redisClient := connectRedis()

	r := gin.Default()

	r.POST("/signup", func(ctx *gin.Context) {
		var body AuthBody
		err := ctx.Bind(&body)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Malformed request body.")
			return
		}
		pwdHash := hashString(body.Password)
		err = redisClient.Set(ctx.Request.Context(), body.Email, pwdHash, 2*time.Hour).Err()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to set the key in database")
			return
		}
		ctx.String(http.StatusOK, "Successfully signed up!")
	})

	r.POST("/login", func(ctx *gin.Context) {
		var body AuthBody
		err := ctx.Bind(&body)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Malformed request body.")
			return
		}

		pwdHash := hashString(body.Password)
		dbHash, err := redisClient.Get(ctx.Request.Context(), body.Email).Result()

		if err == redis.Nil {
			ctx.String(http.StatusBadRequest, "The user does not exist")
		} else if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to get the value from database")
		} else if pwdHash == dbHash {
			ctx.String(http.StatusOK, "Login success!")
		} else {
			ctx.String(http.StatusForbidden, "Invalid credentials.")
		}
	})

	addr := env("LISTEN_ADDR", ":5000")

	log.Fatal(r.Run(addr))
}
