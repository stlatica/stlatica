# stlatica ログ設計

## 概要
Stlaticaバックエンドにおけるログの設計  
目的はエラー調査・パフォーマンスの確認・デバッグ確認（開発用）  
全体的なログの設計をまとめたため、各ログの書き出しパラメータの詳細は別途設計資料を作成するものとする



## ログ方針
- ログは検索・解析等の扱いがしやすいように構造化する
- ログイベントごとに専用の書き出し先（ファイル/テーブルなど）を用意
    -  イベント別のログを連携できるように共通のログIDを設定
- 最低限のログ収集・分析は立てたインスタンス内で実行できるようにし、外部サービスとの連携機能はオプションとして追加設定できるものにする

## フォーマット
JSON  
→ 他サービスとも連携しやすく、大体の人がわかるフォーマットであるため


## ログの種類
- リクエストログ
    - 内容 : APIURL, IP, リクエストパラメータ、ユーザ識別情報（あれば）等
    - 用途 : アクセスの監視、利用状況の分析
- クエリログ
    - 内容 : 実行クエリ, クエリ実行時間, スロークエリの発生等
    - 目的 : パフォーマンスチューニング・DB負荷の監視  
    (mySQL標準のgeneral_log, slow_query_logとは別)
- アプリケーションエラーログ
    - 内容 : エラーメッセージ, スタックトレース等
    - 目的 : システム稼働状況の監視・障害調査
- アプリケーションログ 
    - 内容 : デバッグ用のログの他、アプリケーションにおけるログ全般
    - 目的 : デバッグ・システム稼働状況の確認

## ログの保存先
ログファイルへの書き出しに加えて、外部（クラウドサービスやウェブサービス）連携を必要としない構成として立てたインスタンス内で最低限の構造化されたログ分析をすることを前提に保存先を考える（NoSQL/ElasticSearch）

### 必須構成

ローカルのログファイル + NoSQL/ElasticSearch

- ローカルのログファイル
    - ローカルで完結できる。簡単な調査はここを見るだけでよい
    - tail -f とかでログを流し見できる
    - ログローテーション(logrotate)して1ファイルを大きくさせすぎない
    （"myapp-Y-m-d.log"みたいに命名,つまり日毎）
    - フォーマットをJSONにするので若干読みづらい

- NoSQL/ElasticSearch
    - 構造化データ（JSON）を扱える
    - RDBより書き込みのパフォーマンス面が優位
        - ElasticSearch (候補)
            - 分散型なのでスケールしやすい 
            - 分析機能や可視化ツールと相性が良い
        - MongoDB (候補)
            - 分散型なのでスケールしやすい

### 外部連携
※ 検討中

- Cloud Watch Logs
    -  ローカルのログファイルを設定できるので導入が容易
        - ローテーションしているファイルでも"myapp-*.log"みたいにして監視できるはず

### 採用しなかったもの（一部）
- RDBS(mySQLとか)
    - 頻繁に書き込まれることが想定されるため蓄積とI/O時負荷に懸念
        - 整合性を担保するトランザクション処理はここでは過剰
        - データ蓄積に対する対処をする必要がでてくる
    - (なお) RDBSを含むスケールできるようなサービスを、webサーバの動作（トランザクション）とは独立して使うならそこまで毛嫌いしなくていいのかも

## ログのバックアップ
- ローカルのログファイル
    - ログローテーションしたファイルを圧縮して残す
        -  バッチ処理
- 外部ストレージサービス（S3とかGCS）
    - ローカルのログファイルとか, NoSQLベースならバックアップデータを放り込む

## ログの保持期間
外部ストレージに放り込む場合、そのバックアップはずっと残す。（運用に任せ、システム面でサポートを考えない）
ファイル・NoSQLベースのログは実際に書き込まれる量次第で考える（決め打ちするなら3ヶ月とか1年とか）

## ログレベル
ログの種類別にログレベルとそれに伴う内容/書き出しタイミングは異なる。

- リクエストログ
    - INFO : ほぼこれになりそう
    - WARN : エラーではないが検知しておきたい動作発生時
        - 非推奨APIの実行？ 
- クエリログ
    - ERROR : クエリ実行エラー時
    - WARN : エラーではないが検知しておきたい動作発生 時
        - 定義した閾値を超える実行時間のクエリが実行
    - INFO : ↑以外のクエリ実行
