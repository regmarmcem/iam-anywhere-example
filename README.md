## これはなに
[JAWS-UGコンテナ支部 入門編 #7 初心者大歓迎LT大会](https://jawsug-container.connpass.com/event/253866/)のサンプルコード

発表内容は以下
https://speakerdeck.com/regmarmcem/eks-iam-anywhere

## 使い方
### 前提条件

1. EKS Anywhereのローカルクラスタを構築済
```zsh
➜ CLUSTER_NAME=dev-cluster
➜ eksctl anywhere generate clusterconfig $CLUSTER_NAME --provider docker > $CLUSTER_NAME.yaml
➜ eksctl anywhere create cluster -f $CLUSTER_NAME.yaml
➜ export KUBECONFIG=${PWD}/dev-cluster/generated/dev-cluster.kind.kubeconfig
```
2. DynamoDBテーブルを作成済み
* テーブル名: example-table
* Partition Key: ID(N)

### IAM Roles Anywhereの設定
[発表資料](https://speakerdeck.com/regmarmcem/eks-iam-anywhere)を参照


### アプリケーションのビルド

```zsh
➜ cd dynamodb_client
➜ CGO_ENABLED=0 go build -o dynamodbclient
➜ docker build -t dynamodbclient .
➜ docker save dynamodbclient:latest > dynamodbclient.tar
```

### Kubernetesの設定
```zsh
➜ docker save dynamodbclient:latest > dynamodbclient.tar
➜ docker cp dynamodbclient.tar dev-cluster-eks-a-cluster-control-plane:/
➜ docker exec -it dev-cluster-eks-a-cluster-control-plane bash
bash-4.2# ctr -n k8s.io images import dynamodbclient.tar
➜ kubectl create secret tls dynamodb-secret --cert=dynamodbclient.crt --key=dynamodbclient.key
➜ kubectl apply -f dynamodbclient-cm.yaml # yamlの中身は書き換えておくこと
➜ kubectl apply –f dynamodbclient.yaml
```


