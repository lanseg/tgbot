package telegram

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"
	"net/url"
)

type Response[T any] struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`

	Result T `json:"result"`
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

func (b *TelegramBotImpl) queryApi(apiMethod string, params url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, apiMethod), params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func queryAndUnmarshal[T any](b *TelegramBotImpl, apiMethod string, params url.Values) (T, error) {
	var zero T
	resultBytes, err := b.queryApi(apiMethod, params)
	if err != nil {
		return zero, err
	}
	response := &Response[T]{}
	if err = json.Unmarshal(resultBytes, response); err != nil {
		return zero, fmt.Errorf("Cannot unmarshal the response: %s", err)
	}
	if !response.Ok {
		return zero, fmt.Errorf("Request \"%s\" completed with error %d: %s",
			apiMethod, response.ErrorCode, response.Description)
	}
	return response.Result, nil
}

func (b *TelegramBotImpl) GetUpdates(request *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	params := url.Values{}
	params.Set("offset", fmt.Sprintf("%d", request.Offset))
	params.Set("limit", fmt.Sprintf("%d", request.Limit))
	params.Set("timeout", fmt.Sprintf("%d", request.Timeout))
	for _, upd := range request.AllowedUpdates {
		params.Add("allowed_updates", upd)
	}

	result, err := queryAndUnmarshal[[]*Update](b, "getUpdates", params)
	if err != nil {
		return nil, err
	}
	return &GetUpdatesResponse{
		Result: result,
	}, nil
}