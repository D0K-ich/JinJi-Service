package messages

import (
	"encoding/json"
	"go.uber.org/zap"
)

func ReMarshalMust[TO any](in interface{}) (out TO) {
	var err error
	if out, err = ReMarshal[TO](in); err != nil {
		log.Error("Failed to remarshall in must func", zap.Any("err", err))
		return
	}
	return
}

func ReMarshal[TO any](in interface{}) (out TO, err error) {
	var body []byte
	if body, err = json.Marshal(in); err != nil {
		log.Error("Failed to marshal data", zap.Any("err", err))
		return
	}
	if err = json.Unmarshal(body, &out); err != nil {
		log.Error("Failed to unmarshal data", zap.Any("err", err))
		return
	}
	return
}

func ToMapStringMust(structure any) (out map[string]any) {
	var err error
	if out, err = ToMapStringAny(structure); err != nil {
		log.Error("Failed to ToMapStringMust", zap.Any("err", err))
		return
	}
	return
}

func ToMapStringAny(structure any) (out map[string]any, err error) {
	if out, err = ReMarshal[map[string]any](structure); err != nil {return}
	return
}
