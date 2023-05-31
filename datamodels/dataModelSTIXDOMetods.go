package datamodels

import (
	"fmt"
	"path"
	"reflect"
	"time"

	mstixo "github.com/av-belyakov/methodstixobjects"
)

/*************************************************************************/
/********** 			Domain Objects STIX (МЕТОДЫ)			**********/
/*************************************************************************/

/*func (cpdostix CommonPropertiesDomainObjectSTIX) sanitizeStruct() CommonPropertiesDomainObjectSTIX {
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

	//время модификации объекта
	cpdostix.Modified = time.Now()

	return cpdostix
}

// comparisonTypeCommonFields выполняет сравнение двух объектов типа CommonPropertiesDomainObjectSTIX
func (cpdostix *CommonPropertiesDomainObjectSTIX) comparisonTypeCommonFields(coNew *CommonPropertiesDomainObjectSTIX) (bool, []OldFieldValueObjectType) {
	var (
		isEqual bool = true
		result  []OldFieldValueObjectType
	)

	oldValue := reflect.ValueOf(*cpdostix)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*coNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				continue
			}

			fieldType := typeOfOldValue.Field(i).Type.Name()
			if fieldType == "" {
				fieldType = typeOfOldValue.Field(i).Type.Kind().String()
			}

			result = append(result, OldFieldValueObjectType{
				FeildType: fieldType,
				Path:      path.Join("CommonPropertiesDomainObjectSTIX", typeOfOldValue.Field(i).Name),
				Value:     oldValue.Field(i).Interface(),
			})

			isEqual = false
		}
	}

	return isEqual, result
}*/
// comparisonTypeCommonFields выполняет сравнение двух объектов типа CommonPropertiesDomainObjectSTIX
func comparisonTypeCommonFields(cpdostix, coNew mstixo.CommonPropertiesDomainObjectSTIX) (bool, []OldFieldValueObjectType) {
	var (
		isEqual bool = true
		result  []OldFieldValueObjectType
	)

	oldValue := reflect.ValueOf(cpdostix)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(coNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				continue
			}

			fieldType := typeOfOldValue.Field(i).Type.Name()
			if fieldType == "" {
				fieldType = typeOfOldValue.Field(i).Type.Kind().String()
			}

			result = append(result, OldFieldValueObjectType{
				FeildType: fieldType,
				Path:      path.Join("CommonPropertiesDomainObjectSTIX", typeOfOldValue.Field(i).Name),
				Value:     oldValue.Field(i).Interface(),
			})

			isEqual = false
		}
	}

	return isEqual, result
}

