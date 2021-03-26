package datamodels

import (
	"encoding/json"
	"fmt"
)

/*****************************************************************************/
/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/
/*****************************************************************************/

func (ocprstix *OptionalCommonPropertiesRelationshipObjectSTIX) checkingTypeCommonFields() bool {
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

/* --- RelationshipObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (rstix RelationshipObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &rstix); err != nil {
		return nil, err
	}

	return rstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (rstix RelationshipObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(rstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе RelationshipObjectSTIX
// возвращает ВАЛИДНЫЙ объект RelationshipObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.RelationshipObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (rstix RelationshipObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !rstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

/* --- SightingObjectSTIX --- */

//DecoderJSON выполняет декодирование JSON объекта
func (sstix SightingObjectSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &sstix); err != nil {
		return nil, err
	}

	return sstix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (sstix SightingObjectSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(sstix)

	return &result, err
}

//CheckingTypeFields является валидатором параметров содержащихся в типе SightingObjectSTIX
// возвращает ВАЛИДНЫЙ объект SightingObjectSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.SightingObjectSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method CheckingTypeFields (CheckingTypeFields has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник CheckingTypeFields работает по ссылке)
func (sstix SightingObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}
