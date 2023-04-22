package entity

import (
    "errors"
    "time"

    "github.com/google/uuid"
    tiktoken_go "github.com/j178/tiktoken-go"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}


func NewMessage(role, content string, model *Model) (*Message, error) {
	totalTokens := tiktoken_go.CountTokens(model.GetModelName(), content)
	msg := &Message {
		ID: uuid.New().String(),
		Role: role,
		Content: content,
		Model: model,
		Tokens: totalTokens,
		CreatedAt: time.Now(),
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}
	
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system"  && m.Role != "assistant" {
		return errors.New("invalid role")
	}

	if m.Content == "" {
		return errors.New("content is empty")
	}

	if m.CreatedAt.IsZero() {
		return errors.New("created_at is empty")
	}

	if m.CreatedAt.After(time.Now()) {
		return errors.New("created_at is in the future")
	}

	if m.Model == nil {
		return errors.New("model is empty")
	}

	if m.Model.MaxTokens <= m.Tokens {
		return errors.New("message is too long")
	}

	if m.Model.MaxTokens < 0 {
		return errors.New("model max tokens is negative")
	}

	if m.Model.MaxTokens == 0 {
		return errors.New("model max tokens is zero")
	}

	return nil
}

func (m *Message) GetQtdTokens() int {
	return m.Tokens
}