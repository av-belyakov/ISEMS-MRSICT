package datamodels

import (
	"encoding/json"
	"time"
)

//ModAPIRequestProcessingCommonJSON содержит описание общих полей формата JSON для запросов и ответов модуля ModuleAPIRequestProcessing
// TaskID - ID задачи
// Section - секция, предназначена для разграничения запросов по типам обработчиков. Например, используются следующие типы обработчиков:
//  - "handling stix object" (обработка объектов STIX)
//  - "handling search requests" (обработка поисковых запросов)
//  - "handling reference book" (обработка запросов связанных со справочниками)
//  - "generating reports" (генерирование отчетов)
//  - "formation final documents" (генерирование итоговых документов)
type ModAPIRequestProcessingCommonJSON struct {
	TaskID  string `json:"task_id"`
	Section string `json:"section"`
}

//ModAPIRequestProcessingReqJSON содержит описание формата JSON запроса получаемого через модуль ModuleAPIRequestProcessing
// TaskWasGeneratedAutomatically - задача была сгенерирована автоматически (true - да)
// UserNameGeneratedTask - имя пользователя сгенерировавшего задачу
// RequestDetails - подробности запроса
type ModAPIRequestProcessingReqJSON struct {
	ModAPIRequestProcessingCommonJSON
	TaskWasGeneratedAutomatically bool             `json:"task_was_generated_automatically"`
	UserNameGeneratedTask         string           `json:"user_name_generated_task"`
	RequestDetails                *json.RawMessage `json:"request_details"`
}

//ModAPIRequestProcessingReqHandlingSTIXObjectJSON содержит список произвольных объектов STIX. Информация из данных списков добавляется в
// базу данных, если ее там нет или обновляется, если она там есть. Данный тип применяется ТОЛЬКО при Section:"handling stix object"
type ModAPIRequestProcessingReqHandlingSTIXObjectJSON []*json.RawMessage

//ModAPIRequestProcessingResJSON содержит описание формата JSON ответа, передаваемого через модуль ModuleAPIRequestProcessing
// Description - дополнительное детальное описание. При возникновении ошибки, сюда пишется максимально подробное описание ошибки.
// IsSuccessful - индикатор успешности выполнения задачи (false - неуспешно, true - успешно).
// InformationMessage - информационное сообщение
// AdditionalParameters - дополнительные параметры связанные с выполняемой задачей.
type ModAPIRequestProcessingResJSON struct {
	ModAPIRequestProcessingCommonJSON
	IsSuccessful         bool                                      `json:"is_successful"`
	Description          string                                    `json:"description"`
	InformationMessage   ModAPIRequestProcessingResJSONInfoMsgType `json:"information_message"`
	AdditionalParameters interface{}                               `json:"additional_parameters"`
}

//ModAPIRequestProcessingResJSONInfoMsgType подробное описание InformationMessage
// MsgType - тип информационного сообщения ('info', 'success', 'warning', 'danger')
// Msg - информационное сообщение
type ModAPIRequestProcessingResJSONInfoMsgType struct {
	MsgType string `json:"msg_type"`
	Msg     string `json:"msg"`
}

//CommonModAPIRequestProcessingResJSONSearchReqType содержит описание формата JSON запроса к поисковой машине
// CollectionName - наименование коллекции документов. Для поиска STIX объектов collection_name = "stix object"
// SearchParameters - параметры поиска (для разных коллекций параметры поиска отличаются)
type CommonModAPIRequestProcessingResJSONSearchReqType struct {
	CollectionName   string           `json:"collection_name"`
	SearchParameters *json.RawMessage `json:"search_parameters"`
}

