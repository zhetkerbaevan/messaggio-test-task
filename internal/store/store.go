package store

import (
	"database/sql"

	"github.com/zhetkerbaevan/messaggio-test-task/internal/model"
)

type MessagesStore struct {
	db *sql.DB
}

func NewMessagesStore(db *sql.DB) *MessagesStore {
	return &MessagesStore{db: db}
}

func (s *MessagesStore) CreateMessage(message model.MessagesPayload) (int, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO messages (content) VALUES ($1) RETURNING id", message.Content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *MessagesStore) MarkMessageAsProcessed(id int) error {
	_, err := s.db.Exec("UPDATE messages SET processed=true WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessagesStore) GetStatistics() ([]model.Messages, error) {
	rows, err := s.db.Query("SELECT * FROM messages WHERE processed=$1", true)
	if err != nil {
		return nil, err
	}

	messages := make([]model.Messages, 0)
	for rows.Next() {
		message, err := scanIntoMessages(rows)
		if err != nil {
			return nil, err
		}

		messages = append(messages, *message)
	}
	return messages, nil
}

func scanIntoMessages(rows *sql.Rows) (*model.Messages, error) {
	m := new(model.Messages)
	
	err := rows.Scan(
		&m.Id,
		&m.Content,
		&m.Processed,
		&m.Created_At,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}