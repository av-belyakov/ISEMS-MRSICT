package datamodels

import (
	"encoding/json"
)

//HandlerSTIXObject интерфейс реализующий обработчики для STIX объектов
type HandlerSTIXObject interface {
	ValidatorJSONObject
	DecoderJSONObject
	EncoderJSONObject
}

//ValidatorJSONObject интерфейс реализующий обработчик для валидации STIX объектов
type ValidatorJSONObject interface {
	CheckingTypeFields() bool
}

//DecoderJSONObject интерфейс реализующий обработчик для декодирования JSON объекта в STIX объект
type DecoderJSONObject interface {
	DecoderJSON(*json.RawMessage) (interface{}, error)
}

//EncoderJSONObject интерфейс реализующий обработчик для кодирования STIX объекта в JSON объект
type EncoderJSONObject interface {
	EncoderJSON(interface{}) (*[]byte, error)
}

//ListSTIXObject может содержать любой из STIX объектов с указанием его типа
// DataType - тип STIX объекта
// Data - непосредственно сам STIX объект
type ListSTIXObject struct {
	DataType string
	Data     HandlerSTIXObject
	//	Data     interface{}
}
