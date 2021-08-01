package setdata_telegram_store

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateTelegramUnknown = com.NewMiddleError(errors.New("could not create telegram bot:unknown error"), 500, 150)
	ErrTelegramNotFound      = com.NewMiddleError(errors.New("telegram not found"), 404, 151)
)
