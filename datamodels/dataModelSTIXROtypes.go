package datamodels

import (
	"time"
)

/********** 			Relationship Objects STIX (ТИПЫ)			**********/

//OptionalCommonPropertiesRelationshipObjectSTIX общие, опциональные свойства для все объектов STIX типа Relationship Objects
// SpecVersion - версия STIX спецификации (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Modified - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type OptionalCommonPropertiesRelationshipObjectSTIX struct {
	SpecVersion string    `json:"spec_version" bson:"spec_version"`
	Created     time.Time `json:"created" bson:"created"`
	Modified    time.Time `json:"modified" bson:"modified"`
}

//RelationshipObjectSTIX объект "Relationship", по терминалогии STIX, используется для связывания двух Domain Object STIX (SDO) или Cyber-observable Objects STIX
//  (SCO), чтобы описать, как они связаны друг с другом. Если SDOS и SCOS считаются "узлами" или "вершинами" в графе, то Объекты отношений (SRO)
//  представляют собой "ребра".
// RelationshipType - содержит наименование, используемое для идентификации объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Description - содержит описание более подробной информации об объекте.
// SourceRef - устанавливает идентификатор исходного (исходящего) объекта. Значение ДОЛЖНО быть идентификатором ссылки на SDO или SCO (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// TargetRef - определяет идентификатор целевого (to) объекта. Значение ДОЛЖНО быть идентификатором ссылки на SDO или SCO(ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// StartTime - время, в формате "2016-05-12T08:17:27.000Z". Эта необязательная временная метка представляет собой самое раннее время, в которое существует
//  связь между объектами.
// StopTime - время, в формате "2016-05-12T08:17:27.000Z". Последнее время, в которое существует связь между объектами.
type RelationshipObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesRelationshipObjectSTIX
	RelationshipType string             `json:"relationship_type" bson:"relationship_type"`
	Description      string             `json:"description" bson:"description"`
	SourceRef        IdentifierTypeSTIX `json:"source_ref" bson:"source_ref"`
	TargetRef        IdentifierTypeSTIX `json:"target_ref" bson:"target_ref"`
	StartTime        time.Time          `json:"start_time" bson:"start_time"`
	StopTime         time.Time          `json:"stop_time" bson:"stop_time"`
}

//SightingObjectSTIX объект "Sighting", по терминалогии STIX, это особый тип SRO. Отношение, которое содержит дополнительные свойства, отсутствующие в объекте Relationship.
// Description - содержит более детальную информацию об объекте.
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z". Определяет начало временного окна, в течение которого был замечен SDO, на который ссылается свойство sighting_of_ref.
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z". Определяет конец временного окна, в течение которого был замечен SDO, на который ссылается свойство sighting_of_ref.
// Count - определяет количество раз, когда SDO, на которое ссылается свойство sighting_of_ref, был замечен.
// SightingOfRef - содержит ссылку на объект Domain Object STIX (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// ObservedDataRefs - содержит список идентификационных ссылок на наблюдаемые объекты данных, содержащие необработанные киберданные для данного наблюдения.
// WhereSightedRefs - содержит список идентификационных ссылок на объекты идентификации или местоположения, описывающие сущности или типы сущностей.
// Summary - содержит индикатор информаирующий о том, следует ли считеть данный объект совокупностью сводных данных. Сводные данные представляют собой совокупность предыдущих отчетов
//  о наблюдениях и не должны рассматриваться как первичные исходные данные.
type SightingObjectSTIX struct {
	CommonPropertiesObjectSTIX
	OptionalCommonPropertiesRelationshipObjectSTIX
	Description      string                `json:"description" bson:"description"`
	FirstSeen        time.Time             `json:"first_seen" bson:"first_seen"`
	LastSeen         time.Time             `json:"last_seen" bson:"last_seen"`
	Count            int                   `json:"count" bson:"count"`
	SightingOfRef    IdentifierTypeSTIX    `json:"sighting_of_ref" bson:"sighting_of_ref"`
	ObservedDataRefs []*IdentifierTypeSTIX `json:"observed_data_refs" bson:"observed_data_refs"`
	WhereSightedRefs []*IdentifierTypeSTIX `json:"where_sighted_refs" bson:"where_sighted_refs"`
	Summary          bool                  `json:"summary" bson:"summary"`
}
