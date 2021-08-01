package setdata_telegram_store

import "github.com/google/uuid"

type TelegramService interface {
	CreateTelegramBot(cmd *CreateTelegramBotCommand) (*TelegramBot, error)
	GetTelegramBot(cmd *GetTelegramBotCommand) (*TelegramBot, error)
	ListTelegramBot(cmd *ListTelegramBotCommand) ([]TelegramBot, error)
	DeleteTelegramBot(cmd *DeleteTelegramBotCommand) error
}

type telegramService struct {
	store TelegramStore
}

func (t *telegramService) CreateTelegramBot(cmd *CreateTelegramBotCommand) (*TelegramBot, error) {
	telegramBot := &TelegramBot{Id: uuid.New().String()}
	telegramBot.Name = cmd.Name
	telegramBot.AccessToken = cmd.AccessToken
	return t.store.Create(telegramBot)
}

func (t *telegramService) GetTelegramBot(cmd *GetTelegramBotCommand) (*TelegramBot, error) {
	return t.store.Get(cmd.Id)
}

func (t *telegramService) ListTelegramBot(cmd *ListTelegramBotCommand) ([]TelegramBot, error) {
	return t.store.List()
}

func (t *telegramService) DeleteTelegramBot(cmd *DeleteTelegramBotCommand) error {
	return t.store.Delete(cmd.Id)
}
