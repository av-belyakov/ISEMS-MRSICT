package routingflowsmoduleapirequestprocessing

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"time"

	"ISEMS-MRSICT/commonlibs"
	"ISEMS-MRSICT/datamodels"
	"ISEMS-MRSICT/decoders"
	"ISEMS-MRSICT/memorytemporarystoragecommoninformation"

	"github.com/asaskevich/govalidator"
	mstixo "github.com/av-belyakov/methodstixobjects"
)

// UnmarshalJSONObjectSTIXReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит список объектов STIX
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

// UnmarshalJSONObjectReqSearchParameters декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который содержит параметры для
// выполнения поиска документов в коллекциях БД
func UnmarshalJSONObjectReqSearchParameters(msgReq *json.RawMessage) (datamodels.ModAPIRequestProcessingResJSONSearchReqType, error) {
	var result datamodels.ModAPIRequestProcessingResJSONSearchReqType
	var resultTmp datamodels.CommonModAPIRequestProcessingResJSONSearchReqType

	if msgReq == nil {
		return result, fmt.Errorf("the accepted json document does not contain the request parameters")
	}

	if err := json.Unmarshal(*msgReq, &resultTmp); err != nil {
		return result, err
	}

	switch resultTmp.CollectionName {
	case "stix object":
		var msgSearch datamodels.SearchThroughCollectionSTIXObjectsType
		if err := json.Unmarshal(*resultTmp.SearchParameters, &msgSearch); err != nil {
			return result, err
		}

		result.CollectionName = resultTmp.CollectionName
		result.PaginateParameters = resultTmp.PaginateParameters
		result.SortableField = resultTmp.SortableField
		result.SearchParameters = msgSearch

	case "get list computer threat", "stix object list type grouping":
		var msgSearch struct {
			TypeList string `json:"type_list"`
		}

		if err := json.Unmarshal(*resultTmp.SearchParameters, &msgSearch); err != nil {
			return result, err
		}

		result.CollectionName = resultTmp.CollectionName
		result.PaginateParameters = resultTmp.PaginateParameters
		result.SortableField = resultTmp.SortableField
		result.SearchParameters = msgSearch

	case "differences objects collection":
		var msgSearch struct {
			DocumentID     string `json:"document_id"`
			CollectionName string `json:"collection_name"`
		}

		if err := json.Unmarshal(*resultTmp.SearchParameters, &msgSearch); err != nil {
			return result, err
		}

		result.CollectionName = resultTmp.CollectionName
		result.PaginateParameters = resultTmp.PaginateParameters
		result.SortableField = resultTmp.SortableField
		result.SearchParameters = msgSearch
	}

	return result, nil
}

// UnmarshalJSONReferenceBookReq декодирует JSON документ, поступающий от модуля 'moduleapirequestprocessing', который cодержит список действий со справочной информацией
// и данные необходитимые для выполнения данных действий
func UnmarshalJSONRBookReq(reqDetails *json.RawMessage) (*datamodels.RBookReqParameters, error) {
	var (
		resultReqDetails datamodels.RBookReqParameters
		err              error
	)
	if err = json.Unmarshal(*reqDetails, &resultReqDetails); err != nil {
		return nil, err
	}
	return &resultReqDetails, err
}

