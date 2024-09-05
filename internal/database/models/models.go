package models

type Note struct {
	ID   uint   `gorm:"primaryKey;column:user_id"`
	Name string `gorm:"column:name"`
	Text string `gorm:"column:text"`
}

func (Note) TableName() string {
	return "Notes"
}

type UserNotes struct {
	UserID uint `gorm:"primaryKey;column:user_id"`
	NoteID uint `gorm:"primaryKey;column:note_id"`
	Note   Note `gorm:"foreignKey:NoteID;references:ID"`
}
