package main
 
import (
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/route53"
)
 
type MyEvent struct {
  HostZoneId string `json:"HostZoneId"`
}

type MyResponse struct {
  Message string `json:"Response:"`
}

/*
func init() {
  mySession := session.Must(session.NewSession())
  svc := route53.New(
    mySession,
    aws.NewConfig().WithRegion("us-east-1"),
  )
}
*/

func hello(event MyEvent) (MyResponse, error) {
  mySession := session.Must(session.NewSession())
  svc := route53.New(
    mySession,
    aws.NewConfig().WithRegion("us-east-1"),
  )
  input := &route53.ListResourceRecordSetsInput {
    HostedZoneId: aws.String(event.HostZoneId),
  }
  result, _ := svc.ListResourceRecordSets(input)
  fmt.Printf("result.IsTruncated: %v ", *result.IsTruncated)
  
  if *result.IsTruncated {
    // Truncated, handle pagination
    fmt.Printf("truncated")
  }
  return MyResponse{Message: fmt.Sprintf("OK 200", event.HostZoneId)}, nil
}
 
func main() {
  lambda.Start(hello)
}