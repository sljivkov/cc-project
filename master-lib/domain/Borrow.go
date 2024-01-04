package domain

type Borrow struct {
	ISBN   string `json:"isbn"`
	UserId uint   `json:"user_id"`
}
