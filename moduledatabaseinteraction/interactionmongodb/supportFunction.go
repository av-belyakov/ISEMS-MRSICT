package interactionmongodb

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"time"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"

	"github.com/asaskevich/govalidator"
	mstixo "github.com/av-belyakov/methodstixobjects"
	"github.com/google/uuid"
	ipv4conv "github.com/signalsciences/ipv4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	listTypeSTIXObject = []string{
		"grouping",
		"note",
		"observed-data",
		"opinion",
		"report",
	}
)

// ComparasionListTypeSTIXObject содержит два списка STIX объектов, предназначенных для сравнения
type ComparasionListTypeSTIXObject struct {
	CollectionType   string
	OldList, NewList []*datamodels.ElementSTIXObject
}

// ComparasionListSTIXObject выполняет сравнение двух списков STIX объектов, cписка STIX объектов, полученных из БД и принятых от клиента API
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

// SavingAdditionalNameListSTIXObject сохранение дополнительного наименования в некоторых STIX объектах, имеющих свойства не входящие
// в основную спецификацию STIX 2.0
func SavingAdditionalNameListSTIXObject(currentList, addedList []*datamodels.ElementSTIXObject) []*datamodels.ElementSTIXObject {
	fmt.Println("func 'SavingAdditionalNameListSTIXObject', START...")

	for k, vadd := range addedList {
		if vadd.DataType != "report" {
			continue
		}

		reportAdd, ok := vadd.Data.(datamodels.ReportDomainObjectsSTIX)
		if !ok {
			continue
		}

	DONE:
		for _, vcurrent := range currentList {
			switch vcurrent.DataType {
			case "report":
				if vadd.Data.GetID() != vcurrent.Data.GetID() {
					continue
				}

				reportCurrent, ok := vcurrent.Data.(datamodels.ReportDomainObjectsSTIX)
				if !ok {
					break DONE
				}

				reportAdd.OutsideSpecification.AdditionalName = reportCurrent.OutsideSpecification.AdditionalName

				addedList[k] = &datamodels.ElementSTIXObject{
					DataType: reportAdd.GetType(),
					Data:     reportAdd,
				}
			}
		}
	}

	return addedList
}

type definingTypeSTIXObject struct {
	mstixo.CommonPropertiesObjectSTIX
}

