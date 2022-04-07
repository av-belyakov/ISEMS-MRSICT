package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"ISEMS-MRSICT/commonlibs"
)

/*****************************************************************************/
/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/
/*****************************************************************************/

func (oc *OptionalCommonPropertiesRelationshipObjectSTIX) validateStructCommonFields() bool {
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(oc.SpecVersion)) {
		return false
	}

	if oc.Created.Unix() < 0 {
		return false
	}

	if oc.Modified.Unix() < 0 {
		return false
	}

	return true
}

//ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (oc *OptionalCommonPropertiesRelationshipObjectSTIX) ToStringBeautiful() string {
	var str string
	str += fmt.Sprintf("spec_version: '%s'\n", oc.SpecVersion)
	str += fmt.Sprintf("created: '%v'\n", oc.Created)
	str += fmt.Sprintf("modified: '%v'\n", oc.Modified)

	return str
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

//ValidateStruct является валидатором параметров содержащихся в типе RelationshipObjectSTIX
func (rstix RelationshipObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(relationship--)[0-9a-f|-]+$`).MatchString(rstix.ID)) {
		return false
	}

	if !rstix.validateStructCommonFields() {
		return false
	}

	if !(regexp.MustCompile(`^[0-9a-z|-]+$`).MatchString(rstix.RelationshipType)) {
		return false
	}

	if !rstix.SourceRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !rstix.TargetRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return rstix.validateStructCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (rstix RelationshipObjectSTIX) SanitizeStruct() RelationshipObjectSTIX {
	rstix.Description = commonlibs.StringSanitize(rstix.Description)

	return rstix
}

//GetID возвращает ID STIX объекта
func (rstix RelationshipObjectSTIX) GetID() string {
	return rstix.ID
}

//GetType возвращает Type STIX объекта
func (rstix RelationshipObjectSTIX) GetType() string {
	return rstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа RelationshipObjectSTIX
func (rstix RelationshipObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             rstix.ID,
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
func (rstix RelationshipObjectSTIX) ToStringBeautiful() string {
	str := rstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += rstix.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("relationship_type: '%s'\n", rstix.RelationshipType)
	str += fmt.Sprintf("description: '%s'\n", rstix.Description)
	str += fmt.Sprintf("source_ref: '%v'\n", rstix.SourceRef)
	str += fmt.Sprintf("target_ref: '%v'\n", rstix.TargetRef)
	str += fmt.Sprintf("start_time: '%v'\n", rstix.StartTime)
	str += fmt.Sprintf("stop_time: '%v'\n", rstix.StopTime)

	return str
}

//GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (rstix RelationshipObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   rstix.ID,
		"type": rstix.Type,
	}
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

//ValidateStruct является валидатором параметров содержащихся в типе SightingObjectSTIX
func (sstix SightingObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(sighting--)[0-9a-f|-]+$`).MatchString(sstix.ID)) {
		return false
	}

	if !sstix.validateStructCommonFields() {
		return false
	}

	if !sstix.SightingOfRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(sstix.ObservedDataRefs) > 0 {
		for _, v := range sstix.ObservedDataRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(sstix.WhereSightedRefs) > 0 {
		for _, v := range sstix.WhereSightedRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return sstix.validateStructCommonFields()
}

//SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (sstix SightingObjectSTIX) SanitizeStruct() SightingObjectSTIX {
	sstix.Description = commonlibs.StringSanitize(sstix.Description)

	return sstix
}

//GetID возвращает ID STIX объекта
func (sstix SightingObjectSTIX) GetID() string {
	return sstix.ID
}

//GetType возвращает Type STIX объекта
func (sstix SightingObjectSTIX) GetType() string {
	return sstix.Type
}

//ComparisonTypeCommonFields выполняет сравнение двух объектов типа SightingObjectSTIX
func (sstix SightingObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
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
func (sstix SightingObjectSTIX) ToStringBeautiful() string {
	str := sstix.CommonPropertiesObjectSTIX.ToStringBeautiful()
	str += sstix.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful()
	str += fmt.Sprintf("description: '%s'\n", sstix.Description)
	str += fmt.Sprintf("first_seen: '%v'\n", sstix.FirstSeen)
	str += fmt.Sprintf("last_seen: '%v'\n", sstix.LastSeen)
	str += fmt.Sprintf("count: '%d'\n", sstix.Count)
	str += fmt.Sprintf("sighting_of_ref: '%v'\n", sstix.SightingOfRef)
	str += fmt.Sprintf("observed_data_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\tobserved_data_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(sstix.ObservedDataRefs))
	str += fmt.Sprintf("where_sighted_refs: \n%v", func(l []*IdentifierTypeSTIX) string {
		var str string
		for k, v := range l {
			str += fmt.Sprintf("\twhere_sighted_ref '%d': '%v'\n", k, *v)
		}
		return str
	}(sstix.WhereSightedRefs))
	str += fmt.Sprintf("summary: '%v'\n", sstix.Summary)

	return str
}

//GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (sstix SightingObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   sstix.ID,
		"type": sstix.Type,
	}
}
