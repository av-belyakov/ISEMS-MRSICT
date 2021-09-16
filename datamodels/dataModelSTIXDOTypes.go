package datamodels

import (
	"time"
)

/********** 			Domain Objects STIX (ТИПЫ)			**********/

//CommonPropertiesDomainObjectSTIX свойства общие, для всех объектов STIX
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Modified - время модификации объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// CreatedByRef - содержит идентификатор источника создавшего данный объект
// Revoked - вернуть к текущему состоянию
// Labels - определяет набор терминов, используемых для описания данного объекта
// Сonfidence - определяет уверенность создателя в правильности своих данных. Доверительное значение ДОЛЖНО быть числом
//  в диапазоне 0-100. Если 0 - значение не определено.
// Lang - содержит текстовый код языка, на котором написан контент объекта. Для английского это "en" для русского "ru"
// ExternalReferences - список внешних ссылок не относящихся к STIX информации
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// Extensions - может содержать дополнительную информацию, относящуюся к объекту
type CommonPropertiesDomainObjectSTIX struct {
	SpecVersion        string                     `json:"spec_version" bson:"spec_version" required:"true"`
	Created            time.Time                  `json:"created" bson:"created" required:"true"`
	Modified           time.Time                  `json:"modified" bson:"modified" required:"true"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	Revoked            bool                       `json:"revoked" bson:"revoked"`
	Labels             []string                   `json:"labels" bson:"labels"`
	Сonfidence         int                        `json:"confidence" bson:"confidence"`
	Lang               string                     `json:"lang" bson:"lang"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []IdentifierTypeSTIX       `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   []GranularMarkingsTypeSTIX `json:"granular_markings" bson:"granular_markings"`
	Defanged           bool                       `json:"defanged" bson:"defanged"`
	Extensions         map[string]string          `json:"extensions" bson:"extensions"`
}

//AttackPatternDomainObjectsSTIX объект "Attack Pattern", по терминалогии STIX, описывающий способы компрометации цели
// Name - имя используемое для идентификации "Attack Pattern" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание атаки
// Aliases - альтернативные имена
// KillChainPhases - список цепочки фактов, в которых используется этот шаблон атак
type AttackPatternDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name            string                  `json:"name" bson:"name" required:"true"`
	Description     string                  `json:"description" bson:"description"`
	Aliases         []string                `json:"aliases" bson:"aliases"`
	KillChainPhases KillChainPhasesTypeSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
}

//CampaignDomainObjectsSTIX объект "Campaign", по терминалогии STIX, это набор действий определяющих злонамеренную деятельность или атаки
// Name - имя используемое для идентификации "Campaign" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации "Campaign"
// FirstSeen - время когда компания была впервые обнаружена, в формате "2016-05-12T08:17:27.000Z"
// LastSeen - время когда компания была зафиксирована в последний раз, в формате "2016-05-12T08:17:27.000Z"
// Objective - основная цель, желаемый результат или предполагаемый эффект
type CampaignDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name        string    `json:"name" bson:"name" required:"true"`
	Description string    `json:"description" bson:"description"`
	Aliases     []string  `json:"aliases" bson:"aliases"`
	FirstSeen   time.Time `json:"first_seen" bson:"first_seen"`
	LastSeen    time.Time `json:"last_seen" bson:"last_seen"`
	Objective   string    `json:"objective" bson:"objective"`
}

//CourseOfActionDomainObjectsSTIX объект "Course of Action", по терминалогии STIX, описывающий совокупность действий направленных
//  на предотвращение (защиту) либо реагирование на текущую атаку
// Name - имя используемое для идентификации "Course of Action" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Action - ЗАРЕЗЕРВИРОВАНО
type CourseOfActionDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name        string      `json:"name" bson:"name" required:"true"`
	Description string      `json:"description" bson:"description"`
	Action      interface{} `json:"action" bson:"action"`
}

//GroupingDomainObjectsSTIX объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
// Name - имя используемое для идентификации "Grouping" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Context - краткий дескриптор конкретного контекста, совместно используемого содержимым, на которое ссылается группа. Должно быть одно, из
//  заранее определенных (предложенных) значений (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - указывает на список объектов STIX, на которые ссылается эта группировка (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type GroupingDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name        string               `json:"name" bson:"name" required:"true"`
	Description string               `json:"description" bson:"description"`
	Context     OpenVocabTypeSTIX    `json:"context" bson:"context" required:"true"`
	ObjectRefs  []IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}

//IdentityDomainObjectsSTIX объект "Identity", по терминалогии STIX, содержит основную идентификационную информацию физичиских лиц, организаций и т.д.
// Name - имя используемое для идентификации "Identity" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Roles - список ролей для идентификации действий
// IdentityClass - одно, из заранее определенных (предложенных) значений физического лица или организации
// Sectors - заранее определенный (предложенный) перечень отраслей промышленности, к которой принадлежит физическое лицо или организация
// ContactInformation - любая контактная информация (email, phone number and etc.)
type IdentityDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name               string              `json:"name" bson:"name" required:"true"`
	Description        string              `json:"description" bson:"description"`
	Roles              []string            `json:"roles" bson:"roles"`
	IdentityClass      OpenVocabTypeSTIX   `json:"identity_class" bson:"identity_class"`
	Sectors            []OpenVocabTypeSTIX `json:"sectors" bson:"sectors"`
	ContactInformation string              `json:"contact_information" bson:"contact_information"`
}

//IndicatorDomainObjectsSTIX объект "Indicator", по терминалогии STIX, содержит шаблон который может быть использован для
//  обнаружения подозрительной или вредоносной киберактивности
// Name - имя используемое для идентификации "Indicator" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// IndicatorTypes - заранее определенный (предложенный) перечень категорий индикаторов
// Pattern - шаблон для обнаружения индикаторов (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternType - одно, из заранее определенных (предложенных) значений языкового шаблона, используемого в этом индикаторе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternVersion - версия языка шаблонов
// ValidFrom - время с которого этот индикатор считается валидным, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ValidUntil - время начиная с которого этот индикатор не может считаться валидным, в формате "2016-05-12T08:17:27.000Z"
// KillChainPhases - список цепочки фактов, к которым можно отнести соответствующте индикаторы
type IndicatorDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name            string                  `json:"name" bson:"name" required:"true"`
	Description     string                  `json:"description" bson:"description"`
	IndicatorTypes  []OpenVocabTypeSTIX     `json:"indicator_types" bson:"indicator_types"`
	Pattern         string                  `json:"pattern" bson:"pattern" required:"true"`
	PatternType     OpenVocabTypeSTIX       `json:"pattern_type" bson:"pattern_type" required:"true"`
	PatternVersion  string                  `json:"pattern_version" bson:"pattern_version"`
	ValidFrom       time.Time               `json:"valid_from" bson:"valid_from" required:"true"`
	ValidUntil      time.Time               `json:"valid_until" bson:"valid_until"`
	KillChainPhases KillChainPhasesTypeSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
}

//InfrastructureDomainObjectsSTIX объект "Infrastructure", по терминалогии STIX, содержит описание любых систем,
//  программных служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
// Name - имя используемое для идентификации "Infrastructure" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// InfrastructureTypes - заранее определенный (предложенный) перечень описываемых инфраструктур
// Aliases - альтернативные имена используемые для идентификации этой инфраструктуры
// KillChainPhases - список цепочки фактов, к которым может быть отнесена эта инфраструктура
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данная инфраструктура была впервые замечена за осуществлением вредоносной активности
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данная инфраструктура в последний раз была замечена за осуществлением вредоносной активности
type InfrastructureDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name                string                  `json:"name" bson:"name" required:"true"`
	Description         string                  `json:"description" bson:"description"`
	InfrastructureTypes []OpenVocabTypeSTIX     `json:"infrastructure_types" bson:"infrastructure_types"`
	Aliases             []string                `json:"aliases" bson:"aliases"`
	KillChainPhases     KillChainPhasesTypeSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	FirstSeen           time.Time               `json:"first_seen" bson:"first_seen"`
	LastSeen            time.Time               `json:"last_seen" bson:"last_seen"`
}

//IntrusionSetDomainObjectsSTIX объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный набор враждебного поведения и ресурсов
//  с общими свойствами, который, как считается, управляется одной организацией
// Name - имя используемое для идентификации "Intrusion Set" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации набора вторжения
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный набор вторжения впервые был зафиксирован
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный набор вторжения был зафиксирован в последний раз
// Goals - высокоуровневые цели этого набора вторжения
// ResourceLevel - заранее определенный (предложенный) перечень уровней, на которых обычно работает данный набор вторжений, который, в свою очередь,
//  определяет ресурсы, доступные этому набору вторжений для использования в атаке
// PrimaryMotivation - одно, из заранее определенных (предложенных) перечней причин, мотиваций или целей определяющий данный набор вторжения
// SecondaryMotivations - заранее определенный (предложенный) вторичный перечень причин, мотиваций или целей определяющий данный набор вторжений
type IntrusionSetDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name                 string              `json:"name" bson:"name" required:"true"`
	Description          string              `json:"description" bson:"description"`
	Aliases              []string            `json:"aliases" bson:"aliases"`
	FirstSeen            time.Time           `json:"first_seen" bson:"first_seen"`
	LastSeen             time.Time           `json:"last_seen" bson:"last_seen"`
	Goals                []string            `json:"goals" bson:"goals"`
	ResourceLevel        OpenVocabTypeSTIX   `json:"resource_level" bson:"resource_level"`
	PrimaryMotivation    OpenVocabTypeSTIX   `json:"primary_motivation" bson:"primary_motivation"`
	SecondaryMotivations []OpenVocabTypeSTIX `json:"secondary_motivations" bson:"secondary_motivations"`
}

//LocationDomainObjectsSTIX объект "Location", по терминалогии STIX, содержит описание географического местоположения
// Name - имя используемое для идентификации "Location"
// Description - более подробное описание
// Latitude - широта
// Longitude - долгота
// Precision - определяет точность координат, заданных свойствами широта и долгота (измеряется в метрах)
// Region - один, из заранее определенного (предложенного) перечня регионов
// Country - наименование страны
// AdministrativeArea - административный округ
// City - наименование города
// StreetAddress - физический адрес
// PostalCode - почтовый адрес
type LocationDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name               string            `json:"name" bson:"name"`
	Description        string            `json:"description" bson:"description"`
	Latitude           float32           `json:"latitude" bson:"latitude"`
	Longitude          float32           `json:"longitude" bson:"longitude"`
	Precision          float32           `json:"precision" bson:"precision"`
	Region             OpenVocabTypeSTIX `json:"region" bson:"region"`
	Country            string            `json:"country" bson:"country"`
	AdministrativeArea string            `json:"administrative_area" bson:"administrative_area"`
	City               string            `json:"city" bson:"city"`
	StreetAddress      string            `json:"street_address" bson:"street_address"`
	PostalCode         string            `json:"postal_code" bson:"postal_code"`
}

//MalwareDomainObjectsSTIX объект "Malware", по терминалогии STIX, содержит подробную информацию о функционировании вредоносной программы
// Name - имя используемое для идентификации "Malware"
// Description - более подробное описание
// MalwareTypes - заранее определенный (предложенный) перечень вредоносного ПО
// IsFamily - представляет ли объект семейство вредоносных программ (если true) или экземпляр вредоносного ПО (если false) (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Aliases - альтернативные имена используемые для идентификации этого ПО или семейства ПО
// KillChainPhases - список цепочки фактов, к которым может быть отнесено это вредоносное ПО
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда вредоносное ПО или семейство вредоносных программ были впервые замечины
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда вредоносное ПО или семейство вредоносных программ были замечины в последний раз
// OperatingSystemRefs - перечень идентификаторов ОС в которых может быть выполнено вредоносное ПО или семейство вредоносных программ
// ArchitectureExecutionEnvs - заранее определенный (предложенный) перечень архитектур в которых может быть выполнено вредоносное ПО или семейство
//  вредоносных программ
// ImplementationLanguages - заранее определенный (предложенный) перечень языков программирования, используемых для реализации вредоносного ПО или семейства
//  вредоносных программ
// Capabilities - заранее определенный (предложенный) перечень возможных идентификаторов используемых для обнаружения вредоносного ПО или семейства вредоносных
//  программ
// SampleRefs - определяет список идентификаторов файлов или объектов ассоциируемых с вредоносным ПО или семейством вредоносных программ
type MalwareDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name                      string                  `json:"name" bson:"name"`
	Description               string                  `json:"description" bson:"description"`
	MalwareTypes              []OpenVocabTypeSTIX     `json:"malware_types" bson:"malware_types"`
	IsFamily                  bool                    `json:"is_family" bson:"is_family" required:"true"`
	Aliases                   []string                `json:"aliases" bson:"aliases"`
	KillChainPhases           KillChainPhasesTypeSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	FirstSeen                 time.Time               `json:"first_seen" bson:"first_seen"`
	LastSeen                  time.Time               `json:"last_seen" bson:"last_seen"`
	OperatingSystemRefs       []IdentifierTypeSTIX    `json:"operating_system_refs" bson:"operating_system_refs"`
	ArchitectureExecutionEnvs []OpenVocabTypeSTIX     `json:"architecture_execution_envs" bson:"architecture_execution_envs"`
	ImplementationLanguages   []OpenVocabTypeSTIX     `json:"implementation_languages" bson:"implementation_languages"`
	Capabilities              []OpenVocabTypeSTIX     `json:"capabilities" bson:"capabilities"`
	SampleRefs                []IdentifierTypeSTIX    `json:"sample_refs" bson:"sample_refs"`
}

//MalwareAnalysisDomainObjectsSTIX объект "Malware Analysis", по терминалогии STIX, содержит анализ вредоносных программ захватывающих метаданные
//  и результаты конкретного статического или динамического анализа, выполненного на экземпляре вредоносного ПО или семействе вредоносных программ
// Product - название аналитического ПО использованного для обработки и анализа вредоносного ПО (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Version - версия аналитического ПО
// HostVMRef - идентификатор на описание виртуального окружения использованного для анализа вредоносного ПО
// OperatingSystemRef - идентификатор на описание ОС используемой для динамического анализа вредоносного ПО
// InstalledSoftwareRefs - список идентификаторов ссылающихся на описание любого нестандартного ПО установленного в ОС используемой для динамического
//  анализа вредоносного ПО
// ConfigurationVersion - именованная конфигурация дополнительных параметров конфигурации продукта, используемого для анализа
// Modules - конкретные модули анализа, которые были использованы и сконфигурированы в продукте во время выполнения анализа
// AnalysisEngineVersion - версия аналитического движка или продукта (включая AV-движки), который был использован для выполнения анализа
// AnalysisDefinitionVersion - версия определений анализа, используемых инструментом анализа (включая AV-инструменты)
// Submitted - время, в формате "2016-05-12T08:17:27.000Z", когда вредоносное ПО было впервые отправлено на сканирование или анализ
// AnalysisStarted - время, в формате "2016-05-12T08:17:27.000Z", начала анализа вредоносного ПО
// AnalysisEnded - время, в формате "2016-05-12T08:17:27.000Z", когда был завершен анализ вредоносного ПО
// ResultName - результат классификации или имя, присвоенное экземпляру вредоносного ПО инструментом анализа (сканером)
// Result - один, из заранее определенного (предложенного) перечня результатов классификации, определяется аналитическим инструментом или сканером
// AnalysisScoRefs - список идентификаторов на другие наблюдаемые Domain Objects STIX которые были захвачены в процессе наблюдения
// SampleRef - содержит ссылку на файл, сетевой трафик или объект на основе которого был выполнен анализ вредоносного ПО
// AvResult - результат аналитической обработки (ЭТО ПОЛЕ ЕСТЬ ТОЛЬКО В ПРИМЕРЕ, в описании типа данного поля нет)
type MalwareAnalysisDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Product                   string               `json:"product" bson:"product" required:"true"`
	Version                   string               `json:"version" bson:"version"`
	HostVMRef                 IdentifierTypeSTIX   `json:"host_vm_ref" bson:"host_vm_ref"`
	OperatingSystemRef        IdentifierTypeSTIX   `json:"operating_system_ref" bson:"operating_system_ref"`
	InstalledSoftwareRefs     []IdentifierTypeSTIX `json:"installed_software_refs" bson:"installed_software_refs"`
	ConfigurationVersion      string               `json:"configuration_version" bson:"configuration_version"`
	Modules                   []string             `json:"modules" bson:"modules"`
	AnalysisEngineVersion     string               `json:"analysis_engine_version" bson:"analysis_engine_version"`
	AnalysisDefinitionVersion string               `json:"analysis_definition_version" bson:"analysis_definition_version"`
	Submitted                 time.Time            `json:"submitted" bson:"submitted"`
	AnalysisStarted           time.Time            `json:"analysis_started" bson:"analysis_started"`
	AnalysisEnded             time.Time            `json:"analysis_ended" bson:"analysis_ended"`
	ResultName                string               `json:"result_name" bson:"result_name"`
	Result                    OpenVocabTypeSTIX    `json:"result" bson:"result"`
	AnalysisScoRefs           []IdentifierTypeSTIX `json:"analysis_sco_refs" bson:"analysis_sco_refs"`
	SampleRef                 IdentifierTypeSTIX   `json:"sample_ref" bson:"sample_ref"`
	AvResult                  OpenVocabTypeSTIX    `json:"av_result" bson:"av_result"`
}

//NoteDomainObjectsSTIX объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую текущий контекст анализа либо содержащей
//  результаты дополнительного анализа которые не может быть описан в терминах объектов STIX
// Abstract - краткое изложение содержания записки
// Content - основное содержание (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Authors - список авторов
// ObjectRefs - список идентификаторов на других DO STIX объектов к которым применяется замечание (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type NoteDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Abstract   string               `json:"abstract" bson:"abstract"`
	Content    string               `json:"content" bson:"content" required:"true"`
	Authors    []string             `json:"authors" bson:"authors"`
	ObjectRefs []IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}

//ObservedDataDomainObjectsSTIX объект "Observed Data", по терминалогии STIX, содержит информацию о сущностях связанных с кибер безопасностью, таких как файлы,
//  системы или сети. Наблюдаемые данные это не результат анализа или заключение искусственного интеллекта, это просто сырая информация без какого-либо контекста.
// FirstObserved - время, в формате "2016-05-12T08:17:27.000Z", начала временного окна, в течение которого были замечены данные (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// LastObserved - время, в формате "2016-05-12T08:17:27.000Z", окончание временного окна, в течение которого были замечены данные (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// NumberObserved - количество раз, когда фиксировался каждый наблюдаемый кибер объект SCO, представленный в свойстве ObjectRef (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов на другие наблюдаемые кибер объекты SCO
type ObservedDataDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	FirstObserved  time.Time            `json:"first_observed" bson:"first_observed" required:"true"`
	LastObserved   time.Time            `json:"last_observed" bson:"last_observed" required:"true"`
	NumberObserved int                  `json:"number_observed" bson:"number_observed" required:"true"`
	ObjectRefs     []IdentifierTypeSTIX `json:"object_refs" bson:"object_refs"`
}

//OpinionDomainObjectsSTIX объект "Opinion", по терминалогии STIX, содержит оценку информации в приведенной в каком либо другом объекте STIX, которую произвел
//  другой участник анализа.
// Explanation - объясняет почему обработчик оставил это мнение
// Authors - список авторов этого мнения
// Opinion - мнение обо всех STIX объектах перечисленных в ObjectRefs (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов на другие STIX объекты (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type OpinionDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Explanation string               `json:"explanation" bson:"explanation"`
	Authors     []string             `json:"authors" bson:"authors"`
	Opinion     EnumTypeSTIX         `json:"opinion" bson:"opinion" required:"true"`
	ObjectRefs  []IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}

//ReportDomainObjectsSTIX объект "Report", по терминалогии STIX, содержит совокупность данных об угрозах, сосредоточенных на одной или нескольких темах,
//  таких как описание исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним детали. Применяется для группировки информации
//  связанной с кибер угрозой. Может быть использован для дальнейшей публикации данной информации как истории расследования.
// Name - имя используемое для идентификации "Report" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ReportTypes - заранее определенный (предложенный) перечень возможных типов контента содержащихся в этом отчете
// Published - время, в формате "2016-05-12T08:17:27.000Z", когда данный отчет был официально опубликован (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов STIX объектов, которые ссылаются на этот отчет (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// OutsideSpecification - свойства не входящие в основную спецификацию STIX 2.0
type ReportDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name                 string                     `json:"name" bson:"name" required:"true"`
	Description          string                     `json:"description" bson:"description"`
	ReportTypes          []OpenVocabTypeSTIX        `json:"report_types" bson:"report_types"`
	Published            time.Time                  `json:"published" bson:"published" required:"true"`
	ObjectRefs           []IdentifierTypeSTIX       `json:"object_refs" bson:"object_refs" required:"true"`
	OutsideSpecification ReportOutsideSpecification `json:"outside_specification" bson:"outside_specification"`
}

//ThreatActorDomainObjectsSTIX объект "Threat Actor", по терминалогии STIX, содержит информацию о физических лицах или их группах и организациях
//  которые могут действовать со злым умыслом.
// Name - имя используемое для идентификации "Threat Actor" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ThreatActorTypes - заранее определенный (предложенный) перечень типов субъектов угрозы
// Aliases - альтернативные имена используемые для этого субъекта угроз
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный субъект угроз был впервые зафиксирован
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный субъект угроз был зафиксирован в последний раз
// Roles - заранее определенный (предложенный) перечень возможных ролей субъекта угроз
// Goals - высокоуровневые цели субъекта угроз.
// Sophistication - один, из заранее определенного (предложенного) перечня навыков, специальных знания, специальной подготовки или опыта,
//  которыми должен обладать субъект угрозы, чтобы осуществить атаку
// ResourceLevel - один, из заранее определенного (предложенного) перечня организационных уровней, на котором обычно работает этот субъект угрозы,
//  который, в свою очередь, определяет ресурсы, доступные этому субъекту угрозы для использования в атаке.
// PrimaryMotivation - одна, из заранее определенного (предложенного) перечня причин, мотиваций или целей стоящих за этим субъектом угрозы
// SecondaryMotivations - заранее определенный (предложенный) перечень возможных вторичных причин, мотиваций или целей стоящих за этим субъектом угрозы
// PersonalMotivations - заранее определенный (предложенный) перечень возможных персональных причин, мотиваций или целей стоящих за этим субъектом угрозы
type ThreatActorDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name                 string              `json:"name" bson:"name" required:"true"`
	Description          string              `json:"description" bson:"description"`
	ThreatActorTypes     []OpenVocabTypeSTIX `json:"threat_actor_types" bson:"threat_actor_types"`
	Aliases              []string            `json:"aliases" bson:"aliases"`
	FirstSeen            time.Time           `json:"first_seen" bson:"first_seen"`
	LastSeen             time.Time           `json:"last_seen" bson:"last_seen"`
	Roles                []OpenVocabTypeSTIX `json:"roles" bson:"roles"`
	Goals                []string            `json:"goals" bson:"goals"`
	Sophistication       OpenVocabTypeSTIX   `json:"sophistication" bson:"sophistication"`
	ResourceLevel        OpenVocabTypeSTIX   `json:"resource_level" bson:"resource_level"`
	PrimaryMotivation    OpenVocabTypeSTIX   `json:"primary_motivation" bson:"primary_motivation"`
	SecondaryMotivations []OpenVocabTypeSTIX `json:"secondary_motivations" bson:"secondary_motivations"`
	PersonalMotivations  []OpenVocabTypeSTIX `json:"personal_motivations" bson:"personal_motivations"`
}

//ToolDomainObjectsSTIX объект "Tool", по терминалогии STIX, содержит информацию о легитимном ПО которое может быть использованно для реализации
//  компьютерных угроз
// Name - имя используемое для идентификации "Tool" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ToolTypes - заранее определенный (предложенный) перечень типов инструментов
// Aliases - альтернативные имена используемые для идентификации инструментов
// KillChainPhases - список цепочки фактов, к которым может быть отнесен этот инструмент
// ToolVersion - версия инструмента
type ToolDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name            string                  `json:"name" bson:"name" required:"true"`
	Description     string                  `json:"description" bson:"description"`
	ToolTypes       []OpenVocabTypeSTIX     `json:"tool_types" bson:"tool_types"`
	Aliases         []string                `json:"aliases" bson:"aliases"`
	KillChainPhases KillChainPhasesTypeSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	ToolVersion     string                  `json:"tool_version" bson:"tool_version"`
}

//VulnerabilityDomainObjectsSTIX объект "Vulnerability", по терминологии STIX, содержит описание уязвимостей полученных в результате неверной формализации
//  требований, ошибочном проектировании или некорректной реализации программного кода или логики в ПО, а также в компонентах оборудования
// Name - имя используемое для идентификации "Vulnerability" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
type VulnerabilityDomainObjectsSTIX struct {
	CommonPropertiesObjectSTIX
	CommonPropertiesDomainObjectSTIX
	Name        string `json:"name" bson:"name" required:"true"`
	Description string `json:"description" bson:"description"`
}
