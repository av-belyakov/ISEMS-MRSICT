package interactionmongodb

import (
	"fmt"
	"net"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	"github.com/asaskevich/govalidator"
	ipv4conv "github.com/signalsciences/ipv4"
	"go.mongodb.org/mongo-driver/bson"
)

//CreateSearchQueriesSTIXObject обработчик формирующий поисковые запросы для осуществления поиска по коллекции содержащей документы в формате STIX
func CreateSearchQueriesSTIXObject(sp *datamodels.SearchThroughCollectionSTIXObjectsType) bson.D {
	var (
		documentsType           bson.E
		dataTimeActionForObject bson.E
		createdByRef            bson.E
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

	sizessf := len(sp.SpecificSearchFields)

	if sizessf == 0 {
		return bson.D{
			documentsType,
			dataTimeActionForObject,
			createdByRef,
		}
	}

	//между всеми объектами sp.SpecificSearchFields применяется логика "ИЛИ"
	if sizessf == 1 {
		return bson.D{
			documentsType,
			dataTimeActionForObject,
			createdByRef,
			HandlerSpecificSearchFields(&sp.SpecificSearchFields[0]),
		}
	}
	var ssf bson.A
	for _, v := range sp.SpecificSearchFields {
		ssf = append(ssf, bson.D{HandlerSpecificSearchFields(&v)})
	}

	return bson.D{
		documentsType,
		dataTimeActionForObject,
		createdByRef,
		bson.E{Key: "$or", Value: ssf},
	}
}

//HandlerSpecificSearchFields обработчик поля "specific_search_fields"
func HandlerSpecificSearchFields(ssf *datamodels.SpecificSearchFieldsSTIXObjectType) bson.E {
	var (
		name    bson.D
		aliases bson.D
		seens   bson.D
		roles   bson.D
		country bson.D
		city    bson.D
		number  bson.D
		value   bson.D
	)

	timeFirstSeenIsExist := ssf.FirstSeen.Start.Unix() > 0 && ssf.FirstSeen.End.Unix() > 0
	timeLastSeenIsExist := ssf.LastSeen.Start.Unix() > 0 && ssf.LastSeen.End.Unix() > 0

	if timeFirstSeenIsExist && timeLastSeenIsExist {
		seens = bson.D{{Key: "$or", Value: bson.A{
			bson.D{{Key: "first_seen", Value: bson.D{
				{Key: "$gte", Value: ssf.FirstSeen.Start},
				{Key: "$lte", Value: ssf.FirstSeen.End},
			}}},
			bson.D{{Key: "last_seen", Value: bson.D{
				{Key: "$gte", Value: ssf.LastSeen.Start},
				{Key: "$lte", Value: ssf.LastSeen.End},
			}}},
		}}}
	} else if timeFirstSeenIsExist && !timeLastSeenIsExist {
		seens = bson.D{{Key: "first_seen", Value: bson.D{
			{Key: "$gte", Value: ssf.FirstSeen.Start},
			{Key: "$lte", Value: ssf.FirstSeen.End},
		}}}
	} else if !timeFirstSeenIsExist && timeLastSeenIsExist {
		seens = bson.D{{Key: "last_seen", Value: bson.D{
			{Key: "$gte", Value: ssf.LastSeen.Start},
			{Key: "$lte", Value: ssf.LastSeen.End},
		}}}
	}

	if ssf.Name != "" {
		name = bson.D{{Key: "name", Value: ssf.Name}}
	}

	if len(ssf.Aliases) > 0 {
		aliases = bson.D{{Key: "aliases", Value: bson.D{{Key: "$in", Value: ssf.Aliases}}}}
	}

	if len(ssf.Roles) > 0 {
		roles = bson.D{{Key: "roles", Value: bson.D{{Key: "$in", Value: ssf.Roles}}}}
	}

	if ssf.Country != "" {
		country = bson.D{{Key: "country", Value: ssf.Country}}
	}

	if ssf.City != "" {
		city = bson.D{{Key: "city", Value: ssf.City}}
	}

	if ssf.NumberAutonomousSystem > 0 {
		number = bson.D{{Key: "$eq", Value: bson.D{{Key: "number", Value: ssf.NumberAutonomousSystem}}}}
	}

	if len(ssf.Value) > 0 {
		value = HandlerValueField(ssf.Value)
	}

	return bson.E{
		Key: "$and",
		Value: bson.A{
			name,
			aliases,
			seens,
			roles,
			country,
			city,
			number,
			value,
		},
	}
}

//HandlerValueField обработка поля "value" которое может содержать любые из следующих типов значений: "domain-name", "email-addr", "ipv4-addr",
// "ipv6-addr" или "url"
func HandlerValueField(listValue []string) bson.D {
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

	fmt.Printf("func 'HandlerValueField', LIST URL: '%v'\n", listURL)
	fmt.Printf("func 'HandlerValueField', LIST IPv4 or IPv4Net: '%v'\n", listIPv4)
	fmt.Printf("func 'HandlerValueField', LIST ALL REMAINING: '%v'\n", listAllRemainingOnes)

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
		return bson.D{}
	}

	return bson.D{{Key: "$or", Value: tl}}
}