// CheckSearchSTIXObject выполняет валидацию параметров запроса для поиска информации по STIX объектам
func CheckSearchSTIXObject(req *datamodels.ModAPIRequestProcessingResJSONSearchReqType) (datamodels.ModAPIRequestProcessingResJSONSearchReqType, error) {
	var sortableFieldIsOK bool
	listSortableField := []string{
		"document_type",
		"data_created",
		"data_modified",
		"data_first_seen",
		"data_last_seen",
		"ipv4",
		"ipv6",
		"country",
	}

	//проверяем значение поля по которому будет выполнена сортировка
	if req.SortableField != "" {
		for _, v := range listSortableField {
			if v == req.SortableField {
				sortableFieldIsOK = true

				break
			}
		}

		if !sortableFieldIsOK {
			return *req, fmt.Errorf("invalid field value 'SortableField'")
		}
	}

	sp, ok := req.SearchParameters.(datamodels.SearchThroughCollectionSTIXObjectsType)
	if !ok {
		return *req, fmt.Errorf("type conversion error")
	}

	if len(sp.DocumentsID) > 0 {
		for _, v := range sp.DocumentsID {
			if !(regexp.MustCompile(`^([0-9a-z|-]+)(--)([0-9a-f|-]+)$`).MatchString(v)) {
				return *req, fmt.Errorf("invalid search value accepted in 'DocumentsID' field")
			}
		}
	}

	if len(sp.DocumentsType) > 0 {
		for _, v := range sp.DocumentsType {
			if !(regexp.MustCompile(`^[0-9a-z|-]+$`).MatchString(v)) {
				return *req, fmt.Errorf("invalid search value accepted in 'DocumentsType' field")
			}
		}
	}

	tcsn := sp.Created.Start.Unix()
	tcen := sp.Created.End.Unix()

	if tcsn > 0 && tcen > 0 {
		if tcsn >= tcen {
			return *req, fmt.Errorf("invalid search value accepted in 'Created.Start' or 'Created.End' fields")
		}
	}

	tmsn := sp.Modified.Start.Unix()
	tmen := sp.Modified.End.Unix()

	if tmsn > 0 && tmen > 0 {
		if tmsn >= tmen {
			return *req, fmt.Errorf("invalid search value accepted in 'Modified.Start' or 'Modified.End' fields")
		}
	}

	sp.CreatedByRef = commonlibs.StringSanitize(sp.CreatedByRef)

	//поля не входящие в основную спецификацию STIX
	sp.OutsideSpecificationSearchFields = datamodels.OutsideSpecificationSearchFieldsType{
		DecisionsMadeComputerThreat: commonlibs.StringSanitize(sp.OutsideSpecificationSearchFields.DecisionsMadeComputerThreat),
		ComputerThreatType:          commonlibs.StringSanitize(sp.OutsideSpecificationSearchFields.ComputerThreatType),
	}

	//наличие дополнительных полей
	if len(sp.SpecificSearchFields) == 0 {
		return *req, nil
	}

	for k, v := range sp.SpecificSearchFields {
		if len(v.Value) > 0 {
			if err := checkSearchFieldsValue(v.Value); err != nil {
				return *req, err
			}
		}

		sp.SpecificSearchFields[k].Name = commonlibs.StringSanitize(v.Name)

		if len(v.Aliases) > 0 {
			for key, value := range v.Aliases {
				sp.SpecificSearchFields[k].Aliases[key] = commonlibs.StringSanitize(value)
			}
		}

		tfsn := v.FirstSeen.Start.Unix()
		tfen := v.FirstSeen.End.Unix()
		if tfsn > 0 && tfen > 0 {
			if tfsn >= tfen {
				return *req, fmt.Errorf("invalid search value accepted in 'FirstSeen.Start' or 'FirstSeen.End' fields")
			}
		}

		tlsn := v.LastSeen.Start.Unix()
		tlen := v.FirstSeen.End.Unix()
		if tlsn > 0 && tlen > 0 {
			if tlsn >= tlen {
				return *req, fmt.Errorf("invalid search value accepted in 'LastSeen.Start' or 'LastSeen.End' fields")
			}
		}

		if len(v.Roles) > 0 {
			for key, value := range v.Roles {
				sp.SpecificSearchFields[k].Roles[key] = commonlibs.StringSanitize(value)
			}
		}

		if v.Country != "" {
			if !(regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(v.Country)) {
				return *req, fmt.Errorf("invalid search value accepted in 'Country' field")
			}
		}

		sp.SpecificSearchFields[k].City = commonlibs.StringSanitize(v.City)
	}

	return *req, nil
}

// checkSearchFieldsValue выполняет проверку поля "Value" на соответствие одному из типов значений "domain-name", "email-addr", "ipv4-addr",
// "ipv6-addr" или "url"
func checkSearchFieldsValue(l []string) error {
	for _, v := range l {
		ipCIDR, _, _ := net.ParseCIDR(v)

		if ipCIDR != nil {
			continue
		} else if net.ParseIP(v) != nil {
			continue
		} else if govalidator.IsDNSName(v) {
			continue
		} else if govalidator.IsEmail(v) {
			continue
		} else if govalidator.IsURL(v) {
			continue
		} else {
			return fmt.Errorf("invalid search value accepted in 'Value' field, type undefined")
		}
	}

	return nil
}

