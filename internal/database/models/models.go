package models

type Note struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
	Text string `db:"text"`
}

func (Note) TableName() string {
	return "Notes"
}

type UserNotes struct {
	UserID uint
	NoteID uint
	Note   Note
}
