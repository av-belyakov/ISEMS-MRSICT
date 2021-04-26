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

// VocabularyWorker
// Интерфейс для работы со словарями терминов STIX
/* type VocabulariesWorker interface{
	getElement(elementName string ) (ElementDescription)
	getFullElementDecription(elementName string ) (string)
	getShortElementDecription(elementName string ) (string)
	getShortDecription(elementName string ) (string)
	getFullDecription(elementName string ) (string)
}
*/

// VocabularyElement структура описывающая элемент словаря STIX
type VocabularyElement struct {
	Name             string `json:"name" bson:"name"`
	ShortDescription string `json:"shortDescription" bson:"shortDescription"`
	FullDescription  string `json:"fullDescription" bson:"fullDescription"`
}

// Vocabulary структура описывающая словарь терминов (либо пользовательский)
type Vocabulary struct {
	VocabularyElement
	Elements []VocabularyElement `json:"elements" bson:"elements"`
}

//ReferencesBookReqParameters структура описывающая параметр для  ReferencesBookReq
type ReferencesBookReqParameter struct {
	OP string `json:"op" bson:"op"`
	Vocabulary
}

// Vocabularys структура в которой могут хранятся все "открытые" и  "закрытые" словари STIX
// и пользовательские словари терминов
type Vocabularys []Vocabulary

//ReferencesBookReq - структура описывающая запрос от API к справочной информации
type ReferencesBookReq struct {
	APIRequestProcessingReqJSON
	RequestDetails []ReferencesBookReqParameter `json:"request_details" bson:"request_details"`
}

type ReferencesBookRsp struct {
}
