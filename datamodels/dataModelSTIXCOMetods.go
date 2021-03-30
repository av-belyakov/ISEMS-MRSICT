package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
)

/*********************************************************************************/
/********** 			Cyber-observable Objects STIX (МЕТОДЫ)			**********/
/*********************************************************************************/

/*
//OptionalCommonPropertiesCyberObservableObjectSTIX содержит опциональные общие свойства для Cyber-observable Objects STIX
// SpecVersion - версия STIX спецификации.
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// Extensions - может содержать дополнительную информацию, относящуюся к объекту
type OptionalCommonPropertiesCyberObservableObjectSTIX struct {
	SpecVersion       string                         `json:"spec_version" bson:"spec_version"`
	ObjectMarkingRefs []*IdentifierTypeSTIX          `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings  GranularMarkingsTypeSTIX       `json:"granular_markings" bson:"granular_markings"`
	Defanged          bool                           `json:"defanged" bson:"defanged"`
	Extensions        map[string]*DictionaryTypeSTIX `json:"extensions" bson:"extensions"`
}
*/

func (ocpcstix *OptionalCommonPropertiesCyberObservableObjectSTIX) checkingTypeCommonFields() bool {
	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(ocpcstix.SpecVersion)) {
		return false
	}

	return true
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) sanitizeStruct() OptionalCommonPropertiesCyberObservableObjectSTIX {

	return ocpcstix
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) ToStringBeautiful() string {
	var str string
	str += fmt.Sprintf("spec_version: '%s'\n", ocpcstix.SpecVersion)

	/*
		str += fmt.Sprintf("created: '%v'\n", cp.Created)
		str += fmt.Sprintf("modified: '%v'\n", cp.Modified)
		str += fmt.Sprintf("created_by_ref: '%s'\n", cp.CreatedByRef)
		str += fmt.Sprintf("revoked: '%v'\n", cp.Revoked)
	*/

	return str
}

