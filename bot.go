package tgbot

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`

	Raw    []byte
	Result T `json:"result"`
}

type TelegramBot interface {
	Query(methodName string, body interface{}) ([]byte, error)
	ResolveUrl(filepath string) string
}

type TelegramBotImpl struct {
	TelegramBot

	httpClient *http.Client
	token      string
}

func NewBot(token string) TelegramBot {
	return &TelegramBotImpl{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (b *TelegramBotImpl) ResolveUrl(filepath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.token, filepath)
}

func (b *TelegramBotImpl) Query(apiMethod string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST",
		fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, apiMethod),
		bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := b.httpClient.Do(request)
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	if _, err = io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func queryAndUnmarshal[T any](b TelegramBot, apiMethod string, request interface{}) (*Response[T], error) {
	resultBytes, err := b.Query(apiMethod, request)
	if err != nil {
		return nil, err
	}

	result := &Response[T]{}
	if err = json.Unmarshal(resultBytes, result); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal the response: %s", err)
	}

	result.Raw = resultBytes
	if !result.Ok {
		return nil, fmt.Errorf("Request \"%s\" completed with error %d: %s",
			apiMethod, result.ErrorCode, result.Description)
	}
	return result, nil
}
