package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type TelegramRequestParams struct {
	BotId           string
	Keyboards       []string
	ChatId          string
	MessageThreadId string
	Text            string
	ParseMode       string
}

func SendMessage(t *TelegramRequestParams, clnt *resty.Client) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotId)
	params := make(map[string]string)

	params["chat_id"] = t.ChatId
	params["message_thread_id"] = t.MessageThreadId
	params["text"] = t.Text
	params["parse_mode"] = t.ParseMode

	if len(t.Keyboards) != 0 {
		j, err := json.Marshal(t.Keyboards)
		if err != nil {
			fmt.Println(err)
			return
		}
		params["reply_markup"] = string(j)
	}

	res, err := clnt.R().SetQueryParams(params).Get(url)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Response Info:", res)
}
