package main

import (
	"io/ioutil"
	"net/http"
	"time"
	// "encoding/json"
	// "fmt"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 依存パッケージの読み込み > グローバル定数 > グローバル変数 > init() > main() の順に実行（判定）される

// bitflyerのベースurlとエンドポイント
const bitflyerBaseUrl = "https://api.bitflyer.jp/v1/getticker?product_code="

var e = createMux()

func main() {
	// パスとarticleIndexを紐付けている
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":9000"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, GetBitcoinApi())
}

// 対応出来ていない
type BitflyerBitcoin struct {
	Bitcoin string `json:"product_code"`
	Time    time.Time `json:"timestamp"`
	Id      int `json:"tick_id"`
	BestBid int `json:"best_bid"`
	BestAsk int `json:"best_ask"`
	BestBidSize int `json:"best_bid_size"`
	BestAskSize int `json: "best_ask_size"`
	TotalBidDepth int `json:"total_bid_depth"`
	TotalAskDepth int `json:"total_ask_depth"`
	Ltp int `json:"ltp"`
	Volume int `json:"volume"`
	VolumeByProduct int `json:"volume_by_product"`
}

// BitFlyerのBitcoinのAPIを取得する関数
func GetBitcoinApi() string {

	// GetでWebAPIに対してアクセスする
	resp, err := http.Get(bitflyerBaseUrl + "btc_jpy")
	if err != nil {
		log.Fatal(err)
	}

	// 最後にapiをCloseする
	defer resp.Body.Close()

	// ReadAllは、エラーまたはEOFに達するまで読み込み、読み込んだデータを返す
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// 対応出来ていない
	// var bitcoin BitflyerBitcoin
	// if err := json.Unmarshal(body, &bitcoin); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(body)

	// byte配列をstring型へキャスト
	BitcoinString := string(body)

	return BitcoinString
}
