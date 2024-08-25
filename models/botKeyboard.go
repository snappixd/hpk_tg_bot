package models

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var NumericKeyboard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(
		tu.InlineKeyboardButton("КІ-211"),
		tu.InlineKeyboardButton("КІ-212"),
		tu.InlineKeyboardButton("КІ-213"),
	),
)
