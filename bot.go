package tgbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	ResolveFileUrl(filename string) *url.URL
}

type TelegramBotImpl struct {
	TelegramBot

	httpClient *http.Client
	server     *url.URL
	token      string
}

func NewCustomBot(server string, token string) (TelegramBot, error) {
	url, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	return &TelegramBotImpl{
		token:      token,
		server:     url,
		httpClient: &http.Client{},
	}, nil

}

func NewBot(token string) (TelegramBot, error) {
	return NewCustomBot("https://api.telegram.org", token)
}

func (b *TelegramBotImpl) ResolveFileUrl(filename string) *url.URL {
	return b.server.JoinPath("file", "bot"+b.token, filename)
}

func (b *TelegramBotImpl) Query(apiMethod string, body interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	methodURL := b.server.JoinPath("bot"+b.token, apiMethod)
	request, err := http.NewRequest("POST", methodURL.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := b.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
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
