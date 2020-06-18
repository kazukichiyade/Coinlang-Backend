package main

// 依存パッケージの読み込み > グローバル定数 > グローバル変数 > init() > main() の順に実行（判定）される

func main() {
	route := router.echoStart()

	// Webサーバーをポート番号 9000 で起動する
	route.Logger.Fatal(route.Start(":9000"))
}
