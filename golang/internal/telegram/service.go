package telegram

import (
	"fmt"
	httpclient "logger/internal/http_client"
	messagefactory "logger/internal/telegram/message_factory"
	"logger/src/middlewares"
	logmodel "logger/src/models/log_model"
)

func ReportError(l logmodel.ReportableLog) {
	text := fmt.Sprintf("%s %sError: %s File: %s Line: %s %s Url: %s Phone: %s IP: %s %s",
		messagefactory.ItalicInline("📛 Exception"),
		messagefactory.Text(l.ENV),
		messagefactory.Copy(l.Message),
		messagefactory.Copy(l.File),
		messagefactory.Copy(l.Line),
		messagefactory.Seperator(),
		messagefactory.Copy(l.Url),
		messagefactory.Copy(l.Phone),
		messagefactory.Copy(l.Ip),
		messagefactory.Italic(l.Time),
	)

	client := httpclient.NewHttpClient().Client

	user := middlewares.GetAuthUser()

	t := &TelegramRequestParams{
		BotId:           user.TelegramBotId,
		ChatId:          user.ChatId,
		Text:            text,
		ParseMode:       "html",
		Keyboards:       []string{},
		MessageThreadId: "",
	}

	SendMessage(t, client)
}
