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
		return false
	}

	//валидация содержимого поля CreatedByRef
	if len(fmt.Sprint(cpdostix.CreatedByRef)) > 0 {
		if !(regexp.MustCompile(`^[0-9a-zA-Z-_]+(--)[0-9a-f|-]+$`).MatchString(fmt.Sprint(cpdostix.CreatedByRef))) {
			return false
		}
	}

	//для поля Lang
	if len(cpdostix.Lang) > 0 {
		if !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(cpdostix.Lang) {
			return false
		}
	}
	//вызываем метод проверки полей типа ExternalReferences
	if ok := cpdostix.ExternalReferences.CheckExternalReferencesTypeSTIX(); !ok {
		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(cpdostix.ObjectMarkingRefs) > 0 {
		for _, value := range cpdostix.ObjectMarkingRefs {
			if !value.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
	if ok := cpdostix.GranularMarkings.CheckGranularMarkingsTypeSTIX(); !ok {
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

	//обработка содержимого списка поля ExternalReferences
	cpdostix.ExternalReferences = cpdostix.ExternalReferences.SanitizeStructExternalReferencesTypeSTIX()

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
			str += fmt.Sprintf("\texternal_references element '%d':\n", k)
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
	if !(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(apstix.ID)) {
		return false
	}

	//обязательное поле
	if apstix.Name == "" {

		fmt.Println("AttackPatternDomainObjectsSTIX ERROR 111")

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
			str += fmt.Sprintf("\toperating_system_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(mstix.OperatingSystemRefs))
	str += fmt.Sprintf("architecture_execution_envs: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tarchitecture_execution_env '%d': '%v'\n", k, *v)
		}
		return str
	}(mstix.ArchitectureExecutionEnvs))
	str += fmt.Sprintf("implementation_languages: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\timplementation_language '%d': '%v'\n", k, *v)
		}
		return str
	}(mstix.ImplementationLanguages))
	str += fmt.Sprintf("capabilities: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcapabilitie '%d': '%v'\n", k, *v)
		}
		return str
	}(mstix.Capabilities))
	str += fmt.Sprintf("sample_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsample_ref '%d': '%v'\n", k, *v)
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

	if mastix.Version != "" && !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(mastix.Version)) {
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
	mastix.AvResult = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(mastix.AvResult)))

	return mastix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mastix MalwareAnalysisDomainObjectsSTIX) ToStringBeautiful() string {
	str := mastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += mastix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("product: '%s'\n", mastix.Product)
	str += fmt.Sprintf("version: '%s'\n", mastix.Version)
	str += fmt.Sprintf("host_vm_ref: '%s'\n", mastix.HostVMRef)
	str += fmt.Sprintf("operating_system_ref: '%s'\n", mastix.OperatingSystemRef)
	str += fmt.Sprintf("installed_software_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tinstalled_software_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(mastix.InstalledSoftwareRefs))
	str += fmt.Sprintf("configuration_version: '%s'\n", mastix.ConfigurationVersion)
	str += fmt.Sprintf("modules: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tmodule '%d': '%s'\n", k, v)
		}
		return str
	}(mastix.Modules))
	str += fmt.Sprintf("analysis_engine_version: '%s'\n", mastix.AnalysisEngineVersion)
	str += fmt.Sprintf("analysis_definition_version: '%s'\n", mastix.AnalysisDefinitionVersion)
	str += fmt.Sprintf("submitted: '%v'\n", mastix.Submitted)
	str += fmt.Sprintf("analysis_started: '%v'\n", mastix.AnalysisStarted)
	str += fmt.Sprintf("analysis_ended: '%v'\n", mastix.AnalysisEnded)
	str += fmt.Sprintf("result_name: '%s'\n", mastix.ResultName)
	str += fmt.Sprintf("result: '%s'\n", mastix.Result)
	str += fmt.Sprintf("analysis_sco_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tanalysis_sco_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(mastix.AnalysisScoRefs))
	str += fmt.Sprintf("sample_ref: '%v'\n", mastix.SampleRef)
	str += fmt.Sprintf("av_result: '%v'\n", mastix.AvResult)

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
func (nstix NoteDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(note--)[0-9a-f|-]+$`).MatchString(nstix.ID)) {
		return false
	}

	if !nstix.checkingTypeCommonFields() {
		return false
	}

	if nstix.Content == "" || len(nstix.ObjectRefs) == 0 {
		return false
	}

	for _, v := range nstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (nstix NoteDomainObjectsSTIX) SanitizeStruct() NoteDomainObjectsSTIX {
	nstix.CommonPropertiesDomainObjectSTIX = nstix.sanitizeStruct()

	nstix.Abstract = commonlibs.StringSanitize(nstix.Abstract)
	nstix.Content = commonlibs.StringSanitize(nstix.Content)

	if len(nstix.Authors) > 0 {
		mTmp := make([]string, 0, len(nstix.Authors))
		for _, v := range nstix.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		nstix.Authors = mTmp
	}

	return nstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (nstix NoteDomainObjectsSTIX) ToStringBeautiful() string {
	str := nstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += nstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("abstract: '%s'\n", nstix.Abstract)
	str += fmt.Sprintf("content: '%s'\n", nstix.Content)
	str += fmt.Sprintf("authors: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tauthor '%d': '%s'\n", k, v)
		}
		return str
	}(nstix.Authors))
	str += fmt.Sprintf("object_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(nstix.ObjectRefs))

	return str
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
func (odstix ObservedDataDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(observed-data--)[0-9a-f|-]+$`).MatchString(odstix.ID)) {
		return false
	}

	if !odstix.checkingTypeCommonFields() {
		return false
	}

	if len(odstix.ObjectRefs) > 0 {
		for _, v := range odstix.ObjectRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (odstix ObservedDataDomainObjectsSTIX) SanitizeStruct() ObservedDataDomainObjectsSTIX {
	odstix.CommonPropertiesDomainObjectSTIX = odstix.sanitizeStruct()

	return odstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (odstix ObservedDataDomainObjectsSTIX) ToStringBeautiful() string {
	str := odstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += odstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("first_observed: '%v'\n", odstix.FirstObserved)
	str += fmt.Sprintf("last_observed: '%v'\n", odstix.LastObserved)
	str += fmt.Sprintf("number_observed: '%d'\n", odstix.NumberObserved)
	str += fmt.Sprintf("object_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(odstix.ObjectRefs))

	return str
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
func (ostix OpinionDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(opinion--)[0-9a-f|-]+$`).MatchString(ostix.ID)) {
		return false
	}

	if !ostix.checkingTypeCommonFields() {
		return false
	}

	if ostix.Opinion == "" {
		return false
	}

	if len(ostix.ObjectRefs) == 0 {
		return false
	}

	for _, v := range ostix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ostix OpinionDomainObjectsSTIX) SanitizeStruct() OpinionDomainObjectsSTIX {
	ostix.CommonPropertiesDomainObjectSTIX = ostix.sanitizeStruct()

	ostix.Explanation = commonlibs.StringSanitize(ostix.Explanation)

	if len(ostix.Authors) > 0 {
		mTmp := make([]string, 0, len(ostix.Authors))
		for _, v := range ostix.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		ostix.Authors = mTmp
	}

	ostix.Opinion = EnumTypeSTIX(commonlibs.StringSanitize(string(ostix.Opinion)))

	return ostix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ostix OpinionDomainObjectsSTIX) ToStringBeautiful() string {
	str := ostix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ostix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("explanation: '%v'\n", ostix.Explanation)
	str += fmt.Sprintf("authors: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tauthor '%d': '%s'\n", k, v)
		}
		return str
	}(ostix.Authors))
	str += fmt.Sprintf("opinion: '%v'\n", ostix.Opinion)
	str += fmt.Sprintf("object_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(ostix.ObjectRefs))

	return str
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
func (rstix ReportDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(report--)[0-9a-f|-]+$`).MatchString(rstix.ID)) {
		return false
	}

	if !rstix.checkingTypeCommonFields() {
		return false
	}

	//обязательное поле
	if rstix.Name == "" {
		return false
	}

	//обязательное поле
	if len(rstix.ObjectRefs) == 0 {
		return false
	}

	for _, v := range rstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (rstix ReportDomainObjectsSTIX) SanitizeStruct() ReportDomainObjectsSTIX {
	rstix.CommonPropertiesDomainObjectSTIX = rstix.sanitizeStruct()

	rstix.Name = commonlibs.StringSanitize(rstix.Name)
	rstix.Description = commonlibs.StringSanitize(rstix.Description)
	if len(rstix.ReportTypes) > 0 {
		r := make([]*OpenVocabTypeSTIX, 0, len(rstix.ReportTypes))
		for _, v := range rstix.ReportTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			r = append(r, &tmp)
		}

		rstix.ReportTypes = r
	}

	return rstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (rstix ReportDomainObjectsSTIX) ToStringBeautiful() string {
	str := rstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += rstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", rstix.Name)
	str += fmt.Sprintf("description: '%s'\n", rstix.Description)
	str += fmt.Sprintf("report_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\treport_type '%d': '%v'\n", k, *v)
		}
		return str
	}(rstix.ReportTypes))
	str += fmt.Sprintf("published: '%v'\n", rstix.Published)
	str += fmt.Sprintf("object_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(rstix.ObjectRefs))

	return str
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
func (tastix ThreatActorDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(threat-actor--)[0-9a-f|-]+$`).MatchString(tastix.ID)) {
		return false
	}

	if !tastix.checkingTypeCommonFields() {
		return false
	}

	if tastix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (tastix ThreatActorDomainObjectsSTIX) SanitizeStruct() ThreatActorDomainObjectsSTIX {
	tastix.CommonPropertiesDomainObjectSTIX = tastix.sanitizeStruct()

	tastix.Name = commonlibs.StringSanitize(tastix.Name)
	tastix.Description = commonlibs.StringSanitize(tastix.Description)

	if len(tastix.ThreatActorTypes) > 0 {
		ta := make([]*OpenVocabTypeSTIX, 0, len(tastix.ThreatActorTypes))
		for _, v := range tastix.ThreatActorTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, &tmp)
		}

		tastix.ThreatActorTypes = ta
	}

	if len(tastix.Aliases) > 0 {
		mTmp := make([]string, 0, len(tastix.Aliases))
		for _, v := range tastix.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tastix.Aliases = mTmp
	}

	if len(tastix.Roles) > 0 {
		ta := make([]*OpenVocabTypeSTIX, 0, len(tastix.Roles))
		for _, v := range tastix.Roles {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, &tmp)
		}

		tastix.Roles = ta
	}

	if len(tastix.Goals) > 0 {
		mTmp := make([]string, 0, len(tastix.Goals))
		for _, v := range tastix.Goals {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tastix.Goals = mTmp
	}

	tastix.Sophistication = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.Sophistication)))
	tastix.ResourceLevel = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.ResourceLevel)))
	tastix.PrimaryMotivation = OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.PrimaryMotivation)))

	if len(tastix.SecondaryMotivations) > 0 {
		ta := make([]*OpenVocabTypeSTIX, 0, len(tastix.SecondaryMotivations))
		for _, v := range tastix.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, &tmp)
		}

		tastix.Roles = ta
	}

	if len(tastix.PersonalMotivations) > 0 {
		ta := make([]*OpenVocabTypeSTIX, 0, len(tastix.PersonalMotivations))
		for _, v := range tastix.PersonalMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, &tmp)
		}

		tastix.Roles = ta
	}

	return tastix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (tastix ThreatActorDomainObjectsSTIX) ToStringBeautiful() string {
	str := tastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += tastix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", tastix.Name)
	str += fmt.Sprintf("description: '%s'\n", tastix.Description)
	str += fmt.Sprintf("threat_actor_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tthreat_actor_type '%d': '%v'\n", k, *v)
		}
		return str
	}(tastix.ThreatActorTypes))
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(tastix.Aliases))
	str += fmt.Sprintf("first_seen: '%v'\n", tastix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", tastix.LastSeen)
	str += fmt.Sprintf("roles: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\trole '%d': '%v'\n", k, *v)
		}
		return str
	}(tastix.Roles))
	str += fmt.Sprintf("goals: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tgoal '%d': '%s'\n", k, v)
		}
		return str
	}(tastix.Goals))
	str += fmt.Sprintf("sophistication: '%v'\n", tastix.FirstSeen)
	str += fmt.Sprintf("resource_level: '%v'\n", tastix.LastSeen)
	str += fmt.Sprintf("primary_motivation: '%v'\n", tastix.LastSeen)
	str += fmt.Sprintf("secondary_motivations: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tsecondary_motivation '%d': '%v'\n", k, *v)
		}
		return str
	}(tastix.SecondaryMotivations))
	str += fmt.Sprintf("personal_motivations: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tpersonal_motivation '%d': '%v'\n", k, *v)
		}
		return str
	}(tastix.PersonalMotivations))

	return str
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
func (tstix ToolDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(tool--)[0-9a-f|-]+$`).MatchString(tstix.ID)) {
		return false
	}

	if !tstix.checkingTypeCommonFields() {
		return false
	}

	if tstix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (tstix ToolDomainObjectsSTIX) SanitizeStruct() ToolDomainObjectsSTIX {
	tstix.CommonPropertiesDomainObjectSTIX = tstix.sanitizeStruct()

	tstix.Name = commonlibs.StringSanitize(tstix.Name)
	tstix.Description = commonlibs.StringSanitize(tstix.Description)

	if len(tstix.ToolTypes) > 0 {
		t := make([]*OpenVocabTypeSTIX, 0, len(tstix.ToolTypes))
		for _, v := range tstix.ToolTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			t = append(t, &tmp)
		}

		tstix.ToolTypes = t
	}

	if len(tstix.Aliases) > 0 {
		mTmp := make([]string, 0, len(tstix.Aliases))
		for _, v := range tstix.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tstix.Aliases = mTmp
	}

	tstix.KillChainPhases = tstix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()
	tstix.ToolVersion = commonlibs.StringSanitize(tstix.ToolVersion)

	return tstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (tstix ToolDomainObjectsSTIX) ToStringBeautiful() string {
	str := tstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += tstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()

	str += fmt.Sprintf("name: '%s'\n", tstix.Name)
	str += fmt.Sprintf("description: '%s'\n", tstix.Description)
	str += fmt.Sprintf("tool_types: \n%v", func(l []*OpenVocabTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\ttool_type '%d': '%v'\n", k, *v)
		}
		return str
	}(tstix.ToolTypes))
	str += fmt.Sprintf("aliases: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\taliase '%d': '%s'\n", k, v)
		}
		return str
	}(tstix.Aliases))
	str += fmt.Sprintf("kill_chain_phases: \n%v", func(l []*KillChainPhasesTypeElementSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tkill_chain_phase '%d': '%v'\n", k, *v)
		}
		return str
	}(tstix.KillChainPhases))
	str += fmt.Sprintf("tool_version: '%s'\n", tstix.ToolVersion)

	return str
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
func (vstix VulnerabilityDomainObjectsSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(vulnerability--)[0-9a-f|-]+$`).MatchString(vstix.ID)) {
		return false
	}

	if !vstix.checkingTypeCommonFields() {
		return false
	}

	if vstix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (vstix VulnerabilityDomainObjectsSTIX) SanitizeStruct() VulnerabilityDomainObjectsSTIX {
	vstix.CommonPropertiesDomainObjectSTIX = vstix.sanitizeStruct()

	vstix.Name = commonlibs.StringSanitize(vstix.Name)
	vstix.Description = commonlibs.StringSanitize(vstix.Description)

	return vstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (vstix VulnerabilityDomainObjectsSTIX) ToStringBeautiful() string {
	str := vstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += vstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", vstix.Name)
	str += fmt.Sprintf("description: '%s'\n", vstix.Description)

	return str
}
