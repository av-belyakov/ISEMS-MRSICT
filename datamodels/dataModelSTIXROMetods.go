package datamodels

import (
	"time"
)

/*****************************************************************************/
/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/
/*****************************************************************************/

/* --- RelationshipObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (rstix RelationshipObjectSTIX) GetCurrentObject() interface{} {
	return rstix.RelationshipObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа RelationshipObjectSTIX
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

	return isEqual, cot, nil
}

/* --- SightingObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (sstix SightingObjectSTIX) GetCurrentObject() interface{} {
	return sstix.SightingObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа SightingObjectSTIX
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

	return isEqual, cot, nil
}
