package datamodels

import (
	"time"
)

/********** 			Описание типов для коллекции обеспечивающей логирование объектов STIX 			**********/

//LogChangesStoredValuesObjectSTIX содержит информацию с логами изменений произошедших в объектах STIX
// IDObjectSTIX - id объекта STIX
// DateModification - дата/время модификации объекта STIX
// NameSystemMadeChange - имя системы выполнившей изменение
// UserNameMadeChange - имя пользователя выполнившего изменение
// ListChanges - список изменений
type LogChangesStoredValuesObjectSTIX struct {
	IDObjectSTIX         string                                   `jsoh:"id_object_stix" bson:"id_object_stix"`
	DateModification     time.Time                                `jsoh:"date_modification" bson:"date_modification"`
	SystemNameMadeChange string                                   `jsoh:"name_system_made_change" bson:"name_system_made_change"`
	UserNameMadeChange   string                                   `json:"user_name_made_change" bson:"user_name_made_change"`
	ListChanges          []LogChangesStoredValuesChangeObjectSTIX `json:"list_changes" bson:"list_changes"`
}

//LogChangesStoredValuesChangeObjectSTIX содержит информацию по измененному объекту STIX
// ObjectNameSTIX - наименование объекта STIX подвергшегося изменению
// FieldName - поле объекта STIX подвергшееся изменению
// FieldType - тип данных поля, подвергшееся изменению. Значение этого поля будет использоватся для приведение типа, так как в свойствах
//  PreviousValue и CurrenValue может хранится разные значения. По этому для них приходится использовать значение []byte (вероятнее всего JSON).
// PreviousValue - предыдущее значение содержащееся в поле объекта STIX
// CurrenValue - текущее значение содержащееся в поле объекта STIX
type LogChangesStoredValuesChangeObjectSTIX struct {
	ObjectNameSTIX string `json:"object_name_stix" bson:"object_name_stix"`
	FieldName      string `json:"field_name" bson:"field_name"`
	FieldType      string `json:"field_type" bson:"field_type"`
	PreviousValue  []byte `json:"previous_value" bson:"previous_value"`
	CurrenValue    []byte `json:"current_value" bson:"current_value"`
}
