// Package tiktokshopak
// @author: 杨斯杰
// @date: 2024/2/21 17:14

package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"os"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	appKey := os.Getenv("TTS_APP_KEY")
	if appKey == "" {
		e.Logger.Error("TTS_APP_KEY is not specified")
		os.Exit(1)
	}
	appSecret := os.Getenv("TTS_APP_SECRET")
	if appSecret == "" {
		e.Logger.Error("TTS_APP_SECRET is not specified")
		os.Exit(1)
	}

	// Routes
	e.GET("/auth", func(c echo.Context) error {
		authCode := c.QueryParam("code")
		if authCode == "" {
			return c.String(400, "code not found")
		}
		e.Logger.Info("code is " + authCode)

		state := c.QueryParam("state")
		e.Logger.Info("state is " + state)

		RequestAccessToken(e.Logger, authCode, appKey, appSecret)
		return c.String(200, "")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func RequestAccessToken(logger echo.Logger, authCode, appKey, appSecret string) {
	url := fmt.Sprintf("https://auth.tiktok-shops.com/api/v2/token/get?app_key=%s&app_secret=%s&auth_code=%s&grant_type=authorized_code", appKey, appSecret, authCode)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Error("build request error", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		logger.Error("request error", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Error("read body", err)
		return
	}
	logger.Print(string(body))
}
