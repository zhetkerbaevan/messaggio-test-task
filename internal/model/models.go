package model

import "time"

type Messages struct {
	Id         int `json:"id"`
	Content    string `json:"content"`
	Processed  bool `json:"processed"`
	Created_At time.Time `json:"created_At"`
}

type MessagesStore interface {
	CreateMessage(MessagesPayload) (int, error)
	GetStatistics() (*Statistics, error)
	MarkMessageAsProcessed(int) error
}

type MessagesPayload struct {
	Content    string `json:"content"`
}

type Statistics struct {
    TotalMessages      int `json:"total_messages"`
    ProcessedMessages  int `json:"processed_messages"`
    UnprocessedMessages int `json:"unprocessed_messages"`
}