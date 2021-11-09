package datamodels

import (
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"time"

	"ISEMS-MRSICT/commonlibs"

	"github.com/asaskevich/govalidator"
)

/*********************************************************************************/
/********** 			Cyber-observable Objects STIX (МЕТОДЫ)			**********/
/*********************************************************************************/

func (ocpcstix *OptionalCommonPropertiesCyberObservableObjectSTIX) validateStructCommonFields() bool {
	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(ocpcstix.SpecVersion)) {
		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(ocpcstix.ObjectMarkingRefs) > 0 {
		for _, value := range ocpcstix.ObjectMarkingRefs {
			if !value.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(ocpcstix.GranularMarkings) > 0 {
		for _, value := range ocpcstix.GranularMarkings {
			//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
			if !value.CheckGranularMarkingsTypeSTIX() {
				return false
			}
		}
	}

	return true
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) sanitizeStruct() OptionalCommonPropertiesCyberObservableObjectSTIX {
	//обработка содержимого списка поля Extensions
	if len(ocpcstix.Extensions) > 0 {
		ext := make(map[string]DictionaryTypeSTIX, len(ocpcstix.Extensions))
		for k, v := range ocpcstix.Extensions {
			switch v := v.dictionary.(type) {
			case string:
				ext[k] = DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				ext[k] = DictionaryTypeSTIX{v}
			}
		}
		ocpcstix.Extensions = ext
	}

	return ocpcstix
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) ToStringBeautiful() string {
	var str string
	str += fmt.Sprintf("spec_version: '%s'\n", ocpcstix.SpecVersion)
	str += fmt.Sprintf("object_marking_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_marking_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ocpcstix.ObjectMarkingRefs))
	str += fmt.Sprintf("granular_markings: \n%v", func(l []GranularMarkingsTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tgranular_markings number %d.", k)
			str += fmt.Sprintf("\tlang: '%s'\n", v.Lang)
			str += fmt.Sprintf("\tmarking_ref: '%v'\n", v.MarkingRef)
			str += fmt.Sprintf("\tselectors: \n%v", func(l []string) string {
				var str string
				for k, v := range l {
					str += fmt.Sprintf("\t\tselector '%d': '%s'\n", k, v)
				}
				return str
			}(v.Selectors))
		}
		return str
	}(ocpcstix.GranularMarkings))
	str += fmt.Sprintf("defanged: '%v'\n", ocpcstix.Defanged)
	str += fmt.Sprintf("extensions: \n%v", func(l map[string]DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, v)
		}
		return str
	}(ocpcstix.Extensions))

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

//ValidateStruct является валидатором параметров содержащихся в типе ArtifactCyberObservableObjectSTIX
func (astix ArtifactCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(artifact--)[0-9a-f|-]+$`).MatchString(astix.ID)) {
		return false
	}

	if !astix.validateStructCommonFields() {
		return false
	}

	if astix.PayloadBin != "" {
		if !govalidator.IsBase64(astix.PayloadBin) {
			return false
		}
	}

	if astix.URL != "" {
		if !govalidator.IsURL(astix.URL) {
			return false
		}
	}

	if !astix.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (astix ArtifactCyberObservableObjectSTIX) SanitizeStruct() ArtifactCyberObservableObjectSTIX {
	astix.OptionalCommonPropertiesCyberObservableObjectSTIX = astix.sanitizeStruct()

	astix.MimeType = commonlibs.StringSanitize(astix.MimeType)
	astix.EncryptionAlgorithm = EnumTypeSTIX(commonlibs.StringSanitize(string(astix.EncryptionAlgorithm)))
	astix.DecryptionKey = commonlibs.StringSanitize(astix.DecryptionKey)

	return astix
}

//GetID возвращает ID STIX объекта
func (astix ArtifactCyberObservableObjectSTIX) GetID() string {
	return astix.ID
}

//GetType возвращает Type STIX объекта
func (astix ArtifactCyberObservableObjectSTIX) GetType() string {
	return astix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа ArtifactCyberObservableObjectSTIX
func (astix ArtifactCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             astix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (astix ArtifactCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := astix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += astix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("mime_type: '%s'\n", astix.MimeType)
	str += fmt.Sprintf("payload_bin: '%s'\n", astix.PayloadBin)
	str += fmt.Sprintf("url: '%s'\n", astix.URL)
	str += fmt.Sprintf("hashes: '%v'\n", astix.Hashes)
	str += fmt.Sprintf("encryption_algorithm: '%v'\n", astix.EncryptionAlgorithm)
	str += fmt.Sprintf("decryption_key: '%s'\n", astix.DecryptionKey)

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

//ValidateStruct является валидатором параметров содержащихся в типе AutonomousSystemCyberObservableObjectSTIX
func (asstix AutonomousSystemCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(autonomous-system--)[0-9a-f|-]+$`).MatchString(asstix.ID)) {
		return false
	}

	if !asstix.validateStructCommonFields() {
		return false
	}

	if asstix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (asstix AutonomousSystemCyberObservableObjectSTIX) SanitizeStruct() AutonomousSystemCyberObservableObjectSTIX {
	asstix.OptionalCommonPropertiesCyberObservableObjectSTIX = asstix.sanitizeStruct()

	asstix.Name = commonlibs.StringSanitize(asstix.Name)
	asstix.RIR = commonlibs.StringSanitize(asstix.RIR)

	return asstix
}

//GetID возвращает ID STIX объекта
func (asstix AutonomousSystemCyberObservableObjectSTIX) GetID() string {
	return asstix.ID
}

