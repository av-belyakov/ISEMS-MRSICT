package commonlibs

import (
	"encoding/json"
	"fmt"

	"ISEMS-MRSICT/datamodels"
)

/*
	// - "email-mime-part-type"
	// - "alternate-data-stream-type"
		// - "windows-pe-optional-header-type"
	// - "windows-pe-section-type"
		// - "windows-registry-value-type"
	// - "x509-v3-extensions-type"
*/

//decodingExtensionsSTIX декодирует следующие типы STIX расширений:
// - "archive-ext"
// - "ntfs-ext"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "windows-process-ext"
// - "windows-service-ext"
// - "unix-account-ext"
func decodingExtensionsSTIX(extType string, rawMsg *json.RawMessage) (interface{}, error) {
	var err error
	switch extType {
	case "archive-ext":
		var archiveExt datamodels.ArchiveFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &archiveExt)

		return archiveExt, err
	case "ntfs-ext":
		var ntfsExt datamodels.NTFSFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &ntfsExt)

		return ntfsExt, err
	case "pdf-ext":
		var pdfExt datamodels.PDFFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &pdfExt)

		return pdfExt, err
	case "raster-image-ext":
		var rasterImageExt datamodels.RasterImageFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &rasterImageExt)

		return rasterImageExt, err
	case "windows-pebinary-ext":
		var windowsPebinaryExt datamodels.WindowsPEBinaryFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &windowsPebinaryExt)

		return windowsPebinaryExt, err
	case "http-request-ext":
		var httpRequestExt datamodels.HTTPRequestExtensionSTIX
		err = json.Unmarshal(*rawMsg, &httpRequestExt)

		return httpRequestExt, err
	case "icmp-ext":
		var icmpExt datamodels.ICMPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &icmpExt)

		return icmpExt, err
	case "socket-ext":
		var socketExt datamodels.NetworkSocketExtensionSTIX
		err := json.Unmarshal(*rawMsg, &socketExt)

		return socketExt, err
	case "tcp-ext":
		var tcpExt datamodels.TCPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &tcpExt)

		return tcpExt, err
	case "windows-process-ext":
		var windowsProcessExt datamodels.WindowsProcessExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsProcessExt)

		return windowsProcessExt, err
	case "windows-service-ext":
		var windowsServiceExt datamodels.WindowsServiceExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsServiceExt)

		return windowsServiceExt, err
	case "unix-account-ext":
		var unixAccountExt datamodels.UNIXAccountExtensionSTIX
		err := json.Unmarshal(*rawMsg, &unixAccountExt)

		return unixAccountExt, err
	default:
		return struct{}{}, nil
	}
}

