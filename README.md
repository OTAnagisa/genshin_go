# genshin_go
マイグレ
ディレクトリ移動
```
cd ./db/migrations
```
マイグレファイル作成
```
goose create create_language sql
```
マイグレ実行
```
goose postgres "host=localhost port=5432 user=postgres password=postgres dbname=genshin search_path=test sslmode=disable" up
```
