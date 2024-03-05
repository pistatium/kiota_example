package main

import (
	"context"
	"errors"
	"fmt"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"log"
	netHttp "net/http"
	"time"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	"github.com/pistatium/kiota_example/client"
	"github.com/pistatium/kiota_example/client/models"
	"github.com/pistatium/kiota_example/client/posts"
)

// 生成した kiota のクライアントを使って API を叩くサンプル
func main() {
	ctx := context.Background()
	authProvider := authentication.AnonymousAuthenticationProvider{}
	httpClient := &netHttp.Client{
		CheckRedirect: func(req *netHttp.Request, via []*netHttp.Request) error {
			return netHttp.ErrUseLastResponse
		},
		Timeout: time.Second * 100,
	}
	transport := netHttp.DefaultTransport.(*netHttp.Transport).Clone()
	transport.ForceAttemptHTTP2 = true
	transport.DisableCompression = false
	httpClient.Transport = transport

	fmt.Println(httpClient)
	adapter, err := kiotaHttp.NewNetHttpRequestAdapter(&authProvider)
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
		header := abstractions.NewRequestHeaders()
		result, err := apiClient.Posts().Post(ctx, requestData, &posts.PostsRequestBuilderPostRequestConfiguration{
			Headers: header,
			Options: []abstractions.RequestOption{
				kiotaHttp.NewCompressionOptions(false),
			},
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
		// 作成した Post の ID を出力
		fmt.Printf("saved id: %d\n", *result.GetId())
	}
}
