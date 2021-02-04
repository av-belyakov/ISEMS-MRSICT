package datamodels

import (
	"time"
)

/**********			 Некоторые примитивные типы STIX			 **********/

//ExternalReferencesTypeSTIX тип со списком о внешних ссылках не относящихся к STIX информации
type ExternalReferencesTypeSTIX []*ExternalReferenceTypeElementSTIX //`json:"external-reference" bson:"external-reference"`

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

//HashesTypeSTIX тип содержащий хеш хначения, где <тип_хеша>:<хеш>
type HashesTypeSTIX map[string]string

//IdentifierTypeSTIX тип содержащий уникальный идентификатор UUID, преимущественно версии 4
// Type - тип идентификатора, берется из пространства имен спецификации STIX или формируется самостоятельно, пример: <ipv4_addr> или <external-source> (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор в формате UUID, формируется по принципу <external-source--ff26c055-6336-5bc5-b98d-13d6226742dd> (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Value - может содержать какую либо дополнительную информацию
type IdentifierTypeSTIX struct {
	Type  string `json:"type" bson:"type"`
	ID    string `json:"id" bson:"id"`
	Value string `json:"value" bson:"value"`
}

//KillChainPhasesTypeSTIX тип содержащий цепочки фактов, приведшей к какому либо урону
type KillChainPhasesTypeSTIX []*KillChainPhasesTypeElementSTIX //`json:"kill_chain_phases" bson:"kill_chain_phases"`

//KillChainPhasesTypeElementSTIX тип содержащий набор элементов цепочки фактов, приведшей к какому либо урону
// KillChainName - имя цепочки (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PhaseName - наименование фазы из спецификации STIX, например, "reconnaissance", "pre-attack" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type KillChainPhasesTypeElementSTIX struct {
	KillChainName string `json:"kill_chain_name" bson:"kill_chain_name"`
	PhaseName     string `json:"phase_name" bson:"phase_name"`
}

/********** 			Data Markings STIX 			**********/

//CommonDataMarkingsTypeSTIX общие свойства меток данных
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type CommonDataMarkingsTypeSTIX struct {
	Type        string    `json:"type" bson:"type"`
	SpecVersion string    `json:"spec_version" bson:"spec_version"`
	ID          string    `json:"id" bson:"id"`
	Created     time.Time `json:"created" bson:"created"`
}

//GranularMarkingsTypeSTIX тип определяет как тип MarkingDefinitionObjectSTIX ссылается на свойства "marking_ref" и "lang"
// MarkingRef
// Lang
// Selectors
type GranularMarkingsTypeSTIX struct {
	MarkingRef `json:"marking_ref" bson:"marking_ref"`
	Lang `json:"lang" bson:"lang"`
	Selectors `json:"selectors" bson:"selectors"`
}

//MarkingDefinitionObjectSTIX объект определения маркировки содержит метки данных ссылающиеся на требования к обработке или совместному использованию данных
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

/********** 			Domain Objects STIX 			**********/

//CommonPropertiesObjectSTIX общие, для всех объектов STIX, свойства
// Type - наименование типа шаблона, для этого типа это поле ДОЛЖНО содержать "attack-pattern" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Modified - время модификации объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// CreatedByRef - содержит идентификатор источника создавшего данный объект
// Labels - определяет набор терминов, используемых для описания данного объекта
// Сonfidence - определяет уверенность создателя в правильности своих данных. Доверительное значение ДОЛЖНО быть числом
//  в диапазоне 0-100. Если 0 - значение не определено.
// Lang - содержит текстовый код языка, на котором написан контент объекта. Для английского это "en" для русского "ru"
// ExternalReferences - список внешних ссылок не относящихся к STIX информации
// ObjectMarkingRefs - определяет список свойств идентификаторов объектов определения маркировки, которые применяются к этому объекту
//  хотя оно и является списком типа IdentifierTypeSTIX, но тот в свою очередь ССЫЛАЕТСЯ на объект типа MarkingDefinitionObjectSTIX (marking-definition)
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
type CommonPropertiesObjectSTIX struct {
	Type               string                     `json:"type" bson:"type"`
	SpecVersion        string                     `json:"spec_version" bson:"spec_version"`
	ID                 string                     `json:"id" bson:"id"`
	Created            time.Time                  `json:"created" bson:"created"`
	Modified           time.Timer                 `json:"modified" bson:"modified"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	Labels             []string                   `json:"labels" bson:"labels"`
	Сonfidence         int                        `json:"confidence" bson:"confidence"`
	Lang               string                     `json:"lang" bson:"lang"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
}

//AttackPatternDomainObjectsSTIX объект "шаблон атаки" описывающий способы компрометации цели
// Name - имя используемое для идентификации шаблона атаки (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type AttackPatternDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX

	Name string `json:"name" bson:"name"`
}

//`json:"" bson:""`
