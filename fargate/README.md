# Fargate アプリケーションの継続的デリバリー

## ハンズオン環境の起動

### 1. アクセスキーをそれぞれ変数に設定します

```
$ export AWS_ACCESS_KEY_ID=
$ export AWS_SECRET_ACCESS_KEY=
```

### 2. ハンズオン環境を起動します

```
$ docker run --rm -it -e AWS_DEFAULT_REGION=ap-northeast-1 \
     -e AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v $(pwd):/root/config \
     -p 8080:8080 jawsug/container:fargate-handson
```

### 3. ブラウザで環境に接続します

ブラウザで http://localhost:8080 を開いてください。  
（EC2 を利用している場合は、localhost を EC2 のパブリック IP アドレスに読み替えてください）

パスワードを聞かれるので **jawsug** と入力してください。

### 4. ハンズオンの実施

ハンズオンはブラウザ内の Jupyter notebook で行います。  
以下の順序でハンズオンを進めてください。

```
- 01-provision-aws-resources.ipynb
```

### 5. 後片付け

`Ctrl + C` でコンテナに停止シグナルを送ると、Jupyter notebook から  
`Shutdown this notebook server (y/[n])?` と聞かれます。 
`y` と入力してコンテナを終了しましょう。 