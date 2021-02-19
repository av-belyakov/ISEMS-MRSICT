package datamodels

import (
	"time"
)

/**********			 Некоторые примитивные типы STIX			 **********/

//EnumTypeSTIX тип "enum", по терминологии STIX, является жестко заданным списком терминов, который представлен в виде строки
type EnumTypeSTIX string

//ExternalReferencesTypeSTIX тип "external-reference", по терминалогии STIX, является списком с информацией о внешних ссылках не относящихся к STIX информации
type ExternalReferencesTypeSTIX []*ExternalReferenceTypeElementSTIX

//ExternalReferenceTypeElementSTIX тип содержащий подробную информацию о внешних ссылках, таких как URL, ID и т.д.
// SourceName - имя источника (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - описание
// URL - URL ссылка на внешних источниках
// Hashes - содержит словарь хэшей для содержимого URL-адреса. Это ДОЛЖНО быть предусмотрено при наличии свойства url
// ExternalID - идентификатор на внешних источниках
type ExternalReferenceTypeElementSTIX struct {
	SourceName  string         `json:"source_name" bson:"source_name"`
	Description string         `json:"description" bson:"description"`
	URL         string         `json:"url" bson:"url"`
	Hashes      HashesTypeSTIX `json:"hashes" bson:"hashes"`
	ExternalID  string         `json:"external_id" bson:"external_id"`
}

//HashesTypeSTIX тип "hashes", по терминологии STIX, содержащий хеш значения, где <тип_хеша>:<хеш>
type HashesTypeSTIX map[string]string

//IdentifierTypeSTIX тип "identifier", по терминалогии STIX, содержащий уникальный идентификатор UUID, преимущественно версии 4 при этом ID должен
//начинаться с наименования организации или программного обеспечения сгенерировавшего его. Например, <example-source--ff26c055-6336-5bc5-b98d-13d6226742dd> (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type IdentifierTypeSTIX string

//KillChainPhasesTypeSTIX тип "kill-chain-phases", по терминалогии STIX, содержащий цепочки фактов, приведших к какому либо урону
type KillChainPhasesTypeSTIX []*KillChainPhasesTypeElementSTIX

//KillChainPhasesTypeElementSTIX тип содержащий набор элементов цепочки фактов, приведших к какому либо урону
// KillChainName - имя цепочки (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PhaseName - наименование фазы из спецификации STIX, например, "reconnaissance", "pre-attack" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type KillChainPhasesTypeElementSTIX struct {
	KillChainName string `json:"kill_chain_name" bson:"kill_chain_name"`
	PhaseName     string `json:"phase_name" bson:"phase_name"`
}

//OpenVocabTypeSTIX тип "open-vocab", по терминалогии STIX, содержащий заранее определенное (предложенное) значение
type OpenVocabTypeSTIX string

/********** 			Meta Object STIX 			**********/

/***	 			Language Content STIX 				***/

//LanguageContentTypeSTIX тип "language content", по терминалогии STIX, представляет собой текстовое содержимое для
// объектов STIX, представленных на языках, отличных от языка исходного объекта
// Type - наименование типа объекта, для этого типа это поле ДОЛЖНО содержать "language-content" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Modified - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRef - определяет идентификатор объекта "language content" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectModified - время изменения объекта, к которому применяется данное значение, в формате "2016-05-12T08:17:27.000Z"
// Contents - (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// CreatedByRef - содержит идентификатор источника создавшего данный объект
// Revoked - вернуть к текущему состоянию
// Labels - определяет набор терминов, используемых для описания данного объекта
// Сonfidence - определяет уверенность создателя в правильности своих данных. Доверительное значение ДОЛЖНО быть числом
//  в диапазоне 0-100. Если 0 - значение не определено.
// ExternalReferences - список внешних ссылок не относящихся к STIX информации
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
type LanguageContentTypeSTIX struct {
	Type               string                     `json:"type" bson:"type"`
	ID                 string                     `json:"id" bson:"id"`
	SpecVersion        string                     `json:"spec_version" bson:"spec_version"`
	Created            time.Time                  `json:"created" bson:"created"`
	Modified           time.Time                  `json:"modified" bson:"modified"`
	ObjectRef          IdentifierTypeSTIX         `json:"object_ref" bson:"object_ref"`
	ObjectModified     time.Time                  `json:"object_modified" bson:"object_modified"`
	Contents           map[string]string          `json:"contents" bson:"contents"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	Revoked            bool                       `json:"revoked" bson:"revoked"`
	Labels             []string                   `json:"labels" bson:"labels"`
	Сonfidence         int                        `json:"confidence" bson:"confidence"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
}

