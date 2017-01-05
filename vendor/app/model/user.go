package model

type User struct {
	ID       uint32 `db:"id" json:"id"`
	Phone    string `db:"phone" json:"phone"`
	Password string `db:"password" json:"-"`
}
