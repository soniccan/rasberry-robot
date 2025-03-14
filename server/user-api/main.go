package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("CLOUDRUN_ACCESS_TOKEN")
	if token == "" {
		log.Fatal("CLOUDRUN_ACCESS_TOKENが設定されていません")
	}

	cloudRunURL := os.Getenv("CLOUDRUN_URL")

	req, err := http.NewRequest("GET", cloudRunURL, nil)
	if err != nil {
		log.Fatalf("リクエスト作成エラー: %v", err)
	}
	// AuthorizationヘッダーにBearerトークンを設定
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("リクエスト実行エラー: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("レスポンス読み込みエラー: %v", err)
	}

	log.Printf("レスポンスステータス: %s", resp.Status)
	log.Printf("レスポンス本文: %s", string(body))
}