/* --- AttackPatternDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (apstix AttackPatternDomainObjectsSTIX) GetCurrentObject() interface{} {
	return apstix.AttackPatternDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа AttackPatternDomainObjectsSTIX
func (apstix AttackPatternDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             apstix.ID,
		}
	)

	apNew, ok := newObj.(*mstixo.AttackPatternDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(apstix.AttackPatternDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*apNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("AttackPatternDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("AttackPatternDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- CampaignDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (cstix CampaignDomainObjectsSTIX) GetCurrentObject() interface{} {
	return cstix.CampaignDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа CampaignDomainObjectsSTIX
func (cstix CampaignDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             cstix.ID,
		}
	)

	cNew, ok := newObj.(*mstixo.CampaignDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(cstix.CampaignDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*cNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("CampaignDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("CampaignDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- CourseOfActionDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (cofastix CourseOfActionDomainObjectsSTIX) GetCurrentObject() interface{} {
	return cofastix.CourseOfActionDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа CourseOfActionDomainObjectsSTIX
func (cofastix CourseOfActionDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             cofastix.ID,
		}
	)

	cofaNew, ok := newObj.(*mstixo.CourseOfActionDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(cofastix.CourseOfActionDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*cofaNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("CourseOfActionDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			fmt.Printf("\t____fucn 'ComparisonTypeCommonFields', courseOfAction, typeOfOldValue.Field(i).Type.Name(): %s, typeOfNewValue.Field(i).Type.Name(): %s\n", typeOfOldValue.Field(i).Type.Name(), typeOfNewValue.Field(i).Type.Name())

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("CourseOfActionDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- GroupingDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (gstix GroupingDomainObjectsSTIX) GetCurrentObject() interface{} {
	return gstix.GroupingDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа GroupingDomainObjectsSTIX
func (gstix GroupingDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             gstix.ID,
		}
	)

	gNew, ok := newObj.(*mstixo.GroupingDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(gstix.GroupingDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*gNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("GroupingDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("GroupingDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- IdentityDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (istix IdentityDomainObjectsSTIX) GetCurrentObject() interface{} {
	return istix.IdentityDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа IdentityDomainObjectsSTIX
func (istix IdentityDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             istix.ID,
		}
	)

	iNew, ok := newObj.(*mstixo.IdentityDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(istix.IdentityDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*iNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("IdentityDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("IdentityDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- IndicatorDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (istix IndicatorDomainObjectsSTIX) GetCurrentObject() interface{} {
	return istix.IndicatorDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа IndicatorDomainObjectsSTIX
func (istix IndicatorDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             istix.ID,
		}
	)

	iNew, ok := newObj.(*mstixo.IndicatorDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(istix.IndicatorDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*iNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("IndicatorDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("IndicatorDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- InfrastructureDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (istix InfrastructureDomainObjectsSTIX) GetCurrentObject() interface{} {
	return istix.InfrastructureDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа InfrastructureDomainObjectsSTIX
func (istix InfrastructureDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             istix.ID,
		}
	)

	iNew, ok := newObj.(*mstixo.InfrastructureDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(istix.InfrastructureDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*iNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("InfrastructureDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("InfrastructureDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- IntrusionSetDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (istix IntrusionSetDomainObjectsSTIX) GetCurrentObject() interface{} {
	return istix.IntrusionSetDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа IntrusionSetDomainObjectsSTIX
func (istix IntrusionSetDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             istix.ID,
		}
	)

	iNew, ok := newObj.(*mstixo.IntrusionSetDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(istix.IntrusionSetDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*iNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("IntrusionSetDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("IntrusionSetDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- LocationDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (lstix LocationDomainObjectsSTIX) GetCurrentObject() interface{} {
	return lstix.LocationDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа LocationDomainObjectsSTIX
func (lstix LocationDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             lstix.ID,
		}
	)

	lNew, ok := newObj.(*mstixo.LocationDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(lstix.LocationDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*lNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("LocationDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("LocationDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- MalwareDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (mstix MalwareDomainObjectsSTIX) GetCurrentObject() interface{} {
	return mstix.MalwareDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа MalwareDomainObjectsSTIX
func (mstix MalwareDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             mstix.ID,
		}
	)

	mNew, ok := newObj.(*mstixo.MalwareDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(mstix.MalwareDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*mNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("MalwareDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("MalwareDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- MalwareAnalysisDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (mstix MalwareAnalysisDomainObjectsSTIX) GetCurrentObject() interface{} {
	return mstix.MalwareAnalysisDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа MalwareAnalysisDomainObjectsSTIX
func (mstix MalwareAnalysisDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             mstix.ID,
		}
	)

	mNew, ok := newObj.(*mstixo.MalwareAnalysisDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(mstix.MalwareAnalysisDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*mNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("MalwareAnalysisDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("MalwareAnalysisDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- NoteDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (nstix NoteDomainObjectsSTIX) GetCurrentObject() interface{} {
	return nstix.NoteDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа NoteDomainObjectsSTIX
func (nstix NoteDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             nstix.ID,
		}
	)

	nNew, ok := newObj.(*mstixo.NoteDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(nstix.NoteDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*nNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("NoteDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("NoteDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- ObservedDataDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (ostix ObservedDataDomainObjectsSTIX) GetCurrentObject() interface{} {
	return ostix.ObservedDataDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ObservedDataDomainObjectsSTIX
func (ostix ObservedDataDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ostix.ID,
		}
	)

	oNew, ok := newObj.(*mstixo.ObservedDataDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(ostix.ObservedDataDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*oNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("ObservedDataDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("ObservedDataDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- OpinionDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (ostix OpinionDomainObjectsSTIX) GetCurrentObject() interface{} {
	return ostix.OpinionDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа OpinionDomainObjectsSTIX
func (ostix OpinionDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ostix.ID,
		}
	)

	oNew, ok := newObj.(*mstixo.OpinionDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(ostix.OpinionDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*oNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("OpinionDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("OpinionDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- ReportDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (rstix ReportDomainObjectsSTIX) GetCurrentObject() interface{} {
	return rstix.ReportDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ReportDomainObjectsSTIX
func (rstix ReportDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             rstix.ID,
		}
	)

	rNew, ok := newObj.(*mstixo.ReportDomainObjectsSTIX)
	if !ok {

		fmt.Println("Report func 'ComparisonTypeCommonFields', ERROR ---=== type conversion error ===---")

		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(rstix.ReportDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*rNew)
	typeOfNewValue := newValue.Type()

	fmt.Println("Report func 'ComparisonTypeCommonFields', reflect oldValue:", oldValue, " reflect newValue:", newValue)
	fmt.Println("Report func 'ComparisonTypeCommonFields', reflect typeOfOldValue:", typeOfOldValue, " reflect typeOfNewValue:", typeOfNewValue)

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {

			fmt.Println("Report func 'ComparisonTypeCommonFields', typeOfOldValue.Field(i).Name: ", typeOfOldValue.Field(i).Name, " typeOfNewValue.Field(j).Name:", typeOfNewValue.Field(j).Name)

			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("ReportDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("ReportDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- ThreatActorDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (tastix ThreatActorDomainObjectsSTIX) GetCurrentObject() interface{} {
	return tastix.ThreatActorDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ThreatActorDomainObjectsSTIX
func (tastix ThreatActorDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             tastix.ID,
		}
	)

	taNew, ok := newObj.(*mstixo.ThreatActorDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(tastix.ThreatActorDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*taNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("ThreatActorDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("ThreatActorDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- ToolDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (tstix ToolDomainObjectsSTIX) GetCurrentObject() interface{} {
	return tstix.ToolDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ToolDomainObjectsSTIX
func (tstix ToolDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             tstix.ID,
		}
	)

	tNew, ok := newObj.(*mstixo.ToolDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(tstix.ToolDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*tNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("ToolDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("ToolDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}

/* --- VulnerabilityDomainObjectsSTIX --- */

