package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	"github.com/pistatium/kiota_example/client"
	"github.com/pistatium/kiota_example/client/models"
	"github.com/pistatium/kiota_example/client/posts"
	"log"
	netHttp "net/http"
)

// 生成した kiota のクライアントを使って API を叩くサンプル
func main() {
	ctx := context.Background()
	authProvider := authentication.AnonymousAuthenticationProvider{}
	httpClient := netHttp.Client{}
	adapter, err := kiotaHttp.NewNetHttpRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		&authProvider,
		nil,
		nil,
		// デフォルトの実装だと request body が gzip で圧縮されて 400 Bad Request になるので httpClientを上書き
		// https://github.com/microsoft/kiota-http-go/blob/4760d5dce5fbf430b166e5c67ddbe2a0feb2bf53/pipeline.go#L65
		&httpClient,
	)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v\n", err)
	}

	adapter.SetBaseUrl("http://localhost:8080/api/v1")

	apiClient := client.NewApiClient(adapter)

	// GET /posts にリクエストを送信して Post の一覧を取得
	{
		fmt.Println("# GET /posts")
		// リクエストパラメータの設定
		limit := int32(3)
		params := posts.PostsRequestBuilderGetQueryParameters{
			Limit: &limit,
		}
		// リクエストを送信
		result, err := apiClient.Posts().Get(ctx, &posts.PostsRequestBuilderGetRequestConfiguration{
			QueryParameters: &params,
		})

		if err != nil {
			// OpenAPIで定義されているエラーは errors.As で取得出来る
			var errResp *models.ErrResponse
			if errors.As(err, &errResp) {
				fmt.Printf("request error: %s\n", *errResp.GetDetail())
				return
			}
			fmt.Printf("Error getting inference result: %+v\n", err)
			return
		}
		// Post の一覧を出力
		for _, post := range result.GetPosts() {
			fmt.Printf("id: %d, title: %s, content: %s\n", *post.GetId(), *post.GetTitle(), *post.GetContent())
		}
	}

	// POST /posts にリクエストを送信して Post を作成
	{
		fmt.Println("\n# POST /posts")
		// リクエストデータの作成
		requestData := models.NewPost()
		title := "Hello world"
		requestData.SetTitle(&title)
		content := "This is a test from kiota apiClient."
		requestData.SetContent(&content)

		// リクエストを送信
		result, err := apiClient.Posts().Post(ctx, requestData, nil)

		if err != nil {
			// OpenAPIで定義されているエラーは errors.As で取得出来る
			var errResp *models.ErrResponse
			if errors.As(err, &errResp) {
				fmt.Printf("request error: %s\n", *errResp.GetDetail())
				return
			}
			fmt.Printf("Error getting inference result: %+v\n", err)
			return
		}
		// 作成した Post の ID を出力
		fmt.Printf("saved id: %d\n", *result.GetId())
	}

}