// GetListElementSTIXObject возвращает, из БД, список STIX объектов
func GetListElementSTIXObject(cur *mongo.Cursor) []*datamodels.ElementSTIXObject {
	elements := []*datamodels.ElementSTIXObject{}

	//fmt.Println("func 'GetListElementSTIXObject', START...")

	for cur.Next(context.Background()) {
		var modelType definingTypeSTIXObject
		if err := cur.Decode(&modelType); err != nil {

			//fmt.Println("func 'GetListElementSTIXObject', cur.Decode(&modelType), CONTINUE")

			continue
		}

		//fmt.Println("func 'GetListElementSTIXObject', modelType.Type:", modelType)

		switch modelType.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			tmpObj := mstixo.AttackPatternDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.AttackPatternDomainObjectsSTIX{AttackPatternDomainObjectsSTIX: tmpObj},
			})
		case "campaign":
			tmpObj := mstixo.CampaignDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.CampaignDomainObjectsSTIX{CampaignDomainObjectsSTIX: tmpObj},
			})

		case "course-of-action":
			tmpObj := mstixo.CourseOfActionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.CourseOfActionDomainObjectsSTIX{CourseOfActionDomainObjectsSTIX: tmpObj},
			})

		case "grouping":
			tmpObj := mstixo.GroupingDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.GroupingDomainObjectsSTIX{GroupingDomainObjectsSTIX: tmpObj},
			})

		case "identity":
			tmpObj := mstixo.IdentityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IdentityDomainObjectsSTIX{IdentityDomainObjectsSTIX: tmpObj},
			})

		case "indicator":
			tmpObj := mstixo.IndicatorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IndicatorDomainObjectsSTIX{IndicatorDomainObjectsSTIX: tmpObj},
			})

		case "infrastructure":
			tmpObj := mstixo.InfrastructureDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.InfrastructureDomainObjectsSTIX{InfrastructureDomainObjectsSTIX: tmpObj},
			})

		case "intrusion-set":
			tmpObj := mstixo.IntrusionSetDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IntrusionSetDomainObjectsSTIX{IntrusionSetDomainObjectsSTIX: tmpObj},
			})

		case "location":
			tmpObj := mstixo.LocationDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.LocationDomainObjectsSTIX{LocationDomainObjectsSTIX: tmpObj},
			})

		case "malware":
			tmpObj := mstixo.MalwareDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MalwareDomainObjectsSTIX{MalwareDomainObjectsSTIX: tmpObj},
			})

		case "malware-analysis":
			tmpObj := mstixo.MalwareAnalysisDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MalwareAnalysisDomainObjectsSTIX{MalwareAnalysisDomainObjectsSTIX: tmpObj},
			})

		case "note":
			tmpObj := mstixo.NoteDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.NoteDomainObjectsSTIX{NoteDomainObjectsSTIX: tmpObj},
			})

		case "observed-data":
			tmpObj := mstixo.ObservedDataDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ObservedDataDomainObjectsSTIX{ObservedDataDomainObjectsSTIX: tmpObj},
			})

		case "opinion":
			tmpObj := mstixo.OpinionDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.OpinionDomainObjectsSTIX{OpinionDomainObjectsSTIX: tmpObj},
			})

		case "report":
			tmpObj := mstixo.ReportDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ReportDomainObjectsSTIX{ReportDomainObjectsSTIX: tmpObj},
			})

		case "threat-actor":
			tmpObj := mstixo.ThreatActorDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ThreatActorDomainObjectsSTIX{ThreatActorDomainObjectsSTIX: tmpObj},
			})

		case "tool":
			tmpObj := mstixo.ToolDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ToolDomainObjectsSTIX{ToolDomainObjectsSTIX: tmpObj},
			})

		case "vulnerability":
			tmpObj := mstixo.VulnerabilityDomainObjectsSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.VulnerabilityDomainObjectsSTIX{VulnerabilityDomainObjectsSTIX: tmpObj},
			})

		/* *** Relationship Objects *** */
		case "relationship":
			tmpObj := mstixo.RelationshipObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.RelationshipObjectSTIX{RelationshipObjectSTIX: tmpObj},
			})

		case "sighting":
			tmpObj := mstixo.SightingObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.SightingObjectSTIX{SightingObjectSTIX: tmpObj},
			})

		/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			tmpObj := mstixo.ArtifactCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ArtifactCyberObservableObjectSTIX{ArtifactCyberObservableObjectSTIX: tmpObj},
			})

		case "autonomous-system":
			tmpObj := mstixo.AutonomousSystemCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.AutonomousSystemCyberObservableObjectSTIX{AutonomousSystemCyberObservableObjectSTIX: tmpObj},
			})

		case "directory":
			tmpObj := mstixo.DirectoryCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.DirectoryCyberObservableObjectSTIX{DirectoryCyberObservableObjectSTIX: tmpObj},
			})

		case "domain-name":
			tmpObj := mstixo.DomainNameCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.DomainNameCyberObservableObjectSTIX{DomainNameCyberObservableObjectSTIX: tmpObj},
			})

		case "email-addr":
			tmpObj := mstixo.EmailAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.EmailAddressCyberObservableObjectSTIX{EmailAddressCyberObservableObjectSTIX: tmpObj},
			})

		case "email-message":
			tmpObj := mstixo.EmailMessageCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.EmailMessageCyberObservableObjectSTIX{EmailMessageCyberObservableObjectSTIX: tmpObj},
			})

		case "file":
			tmpObj := mstixo.FileCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.FileCyberObservableObjectSTIX{FileCyberObservableObjectSTIX: tmpObj},
			})

		case "ipv4-addr":
			tmpObj := datamodels.IPv4AddressCyberObservableSimilarObjectSTIX{}
			//tmpObj := datamodels.IPv4AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			/*resolvesToRefs := make([]mstixo.IdentifierTypeSTIX, 0, len(tmpObj.ResolvesToRefs))
			for _, v := range tmpObj.ResolvesToRefs {
				resolvesToRefs = append(resolvesToRefs, v)
			}*/

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data: datamodels.IPv4AddressCyberObservableObjectSTIX{
					IPv4AddressCyberObservableObjectSTIX: mstixo.IPv4AddressCyberObservableObjectSTIX{
						CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
							Type: tmpObj.CommonPropertiesObjectSTIX.Type,
							ID:   tmpObj.CommonPropertiesObjectSTIX.ID,
						},
						OptionalCommonPropertiesCyberObservableObjectSTIX: mstixo.OptionalCommonPropertiesCyberObservableObjectSTIX{
							SpecVersion:       tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.SpecVersion,
							ObjectMarkingRefs: tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.ObjectMarkingRefs,
							GranularMarkings:  tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.GranularMarkings,
							Defanged:          tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.Defanged,
						},
						Value:          tmpObj.Value,
						ResolvesToRefs: tmpObj.ResolvesToRefs,
						BelongsToRefs:  tmpObj.BelongsToRefs,
					},
				},
			})

			/*
				elements = append(elements, &datamodels.ElementSTIXObject{
					DataType: modelType.Type,
					Data: datamodels.IPv4AddressCyberObservableObjectSTIX{
						mstixo.IPv4AddressCyberObservableObjectSTIX{
							CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
								Type: tmpObj.CommonPropertiesObjectSTIX.Type,
								ID:   tmpObj.CommonPropertiesObjectSTIX.ID,
							},
							OptionalCommonPropertiesCyberObservableObjectSTIX: mstixo.OptionalCommonPropertiesCyberObservableObjectSTIX{
								SpecVersion:       tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.SpecVersion,
								ObjectMarkingRefs: tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.ObjectMarkingRefs,
								GranularMarkings:  tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.GranularMarkings,
								Defanged:          tmpObj.OptionalCommonPropertiesCyberObservableObjectSTIX.Defanged,
							},
							Value:          tmpObj.Value,
							ResolvesToRefs: tmpObj.ResolvesToRefs,
							BelongsToRefs:  tmpObj.BelongsToRefs,
						},
					},
				})
			*/

		case "ipv6-addr":
			tmpObj := mstixo.IPv6AddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.IPv6AddressCyberObservableObjectSTIX{IPv6AddressCyberObservableObjectSTIX: tmpObj},
			})

		case "mac-addr":
			tmpObj := mstixo.MACAddressCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MACAddressCyberObservableObjectSTIX{MACAddressCyberObservableObjectSTIX: tmpObj},
			})

		case "mutex":
			tmpObj := mstixo.MutexCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.MutexCyberObservableObjectSTIX{MutexCyberObservableObjectSTIX: tmpObj},
			})

		case "network-traffic":
			tmpObj := mstixo.NetworkTrafficCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.NetworkTrafficCyberObservableObjectSTIX{NetworkTrafficCyberObservableObjectSTIX: tmpObj},
			})

		case "process":
			tmpObj := mstixo.ProcessCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.ProcessCyberObservableObjectSTIX{ProcessCyberObservableObjectSTIX: tmpObj},
			})

		case "software":
			tmpObj := mstixo.SoftwareCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.SoftwareCyberObservableObjectSTIX{SoftwareCyberObservableObjectSTIX: tmpObj},
			})

		case "url":
			tmpObj := mstixo.URLCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.URLCyberObservableObjectSTIX{URLCyberObservableObjectSTIX: tmpObj},
			})

		case "user-account":
			tmpObj := mstixo.UserAccountCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.UserAccountCyberObservableObjectSTIX{UserAccountCyberObservableObjectSTIX: tmpObj},
			})

		case "windows-registry-key":
			tmpObj := mstixo.WindowsRegistryKeyCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{WindowsRegistryKeyCyberObservableObjectSTIX: tmpObj},
			})

		case "x509-certificate":
			tmpObj := mstixo.X509CertificateCyberObservableObjectSTIX{}
			err := cur.Decode(&tmpObj)
			if err != nil {
				break
			}

			elements = append(elements, &datamodels.ElementSTIXObject{
				DataType: modelType.Type,
				Data:     datamodels.X509CertificateCyberObservableObjectSTIX{X509CertificateCyberObservableObjectSTIX: tmpObj},
			})
		}
	}

	//fmt.Println("func 'GetListElementSTIXObject' STOP elements ", elements)

	return elements
}

