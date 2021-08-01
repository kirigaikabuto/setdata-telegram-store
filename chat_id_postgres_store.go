package setdata_telegram_store

import (
	"database/sql"
	"log"
	"strconv"
)

var chatIdPostgresQueries = []string{
	`create table if not exists chat_ids(
		id text,
		telegram_bot_id text,
		value text,
		primary key(id)
	);`,
}

type chatIdStore struct {
	db *sql.DB
}

func NewPostgresChatIdStore(cfg PostgresConfig) (ChatIdStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range chatIdPostgresQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &chatIdStore{db: db}
	return store, nil
}

func (c *chatIdStore) Create(ch *ChatId) (*ChatId, error) {
	query := "insert into chat_ids (id, telegram_bot_id, value) values ($1, $2, $3)"
	result, err := c.db.Exec(query, ch.Id, ch.TelegramBotId, ch.Value)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateTelegramUnknown
	}
	return ch, nil
}

func (c *chatIdStore) List(telegramBotId string) ([]ChatId, error) {
	items := []ChatId{}
	query := "select id, telegram_bot_id, value from chat_ids "
	var values []interface{}
	cnt := 1
	if telegramBotId != "" {
		query = query + "where telegram_bot_id = $" + strconv.Itoa(cnt)
		values = append(values, telegramBotId)
		cnt += 1
	}
	rows, err := c.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		item := ChatId{}
		err = rows.Scan(&item.Id, &item.TelegramBotId, &item.Value)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
