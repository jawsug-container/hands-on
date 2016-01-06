01. Get started with Docker
===============================

.. role:: blue
.. role:: green
.. role:: captions
.. role:: commands
.. role:: results

.. raw:: html

  <style>
    blockquote {
      margin: 0 0 30px 30px;
    }
    .blue {
      color: blue;
    }
    .green {
      color: #20AB20;
    }
    .captions {
      display: inline-block;
      margin: 0 0 -17px 0;
    }
    .commands {
      color: #FF6F6F;
      font-weight: bold;
      display: inline-block;
      margin: 0 0 -17px 0;
    }
    .results {
      color: silver;
      display: inline-block;
      margin: 0 0 -17px 0;
    }
  </style>

1. Docker のインストール
-------------------------------

1.1. Docker Toolboxを |docker-toolbox| からインストールします

.. |docker-toolbox| raw:: html

   <a class="reference external" href="https://www.docker.com/docker-toolbox" target="_blank">ここ</a>

1.2. Docker Quickstart Terminal を起動します

  :captions:`こんな画面が立ち上がります`

  .. parsed-literal::

                            ##         .
                      ## ## ##        ==
                   ## ## ## ## ##    ===
               /"""""""""""""""""\\___/ ===
          ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~ /  ===- ~~~
               \\______ o           __/
                 \\    \\         __/
                  \\____\\_______/

    :blue:`docker` is configured to use the :green:`default` machine with IP :green:`192.168.99.100`
    For help getting started, check out the docs at https://docs.docker.com

    macosx $

  .. important::

    もし *'default machine with IP'* が空だったら、VirtualBoxを起動し
    **default** という名前の仮想マシンをスタートします。ターミナルも再起動！

1.3. docker のバージョンを確認してみましょう

  :commands:`コマンド`

  .. code-block:: bash

    docker version

  :results:`結果例`

  .. code-block:: bash

    Client:
     Version:      1.9.1
     API version:  1.21
     Go version:   go1.4.3
     Git commit:   a34a1d5
     Built:        Fri Nov 20 17:56:04 UTC 2015
     OS/Arch:      darwin/amd64

    Server:
     Version:      1.9.1
     API version:  1.21
     Go version:   go1.4.3
     Git commit:   a34a1d5
     Built:        Fri Nov 20 17:56:04 UTC 2015
     OS/Arch:      linux/amd64

2. Nginx の起動
-------------------------------

2.1. nginx を起動してみます

  :commands:`コマンド`

  .. code-block:: bash

    docker run -p 80:80 nginx

2.2. ブラウザから確認

  | Docker Quickstart Terminal 起動時に表示される **default machine with IP** を確認し
  | その IP アドレスをブラウザに入力してみましょう。
