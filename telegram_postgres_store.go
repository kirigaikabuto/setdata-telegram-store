package setdata_telegram_store

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var telegramPostgresQueries = []string{
	`create table if not exists telegram(
		id text,
		name text,
		access_token text,
		primary key(id)
	);`,
}

type telegramStore struct {
	db *sql.DB
}

func NewPostgresTelegramStore(cfg PostgresConfig) (TelegramStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range telegramPostgresQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &telegramStore{db: db}
	//create default values
	store.Create(&TelegramBot{
		Id:          "123456789",
		Name:        "main_telegram",
		AccessToken: "1930340256:AAFE0Hwv0cA6o5DJviO66DkfJa5JcPHUChY",
	})
	return store, nil
}

func (t *telegramStore) Create(tel *TelegramBot) (*TelegramBot, error) {
	query := "insert into telegram (id, name, access_token) values ($1, $2, $3)"
	result, err := t.db.Exec(query, tel.Id, tel.Name, tel.AccessToken)
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
	return tel, nil
}

func (t *telegramStore) Get(id string) (*TelegramBot, error) {
	telegramBot := &TelegramBot{}
	query := "select id, name, access_token from telegram where id = $1 limit 1"
	err := t.db.QueryRow(query, id).Scan(&telegramBot.Id, &telegramBot.Name, &telegramBot.AccessToken)
	if err == sql.ErrNoRows {
		return nil, ErrTelegramNotFound
	} else if err != nil {
		return nil, err
	}
	return telegramBot, nil
}

func (t *telegramStore) List() ([]TelegramBot, error) {
	telegramBots := []TelegramBot{}
	query := "select id, name, access_token from telegram "
	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := TelegramBot{}
		err = rows.Scan(&item.Id, &item.Name, &item.AccessToken)
		if err != nil {
			return nil, err
		}
		telegramBots = append(telegramBots, item)
	}
	return telegramBots, nil
}

func (t *telegramStore) Delete(id string) error {
	query := "delete from telegram where id = $1"
	result, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrTelegramNotFound
	}
	return nil
}
