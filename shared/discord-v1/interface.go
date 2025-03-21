package discord

import "time"

type Hook interface {
	OnInit() error
	OnCreateChatMessage(message ChatMessage) error
	OnCreateInteraction(interaction Interaction) error
	OnEvent(event string) error
}

type CreateChatMessageHelper interface {
	SendMessage(message ChatMessage) error
	SendMessageToChannel(channelID string, message ChatMessage) error
}

type ChatMessage struct {
	// The ID of the message.
	ID string

	// The ID of the channel in which the message was sent.
	ChannelID string

	// The ID of the guild in which the message was sent.
	GuildID string

	// The content of the message.
	Content string

	// The time at which the messsage was sent.
	// CAUTION: this field may be removed in a
	// future API version; it is safer to calculate
	// the creation time via the ID.
	Timestamp time.Time

	// The time at which the last edit of the message
	// occurred, if it has been edited.
	EditedTimestamp *time.Time

	// The roles mentioned in the message.
	MentionRoles []string

	// Whether the message is text-to-speech.
	TTS bool

	// Whether the message mentions everyone.
	MentionEveryone bool

	// The flags of the message, which describe extra features of a message.
	// This is a combination of bit masks; the presence of a certain permission can
	// be checked by performing a bitwise AND between this int and the flag.
	MessageFlags int
}

type Interaction struct {
	// The ID of the interaction.
	ID string

	// The ID of where the interaction was triggered.
	GuildID string

	// The ID of the channel in which the interaction was triggered.
	ChannelID string

	// Message that triggered the interaction.
	Message string
}

// UsePartialHooks is a partial implementation of the hook Interface.
//
// It is useful for embedding in a struct that only needs to
// implement a subset of the hook Interface.
type UsePartialHooks struct{}

func (u *UsePartialHooks) OnInit() error {
	return nil
}

func (u *UsePartialHooks) OnCreateChatMessage(message ChatMessage) error {
	return nil
}

func (u *UsePartialHooks) OnCreateInteraction(interaction Interaction) error {
	return nil
}

func (u *UsePartialHooks) OnEvent(event string) error {
	return nil
}
