package markups

import (
	"fmt"

	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database/models"
)

func GetNotesKeyboard(notes []models.Note) *dto.InlineKeyboardMarkup {
	raw := 1
	resultKeyboard := &dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]*dto.InlineKeyboardButton{
			{{Text: "📝 создать заметку", CallbackData: "create_note"}},
			{{Text: "<", CallbackData: GetNoteListCallback}, {Text: ">", CallbackData: GetNoteListCallback}},
		},
	}

	for i, note := range notes {
		if i%2 == 0 {
			raw++
			resultKeyboard.InlineKeyboard = append(resultKeyboard.InlineKeyboard, []*dto.InlineKeyboardButton{})
		}
		newButton := &dto.InlineKeyboardButton{Text: note.Name, CallbackData: GetNoteCallback + fmt.Sprintf(" %d", note.ID)}
		resultKeyboard.InlineKeyboard[raw] = append(resultKeyboard.InlineKeyboard[raw], newButton)
	}
	return resultKeyboard
}

var (
	GoBackKeyboard = &dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]*dto.InlineKeyboardButton{
			{{Text: "к заметкам", CallbackData: GetNoteListCallback + " 0"}},
		},
	}
)
