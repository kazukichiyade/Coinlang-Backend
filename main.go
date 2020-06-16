package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 依存パッケージの読み込み > グローバル定数 > グローバル変数 > init() > main() の順に実行（判定）される

// bitflyerのベースurlとエンドポイント
const bitflyerBaseURL = "https://api.bitflyer.jp/v1/getticker?product_code="

// BitflyerCoin is Bitflyerにあるコインの構造体 (jsonのkeyの表記を変更できるかつkeyと一致しないvalueは保持しない)
type BitflyerCoin struct {
	Coin            string  `json:"product_code"` // コイン名
	Time            string  `json:"timestamp"`    // 時間
	BestBid         float64 `json:"best_bid"`     // 売値
	BestAsk         float64 `json:"best_ask"`     // 買値
	LastTradedPrice float64 `json:"ltp"`          // 最終取引価格
}

var e = echoStart()

func main() {
	// `/` というパス（URL）と `articleIndex` という処理を結びつける
	e.GET("/", articleIndex)

	// Webサーバーをポート番号 9000 で起動する
	e.Logger.Fatal(e.Start(":9000"))
}

func echoStart() *echo.Echo {
	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {

	// ステータスコード 200 で、GetBitcoinAPI関数で取得した文字列をレスポンス
	return c.JSON(http.StatusOK, GetBitcoinAPI())
}

// GetBitcoinAPI is BitFlyerのBitcoinのAPIを取得する関数
func GetBitcoinAPI() []BitflyerCoin {

	// WebAPIに対してアクセスする
	resp, err := http.Get(bitflyerBaseURL + "btc_jpy")
	if err != nil {
		log.Fatal(err)
	}

	// 最後にWebAPIに対してのアクセスをCloseする
	defer resp.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var bitcoin []BitflyerCoin
	if err := json.Unmarshal(body, &bitcoin); err != nil {
		log.Fatal(err)
	}

	// var bitcoin BitflyerCoin
	// if err := json.Unmarshal(body, &bitcoin); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(bitcoin)

	// u := new(BitflyerCoin)
	// err = json.Unmarshal(body, u)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(body)

	// byte配列をstring型へキャスト
	// BitcoinString := string(bitcoin)

	return bitcoin
}
