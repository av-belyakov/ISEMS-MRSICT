package datamodels

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("func 'checkingTypeCommonFields', START...")

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

/* --- ArtifactCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ArtifactCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект ArtifactCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ArtifactCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (astix ArtifactCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !astix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- AutonomousSystemCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе AutonomousSystemCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект AutonomousSystemCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.AutonomousSystemCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (asstix AutonomousSystemCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !asstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- DirectoryCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе DirectoryCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект DirectoryCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.DirectoryCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (dstix DirectoryCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !dstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- DomainNameCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе DomainNameCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект DomainNameCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.DomainNameCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (dnstix DomainNameCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !dnstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
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

//CheckingTypeFields является валидатором параметров содержащихся в типе EmailAddressCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект EmailAddressCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.EmailAddressCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (eastix EmailAddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !eastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- EmailMessageCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе EmailMessageCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект EmailMessageCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.EmailMessageCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (emstix EmailMessageCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !emstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//GetFileCyberObservableObjectSTIX возвращает объект типа FileCyberObservableObjectSTIX
func (fstix *FileCyberObservableObjectSTIX) GetFileCyberObservableObjectSTIX() *FileCyberObservableObjectSTIX {
	return fstix
}

/* --- EmailMessageCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе FileCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект FileCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.FileCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (fstix FileCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !fstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

//GetIPv4AddressCyberObservableObjectSTIX выполняет объект типа IPv4AddressCyberObservableObjectSTIX
func (fstix *IPv4AddressCyberObservableObjectSTIX) GetIPv4AddressCyberObservableObjectSTIX() *IPv4AddressCyberObservableObjectSTIX {
	return fstix
}

/* --- IPv4AddressCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе IPv4AddressCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект IPv4AddressCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.IPv4AddressCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (ip4stix IPv4AddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !ip4stix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- IPv6AddressCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе IPv6AddressCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект IPv6AddressCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.IPv6AddressCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (ip6stix IPv6AddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !ip6stix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- MACAddressCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе MACAddressCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект MACAddressCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.MACAddressCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (macstix MACAddressCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !macstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- MutexCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе MutexCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект MutexCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.MutexCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (mstix MutexCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !mstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- NetworkTrafficCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе NetworkTrafficCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект NetworkTrafficCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.NetworkTrafficCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (ntstix NetworkTrafficCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !ntstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- ProcessCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе ProcessCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект ProcessCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.ProcessCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (pstix ProcessCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !pstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- SoftwareCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе SoftwareCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект SoftwareCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.SoftwareCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (sstix SoftwareCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- URLCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе URLCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект URLCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.URLCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (urlstix URLCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !urlstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- UserAccountCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе UserAccountCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект UserAccountCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.UserAccountCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (uastix UserAccountCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !uastix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- WindowsRegistryKeyCyberObservableObjectSTIX --- */

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

//CheckingTypeFields является валидатором параметров содержащихся в типе WindowsRegistryKeyCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект WindowsRegistryKeyCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.WindowsRegistryKeyCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !wrkstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- X509CertificateCyberObservableObjectSTIX --- */

//CheckingTypeFields является валидатором параметров содержащихся в типе X509CertificateCyberObservableObjectSTIX
// возвращает ВАЛИДНЫЙ объект X509CertificateCyberObservableObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.X509CertificateCyberObservableObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (x509sstix X509CertificateCyberObservableObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !x509sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
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
