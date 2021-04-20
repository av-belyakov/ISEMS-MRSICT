package interactionmongodb

import (
	"context"

	"ISEMS-MRSICT/datamodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//ComparasionListTypeSTIXObject содержит два списка STIX объектов, предназначенных для сравнения
type ComparasionListTypeSTIXObject struct {
	CollectionType   string
	OldList, NewList []*datamodels.ElementSTIXObject
}

//ComparasionListSTIXObject выполняет сравнение двух списков STIX объектов, cписка STIX объектов, полученных из БД и принятых от клиента API
func ComparasionListSTIXObject(clt ComparasionListTypeSTIXObject) []datamodels.DifferentObjectType {
	var (
		listDifferentResult []datamodels.DifferentObjectType
		dot                 datamodels.DifferentObjectType
		err                 error
		isEqual             = true
	)

	for _, vo := range clt.OldList {
		for _, vn := range clt.NewList {
			if vo.DataType != vn.DataType {
				continue
			}

			switch vn.DataType {
			//только для Domain Objects STIX
			case "attack-pattern":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetAttackPatternDomainObjectsSTIX(), clt.CollectionType)
			case "campaign":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetCampaignDomainObjectsSTIX(), clt.CollectionType)
			case "course-of-action":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetCourseOfActionDomainObjectsSTIX(), clt.CollectionType)
			case "grouping":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetGroupingDomainObjectsSTIX(), clt.CollectionType)
			case "identity":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIdentityDomainObjectsSTIX(), clt.CollectionType)
			case "indicator":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIndicatorDomainObjectsSTIX(), clt.CollectionType)
			case "infrastructure":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetInfrastructureDomainObjectsSTIX(), clt.CollectionType)
			case "intrusion-set":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetIntrusionSetDomainObjectsSTIX(), clt.CollectionType)
			case "location":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetLocationDomainObjectsSTIX(), clt.CollectionType)
			case "malware":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetMalwareDomainObjectsSTIX(), clt.CollectionType)
			case "malware-analysis":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetMalwareAnalysisDomainObjectsSTIX(), clt.CollectionType)
			case "note":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetNoteDomainObjectsSTIX(), clt.CollectionType)
			case "observed-data":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetObservedDataDomainObjectsSTIX(), clt.CollectionType)
			case "opinion":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetOpinionDomainObjectsSTIX(), clt.CollectionType)
			case "report":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetReportDomainObjectsSTIX(), clt.CollectionType)
			case "threat-actor":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetThreatActorDomainObjectsSTIX(), clt.CollectionType)
			case "tool":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetToolDomainObjectsSTIX(), clt.CollectionType)
			case "vulnerability":
				isEqual, dot, err = vo.Data.ComparisonTypeCommonFields(vn.GetVulnerabilityDomainObjectsSTIX(), clt.CollectionType)
			}

			if err != nil {
				continue
			}

			if isEqual {
				continue
			}

			listDifferentResult = append(listDifferentResult, dot)
		}
	}

	return listDifferentResult
}

type definingTypeSTIXObject struct {
	datamodels.CommonPropertiesObjectSTIX
}

//GetListElementSTIXObject возвращает, из БД, список STIX объектов
func GetListElementSTIXObject(cur *mongo.Cursor) []*datamodels.ElementSTIXObject {
	elements := []*datamodels.ElementSTIXObject{}

	for cur.Next(context.Background()) {
		var modelType definingTypeSTIXObject
		if err := cur.Decode(&modelType); err != nil {
			continue
		}

		switch modelType.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			tmpObj := datamodels.AttackPatternDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})
		case "campaign":
			tmpObj := datamodels.CampaignDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "course-of-action":
			tmpObj := datamodels.CourseOfActionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "grouping":
			tmpObj := datamodels.GroupingDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "identity":
			tmpObj := datamodels.IdentityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "indicator":
			tmpObj := datamodels.IndicatorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "infrastructure":
			tmpObj := datamodels.InfrastructureDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "intrusion-set":
			tmpObj := datamodels.IntrusionSetDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "location":
			tmpObj := datamodels.LocationDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "malware":
			tmpObj := datamodels.MalwareDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "malware-analysis":
			tmpObj := datamodels.MalwareAnalysisDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "note":
			tmpObj := datamodels.NoteDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "observed-data":
			tmpObj := datamodels.ObservedDataDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "opinion":
			tmpObj := datamodels.OpinionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "report":
			tmpObj := datamodels.ReportDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "threat-actor":
			tmpObj := datamodels.ThreatActorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "tool":
			tmpObj := datamodels.ToolDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "vulnerability":
			tmpObj := datamodels.VulnerabilityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		/* *** Relationship Objects *** */
		case "relationship":
			tmpObj := datamodels.RelationshipObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "sighting":
			tmpObj := datamodels.SightingObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			tmpObj := datamodels.ArtifactCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "autonomous-system":
			tmpObj := datamodels.AutonomousSystemCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "directory":
			tmpObj := datamodels.DirectoryCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "domain-name":
			tmpObj := datamodels.DomainNameCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "email-addr":
			tmpObj := datamodels.EmailAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "email-message":
			tmpObj := datamodels.EmailMessageCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "file":
			tmpObj := datamodels.FileCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "ipv4-addr":
			tmpObj := datamodels.IPv4AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "ipv6-addr":
			tmpObj := datamodels.IPv6AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "mac-addr":
			tmpObj := datamodels.MACAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "mutex":
			tmpObj := datamodels.MutexCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "network-traffic":
			tmpObj := datamodels.NetworkTrafficCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "process":
			tmpObj := datamodels.ProcessCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "software":
			tmpObj := datamodels.SoftwareCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "url":
			tmpObj := datamodels.URLCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "user-account":
			tmpObj := datamodels.UserAccountCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "windows-registry-key":
			tmpObj := datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})

		case "x509-certificate":
			tmpObj := datamodels.X509CertificateCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     tmpObj,
			})
		}
	}

	return elements
}

//FindSTIXObjectByID выполняет поиск в БД, STIX объектов по их ID и возвращает список STIX объектов типа datamodels.ElementSTIXObject
func FindSTIXObjectByID(qp QueryParameters, listID []string) ([]*datamodels.ElementSTIXObject, error) {
	var objID primitive.A

	for _, v := range listID {
		objID = append(objID, v)
	}

	cur, err := qp.Find((bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: objID}}}}))
	if err != nil {
		return nil, err
	}

	lr := GetListElementSTIXObject(cur)

	return lr, nil
}

//ReplacementElementsSTIXObject выполняет замену в БД, списка STIX объектов или добовляет новые объекты если их нет в БД
func ReplacementElementsSTIXObject(qp QueryParameters, l []*datamodels.ElementSTIXObject) error {
	listSize := len(l)
	listObj := make([]interface{}, 0, listSize)
	reqDeleteID := primitive.A{}

	for _, v := range l {
		reqDeleteID = append(reqDeleteID, v.Data.GetID())
		listObj = append(listObj, v.Data)
	}

	_, err := qp.DeleteManyData(bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: reqDeleteID}}}})
	if err != nil {
		return err
	}

	_, err = qp.InsertData(listObj)
	if err != nil {
		return err
	}

	return nil
}
