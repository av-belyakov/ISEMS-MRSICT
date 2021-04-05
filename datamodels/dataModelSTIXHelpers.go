package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	govalidator "github.com/asaskevich/govalidator"

	"ISEMS-MRSICT/commonlibs"
)

/**********			 Некоторые примитивные типы STIX			 **********/

//EnumTypeSTIX тип "enum", по терминологии STIX, является жестко заданным списком терминов, который представлен в виде строки
type EnumTypeSTIX string

//ExternalReferencesTypeSTIX тип "external-reference", по терминалогии STIX, является списком с информацией о внешних ссылках не относящихся к STIX информации
type ExternalReferencesTypeSTIX []*ExternalReferenceTypeElementSTIX

//CheckExternalReferencesTypeSTIX выполняет проверку значений типа ExternalReferencesTypeSTIX
func (ertstix *ExternalReferencesTypeSTIX) CheckExternalReferencesTypeSTIX() bool {
	if len(*ertstix) == 0 {
		return true
	}

	for _, v := range *ertstix {
		if !v.CheckExternalReferenceTypeElementSTIX() {
			return false
		}
	}

	return true
}

//SanitizeStructExternalReferencesTypeSTIX для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ertstix ExternalReferencesTypeSTIX) SanitizeStructExternalReferencesTypeSTIX() ExternalReferencesTypeSTIX {
	size := len(ertstix)
	if size == 0 {
		return ertstix
	}

	ert := make([]*ExternalReferenceTypeElementSTIX, 0, size)
	for _, v := range ertstix {
		tmp := v.SanitizeStructExternalReferenceTypeElementSTIX()
		ert = append(ert, &tmp)
	}

	return ert
}

//ExternalReferenceTypeElementSTIX тип содержащий подробную информацию о внешних ссылках, таких как URL, ID и т.д.
// SourceName - имя источника (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - описание
// URL - URL ссылка на внешних источниках
// Hashes - содержит словарь хэшей для содержимого URL-адреса. Это ДОЛЖНО быть предусмотрено при наличии свойства url
// ExternalID - идентификатор на внешних источниках
type ExternalReferenceTypeElementSTIX struct {
	SourceName  string         `json:"source_name" bson:"source_name" required:"true"`
	Description string         `json:"description" bson:"description"`
	URL         string         `json:"url" bson:"url"`
	Hashes      HashesTypeSTIX `json:"hashes" bson:"hashes"`
	ExternalID  string         `json:"external_id" bson:"external_id"`
}

//CheckExternalReferenceTypeElementSTIX выполняет проверку значений типа ExternalReferenceTypeElementSTIX
func (ertestix *ExternalReferenceTypeElementSTIX) CheckExternalReferenceTypeElementSTIX() bool {
	if ertestix.URL != "" && !govalidator.IsURL(ertestix.URL) {
		return false
	}

	return true
}

//SanitizeStructExternalReferenceTypeElementSTIX выполняет проверку значений типа ExternalReferenceTypeElementSTIX
func (ertestix ExternalReferenceTypeElementSTIX) SanitizeStructExternalReferenceTypeElementSTIX() ExternalReferenceTypeElementSTIX {
	return ExternalReferenceTypeElementSTIX{
		SourceName:  commonlibs.StringSanitize(ertestix.SourceName),
		Description: commonlibs.StringSanitize(ertestix.Description),
		URL:         ertestix.URL,
		Hashes:      ertestix.Hashes,
		ExternalID:  commonlibs.StringSanitize(ertestix.ExternalID),
	}
}

//HashesTypeSTIX тип "hashes", по терминологии STIX, содержащий хеш значения, где <тип_хеша>:<хеш>
type HashesTypeSTIX map[string]string

//CheckHashesTypeSTIX выполняет проверку значений типа HashesTypeSTIX
func (htstix *HashesTypeSTIX) CheckHashesTypeSTIX() bool {
	if len(*htstix) == 0 {
		return true
	}

	pattern := regexp.MustCompile(`^[0-9a-zA-Z-_=]+$`)
	for k, v := range *htstix {
		if !pattern.MatchString(k) {
			return false
		}

		if !pattern.MatchString(v) {
			return false
		}
	}

	return true
}