// CheckSTIXObjects выполняет валидацию списка STIX объектов
func CheckSTIXObjects(l []*datamodels.ElementSTIXObject) error {
	for _, item := range l {
		if item.Data.ValidateStruct() {
			continue
		}

		return fmt.Errorf("one or more STIX objects are invalid")
	}

	return nil
}

func VerifyOutsideSpecificationFields(
	l []*datamodels.ElementSTIXObject,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	clientName string) {
	verifyDecisionsMadeComputerThreat := func(valueBeChecked string) string {
		result := ""

		if valueBeChecked == "" {
			return result
		}

		ldm, err := tst.GetListDecisionsMade()
		if err != nil {
			return result
		}

		var isExist bool
		for k := range ldm {
			if k == valueBeChecked {
				isExist = true

				break
			}
		}

		if isExist {
			return valueBeChecked
		}

		return result
	}

	verifyComputerThreatType := func(valueBeChecked string) string {
		result := ""

		if valueBeChecked == "" {
			return result
		}

		ldm, err := tst.GetListComputerThreat()
		if err != nil {
			return result
		}

		var isExist bool
		for k := range ldm {
			if k == valueBeChecked {
				isExist = true

				break
			}
		}

		if isExist {
			return valueBeChecked
		}

		return result
	}

	rand.Seed(82)

	for k, v := range l {
		switch v.DataType {
		case "report":
			e, ok := v.Data.(datamodels.ReportDomainObjectsSTIX)
			if !ok {
				continue
			}

			salt := fmt.Sprint(rand.Intn(10000))

			e.OutsideSpecification = mstixo.ReportOutsideSpecification{
				AdditionalName:              fmt.Sprintf("report--%s-%v.%v", clientName, time.Now().Unix(), salt),
				DecisionsMadeComputerThreat: verifyDecisionsMadeComputerThreat(e.OutsideSpecification.DecisionsMadeComputerThreat),
				ComputerThreatType:          verifyComputerThreatType(e.OutsideSpecification.ComputerThreatType),
			}

			l[k] = &datamodels.ElementSTIXObject{
				DataType: v.DataType,
				Data:     e,
			}
		}
	}
}

