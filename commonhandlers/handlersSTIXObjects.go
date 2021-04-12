package commonhandlers

import (
	"ISEMS-MRSICT/datamodels"
)

//GetListIDFromListSTIXObjects возвращает список ID STIX объектов
func GetListIDFromListSTIXObjects(l []*datamodels.ElementSTIXObject) []string {
	list := make([]string, 0, len(l))

	for _, v := range l {
		list = append(list, v.Data.GetID())
	}

	return list
}