//DecoderFromJSONToSTIXObject декодирует STIX сообщения формата JSON в STIX объект. Выполняется дикодирование следующих типов STIX объектов:
//  1. Для Domain Objects STIX
// - "attack-pattern"
// - "campaign"
// - "course-of-action"
// - "grouping"
// - "identity"
// - "indicator"
// - "infrastructure"
// - "intrusion-set"
// - "location"
// - "malware"
// - "malware-analysis"
// - "note"
// - "observed-data"
// - "opinion"
// - "report"
// - "threat-actor"
// - "tool"
// - "vulnerability"
//  2. Для Relationship Objects STIX
// - "relationship"
// - "sighting"
//  3. Для Cyber-observable Objects STIX
// - "artifact"
// - "autonomous-system"
// - "directory"
// - "domain-name"
// - "email-addr"
// - "email-message"
// - "file"
// - "ipv4-addr"
// - "ipv6-addr"
// - "mac-addr"
// - "mutex"
// - "network-traffic"
// - "process"
// - "software"
// - "url"
// - "user-account"
// - "windows-registry-key"
// - "x509-certificate"
func DecoderFromJSONToSTIXObject(objectType string, rawMessage *json.RawMessage) (interface{}, string, error) {
	ListFuncDecoderFromJSONToSTIXObject := map[string]func(*json.RawMessage) (interface{}, string, error){
		/* *** Domain Objects STIX *** */
		"attack-pattern": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.AttackPatternDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"campaign": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.CampaignDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"course-of-action": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.CourseOfActionDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"grouping": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.GroupingDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"identity": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.IdentityDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"indicator": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.IndicatorDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"infrastructure": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.InfrastructureDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"intrusion-set": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.IntrusionSetDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"location": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.LocationDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"malware": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.MalwareDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"malware-analysis": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.MalwareAnalysisDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"note": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.NoteDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"observed-data": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.ObservedDataDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"opinion": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.OpinionDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"report": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.ReportDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"threat-actor": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.ThreatActorDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"tool": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.ToolDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		"vulnerability": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.VulnerabilityDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "domain object stix", err
			}

			return object, "domain object stix", nil
		},
		/* *** Relationship Objects *** */
		"relationship": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.RelationshipObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "relationship object stix", err
			}

			return object, "relationship object stix", nil
		},
		"sighting": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.SightingObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "relationship object stix", err
			}

			return object, "relationship object stix", nil
		},
		/* *** Cyber-observable Objects STIX *** */
		"artifact": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.ArtifactCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"autonomous-system": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.AutonomousSystemCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"directory": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.DirectoryCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"domain-name": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.DomainNameCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"email-addr": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.EmailAddressCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"email-message": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.EmailMessageCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"file": func(msg *json.RawMessage) (interface{}, string, error) {
			var commonObject datamodels.CommonFileCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &commonObject); err != nil {
				return nil, "cyber observable object stix", err
			}

			fcoostix := datamodels.FileCyberObservableObjectSTIX{
				CommonFileCyberObservableObjectSTIX: commonObject,
			}

			if len(commonObject.Extensions) == 0 {
				return fcoostix, "cyber observable object stix", nil
			}

			ext := map[string]*interface{}{}
			for key, value := range commonObject.Extensions {
				e, err := decodingExtensionsSTIX(key, value)
				if err != nil {
					continue
				}

				ext[key] = &e
			}

			fcoostix.Extensions = ext

			return fcoostix, "cyber observable object stix", nil
		},
		"ipv4-addr": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.IPv4AddressCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"ipv6-addr": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.IPv6AddressCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"mac-addr": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.MACAddressCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"mutex": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.MutexCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"network-traffic": func(msg *json.RawMessage) (interface{}, string, error) {
			var commonObject datamodels.CommonNetworkTrafficCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &commonObject); err != nil {
				return nil, "cyber observable object stix", err
			}

			ntcostix := datamodels.NetworkTrafficCyberObservableObjectSTIX{
				CommonNetworkTrafficCyberObservableObjectSTIX: commonObject,
			}

			if len(commonObject.Extensions) == 0 {
				return ntcostix, "cyber observable object stix", nil
			}

			ext := map[string]*interface{}{}
			for key, value := range commonObject.Extensions {
				e, err := decodingExtensionsSTIX(key, value)
				if err != nil {
					continue
				}

				ext[key] = &e
			}

			ntcostix.Extensions = ext

			return ntcostix, "cyber observable object stix", nil
		},
		"process": func(msg *json.RawMessage) (interface{}, string, error) {
			var commonObject datamodels.CommonProcessCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &commonObject); err != nil {
				return nil, "cyber observable object stix", err
			}

			pcoostix := datamodels.ProcessCyberObservableObjectSTIX{
				CommonProcessCyberObservableObjectSTIX: commonObject,
			}

			if len(commonObject.Extensions) == 0 {
				return pcoostix, "cyber observable object stix", nil
			}

			ext := map[string]*interface{}{}
			for key, value := range commonObject.Extensions {
				e, err := decodingExtensionsSTIX(key, value)
				if err != nil {
					continue
				}

				ext[key] = &e
			}
			pcoostix.Extensions = ext

			return pcoostix, "cyber observable object stix", nil
		},
		"software": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.SoftwareCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"url": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.URLCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"user-account": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.UserAccountCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"windows-registry-key": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.WindowsRegistryKeyCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
		"x509-certificate": func(msg *json.RawMessage) (interface{}, string, error) {
			var object datamodels.X509CertificateCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, "cyber observable object stix", err
			}

			return object, "cyber observable object stix", nil
		},
	}

	return ListFuncDecoderFromJSONToSTIXObject[objectType](rawMessage)
}

