package datamodels

import (
	"encoding/json"
	"fmt"
)

//DecoderJSON выполняет декодирование JSON объекта
func (apdostix *AttackPatternDomainObjectsSTIX) DecoderJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, apdostix); err != nil {
		return nil, err
	}

	return apdostix, nil
}

//EncoderJSON выполняет кодирование в JSON объект
func (apdostix *AttackPatternDomainObjectsSTIX) EncoderJSON(interface{}) (*[]byte, error) {
	return &[]byte{}, nil
}

//CheckingTypeFields является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
func (apdostix *AttackPatternDomainObjectsSTIX) CheckingTypeFields() bool {
	fmt.Println("func 'CheckingTypeFields', START...")

	if !apdostix.checkingTypeCommonFields() {
		return false
	}

	//тут проверяем остальные параметры, не входящие в тип CommonPropertiesDomainObjectSTIX

	return true
}

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
	SpecVersion        string                     `json:"spec_version" bson:"spec_version" required:"true"`
	Created            time.Time                  `json:"created" bson:"created" required:"true"`
	Modified           time.Time                  `json:"modified" bson:"modified" required:"true"`
	CreatedByRef       IdentifierTypeSTIX         `json:"created_by_ref" bson:"created_by_ref"`
	Revoked            bool                       `json:"revoked" bson:"revoked"`
	Labels             []string                   `json:"labels" bson:"labels"`
	Сonfidence         int                        `json:"confidence" bson:"confidence"`
	Lang               string                     `json:"lang" bson:"lang"`
	ExternalReferences ExternalReferencesTypeSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []*IdentifierTypeSTIX      `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   GranularMarkingsTypeSTIX   `json:"granular_markings" bson:"granular_markings"`
	Defanged           bool                       `json:"defanged" bson:"defanged"`
	Extensions         map[string]string          `json:"extensions" bson:"extensions"`
}
*/
func (cpdostix *CommonPropertiesDomainObjectSTIX) checkingTypeCommonFields() bool {
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