/***	 			Data Markings STIX 				***/

//CommonDataMarkingsTypeSTIX общие свойства меток данных
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type CommonDataMarkingsTypeSTIX struct {
	SpecVersion string    `json:"spec_version" bson:"spec_version"`
	ID          string    `json:"id" bson:"id"`
	Created     time.Time `json:"created" bson:"created"`
}

//GranularMarkingsTypeSTIX тип "granular_markings", по терминалогии STIX, представляет собой набор маркеров ссылающихся на свойства "marking_ref" и "lang"
// Lang - идентифицирует язык соответствующим маркером
// MarkingRef - определяет идентификатор объекта "marking-definition"
// Selectors - определяет список селекторов для содержимого объекта STIX, к которому применяется это свойство
type GranularMarkingsTypeSTIX struct {
	Lang       string             `json:"lang" bson:"lang"`
	MarkingRef IdentifierTypeSTIX `json:"marking_ref" bson:"marking_ref"`
	Selectors  []string           `json:"selectors" bson:"selectors"`
}

//MarkingDefinitionObjectSTIX объект "marking-definition", по терминалогии STIX, содержит метки данных ссылающиеся на требования к обработке
// или совместному использованию данных
// Type - наименование типа объекта, для этого типа это поле ДОЛЖНО содержать "marking-definition" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - наименование маркера
// DefinitionType - определение типа, должно содержать "statement" или "tlp"
// Definition - содержит маркер в виде объекта вида или { "statement": "Copyright 2019, Example Corp" } или { "tlp": "white" }, где
//  "white" может быть заменен на "green", "amber", "red"
// CreatedByRef - содержит идентификатор источника создавшего данный объект
// ExternalReferences - список внешних ссылок не относящихся к STIX информации
// ObjectMarkingRefs - определяет список свойств идентификаторов объектов определения маркировки, которые применяются к этому объекту
//  хотя оно и является списком типа IdentifierTypeSTIX, но тот в свою очередь ССЫЛАЕТСЯ на объект типа MarkingDefinitionObjectSTIX (marking-definition)
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
type MarkingDefinitionObjectSTIX struct {
	CommonDataMarkingsTypeSTIX
	Type               string                     `json:"type" bson:"type"`
	Name               string                     `json:"name" bson:"name"`
	DefinitionType     string                     `json:"definition_type" bson:"definition_type"`
	Definition         map[string]string          `json:"definition" bson:"definition"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
}

/********** 			Bundle Object STIX 			**********/

//BundleObjectSTIX объект "bundle", по терминалогии STIX, содержит коллекцию произвольных STIX объектов сгруппированных вместе в единый контейнер. Подобная Связка
// не имеет никакого семантического значения, и объекты, содержащиеся в Связке, не считаются связанными в силу того, что они находятся в одной Связке.
// Type - наименование типа объекта, для этого типа это поле ДОЛЖНО содержать "bundle" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Objects - содержит список любых STIX объектов
type BundleObjectSTIX struct {
	Type    string        `json:"" bson:""`
	ID      string        `json:"" bson:""`
	Objects []interface{} `json:"objects" bson:"objects"`
}
