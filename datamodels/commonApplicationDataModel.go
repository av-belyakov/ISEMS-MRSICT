package datamodels

//ListSTIXObject может содержать любой из STIX объектов с указанием его типа
// DataType - тип STIX объекта
// Data - непосредственно сам STIX объект
type ListSTIXObject struct {
	DataType string
	Data     interface{}
}
