package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

// GetBitcoinAPI is BitFlyerのBitcoinのAPIを取得する関数
func GetBitcoinAPI() BitflyerCoin {

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

	var bitcoin BitflyerCoin
	if err := json.Unmarshal(body, &bitcoin); err != nil {
		log.Fatal(err)
	}

	fmt.Println(bitcoin)
	return bitcoin
}
