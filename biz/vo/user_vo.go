package vo

import "time"

type UserVO struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	NickName  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Sign      string    `json:"sign"`
	CreatedAt time.Time `json:"created_at"`
}
