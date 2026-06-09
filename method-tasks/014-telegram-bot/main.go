package main

import "fmt"

type TelegramBot struct {
	Name string
	IsRunning bool
	Messages int
}

func (t *TelegramBot) Start() {
	if t.IsRunning {
		fmt.Println("Bot is already started!")
		return
	}

	t.IsRunning =true
	fmt.Println("Bot is started")
}

func (t *TelegramBot) Stop() {
	if !t.IsRunning {
		fmt.Println("Bot is already off")
		return
	}
	t.IsRunning = false
	fmt.Println("Bot turned off")
}

func (t *TelegramBot) SendMessage(){
	if !t.IsRunning {
		fmt.Println("Turn on the bot first")
		return
	}
	t.Messages ++
	fmt.Println("Message sent")
}

func (t *TelegramBot) PrintInfo(){
	fmt.Println("Bot name:", t.Name)
	fmt.Println("Status:", t.IsRunning)
	fmt.Println("messages already sent", t.Messages)
}

func main() {
	bot := TelegramBot{
	Name:      "AlertBot",
	IsRunning: false,
	Messages:  0,
}

bot.PrintInfo()

bot.SendMessage()
bot.Start()
bot.SendMessage()
bot.SendMessage()
bot.Stop()
bot.SendMessage()

bot.PrintInfo()

}