package main

import "fmt"

type Notifier interface {
	Notify(message string)
}

type TelegramNotifier struct {
	ChatId string
}
type WhatsUpNotifier struct {
	ChatId string
}
type EmailNotifier struct {
	ChatId string
}

func (t *TelegramNotifier) Notify(message string) {
	fmt.Println("Telegram alert:", message)
}

func (w *WhatsUpNotifier) Notify(message string) {
	fmt.Println("Whatsup alert:", message)
}

func (e *EmailNotifier) Notify(message string) {
	fmt.Println("Email alert:", message)
}

func SendAlert(n Notifier, message string) {
	n.Notify(message)
}


func main() {
	telegram := TelegramNotifier{ChatId: "idTelegram"}
	whatsup := WhatsUpNotifier{ChatId: "idWhatsup"}
	email := EmailNotifier{ChatId: "idEmail"}

	SendAlert(&telegram,"High CPU usage" )
	SendAlert(&whatsup,"High CPU usage")
	SendAlert(&email,"High CPU usage")
}