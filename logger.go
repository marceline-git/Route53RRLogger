package main
 
import (
  "fmt"
  "os"
  "strings"
  "time"
  "encoding/json"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/route53"
  "github.com/aws/aws-sdk-go/service/s3"
)
 
type MyEvent struct {
  HostZoneId string `json:"HostZoneId"`
  BucketName string `json:"BucketName"`
}

type MyResponse struct {
  Message string `json:"Response:"`
}

func getRecord(id string) (*route53.ListResourceRecordSetsOutput, error) {
  mySession := session.Must(session.NewSession())
  svc := route53.New(
    mySession,
    aws.NewConfig().WithRegion("us-east-1"),
  )
  input := &route53.ListResourceRecordSetsInput {
    HostedZoneId: aws.String(id),
  }
  result, err := svc.ListResourceRecordSets(input)
  if *result.IsTruncated {
    // Truncated, handle pagination
    fmt.Printf("truncated")
  }
  return result, err
}

func putObject(bucketName string, data *route53.ListResourceRecordSetsOutput) (*s3.PutObjectOutput, error) {
  mySession := session.Must(session.NewSession())
  svc := s3.New(
    mySession,
    aws.NewConfig().WithRegion("ap-northeast-1"),
  )

  jsonBytes, err := json.Marshal(*data)
  jsonFile, err := os.Create("/tmp/tmp.json")
  if err != nil {
          panic(err)
  }
  defer jsonFile.Close()
  jsonFile.Write(jsonBytes)
  jsonFile.Close()

  prefix := time.Now().Format("/2006/01/02/")
  filename := time.Now().Format("2006-0102-1504") + ".json"

  input := &s3.PutObjectInput{
    Body:   aws.ReadSeekCloser(strings.NewReader("/tmp/tmp.json")),
    Bucket: aws.String(bucketName),
    Key:    aws.String("/rrlogs"+prefix + filename),
  }
  return svc.PutObject(input)
}

func logger(event MyEvent) (MyResponse, error) {
  data,_ := getRecord(event.HostZoneId)
  response,_ := putObject(event.BucketName, data)
  fmt.Printf("res: %v", response)
  return MyResponse{Message: fmt.Sprintf("OK 200")}, nil
}
 
func main() {
  lambda.Start(logger)
}