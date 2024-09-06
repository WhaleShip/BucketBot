package dto

type ResponseMessage struct {
	Chat_id int                  `json:"chat_id"`
	Text    string               `json:"text"`
	Markup  InlineKeyboardMarkup `json:"reply_markup"`
}

type ResponseEditMessage struct {
	Message_id int64 `json:"message_id"`
	*ResponseMessage
}