//ModAPIRequestProcessingResJSONSearchReqType содержит описание формата JSON запроса к поисковой машине
// CollectionName - наименование коллекции документов. Для поиска STIX объектов collection_name = "stix object"
// PaginateParameters - параметры разбиения на страницы
//  MaxPartNum - размер части, то есть максимальное количество найденных элементов, которое может содержаться в одном ответе
//	CurrentPartNumber - номер текущей части (0 или 1 считаются за первую часть)
// SearchParameters - параметры поиска (для разных коллекций параметры поиска отличаются)
type ModAPIRequestProcessingResJSONSearchReqType struct {
	CollectionName     string `json:"collection_name"`
	PaginateParameters struct {
		MaxPartNum        int `json:"max_part_num"`
		CurrentPartNumber int `json:"current_part_number"`
	} `json:"paginate_parameters"`
	SearchParameters interface{} `json:"search_parameters"`
}

//SearchThroughCollectionSTIXObjectsType содержит описание формата JSON запроса для поиска информации о STIX объектах
// DocumentsID - список идентификаторов документов STIX (если данный параметр СОДЕРЖИТ список идентификаторов то остальные параметры
//  НЕ УЧИТЫВАЮТСЯ в ходе выполнения поискового запроса)
// DocumentsType - список наименованиq типjd документов STIX объектов, название которых содержится в поле "type" любого STIX документа
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (только для DO или RO STIX)
// Modified - время модификации объекта, в формате "2016-05-12T08:17:27.000Z" (только для DO или RO STIX)
// CreatedByRef - содержит идентификатор источника создавшего данный объект (только для DO STIX)
// SpecificSearchFields - содержит список специфичных полй объектов
type SearchThroughCollectionSTIXObjectsType struct {
	DocumentsID   []string `json:"documents_id"`
	DocumentsType []string `json:"documents_type"`
	Created       struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"created"`
	Modified struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"modified"`
	CreatedByRef         string                               `json:"created_by_ref"`
	SpecificSearchFields []SpecificSearchFieldsSTIXObjectType `json:"specific_search_fields"`
}

//SpecificSearchFieldsSTIXObjectType содержит набор полей для поиска, которые являются специфичными для объектов STIX
// ObjectName - содержит наименование STIX объекта
// SearchFields - содержит перечень уникальных полей соответствующих определенному STIX объекту, внутри данного объекта действует логика "И"
type SpecificSearchFieldsSTIXObjectType struct {
	ObjectName   string                     `json:"object_name"`
	SearchFields SearchFieldsSTIXObjectType `json:"search_fields"`
}

//SearchFieldsSTIXObjectType содержит перечень полей STIX объекта по которым выполняется поиск, не все поля для каждого STIX объекта будут заполнятся
// У некоторых STIX объектов может не хватать одного или более полей
// Name - имя используемое для идентификации типа STIX объекта
// Aliases - альтернативные имена используемые для идентификации типа STIX объекта
// FirstSeen - интервал времени когда сущность была обнаружена впервые
// LastSeen - интервал времени когда сущность была обнаружена в последний раз
// Roles - список ролей для идентификации действий
// Country - наименование страны
// City - наименование города
// URL - унифицированный указатель ресурса
// Number - номер для идентификации
// Value - список каких либо значений
type SearchFieldsSTIXObjectType struct {
	Name      string   `json:"name"`
	Aliases   []string `json:"aliases"`
	FirstSeen struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"first_seen"`
	LastSeen struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	} `json:"last_seen"`
	Roles   []string `json:"roles"`
	Country string   `json:"country"`
	City    string   `json:"city"`
	URL     string   `json:"url"`
	Number  int      `json:"number"`
	Value   []string `json:"value"`
}

/*
	Example!!! для передачи информации частями
//TransmittingRequestedInformationModAPIRequestProcessingResJSON содержит информацию передаваемую в ответ на запрашиваемые данные
// TotalNumberParts - общее количество частей
// GivenSizePart - заданный размер части
// NumberTransmittedPart - номер передаваемой части
// TransmittedData - передаваемые данные
type TransmittingRequestedInformationModAPIRequestProcessingResJSON struct {
	TotalNumberParts      int         `json:"total_number_parts"`
	GivenSizePart         int         `json:"given_size_part"`
	NumberTransmittedPart int         `json:"number_transmitted_part"`
	TransmittedData       interface{} `json:"transmitted_data"` // пока 'interface{}' так как не знаю тип передаваемых данных
}
*/
