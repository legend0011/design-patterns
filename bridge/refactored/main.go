package main

import "fmt"

// 抽象和实现解耦，各自独立开发
//	就像这里的MsgSender 和 Notification，可以各自独立开发 然后再随意组合

// MsgSender is Implementation (上面的实现，是干活的)
type MsgSender interface {
	send(msg string)
}

type TelephoneMsgSender struct{}

func (t TelephoneMsgSender) send(msg string) {
	fmt.Println("send telephone msg..." + msg)
}

type EmailMsgSender struct{}

func (t EmailMsgSender) send(msg string) {
	fmt.Println("send email msg..." + msg)
}

type WechatMsgSender struct{}

func (t WechatMsgSender) send(msg string) {
	fmt.Println("send wechat msg..." + msg)
}

// Notification is Abstraction
type Notification interface {
	notify(msg string)
}

type SevereNotification struct {
	msgSender MsgSender
}

func (n SevereNotification) notify(msg string) {
	n.msgSender.send("severe")
}

type UrgentNotification struct {
	msgSender MsgSender
}

func (n UrgentNotification) notify(msg string) {
	n.msgSender.send("urgent")
}

type TrivialNotification struct {
	msgSender MsgSender
}

func (n TrivialNotification) notify(msg string) {
	n.msgSender.send("trivial")
}

func main() {
	urgentTelephoneNotification := UrgentNotification{
		msgSender: TelephoneMsgSender{},
	}
	urgentTelephoneNotification.notify("")

	trivialWechatNotification := TrivialNotification{
		msgSender: WechatMsgSender{},
	}
	trivialWechatNotification.notify("")
}
