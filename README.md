# TODO
- そもそも完成してない
- データストア
- Pagenation

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