//GetType возвращает Type STIX объекта
func (asstix AutonomousSystemCyberObservableObjectSTIX) GetType() string {
	return asstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа AutonomousSystemCyberObservableObjectSTIX
func (asstix AutonomousSystemCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             asstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (asstix AutonomousSystemCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := asstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += asstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("number: '%d'\n", asstix.Number)
	str += fmt.Sprintf("name: '%s'\n", asstix.Name)
	str += fmt.Sprintf("rir: '%s'\n", asstix.RIR)

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

//ValidateStruct является валидатором параметров содержащихся в типе DirectoryCyberObservableObjectSTIX
func (dstix DirectoryCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(directory--)[0-9a-f|-]+$`).MatchString(dstix.ID)) {
		return false
	}

	if !dstix.validateStructCommonFields() {
		return false
	}

	if dstix.Path == "" {
		return false
	}

	isUnixPath := govalidator.IsUnixFilePath(dstix.Path)
	isWinPath := govalidator.IsWinFilePath(dstix.Path)
	if !isUnixPath && !isWinPath {
		return false
	}

	if len(dstix.ContainsRefs) > 0 {
		for _, v := range dstix.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (dstix DirectoryCyberObservableObjectSTIX) SanitizeStruct() DirectoryCyberObservableObjectSTIX {
	dstix.OptionalCommonPropertiesCyberObservableObjectSTIX = dstix.sanitizeStruct()

	dstix.PathEnc = commonlibs.StringSanitize(dstix.PathEnc)

	return dstix
}

//GetID возвращает ID STIX объекта
func (dstix DirectoryCyberObservableObjectSTIX) GetID() string {
	return dstix.ID
}

//GetType возвращает Type STIX объекта
func (dstix DirectoryCyberObservableObjectSTIX) GetType() string {
	return dstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа DirectoryCyberObservableObjectSTIX
func (dstix DirectoryCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             dstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dstix DirectoryCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("path: '%s'\n", dstix.Path)
	str += fmt.Sprintf("path_enc: '%s'\n", dstix.PathEnc)
	str += fmt.Sprintf("ctime: '%v'\n", dstix.Ctime)
	str += fmt.Sprintf("mtime: '%v'\n", dstix.Mtime)
	str += fmt.Sprintf("atime: '%s'\n", dstix.Atime)
	str += fmt.Sprintf("contains_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcontains_ref '%d': '%v'\n", k, v)
		}
		return str
	}(dstix.ContainsRefs))

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

//ValidateStruct является валидатором параметров содержащихся в типе DomainNameCyberObservableObjectSTIX
func (dnstix DomainNameCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(domain-name--)[0-9a-f|-]+$`).MatchString(dnstix.ID)) {
		return false
	}

	if !dnstix.validateStructCommonFields() {
		return false
	}

	if !govalidator.IsDNSName(dnstix.Value) {
		return false
	}

	if len(dnstix.ResolvesToRefs) > 0 {
		for _, v := range dnstix.ResolvesToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (dnstix DomainNameCyberObservableObjectSTIX) SanitizeStruct() DomainNameCyberObservableObjectSTIX {
	dnstix.OptionalCommonPropertiesCyberObservableObjectSTIX = dnstix.sanitizeStruct()

	return dnstix
}

//GetID возвращает ID STIX объекта
func (dnstix DomainNameCyberObservableObjectSTIX) GetID() string {
	return dnstix.ID
}

//GetType возвращает Type STIX объекта
func (dnstix DomainNameCyberObservableObjectSTIX) GetType() string {
	return dnstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа DomainNameCyberObservableObjectSTIX
func (dnstix DomainNameCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             dnstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dnstix DomainNameCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dnstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dnstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", dnstix.Value)
	str += fmt.Sprintf("resolves_to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tresolves_to_ref '%d': '%v'\n", k, v)
		}
		return str
	}(dnstix.ResolvesToRefs))

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

//ValidateStruct является валидатором параметров содержащихся в типе EmailAddressCyberObservableObjectSTIX
func (eastix EmailAddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(email-addr--)[0-9a-f|-]+$`).MatchString(eastix.ID)) {
		return false
	}

	if !eastix.validateStructCommonFields() {
		return false
	}

	if !govalidator.IsEmail(eastix.Value) {
		return false
	}

	if !eastix.BelongsToRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (eastix EmailAddressCyberObservableObjectSTIX) SanitizeStruct() EmailAddressCyberObservableObjectSTIX {
	eastix.OptionalCommonPropertiesCyberObservableObjectSTIX = eastix.sanitizeStruct()

	eastix.DisplayName = commonlibs.StringSanitize(eastix.DisplayName)

	return eastix
}

//GetID возвращает ID STIX объекта
func (eastix EmailAddressCyberObservableObjectSTIX) GetID() string {
	return eastix.ID
}

//GetType возвращает Type STIX объекта
func (eastix EmailAddressCyberObservableObjectSTIX) GetType() string {
	return eastix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа EmailAddressCyberObservableObjectSTIX
func (eastix EmailAddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             eastix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (eastix EmailAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := eastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += eastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", eastix.Value)
	str += fmt.Sprintf("display_name: '%s'\n", eastix.DisplayName)
	str += fmt.Sprintf("belongs_to_ref: '%v'\n", eastix.BelongsToRef)

	return str
}

/* --- EmailMessageCyberObservableObjectSTIX --- */

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

//ValidateStruct является валидатором параметров содержащихся в типе EmailMessageCyberObservableObjectSTIX
func (emstix EmailMessageCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(email-message--)[0-9a-f|-]+$`).MatchString(emstix.ID)) {
		return false
	}

	if !emstix.validateStructCommonFields() {
		return false
	}

	if !emstix.FromRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !emstix.SenderRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(emstix.ToRefs) > 0 {
		for _, v := range emstix.ToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(emstix.CcRefs) > 0 {
		for _, v := range emstix.CcRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(emstix.BccRefs) > 0 {
		for _, v := range emstix.BccRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(emstix.BodyMultipart) > 0 {
		for _, v := range emstix.BodyMultipart {
			if !v.BodyRawRef.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !emstix.RawEmailRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (emstix EmailMessageCyberObservableObjectSTIX) SanitizeStruct() EmailMessageCyberObservableObjectSTIX {
	emstix.OptionalCommonPropertiesCyberObservableObjectSTIX = emstix.sanitizeStruct()

	emstix.ContentType = commonlibs.StringSanitize(emstix.ContentType)
	emstix.MessageID = commonlibs.StringSanitize(emstix.MessageID)
	emstix.Subject = commonlibs.StringSanitize(emstix.Subject)

	if len(emstix.ReceivedLines) > 0 {
		tmp := make([]string, 0, len(emstix.ReceivedLines))
		for _, v := range emstix.ReceivedLines {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}
		emstix.ReceivedLines = tmp
	}

	if len(emstix.AdditionalHeaderFields) > 0 {
		tmp := make(map[string]DictionaryTypeSTIX, len(emstix.AdditionalHeaderFields))
		for k, v := range emstix.AdditionalHeaderFields {
			switch v := v.dictionary.(type) {
			case string:
				tmp[k] = DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = DictionaryTypeSTIX{v}
			}
		}
		emstix.AdditionalHeaderFields = tmp
	}

	emstix.Body = commonlibs.StringSanitize(emstix.Body)

	if len(emstix.BodyMultipart) > 0 {
		tmp := make([]EmailMIMEPartTypeSTIX, 0, len(emstix.BodyMultipart))
		for _, v := range emstix.BodyMultipart {
			tmp = append(tmp, EmailMIMEPartTypeSTIX{
				Body:               commonlibs.StringSanitize(v.Body),
				BodyRawRef:         v.BodyRawRef,
				ContentType:        commonlibs.StringSanitize(v.ContentType),
				ContentDisposition: commonlibs.StringSanitize(v.ContentDisposition),
			})
		}
		emstix.BodyMultipart = tmp
	}

	return emstix
}

//GetID возвращает ID STIX объекта
func (emstix EmailMessageCyberObservableObjectSTIX) GetID() string {
	return emstix.ID
}

//GetType возвращает Type STIX объекта
func (emstix EmailMessageCyberObservableObjectSTIX) GetType() string {
	return emstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа EmailMessageCyberObservableObjectSTIX
func (emstix EmailMessageCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             emstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (emstix EmailMessageCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := emstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += emstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("is_multipart: '%v'\n", emstix.IsMultipart)
	str += fmt.Sprintf("date: '%v'\n", emstix.Date)
	str += fmt.Sprintf("content_type: '%s'\n", emstix.ContentType)
	str += fmt.Sprintf("from_ref: '%v'\n", emstix.FromRef)
	str += fmt.Sprintf("sender_ref: '%v'\n", emstix.SenderRef)
	str += fmt.Sprintf("to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tto_ref '%d': '%v'\n", k, v)
		}
		return str
	}(emstix.ToRefs))
	str += fmt.Sprintf("cc_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcc_ref '%d': '%v'\n", k, v)
		}
		return str
	}(emstix.CcRefs))
	str += fmt.Sprintf("bcc_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tbcc_ref '%d': '%v'\n", k, v)
		}
		return str
	}(emstix.BccRefs))
	str += fmt.Sprintf("message_id: '%v'\n", emstix.MessageID)
	str += fmt.Sprintf("subject: '%v'\n", emstix.Subject)
	str += fmt.Sprintf("received_lines: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\treceived_line '%d': '%s'\n", k, v)
		}
		return str
	}(emstix.ReceivedLines))
	str += fmt.Sprintf("additional_header_fields: \n%v", func(l map[string]DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, v)
		}
		return str
	}(emstix.AdditionalHeaderFields))
	str += fmt.Sprintf("body: '%v'\n", emstix.Body)
	str += fmt.Sprintf("body_multipart: \n%v", func(l []EmailMIMEPartTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tbody_multipart '%d':\n", k)
			str += fmt.Sprintf("\t\tbody: '%s'\n", v.Body)
			str += fmt.Sprintf("\t\tbody_raw_ref: '%s'\n", v.BodyRawRef)
			str += fmt.Sprintf("\t\tcontent_type: '%s'\n", v.ContentType)
			str += fmt.Sprintf("\t\tcontent_disposition: '%s'\n", v.ContentDisposition)
		}
		return str
	}(emstix.BodyMultipart))
	str += fmt.Sprintf("raw_email_ref: '%v'\n", emstix.RawEmailRef)

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
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Hashes:             commonObject.Hashes,
		Size:               commonObject.Size,
		Name:               commonObject.Name,
		NameEnc:            commonObject.NameEnc,
		MagicNumberHex:     commonObject.MagicNumberHex,
		MimeType:           commonObject.MimeType,
		Ctime:              commonObject.Ctime,
		Mtime:              commonObject.Mtime,
		Atime:              commonObject.Atime,
		ParentDirectoryRef: commonObject.ParentDirectoryRef,
	}

	if len(commonObject.Extensions) == 0 {
		return fstix, nil
	}

	ext := map[string]interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = e
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

