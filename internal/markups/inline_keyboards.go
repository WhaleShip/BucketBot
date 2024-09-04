package markups

import "github.com/WhaleShip/BucketBot/dto"

func GetNotesKeyboard() dto.InlineKeyboardMarkup {
	return dto.InlineKeyboardMarkup{
		InlineKeyboard: [][]dto.InlineKeyboardButton{
			{{Text: "заметка", CallbackData: "action_1"}},
			{{Text: "📝 создать заметку", CallbackData: "create_note"}},
			{{Text: "<", CallbackData: "go_prev"}, {Text: ">", CallbackData: "go_next"}},
		},
	}
}
