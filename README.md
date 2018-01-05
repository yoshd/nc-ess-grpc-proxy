# nc-ess-grpc-proxy

## Client samples

### Go

See Protocol Buffers Reference [Go Generated Code](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

* Go client code sample

```
package main

import (
	"github.com/yoshd/nc-ess-grpc-proxy/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	host     string   = "YourESSProxyHost"
	port     string   = ":YourESSProxyPort"
	username string   = "YourESSSMTPUsername"
	password string   = "YourESSSMTPPassword"
	fromAddr string   = "YourFromMailAddr"
	toAddrs  []string = []string{"YourToMailAddrs"}
	subject  string   = "YourMailSubject"
	body     string   = "YourMailBody"
)

func main() {
	target := host + port
	conn, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connection err: %v", err)
	}

	defer conn.Close()

	c := pb.NewESSProxyClient(conn)

	auth := pb.AuthInfo{Username: username, Password: pasword}
	s := pb.SendEmailRequest{
		Auth:     &auth,
		FromAddr: fromAddr,
		ToAddrs:  toAddrs,
		Subject:  subject,
		Body:     body,
	}
	re, err := c.SendEmail(context.Background(), &s)
	if err != nil {
		log.Fatalf("PRC err: %v", err)
	}
	log.Printf("Result: %s", re.Result)
}

```
