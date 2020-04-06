package botgolang

import "fmt"

//go:generate easyjson -all chat.go

type ChatAction = string

const (
	TypingAction  ChatAction = "typing"
	LookingAction ChatAction = "looking"
)

type ChatType = string

const (
	Private ChatType = "private"
	Group   ChatType = "group"
	Channel ChatType = "channel"
)

type Chat struct {
	client *Client
	// Id of the chat
	ID string `json:"chatId"`

	// Type of the chat: channel, group or private
	Type ChatType `json:"type"`

	// First name of the user
	FirstName string `json:"firstName"`

	// Last name of the user
	LastName string `json:"lastName"`

	// Nick of the user
	Nick string `json:"nick"`

	// User about or group/channel description
	About string `json:"about"`

	// Rules of the group/channel
	Rules string `json:"rules"`

	// Flag that indicates that requested chat is the bot
	IsBot bool `json:"isBot"`

	// Title of the chat
	Title string `json:"title"`

	// Is this chat public?
	Public bool `json:"public"`

	// Is this chat has join moderation?
	JoinModeration bool `json:"joinModeration"`

	// You can send this link to all your friends
	InviteLink string `json:"inviteLink"`

	// Chat admins list
	Admins []Admin `json:"admins"`
}

func (c *Chat) resolveID() string {
	switch c.Type {
	case Private:
		return c.Nick
	default:
		return c.ID
	}
}

func (c *Chat) SendActions(actions ...ChatAction) error {
	fmt.Println(c.resolveID())
	return c.client.SendChatActions(c.resolveID(), actions...)
}