// GetCurrentObject возвращает текущий объект
func (vstix VulnerabilityDomainObjectsSTIX) GetCurrentObject() interface{} {
	return vstix.VulnerabilityDomainObjectsSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа VulnerabilityDomainObjectsSTIX
func (vstix VulnerabilityDomainObjectsSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             vstix.ID,
		}
	)

	vNew, ok := newObj.(*mstixo.VulnerabilityDomainObjectsSTIX)
	if !ok {
		return isEqual, cot, fmt.Errorf("type conversion error")
	}

	oldValue := reflect.ValueOf(vstix.VulnerabilityDomainObjectsSTIX)
	typeOfOldValue := oldValue.Type()

	newValue := reflect.ValueOf(*vNew)
	typeOfNewValue := newValue.Type()

	for i := 0; i < oldValue.NumField(); i++ {
		for j := 0; j < newValue.NumField(); j++ {
			if typeOfOldValue.Field(i).Name != typeOfNewValue.Field(j).Name {
				continue
			}

			if typeOfOldValue.Field(i).Name == "CommonPropertiesDomainObjectSTIX" {
				//привести значение к типу CommonPropertiesDomainObjectSTIX
				cpdoOld, ok := oldValue.Field(i).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				cpdoNew, ok := newValue.Field(j).Interface().(mstixo.CommonPropertiesDomainObjectSTIX)
				if !ok {
					return false, cot, fmt.Errorf("type conversion error")
				}

				ok, result := comparisonTypeCommonFields(cpdoOld, cpdoNew)
				if ok {
					continue
				}

				isEqual = false

				for _, v := range result {
					cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
						FeildType: v.FeildType,
						Path:      path.Join("VulnerabilityDomainObjectsSTIX", v.Path),
						Value:     v.Value,
					})
				}

				continue
			}

			if !reflect.DeepEqual(oldValue.Field(i).Interface(), newValue.Field(j).Interface()) {
				cot.FieldList = append(cot.FieldList, OldFieldValueObjectType{
					FeildType: typeOfOldValue.Field(i).Type.Name(),
					Path:      path.Join("VulnerabilityDomainObjectsSTIX", typeOfOldValue.Field(i).Name),
					Value:     oldValue.Field(i).Interface(),
				})

				isEqual = false
			}
		}
	}

	return isEqual, cot, nil
}
