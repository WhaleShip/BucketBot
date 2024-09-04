package dto

type Response struct {
	Chat_id int64                `json:"chat_id"`
	Text    string               `json:"text"`
	Markup  InlineKeyboardMarkup `json:"reply_markup"`
}
