# connpass-keyword-bot-v1

名古屋で開催されるイベントを通知してくれる LINEbot。CDK を用いて Lamdba 関数をデプロイしています。

## 環境

- go: v1.20
- gomock: v0.3.0

## セットアップ

Lambda の起動

```bash
make start-lambda
```

デプロイ

```bash
make deploy
```

SAM テンプレートの作成

```bash
make create-sam-template
```

モックの作成

```bash
mockgen -source=pkg/domain/repository/event_repository_interface.go -destination=./mocks/repository/event_repository_mock.go
```
