package model

type Memo struct {
	Id        uint32 `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"createdAt" json:"createdAt"`
	Uid       uint32 `db:"uid" json:"uid"`
}
