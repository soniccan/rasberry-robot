package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// .envからCLOUDRUN_URLを取得
		cloudRunURL := os.Getenv("CLOUDRUN_URL")
		if cloudRunURL == "" {
			http.Error(w, "CLOUDRUN_URLが設定されていません", http.StatusInternalServerError)
			return
		}
		// 正常なHTTPレスポンスを返す
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Cloud Run URL: " + cloudRunURL))
		if err != nil {
			log.Printf("レスポンスの書き込み中にエラーが発生しました: %v", err)
		}
	})

	log.Println("サーバーがポート8080で起動しています...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
