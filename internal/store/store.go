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

func (s *MessagesStore) GetStatistics() (*model.Statistics, error) {
	var stats model.Statistics
	//Get number of all messages
    err := s.db.QueryRow("SELECT COUNT(*) FROM messages").Scan(&stats.TotalMessages)
    if err != nil {
        return nil, err
    }

    //Get number of processed messages
    err = s.db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = $1", true).Scan(&stats.ProcessedMessages)
    if err != nil {
        return nil, err
    }

    //Get number of unprocessed messages
    err = s.db.QueryRow("SELECT COUNT(*) FROM messages WHERE processed = $1", false).Scan(&stats.UnprocessedMessages)
    if err != nil {
        return nil, err
    }

    return &stats, nil
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