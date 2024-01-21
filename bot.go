package tgbot

import (
	"bytes"
	"fmt"
	"io"

	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func getRequestValues(request interface{}) map[string]string {
	result := map[string]string{}
	elem := reflect.ValueOf(request).Elem()
	typeDef := elem.Type()
	for i := 0; i < typeDef.NumField(); i++ {
		field := elem.Field(i)
		name := typeDef.Field(i).Tag.Get("json")
		if field.Kind() == reflect.Slice {
			sliced := []string{}
			for i := 0; i < field.Len(); i++ {
				sliced = append(sliced, fmt.Sprintf("%v", field.Index(i)))
			}
			result[name] = strings.Join(sliced, ",")
		} else {
			result[name] = fmt.Sprintf("%v", field)
		}
	}
	return result
}

type Response[T any] struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`

	Result T `json:"result"`
}

type TelegramBot interface {
	query(methodName string, params map[string]string) ([]byte, error)
}

type TelegramBotImpl struct {
	TelegramBot

	httpClient *http.Client
	token      string
}

func NewBot(token string) *TelegramBotImpl {
	return &TelegramBotImpl{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (b *TelegramBotImpl) query(apiMethod string, body map[string]string) ([]byte, error) {
	params := url.Values{}
	for k, v := range body {
		params.Set(k, v)
	}

	resp, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/%s", b.token, apiMethod), params)
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
	resultBytes, err := b.query(apiMethod, getRequestValues(request))
	if err != nil {
		return nil, err
	}

	result := &Response[T]{}
	if err = json.Unmarshal(resultBytes, result); err != nil {
		return nil, fmt.Errorf("Cannot unmarshal the response: %s", err)
	}

	if !result.Ok {
		return nil, fmt.Errorf("Request \"%s\" completed with error %d: %s",
			apiMethod, result.ErrorCode, result.Description)
	}
	return result, nil
}

func (b *TelegramBotImpl) GetUpdates(request *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	apiResponse, err := queryAndUnmarshal[[]*Update](b, "getUpdates", request)
	if err != nil {
		return nil, err
	}
	return &GetUpdatesResponse{
		Result: apiResponse.Result,
	}, nil
}
