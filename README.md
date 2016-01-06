# JAWS-UG コンテナ支部: ハンズオン資料集

http://handson.jawsug-container.org.s3.amazonaws.com/index.html  
ここのもとになっているソースファイルです。

## ドキュメント更新手順

### セットアップ

Docker Toolboxを [ここ](https://www.docker.com/docker-toolbox) からインストールします。  
Pythonはもとより、Sphinxなどは一切不要です。

### 確認用 Web サーバ起動

```
cd /go/to/this/directory
docker-compose up -d
```

### Sphinxドキュメントの HTML化

```
cd /go/to/this/directory
docker-compose -f sphinx.yml run clean-html
```

## コントリビューション

大歓迎です！

1. Forkします ([https://github.com/jawsug-container/hands-on.git/fork](https://github.com/jawsug-container/hands-on.git/fork))
2. featureブランチを作ります
3. 変更を commitしてください
4. ローカルの変更を masterブランチに対し rebaseします
5. Pull Requestをお願いします
