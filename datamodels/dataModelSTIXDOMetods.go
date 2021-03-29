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

	//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
	if ok := cpdostix.GranularMarkings.CheckGranularMarkingsTypeSTIX(); !ok {

		fmt.Println("\ttype CommonPropertiesDomainObjectSTIX - ERROR: 555")

		return false
	}

	return true
}

func (cpdostix CommonPropertiesDomainObjectSTIX) sanitizeStruct() CommonPropertiesDomainObjectSTIX {
	//обработка содержимого списка поля Labels
	if len(cpdostix.Labels) > 0 {
		nl := make([]string, 0, len(cpdostix.Labels))

		for _, l := range cpdostix.Labels {
			nl = append(nl, commonlibs.StringSanitize(l))
		}

		cpdostix.Labels = nl
	}

	//проверяем поле ObjectMarkingRefs
	if len(cpdostix.ObjectMarkingRefs) > 0 {
		newObjectMarkingRefs := make([]*IdentifierTypeSTIX, 0, len(cpdostix.ObjectMarkingRefs))
		for _, value := range cpdostix.ObjectMarkingRefs {
			tmpRes := commonlibs.StringSanitize(fmt.Sprint(value))
			value.AddValue(tmpRes)
			newObjectMarkingRefs = append(newObjectMarkingRefs, value)
		}
		cpdostix.ObjectMarkingRefs = newObjectMarkingRefs
	}

	//обработка содержимого списка поля Extension
	if len(cpdostix.Extensions) > 0 {
		newExtension := make(map[string]string, len(cpdostix.Extensions))
		for extKey, extValue := range cpdostix.Extensions {
			newExtension[extKey] = commonlibs.StringSanitize(extValue)
		}
		cpdostix.Extensions = newExtension
	}

	return cpdostix
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
			str += fmt.Sprintf("\texternal_references element '%d'\n", k)
			str += fmt.Sprintf("\tsource_name: '%s'\n", v.SourceName)
			str += fmt.Sprintf("\tdescription: '%s'\n", v.Description)
			str += fmt.Sprintf("\turl: '%s'\n", v.URL)
			str += fmt.Sprintf("\thashes: '%s'\n", v.Hashes)
			str += fmt.Sprintf("\texternal_id: '%s'\n", v.ExternalID)
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
	if !(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(apstix.ID)) {
		return false
	}

	//обязательное поле
	if apstix.Name == "" {
		return false
	}

	return apstix.checkingTypeCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (apstix AttackPatternDomainObjectsSTIX) SanitizeStruct() AttackPatternDomainObjectsSTIX {
	apstix.CommonPropertiesDomainObjectSTIX = apstix.sanitizeStruct()

	apstix.Name = commonlibs.StringSanitize(apstix.Name)
	apstix.Description = commonlibs.StringSanitize(apstix.Description)

	if len(apstix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(apstix.Aliases))
		for _, v := range apstix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		apstix.Aliases = aliasesTmp
	}

	apstix.KillChainPhases = apstix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	return apstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (apstix AttackPatternDomainObjectsSTIX) ToStringBeautiful() string {
	str := apstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += apstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", apstix.Name)
	str += fmt.Sprintf("description: '%s'\n", apstix.Description)
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(apstix.Aliases))
	str += fmt.Sprintf("kill_chain_phases: \n%v", func(l KillChainPhasesTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tkey:'%v' kill_chain_name: '%s'\n", k, v.KillChainName)
			str += fmt.Sprintf("\tkey:'%v' phase_name: '%s'\n", k, v.PhaseName)
		}
		return str
	}(apstix.KillChainPhases))

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
func (cstix CampaignDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(campaign--)[0-9a-f|-]+$`).MatchString(cstix.ID)) {
		return false
	}

	return cstix.checkingTypeCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (cstix CampaignDomainObjectsSTIX) SanitizeStruct() CampaignDomainObjectsSTIX {
	cstix.CommonPropertiesDomainObjectSTIX = cstix.sanitizeStruct()

	cstix.Name = commonlibs.StringSanitize(cstix.Name)
	cstix.Description = commonlibs.StringSanitize(cstix.Description)

	if len(cstix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(cstix.Aliases))
		for _, v := range cstix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		cstix.Aliases = aliasesTmp
	}

	cstix.Objective = commonlibs.StringSanitize(cstix.Objective)

	return cstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (cstix CampaignDomainObjectsSTIX) ToStringBeautiful() string {
	str := cstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += cstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", cstix.Name)
	str += fmt.Sprintf("description: '%s'\n", cstix.Description)
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(cstix.Aliases))
	str += fmt.Sprintf("first_seen: '%v'\n", cstix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", cstix.LastSeen)
	str += fmt.Sprintf("objective: '%s'\n", cstix.Objective)

	return str
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
func (castix CourseOfActionDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(course-of-action--)[0-9a-f|-]+$`).MatchString(castix.ID)) {
		return false
	}

	//обязательное поле
	if castix.Name == "" {
		return false
	}

	return castix.checkingTypeCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (castix CourseOfActionDomainObjectsSTIX) SanitizeStruct() CourseOfActionDomainObjectsSTIX {
	castix.CommonPropertiesDomainObjectSTIX = castix.sanitizeStruct()

	castix.Name = commonlibs.StringSanitize(castix.Name)
	castix.Description = commonlibs.StringSanitize(castix.Description)
	//cstix.Action - ЗАРЕЗЕРВИРОВАНО

	return castix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (castix CourseOfActionDomainObjectsSTIX) ToStringBeautiful() string {
	str := castix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += castix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", castix.Name)
	str += fmt.Sprintf("description: '%s'\n", castix.Description)
	str += fmt.Sprintf("action: '%v'\n", castix.Action)

	return str
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
func (gstix GroupingDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(grouping--)[0-9a-f|-]+$`).MatchString(gstix.ID)) {
		return false
	}

	if !gstix.checkingTypeCommonFields() {
		return false
	}

	//обязательное поле
	if gstix.Context == "" {
		return false
	}

	//обязательное поле
	if len(gstix.ObjectRefs) == 0 {
		return false
	}

	for _, v := range gstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (gstix GroupingDomainObjectsSTIX) SanitizeStruct() GroupingDomainObjectsSTIX {
	gstix.CommonPropertiesDomainObjectSTIX = gstix.sanitizeStruct()

	gstix.Name = commonlibs.StringSanitize(gstix.Name)
	gstix.Description = commonlibs.StringSanitize(gstix.Description)
	gstix.Context = gstix.Context.SanitizeStructOpenVocabTypeSTIX()

	return gstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (gstix GroupingDomainObjectsSTIX) ToStringBeautiful() string {
	str := gstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += gstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", gstix.Name)
	str += fmt.Sprintf("description: '%s'\n", gstix.Description)
	str += fmt.Sprintf("context: '%s'\n", gstix.Context)
	str += fmt.Sprintf("object_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(gstix.ObjectRefs))

	return str
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
func (istix IdentityDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(identity--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	//обязательное поле
	if istix.Name == "" {
		return false
	}

	return istix.checkingTypeCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix IdentityDomainObjectsSTIX) SanitizeStruct() IdentityDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.sanitizeStruct()

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.Roles) > 0 {
		rolesTmp := make([]string, 0, len(istix.Roles))
		for _, v := range istix.Roles {
			rolesTmp = append(rolesTmp, commonlibs.StringSanitize(v))
		}
		istix.Roles = rolesTmp
	}

	istix.IdentityClass = istix.IdentityClass.SanitizeStructOpenVocabTypeSTIX()

	if len(istix.Sectors) > 0 {
		sectorsTmp := make([]*OpenVocabTypeSTIX, 0, len(istix.Sectors))

		for _, v := range istix.Sectors {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			sectorsTmp = append(sectorsTmp, &tmp)
		}

		istix.Sectors = sectorsTmp
	}

	istix.ContactInformation = commonlibs.StringSanitize(istix.ContactInformation)

	return istix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix IdentityDomainObjectsSTIX) ToStringBeautiful() string {
	str := istix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", istix.Name)
	str += fmt.Sprintf("description: '%s'\n", istix.Description)
	str += fmt.Sprintf("roles: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\trole '%d': '%v'\n", k, v)
		}
		return str
	}(istix.Roles))
	str += fmt.Sprintf("identity_class: '%s'\n", istix.IdentityClass)
	str += fmt.Sprintf("sectors: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsector '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.Sectors))
	str += fmt.Sprintf("contact_information: '%s'\n", istix.ContactInformation)

	return str
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
func (istix IndicatorDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(indicator--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//необходимое поле
	if istix.Pattern == "" {
		return false
	}

	//обязательное поле
	if istix.PatternType == "" {
		return false
	}

	//обязательное поле
	if istix.ValidFrom.Unix() < 0 {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix IndicatorDomainObjectsSTIX) SanitizeStruct() IndicatorDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.sanitizeStruct()

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.IndicatorTypes) > 0 {
		it := make([]*OpenVocabTypeSTIX, 0, len(istix.IndicatorTypes))

		for _, v := range istix.IndicatorTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			it = append(it, &tmp)
		}

		istix.IndicatorTypes = it
	}

	istix.Pattern = commonlibs.StringSanitize(istix.Pattern)
	istix.PatternType = istix.PatternType.SanitizeStructOpenVocabTypeSTIX()
	istix.PatternVersion = commonlibs.StringSanitize(istix.PatternVersion)
	istix.KillChainPhases = istix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	return istix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix IndicatorDomainObjectsSTIX) ToStringBeautiful() string {
	str := istix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", istix.Name)
	str += fmt.Sprintf("description: '%s'\n", istix.Description)
	str += fmt.Sprintf("indicator_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tindicator_type '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.IndicatorTypes))
	str += fmt.Sprintf("pattern: '%s'\n", istix.Pattern)
	str += fmt.Sprintf("pattern_type: '%s'\n", istix.PatternType)
	str += fmt.Sprintf("pattern_version: '%s'\n", istix.PatternVersion)
	str += fmt.Sprintf("valid_from: '%v'\n", istix.ValidFrom)
	str += fmt.Sprintf("valid_until: '%v'\n", istix.ValidUntil)
	str += fmt.Sprintf("sectors: \n%v", func(l []*KillChainPhasesTypeElementSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsector '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.KillChainPhases))

	return str
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
func (istix InfrastructureDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(infrastructure--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//обязательное поле
	if istix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix InfrastructureDomainObjectsSTIX) SanitizeStruct() InfrastructureDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.sanitizeStruct()

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.InfrastructureTypes) > 0 {
		it := make([]*OpenVocabTypeSTIX, 0, len(istix.InfrastructureTypes))

		for _, v := range istix.InfrastructureTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			it = append(it, &tmp)
		}

		istix.InfrastructureTypes = it
	}

	if len(istix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(istix.Aliases))
		for _, v := range istix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		istix.Aliases = aliasesTmp
	}

	istix.KillChainPhases = istix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	return istix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix InfrastructureDomainObjectsSTIX) ToStringBeautiful() string {
	str := istix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", istix.Name)
	str += fmt.Sprintf("description: '%s'\n", istix.Description)
	str += fmt.Sprintf("infrastructure_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tinfrastructure_type '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.InfrastructureTypes))
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(istix.Aliases))
	str += fmt.Sprintf("sectors: \n%v", func(l []*KillChainPhasesTypeElementSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsector '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.KillChainPhases))
	str += fmt.Sprintf("first_seen: '%v'\n", istix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", istix.LastSeen)

	return str
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
func (istix IntrusionSetDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(intrusion-set--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	if !istix.checkingTypeCommonFields() {
		return false
	}

	//обязательное поле
	if istix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix IntrusionSetDomainObjectsSTIX) SanitizeStruct() IntrusionSetDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.sanitizeStruct()

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(istix.Aliases))
		for _, v := range istix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		istix.Aliases = aliasesTmp
	}

	if len(istix.Goals) > 0 {
		goalsTmp := make([]string, 0, len(istix.Goals))
		for _, v := range istix.Goals {
			goalsTmp = append(goalsTmp, commonlibs.StringSanitize(v))
		}
		istix.Goals = goalsTmp
	}

	istix.ResourceLevel = istix.ResourceLevel.SanitizeStructOpenVocabTypeSTIX()
	istix.PrimaryMotivation = istix.PrimaryMotivation.SanitizeStructOpenVocabTypeSTIX()

	if len(istix.SecondaryMotivations) > 0 {
		sm := make([]*OpenVocabTypeSTIX, 0, len(istix.SecondaryMotivations))

		for _, v := range istix.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			sm = append(sm, &tmp)
		}

		istix.SecondaryMotivations = sm
	}

	return istix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix IntrusionSetDomainObjectsSTIX) ToStringBeautiful() string {
	str := istix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", istix.Name)
	str += fmt.Sprintf("description: '%s'\n", istix.Description)
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(istix.Aliases))
	str += fmt.Sprintf("first_seen: '%v'\n", istix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", istix.LastSeen)
	str += fmt.Sprintf("goals: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tgoal '%d': '%s'\n", k, v)
		}
		return str
	}(istix.Goals))
	str += fmt.Sprintf("resource_level: '%s'\n", istix.FirstSeen)
	str += fmt.Sprintf("primary_motivation: '%s'\n", istix.LastSeen)
	str += fmt.Sprintf("secondary_motivations: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsecondary_motivation '%d': '%v'\n", k, *v)
		}
		return str
	}(istix.SecondaryMotivations))

	return str
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
func (lstix LocationDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(location--)[0-9a-f|-]+$`).MatchString(lstix.ID)) {
		return false
	}

	if !lstix.checkingTypeCommonFields() {
		return false
	}

	if (lstix.Latitude > 90.0) || (lstix.Latitude < -90.0) {
		return false
	}

	if (lstix.Longitude > 180.0) || (lstix.Longitude < -180.0) {
		return false
	}

	if !(regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(lstix.Country)) {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (lstix LocationDomainObjectsSTIX) SanitizeStruct() LocationDomainObjectsSTIX {
	lstix.CommonPropertiesDomainObjectSTIX = lstix.sanitizeStruct()

	lstix.Name = commonlibs.StringSanitize(lstix.Name)
	lstix.Description = commonlibs.StringSanitize(lstix.Description)
	lstix.Region = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(lstix.Region)))
	lstix.AdministrativeArea = commonlibs.StringSanitize(lstix.AdministrativeArea)
	lstix.City = commonlibs.StringSanitize(lstix.City)
	lstix.StreetAddress = commonlibs.StringSanitize(lstix.StreetAddress)
	lstix.PostalCode = commonlibs.StringSanitize(lstix.PostalCode)

	return lstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (lstix LocationDomainObjectsSTIX) ToStringBeautiful() string {
	str := lstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += lstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", lstix.Name)
	str += fmt.Sprintf("description: '%s'\n", lstix.Description)
	str += fmt.Sprintf("latitude: '%v'\n", lstix.Latitude)
	str += fmt.Sprintf("longitude: '%v'\n", lstix.Longitude)
	str += fmt.Sprintf("precision: '%v'\n", lstix.Precision)
	str += fmt.Sprintf("region: '%s'\n", lstix.Region)
	str += fmt.Sprintf("country: '%s'\n", lstix.Country)
	str += fmt.Sprintf("administrative_area: '%s'\n", lstix.AdministrativeArea)
	str += fmt.Sprintf("city: '%s'\n", lstix.City)
	str += fmt.Sprintf("street_address: '%s'\n", lstix.StreetAddress)
	str += fmt.Sprintf("postal_code: '%s'\n", lstix.PostalCode)

	return str
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
func (mstix MalwareDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(malware--)[0-9a-f|-]+$`).MatchString(mstix.ID)) {
		return false
	}

	if !mstix.checkingTypeCommonFields() {
		return false
	}

	if len(mstix.OperatingSystemRefs) > 0 {
		for _, v := range mstix.OperatingSystemRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(mstix.SampleRefs) > 0 {
		for _, v := range mstix.SampleRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mstix MalwareDomainObjectsSTIX) SanitizeStruct() MalwareDomainObjectsSTIX {
	mstix.CommonPropertiesDomainObjectSTIX = mstix.sanitizeStruct()

	mstix.Name = commonlibs.StringSanitize(mstix.Name)
	mstix.Description = commonlibs.StringSanitize(mstix.Description)

	if len(mstix.MalwareTypes) > 0 {
		mt := make([]*OpenVocabTypeSTIX, 0, len(mstix.MalwareTypes))

		for _, v := range mstix.MalwareTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			mt = append(mt, &tmp)
		}

		mstix.MalwareTypes = mt
	}

	if len(mstix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(mstix.Aliases))
		for _, v := range mstix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		mstix.Aliases = aliasesTmp
	}

	mstix.KillChainPhases = mstix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	if len(mstix.ArchitectureExecutionEnvs) > 0 {
		aee := make([]*OpenVocabTypeSTIX, 0, len(mstix.ArchitectureExecutionEnvs))

		for _, v := range mstix.ArchitectureExecutionEnvs {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			aee = append(aee, &tmp)
		}

		mstix.ArchitectureExecutionEnvs = aee
	}

	if len(mstix.ImplementationLanguages) > 0 {
		il := make([]*OpenVocabTypeSTIX, 0, len(mstix.ImplementationLanguages))

		for _, v := range mstix.ImplementationLanguages {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			il = append(il, &tmp)
		}

		mstix.ImplementationLanguages = il
	}

	if len(mstix.Capabilities) > 0 {
		c := make([]*OpenVocabTypeSTIX, 0, len(mstix.Capabilities))

		for _, v := range mstix.Capabilities {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			c = append(c, &tmp)
		}

		mstix.Capabilities = c
	}

	return mstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mstix MalwareDomainObjectsSTIX) ToStringBeautiful() string {
	str := mstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += mstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", mstix.Name)
	str += fmt.Sprintf("description: '%s'\n", mstix.Description)
	str += fmt.Sprintf("malware_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tmalware_type '%d': '%v'\n", k, *v)
		}
		return str
	}(mstix.MalwareTypes))
	str += fmt.Sprintf("is_family: '%v'\n", mstix.IsFamily)
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(mstix.Aliases))
	str += fmt.Sprintf("kill_chain_phases: \n%v", func(l KillChainPhasesTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tkey:'%v' kill_chain_name: '%s'\n", k, v.KillChainName)
			str += fmt.Sprintf("\tkey:'%v' phase_name: '%s'\n", k, v.PhaseName)
		}
		return str
	}(mstix.KillChainPhases))
	str += fmt.Sprintf("first_seen: '%v'\n", mstix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", mstix.LastSeen)
	str += fmt.Sprintf("operating_system_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\toperating_system_ref '%d': '%v'\n", k, v)
		}
		return str
	}(mstix.OperatingSystemRefs))
	str += fmt.Sprintf("architecture_execution_envs: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tarchitecture_execution_env '%d': '%v'\n", k, v)
		}
		return str
	}(mstix.ArchitectureExecutionEnvs))
	str += fmt.Sprintf("implementation_languages: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\timplementation_language '%d': '%v'\n", k, v)
		}
		return str
	}(mstix.ImplementationLanguages))
	str += fmt.Sprintf("capabilities: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcapabilitie '%d': '%v'\n", k, v)
		}
		return str
	}(mstix.Capabilities))
	str += fmt.Sprintf("sample_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsample_ref '%d': '%v'\n", k, v)
		}
		return str
	}(mstix.SampleRefs))

	return str
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
func (mastix MalwareAnalysisDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(malware-analysis--)[0-9a-f|-]+$`).MatchString(mastix.ID)) {
		return false
	}

	if !mastix.checkingTypeCommonFields() {
		return false
	}

	if mastix.Product == "" {
		return false
	}

	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(mastix.Version)) {
		return false
	}

	if !mastix.HostVMRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !mastix.OperatingSystemRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(mastix.InstalledSoftwareRefs) > 0 {
		for _, v := range mastix.InstalledSoftwareRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(mastix.AnalysisScoRefs) > 0 {
		for _, v := range mastix.AnalysisScoRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !mastix.SampleRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mastix MalwareAnalysisDomainObjectsSTIX) SanitizeStruct() MalwareAnalysisDomainObjectsSTIX {
	mastix.CommonPropertiesDomainObjectSTIX = mastix.sanitizeStruct()

	mastix.Product = commonlibs.StringSanitize(mastix.Product)
	mastix.ConfigurationVersion = commonlibs.StringSanitize(mastix.ConfigurationVersion)
	if len(mastix.Modules) > 0 {
		mTmp := make([]string, 0, len(mastix.Modules))
		for _, v := range mastix.Modules {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		mastix.Modules = mTmp
	}
	mastix.AnalysisEngineVersion = commonlibs.StringSanitize(mastix.AnalysisEngineVersion)
	mastix.AnalysisDefinitionVersion = commonlibs.StringSanitize(mastix.AnalysisDefinitionVersion)
	mastix.ResultName = commonlibs.StringSanitize(mastix.ResultName)
	mastix.Result = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(mastix.Result)))

	/*

	   Надо все проверить еще раз и сделать ToStringBeautiful прежде чем тестировать

	*/

	return mastix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mastix MalwareAnalysisDomainObjectsSTIX) ToStringBeautiful() string {
	str := mastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += mastix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	/*
		Product                   string                `json:"product" bson:"product" required:"true"`
			Version                   string                `json:"version" bson:"version"`
			HostVMRef                 IdentifierTypeSTIX    `json:"host_vm_ref" bson:"host_vm_ref"`
			OperatingSystemRef        IdentifierTypeSTIX    `json:"operating_system_ref" bson:"operating_system_ref"`
			InstalledSoftwareRefs     []*IdentifierTypeSTIX `json:"installed_software_refs" bson:"installed_software_refs"`
			ConfigurationVersion      string                `json:"configuration_version" bson:"configuration_version"`
			Modules                   []string              `json:"modules" bson:"modules"`
			AnalysisEngineVersion     string                `json:"analysis_engine_version" bson:"analysis_engine_version"`
			AnalysisDefinitionVersion string                `json:"analysis_definition_version" bson:"analysis_definition_version"`
			Submitted                 time.Time             `json:"submitted" bson:"submitted"`
			AnalysisStarted           time.Time             `json:"analysis_started" bson:"analysis_started"`
			AnalysisEnded             time.Time             `json:"analysis_ended" bson:"analysis_ended"`
			ResultName                string                `json:"result_name" bson:"result_name"`
			Result                    OpenVocabTypeSTIX     `json:"result" bson:"result"`
			AnalysisScoRefs           []*IdentifierTypeSTIX `json:"analysis_sco_refs" bson:"analysis_sco_refs"`
			SampleRef                 IdentifierTypeSTIX    `json:"sample_ref" bson:"sample_ref"`


				str += fmt.Sprintf("name: '%s'\n", mstix.Name)
				str += fmt.Sprintf("description: '%s'\n", mstix.Description)
				str += fmt.Sprintf("malware_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\tmalware_type '%d': '%v'\n", k, *v)
					}
					return str
				}(mstix.MalwareTypes))
				str += fmt.Sprintf("is_family: '%v'\n", mstix.IsFamily)
				str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
					}
					return str
				}(mstix.Aliases))
				str += fmt.Sprintf("kill_chain_phases: \n%v", func(l KillChainPhasesTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\tkey:'%v' kill_chain_name: '%s'\n", k, v.KillChainName)
						str += fmt.Sprintf("\tkey:'%v' phase_name: '%s'\n", k, v.PhaseName)
					}
					return str
				}(mstix.KillChainPhases))
				str += fmt.Sprintf("first_seen: '%v'\n", mstix.FirstSeen)
				str += fmt.Sprintf("last_seen: '%v'\n", mstix.LastSeen)
				str += fmt.Sprintf("operating_system_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\toperating_system_ref '%d': '%v'\n", k, v)
					}
					return str
				}(mstix.OperatingSystemRefs))
				str += fmt.Sprintf("architecture_execution_envs: \n%v", func(l []*OpenVocabTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\tarchitecture_execution_env '%d': '%v'\n", k, v)
					}
					return str
				}(mstix.ArchitectureExecutionEnvs))
				str += fmt.Sprintf("implementation_languages: \n%v", func(l []*OpenVocabTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\timplementation_language '%d': '%v'\n", k, v)
					}
					return str
				}(mstix.ImplementationLanguages))
				str += fmt.Sprintf("capabilities: \n%v", func(l []*OpenVocabTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\tcapabilitie '%d': '%v'\n", k, v)
					}
					return str
				}(mstix.Capabilities))
				str += fmt.Sprintf("sample_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
					var str string
					for k, v := range l {
						str += fmt.Sprintf("\tsample_ref '%d': '%v'\n", k, v)
					}
					return str
				}(mstix.SampleRefs))
	*/
	return str
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
