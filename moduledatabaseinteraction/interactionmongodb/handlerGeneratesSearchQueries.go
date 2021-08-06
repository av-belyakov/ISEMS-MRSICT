package interactionmongodb

import (
	"fmt"
	"net"
	"reflect"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	"github.com/asaskevich/govalidator"
	ipv4conv "github.com/signalsciences/ipv4"
	"go.mongodb.org/mongo-driver/bson"
)

//CreateSearchQueriesSTIXObject обработчик формирующий поисковые запросы для осуществления поиска по коллекции содержащей документы в формате STIX
func CreateSearchQueriesSTIXObject(sp *datamodels.SearchThroughCollectionSTIXObjectsType) bson.D {
	var (
		documentsType                        bson.E
		dataTimeActionForObject              bson.E
		createdByRef                         bson.E
		outsideSpecificationSearchFields     bson.E
		outsideSpecificationSearchFieldsList bson.A
		ssf                                  bson.A
		listTypeName                         = map[string]string{
			"ComputerThreatType":          "computer_threat_type",
			"DecisionsMadeComputerThreat": "decisions_made_computer_threat",
		}
	)

	if len(sp.DocumentsID) > 0 {
		return bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: sp.DocumentsID}}}}
	}

	if len(sp.DocumentsType) > 0 {
		documentsType = bson.E{Key: "commonpropertiesobjectstix.type", Value: bson.D{{Key: "$in", Value: sp.DocumentsType}}}
	}

	timeCreateIsExist := sp.Created.Start.Unix() > 0 && sp.Created.End.Unix() > 0
	timeModifiedIsExist := sp.Modified.Start.Unix() > 0 && sp.Modified.End.Unix() > 0

	if timeCreateIsExist && timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: "commonpropertiesdomainobjectstix.created", Value: bson.D{
				{Key: "$gte", Value: sp.Created.Start},
				{Key: "$lte", Value: sp.Created.End},
			}}},
			bson.D{{Key: "commonpropertiesdomainobjectstix.modified", Value: bson.D{
				{Key: "$gte", Value: sp.Modified.Start},
				{Key: "$lte", Value: sp.Modified.End},
			}}},
		}}
	} else if !timeCreateIsExist && timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "commonpropertiesdomainobjectstix.modified", Value: bson.D{
			{Key: "$gte", Value: sp.Modified.Start},
			{Key: "$lte", Value: sp.Modified.End},
		}}
	} else if timeCreateIsExist && !timeModifiedIsExist {
		dataTimeActionForObject = bson.E{Key: "commonpropertiesdomainobjectstix.created", Value: bson.D{
			{Key: "$gte", Value: sp.Created.Start},
			{Key: "$lte", Value: sp.Created.End},
		}}
	}

	if sp.CreatedByRef != "" {
		createdByRef = bson.E{Key: "commonpropertiesdomainobjectstix.created_by_ref", Value: sp.CreatedByRef}
	}

	ossfValue := reflect.ValueOf(sp.OutsideSpecificationSearchFields)
	ossfType := ossfValue.Type()
	for i := 0; i < ossfType.NumField(); i++ {
		switch ossfType.Field(i).Type.Kind() {
		case reflect.String:
			if ossfValue.Field(i).String() != "" {
				outsideSpecificationSearchFields = bson.E{Key: fmt.Sprintf("outside_specification.%s", listTypeName[ossfType.Field(i).Name]), Value: ossfValue.Field(i).String()}
				outsideSpecificationSearchFieldsList = append(outsideSpecificationSearchFieldsList, bson.D{outsideSpecificationSearchFields})
			}

			//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		}
	}

	countossfl := len(outsideSpecificationSearchFieldsList)
	if countossfl > 1 {
		outsideSpecificationSearchFields = bson.E{Key: "$or", Value: outsideSpecificationSearchFieldsList}
	}

	if len(sp.SpecificSearchFields) == 0 {
		return bson.D{
			documentsType,
			dataTimeActionForObject,
			createdByRef,
			outsideSpecificationSearchFields,
		}
	}

	for _, v := range sp.SpecificSearchFields {
		ssf = append(ssf, HandlerSpecificSearchFields(sp.DocumentsType, &v))
	}

	return bson.D{
		documentsType,
		dataTimeActionForObject,
		createdByRef,
		outsideSpecificationSearchFields,
		bson.E{Key: "$or", Value: ssf},
	}
}