//IdentifierTypeSTIX тип "identifier", по терминалогии STIX, содержащий уникальный идентификатор UUID, преимущественно версии 4 при этом ID должен
//начинаться с наименования организации или программного обеспечения сгенерировавшего его. Например, <example-source--ff26c055-6336-5bc5-b98d-13d6226742dd> (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type IdentifierTypeSTIX string

//CheckIdentifierTypeSTIX выполняет проверку значения типа IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) CheckIdentifierTypeSTIX() bool {
	if len(*itstix) == 0 {
		return true
	}

	return regexp.MustCompile(`^[0-9a-zA-Z-_]+(--)[0-9a-f-]+$`).MatchString(fmt.Sprint(*itstix))
}

//AddValue добавляет значение в тип IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) AddValue(str string) {
	var i IdentifierTypeSTIX = IdentifierTypeSTIX(str)
	itstix = &i
}

//KillChainPhasesTypeSTIX тип "kill-chain-phases", по терминалогии STIX, содержащий цепочки фактов, приведших к какому либо урону
type KillChainPhasesTypeSTIX []*KillChainPhasesTypeElementSTIX

//SanitizeStructKillChainPhasesTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (kcptstix KillChainPhasesTypeSTIX) SanitizeStructKillChainPhasesTypeSTIX() KillChainPhasesTypeSTIX {
	for k := range kcptstix {
		kcptstix[k].SanitizeStructKillChainPhasesTypeElementSTIX()
	}

	return kcptstix
}

//KillChainPhasesTypeElementSTIX тип содержащий набор элементов цепочки фактов, приведших к какому либо урону
// KillChainName - имя цепочки (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PhaseName - наименование фазы из спецификации STIX, например, "reconnaissance", "pre-attack" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type KillChainPhasesTypeElementSTIX struct {
	KillChainName string `json:"kill_chain_name" bson:"kill_chain_name" required:"true"`
	PhaseName     string `json:"phase_name" bson:"phase_name" required:"true"`
}

//SanitizeStructKillChainPhasesTypeElementSTIX выполлняет проверку значения типа KillChainPhasesTypeElementSTIX
func (kcptestix KillChainPhasesTypeElementSTIX) SanitizeStructKillChainPhasesTypeElementSTIX() KillChainPhasesTypeElementSTIX {
	if kcptestix.KillChainName == "" {
		return kcptestix
	}

	if kcptestix.PhaseName == "" {
		return kcptestix
	}

	kcptestix.KillChainName = commonlibs.StringSanitize(kcptestix.KillChainName)
	kcptestix.PhaseName = commonlibs.StringSanitize(kcptestix.PhaseName)

	return kcptestix
}

//OpenVocabTypeSTIX тип "open-vocab", по терминалогии STIX, содержащий заранее определенное (предложенное) значение
type OpenVocabTypeSTIX string

//SanitizeStructOpenVocabTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (ovtstix OpenVocabTypeSTIX) SanitizeStructOpenVocabTypeSTIX() OpenVocabTypeSTIX {
	return OpenVocabTypeSTIX(commonlibs.StringSanitize(fmt.Sprint(ovtstix)))
}

//DictionaryTypeSTIX тип "dictionary", по терминалогии STIX, содержащий значения любых типов
type DictionaryTypeSTIX struct {
	dictionary interface{}
}

