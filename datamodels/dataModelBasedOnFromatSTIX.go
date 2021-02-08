package datamodels

import (
	"time"
)

/**********			 Некоторые примитивные типы STIX			 **********/

//ExternalReferencesTypeSTIX тип "external-reference", по терминалогии STIX является списком с информацией о внешних ссылках не относящихся к STIX информации
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

//HashesTypeSTIX тип "hashes", по терминалогии STIX, содержащий хеш хначения, где <тип_хеша>:<хеш>
type HashesTypeSTIX map[string]string

//IdentifierTypeSTIX тип "identifier", по терминалогии STIX, содержащий уникальный идентификатор UUID, преимущественно версии 4 при этом ID должен
//начинаться с наименования организации или программного обеспечения сгенерировавшего его. Например, <example-source--ff26c055-6336-5bc5-b98d-13d6226742dd> (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type IdentifierTypeSTIX string

//KillChainPhasesTypeSTIX тип "kill-chain-phases", по терминалогии STIX, содержащий цепочки фактов, приведших к какому либо урону
type KillChainPhasesTypeSTIX []*KillChainPhasesTypeElementSTIX //`json:"kill_chain_phases" bson:"kill_chain_phases"`

//KillChainPhasesTypeElementSTIX тип содержащий набор элементов цепочки фактов, приведших к какому либо урону
// KillChainName - имя цепочки (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PhaseName - наименование фазы из спецификации STIX, например, "reconnaissance", "pre-attack" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type KillChainPhasesTypeElementSTIX struct {
	KillChainName string `json:"kill_chain_name" bson:"kill_chain_name"`
	PhaseName     string `json:"phase_name" bson:"phase_name"`
}

//OpenVocabTypeSTIX тип "open-vocab", по терминалогии STIX, содержащий заранее определенное (предложенное) значение
type OpenVocabTypeSTIX string

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

//GranularMarkingsTypeSTIX тип "granular_markings", по терминалогии STIX, представляет собой набор маркеров ссылающихся на свойства "marking_ref" и "lang"
// Lang - идентифицирует язык соответствующим маркером
// MarkingRef - определяет идентификатор объекта "marking-definition"
// Selectors - определяет список селекторов для содержимого объекта STIX, к которому применяется это свойство
type GranularMarkingsTypeSTIX struct {
	Lang       string             `json:"lang" bson:"lang"`
	MarkingRef IdentifierTypeSTIX `json:"marking_ref" bson:"marking_ref"`
	Selectors  []string           `json:"selectors" bson:"selectors"`
}

//MarkingDefinitionObjectSTIX объект "marking-definition", по терминалогии STIX, содержит метки данных ссылающиеся на требования к обработке или совместному использованию данных
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

//CommonPropertiesObjectSTIX свойства общие, для всех объектов STIX, свойства
// Type - наименование типа шаблона (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//  Type должен содержать одно из следующих значений:
//  - "attack-pattern"
//  - "campaign"
//  - "course-of-action"
//  - "grouping"
//  - "identity"
//  - "indicator"
//  - "infrastructure"
//  - "intrusion-set"
//  -
//  -
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
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// Extensions - может содержать дополнительную информацию, относящуюся к объекту
type CommonPropertiesObjectSTIX struct {
	Type               string                     `json:"type" bson:"type"`
	SpecVersion        string                     `json:"spec_version" bson:"spec_version"`
	ID                 string                     `json:"id" bson:"id"`
	Created            time.Time                  `json:"created" bson:"created"`
	Modified           time.Time                  `json:"modified" bson:"modified"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	Labels             []string                   `json:"labels" bson:"labels"`
	Сonfidence         int                        `json:"confidence" bson:"confidence"`
	Lang               string                     `json:"lang" bson:"lang"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
	Defanged           bool                       `json:"defanged" bson:"defanged"`
	Extensions         map[string]string          `json:"extensions" bson:"extensions"`
}

//AttackPatternDomainObjectsSTIX объект "Attack Pattern", по терминалогии STIX, описывающий способы компрометации цели
// Name - имя используемое для идентификации "Attack Pattern" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание атаки
// Aliases - альтернативные имена
// KillChainPhases - список цепочки фактов, приведших к урону
type AttackPatternDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name            string                            `json:"name" bson:"name"`
	Description     string                            `json:"description" bson:"description"`
	Aliases         []string                          `json:"aliases" bson:"aliases"`
	KillChainPhases []*KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
}

//CampaignDomainObjectsSTIX объект "Campaign", по терминалогии STIX, набор действий определяющих злонамеренную деятельность или атаки
// Name - имя используемое для идентификации "Campaign" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации "Campaign"
// FirstSeen - время когда компания была впервые обнаружена
// LastSeen - время когда компания была зафиксирована в последний раз
// Objective - основная цель, желаемый результат или предполагаемый эффект
type CampaignDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Aliases     []string  `json:"aliases" bson:"aliases"`
	FirstSeen   time.Time `json:"first_seen" bson:"first_seen"`
	LastSeen    time.Time `json:"last_seen" bson:"last_seen"`
	Objective   string    `json:"objective" bson:"objective"`
}

//CourseOfActionDomainObjectsSTIX объект "Course of Action", по терминалогии STIX, описывающий совокупность действий направленных
// на предотвращение (защиту) либо реагирование на текущую атаку
// Name - имя используемое для идентификации "Course of Action" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Action - ЗАРЕЗЕРВИРОВАНО
type CourseOfActionDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Action      interface{} `json:"action" bson:"action"`
}

