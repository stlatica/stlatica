# stlatica ドメインモデル図


<img src="domeinModel.svg" />


## 各クラス詳細

### User

本サービスにおけるユーザー。ActivityPubにおけるActor

### Activity

投稿およびいいねと類似アクション

### INBOX

ActivityPubに規定された受信コレクション

### OUTBOX

ActivityPubに規定された投稿コレクション

### UserStream

ユーザーに表示するタイムライン。サービス側で表示順序や優先度を調整する想定

### stlaticaClient

公式クライアント

### ActivityPub conformant Client 

ActivityPubに規定された準拠サードパーティクライアント

### ActivityPub conformant Federated Server 

ActivityPubに規定された連合サーバーのActor

### Actor of ActivityPub conformant Federated Server 

ActivityPubに規定された連合サーバーのActor