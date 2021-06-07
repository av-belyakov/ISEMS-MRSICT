package datamodels

/********** 			STIX Vocabulary  			**********/

// Словари STIX имена типов  которых заканчиваются на '-ov' предоставляют список общепринятых
// в отрасли терминов в качестве руководства для пользователя, данные списки терминов не являются открытыми ( 'open' ).
// Это значит что и пользователь при определении значения поля объекта STIX не ограниченможет только значениями
// закрепленными в данном словаре но так же может дополнять его своими терминами и использовать и применять их.
// Если в описании полей STIX-объектов присутствует тип open-vocab это означает что пользователю при заполнеии данного поля
// следует применять занчения одного из словарей данного типа.

// Словари STIX, имена типов которых заканчиваются на '-enum', являются "закрытыми". В описании полей STIX-объектов
// данные словари обозначаются типом "enum". Пользователю при инициализации значения такого поля в STIX-объекте должен
// использовать только значения закрепленные в данном словаре.

// Ниже описаны структуры данных описывающие словари STIX
// названия которых применяются в какчестве значений полей STIX-объектов

// VocabularyElement структура описывающая элемент словаря STIX
type VocElement struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty" bson:"shortDescription,omitempty"`
	FullDescription  string `json:"fullDescription,omitempty" bson:"fullDescription,omitempty"`
}

type VocElements []VocElement

// Vocabulary структура описывающая словарь терминов (либо пользовательский)
type Vocabulary struct {
	VocElement
	Elements VocElements `json:"elements,omitempty" bson:"elements,omitempty"`
}

// Vocabularys структура в которой могут хранятся все "открытые" и  "закрытые" словари STIX
// и пользовательские словари терминов
type Vocabularys []Vocabulary

//APIRequestProcessingReqJSON содержит описание формата JSON запроса получаемого через модуль ModuleAPIRequestProcessing
// моя версия
type APIRequestProcessingReqJSON struct {
	ModAPIRequestProcessingCommonJSON
	TaskWasGeneratedAutomatically bool   `json:"task_was_generated_automatically"`
	UserNameGeneratedTask         string `json:"user_name_generated_task"`
}

////////////////////////////////////////////////////////////////////////////////////
//Ниже приводятся структы описывающие поля запроса от  API к справочной информации
////////////////////////////////////////////////////////////////////////////////////

//RBookReqParameters структура описывающая параметры для  RBookReq
type RBookReqParameter struct {
	OP string `json:"op" bson:"op"`
	Vocabulary
}

//RBookReqParameters - срез структур описывающих параметры для  RBookReq
type RBookReqParameters []RBookReqParameter

//RBookReq - структура описывающая запрос от API к справочной информации
//   Планировалась для Plain Notation
//type RBookReq struct {
//	APIRequestProcessingReqJSON
//	RequestDetails []RBookReqParameter `json:"request_details" bson:"request_details"`
//}

//////////////////////////////////////////////////////////////////////////////////////
//Ниже приводятся структы описывающие поля ответа обработки запросов к Referencr Book
//////////////////////////////////////////////////////////////////////////////////////

//RBookRespParameter - структура описывающая ответ содержащий результаты обработки запроса к справочной информации
type RBookRespParameter struct {
	OP           string `json:"op" bson:"op"`
	IsSuccessful bool   `json:"is_succsessful" bson:"is_succsessful"`
	Description  string `json:"description" bson:"description"`
	Vocabulary
}

//RBookRespParameters - срез структур описывающих параметры для  RBookResp
type RBookRespParameters []RBookRespParameter

//RBookResp - структура описывающая запрос от API к справочной информации

//type RBookResp struct {
//	APIRequestProcessingReqJSON
//	RequestDetails []RBookReqParameter `json:"request_details" bson:"request_details"`
//}

//////////////////////////////////////////////////////////////////////////////////////
//Ниже приводятся интерфейсы
//////////////////////////////////////////////////////////////////////////////////////

//Sanitizer - интерфейсный тип санитаризатор
//(Осуществляет замену нежелательных символов в полях объектов)
type Sanitizer interface {
	Sanitize()
}

//Validator - интерфейсный тип Валидатор
type Validator interface {
	IsValid() (bool, error)
}

// Тип для хранения списка команд применяемых для работы со словарями
type VocOperations []string

var CommandSet VocOperations = VocOperations{"add", "get", "remove", "replace"}
