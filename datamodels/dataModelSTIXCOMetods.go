package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"

	"ISEMS-MRSICT/commonlibs"

	"github.com/asaskevich/govalidator"
)

/*********************************************************************************/
/********** 			Cyber-observable Objects STIX (МЕТОДЫ)			**********/
/*********************************************************************************/

func (ocpcstix *OptionalCommonPropertiesCyberObservableObjectSTIX) checkingTypeCommonFields() bool {
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

	//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
	if ok := ocpcstix.GranularMarkings.CheckGranularMarkingsTypeSTIX(); !ok {
		return false
	}

	return true
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) sanitizeStruct() OptionalCommonPropertiesCyberObservableObjectSTIX {
	//обработка содержимого списка поля Extensions
	if len(ocpcstix.Extensions) > 0 {
		ext := make(map[string]*DictionaryTypeSTIX, len(ocpcstix.Extensions))
		for k, v := range ocpcstix.Extensions {
			switch v := v.dictionary.(type) {
			case string:
				ext[k] = &DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				ext[k] = &DictionaryTypeSTIX{v}
			}
		}
		ocpcstix.Extensions = ext
	}

	return ocpcstix
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) ToStringBeautiful() string {
	var str string
	str += fmt.Sprintf("spec_version: '%s'\n", ocpcstix.SpecVersion)
	str += fmt.Sprintf("object_marking_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobject_marking_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(ocpcstix.ObjectMarkingRefs))
	str += fmt.Sprintln("granular_markings:")
	str += fmt.Sprintf("\tlang: '%s'\n", ocpcstix.GranularMarkings.Lang)
	str += fmt.Sprintf("\tmarking_ref: '%v'\n", ocpcstix.GranularMarkings.MarkingRef)
	str += fmt.Sprintf("defanged: '%v'\n", ocpcstix.Defanged)
	str += fmt.Sprintf("extensions: \n%v", func(l map[string]*DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, *v)
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

//CheckingTypeFields является валидатором параметров содержащихся в типе ArtifactCyberObservableObjectSTIX
func (astix ArtifactCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(artifact--)[0-9a-f|-]+$`).MatchString(astix.ID)) {
		return false
	}

	if !astix.checkingTypeCommonFields() {
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

//CheckingTypeFields является валидатором параметров содержащихся в типе AutonomousSystemCyberObservableObjectSTIX
func (asstix AutonomousSystemCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(autonomous-system--)[0-9a-f|-]+$`).MatchString(asstix.ID)) {
		return false
	}

	if !asstix.checkingTypeCommonFields() {
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

//CheckingTypeFields является валидатором параметров содержащихся в типе DirectoryCyberObservableObjectSTIX
func (dstix DirectoryCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(directory--)[0-9a-f|-]+$`).MatchString(dstix.ID)) {
		return false
	}

	if !dstix.checkingTypeCommonFields() {
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

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dstix DirectoryCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("path: '%s'\n", dstix.Path)
	str += fmt.Sprintf("path_enc: '%s'\n", dstix.PathEnc)
	str += fmt.Sprintf("ctime: '%v'\n", dstix.Ctime)
	str += fmt.Sprintf("mtime: '%v'\n", dstix.Mtime)
	str += fmt.Sprintf("atime: '%s'\n", dstix.Atime)
	str += fmt.Sprintf("contains_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcontains_ref '%d': '%v'\n", k, *v)
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

//CheckingTypeFields является валидатором параметров содержащихся в типе DomainNameCyberObservableObjectSTIX
func (dnstix DomainNameCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(domain-name--)[0-9a-f|-]+$`).MatchString(dnstix.ID)) {
		return false
	}

	if !dnstix.checkingTypeCommonFields() {
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

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (dnstix DomainNameCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := dnstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += dnstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", dnstix.Value)
	str += fmt.Sprintf("resolves_to_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tresolves_to_ref '%d': '%v'\n", k, *v)
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

//CheckingTypeFields является валидатором параметров содержащихся в типе EmailAddressCyberObservableObjectSTIX
func (eastix EmailAddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	if !(regexp.MustCompile(`^(email-addr--)[0-9a-f|-]+$`).MatchString(eastix.ID)) {
		return false
	}

	if !eastix.checkingTypeCommonFields() {
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

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (eastix EmailAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := eastix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += eastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("value: '%s'\n", eastix.Value)
	str += fmt.Sprintf("display_name: '%s'\n", eastix.DisplayName)
	str += fmt.Sprintf("belongs_to_ref: '%v'\n", eastix.BelongsToRef)

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
		tmp := make(map[string]*DictionaryTypeSTIX, len(emstix.AdditionalHeaderFields))
		for k, v := range emstix.AdditionalHeaderFields {
			switch v := v.dictionary.(type) {
			case string:
				tmp[k] = &DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = &DictionaryTypeSTIX{v}
			}
		}
		emstix.AdditionalHeaderFields = tmp
	}

	emstix.Body = commonlibs.StringSanitize(emstix.Body)

	if len(emstix.BodyMultipart) > 0 {
		tmp := make([]*EmailMIMEPartTypeSTIX, 0, len(emstix.BodyMultipart))
		for _, v := range emstix.BodyMultipart {
			tmp = append(tmp, &EmailMIMEPartTypeSTIX{
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

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (emstix EmailMessageCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := emstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += emstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("is_multipart: '%v'\n", emstix.IsMultipart)
	str += fmt.Sprintf("date: '%v'\n", emstix.Date)
	str += fmt.Sprintf("content_type: '%s'\n", emstix.ContentType)
	str += fmt.Sprintf("from_ref: '%v'\n", emstix.FromRef)
	str += fmt.Sprintf("sender_ref: '%v'\n", emstix.SenderRef)
	str += fmt.Sprintf("to_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tto_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(emstix.ToRefs))
	str += fmt.Sprintf("cc_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcc_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(emstix.CcRefs))
	str += fmt.Sprintf("bcc_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tbcc_ref '%d': '%v'\n", k, *v)
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
	str += fmt.Sprintf("additional_header_fields: \n%v", func(l map[string]*DictionaryTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, *v)
		}
		return str
	}(emstix.AdditionalHeaderFields))
	str += fmt.Sprintf("body: '%v'\n", emstix.Body)
	str += fmt.Sprintf("body_multipart: \n%v", func(l []*EmailMIMEPartTypeSTIX) string {
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
	if len(fstix.Extensions) > 0 {
		tmp := make(map[string]*interface{}, len(fstix.Extensions))
		for k, v := range fstix.Extensions {
			switch v := (*v).(type) {
			case string:
				var str interface{} = commonlibs.StringSanitize(string(v))
				tmp[k] = &str
			default:
				tmp[k] = &v
			}
		}
		fstix.Extensions = tmp
	}

	return fstix
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
	str += fmt.Sprintf("contains_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tcontains_ref '%d': '%v'\n", k, v)
		}
		return str
	}(fstix.ContainsRefs))
	str += fmt.Sprintf("content_ref: '%v'\n", fstix.ContentRef)
	str += fmt.Sprintf("extensions: \n%v", func(l map[string]*interface{}) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\t'%s': '%v'\n", k, *v)
		}
		return str
	}(fstix.Extensions))

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

//checkingExtensionsSTIX выполняет проверку полей следующих типов STIX расширений:
// - "archive-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "windows-service-ext"
func checkingExtensionsSTIX(extType interface{}) bool {
	switch et := extType.(type) {
	case ArchiveFileExtensionSTIX:
		for _, v := range et.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}

	case NTFSFileExtensionSTIX:
		if len(et.AlternateDataStreams) == 0 {
			return true
		}

		for _, v := range et.AlternateDataStreams {
			if !v.Hashes.CheckHashesTypeSTIX() {
				return false
			}
		}

	case WindowsPEBinaryFileExtensionSTIX:
		if !et.FileHeaderHashes.CheckHashesTypeSTIX() {
			return false
		}

	case HTTPRequestExtensionSTIX:
		if !et.MessageBodyDataRef.CheckIdentifierTypeSTIX() {
			return false
		}

	case WindowsServiceExtensionSTIX:
		for _, v := range et.ServiceDllRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

//sanitizeExtensionsSTIX для ряда полей следующих типов STIX расширений:
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
// выполняет замену некоторых специальных символов на их HTML код
func sanitizeExtensionsSTIX(extType interface{}) interface{} {
	sanitizeDictionaryList := func(l map[string]*DictionaryTypeSTIX) map[string]*DictionaryTypeSTIX {
		size := len(l)

		if size == 0 {
			return map[string]*DictionaryTypeSTIX{}
		}

		tmp := make(map[string]*DictionaryTypeSTIX, size)
		for k, v := range l {
			switch v := v.dictionary.(type) {
			case string:
				tmp[k] = &DictionaryTypeSTIX{commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = &DictionaryTypeSTIX{v}
			}
		}

		return tmp
	}

	switch et := extType.(type) {
	case ArchiveFileExtensionSTIX:
		return commonlibs.StringSanitize(et.Comment)

	case NTFSFileExtensionSTIX:
		var tmpType NTFSFileExtensionSTIX
		tmpType.SID = commonlibs.StringSanitize(et.SID)

		size := len(et.AlternateDataStreams)

		if size == 0 {
			return tmpType
		}

		ads := make([]*AlternateDataStreamTypeSTIX, 0, size)
		for _, v := range et.AlternateDataStreams {
			ads = append(ads, &AlternateDataStreamTypeSTIX{
				Name:   commonlibs.StringSanitize(v.Name),
				Hashes: v.Hashes,
				Size:   v.Size,
			})
		}
		tmpType.AlternateDataStreams = ads

		return tmpType

	case PDFFileExtensionSTIX:
		return PDFFileExtensionSTIX{
			Version:          commonlibs.StringSanitize(et.Version),
			IsOptimized:      et.IsOptimized,
			DocumentInfoDict: sanitizeDictionaryList(et.DocumentInfoDict),
			Pdfid0:           commonlibs.StringSanitize(et.Pdfid0),
			Pdfid1:           commonlibs.StringSanitize(et.Pdfid1),
		}

	case RasterImageFileExtensionSTIX:
		return RasterImageFileExtensionSTIX{
			ImageHeight:  et.ImageHeight,
			ImageWidth:   et.ImageWidth,
			BitsPerPixel: et.BitsPerPixel,
			ExifTags:     sanitizeDictionaryList(et.ExifTags),
		}

	case WindowsPEBinaryFileExtensionSTIX:

	case HTTPRequestExtensionSTIX:

	case ICMPExtensionSTIX:

	case NetworkSocketExtensionSTIX:

	case TCPExtensionSTIX:

	case WindowsProcessExtensionSTIX:

	case WindowsServiceExtensionSTIX:

	case UNIXAccountExtensionSTIX:

	}

	return extType
}
