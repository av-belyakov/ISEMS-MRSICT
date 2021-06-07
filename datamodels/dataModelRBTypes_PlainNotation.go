package datamodels

//Типы данных для обработки JSON-документа находящегося в плоской нотации, типа JSON-Patch

// RefernceBookReq
// Основная структура полей запроса на обработку справочника
type CommonRefernceBookReq struct {
	OP string `json:"op" bson:"op"`
	//	parameters json.RawText
}

type NestedMapType map[string]NestedWorker

type NestedMapIntType map[int]NestedWorker

type NestedSliceType []NestedWorker

type RefernceBookGetReq struct {
	CommonRefernceBookReq
	parameters map[string]NestedSliceType `json:"parameters" bson:"parameters"`
}

type RefernceBookReplaceReq struct {
	CommonRefernceBookReq
	parameters map[string]NestedSliceType
}

type RefernceBookRemoveReq struct {
	CommonRefernceBookReq
	parameters map[string]NestedSliceType
}

type RefernceBookAddReq struct {
	CommonRefernceBookReq
	paremeters map[string]NestedSliceType
}

type NestedBoolType bool

type NestedStringType string

type NestedDigitType float64

type NestedWorker interface {
	GetValue() interface{}
}


type ReferencesBookReq1 struct {
	APIRequestProcessingReqJSON
	RequestDetails RequestDetailsRefBookType `json:"request_details" bson:"request_details"`
}

type RequestDetailsRefBookType struct {
	CommonRefernceBookReq
	Parameters NestedMapType `json:"parameters" bson:"parameters"`
}

//RefernceBooker
// Интерфейс для для работы c запросами справочной информации
// RefernceBook JSON API
type RefernceBooker interface {
	GetParameters()
	GetCommand() string
	Pack() []byte
	Unpack(interface{})
}