func GetListSTIXObjectFromJSON(list []*json.RawMessage) ([]*datamodels.ListSTIXObject, error) {
	var result []*datamodels.ListSTIXObject = make([]*datamodels.ListSTIXObject, 0, len(list))
	var commonPropertiesObjectSTIX datamodels.CommonPropertiesObjectSTIX

	for _, item := range list {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			return result, nil
		}

		switch commonPropertiesObjectSTIX.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			var ap datamodels.AttackPatternDomainObjectsSTIX
			elem, err := ap.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.AttackPatternDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})
		case "campaign":
			var c datamodels.CampaignDomainObjectsSTIX
			elem, err := c.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.CampaignDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "course-of-action":
			var ca datamodels.CourseOfActionDomainObjectsSTIX
			elem, err := ca.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.CourseOfActionDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "grouping":
			var g datamodels.GroupingDomainObjectsSTIX
			elem, err := g.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.GroupingDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "identity":
			var i datamodels.IdentityDomainObjectsSTIX
			elem, err := i.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.IdentityDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "indicator":
			var i datamodels.IndicatorDomainObjectsSTIX
			elem, err := i.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.IndicatorDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "infrastructure":
			var i datamodels.InfrastructureDomainObjectsSTIX
			elem, err := i.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.InfrastructureDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "intrusion-set":
			var is datamodels.IntrusionSetDomainObjectsSTIX
			elem, err := is.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.IntrusionSetDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "location":
			var l datamodels.LocationDomainObjectsSTIX
			elem, err := l.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.LocationDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "malware":
			var m datamodels.MalwareDomainObjectsSTIX
			elem, err := m.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.MalwareDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "malware-analysis":
			var ma datamodels.MalwareAnalysisDomainObjectsSTIX
			elem, err := ma.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.MalwareAnalysisDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "note":
			var n datamodels.NoteDomainObjectsSTIX
			elem, err := n.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.NoteDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "observed-data":
			var od datamodels.ObservedDataDomainObjectsSTIX
			elem, err := od.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ObservedDataDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "opinion":
			var o datamodels.OpinionDomainObjectsSTIX
			elem, err := o.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.OpinionDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "report":
			var r datamodels.ReportDomainObjectsSTIX
			elem, err := r.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ReportDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "threat-actor":
			var ta datamodels.ThreatActorDomainObjectsSTIX
			elem, err := ta.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ThreatActorDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "tool":
			var t datamodels.ToolDomainObjectsSTIX
			elem, err := t.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ToolDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "vulnerability":
			var v datamodels.VulnerabilityDomainObjectsSTIX
			elem, err := v.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.VulnerabilityDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

			/* *** Relationship Objects *** */
		case "relationship":
			var r datamodels.RelationshipObjectSTIX
			elem, err := r.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.RelationshipObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "sighting":
			var s datamodels.SightingObjectSTIX
			elem, err := s.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.SightingObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

			/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			var a datamodels.ArtifactCyberObservableObjectSTIX
			elem, err := a.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ArtifactCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "autonomous-system":
			var as datamodels.AutonomousSystemCyberObservableObjectSTIX
			elem, err := as.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.AutonomousSystemCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "directory":
			var d datamodels.DirectoryCyberObservableObjectSTIX
			elem, err := d.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.DirectoryCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "domain-name":
			var dn datamodels.DomainNameCyberObservableObjectSTIX
			elem, err := dn.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.DomainNameCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "email-addr":
			var ea datamodels.EmailAddressCyberObservableObjectSTIX
			elem, err := ea.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.EmailAddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "email-message":
			var em datamodels.EmailMessageCyberObservableObjectSTIX
			elem, err := em.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.EmailMessageCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "file":
			var f datamodels.FileCyberObservableObjectSTIX
			elem, err := f.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.FileCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "ipv4-addr":
			var ip4 datamodels.IPv4AddressCyberObservableObjectSTIX
			elem, err := ip4.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.IPv4AddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "ipv6-addr":
			var ip6 datamodels.IPv6AddressCyberObservableObjectSTIX
			elem, err := ip6.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.IPv6AddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "mac-addr":
			var mac datamodels.MACAddressCyberObservableObjectSTIX
			elem, err := mac.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.MACAddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "mutex":
			var m datamodels.MutexCyberObservableObjectSTIX
			elem, err := m.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.MutexCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "network-traffic":
			var nt datamodels.NetworkTrafficCyberObservableObjectSTIX
			elem, err := nt.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.NetworkTrafficCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "process":
			var p datamodels.ProcessCyberObservableObjectSTIX
			elem, err := p.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.ProcessCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "software":
			var s datamodels.SoftwareCyberObservableObjectSTIX
			elem, err := s.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.SoftwareCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "url":
			var url datamodels.URLCyberObservableObjectSTIX
			elem, err := url.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.URLCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "user-account":
			var ua datamodels.UserAccountCyberObservableObjectSTIX
			elem, err := ua.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.UserAccountCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "windows-registry-key":
			var wrk datamodels.WindowsRegistryKeyCyberObservableObjectSTIX
			elem, err := wrk.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.WindowsRegistryKeyCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		case "x509-certificate":
			var x509 datamodels.X509CertificateCyberObservableObjectSTIX
			elem, err := x509.DecoderJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(datamodels.X509CertificateCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("Error: type conversion error")
			}

			result = append(result, &datamodels.ListSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		}
	}
	return result, nil
}