/* --- ArtifactCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (astix ArtifactCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &astix); err != nil {
		return nil, err
	}

	return astix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (astix ArtifactCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(astix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе ArtifactCyberObservableObjectSTIX
func (astix ArtifactCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(artifact--)[0-9a-f|-]+$`).MatchString(astix.ID)) {
		return false
	}

	if !astix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (astix ArtifactCyberObservableObjectSTIX) SanitizeStruct() ArtifactCyberObservableObjectSTIX {
	astix.OptionalCommonPropertiesCyberObservableObjectSTIX = astix.sanitizeStruct()

	/*
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
	*/

	return astix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (astix ArtifactCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := astix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += astix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- AutonomousSystemCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (asstix AutonomousSystemCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &asstix); err != nil {
		return nil, err
	}

	return asstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (asstix AutonomousSystemCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(asstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AutonomousSystemCyberObservableObjectSTIX
func (asstix AutonomousSystemCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(autonomous-system--)[0-9a-f|-]+$`).MatchString(asstix.ID)) {
		return false
	}

	if !asstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (asstix AutonomousSystemCyberObservableObjectSTIX) SanitizeStruct() AutonomousSystemCyberObservableObjectSTIX {
	asstix.OptionalCommonPropertiesCyberObservableObjectSTIX = asstix.sanitizeStruct()

	/*
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
	*/

	return asstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (asstix AutonomousSystemCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := asstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += asstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- DirectoryCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (dstix DirectoryCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &dstix); err != nil {
		return nil, err
	}

	return dstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (dstix DirectoryCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(dstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе DirectoryCyberObservableObjectSTIX
func (dstix DirectoryCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(directory--)[0-9a-f|-]+$`).MatchString(dstix.ID)) {
		return false
	}

	if !dstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (dstix DirectoryCyberObservableObjectSTIX) SanitizeStruct() DirectoryCyberObservableObjectSTIX {
	dstix.OptionalCommonPropertiesCyberObservableObjectSTIX = dstix.sanitizeStruct()

	/*
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
	*/

	return dstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dstix DirectoryCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- DomainNameCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (dnstix DomainNameCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &dnstix); err != nil {
		return nil, err
	}

	return dnstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (dnstix DomainNameCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(dnstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе DomainNameCyberObservableObjectSTIX
func (dnstix DomainNameCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(domain-name--)[0-9a-f|-]+$`).MatchString(dnstix.ID)) {
		return false
	}

	if !dnstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (dnstix DomainNameCyberObservableObjectSTIX) SanitizeStruct() DomainNameCyberObservableObjectSTIX {
	dnstix.OptionalCommonPropertiesCyberObservableObjectSTIX = dnstix.sanitizeStruct()

	/*
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
	*/

	return dnstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dnstix DomainNameCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dnstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dnstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- EmailAddressCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (eastix EmailAddressCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &eastix); err != nil {
		return nil, err
	}

	return eastix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (eastix EmailAddressCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(eastix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе EmailAddressCyberObservableObjectSTIX
func (eastix EmailAddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(email-addr--)[0-9a-f|-]+$`).MatchString(eastix.ID)) {
		return false
	}

	if !eastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (eastix EmailAddressCyberObservableObjectSTIX) SanitizeStruct() EmailAddressCyberObservableObjectSTIX {
	eastix.OptionalCommonPropertiesCyberObservableObjectSTIX = eastix.sanitizeStruct()

	/*
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
	*/

	return eastix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (eastix EmailAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := eastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += eastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- EmailAddressCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (emstix EmailMessageCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &emstix); err != nil {
		return nil, err
	}

	return emstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (emstix EmailMessageCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(emstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе EmailMessageCyberObservableObjectSTIX
func (emstix EmailMessageCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(email-message--)[0-9a-f|-]+$`).MatchString(emstix.ID)) {
		return false
	}

	if !emstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (emstix EmailMessageCyberObservableObjectSTIX) SanitizeStruct() EmailMessageCyberObservableObjectSTIX {
	emstix.OptionalCommonPropertiesCyberObservableObjectSTIX = emstix.sanitizeStruct()

	/*
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
	*/

	return emstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (emstix EmailMessageCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := emstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += emstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- FileCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (fstix FileCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonFileCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	fstix = FileCyberObservableObjectSTIX{
		CommonFileCyberObservableObjectSTIX: commonObject,
	}

	if len(commonObject.Extensions) == 0 {
		return fstix, nil
	}

	ext := map[string]*interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = &e
	}

	fstix.Extensions = ext

	return fstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (fstix FileCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(fstix)

	return &result, err
}

//GetFileCyberObservableObjectSTIX возвращает объект типа FileCyberObservableObjectSTIX
func (fstix *FileCyberObservableObjectSTIX) GetFileCyberObservableObjectSTIX() *FileCyberObservableObjectSTIX {
	return fstix
}

//CheckingTypeFields является валидатором параметров содержащихся в типе FileCyberObservableObjectSTIX
func (fstix FileCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(file--)[0-9a-f|-]+$`).MatchString(fstix.ID)) {
		return false
	}

	if !fstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (fstix FileCyberObservableObjectSTIX) SanitizeStruct() FileCyberObservableObjectSTIX {
	fstix.OptionalCommonPropertiesCyberObservableObjectSTIX = fstix.sanitizeStruct()

	/*
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
	*/

	return fstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (fstix FileCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := fstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += fstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- IPv4AddressCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (ip4stix IPv4AddressCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &ip4stix); err != nil {
		return nil, err
	}

	return ip4stix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (ip4stix IPv4AddressCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ip4stix)

	return &result, err
}

//GetIPv4AddressCyberObservableObjectSTIX выполняет объект типа IPv4AddressCyberObservableObjectSTIX
func (fstix *IPv4AddressCyberObservableObjectSTIX) GetIPv4AddressCyberObservableObjectSTIX() *IPv4AddressCyberObservableObjectSTIX {
	return fstix
}

//CheckingTypeFields является валидатором параметров содержащихся в типе IPv4AddressCyberObservableObjectSTIX
func (ip4stix IPv4AddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(ipv4-addr--)[0-9a-f|-]+$`).MatchString(ip4stix.ID)) {
		return false
	}

	if !ip4stix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ip4stix IPv4AddressCyberObservableObjectSTIX) SanitizeStruct() IPv4AddressCyberObservableObjectSTIX {
	ip4stix.OptionalCommonPropertiesCyberObservableObjectSTIX = ip4stix.sanitizeStruct()

	/*
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
	*/

	return ip4stix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ip4stix IPv4AddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ip4stix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ip4stix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- IPv6AddressCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (ip6stix IPv6AddressCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &ip6stix); err != nil {
		return nil, err
	}

	return ip6stix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (ip6stix IPv6AddressCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ip6stix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе IPv6AddressCyberObservableObjectSTIX
func (ip6stix IPv6AddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(ipv6-addr--)[0-9a-f|-]+$`).MatchString(ip6stix.ID)) {
		return false
	}

	if !ip6stix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ip6stix IPv6AddressCyberObservableObjectSTIX) SanitizeStruct() IPv6AddressCyberObservableObjectSTIX {
	ip6stix.OptionalCommonPropertiesCyberObservableObjectSTIX = ip6stix.sanitizeStruct()

	/*
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
	*/

	return ip6stix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ip6stix IPv6AddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ip6stix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ip6stix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- MACAddressCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (macstix MACAddressCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &macstix); err != nil {
		return nil, err
	}

	return macstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (macstix MACAddressCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(macstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе MACAddressCyberObservableObjectSTIX
func (macstix MACAddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(mac-addr--)[0-9a-f|-]+$`).MatchString(macstix.ID)) {
		return false
	}

	if !macstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (macstix MACAddressCyberObservableObjectSTIX) SanitizeStruct() MACAddressCyberObservableObjectSTIX {
	macstix.OptionalCommonPropertiesCyberObservableObjectSTIX = macstix.sanitizeStruct()

	/*
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
	*/

	return macstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (macstix MACAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := macstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += macstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- MutexCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (mstix MutexCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &mstix); err != nil {
		return nil, err
	}

	return mstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (mstix MutexCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(mstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе MutexCyberObservableObjectSTIX
func (mstix MutexCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(mutex--)[0-9a-f|-]+$`).MatchString(mstix.ID)) {
		return false
	}

	if !mstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mstix MutexCyberObservableObjectSTIX) SanitizeStruct() MutexCyberObservableObjectSTIX {
	mstix.OptionalCommonPropertiesCyberObservableObjectSTIX = mstix.sanitizeStruct()

	/*
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
	*/

	return mstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mstix MutexCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := mstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += mstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- NetworkTrafficCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (ntstix NetworkTrafficCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonNetworkTrafficCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	ntstix = NetworkTrafficCyberObservableObjectSTIX{
		CommonNetworkTrafficCyberObservableObjectSTIX: commonObject,
	}

	if len(commonObject.Extensions) == 0 {
		return ntstix, nil
	}

	ext := map[string]*interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = &e
	}

	ntstix.Extensions = ext

	return ntstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (ntstix NetworkTrafficCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ntstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе NetworkTrafficCyberObservableObjectSTIX
func (ntstix NetworkTrafficCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(network-traffic--)[0-9a-f|-]+$`).MatchString(ntstix.ID)) {
		return false
	}

	if !ntstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ntstix NetworkTrafficCyberObservableObjectSTIX) SanitizeStruct() NetworkTrafficCyberObservableObjectSTIX {
	ntstix.OptionalCommonPropertiesCyberObservableObjectSTIX = ntstix.sanitizeStruct()

	/*
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
	*/

	return ntstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ntstix NetworkTrafficCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ntstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ntstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- ProcessCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (pstix ProcessCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonProcessCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	pstix = ProcessCyberObservableObjectSTIX{
		CommonProcessCyberObservableObjectSTIX: commonObject,
	}

	if len(commonObject.Extensions) == 0 {
		return pstix, nil
	}

	ext := map[string]*interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = &e
	}
	pstix.Extensions = ext

	return pstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (pstix ProcessCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(pstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе ProcessCyberObservableObjectSTIX
func (pstix ProcessCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(process--)[0-9a-f|-]+$`).MatchString(pstix.ID)) {
		return false
	}

	if !pstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (pstix ProcessCyberObservableObjectSTIX) SanitizeStruct() ProcessCyberObservableObjectSTIX {
	pstix.OptionalCommonPropertiesCyberObservableObjectSTIX = pstix.sanitizeStruct()

	/*
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
	*/

	return pstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (pstix ProcessCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := pstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += pstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- SoftwareCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (sstix SoftwareCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &sstix); err != nil {
		return nil, err
	}

	return sstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (sstix SoftwareCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(sstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе SoftwareCyberObservableObjectSTIX
func (sstix SoftwareCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(software--)[0-9a-f|-]+$`).MatchString(sstix.ID)) {
		return false
	}

	if !sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (sstix SoftwareCyberObservableObjectSTIX) SanitizeStruct() SoftwareCyberObservableObjectSTIX {
	sstix.OptionalCommonPropertiesCyberObservableObjectSTIX = sstix.sanitizeStruct()

	/*
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
	*/

	return sstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (sstix SoftwareCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := sstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += sstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- URLCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (urlstix URLCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &urlstix); err != nil {
		return nil, err
	}

	return urlstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (urlstix URLCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(urlstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе URLCyberObservableObjectSTIX
func (urlstix URLCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(url--)[0-9a-f|-]+$`).MatchString(urlstix.ID)) {
		return false
	}

	if !urlstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (urlstix URLCyberObservableObjectSTIX) SanitizeStruct() URLCyberObservableObjectSTIX {
	urlstix.OptionalCommonPropertiesCyberObservableObjectSTIX = urlstix.sanitizeStruct()

	/*
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
	*/

	return urlstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (urlstix URLCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := urlstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += urlstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- UserAccountCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (uastix UserAccountCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &uastix); err != nil {
		return nil, err
	}

	return uastix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (uastix UserAccountCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(uastix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе UserAccountCyberObservableObjectSTIX
func (uastix UserAccountCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(user-account--)[0-9a-f|-]+$`).MatchString(uastix.ID)) {
		return false
	}

	if !uastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (uastix UserAccountCyberObservableObjectSTIX) SanitizeStruct() UserAccountCyberObservableObjectSTIX {
	uastix.OptionalCommonPropertiesCyberObservableObjectSTIX = uastix.sanitizeStruct()

	/*
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
	*/

	return uastix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (uastix UserAccountCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := uastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += uastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- WindowsRegistryKeyCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &wrkstix); err != nil {
		return nil, err
	}

	return wrkstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(wrkstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе WindowsRegistryKeyCyberObservableObjectSTIX
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(windows-registry-key--)[0-9a-f|-]+$`).MatchString(wrkstix.ID)) {
		return false
	}

	if !wrkstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) SanitizeStruct() WindowsRegistryKeyCyberObservableObjectSTIX {
	wrkstix.OptionalCommonPropertiesCyberObservableObjectSTIX = wrkstix.sanitizeStruct()

	/*
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
	*/

	return wrkstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := wrkstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += wrkstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

/* --- X509CertificateCyberObservableObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &x509sstix); err != nil {
		return nil, err
	}

	return x509sstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (x509sstix X509CertificateCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(x509sstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе X509CertificateCyberObservableObjectSTIX
func (x509sstix X509CertificateCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(x509-certificate--)[0-9a-f|-]+$`).MatchString(x509sstix.ID)) {
		return false
	}

	if !x509sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (x509sstix X509CertificateCyberObservableObjectSTIX) SanitizeStruct() X509CertificateCyberObservableObjectSTIX {
	x509sstix.OptionalCommonPropertiesCyberObservableObjectSTIX = x509sstix.sanitizeStruct()

	/*
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
	*/

	return x509sstix
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (x509sstix X509CertificateCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := x509sstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += x509sstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()

	/*
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
	*/

	return str
}

//decodingExtensionsSTIX декодирует следующие типы STIX расширений:
// - "archive-ext"
// - "ntfs-ext"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "windows-process-ext"
// - "windows-service-ext"
// - "unix-account-ext"
func decodingExtensionsSTIX(extType string, rawMsg *json.RawMessage) (interface{}, error) {
	var err error
	switch extType {
	case "archive-ext":
		var archiveExt ArchiveFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &archiveExt)

		return archiveExt, err
	case "ntfs-ext":
		var ntfsExt NTFSFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &ntfsExt)

		return ntfsExt, err
	case "pdf-ext":
		var pdfExt PDFFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &pdfExt)

		return pdfExt, err
	case "raster-image-ext":
		var rasterImageExt RasterImageFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &rasterImageExt)

		return rasterImageExt, err
	case "windows-pebinary-ext":
		var windowsPebinaryExt WindowsPEBinaryFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &windowsPebinaryExt)

		return windowsPebinaryExt, err
	case "http-request-ext":
		var httpRequestExt HTTPRequestExtensionSTIX
		err = json.Unmarshal(*rawMsg, &httpRequestExt)

		return httpRequestExt, err
	case "icmp-ext":
		var icmpExt ICMPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &icmpExt)

		return icmpExt, err
	case "socket-ext":
		var socketExt NetworkSocketExtensionSTIX
		err := json.Unmarshal(*rawMsg, &socketExt)

		return socketExt, err
	case "tcp-ext":
		var tcpExt TCPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &tcpExt)

		return tcpExt, err
	case "windows-process-ext":
		var windowsProcessExt WindowsProcessExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsProcessExt)

		return windowsProcessExt, err
	case "windows-service-ext":
		var windowsServiceExt WindowsServiceExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsServiceExt)

		return windowsServiceExt, err
	case "unix-account-ext":
		var unixAccountExt UNIXAccountExtensionSTIX
		err := json.Unmarshal(*rawMsg, &unixAccountExt)

		return unixAccountExt, err
	default:
		return struct{}{}, nil
	}
}
