package themap

import (
	"fmt"
)

// CreateUser registers a user at theMAP
func (p *Payment) CreateUser(ip, phone, email string) error {

	var err error

	p.User = User{IP: ip, Phone: phone, Email: email}

	err = proceedRequest("POST", "/createUser", p)
	if err != nil {
		return err
	}

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err

}
