# Go よく使うコマンド

## テスト

```bash
go test ./...          # 全パッケージのテストを実行
go test -v ./...       # 詳細表示（テスト名・結果が見える）
go test -run TestMultiplication ./...  # 特定のテストだけ実行
```

## ビルド・実行

```bash
go build ./...         # ビルド（エラー確認だけに使うことも）
go run main.go         # ファイルを直接実行（main パッケージのみ）
```

## 依存管理

```bash
go mod tidy            # 不要な依存を削除・必要なものを追加
go get パッケージ名     # パッケージを追加
```

## フォーマット・静的解析

```bash
go fmt ./...           # コードを自動整形
go vet ./...           # 怪しいコードを警告（軽い静的解析）
```
