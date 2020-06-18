package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echoStart()

func echoStart() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// `/` というパス（URL）と `articleIndex` という処理を結びつける
	e.GET("/", handler.articleIndexAPI)

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}
