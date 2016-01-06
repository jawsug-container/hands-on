# JAWS-UG コンテナ支部: ハンズオン資料集

http://handson.jawsug-container.org.s3.amazonaws.com/index.html  
ここのもとになっているソースファイルです。

## ドキュメント更新手順

### セットアップ

Docker Toolboxを [ここ](https://www.docker.com/docker-toolbox) からインストールします

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
