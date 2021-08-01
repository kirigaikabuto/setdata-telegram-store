package setdata_telegram_store

type CreateTelegramBotCommand struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

func (cmd *CreateTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).CreateTelegramBot(cmd)
}

type GetTelegramBotCommand struct {
	Id string `json:"id"`
}

func (cmd *GetTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).GetTelegramBot(cmd)
}

type DeleteTelegramBotCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(TelegramService).DeleteTelegramBot(cmd)
}

type ListTelegramBotCommand struct {
}

func (cmd *ListTelegramBotCommand) Exec(service interface{}) (interface{}, error) {
	return service.(TelegramService).ListTelegramBot(cmd)
}
