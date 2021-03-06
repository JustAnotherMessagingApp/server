package main

import (
	"errors"
	"github.com/twinj/uuid"
)

var (
	ErrMessageObjectNil = errors.New("api: message nil")
)

// Message holds message data and nothing more.
// Due to the high volume of messages we use a UUID to provide the required
// keyspace for the Id of all the messages. For messages we suggest a version 1
// UUID based on a timestamp to keep it sequential (CSPRNG not needed for
// message IDs.
type Message struct {
	Id   string
	Body string
}

// NewMessage takes a string and returns a new Message with a unique Id. Message
// must be saved for persistence.
func NewMessage(body string) *Message {
	return &Message{Id: uuid.NewV1().String(), Body: body}
}

// Save will store the message in the database and will return an error in case
// of DB failure.
func (m *Message) Save() error {
	return db.SaveMessage(m)
}

// Delete will remove any matching record from the database permanently and will
// return an error in case of DB failure.
func (m *Message) Delete() error {
	return db.DeleteMessage(m)
}
