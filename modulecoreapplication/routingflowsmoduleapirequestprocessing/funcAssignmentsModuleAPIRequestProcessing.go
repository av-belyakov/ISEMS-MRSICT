package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/decoders"
)

//UnmarshalJSONObjectSTIXReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит список объектов STIX
func UnmarshalJSONObjectSTIXReq(msgReq datamodels.ModAPIRequestProcessingReqJSON) ([]*datamodels.ElementSTIXObject, error) {
	var listSTIXObjectJSON []*json.RawMessage

	if err := json.Unmarshal(*msgReq.RequestDetails, &listSTIXObjectJSON); err != nil {
		return nil, err
	}

	listResults, err := decoders.GetListSTIXObjectFromJSON(listSTIXObjectJSON)
	if err != nil {
		return nil, err
	}

	return listResults, nil
}

//CheckSTIXObjects выполняет валидацию списка STIX объектов
func CheckSTIXObjects(l []*datamodels.ElementSTIXObject) error {
	for _, item := range l {
		if item.Data.CheckingTypeFields() {
			continue
		}

		fmt.Printf("Error checking type STIX object: '%s'\n", item.DataType)

		return fmt.Errorf("one or more STIX objects are invalid")
	}

	return nil
}

//SanitizeSTIXObject выполняем санитаризацию полученных STIX объектов
func SanitizeSTIXObject(l []*datamodels.ElementSTIXObject) []*datamodels.ElementSTIXObject {
	var elem datamodels.HandlerSTIXObject
	listElements := make([]*datamodels.ElementSTIXObject, 0, len(l))

	for _, item := range l {
		switch element := item.Data.(type) {
		case datamodels.AttackPatternDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.CampaignDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.CourseOfActionDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.GroupingDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.IdentityDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.IndicatorDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.InfrastructureDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.IntrusionSetDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.LocationDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.MalwareDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.MalwareAnalysisDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.NoteDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ObservedDataDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.OpinionDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ReportDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ThreatActorDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ToolDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.VulnerabilityDomainObjectsSTIX:
			elem = element.SanitizeStruct()

		case datamodels.RelationshipObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.SightingObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ArtifactCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.AutonomousSystemCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.DirectoryCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.DomainNameCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.EmailAddressCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.EmailMessageCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.FileCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.IPv4AddressCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.IPv6AddressCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.MACAddressCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.MutexCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.NetworkTrafficCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.ProcessCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.SoftwareCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.URLCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.UserAccountCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.WindowsRegistryKeyCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		case datamodels.X509CertificateCyberObservableObjectSTIX:
			elem = element.SanitizeStruct()

		}

		if elem == nil {
			continue
		}

		listElements = append(listElements, &datamodels.ElementSTIXObject{
			DataType: item.DataType,
			Data:     elem,
		})
	}

	return listElements
}
