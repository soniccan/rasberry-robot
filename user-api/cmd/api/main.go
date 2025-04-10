package main

import (
	"context"
	"log"
	"net/http"
	"os"

	cloudLogging "cloud.google.com/go/logging"
	"github.com/joho/godotenv"
	"github.com/rasberry-robot/user-api/pkg/infrastructure/gcp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}

	gcp, err := gcp.NewGCP(context.Background(), os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatalf("GCPの初期化に失敗しました: %v", err)
	}

	gcp.CloudLogging().Log(context.Background(), "user-api", cloudLogging.Default, "テストログ")

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