//UnmarshalJSON дополнительный метод декодирования
func (dtstix *DictionaryTypeSTIX) UnmarshalJSON(data []byte) error {
	var (
		err       error
		str       string
		num       int
		listStr   []string
		listInt   []int
		mapStrStr map[string]string
		mapStrInt map[string]int
		mapIntInt map[int]int
		mapIntStr map[int]string
	)

	if err = json.Unmarshal(data, &str); err == nil {
		dtstix.dictionary = str

		return nil
	}
	if err = json.Unmarshal(data, &num); err == nil {
		dtstix.dictionary = num
		return nil
	}
	if err = json.Unmarshal(data, &listStr); err == nil {
		dtstix.dictionary = listStr
		return nil
	}
	if err = json.Unmarshal(data, &listInt); err == nil {
		dtstix.dictionary = listInt
		return nil
	}
	if err = json.Unmarshal(data, &mapStrStr); err == nil {
		dtstix.dictionary = mapStrStr
		return nil
	}
	if err = json.Unmarshal(data, &mapStrInt); err == nil {
		dtstix.dictionary = mapStrInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntInt); err == nil {
		dtstix.dictionary = mapIntInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntStr); err == nil {
		dtstix.dictionary = mapIntStr
		return nil
	}
	//		dtstix.dictionary = fmt.Sprintln(mapIntStr)
	dtstix.dictionary = data

	return fmt.Errorf("JSON message parsing error, undefined value found in the DictionaryTypeSTIX type")
}

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
	Type               string                     `json:"type" bson:"type" required:"true"`
	ID                 string                     `json:"id" bson:"id" required:"true"`
	SpecVersion        string                     `json:"spec_version" bson:"spec_version" required:"true"`
	Created            time.Time                  `json:"created" bson:"created" required:"true"`
	Modified           time.Time                  `json:"modified" bson:"modified" required:"true"`
	ObjectRef          IdentifierTypeSTIX         `json:"object_ref" bson:"object_ref" required:"true"`
	ObjectModified     time.Time                  `json:"object_modified" bson:"object_modified"`
	Contents           map[string]string          `json:"contents" bson:"contents" required:"true"`
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
	SpecVersion string    `json:"spec_version" bson:"spec_version" required:"true"`
	ID          string    `json:"id" bson:"id" required:"true"`
	Created     time.Time `json:"created" bson:"created" required:"true"`
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

func (gmtstix *GranularMarkingsTypeSTIX) CheckGranularMarkingsTypeSTIX() bool {
	//для поля Lang
	if gmtstix.Lang != "" && !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(gmtstix.Lang) {
		return false
	}

	if !gmtstix.MarkingRef.CheckIdentifierTypeSTIX() {
		return false
	}

	selectorTmp := make([]string, 0, len(gmtstix.Selectors))
	for _, v := range gmtstix.Selectors {
		selectorTmp = append(selectorTmp, commonlibs.StringSanitize(v))
	}
	gmtstix.Selectors = selectorTmp

	return true
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
	Type               string                     `json:"type" bson:"type" required:"true"`
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
	Type    string        `json:"type" bson:"type" required:"true"`
	ID      string        `json:"id" bson:"id" required:"true"`
	Objects []interface{} `json:"objects" bson:"objects"`
}

/********** 			Свойства общие, для всех объектов STIX 			**********/

//CommonPropertiesObjectSTIX свойства общие, для всех объектов STIX
// Type - наименование типа шаблона (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//  Type должен содержать одно из следующих значений:
//  1. Для Domain Objects STIX
// - "attack-pattern"
// - "campaign"
// - "course-of-action"
// - "grouping"
// - "identity"
// - "indicator"
// - "infrastructure"
// - "intrusion-set"
// - "location"
// - "malware"
// - "malware-analysis"
// - "note"
// - "observed-data"
// - "opinion"
// - "report"
// - "threat-actor"
// - "tool"
// - "vulnerability"
//  2. Для Relationship Objects STIX
// - "relationship"
// - "sighting"
//  3. Для Cyber-observable Objects STIX
// - "artifact"
// - "autonomous-system"
// - "directory"
// - "domain-name"
// - "email-addr"
// - "email-message"
// - "file"
// - "ipv4-addr"
// - "ipv6-addr"
// - "mac-addr"
// - "mutex"
// - "network-traffic"
// - "process"
// - "software"
// - "url"
// - "user-account"
// - "windows-registry-key"
// - "x509-certificate"
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type CommonPropertiesObjectSTIX struct {
	Type string `json:"type" bson:"type" required:"true"`
	ID   string `json:"id" bson:"id" required:"true"`
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (cp CommonPropertiesObjectSTIX) ToStringBeautiful() string {
	return fmt.Sprintf("type: '%s'\nid: '%s'\n", cp.Type, cp.ID)
}
