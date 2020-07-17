<<<<<<< HEAD
# Go langの勉強する
Go langを勉強するため、HTTPサーバを作る

## フレームワークの種類
* Gin
軽量、高速処理、堅牢
* Echo
軽量、RestAPI向け、読みやすい、Ginより高速
* iris
ベンチマーク最速、サンプル多い、癖少ない
* Revel
フルスタックWebフレームワーク、MVC、自動コンパイル
* Beego
MVC、自動テスト、専用CLI

## Ginによるサンプルアプリ
[ここ](/gin_test)ではGinによるWebアプリを作る。

### モジュールのインストール
1. ```go mod init```で初期化
golangでのモジュール管理にはgo modを用いる。
次のようにして実行する。
```
mkdir gin_test
cd gin_test
go mod init Matchapy/gin_test
```
注意:$GOPATH外で作業する場合はinit時の引数にモジュール名を追加する

2. main.goのimportに使用するモジュールを書く
今回はGinを用いるので
```
import (
    "net/http"
    "github.com/gin-gonic/gin"
)
```
とする

3. ```go build```でモジュールを自動インストールする
4. main.goを書く。
5. views/index.html, assets/style.cssを配置

### ディレクトリ構造を整理
```
gin_test
    ├ assets
    │     └ [CSSなどを置く]
    ├ config
    │     ├ [データベース周辺]
    │     └ dummy_db.go
    ├ crypto
    │     ├ [暗号化に関わる重要なヘルパー]
    │     └ crypto.go
    ├ helpers
    │     ├ [validation等のその他のヘルパー]
    │     └ helpers.go
    ├ models
    │     └ [モデルを置く]
    ├ routes
    │     ├ [ルーターから呼び出されるハンドラ]
    │     ├ routes.go
    │     └ user_routes.go
    ├ sessions
    │     ├ [セッション管理]
    │     └ dummy_sessions.go
    ├ views
    │     └ [HTMLテンプレートを置く]
    └ main.go

```
### routesを使うようにmain.goを変更
```
import (
    "github.com/gin-gonic/gin"
    "Matchapy/gin_test/routes"
)
```

## 参考
[Modulesの使い方](https://qiita.com/uchiko/items/64fb3020dd64cf211d4e)
[Webアプリ初心者がGo言語でサーバサイド](https://qiita.com/wsuzume/items/75d5c0cd2dd5a1963b9e)
=======
# Go langを勉強する
Go langを勉強するため、HTTPサーバを作る
>>>>>>> 8e2090ca1dee01a7fb583b5da7cf461a4664b019
