package gohelper

import (
	"fmt"

	"github.com/faridyusof727/go-helper/v2/lib"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func ErrorTelegramPrint(apikey string, chatId int64, errorStack error) error {
	bot := lib.GetTelegramInstance(apikey)
	err, ok := errorStack.(stackTracer)

	if !ok {
		return errors.New("errors doesn't implement stackTracer ref: github.com/pkg/errors")
	}

	stack := err.StackTrace()

	msg := tgbotapi.NewMessage(chatId, "")
	msg.Text = fmt.Sprintf("<b>%s</b>\n<pre>%+v</pre>", errorStack.Error(), stack)
	msg.ParseMode = "html"

	if _, err := bot.Send(msg); err != nil {
		return err
	}

	return nil
}
