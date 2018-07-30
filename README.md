# ハンズオン

[![jawsug/container](http://dockeri.co/image/jawsug/container)](https://hub.docker.com/r/jawsug/container/)

Supported tags and respective `Dockerfile` links:

・fargate-handson ([Dockerfile.fargate](https://github.com/jawsug-container/hands-on/blob/master/Dockerfile.fargate))  

## 前提条件

- 端末（Windows, Mac, ..）で Docker が利用できること
- 利用可能な AWS アカウント、Administrator 権限のある IAM ユーザ、そのアクセスキーがあること

### Docker が利用できること

Docker Store からインストーラーをダウンロードし、お手もとの環境へ Docker をインストールしてください。  
Windows へのインストールや動作確認がうまくできない場合は、EC2 の利用をご検討ください。  
（EC2 を利用する場合、パブリック IP アドレスの取得、8080 番ポートの解放が必要です）  

- Mac は [こちら](https://store.docker.com/editions/community/docker-ce-desktop-mac)
- Windows は [こちら](https://store.docker.com/editions/community/docker-ce-desktop-windows)
- EC2 は [こちら](https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/docker-basics.html)

以下のコマンドで、Client と Server の情報が正常に返ってくることを確認してください。

```
$ docker version
```

応答例）

```
Client:
 Version:      18.03.1-ce
 API version:  1.37
 Go version:   go1.9.5
 Git commit:   9ee9f40
 Built:        Thu Apr 26 07:13:02 2018
 OS/Arch:      darwin/amd64
 Experimental: false
 Orchestrator: swarm

Server:
 Engine:
  Version:      18.03.1-ce
  API version:  1.37 (minimum version 1.12)
  Go version:   go1.9.5
  Git commit:   9ee9f40
  Built:        Thu Apr 26 07:22:38 2018
  OS/Arch:      linux/amd64
  Experimental: true
```

### AWS の準備ができていること

#### 1. AWS アカウントの開設

以下のサイトを参考に、AWS アカウントを用意してください。  
https://aws.amazon.com/jp/register-flow/

#### 2. 重要！「AWS 利用開始時に最低限おさえておきたい 10 のこと」の確認

以下のスライドを読み、必要に応じて初期設定を変更してください。  
https://www.slideshare.net/AmazonWebServicesJapan/20180403-aws-white-belt-online-seminar-aws10

#### 3. Administrator 権限のある IAM ユーザと、そのアクセスキーの発行

IAM（Identity and Access Management）とは、AWS でのユーザーや権限を管理するサービスです。  
AWS をより安全に利用するために、初期設定されたルートユーザーとは別に、IAM ユーザーを用意します。

3.1. IAM 管理者ユーザーを作成します

https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/getting-started_create-admin-group.html

3.2. アクセスキーを発行します

発行したアクセスキー・シークレットキーはどこか安全ば場所に保管しておいてください。  
インターネット上に流出させないよう、[git-secrets](https://github.com/awslabs/git-secrets) などをご検討ください。
https://docs.aws.amazon.com/ja_jp/IAM/latest/UserGuide/id_credentials_access-keys.html

#### 4. EC2 を起動・停止する

AWS では開設直後のアカウントの場合、仮想サーバーの起動数などに制限がかかっていることがあります。  
以下の手順に従い、 **t2.micro** でのインスタンス起動、設定、接続、および終了を一度実行してください。

https://aws.amazon.com/jp/getting-started/tutorials/launch-a-virtual-machine/


## コンテンツ

- [Fargate アプリケーションの継続的デリバリー](https://github.com/jawsug-container/hands-on/blob/master/fargate/README.md)
