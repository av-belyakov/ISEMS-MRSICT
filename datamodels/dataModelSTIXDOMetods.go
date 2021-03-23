package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
)

/********** 			Domain Objects STIX (МЕТОДЫ)			**********/

/*
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
	+ SpecVersion        string                     `json:"spec_version" bson:"spec_version" required:"true"`
	no Created            time.Time                  `json:"created" bson:"created" required:"true"`
	no Modified           time.Time                  `json:"modified" bson:"modified" required:"true"`
	+ CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	no Revoked            bool                       `json:"revoked" bson:"revoked"`
	- Labels             []string                   `json:"labels" bson:"labels"`
	no Сonfidence         int                        `json:"confidence" bson:"confidence"`
	+ Lang               string                     `json:"lang" bson:"lang"`
	hisself ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	hisself ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	hisself GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
	no Defanged           bool                       `json:"defanged" bson:"defanged"`
	- Extensions         map[string]string          `json:"extensions" bson:"extensions"`
}
*/

func (cpdostix *CommonPropertiesDomainObjectSTIX) checkingTypeCommonFields() bool {
	fmt.Println("func 'checkingTypeCommonFields', START...")

	//для поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(cpdostix.SpecVersion)) {
		return false
	}

	//для поля CreatedByRef
	if len(cpdostix.CreatedByRef.String()) > 0 {
		if !(regexp.MustCompile(`^[0-9a-zA-Z]+(--)[0-9a-f|-]+$`).MatchString(cpdostix.CreatedByRef.String())) {
			return false
		}
	}

	//дезинфекция содержимого списка поля Labels
	if len(cpdostix.Labels) > 0 {
		nl := make([]string, len(cpdostix.Labels))

		/*for _, l := range cpdostix.Labels {
			nl = append(nl, validation.ReplacingCharactersString(l))
		}*/

		cpdostix.Labels = nl
	}

	//для поля Lang
	if !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(cpdostix.Lang) {
		return false
	}

	/*
		if cpdostix.CreatedByRef {
			return false
		}*

			//rtype := reflect.TypeOf(testTypeOne.Extensions)
			/*
				валидация строк:
				- удаление (замена) нежелательных символов или вырожений
				- проверка строк на наличие ключевых строк в начале строки (для некоторых строк).
				 Например для поля ID строка должна начинатся с названия типа и _ 'location_ggeg777d377377e7f'
	*/

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (apstix AttackPatternDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &apstix); err != nil {
		return apstix, err
	}

	return apstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (apstix AttackPatternDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(apstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (apstix AttackPatternDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !apstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (cstix CampaignDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &cstix); err != nil {
		return nil, err
	}

	return cstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (cstix CampaignDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(cstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (cstix CampaignDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !cstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (castix CourseOfActionDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &castix); err != nil {
		return nil, err
	}

	return castix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (castix CourseOfActionDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(castix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (castix CourseOfActionDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !castix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (gstix GroupingDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &gstix); err != nil {
		return nil, err
	}

	return gstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (gstix GroupingDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(gstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (gstix GroupingDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !gstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (istix IdentityDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (istix IdentityDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (istix IdentityDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (istix IndicatorDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (istix IndicatorDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (istix IndicatorDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (istix InfrastructureDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (istix InfrastructureDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (istix InfrastructureDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (istix IntrusionSetDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (istix IntrusionSetDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (istix IntrusionSetDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (lstix LocationDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &lstix); err != nil {
		return nil, err
	}

	return lstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (lstix LocationDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(lstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (lstix LocationDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !lstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (mstix MalwareDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &mstix); err != nil {
		return nil, err
	}

	return mstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (mstix MalwareDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(mstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (mstix MalwareDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !mstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (mastix MalwareAnalysisDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &mastix); err != nil {
		return nil, err
	}

	return mastix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (mastix MalwareAnalysisDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(mastix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (mastix MalwareAnalysisDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !mastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (nstix NoteDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &nstix); err != nil {
		return nil, err
	}

	return nstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (nstix NoteDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(nstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (nstix NoteDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !nstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (odstix ObservedDataDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &odstix); err != nil {
		return nil, err
	}

	return odstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (odstix ObservedDataDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(odstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (odstix ObservedDataDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !odstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (ostix OpinionDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &ostix); err != nil {
		return nil, err
	}

	return ostix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (ostix OpinionDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ostix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (ostix OpinionDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !ostix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (rstix ReportDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &rstix); err != nil {
		return nil, err
	}

	return rstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (rstix ReportDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(rstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (rstix ReportDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !rstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (tastix ThreatActorDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &tastix); err != nil {
		return nil, err
	}

	return tastix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (tastix ThreatActorDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(tastix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (tastix ThreatActorDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !tastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (tstix ToolDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &tstix); err != nil {
		return nil, err
	}

	return tstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (tstix ToolDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(tstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (tstix ToolDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !tstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//DecoderJSON выполняет декодирование JSON объекта
func (vstix VulnerabilityDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &vstix); err != nil {
		return nil, err
	}

	return vstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (vstix VulnerabilityDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(vstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (vstix VulnerabilityDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !vstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}
