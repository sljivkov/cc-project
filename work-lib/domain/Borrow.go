package domain

type Borrow struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	DateOfBorrowing string `json:"date_of_borrowing"`
	UserId          uint   `json:"user_id"`
}
