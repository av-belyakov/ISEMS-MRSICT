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
// PaginateParameters - параметры разбиения на страницы
//  MaxPartNum - размер части, то есть максимальное количество найденных элементов, которое может содержаться в одном ответе
//	CurrentPartNumber - номер текущей части (0 или 1 считаются за первую часть)
// SortParameters - параметр в котором можно указать значение, по которому будет выполнятся сортировка полей. Параметр должен содержать
//  одно из следующих значений: "document_type", "data_created", "data_modified", "data_first_seen", "data_last_seen", "ipv4", "ipv6",
//  "country". Если данное поле пустое, то сортировка будет выполнятся по свойству ObjectId БД MongoDB
// SearchParameters - параметры поиска (для разных коллекций параметры поиска отличаются)
type CommonModAPIRequestProcessingResJSONSearchReqType struct {
	CollectionName     string `json:"collection_name"`
	PaginateParameters struct {
		MaxPartNum        int `json:"max_part_num"`
		CurrentPartNumber int `json:"current_part_number"`
	} `json:"paginate_parameters"`
	SortParameters   string           `json:"sort_parameters"`
	SearchParameters *json.RawMessage `json:"search_parameters"`
}

//ModAPIRequestProcessingResJSONSearchReqType содержит описание формата JSON запроса к поисковой машине
// CollectionName - наименование коллекции документов. Для поиска STIX объектов collection_name = "stix object"
// PaginateParameters - параметры разбиения на страницы
//  MaxPartNum - размер части, то есть максимальное количество найденных элементов, которое может содержаться в одном ответе
//	CurrentPartNumber - номер текущей части (0 или 1 считаются за первую часть)
// SortableField - параметр в котором можно указать значение, по которому будет выполнятся сортировка полей. Параметр должен содержать
//  одно из следующих значений: "document_type", "data_created", "data_modified", "data_first_seen", "data_last_seen", "ipv4", "ipv6",
//  "country". Если данное поле пустое, то сортировка будет выполнятся по свойству ObjectId БД MongoDB
// SearchParameters - параметры поиска (для разных коллекций параметры поиска отличаются)
type ModAPIRequestProcessingResJSONSearchReqType struct {
	CollectionName     string `json:"collection_name"`
	PaginateParameters struct {
		MaxPartNum        int `json:"max_part_num"`
		CurrentPartNumber int `json:"current_part_number"`
	} `json:"paginate_parameters"`
	SortableField    string      `json:"sortable field"`
	SearchParameters interface{} `json:"search_parameters"`
}

//SearchThroughCollectionSTIXObjectsType содержит описание формата JSON запроса для поиска информации о STIX объектах
// DocumentsID - список идентификаторов документов STIX (если данный параметр СОДЕРЖИТ список идентификаторов то остальные параметры
//  НЕ УЧИТЫВАЮТСЯ в ходе выполнения поискового запроса)
// DocumentsType - список наименованиq типjd документов STIX объектов, название которых содержится в поле "type" любого STIX документа
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (только для DO или RO STIX)
// Modified - время модификации объекта, в формате "2016-05-12T08:17:27.000Z" (только для DO или RO STIX)
// CreatedByRef - содержит идентификатор источника создавшего данный объект (только для DO STIX)
// SpecificSearchFields - содержит список специфичных полей объектов (при этом используется логика "ИЛИ" между объектами предназначенными для поиска)
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

//SpecificSearchFieldsSTIXObjectType содержит набор полей для поиска, которые являются специфичными для объектов STIX (между полями используется
//  логика "И")
// Name - имя используемое для идентификации типа STIX объекта
// Aliases - альтернативные имена используемые для идентификации типа STIX объекта
// FirstSeen - интервал времени когда сущность была обнаружена впервые
// LastSeen - интервал времени когда сущность была обнаружена в последний раз
// Roles - список ролей для идентификации действий
// Country - наименование страны
// City - наименование города
// URL - унифицированный указатель ресурса
// NumberAutonomousSystem - номер для идентификации автономной системы
// Value - может содержать какое либо из следующих типов значений: "domain-name", "email-addr", "ipv4-addr", "ipv6-addr" или "url". Или все эти
//  значения в перемешку. Между значениями в поле 'Value' используется лиогика "ИЛИ".
type SpecificSearchFieldsSTIXObjectType struct {
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
	Roles                  []string `json:"roles"`
	Country                string   `json:"country"`
	City                   string   `json:"city"`
	NumberAutonomousSystem int      `json:"number_autonomous_system"`
	Value                  []string `json:"value"`
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
