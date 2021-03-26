package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"

	"ISEMS-MRSICT/commonlibs"
)

/*************************************************************************/
/********** 			Domain Objects STIX (МЕТОДЫ)			**********/
/*************************************************************************/

func (cpdostix *CommonPropertiesDomainObjectSTIX) checkingTypeCommonFields() bool {
	fmt.Println("func 'checkingTypeCommonFields', START...")

	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(cpdostix.SpecVersion)) {

		fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 111")

		return false
	}

	//валидация содержимого поля CreatedByRef
	if len(fmt.Sprint(cpdostix.CreatedByRef)) > 0 {
		if !(regexp.MustCompile(`^[0-9a-zA-Z-_]+(--)[0-9a-f|-]+$`).MatchString(fmt.Sprint(cpdostix.CreatedByRef))) {

			fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 222")

			return false
		}
	}

	//обработка содержимого списка поля Labels
	if len(cpdostix.Labels) > 0 {
		nl := make([]string, len(cpdostix.Labels))

		for _, l := range cpdostix.Labels {
			nl = append(nl, commonlibs.StringSanitize(l))
		}

		cpdostix.Labels = nl
	}

	//для поля Lang
	if len(cpdostix.Lang) > 0 {
		if !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(cpdostix.Lang) {

			fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 333")

			return false
		}
	}
	//вызываем метод проверки полей типа ExternalReferences
	if ok := cpdostix.ExternalReferences.CheckExternalReferencesTypeSTIX(); !ok {

		fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 444")

		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(cpdostix.ObjectMarkingRefs) > 0 {
		newObjectMarkingRefs := make([]*IdentifierTypeSTIX, len(cpdostix.ObjectMarkingRefs))
		for _, value := range cpdostix.ObjectMarkingRefs {
			tmpRes := commonlibs.StringSanitize(fmt.Sprint(value))
			value.AddValue(tmpRes)
			newObjectMarkingRefs = append(newObjectMarkingRefs, value)
		}
		cpdostix.ObjectMarkingRefs = newObjectMarkingRefs
	}

	//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
	if ok := cpdostix.GranularMarkings.CheckGranularMarkingsTypeSTIX(); !ok {

		fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 555")

		return false
	}

	//обработка содержимого списка поля Extension
	newExtension := make(map[string]string, len(cpdostix.Extensions))
	for extKey, extValue := range cpdostix.Extensions {
		newExtension[extKey] = commonlibs.StringSanitize(extValue)
	}
	cpdostix.Extensions = newExtension

	return true
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (cp CommonPropertiesDomainObjectSTIX) ToStringBeautiful() string {
	var str string
	str += fmt.Sprintf("spec_version: '%s'\n", cp.SpecVersion)
	str += fmt.Sprintf("created: '%v'\n", cp.Created)
	str += fmt.Sprintf("modified: '%v'\n", cp.Modified)
	str += fmt.Sprintf("created_by_ref: '%s'\n", cp.CreatedByRef)
	str += fmt.Sprintf("revoked: '%v'\n", cp.Revoked)
	str += fmt.Sprintf("labels: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tlabel '%d': '%s'\n", k, v)
		}
		return str
	}(cp.Labels))
	str += fmt.Sprintf("external_references: \n%v", func(l []*ExternalReferenceTypeElementSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t\texternal_references element '%d'\n", k)
			str += fmt.Sprintf("\t\tsource_name: '%s'\n", v.SourceName)
			str += fmt.Sprintf("\t\tdescription: '%s'\n", v.Description)
			str += fmt.Sprintf("\t\turl: '%s'\n", v.URL)
			str += fmt.Sprintf("\t\thashes: '%s'\n", v.Hashes)
			str += fmt.Sprintf("\t\texternal_id: '%s'\n", v.ExternalID)
		}
		return str
	}(cp.ExternalReferences))
	str += fmt.Sprintf("object_marking_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tref '%d': '%v'\n", k, v)
		}
		return str
	}(cp.ObjectMarkingRefs))
	str += fmt.Sprintln("granular_markings:")
	str += fmt.Sprintf("\tlang: '%s'\n", cp.GranularMarkings.Lang)
	str += fmt.Sprintf("\tmarking_ref: '%v'\n", cp.GranularMarkings.MarkingRef)
	str += fmt.Sprintf("\tselectors: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t\tselector '%d': '%s'\n", k, v)
		}
		return str
	}(cp.GranularMarkings.Selectors))
	str += fmt.Sprintf("defanged: '%v'\n", cp.Defanged)
	str += fmt.Sprintf("extensions: \n%v", func(l map[string]string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%s'\n", k, v)
		}
		return str
	}(cp.Extensions))

	return str
}

