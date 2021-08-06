package datamodels

import (
	"encoding/json"
	"time"
)

//HandlerSTIXObject интерфейс реализующий обработчики для STIX объектов
type HandlerSTIXObject interface {
	DecoderJSONObject
	EncoderJSONObject
	ValidatorJSONObject
	GetterParametersSTIXObject
	ComparatorSTIXObject
}

//DecoderJSONObject интерфейс реализующий обработчик для декодирования JSON объекта в STIX объект
type DecoderJSONObject interface {
	DecoderJSON(*json.RawMessage) (interface{}, error)
}

//EncoderJSONObject интерфейс реализующий обработчик для кодирования STIX объекта в JSON объект
type EncoderJSONObject interface {
	EncoderJSON(interface{}) (*[]byte, error)
}

//ValidatorJSONObject интерфейс реализующий обработчик для валидации STIX объектов
type ValidatorJSONObject interface {
	ValidateStruct() bool
}

//GetterParametersSTIXObject интерфейс реализующий обработчик для получения ID STIX объекта
type GetterParametersSTIXObject interface {
	GetID() string
	GetType() string
}

//ComparatorSTIXObject интерфейс реализующий обработчик для сравнения STIX объектов одного типа
type ComparatorSTIXObject interface {
	ComparisonTypeCommonFields(interface{}, string) (bool, DifferentObjectType, error)
}

//ToBeautifulOutputConverter интерфейс реализующий обработчик для красивого представления данных хранящихся в пользовательской структуре
type ToBeautifulOutputConverter interface {
	ToStringBeautiful() string
}

type GetIPv4AddressCyberObservableObjectSTIX interface {
	GetIPv4AddressCyberObservableObjectSTIX() *IPv4AddressCyberObservableObjectSTIX
}

type GetFileCyberObservableObjectSTIX interface {
	GetFileCyberObservableObjectSTIX() *FileCyberObservableObjectSTIX
}

//ElementSTIXObject может содержать любой из STIX объектов с указанием его типа
// DataType - тип STIX объекта
// Data - непосредственно сам STIX объект
type ElementSTIXObject struct {
	DataType string
	Data     HandlerSTIXObject
}

//DifferentObjectType содержит перечисление полей и их значения, которые были изменены в произвольном типе
// SourceReceivingChanges - источник от которого были получены изменения
// ModifiedTime - время выполнения модификации
// CollectionName - наименование коллекции в которой выполнялись модификации
// DocumentID - идентификатор документа в котором выполнялись модификации
// FieldList - перечень полей подвергшихся изменениям
type DifferentObjectType struct {
	SourceReceivingChanges string                    `bson:"source_receiving_changes"`
	ModifiedTime           time.Time                 `bson:"modified_time"`
	CollectionName         string                    `bson:"collection_name"`
	DocumentID             string                    `bson:"document_id"`
	FieldList              []OldFieldValueObjectType `bson:"field_list"`
}

//OldFieldValueObjectType содержит старое значение полей, до их модификации
// FeildType - тип поля
// Path - полный путь к объекту подвергшемуся модификации
// Value - предыдущее значение поля, которое подверглось модификации
type OldFieldValueObjectType struct {
	FeildType string      `bson:"feild_type"`
	Path      string      `bson:"path"`
	Value     interface{} `bson:"value"`
}

//ShortDescriptionElementComputerThreat содержит краткое описание элемента 'grouping' содержащего списки STIX объектов типа 'report' компьютерных
// угроз
type ShortDescriptionElementGroupingComputerThreat struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	CountObjectRefs int    `json:"count_object_refs"`
}

/********** 			Domain Objects STIX			**********/

