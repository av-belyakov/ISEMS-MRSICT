package datamodels

import (
	"encoding/json"
	"fmt"
)

/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/

/*
//OptionalCommonPropertiesRelationshipObjectSTIX общие, опциональные свойства для все объектов STIX типа Relationship Objects
// SpecVersion - версия STIX спецификации (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Modified - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type OptionalCommonPropertiesRelationshipObjectSTIX struct {
	SpecVersion string    `json:"spec_version" bson:"spec_version"`
	Created     time.Time `json:"created" bson:"created"`
	Modified    time.Time `json:"modified" bson:"modified"`
}
*/

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

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (rstix RelationshipObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !rstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

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

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (sstix SightingObjectSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !sstix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}
