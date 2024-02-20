# Kiotaを使ってGoのAPIクライアントを動かすサンプル

## Kiota とは

https://learn.microsoft.com/ja-jp/openapi/kiota/

Microsoftが公開している、OpenAPIからAPIクライアントを自動生成するツールです。
OpenAPI(Swagger)の定義からGoやPythonなどのAPIクライアントを自動生成することができます。

このレポジトリではKiotaを使ってGoのAPIクライアントを生成し、実際にAPIを叩くサンプルを作成しています。 


## ディレクトリ構成
* /webapi/
  * サンプルで作ったAPIサーバー
  * webapi/docs/swagger.yaml にAPIの定義が書いてあります
    * swagger.yaml は swag によってアプリケーションのコードから生成しています
* /client/
  * swagger.yaml から自動生成したGoのAPIクライアント
* /client_usage.go
  * /client/ で生成したAPIクライアントを使ってAPIを叩くサンプル
* Makefile
  * WebAPIの実行、APIクライアントの生成などを行うためのMakefile

## 動かし方

```shell
# APIサーバーを起動
make run_webapi

# 別のターミナルで client_usage.go を実行し、APIにアクセスしてみる
make run_example_client
```

## APIの仕様を変更した場合

```shell
# webapi の中の Go のコードを変更する
# swag を使っているのでソース内のコメントを元に swagger.yaml が作られます

# swagger.yaml の再生成
make run_webapi

# swagger.yaml から APIクライアントを再生成
make generate_client

# 必要に応じて client_usage.go を修正
```