//GroupingDomainObjectsSTIX объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
// Name - имя используемое для идентификации "Course of Action" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Context - заранее определенное (предложенное) значение (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - указывает объекты STIX, на которые ссылается эта группировка (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type GroupingDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name        string                `json:"name" bson:"name"`
	Description string                `json:"description" bson:"description"`
	Context     OpenVocabTypeSTIX     `json:"context" bson:"context"`
	ObjectRefs  []*IdentifierTypeSTIX `json:"object_refs" bson:"object_refs"`
}

//IdentityDomainObjectsSTIX объект "Identity", по терминалогии STIX, содержит основную идентификационную информацию физичиских лиц, организаций и т.д.
// Name - имя используемое для идентификации "Identity" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Roles - список ролей для идентификации действий
// IdentityClass - заранее определенное (предложенное) значение физического лица или организации
// Sectors - заранее определенный (предложенный) перечень отраслей промышленности, к которой принадлежит физическое лицо или организация
// ContactInformation -
type IdentityDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name               string               `json:"name" bson:"name"`
	Description        string               `json:"description" bson:"description"`
	Roles              []string             `json:"roles" bson:"roles"`
	IdentityClass      OpenVocabTypeSTIX    `json:"identity_class" bson:"identity_class"`
	Sectors            []*OpenVocabTypeSTIX `json:"sectors" bson:"sectors"`
	ContactInformation string               `json:"contact_information" bson:"contact_information"`
}

//IndicatorDomainObjectsSTIX объект "Indicator", по терминалогии STIX, содержит шаблон который может быть использован для
// обнаружения подозрительной или вредоносной киберактивности
// Name - имя используемое для идентификации "Indicator" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// IndicatorTypes - заранее определенный (предложенный) перечень категорий индикаторов
// Pattern - шаблон для обнаружения индикаторов (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternType - языковой шаблон используемый в этом индикаторе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternVersion - версия языка шаблонов
// ValidFrom - время с которого этот индикатор считается валидным (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ValidUntil - время начиная с которого этот индикатор не может считаться валидным
// KillChainPhases - список цепочки фактов, которые соответствуют индикатору
type IndicatorDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name            string                            `json:"name" bson:"name"`
	Description     string                            `json:"description" bson:"description"`
	IndicatorTypes  []*OpenVocabTypeSTIX              `json:"indicator_types" bson:"indicator_types"`
	Pattern         string                            `json:"pattern" bson:"pattern"`
	PatternType     OpenVocabTypeSTIX                 `json:"pattern_type" bson:"pattern_type"`
	PatternVersion  string                            `json:"pattern_version" bson:"pattern_version"`
	ValidFrom       time.Time                         `json:"valid_from" bson:"valid_from"`
	ValidUntil      time.Time                         `json:"valid_until" bson:"valid_until"`
	KillChainPhases []*KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
}

//InfrastructureDomainObjectsSTIX объект "Infrastructure", по терминалогии STIX, содержит описание любых систем,
//  программных служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
// Name - имя используемое для идентификации "Infrastructure" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// InfrastructureTypes - заранее определенный (предложенный) перечень описываемых инфраструктур
// Aliases - альтернативные имена используемые для идентификации этой инфраструктуры
// KillChainPhases - список цепочки фактов, для которых используется эта инфраструктура
// FirstSeen - время, когда данная инфраструктура была впервые замечена за осуществлением вредоносной активности
// LastSeen - время, когда данная инфраструктура в последний раз была замечена за осуществлением вредоносной активности
type InfrastructureDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name                string                            `json:"name" bson:"name"`
	Description         string                            `json:"description" bson:"description"`
	InfrastructureTypes []*OpenVocabTypeSTIX              `json:"infrastructure_types" bson:"infrastructure_types"`
	Aliases             []string                          `json:"aliases" bson:"aliases"`
	KillChainPhases     []*KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	FirstSeen           time.Time                         `json:"first_seen" bson:"first_seen"`
	LastSeen            time.Time                         `json:"last_seen" bson:"last_seen"`
}

//IntrusionSetDomainObjectsSTIX объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный набор враждебного поведения и ресурсов
//  с общими свойствами, который, как считается, управляется одной организацией
// Name - имя используемое для идентификации "Intrusion Set" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации набора вторжения
// FirstSeen - время, когда данный набор вторжения впервые был зафиксирован
// LastSeen - время, когда данный набор вторжения был зафиксирован в последний раз
// Goals - высокоуровневые цели этого набора вторжения
// ResourceLevel - заранее определенный (предложенный) перечень уровней, на которых обычно работает данный набор вторжений, который, в свою очередь,
//  определяет ресурсы, доступные этому набору вторжений для использования в атаке
// PrimaryMotivation - заранее определенный (предложенный) перечень причин, мотиваций или целей определяющий данный набор вторжений
// SecondaryMotivations - заранее определенный (предложенный) вторичный перечень причин, мотиваций или целей определяющий данный набор вторжений
type IntrusionSetDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	Name                 string               `json:"name" bson:"name"`
	Description          string               `json:"description" bson:"description"`
	Aliases              []string             `json:"aliases" bson:"aliases"`
	FirstSeen            time.Time            `json:"first_seen" bson:"first_seen"`
	LastSeen             time.Time            `json:"last_seen" bson:"last_seen"`
	Goals                []string             `json:"goals" bson:"goals"`
	ResourceLevel        OpenVocabTypeSTIX    `json:"resource_level" bson:"resource_level"`
	PrimaryMotivation    OpenVocabTypeSTIX    `json:"primary_motivation" bson:"primary_motivation"`
	SecondaryMotivations []*OpenVocabTypeSTIX `json:"secondary_motivations" bson:"secondary_motivations"`
}

//`json:"" bson:""`
