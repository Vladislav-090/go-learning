package main

import "fmt"

type ChatMembersRegistry interface {
	AddMember(chat string, member string)
	DeleteMember(chat string, member string)
	GetMembersFromChat(chat string) []string
	ViewAllChats()
	ChatsCount() int
}

type ChatMembersStruct struct {
	Chats map[string][]string
}

func (c *ChatMembersStruct) AddMember(chat string, member string) {
	if _, exist := c.Chats[chat]; !exist {
		c.Chats[chat] = []string{}
		fmt.Println("New Chat created without members!", chat)
	}

	for _, currentMember := range c.Chats[chat] {
		if currentMember == member {
			fmt.Println("Member already exist!", member)
			return
		}
	}

	c.Chats[chat] = append(c.Chats[chat], member)
	fmt.Println("Member added to chat", chat, member)
}

func (c *ChatMembersStruct) DeleteMember(chat string, member string) {
	members, exist := c.Chats[chat]
	if !exist {
		fmt.Println("Chat not found!", chat)
		return
	}
	for i, currentMember := range members {
		if currentMember == member {
			members = append(members[:i], members[i+1:]...)
			c.Chats[chat] = members
			fmt.Println("Member deleted", member)
			return
		}
	}
	fmt.Println("Member not found!", member)
}

func (c *ChatMembersStruct) GetMembersFromChat(chat string) []string {
	return c.Chats[chat]
}

func (c *ChatMembersStruct) ViewAllChats() {
	for chat, members := range c.Chats {
		fmt.Printf("Members : %v in chat: %s \n", members, chat)
	}
}

func (c *ChatMembersStruct) ChatsCount() int {
	return len(c.Chats)
}

func PrintInfo(c ChatMembersRegistry, chat string, member string) {
	c.ViewAllChats()
	fmt.Println("Count of chats:", c.ChatsCount())

	c.AddMember(chat, member)
	fmt.Println("Uploaded chat", c.GetMembersFromChat(chat))
	fmt.Println("Count of chats:", c.ChatsCount())

	c.DeleteMember(chat, member)
	c.ViewAllChats()
	fmt.Println("Count of chats:", c.ChatsCount())
}

func main() {
	chats := ChatMembersStruct{
		Chats: map[string][]string{
			"Whatsup":  {"Vladislav", "Viola", "Afina"},
			"Telegram": {"Dima", "Bruce", "Tanya"},
		},
	}
	PrintInfo(&chats, "Whatsup", "Semen")
}
