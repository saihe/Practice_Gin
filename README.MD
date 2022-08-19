# Practice_Gin

Ginの練習を行なっていたが、oauth2-proxyを使用してAzure ADの認証を行なってから画面を表示するアプリを docker compose で実装した。  
表示する文字列は Ginから取得している。

## Azure ADの設定

1. アプリの登録を行う
2. 認証メニューで 種類をwebアプリ、リダイレクトURLを http://localhost に設定
3. 追加したリダイレクトURLにさらに http://localhost/home を追加
4. テナントIDをコピー

## oauth2-proxyの設定

oauth.confに値を埋め込む

1. クライアントID
2. クライアントシークレット
3. テナントID

## 動作確認

1. docker compose up -d
2. http://localhost へアクセス
3. ログインすると Herro Gin表示される
