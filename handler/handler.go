package handler

import (
	"net/http"

	"coinlang-Backend/api"

	"github.com/labstack/echo"
)

func articleIndexAPI(c echo.Context) error {

	// ステータスコード 200 で、GetBitcoinAPI関数で取得した文字列をレスポンス
	return c.JSON(http.StatusOK, api.GetBitcoinAPI())
}