/* --- AttackPatternDomainObjectsSTIX --- */

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
// возвращает ВАЛИДНЫЙ объект AttackPatternDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.AttackPatternDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (apstix AttackPatternDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(apstix.ID)) {
		return false
	}

	if !apstix.checkingTypeCommonFields() {
		return false
	}

	apstix.Name = commonlibs.StringSanitize(apstix.Name)
	apstix.Description = commonlibs.StringSanitize(apstix.Description)

	aliasesTmp := make([]string, 0, len(apstix.Aliases))
	for _, v := range apstix.Aliases {
		aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
	}
	apstix.Aliases = aliasesTmp

	return apstix.KillChainPhases.CheckKillChainPhasesTypeSTIX()
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ap AttackPatternDomainObjectsSTIX) ToStringBeautiful() string {
	str := ap.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ap.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", ap.Name)
	str += fmt.Sprintf("description: '%s'\n", ap.Description)
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(ap.Aliases))
	str += fmt.Sprintf("kill_chain_phases: \n%v", func(l KillChainPhasesTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tkey:'%v' kill_chain_name: '%s'\n", k, v.KillChainName)
			str += fmt.Sprintf("\tkey:'%v' phase_name: '%s'\n", k, v.PhaseName)
		}
		return str
	}(ap.KillChainPhases))

	return str
}

/* --- CampaignDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе CampaignDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект CampaignDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.CampaignDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (cstix CampaignDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !cstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- CourseOfActionDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе CourseOfActionDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект CourseOfActionDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.CourseOfActionDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (castix CourseOfActionDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !castix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- GroupingDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе GroupingDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект GroupingDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.GroupingDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (gstix GroupingDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !gstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- IdentityDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе IdentityDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект IdentityDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.IdentityDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (istix IdentityDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- IndicatorDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе IndicatorDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект IndicatorDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.IndicatorDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (istix IndicatorDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- InfrastructureDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе InfrastructureDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект InfrastructureDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.InfrastructureDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (istix InfrastructureDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- IntrusionSetDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе IntrusionSetDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект IntrusionSetDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.IntrusionSetDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (istix IntrusionSetDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- LocationDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе LocationDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект LocationDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.LocationDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (lstix LocationDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !lstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- MalwareDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе MalwareDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект MalwareDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.MalwareDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (mstix MalwareDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !mstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- MalwareAnalysisDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе MalwareAnalysisDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект MalwareAnalysisDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.MalwareAnalysisDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (mastix MalwareAnalysisDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !mastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- NoteDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе NoteDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект NoteDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.NoteDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (nstix NoteDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !nstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- ObservedDataDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ObservedDataDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект ObservedDataDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ObservedDataDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (odstix ObservedDataDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !odstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- OpinionDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе OpinionDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект OpinionDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.OpinionDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (ostix OpinionDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !ostix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- ReportDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ReportDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект ReportDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ReportDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (rstix ReportDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !rstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- ThreatActorDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ThreatActorDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект ThreatActorDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ThreatActorDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (tastix ThreatActorDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !tastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- ToolDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ToolDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект ToolDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ToolDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (tstix ToolDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !tstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- VulnerabilityDomainObjectsSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе VulnerabilityDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект VulnerabilityDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.VulnerabilityDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (vstix VulnerabilityDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !vstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}
