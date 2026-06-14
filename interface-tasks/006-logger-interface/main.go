package main

import "fmt"

type Logger interface{
	Log(message string)
}

type ConsoleLogger struct {
	Name string
}

type FileLogger struct {
	Name string
}

type TelegramLogger struct {
	Name string
}

func (t *TelegramLogger) Log(message string) {
	fmt.Println("Telegram Log:", message)
}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("Console Log:", message)
}

func (f *FileLogger) Log(message string) {
	fmt.Println("File Log:", message)
}

func WriteLog(l Logger, message string) {
	l.Log(message)
} 

func main() {
	WriteLog(&TelegramLogger{Name: "Telegram"}, "log has been written")
	WriteLog(&FileLogger{Name: "File"}, "log has been written")
	WriteLog(&ConsoleLogger{Name: "Console"}, "log has been written")
}