// SanitizeSTIXObject выполняем санитаризацию полученных STIX объектов
func SanitizeSTIXObject(l []*datamodels.ElementSTIXObject) []*datamodels.ElementSTIXObject {
	var elem datamodels.HandlerSTIXObject
	listElements := make([]*datamodels.ElementSTIXObject, 0, len(l))

	for _, item := range l {
		switch element := item.Data.(type) {
		case datamodels.AttackPatternDomainObjectsSTIX:
			elem = datamodels.AttackPatternDomainObjectsSTIX{AttackPatternDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.CampaignDomainObjectsSTIX:
			elem = datamodels.CampaignDomainObjectsSTIX{CampaignDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.CourseOfActionDomainObjectsSTIX:
			elem = datamodels.CourseOfActionDomainObjectsSTIX{CourseOfActionDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.GroupingDomainObjectsSTIX:
			elem = datamodels.GroupingDomainObjectsSTIX{GroupingDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.IdentityDomainObjectsSTIX:
			elem = datamodels.IdentityDomainObjectsSTIX{IdentityDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.IndicatorDomainObjectsSTIX:
			elem = datamodels.IndicatorDomainObjectsSTIX{IndicatorDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.InfrastructureDomainObjectsSTIX:
			elem = datamodels.InfrastructureDomainObjectsSTIX{InfrastructureDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.IntrusionSetDomainObjectsSTIX:
			elem = datamodels.IntrusionSetDomainObjectsSTIX{IntrusionSetDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.LocationDomainObjectsSTIX:
			elem = datamodels.LocationDomainObjectsSTIX{LocationDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.MalwareDomainObjectsSTIX:
			elem = datamodels.MalwareDomainObjectsSTIX{MalwareDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.MalwareAnalysisDomainObjectsSTIX:
			elem = datamodels.MalwareAnalysisDomainObjectsSTIX{MalwareAnalysisDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.NoteDomainObjectsSTIX:
			elem = datamodels.NoteDomainObjectsSTIX{NoteDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.ObservedDataDomainObjectsSTIX:
			elem = datamodels.ObservedDataDomainObjectsSTIX{ObservedDataDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.OpinionDomainObjectsSTIX:
			elem = datamodels.OpinionDomainObjectsSTIX{OpinionDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.ReportDomainObjectsSTIX:
			elem = datamodels.ReportDomainObjectsSTIX{ReportDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.ThreatActorDomainObjectsSTIX:
			elem = datamodels.ThreatActorDomainObjectsSTIX{ThreatActorDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.ToolDomainObjectsSTIX:
			elem = datamodels.ToolDomainObjectsSTIX{ToolDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.VulnerabilityDomainObjectsSTIX:
			elem = datamodels.VulnerabilityDomainObjectsSTIX{VulnerabilityDomainObjectsSTIX: element.SanitizeStruct()}

		case datamodels.RelationshipObjectSTIX:
			elem = datamodels.RelationshipObjectSTIX{RelationshipObjectSTIX: element.SanitizeStruct()}

		case datamodels.SightingObjectSTIX:
			elem = datamodels.SightingObjectSTIX{SightingObjectSTIX: element.SanitizeStruct()}

		case datamodels.ArtifactCyberObservableObjectSTIX:
			elem = datamodels.ArtifactCyberObservableObjectSTIX{ArtifactCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.AutonomousSystemCyberObservableObjectSTIX:
			elem = datamodels.AutonomousSystemCyberObservableObjectSTIX{AutonomousSystemCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.DirectoryCyberObservableObjectSTIX:
			elem = datamodels.DirectoryCyberObservableObjectSTIX{DirectoryCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.DomainNameCyberObservableObjectSTIX:
			elem = datamodels.DomainNameCyberObservableObjectSTIX{DomainNameCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.EmailAddressCyberObservableObjectSTIX:
			elem = datamodels.EmailAddressCyberObservableObjectSTIX{EmailAddressCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.EmailMessageCyberObservableObjectSTIX:
			elem = datamodels.EmailMessageCyberObservableObjectSTIX{EmailMessageCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.FileCyberObservableObjectSTIX:
			elem = datamodels.FileCyberObservableObjectSTIX{FileCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.IPv4AddressCyberObservableObjectSTIX:
			elem = datamodels.IPv4AddressCyberObservableObjectSTIX{IPv4AddressCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.IPv6AddressCyberObservableObjectSTIX:
			elem = datamodels.IPv6AddressCyberObservableObjectSTIX{IPv6AddressCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.MACAddressCyberObservableObjectSTIX:
			elem = datamodels.MACAddressCyberObservableObjectSTIX{MACAddressCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.MutexCyberObservableObjectSTIX:
			elem = datamodels.MutexCyberObservableObjectSTIX{MutexCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.NetworkTrafficCyberObservableObjectSTIX:
			elem = datamodels.NetworkTrafficCyberObservableObjectSTIX{NetworkTrafficCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.ProcessCyberObservableObjectSTIX:
			elem = datamodels.ProcessCyberObservableObjectSTIX{ProcessCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.SoftwareCyberObservableObjectSTIX:
			elem = datamodels.SoftwareCyberObservableObjectSTIX{SoftwareCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.URLCyberObservableObjectSTIX:
			elem = datamodels.URLCyberObservableObjectSTIX{URLCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.UserAccountCyberObservableObjectSTIX:
			elem = datamodels.UserAccountCyberObservableObjectSTIX{UserAccountCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.WindowsRegistryKeyCyberObservableObjectSTIX:
			elem = datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{WindowsRegistryKeyCyberObservableObjectSTIX: element.SanitizeStruct()}

		case datamodels.X509CertificateCyberObservableObjectSTIX:
			elem = datamodels.X509CertificateCyberObservableObjectSTIX{X509CertificateCyberObservableObjectSTIX: element.SanitizeStruct()}

		}

		if elem == nil {
			continue
		}

		listElements = append(listElements, &datamodels.ElementSTIXObject{
			DataType: item.DataType,
			Data:     elem,
		})
	}

	/*for _, item := range l {
		switch element := item.Data.(type) {
		case datamodels.AttackPatternDomainObjectsSTIX:
			//e := element.
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
	}*/

	return listElements
}

// SanitizeReqRBObject выполняем санитаризацию запросов к объектам справочникам
func SanitizeReqRBObject(rbrs *datamodels.RBookReqParameters) *datamodels.RBookReqParameters {
	rbrs.Sanitize()
	return rbrs
}
