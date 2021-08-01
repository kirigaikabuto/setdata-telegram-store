package setdata_telegram_store

type ChatIdStore interface {
	Create(ch *ChatId) (*ChatId, error)
	List(telegramBotId string) ([]ChatId, error)
}
