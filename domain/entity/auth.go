package entity

import "errors"

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *LoginReq) Validate() error {
	if req.Username == "" {
		return errors.New("username is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
