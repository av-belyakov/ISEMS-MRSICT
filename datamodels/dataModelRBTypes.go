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

//RBookReqParameters структура описывающая параметры для  RBookReq
type RBookReqParameter struct {
	OP string `json:"op" bson:"op"`
	Vocabulary
}

//RBookReqParameters - срез структур описывающих параметры для  RBookReq
type RBookReqParameters []RBookReqParameter

// Vocabularys структура в которой могут хранятся все "открытые" и  "закрытые" словари STIX
// и пользовательские словари терминов
type Vocabularys []Vocabulary

//RBookReq - структура описывающая запрос от API к справочной информации
type RBookReq struct {
	APIRequestProcessingReqJSON
	RequestDetails []RBookReqParameter `json:"request_details" bson:"request_details"`
}

type RBookRsp struct {
}

//Sanitizer - интерфейсный тип санитаризатор
//(Осуществляет замену нежелательных символов в полях объектов)
type Sanitizer interface {
	Sanitize()
}

//Validator - интерфейсный тип Валидатор
type Validator interface {
	IsValid() (bool, error)
}

var Commands = []string{"add", "get", "remove", "replace"}