//ValidateStruct является валидатором параметров содержащихся в типе FileCyberObservableObjectSTIX
func (fstix FileCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(file--)[0-9a-f|-]+$`).MatchString(fstix.ID)) {
		return false
	}

	if !fstix.validateStructCommonFields() {
		return false
	}

	if !fstix.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	if !fstix.ParentDirectoryRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(fstix.ContainsRefs) > 0 {
		for _, v := range fstix.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !fstix.ContentRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(fstix.Extensions) > 0 {
		for _, v := range fstix.Extensions {
			if !checkingExtensionsSTIX(v) {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (fstix FileCyberObservableObjectSTIX) SanitizeStruct() FileCyberObservableObjectSTIX {
	fstix.OptionalCommonPropertiesCyberObservableObjectSTIX = fstix.sanitizeStruct()

	fstix.Name = commonlibs.StringSanitize(fstix.Name)
	fstix.NameEnc = commonlibs.StringSanitize(fstix.NameEnc)
	fstix.MagicNumberHex = commonlibs.StringSanitize(fstix.MagicNumberHex)
	fstix.MimeType = commonlibs.StringSanitize(fstix.MimeType)

	esize := len(fstix.Extensions)
	if esize == 0 {
		return fstix
	}

	tmp := make(map[string]interface{}, esize)
	for k, v := range fstix.Extensions {
		result := sanitizeExtensionsSTIX(v)
		tmp[k] = result
	}
	fstix.Extensions = tmp

	return fstix
}

//GetID возвращает ID STIX объекта
func (fstix FileCyberObservableObjectSTIX) GetID() string {
	return fstix.ID
}

//GetType возвращает Type STIX объекта
func (fstix FileCyberObservableObjectSTIX) GetType() string {
	return fstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа FileCyberObservableObjectSTIX
func (fstix FileCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             fstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (fstix FileCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := fstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += fstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("hashes: '%v'\n", fstix.Hashes)
	str += fmt.Sprintf("size: '%d'\n", fstix.Size)
	str += fmt.Sprintf("name: '%s'\n", fstix.Name)
	str += fmt.Sprintf("name_enc: '%s'\n", fstix.NameEnc)
	str += fmt.Sprintf("magic_number_hex: '%s'\n", fstix.MagicNumberHex)
	str += fmt.Sprintf("mime_type: '%s'\n", fstix.MimeType)
	str += fmt.Sprintf("ctime: '%v'\n", fstix.Ctime)
	str += fmt.Sprintf("mtime: '%v'\n", fstix.Mtime)
	str += fmt.Sprintf("atime: '%v'\n", fstix.Atime)
	str += fmt.Sprintf("parent_directory_ref: '%v'\n", fstix.ParentDirectoryRef)
	str += fmt.Sprintf("contains_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcontains_ref '%d': '%v'\n", k, v)
		}
		return str
	}(fstix.ContainsRefs))
	str += fmt.Sprintf("content_ref: '%v'\n", fstix.ContentRef)
	str += fmt.Sprintln("extensions:")
	for k, v := range fstix.Extensions {
		str += fmt.Sprintf("\t%s:\n%v\n", k, toStringBeautiful(v))
	}

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

//ValidateStruct является валидатором параметров содержащихся в типе IPv4AddressCyberObservableObjectSTIX
func (ip4stix IPv4AddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(ipv4-addr--)[0-9a-f|-]+$`).MatchString(ip4stix.ID)) {
		return false
	}

	if !ip4stix.validateStructCommonFields() {
		return false
	}

	if ip4stix.Value == "" {
		return false
	}

	isIPv4 := commonlibs.IsIPv4Address(ip4stix.Value)
	isNetworkIPv4 := commonlibs.IsComputerNetAddrIPv4Range(ip4stix.Value)
	if !isIPv4 && !isNetworkIPv4 {
		return false
	}

	if len(ip4stix.ResolvesToRefs) > 0 {
		for _, v := range ip4stix.ResolvesToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(ip4stix.BelongsToRefs) > 0 {
		for _, v := range ip4stix.BelongsToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ip4stix IPv4AddressCyberObservableObjectSTIX) SanitizeStruct() IPv4AddressCyberObservableObjectSTIX {
	ip4stix.OptionalCommonPropertiesCyberObservableObjectSTIX = ip4stix.sanitizeStruct()

	return ip4stix
}

//GetID возвращает ID STIX объекта
func (ip4stix IPv4AddressCyberObservableObjectSTIX) GetID() string {
	return ip4stix.ID
}

//GetType возвращает Type STIX объекта
func (ip4stix IPv4AddressCyberObservableObjectSTIX) GetType() string {
	return ip4stix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа IPv4AddressCyberObservableObjectSTIX
func (ipv4stix IPv4AddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ipv4stix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ip4stix IPv4AddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ip4stix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ip4stix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", ip4stix.Value)
	str += fmt.Sprintf("resolves_to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tresolves_to_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ip4stix.ResolvesToRefs))
	str += fmt.Sprintf("belongs_to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tbelongs_to_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ip4stix.BelongsToRefs))

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

//ValidateStruct является валидатором параметров содержащихся в типе IPv6AddressCyberObservableObjectSTIX
func (ip6stix IPv6AddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(ipv6-addr--)[0-9a-f|-]+$`).MatchString(ip6stix.ID)) {
		return false
	}

	if !ip6stix.validateStructCommonFields() {
		return false
	}

	if ip6stix.Value == "" {
		return false
	}

	if ipv6Addr, _, err := net.ParseCIDR(ip6stix.Value); err == nil {
		if !govalidator.IsIPv6(ipv6Addr.String()) {
			return false
		}
	} else {
		if !govalidator.IsIPv6(ip6stix.Value) {
			return false
		}
	}

	if len(ip6stix.ResolvesToRefs) > 0 {
		for _, v := range ip6stix.ResolvesToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(ip6stix.BelongsToRefs) > 0 {
		for _, v := range ip6stix.BelongsToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ip6stix IPv6AddressCyberObservableObjectSTIX) SanitizeStruct() IPv6AddressCyberObservableObjectSTIX {
	ip6stix.OptionalCommonPropertiesCyberObservableObjectSTIX = ip6stix.sanitizeStruct()

	return ip6stix
}

//GetID возвращает ID STIX объекта
func (ip6stix IPv6AddressCyberObservableObjectSTIX) GetID() string {
	return ip6stix.ID
}

//GetType возвращает Type STIX объекта
func (ip6stix IPv6AddressCyberObservableObjectSTIX) GetType() string {
	return ip6stix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа IPv6AddressCyberObservableObjectSTIX
func (ipv6stix IPv6AddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ipv6stix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ip6stix IPv6AddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ip6stix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ip6stix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", ip6stix.Value)
	str += fmt.Sprintf("resolves_to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tresolves_to_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ip6stix.ResolvesToRefs))
	str += fmt.Sprintf("belongs_to_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tbelongs_to_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ip6stix.BelongsToRefs))
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

//ValidateStruct является валидатором параметров содержащихся в типе MACAddressCyberObservableObjectSTIX
func (macstix MACAddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(mac-addr--)[0-9a-f|-]+$`).MatchString(macstix.ID)) {
		return false
	}

	if !macstix.validateStructCommonFields() {
		return false
	}

	if !govalidator.IsMAC(macstix.Value) {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (macstix MACAddressCyberObservableObjectSTIX) SanitizeStruct() MACAddressCyberObservableObjectSTIX {
	macstix.OptionalCommonPropertiesCyberObservableObjectSTIX = macstix.sanitizeStruct()

	return macstix
}

//GetID возвращает ID STIX объекта
func (macstix MACAddressCyberObservableObjectSTIX) GetID() string {
	return macstix.ID
}

//GetType возвращает Type STIX объекта
func (macstix MACAddressCyberObservableObjectSTIX) GetType() string {
	return macstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа MACAddressCyberObservableObjectSTIX
func (macstix MACAddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             macstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (macstix MACAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := macstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += macstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", macstix.Value)

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

//ValidateStruct является валидатором параметров содержащихся в типе MutexCyberObservableObjectSTIX
func (mstix MutexCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(mutex--)[0-9a-f|-]+$`).MatchString(mstix.ID)) {
		return false
	}

	if !mstix.validateStructCommonFields() {
		return false
	}

	if mstix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mstix MutexCyberObservableObjectSTIX) SanitizeStruct() MutexCyberObservableObjectSTIX {
	mstix.OptionalCommonPropertiesCyberObservableObjectSTIX = mstix.sanitizeStruct()
	mstix.Name = commonlibs.StringSanitize(mstix.Name)

	return mstix
}

//GetID возвращает ID STIX объекта
func (mstix MutexCyberObservableObjectSTIX) GetID() string {
	return mstix.ID
}

//GetType возвращает Type STIX объекта
func (mstix MutexCyberObservableObjectSTIX) GetType() string {
	return mstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа MutexCyberObservableObjectSTIX
func (mstix MutexCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             mstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mstix MutexCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := mstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += mstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", mstix.Name)

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
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Start:        commonObject.Start,
		End:          commonObject.End,
		IsActive:     commonObject.IsActive,
		SrcRef:       commonObject.SrcRef,
		DstRef:       commonObject.DstRef,
		SrcPort:      commonObject.SrcPort,
		DstPort:      commonObject.DstPort,
		Protocols:    commonObject.Protocols,
		SrcByteCount: commonObject.SrcByteCount,
		DstByteCount: commonObject.DstByteCount,
	}

	if len(commonObject.Extensions) == 0 {
		return ntstix, nil
	}

	ext := map[string]interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = e
	}

	ntstix.Extensions = ext

	return ntstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (ntstix NetworkTrafficCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ntstix)

	return &result, err
}

//ValidateStruct является валидатором параметров содержащихся в типе NetworkTrafficCyberObservableObjectSTIX
func (ntstix NetworkTrafficCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(network-traffic--)[0-9a-f|-]+$`).MatchString(ntstix.ID)) {
		return false
	}

	if !ntstix.validateStructCommonFields() {
		return false
	}

	if len(ntstix.Protocols) == 0 {
		return false
	}

	if !ntstix.SrcRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !ntstix.DstRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !ntstix.SrcPayloadRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !ntstix.DstPayloadRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(ntstix.EncapsulatesRefs) > 0 {
		for _, v := range ntstix.EncapsulatesRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !ntstix.EncapsulatedByRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(ntstix.Extensions) > 0 {
		for _, v := range ntstix.Extensions {
			if !checkingExtensionsSTIX(v) {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ntstix NetworkTrafficCyberObservableObjectSTIX) SanitizeStruct() NetworkTrafficCyberObservableObjectSTIX {
	ntstix.OptionalCommonPropertiesCyberObservableObjectSTIX = ntstix.sanitizeStruct()

	esize := len(ntstix.Extensions)
	if esize == 0 {
		return ntstix
	}

	sizeProtocols := len(ntstix.Protocols)
	if sizeProtocols == 0 {
		tmp := make([]string, 0, sizeProtocols)
		for _, v := range ntstix.Protocols {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}
		ntstix.Protocols = tmp
	}

	sizeIPFix := len(ntstix.IPFix)
	if sizeIPFix > 0 {
		tmp := make(map[string]DictionaryTypeSTIX, sizeIPFix)
		for k, v := range ntstix.IPFix {
			switch v := v.dictionary.(type) {
			case string:
				tmp[k] = DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = DictionaryTypeSTIX{v}
			}
		}
		ntstix.IPFix = tmp
	}

	tmp := make(map[string]interface{}, esize)
	for k, v := range ntstix.Extensions {
		result := sanitizeExtensionsSTIX(v)
		tmp[k] = result
	}
	ntstix.Extensions = tmp

	return ntstix
}

//GetID возвращает ID STIX объекта
func (ntstix NetworkTrafficCyberObservableObjectSTIX) GetID() string {
	return ntstix.ID
}

//GetType возвращает Type STIX объекта
func (ntstix NetworkTrafficCyberObservableObjectSTIX) GetType() string {
	return ntstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа NetworkTrafficCyberObservableObjectSTIX
func (ntstix NetworkTrafficCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ntstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ntstix NetworkTrafficCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := ntstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += ntstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("start: '%v'\n", ntstix.Start)
	str += fmt.Sprintf("end: '%v'\n", ntstix.End)
	str += fmt.Sprintf("is_active: '%v'\n", ntstix.IsActive)
	str += fmt.Sprintf("src_ref: '%v'\n", ntstix.SrcRef)
	str += fmt.Sprintf("dst_ref: '%v'\n", ntstix.DstRef)
	str += fmt.Sprintf("src_port: '%d'\n", ntstix.SrcPort)
	str += fmt.Sprintf("dst_port: '%d'\n", ntstix.DstPort)
	str += fmt.Sprintf("protocols: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tprotocol '%d': '%s'\n", k, v)
		}
		return str
	}(ntstix.Protocols))
	str += fmt.Sprintf("src_byte_count: '%d'\n", ntstix.SrcByteCount)
	str += fmt.Sprintf("dst_byte_count: '%d'\n", ntstix.DstByteCount)
	str += fmt.Sprintf("src_packets: '%d'\n", ntstix.SrcPackets)
	str += fmt.Sprintf("dst_packets: '%d'\n", ntstix.DstPackets)
	str += fmt.Sprintf("ipfix: \n%v", func(l map[string]DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, v)
		}
		return str
	}(ntstix.IPFix))
	str += fmt.Sprintf("src_payload_ref: '%v'\n", ntstix.SrcPayloadRef)
	str += fmt.Sprintf("dst_payload_ref: '%v'\n", ntstix.DstPayloadRef)
	str += fmt.Sprintf("encapsulates_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tencapsulates_ref '%d': '%v'\n", k, v)
		}
		return str
	}(ntstix.EncapsulatesRefs))
	str += fmt.Sprintf("encapsulated_by_ref: '%v'\n", ntstix.EncapsulatedByRef)
	str += fmt.Sprintln("extensions:")
	for k, v := range ntstix.Extensions {
		str += fmt.Sprintf("\t%s:\n%v\n", k, toStringBeautiful(v))
	}

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
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		IsHidden:             commonObject.IsHidden,
		PID:                  commonObject.PID,
		CreatedTime:          commonObject.CreatedTime,
		Cwd:                  commonObject.Cwd,
		CommandLine:          commonObject.CommandLine,
		EnvironmentVariables: commonObject.EnvironmentVariables,
		OpenedConnectionRefs: commonObject.OpenedConnectionRefs,
		CreatorUserRef:       commonObject.CreatorUserRef,
		ImageRef:             commonObject.ImageRef,
	}

	if len(commonObject.Extensions) == 0 {
		return pstix, nil
	}

	ext := map[string]interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := decodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = e
	}
	pstix.Extensions = ext

	return pstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (pstix ProcessCyberObservableObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(pstix)

	return &result, err
}

//ValidateStruct является валидатором параметров содержащихся в типе ProcessCyberObservableObjectSTIX
func (pstix ProcessCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(process--)[0-9a-f|-]+$`).MatchString(pstix.ID)) {
		return false
	}

	if !pstix.validateStructCommonFields() {
		return false
	}

	if len(pstix.OpenedConnectionRefs) > 0 {
		for _, v := range pstix.OpenedConnectionRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !pstix.CreatorUserRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !pstix.ImageRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !pstix.ParentRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(pstix.ChildRefs) > 0 {
		for _, v := range pstix.ChildRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(pstix.Extensions) > 0 {
		for _, v := range pstix.Extensions {
			if !checkingExtensionsSTIX(v) {
				return false
			}
		}
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (pstix ProcessCyberObservableObjectSTIX) SanitizeStruct() ProcessCyberObservableObjectSTIX {
	pstix.OptionalCommonPropertiesCyberObservableObjectSTIX = pstix.sanitizeStruct()

	pstix.Cwd = commonlibs.StringSanitize(pstix.Cwd)
	pstix.CommandLine = commonlibs.StringSanitize(pstix.CommandLine)

	sizeEnv := len(pstix.EnvironmentVariables)
	if sizeEnv > 0 {
		tmp := make(map[string]DictionaryTypeSTIX, sizeEnv)
		for k, v := range pstix.EnvironmentVariables {
			switch v := v.dictionary.(type) {
			case string:
				tmp[k] = DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = DictionaryTypeSTIX{v}
			}
		}
		pstix.EnvironmentVariables = tmp
	}

	esize := len(pstix.Extensions)
	tmp := make(map[string]interface{}, esize)
	for k, v := range pstix.Extensions {
		result := sanitizeExtensionsSTIX(v)
		tmp[k] = result
	}
	pstix.Extensions = tmp

	return pstix
}

//GetID возвращает ID STIX объекта
func (pstix ProcessCyberObservableObjectSTIX) GetID() string {
	return pstix.ID
}

//GetType возвращает Type STIX объекта
func (pstix ProcessCyberObservableObjectSTIX) GetType() string {
	return pstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа ProcessCyberObservableObjectSTIX
func (pstix ProcessCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             pstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (pstix ProcessCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := pstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += pstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("is_hidden: '%v'\n", pstix.IsHidden)
	str += fmt.Sprintf("pid: '%d'\n", pstix.PID)
	str += fmt.Sprintf("created_time: '%v'\n", pstix.CreatedTime)
	str += fmt.Sprintf("cwd: '%s'\n", pstix.Cwd)
	str += fmt.Sprintf("command_line: '%s'\n", pstix.CommandLine)
	str += fmt.Sprintf("environment_variables: \n%v", func(l map[string]DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, v)
		}
		return str
	}(pstix.EnvironmentVariables))
	str += fmt.Sprintf("opened_connection_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\topened_connection_ref '%d': '%v'\n", k, v)
		}
		return str
	}(pstix.OpenedConnectionRefs))
	str += fmt.Sprintf("creator_user_ref: '%v'\n", pstix.CreatorUserRef)
	str += fmt.Sprintf("image_ref: '%v'\n", pstix.ImageRef)
	str += fmt.Sprintf("parent_ref: '%v'\n", pstix.ParentRef)
	str += fmt.Sprintf("child_refs: \n%v", func(l []IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tchild_ref '%d': '%v'\n", k, v)
		}
		return str
	}(pstix.ChildRefs))
	str += fmt.Sprintln("extensions:")
	for k, v := range pstix.Extensions {
		str += fmt.Sprintf("\t%s:\n%v\n", k, toStringBeautiful(v))
	}

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

//ValidateStruct является валидатором параметров содержащихся в типе SoftwareCyberObservableObjectSTIX
func (sstix SoftwareCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(software--)[0-9a-f|-]+$`).MatchString(sstix.ID)) {
		return false
	}

	if !sstix.validateStructCommonFields() {
		return false
	}

	if sstix.Name == "" {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (sstix SoftwareCyberObservableObjectSTIX) SanitizeStruct() SoftwareCyberObservableObjectSTIX {
	sstix.OptionalCommonPropertiesCyberObservableObjectSTIX = sstix.sanitizeStruct()

	sstix.Name = commonlibs.StringSanitize(sstix.Name)
	sstix.CPE = commonlibs.StringSanitize(sstix.CPE)
	sstix.SwID = commonlibs.StringSanitize(sstix.SwID)
	sizel := len(sstix.Languages)

	if sizel > 0 {
		tmp := make([]string, 0, sizel)
		for _, v := range sstix.Languages {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}
		sstix.Languages = tmp
	}

	sstix.Vendor = commonlibs.StringSanitize(sstix.Vendor)
	sstix.Version = commonlibs.StringSanitize(sstix.Version)

	return sstix
}

//GetID возвращает ID STIX объекта
func (sstix SoftwareCyberObservableObjectSTIX) GetID() string {
	return sstix.ID
}

//GetType возвращает Type STIX объекта
func (sstix SoftwareCyberObservableObjectSTIX) GetType() string {
	return sstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа SoftwareCyberObservableObjectSTIX
func (sstix SoftwareCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             sstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (sstix SoftwareCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := sstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += sstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("name: '%s'\n", sstix.Name)
	str += fmt.Sprintf("cpe: '%s'\n", sstix.CPE)
	str += fmt.Sprintf("swid: '%s'\n", sstix.SwID)
	str += fmt.Sprintf("languages: \n%v", func(l []string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tlanguage '%d': '%s'\n", k, v)
		}
		return str
	}(sstix.Languages))
	str += fmt.Sprintf("vendor: '%s'\n", sstix.Vendor)
	str += fmt.Sprintf("version: '%s'\n", sstix.Version)

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

//ValidateStruct является валидатором параметров содержащихся в типе URLCyberObservableObjectSTIX
func (urlstix URLCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(url--)[0-9a-f|-]+$`).MatchString(urlstix.ID)) {
		return false
	}

	if !urlstix.validateStructCommonFields() {
		return false
	}

	if !govalidator.IsURL(urlstix.Value) {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (urlstix URLCyberObservableObjectSTIX) SanitizeStruct() URLCyberObservableObjectSTIX {
	urlstix.OptionalCommonPropertiesCyberObservableObjectSTIX = urlstix.sanitizeStruct()

	return urlstix
}

//GetID возвращает ID STIX объекта
func (urlstix URLCyberObservableObjectSTIX) GetID() string {
	return urlstix.ID
}

//GetType возвращает Type STIX объекта
func (urlstix URLCyberObservableObjectSTIX) GetType() string {
	return urlstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа URLCyberObservableObjectSTIX
func (urlstix URLCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             urlstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (urlstix URLCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := urlstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += urlstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", urlstix.Value)

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

//ValidateStruct является валидатором параметров содержащихся в типе UserAccountCyberObservableObjectSTIX
func (uastix UserAccountCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(user-account--)[0-9a-f|-]+$`).MatchString(uastix.ID)) {
		return false
	}

	if !uastix.validateStructCommonFields() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (uastix UserAccountCyberObservableObjectSTIX) SanitizeStruct() UserAccountCyberObservableObjectSTIX {
	uastix.OptionalCommonPropertiesCyberObservableObjectSTIX = uastix.sanitizeStruct()

	uastix.UserID = commonlibs.StringSanitize(uastix.UserID)
	uastix.Credential = commonlibs.StringSanitize(uastix.Credential)
	uastix.AccountLogin = commonlibs.StringSanitize(uastix.AccountLogin)
	uastix.AccountType = uastix.AccountType.SanitizeStructOpenVocabTypeSTIX()
	uastix.DisplayName = commonlibs.StringSanitize(uastix.DisplayName)

	esize := len(uastix.Extensions)
	tmp := make(map[string]UNIXAccountExtensionSTIX, esize)
	for k, v := range uastix.Extensions {
		result := sanitizeExtensionsSTIX(v)
		if ct, ok := result.(UNIXAccountExtensionSTIX); ok {
			tmp[k] = ct
		}
	}
	uastix.Extensions = tmp

	return uastix
}

//GetID возвращает ID STIX объекта
func (uastix UserAccountCyberObservableObjectSTIX) GetID() string {
	return uastix.ID
}

//GetType возвращает Type STIX объекта
func (uastix UserAccountCyberObservableObjectSTIX) GetType() string {
	return uastix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа UserAccountCyberObservableObjectSTIX
func (uastix UserAccountCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             uastix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (uastix UserAccountCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := uastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += uastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("user_id: '%s'\n", uastix.UserID)
	str += fmt.Sprintf("credential: '%s'\n", uastix.Credential)
	str += fmt.Sprintf("account_login: '%s'\n", uastix.AccountLogin)
	str += fmt.Sprintf("account_type: '%v'\n", uastix.AccountType)
	str += fmt.Sprintf("display_name: '%s'\n", uastix.DisplayName)
	str += fmt.Sprintf("is_service_account: '%v'\n", uastix.IsServiceAccount)
	str += fmt.Sprintf("is_privileged: '%v'\n", uastix.IsPrivileged)
	str += fmt.Sprintf("can_escalate_privs: '%v'\n", uastix.CanEscalatePrivs)
	str += fmt.Sprintf("is_disabled: '%v'\n", uastix.IsDisabled)
	str += fmt.Sprintf("account_created: '%v'\n", uastix.AccountCreated)
	str += fmt.Sprintf("account_expires: '%v'\n", uastix.AccountExpires)
	str += fmt.Sprintf("credential_last_changed: '%v'\n", uastix.CredentialLastChanged)
	str += fmt.Sprintf("account_first_login: '%v'\n", uastix.AccountFirstLogin)
	str += fmt.Sprintf("account_last_login: '%v'\n", uastix.AccountLastLogin)
	str += fmt.Sprintln("extensions:")
	for k, v := range uastix.Extensions {
		str += fmt.Sprintf("\t%s:\n%v\n", k, toStringBeautiful(v))
	}

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

//ValidateStruct является валидатором параметров содержащихся в типе WindowsRegistryKeyCyberObservableObjectSTIX
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(windows-registry-key--)[0-9a-f|-]+$`).MatchString(wrkstix.ID)) {
		return false
	}

	if !wrkstix.validateStructCommonFields() {
		return false
	}

	if !wrkstix.CreatorUserRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) SanitizeStruct() WindowsRegistryKeyCyberObservableObjectSTIX {
	wrkstix.OptionalCommonPropertiesCyberObservableObjectSTIX = wrkstix.sanitizeStruct()

	wrkstix.Key = commonlibs.StringSanitize(wrkstix.Key)

	sizev := len(wrkstix.Values)
	if sizev > 0 {
		tmp := make([]WindowsRegistryValueTypeSTIX, 0, sizev)

		for _, v := range wrkstix.Values {
			tmp = append(tmp, v.SanitizeStructWindowsRegistryValueTypeSTIX())
		}

		wrkstix.Values = tmp
	}

	return wrkstix
}

//GetID возвращает ID STIX объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) GetID() string {
	return wrkstix.ID
}

//GetType возвращает Type STIX объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) GetType() string {
	return wrkstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа WindowsRegistryKeyCyberObservableObjectSTIX
func (wrstix WindowsRegistryKeyCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             wrstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := wrkstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += wrkstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("key: '%s'\n", wrkstix.Key)
	str += fmt.Sprintf("values: \n%v", func(l []WindowsRegistryValueTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tvalue '%d':\n", k)
			str += fmt.Sprintf("\t\tname: '%s'\n", v.Name)
			str += fmt.Sprintf("\t\tdata: '%s'\n", v.Data)
			str += fmt.Sprintf("\t\tdata_type: '%s'\n", v.DataType)
		}
		return str
	}(wrkstix.Values))
	str += fmt.Sprintf("modified_time: '%v'\n", wrkstix.ModifiedTime)
	str += fmt.Sprintf("creator_user_ref: '%v'\n", wrkstix.CreatorUserRef)
	str += fmt.Sprintf("number_of_subkeys: '%d'\n", wrkstix.NumberOfSubkeys)

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

//ValidateStruct является валидатором параметров содержащихся в типе X509CertificateCyberObservableObjectSTIX
func (x509sstix X509CertificateCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(x509-certificate--)[0-9a-f|-]+$`).MatchString(x509sstix.ID)) {
		return false
	}

	if !x509sstix.validateStructCommonFields() {
		return false
	}

	if !x509sstix.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	return true
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (x509sstix X509CertificateCyberObservableObjectSTIX) SanitizeStruct() X509CertificateCyberObservableObjectSTIX {
	x509sstix.OptionalCommonPropertiesCyberObservableObjectSTIX = x509sstix.sanitizeStruct()

	x509sstix.Version = commonlibs.StringSanitize(x509sstix.Version)
	x509sstix.SerialNumber = commonlibs.StringSanitize(x509sstix.SerialNumber)
	x509sstix.SignatureAlgorithm = commonlibs.StringSanitize(x509sstix.SignatureAlgorithm)
	x509sstix.Issuer = commonlibs.StringSanitize(x509sstix.Issuer)
	x509sstix.Subject = commonlibs.StringSanitize(x509sstix.Subject)
	x509sstix.SubjectPublicKeyAlgorithm = commonlibs.StringSanitize(x509sstix.SubjectPublicKeyAlgorithm)
	x509sstix.SubjectPublicKeyModulus = commonlibs.StringSanitize(x509sstix.SubjectPublicKeyModulus)

	x509sstix.X509V3Extensions = x509sstix.X509V3Extensions.SanitizeStructX509V3ExtensionsTypeSTIX()

	return x509sstix
}

//GetID возвращает ID STIX объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) GetID() string {
	return x509sstix.ID
}

//GetType возвращает Type STIX объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) GetType() string {
	return x509sstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа X509CertificateCyberObservableObjectSTIX
func (x509sstix X509CertificateCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             x509sstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (x509sstix X509CertificateCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := x509sstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += x509sstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("is_self_signed: '%v'\n", x509sstix.IsSelfSigned)
	str += fmt.Sprintf("hashes: \n%v", func(l map[string]string) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, v)
		}
		return str
	}(x509sstix.Hashes))
	str += fmt.Sprintf("version: '%s'\n", x509sstix.Version)
	str += fmt.Sprintf("serial_number: '%s'\n", x509sstix.SerialNumber)
	str += fmt.Sprintf("signature_algorithm: '%s'\n", x509sstix.SignatureAlgorithm)
	str += fmt.Sprintf("issuer: '%s'\n", x509sstix.Issuer)
	str += fmt.Sprintf("validity_not_before: '%v'\n", x509sstix.ValidityNotBefore)
	str += fmt.Sprintf("validity_not_after: '%v'\n", x509sstix.ValidityNotAfter)
	str += fmt.Sprintf("subject: '%s'\n", x509sstix.Subject)
	str += fmt.Sprintf("subject_public_key_algorithm: '%s'\n", x509sstix.SubjectPublicKeyAlgorithm)
	str += fmt.Sprintf("subject_public_key_modulus: '%s'\n", x509sstix.SubjectPublicKeyModulus)
	str += fmt.Sprintf("subject_public_key_exponent: '%v'\n", x509sstix.SubjectPublicKeyExponent)
	str += fmt.Sprintln("x509_v3_extensions:")
	str += fmt.Sprintf("\tbasic_constraints: '%s'\n", x509sstix.X509V3Extensions.BasicConstraints)
	str += fmt.Sprintf("\tname_constraints: '%s'\n", x509sstix.X509V3Extensions.NameConstraints)
	str += fmt.Sprintf("\tpolicy_contraints: '%s'\n", x509sstix.X509V3Extensions.PolicyContraints)
	str += fmt.Sprintf("\tkey_usage: '%s'\n", x509sstix.X509V3Extensions.KeyUsage)
	str += fmt.Sprintf("\textended_key_usage: '%s'\n", x509sstix.X509V3Extensions.ExtendedKeyUsage)
	str += fmt.Sprintf("\tsubject_key_identifier: '%s'\n", x509sstix.X509V3Extensions.SubjectKeyIdentifier)
	str += fmt.Sprintf("\tauthority_key_identifier: '%s'\n", x509sstix.X509V3Extensions.AuthorityKeyIdentifier)
	str += fmt.Sprintf("\tsubject_alternative_name: '%s'\n", x509sstix.X509V3Extensions.SubjectAlternativeName)
	str += fmt.Sprintf("\tissuer_alternative_name: '%s'\n", x509sstix.X509V3Extensions.IssuerAlternativeName)
	str += fmt.Sprintf("\tsubject_directory_attributes: '%s'\n", x509sstix.X509V3Extensions.SubjectDirectoryAttributes)
	str += fmt.Sprintf("\tcrl_distribution_points: '%s'\n", x509sstix.X509V3Extensions.CrlDistributionPoints)
	str += fmt.Sprintf("\tinhibit_any_policy: '%s'\n", x509sstix.X509V3Extensions.InhibitAnyPolicy)
	str += fmt.Sprintf("\tprivate_key_usage_period_not_before: '%v'\n", x509sstix.X509V3Extensions.PrivateKeyUsagePeriodNotBefore)
	str += fmt.Sprintf("\tprivate_key_usage_period_not_after: '%v'\n", x509sstix.X509V3Extensions.PrivateKeyUsagePeriodNotAfter)
	str += fmt.Sprintf("\tcertificate_policies: '%s'\n", x509sstix.X509V3Extensions.CertificatePolicies)
	str += fmt.Sprintf("\tpolicy_mappings: '%s'\n", x509sstix.X509V3Extensions.PolicyMappings)

	return str
}
