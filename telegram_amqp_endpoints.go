package setdata_telegram_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type TelegramAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func (t *TelegramAmqpEndpoints) CreateTelegramBotAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateTelegramBotCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := t.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (t *TelegramAmqpEndpoints) ListTelegramBotAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListTelegramBotCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := t.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (t *TelegramAmqpEndpoints) GetTelegramBotAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetTelegramBotCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := t.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (t *TelegramAmqpEndpoints) DeleteTelegramBotAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteTelegramBotCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := t.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}