// FindSTIXObjectByID выполняет поиск в БД, STIX объектов по их ID и возвращает список STIX объектов типа datamodels.ElementSTIXObject
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

// ReplacementElementsSTIXObject выполняет замену в БД, списка STIX объектов или добовляет новые объекты если их нет в БД
func ReplacementElementsSTIXObject(qp QueryParameters, l []*datamodels.ElementSTIXObject) error {
	listSize := len(l)
	listObj := make([]interface{}, 0, listSize)
	reqDeleteID := primitive.A{}

	for _, v := range l {
		var hmax, hmin uint32
		reqDeleteID = append(reqDeleteID, v.Data.GetID())

		//добавляем поля HostMin и HostMax с цифровым минимальным и максимальным значением IPv4 (это нужно для быстрого поиска в БД)
		if v.Data.GetType() == "ipv4-addr" {
			ipv4addr, ok := v.Data.(datamodels.IPv4AddressCyberObservableObjectSTIX)
			if !ok {
				continue
			}

			if hostMin, hostMax, err := ipv4conv.CIDR2Range(ipv4addr.Value); err == nil {
				hmax, _ = commonlibs.Ip2long(hostMax)
				hmin, _ = commonlibs.Ip2long(hostMin)
			} else {
				hmax, _ = commonlibs.Ip2long(ipv4addr.Value)
				hmin, _ = commonlibs.Ip2long(ipv4addr.Value)
			}

			listObj = append(listObj, datamodels.IPv4AddressCyberObservableSimilarObjectSTIX{
				CommonPropertiesObjectSTIX:                        ipv4addr.CommonPropertiesObjectSTIX,
				OptionalCommonPropertiesCyberObservableObjectSTIX: ipv4addr.OptionalCommonPropertiesCyberObservableObjectSTIX,
				HostMin:        hmin,
				HostMax:        hmax,
				Value:          ipv4addr.Value,
				ResolvesToRefs: ipv4addr.ResolvesToRefs,
				BelongsToRefs:  ipv4addr.BelongsToRefs,
			})

			continue
		}

		//убираем "0000" из актетов IPv6, например было "2001:0db8:85a3:0000:0000:8a2e:0370:7334", стало "2001:0db8:85a3::8a2e:0370:7334"
		if v.Data.GetType() == "ipv6-addr" {
			ipv6addr, ok := v.Data.(datamodels.IPv6AddressCyberObservableObjectSTIX)
			if !ok {
				continue
			}

			var ip = ipv6addr.Value

			if ipv6Addr, _, err := net.ParseCIDR(ipv6addr.Value); err == nil {
				if !govalidator.IsIPv6(ipv6Addr.String()) {
					continue
				}
			} else {
				ipv6 := net.ParseIP(ipv6addr.Value)
				if ipv6 == nil {
					continue
				}

				ip = ipv6.To16().String()
			}

			listObj = append(listObj, mstixo.IPv6AddressCyberObservableObjectSTIX{
				CommonPropertiesObjectSTIX:                        ipv6addr.CommonPropertiesObjectSTIX,
				OptionalCommonPropertiesCyberObservableObjectSTIX: ipv6addr.OptionalCommonPropertiesCyberObservableObjectSTIX,
				Value:          ip,
				ResolvesToRefs: ipv6addr.ResolvesToRefs,
				BelongsToRefs:  ipv6addr.BelongsToRefs,
			})

			continue
		}

		listObj = append(listObj, v.Data.GetCurrentObject())

		/*if v.DataType == "report" {
			fmt.Printf("func 'ReplacementElementsSTIXObject', DataType:%s, v.Data.GetCurrentObject():%v \n", v.DataType, v.Data.GetCurrentObject())
		}*/
	}

	_, err := qp.DeleteManyData(bson.D{{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: reqDeleteID}}}}, options.Delete())
	if err != nil {
		return err
	}

	_, err = qp.InsertData(listObj, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "commonpropertiesobjectstix.type", Value: 1},
				{Key: "commonpropertiesobjectstix.id", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source_ref", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

// FindRBObjectByName выполняет поиск в БД, Reference Book объектов по их ID и возвращает список объектов типа datamodels.Vocabulary - справочник
func FindRBObjectsByNames(qp QueryParameters, listNames []string) (datamodels.Vocabularys, error) {
	//var objID primitive.A
	//for _, v := range listNames {
	//	obj
	//}
	return nil, nil
}

// FilterEditabelRB - функция проверки и фильтрации объектов RB на то что они являются редактируемыми
func FilterEditabelRB(listRB datamodels.Vocabularys) (datamodels.Vocabularys, datamodels.Vocabularys) {
	var (
		listNotEditable datamodels.Vocabularys
		listEditable    datamodels.Vocabularys
	)

	for _, v := range listRB {
		if regexp.MustCompile("-enum$").MatchString(v.Name) {
			listNotEditable = append(listNotEditable, v)
		}
		if regexp.MustCompile("-ov$").MatchString(v.Name) {
			listEditable = append(listEditable, v)
		}
	}
	listRB = listEditable
	return listEditable, listNotEditable
}

// ComparasionListRBbject - функция поэлементного сравнения вдух списков RB-объектов
func ComparasionListRBbject(compList1 datamodels.Vocabularys, compList2 datamodels.Vocabularys) []datamodels.DifferentObjectType {
	return nil
}

// GetIDGroupingObjectSTIX проверяет наличие Grouping STIX DO объектов с заданными именами и при необходимости создает их. Возвращает список
// идентификаторов STIX DO объектов типа Grouping и название объекта.
func GetIDGroupingObjectSTIX(qp QueryParameters, listSearch map[string]datamodels.StorageApplicationCommonListType) (map[string]datamodels.StorageApplicationCommonListType, error) {
	var (
		isTrue     bool
		ls         []string
		listInsert []interface{}
	)
	listID := map[string]datamodels.StorageApplicationCommonListType{}

	for k := range listSearch {
		ls = append(ls, k)
	}

	//получить все найденные документы
	cur, err := qp.Find(bson.D{{Key: "name", Value: bson.D{{Key: "$in", Value: ls}}}})
	if err != nil {
		return listID, err
	}

	listTypeStatus := GetListGroupingObjectSTIX(cur)

	for ko, vo := range listSearch {
		for _, vt := range listTypeStatus {
			if ko == vt.Name {
				listID[ko] = datamodels.StorageApplicationCommonListType{
					ID:          vt.ID,
					Description: vo.Description,
				}

				isTrue = true

				continue
			}
		}

		context := mstixo.OpenVocabTypeSTIX("suspicious-activity")
		if !isTrue {
			id := uuid.NewString()
			if (ko == "successfully implemented computer threat") || (ko == "unsuccessfully computer threat") || (ko == "false positive") {
				context = mstixo.OpenVocabTypeSTIX(ko)
			}

			listInsert = append(listInsert, datamodels.GroupingDomainObjectsSTIX{
				GroupingDomainObjectsSTIX: mstixo.GroupingDomainObjectsSTIX{
					CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
						Type: "grouping",
						ID:   fmt.Sprintf("grouping--%s", id),
					},
					CommonPropertiesDomainObjectSTIX: mstixo.CommonPropertiesDomainObjectSTIX{
						SpecVersion: "2.1",
						Created:     time.Now(),
					},
					Name:        ko,
					Description: vo.Description,
					Context:     context,
					ObjectRefs:  []mstixo.IdentifierTypeSTIX{},
				},
			})

			listID[ko] = datamodels.StorageApplicationCommonListType{
				ID:          id,
				Description: vo.Description,
			}
		}

		isTrue = false
	}

	if len(listInsert) == 0 {
		return listID, nil
	}

	_, err = qp.InsertData(listInsert, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "commonpropertiesobjectstix.type", Value: 1},
				{Key: "commonpropertiesobjectstix.id", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source_ref", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	})

	return listID, err
}

// GetListGroupingObjectSTIX возвращает из БД список STIX DO объектов типа Grouping
func GetListGroupingObjectSTIX(cur *mongo.Cursor) []datamodels.GroupingDomainObjectsSTIX {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var list []datamodels.GroupingDomainObjectsSTIX

	for cur.Next(ctx) {
		var gdostix datamodels.GroupingDomainObjectsSTIX
		if err := cur.Decode(&gdostix); err != nil {
			continue
		}

		list = append(list, gdostix)
	}

	return list
}

// GetListGroupingComputerThreat обрабатывает список STIX DO объектов типа Grouping и возвращает набор элементов содержащий лишь некоторые поля
// из данного объекта, а также подсчитывает количество элементов в поле object_ref
func GetListGroupingComputerThreat(cur *mongo.Cursor) []datamodels.ShortDescriptionElementGroupingComputerThreat {
	var (
		list               []datamodels.GroupingDomainObjectsSTIX
		listComputerThreat []datamodels.ShortDescriptionElementGroupingComputerThreat
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for cur.Next(ctx) {
		var gdostix datamodels.GroupingDomainObjectsSTIX
		if err := cur.Decode(&gdostix); err != nil {
			continue
		}

		list = append(list, gdostix)
	}

	for _, v := range list {
		listComputerThreat = append(listComputerThreat, datamodels.ShortDescriptionElementGroupingComputerThreat{
			ID:              v.ID,
			Type:            v.Type,
			Name:            v.Name,
			Description:     v.Description,
			CountObjectRefs: len(v.ObjectRefs),
		})
	}

	return listComputerThreat
}

// getPropertyObjectRefs вспомогоательная функция для получения списка идентификаторов объектов STIX содержащихся с свойстве 'object_refs'
// некоторых объектов STIX
func getPropertyObjectRefs(element *datamodels.ElementSTIXObject) ([]mstixo.IdentifierTypeSTIX /*[]datamodels.IdentifierTypeSTIX*/, error) {
	//var or []datamodels.IdentifierTypeSTIX
	var or []mstixo.IdentifierTypeSTIX

	switch element.DataType {
	case "grouping":
		obj, ok := element.Data.(datamodels.GroupingDomainObjectsSTIX)
		if !ok {
			return or, fmt.Errorf("conversion error")
		}

		or = obj.ObjectRefs

	case "note":
		obj, ok := element.Data.(datamodels.NoteDomainObjectsSTIX)
		if !ok {
			return or, fmt.Errorf("conversion error")
		}

		or = obj.ObjectRefs

	case "observed-data":
		obj, ok := element.Data.(datamodels.ObservedDataDomainObjectsSTIX)
		if !ok {
			return or, fmt.Errorf("conversion error")
		}

		or = obj.ObjectRefs

	case "opinion":
		obj, ok := element.Data.(datamodels.OpinionDomainObjectsSTIX)
		if !ok {
			return or, fmt.Errorf("conversion error")
		}

		or = obj.ObjectRefs

	case "report":
		obj, ok := element.Data.(datamodels.ReportDomainObjectsSTIX)
		if !ok {
			return or, fmt.Errorf("conversion error")
		}

		or = obj.ObjectRefs

	}

	return or, nil
}

// CreatingAdditionalRelationshipSTIXObject создает дополнительные STIX объекты типа 'relationship', обеспечивающие обратные связи для STIX объектов
// перечисленных в свойстве 'object_refs' и содержащемся в таких STIX объектах, как 'grouping', 'note', 'observed-data', 'opinion', 'report'
func CreatingAdditionalRelationshipSTIXObject(qp QueryParameters, l []*datamodels.ElementSTIXObject) ([]*datamodels.ElementSTIXObject, error) {
	var (
		listIDTargetRef                  = []string{}
		modelRelationship                datamodels.RelationshipObjectSTIX
		listRelationshipSTIXObject       = []datamodels.RelationshipObjectSTIX{}
		createListRelationshipSTIXObject = []datamodels.RelationshipObjectSTIX{}
		listFoundRelationshipSTIXObject  = []datamodels.RelationshipObjectSTIX{}
		listTrueSTIXObject               = map[string]struct {
			ObjectRefs []mstixo.IdentifierTypeSTIX
		}{}
	)

	//поиск объектов типа 'grouping', 'note', 'observed-data', 'opinion' или 'report' среди объектов полученных от пользователя
	// и сохранение ссылки из свойства ObjectRef данных объектов, в отдельный объект
	for _, v := range l {
		if v.DataType == "relationship" {
			if rso, ok := v.Data.(datamodels.RelationshipObjectSTIX); ok {
				listRelationshipSTIXObject = append(listRelationshipSTIXObject, rso)
			}

			continue
		}

		for _, vType := range listTypeSTIXObject {
			if vType == v.DataType {
				or, err := getPropertyObjectRefs(v)
				if err != nil {
					return l, err
				}

				listTrueSTIXObject[v.Data.GetID()] = struct {
					ObjectRefs []mstixo.IdentifierTypeSTIX
				}{ObjectRefs: or}
				listIDTargetRef = append(listIDTargetRef, v.Data.GetID())

				break
			}
		}
	}

	//поиск в БД объектов типа 'relationship', где свойство target_ref будет равно ID одного из объектов типа: 'grouping', 'report',
	// 'note', 'observed-data', 'opinion'
	cur, err := qp.Find(bson.D{
		bson.E{Key: "commonpropertiesobjectstix.type", Value: "relationship"},
		bson.E{Key: "target_ref", Value: bson.D{{Key: "$in", Value: listIDTargetRef}}}})
	if err != nil {
		return l, err
	}

	for cur.Next(context.Background()) {
		if err := cur.Decode(&modelRelationship); err != nil {
			continue
		}

		listFoundRelationshipSTIXObject = append(listFoundRelationshipSTIXObject, modelRelationship)
	}

	//поиск в найденных объектах типа 'relationship' совпадений, ID в свойстве 'target_ref' должно соответствовать ID одному из объектов типа:
	// 'grouping', 'report', 'note', 'observed' или 'opinion', а ID в свойстве 'source_ref' должно соответствовать одному из ID в свойстве
	// 'object_ref' объектов типа: 'grouping', 'report', 'note', 'observed' или 'opinion' если совпадения нет, то необходимо создать объект типа
	// 'relateonship', обеспечивающий обратные связи
	for id, objRef := range listTrueSTIXObject {
		for _, idor := range objRef.ObjectRefs {
			tmpRelationship := datamodels.RelationshipObjectSTIX{
				RelationshipObjectSTIX: mstixo.RelationshipObjectSTIX{
					CommonPropertiesObjectSTIX: mstixo.CommonPropertiesObjectSTIX{
						Type: "relationship",
						ID:   fmt.Sprintf("relationship--%s", uuid.NewString()),
					},
					OptionalCommonPropertiesRelationshipObjectSTIX: mstixo.OptionalCommonPropertiesRelationshipObjectSTIX{
						SpecVersion: "2.1",
						Created:     time.Now(),
						Modified:    time.Now(),
					},
					Description: "an automatically created object for establishing feedbacks",
					SourceRef:   idor,
					TargetRef:   mstixo.IdentifierTypeSTIX(id),
				},
			}

			//поиск по списку объектов типа 'relationship' полученных от пользователя
			if len(listRelationshipSTIXObject) != 0 {
				for _, v := range listRelationshipSTIXObject {
					if (v.SourceRef == idor) && (v.TargetRef == mstixo.IdentifierTypeSTIX(id)) {
						tmpRelationship = datamodels.RelationshipObjectSTIX{}

						break
					}
				}
			}

			//поиск по списку объектов типа 'relationship' полученных из БД
			for _, vrs := range listFoundRelationshipSTIXObject {
				if id != string(vrs.TargetRef) {
					continue
				}

				if idor == vrs.SourceRef {
					tmpRelationship = datamodels.RelationshipObjectSTIX{}

					break
				}
			}

			if tmpRelationship.ID != "" {
				createListRelationshipSTIXObject = append(createListRelationshipSTIXObject, tmpRelationship)
			}
		}
	}

	//добавляем вновь созданные объекты типа 'relationship' в основной список объектов, который был получен от пользователя
	// и котороый необходимо добавить в БД
	for _, v := range createListRelationshipSTIXObject {
		l = append(l, &datamodels.ElementSTIXObject{
			DataType: v.Type,
			Data:     v,
		})
	}

	return l, nil
}

// DeleteOldRelationshipSTIXObject удаляет дополнительные STIX объекты типа 'relationship', обеспечивающие обратные связи для STIX объектов
// идентификаторы которых содержатся в свойстве 'object_ref'
func DeleteOldRelationshipSTIXObject(qp QueryParameters, l []*datamodels.ElementSTIXObject) error {
	var (
		listIDTargetRef      = []string{}
		listSearchParameters bson.A
		listTrueSTIXObject   = map[string]struct {
			ObjectRefs []mstixo.IdentifierTypeSTIX
		}{}
		listIDDelRelationshipSTIXObject = []struct {
			SourceRef, TargetRef string
		}{}
	)

	//поиск объектов типа "grouping", "note", "observed-data", "opinion", "report" среди объектов полученных от пользователя
	// и сохранение ссылки из свойства ObjectRef данных объектов, в отдельный объект
	for _, v := range l {
		for _, vType := range listTypeSTIXObject {
			if vType == v.DataType {
				or, err := getPropertyObjectRefs(v)
				if err != nil {
					return err
				}

				listTrueSTIXObject[v.Data.GetID()] = struct {
					ObjectRefs []mstixo.IdentifierTypeSTIX
				}{ObjectRefs: or}
				listIDTargetRef = append(listIDTargetRef, v.Data.GetID())

				break
			}
		}
	}

	//поиск в БД объектов типа: 'grouping', 'report', 'note', 'observed-data', 'opinion'
	cur, err := qp.Find(bson.D{bson.E{Key: "commonpropertiesobjectstix.id", Value: bson.D{{Key: "$in", Value: listIDTargetRef}}}})
	if err != nil {
		return err
	}

	for _, v := range GetListElementSTIXObject(cur) {
		for id, lor := range listTrueSTIXObject {
			if v.Data.GetID() != id {
				continue
			}

			var listObjectRefs []mstixo.IdentifierTypeSTIX

			switch v.DataType {
			case "grouping":
				data := v.GetGroupingDomainObjectsSTIX()
				listObjectRefs = data.ObjectRefs

			case "note":
				data := v.GetNoteDomainObjectsSTIX()
				listObjectRefs = data.ObjectRefs

			case "observed-data":
				data := v.GetObservedDataDomainObjectsSTIX()
				listObjectRefs = data.ObjectRefs

			case "opinion":
				data := v.GetOpinionDomainObjectsSTIX()
				listObjectRefs = data.ObjectRefs

			case "report":
				data := v.GetReportDomainObjectsSTIX()
				listObjectRefs = data.ObjectRefs

			}

			if len(listObjectRefs) == 0 {
				continue
			}

			for _, value := range listObjectRefs {
				isExist := false
				for _, idor := range lor.ObjectRefs {
					if value == idor {
						isExist = true

						break
					}
				}

				if isExist {
					continue
				}

				listIDDelRelationshipSTIXObject = append(listIDDelRelationshipSTIXObject, struct {
					SourceRef string
					TargetRef string
				}{
					SourceRef: string(value),
					TargetRef: id,
				})
			}
		}
	}

	for _, elem := range listIDDelRelationshipSTIXObject {
		listSearchParameters = append(listSearchParameters, bson.D{
			bson.E{Key: "source_ref", Value: elem.SourceRef},
			bson.E{Key: "target_ref", Value: elem.TargetRef},
		})
	}

	if len(listSearchParameters) == 0 {
		return nil
	}

	if _, err := qp.DeleteManyData(bson.D{
		bson.E{Key: "commonpropertiesobjectstix.type", Value: "relationship"},
		bson.E{Key: "$or", Value: listSearchParameters}},
		options.Delete()); err != nil {
		return err
	}

	return nil
}

// switchMSGType - функция заполняющая одно из информационных полей cообщения
// распознавая тип объекта передаваемого в нее
func switchMSGType(msg *datamodels.ModuleDataBaseInteractionChannel, m interface{}) bool {
	msg.ErrorMessage = datamodels.ErrorDataTypePassedThroughChannels{}
	msg.InformationMessage = datamodels.InformationDataTypePassedThroughChannels{}
	switch m.(type) {
	case datamodels.ErrorDataTypePassedThroughChannels:
		msg.ErrorMessage = m.(datamodels.ErrorDataTypePassedThroughChannels)
		return true
	case datamodels.InformationDataTypePassedThroughChannels:
		msg.InformationMessage = m.(datamodels.InformationDataTypePassedThroughChannels)
		return true
	default:
		return false
	}
}
