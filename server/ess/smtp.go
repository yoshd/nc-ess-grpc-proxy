package ess

import (
	"log"
	"net/smtp"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/yoshd/nc-ess-grpc-proxy/pb"
)

const (
	SMTPServer = "ess-smtp.cloud.nifty.com"
	SMTPPort = 587
)

func (s *proxy) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	auth := smtp.PlainAuth(
		"",
		in.Auth.Username,
		in.Auth.Password,
		SMTPServer,
	)
	text := makeMailText(in.FromAddr, in.ToAddrs, in.Subject, in.Body)
	err := smtp.SendMail(
		SMTPServer+":"+strconv.Itoa(SMTPPort),
		auth,
		in.FromAddr,
		in.ToAddrs,
		[]byte(text),
	)
	if err != nil {
		log.Fatal(err)
		return &pb.SendEmailReply{Result: "NG"}, err
	}
	log.Println("Result: OK")
	return &pb.SendEmailReply{Result: "OK"}, nil
}

func makeMailText(fromAddr string, toAddrs []string, subject string, body string) string {
	text := "From: " + fromAddr + "\r\n" +
		"To: " + strings.Join(toAddrs, ",") + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body
	return text
}
