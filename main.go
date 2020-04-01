package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, GetBitcoinApi())
}

//type BitflyerBitcoin struct {
//	Bitcoin string
//	Id      string
//	Time    time.Time
//}

//var coin BitflyerBitcoin

// BitFlyerのBitcoinのAPIを取得する関数
func GetBitcoinApi() string {

	// GetでWebAPIに対してアクセスする
	resp, err := http.Get(bitflyerBaseUrl + "btc_jpy")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	// 最後にapiをCloseする
	defer resp.Body.Close()

	// ReadAllは、エラーまたはEOFに達するまで読み込み、読み込んだデータを返す
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(byteArray)

	/*var bitcoin bitflyerBitcoins
	err = json.Unmarshal(body, &bicoin)
	if err != nil {
		log.Fatal(err)
	}*/

	// byte配列をstring型へキャスト
	castString := string(byteArray)

	// a := json.Unmarshal(byteArray)
	//fmt.Println(castString)

	//resp2 := json.Unmarshal(byteArray, &coin)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(reflect.TypeOf(resp2))

	return castString
}
