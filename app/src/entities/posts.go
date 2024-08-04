package entities

import "errors"

type Posts struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

func (p *Posts) Validate() error {
	if err := p.checkID(); err != nil {
		return err
	}
	if err := p.checkUserID(); err != nil {
		return err
	}
	if err := p.checkTitle(); err != nil {
		return err
	}
	if err := p.checkContent(); err != nil {
		return err
	}
	return nil
}

func (p *Posts) checkID() error {
	if p.ID == "" {
		return errors.New("IDが空です")
	}
	return nil
}

func (p *Posts) checkUserID() error {
	if p.UserID == "" {
		return errors.New("UserIDは必須です")
	}
	return nil
}

func (p *Posts) checkTitle() error {
	if p.Title == "" {
		return errors.New("タイトルは必須です")
	}
	return nil
}

func (p *Posts) checkContent() error {
	if p.Content == "" {
		return errors.New("コンテンツは必須です")
	}
	return nil
}
