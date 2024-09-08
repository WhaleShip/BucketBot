package markups

import (
	"fmt"

	"github.com/WhaleShip/BucketBot/dto"
	"github.com/WhaleShip/BucketBot/internal/database/models"
)

func GetNotesKeyboard(notes []models.Note, offset int) *dto.InlineKeyboardMarkup {
	var prevPageCallbackData string
	raw := 1
	if offset-NotesPerScreenCount < 0 {
		prevPageCallbackData = NoPageCallback
	} else {
		prevPageCallbackData = GetNoteListCallback + fmt.Sprintf(" %d", offset-NotesPerScreenCount)
	}

	nextPageCallbackData := GetNoteListCallback + fmt.Sprintf(" %d", offset+NotesPerScreenCount)
	resultKeyboard := &dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]*dto.InlineKeyboardButton{
			{{Text: "📝 создать заметку", CallbackData: CreateNoteCallback}},
			{{Text: "<", CallbackData: prevPageCallbackData}, {Text: ">", CallbackData: nextPageCallbackData}},
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

func GetNoteGoBackKeyboard(noteID int) *dto.InlineKeyboardMarkup {
	nodeGoBackKeyboard := &dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]*dto.InlineKeyboardButton{
			{{Text: "удалить заметку", CallbackData: DeleteNote + fmt.Sprintf(" %d", noteID)}},
			{{Text: "к заметкам", CallbackData: GetNoteListCallback + " 0"}},
		},
	}
	return nodeGoBackKeyboard
}

var (
	GoBackKeyboard = &dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]*dto.InlineKeyboardButton{
			{{Text: "к заметкам", CallbackData: GetNoteListCallback + " 0"}},
		},
	}
)
