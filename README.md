# What's this
Lambda 関数 Golang runtime の練習。  
Route 53 RRset の情報を取得し、どこかに保存する。  
とりあえず struct to json (および json file) も学びたかったので S3 に保存することにした。

# TODO
- timestamped S3 path
- Pagination
- response
- error handling

# 準備
## パッケージのインストール
```
$ go get -u github.com/aws/aws-lambda-go/lambda
$ go get -u github.com/aws/aws-sdk-go
```

# 設計
ListResourceRecordSets API から情報を取ってくる。定期的に。  
ほんとは API call rate の制限を受けないようにしたい…。  

https://docs.aws.amazon.com/sdk-for-go/api/service/route53/#Route53.ListResourceRecordSets

https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListResourceRecordSets.html

# Testing
json で情報を渡す。  
HostZoneId と BucketName が今のところ必要。

```json
{
  "HostZoneId": "Z123EXAMPLEAA",
  "BucketName": "sample-bucket"
}
```
