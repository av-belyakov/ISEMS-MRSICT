package commonlibs

import (
	"encoding/json"

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