- アプリケーションエラーログ
    - ERROR : エラー発生時
    - TRACE : トレース（エラーメッセージを数行に分けるなら使う）
    - WARN : エラーではないが検知しておきたい動作発生時 
- アプリケーションログ 
    - INFO : 開発用じゃなくても必要なもの。例えばバッチのログとか他システムとの連携に関するログとかあれば
    - DEBUG : 開発用に追加するもの。


## ログパラメータ
共通＋各イベントごとのパラメータを設定

### 共通
- ログレベル
- ログ個別ID
    - ログを一意に定めるもの
- ログID
    - 同タイミングで発生した各種ログを紐付けやすくする  
    （ログイベントごとに書き出しファイルを変えることもある設計のため）
    - 順番に並んでくれているとうれしい気がするのでULIDとか?
- 処理実行時間
    - 書き出しの瞬間ではなく処理を開始したタイミングのタイムスタンプ
- タグ（ログイベント）
    - ログイベントの分類
        - リクエスト/クエリ/アプリエラー/アプリ
- リクエストURL (エンドポイント)
    - バッチ処理とか内部処理のエラーなら空 
- リクエストIP
- ユーザ識別ID（actor_id）
    - とれないときは空か0 

### リクエストログ
- リクエストパラメータ
    - データ量が多くなりやすいのでON/OFFできたほうがよさそう 
- UserAgent

### クエリログ
- クエリ内容
- 実行にかかった時間

### アプリケーション(エラー)ログ
- エラーコード
- メッセージ

### フロントエンド側エラーログ
※ フロントからエラーとして受けたいログ送信を受け付けて書き出しする場合
- 文字列 (エラーメッセージ)

### 機密情報の取り扱い（maskingなど）
どんな情報を扱うか、参照の必要性に応じてするしないの判断は都度したほうがよいという前提で、方針は下記
- 完全にわからなくすべきものは適当な文字に置き換える
- ある程度区別はつけたいみたいなシチュエーションのときは一方向ハッシュをかける

## ログ収集ツール
ログ分析用データストア/外部サービスへのログ転送のため  
ログ収集ツールにあまりこだわりがないのでメジャーなでよいのでは(思考停止)

- Fluentd (採用?)
    - メジャーなログ収集ツールA
    - 比較的軽量
    - 単体で動作する
    - ログデータをバッファに持ち書き込みに失敗しても再実行できる
    - プラグインも多く複数のサービスと連携できそう

- logstash (候補)
    - メジャーなログ収集ツールB
    - Elastic Stack (ELK) のコンポーネントの一つなので、分析やデータストアにElasticSearchやkibanaが使われるならシナジーが一番高そうではある

## ログ分析ツール
可視化ツールに現状のこだわりがないのでメジャーなでよいのでは(思考停止2)

- Grafana (採用？)
    - メジャーな可視化ツールA
    - プラグインの導入込みで複数のデータストアに対応
- kibana (候補)
    - メジャーな可視化ツールB
    - ElasticSearchでしか使えないらしい。多少融通が利くほうがよいので不採用の方向（せっかくElasticSearchつかうならこっちにしようぜという判断も一理ある）

分析用にローカルで用意するデータストアとの相性にもよりそうなのでいったん検討中としたい    


### ログ分析の工程について

- [1] ログファイルを収集ツールのソースとする場合
    - webサーバでログファイル書き込み→収集ツールが検知した書き込みを転送→分析用データストア/外部サービス

- [2] HTTPでログ情報を収集ツールに送信する場合
    - webサーバでログ情報を収集ツールへ送信→収集ツールがタグごとに転送→分析用データストア/外部サービス

[1]はログファイルの書き込みまではwebサーバが実行する。
書き込みに失敗することがあると、ログは失われる。
[2]は取り合えず情報は送信するので、情報を受け取ることができれば書き込みを収集ツールに任せられる。HTTP通信のためのオーバヘッドは発生する。ツールをホストしているサーバが落ちたりするとログは失われる。ローカルログファイルの書き込み自体はこの工程を踏まないほうが無駄はなさそう...


### TODO
- 外部サービスとの連携
    - スコープがさらに広がるので、ここでは検討中とする 
- ローカル分析用のデータストア/サービス
    - 未決定な項目が多め。分析用と割り切ってスコープを区切ってもいいかも 