//GetAttackPatternDomainObjectsSTIX возвращает объект STIX типа 'attack-pattern'
func (estix *ElementSTIXObject) GetAttackPatternDomainObjectsSTIX() *AttackPatternDomainObjectsSTIX {
	if result, ok := estix.Data.(AttackPatternDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetCampaignDomainObjectsSTIX возвращает объект STIX типа 'campaign'
func (estix *ElementSTIXObject) GetCampaignDomainObjectsSTIX() *CampaignDomainObjectsSTIX {
	if result, ok := estix.Data.(CampaignDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetCourseOfActionDomainObjectsSTIX возвращает объект STIX типа 'course-of-action'
func (estix *ElementSTIXObject) GetCourseOfActionDomainObjectsSTIX() *CourseOfActionDomainObjectsSTIX {
	if result, ok := estix.Data.(CourseOfActionDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetGroupingDomainObjectsSTIX возвращает объект STIX типа 'grouping'
func (estix *ElementSTIXObject) GetGroupingDomainObjectsSTIX() *GroupingDomainObjectsSTIX {
	if result, ok := estix.Data.(GroupingDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetIdentityDomainObjectsSTIX возвращает объект STIX типа 'identity'
func (estix *ElementSTIXObject) GetIdentityDomainObjectsSTIX() *IdentityDomainObjectsSTIX {
	if result, ok := estix.Data.(IdentityDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetIndicatorDomainObjectsSTIX возвращает объект STIX типа 'indicator'
func (estix *ElementSTIXObject) GetIndicatorDomainObjectsSTIX() *IndicatorDomainObjectsSTIX {
	if result, ok := estix.Data.(IndicatorDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetInfrastructureDomainObjectsSTIX возвращает объект STIX типа 'infrastructure'
func (estix *ElementSTIXObject) GetInfrastructureDomainObjectsSTIX() *InfrastructureDomainObjectsSTIX {
	if result, ok := estix.Data.(InfrastructureDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetIntrusionSetDomainObjectsSTIX возвращает объект STIX типа 'intrusion-set'
func (estix *ElementSTIXObject) GetIntrusionSetDomainObjectsSTIX() *IntrusionSetDomainObjectsSTIX {
	if result, ok := estix.Data.(IntrusionSetDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetLocationDomainObjectsSTIX возвращает объект STIX типа 'location'
func (estix *ElementSTIXObject) GetLocationDomainObjectsSTIX() *LocationDomainObjectsSTIX {
	if result, ok := estix.Data.(LocationDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetMalwareDomainObjectsSTIX возвращает объект STIX типа 'malware'
func (estix *ElementSTIXObject) GetMalwareDomainObjectsSTIX() *MalwareDomainObjectsSTIX {
	if result, ok := estix.Data.(MalwareDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetMalwareAnalysisDomainObjectsSTIX возвращает объект STIX типа 'malware-analysis'
func (estix *ElementSTIXObject) GetMalwareAnalysisDomainObjectsSTIX() *MalwareAnalysisDomainObjectsSTIX {
	if result, ok := estix.Data.(MalwareAnalysisDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetNoteDomainObjectsSTIX возвращает объект STIX типа 'note'
func (estix *ElementSTIXObject) GetNoteDomainObjectsSTIX() *NoteDomainObjectsSTIX {
	if result, ok := estix.Data.(NoteDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetObservedDataDomainObjectsSTIX возвращает объект STIX типа 'observed-data'
func (estix *ElementSTIXObject) GetObservedDataDomainObjectsSTIX() *ObservedDataDomainObjectsSTIX {
	if result, ok := estix.Data.(ObservedDataDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetOpinionDomainObjectsSTIX возвращает объект STIX типа 'opinion'
func (estix *ElementSTIXObject) GetOpinionDomainObjectsSTIX() *OpinionDomainObjectsSTIX {
	if result, ok := estix.Data.(OpinionDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetReportDomainObjectsSTIX возвращает объект STIX типа 'report'
func (estix *ElementSTIXObject) GetReportDomainObjectsSTIX() *ReportDomainObjectsSTIX {
	if result, ok := estix.Data.(ReportDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetThreatActorDomainObjectsSTIX возвращает объект STIX типа 'threat-actor'
func (estix *ElementSTIXObject) GetThreatActorDomainObjectsSTIX() *ThreatActorDomainObjectsSTIX {
	if result, ok := estix.Data.(ThreatActorDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetToolDomainObjectsSTIX возвращает объект STIX типа 'tool'
func (estix *ElementSTIXObject) GetToolDomainObjectsSTIX() *ToolDomainObjectsSTIX {
	if result, ok := estix.Data.(ToolDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

//GetVulnerabilityDomainObjectsSTIX возвращает объект STIX типа 'vulnerability'
func (estix *ElementSTIXObject) GetVulnerabilityDomainObjectsSTIX() *VulnerabilityDomainObjectsSTIX {
	if result, ok := estix.Data.(VulnerabilityDomainObjectsSTIX); ok {
		return &result
	}

	return nil
}

/********** 			Relationship Objects STIX			**********/

//GetRelationshipObjectSTIX возвращает объект STIX типа 'relationship'
func (estix *ElementSTIXObject) GetRelationshipObjectSTIX() *RelationshipObjectSTIX {
	if result, ok := estix.Data.(RelationshipObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetSightingObjectSTIX возвращает объект STIX типа 'sighting'
func (estix *ElementSTIXObject) GetSightingObjectSTIX() *SightingObjectSTIX {
	if result, ok := estix.Data.(SightingObjectSTIX); ok {
		return &result
	}

	return nil
}

/********** 			Cyber-observable Objects STIX			**********/

//GetArtifactCyberObservableObjectSTIX возвращает объект STIX типа 'artifact'
func (estix *ElementSTIXObject) GetArtifactCyberObservableObjectSTIX() *ArtifactCyberObservableObjectSTIX {
	if result, ok := estix.Data.(ArtifactCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetAutonomousSystemCyberObservableObjectSTIX возвращает объект STIX типа 'autonomous-system'
func (estix *ElementSTIXObject) GetAutonomousSystemCyberObservableObjectSTIX() *AutonomousSystemCyberObservableObjectSTIX {
	if result, ok := estix.Data.(AutonomousSystemCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetDirectoryCyberObservableObjectSTIX возвращает объект STIX типа 'directory'
func (estix *ElementSTIXObject) GetDirectoryCyberObservableObjectSTIX() *DirectoryCyberObservableObjectSTIX {
	if result, ok := estix.Data.(DirectoryCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetDomainNameCyberObservableObjectSTIX возвращает объект STIX типа 'domain-name'
func (estix *ElementSTIXObject) GetDomainNameCyberObservableObjectSTIX() *DomainNameCyberObservableObjectSTIX {
	if result, ok := estix.Data.(DomainNameCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetEmailAddressCyberObservableObjectSTIX возвращает объект STIX типа 'email-addr'
func (estix *ElementSTIXObject) GetEmailAddressCyberObservableObjectSTIX() *EmailAddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(EmailAddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetEmailMessageCyberObservableObjectSTIX возвращает объект STIX типа 'email-message'
func (estix *ElementSTIXObject) GetEmailMessageCyberObservableObjectSTIX() *EmailMessageCyberObservableObjectSTIX {
	if result, ok := estix.Data.(EmailMessageCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetFileCyberObservableObjectSTIX возвращает объект STIX типа 'file'
func (estix *ElementSTIXObject) GetFileCyberObservableObjectSTIX() *FileCyberObservableObjectSTIX {
	if result, ok := estix.Data.(FileCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetIPv4AddressCyberObservableObjectSTIX возвращает объект STIX типа 'ipv4-addr'
func (estix *ElementSTIXObject) GetIPv4AddressCyberObservableObjectSTIX() *IPv4AddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(IPv4AddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetIPv6AddressCyberObservableObjectSTIX возвращает объект STIX типа 'ipv6-addr'
func (estix *ElementSTIXObject) GetIPv6AddressCyberObservableObjectSTIX() *IPv6AddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(IPv6AddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetMACAddressCyberObservableObjectSTIX возвращает объект STIX типа 'mac-addr'
func (estix *ElementSTIXObject) GetMACAddressCyberObservableObjectSTIX() *MACAddressCyberObservableObjectSTIX {
	if result, ok := estix.Data.(MACAddressCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetMutexCyberObservableObjectSTIX возвращает объект STIX типа 'mutex'
func (estix *ElementSTIXObject) GetMutexCyberObservableObjectSTIX() *MutexCyberObservableObjectSTIX {
	if result, ok := estix.Data.(MutexCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetNetworkTrafficCyberObservableObjectSTIX возвращает объект STIX типа 'network-traffic'
func (estix *ElementSTIXObject) GetNetworkTrafficCyberObservableObjectSTIX() *NetworkTrafficCyberObservableObjectSTIX {
	if result, ok := estix.Data.(NetworkTrafficCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetProcessCyberObservableObjectSTIX возвращает объект STIX типа 'process'
func (estix *ElementSTIXObject) GetProcessCyberObservableObjectSTIX() *ProcessCyberObservableObjectSTIX {
	if result, ok := estix.Data.(ProcessCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetSoftwareCyberObservableObjectSTIX возвращает объект STIX типа 'software'
func (estix *ElementSTIXObject) GetSoftwareCyberObservableObjectSTIX() *SoftwareCyberObservableObjectSTIX {
	if result, ok := estix.Data.(SoftwareCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetURLCyberObservableObjectSTIX возвращает объект STIX типа 'url'
func (estix *ElementSTIXObject) GetURLCyberObservableObjectSTIX() *URLCyberObservableObjectSTIX {
	if result, ok := estix.Data.(URLCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetUserAccountCyberObservableObjectSTIX возвращает объект STIX типа 'user-account'
func (estix *ElementSTIXObject) GetUserAccountCyberObservableObjectSTIX() *UserAccountCyberObservableObjectSTIX {
	if result, ok := estix.Data.(UserAccountCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetWindowsRegistryKeyCyberObservableObjectSTIX возвращает объект STIX типа 'windows-registry-key'
func (estix *ElementSTIXObject) GetWindowsRegistryKeyCyberObservableObjectSTIX() *WindowsRegistryKeyCyberObservableObjectSTIX {
	if result, ok := estix.Data.(WindowsRegistryKeyCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}

//GetX509CertificateCyberObservableObjectSTIX возвращает объект STIX типа 'x509-certificate'
func (estix *ElementSTIXObject) GetX509CertificateCyberObservableObjectSTIX() *X509CertificateCyberObservableObjectSTIX {
	if result, ok := estix.Data.(X509CertificateCyberObservableObjectSTIX); ok {
		return &result
	}

	return nil
}
