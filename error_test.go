package gohelper_test

import (
	"testing"

	gohelper "github.com/faridyusof727/go-helper/v2"
	"github.com/pkg/errors"
)

func TestTelegramStack(t *testing.T) {
	testErr := errors.New("testing error for telegram")

	apikey := ""
	chatId := int64(1) // can be a group or direct chat

	err := gohelper.ErrorTelegramPrint(apikey, chatId, testErr)
	if err != nil {
		t.Error(err.Error())
	}
}