//HandlerSpecificSearchFields обработчик поля "specific_search_fields"
func HandlerSpecificSearchFields(ldt []string, ssf *datamodels.SpecificSearchFieldsSTIXObjectType) bson.D {
	var (
		name      bson.E
		aliases   bson.E
		seens     bson.E
		published bson.E
		roles     bson.E
		country   bson.E
		city      bson.E
		number    bson.E
		value     bson.E
	)

	timeFirstSeenIsExist := ssf.FirstSeen.Start.Unix() > 0 && ssf.FirstSeen.End.Unix() > 0
	timeLastSeenIsExist := ssf.LastSeen.Start.Unix() > 0 && ssf.LastSeen.End.Unix() > 0

	if timeFirstSeenIsExist && timeLastSeenIsExist {
		seens = bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: "first_seen", Value: bson.D{
				{Key: "$gte", Value: ssf.FirstSeen.Start},
				{Key: "$lte", Value: ssf.FirstSeen.End},
			}}},
			bson.D{{Key: "last_seen", Value: bson.D{
				{Key: "$gte", Value: ssf.LastSeen.Start},
				{Key: "$lte", Value: ssf.LastSeen.End},
			}}},
		}}
	} else if timeFirstSeenIsExist && !timeLastSeenIsExist {
		seens = bson.E{Key: "first_seen", Value: bson.D{
			{Key: "$gte", Value: ssf.FirstSeen.Start},
			{Key: "$lte", Value: ssf.FirstSeen.End},
		}}
	} else if !timeFirstSeenIsExist && timeLastSeenIsExist {
		seens = bson.E{Key: "last_seen", Value: bson.D{
			{Key: "$gte", Value: ssf.LastSeen.Start},
			{Key: "$lte", Value: ssf.LastSeen.End},
		}}
	}

	/*
		параметр Published есть только в объекте STIX DO Report и отвечает за 'закрытие' объекта Report, с помощью данного параметра будет
		осуществлятся поиск по 'открытым' и 'закрытым' отчетам.
		Поиск по полю 'published' будет выполнятся только если поле 'documents_type' содержит один элемент, а тип STIX DO будет равен "report".
		В остальных случаях поиск по полю 'published' будет игнорироватся
	*/
	if len(ldt) == 1 && ldt[0] == "report" {
		if ssf.Published.Unix() > 0 {
			published = bson.E{Key: "published", Value: bson.D{{Key: "$gt", Value: ssf.Published}}}
		} else {
			published = bson.E{Key: "published", Value: bson.D{{Key: "$lte", Value: ssf.Published}}}
		}
	}

	if ssf.Name != "" {
		name = bson.E{Key: "name", Value: ssf.Name}
	}

	if len(ssf.Aliases) > 0 {
		aliases = bson.E{Key: "aliases", Value: bson.D{{Key: "$in", Value: ssf.Aliases}}}
	}

	if len(ssf.Roles) > 0 {
		roles = bson.E{Key: "roles", Value: bson.D{{Key: "$in", Value: ssf.Roles}}}
	}

	if ssf.Country != "" {
		country = bson.E{Key: "country", Value: ssf.Country}
	}

	if ssf.City != "" {
		city = bson.E{Key: "city", Value: ssf.City}
	}

	if ssf.NumberAutonomousSystem > 0 {
		number = bson.E{Key: "$eq", Value: bson.D{{Key: "number", Value: ssf.NumberAutonomousSystem}}}
	}

	if len(ssf.Value) > 0 {
		value = HandlerValueField(ssf.Value)
	}

	return bson.D{
		name,
		aliases,
		seens,
		published,
		roles,
		country,
		city,
		number,
		value,
	}
}

//HandlerValueField обработка поля "value" которое может содержать любые из следующих типов значений: "domain-name", "email-addr", "ipv4-addr",
// "ipv6-addr" или "url"
func HandlerValueField(listValue []string) bson.E {
	var (
		tl                            bson.A
		listURL, listAllRemainingOnes = []string{}, []string{}
		listIPv4                      = []struct {
			start uint32
			end   uint32
		}{}
	)

	//так как поле "Value" может содержать любой из типов значений: "domain-name", "email-addr", "ipv4-addr", "ipv6-addr" или "url"
	// необходимо отделить тип "url" и "ipv4-addr" от всех остальных типов
	for _, v := range listValue {
		if commonlibs.IsIPv4Address(v) {
			if ipInt, err := commonlibs.Ip2long(v); err == nil {
				listIPv4 = append(listIPv4, struct {
					start uint32
					end   uint32
				}{ipInt, ipInt})
			}
		} else if commonlibs.IsComputerNetAddrIPv4Range(v) {
			if hostMin, hostMax, err := ipv4conv.CIDR2Range(v); err == nil {
				min, _ := commonlibs.Ip2long(hostMin)
				max, _ := commonlibs.Ip2long(hostMax)

				listIPv4 = append(listIPv4, struct {
					start uint32
					end   uint32
				}{min, max})
			}
		} else if govalidator.IsURL(v) {
			listURL = append(listURL, v)
		} else {
			ipv6 := net.ParseIP(v)
			if ipv6 == nil {
				listAllRemainingOnes = append(listAllRemainingOnes, v)

				continue
			}

			listAllRemainingOnes = append(listAllRemainingOnes, ipv6.To16().String())
		}
	}

	sizeListURL := len(listURL)
	sizeListIPv4 := len(listIPv4)
	sizeListAll := len(listAllRemainingOnes)

	//обработка только URL
	if sizeListURL > 0 {
		for _, v := range listURL {
			tl = append(tl, bson.A{bson.D{{Key: "url", Value: v}}, bson.D{{Key: "value", Value: v}}}...)
		}
	}

	//обработка только IPv4 или диапазонов IPv4
	if sizeListIPv4 > 0 {
		for _, v := range listIPv4 {
			tl = append(tl, bson.D{{Key: "$and", Value: bson.A{
				bson.D{{Key: "host_min", Value: bson.D{
					{Key: "$gte", Value: v.start}}}},
				bson.D{{Key: "host_max", Value: bson.D{
					{Key: "$lte", Value: v.end}}}},
			}}})
		}
	}

	//обработка остальных значений
	if sizeListAll > 0 {
		for _, v := range listAllRemainingOnes {
			tl = append(tl, bson.D{{Key: "value", Value: v}})
		}
	}

	if len(tl) == 0 {
		return bson.E{}
	}

	return bson.E{Key: "$or", Value: tl}
}
