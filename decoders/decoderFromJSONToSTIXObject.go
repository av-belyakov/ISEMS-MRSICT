package decoders

import (
	"encoding/json"
	"fmt"
	"strings"

	"ISEMS-MRSICT/datamodels"

	mstixo "github.com/av-belyakov/methodstixobjects"
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
/*func decodingExtensionsSTIX(extType string, rawMsg *json.RawMessage) (interface{}, error) {
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
}*/

func GetListSTIXObjectFromJSON(list []*json.RawMessage) ([]*datamodels.ElementSTIXObject, error) {
	var result []*datamodels.ElementSTIXObject = make([]*datamodels.ElementSTIXObject, 0, len(list))
	var commonPropertiesObjectSTIX mstixo.CommonPropertiesObjectSTIX

	for _, item := range list {
		err := json.Unmarshal(*item, &commonPropertiesObjectSTIX)
		if err != nil {
			return result, nil
		}

		switch commonPropertiesObjectSTIX.Type {
		/* *** Domain Objects STIX *** */
		case "attack-pattern":
			var ap datamodels.AttackPatternDomainObjectsSTIX
			elem, err := ap.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.AttackPatternDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'attack-pattern' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.AttackPatternDomainObjectsSTIX{AttackPatternDomainObjectsSTIX: e},
			})
		case "campaign":
			var c datamodels.CampaignDomainObjectsSTIX
			elem, err := c.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.CampaignDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'campaign' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.CampaignDomainObjectsSTIX{CampaignDomainObjectsSTIX: e},
			})

		case "course-of-action":
			var ca datamodels.CourseOfActionDomainObjectsSTIX
			elem, err := ca.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.CourseOfActionDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'course-of-action' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.CourseOfActionDomainObjectsSTIX{CourseOfActionDomainObjectsSTIX: e},
			})

		case "grouping":
			var g datamodels.GroupingDomainObjectsSTIX
			elem, err := g.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.GroupingDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'grouping' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.GroupingDomainObjectsSTIX{GroupingDomainObjectsSTIX: e},
			})

		case "identity":
			var i datamodels.IdentityDomainObjectsSTIX
			elem, err := i.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.IdentityDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'identity' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.IdentityDomainObjectsSTIX{IdentityDomainObjectsSTIX: e},
			})

		case "indicator":
			var i datamodels.IndicatorDomainObjectsSTIX
			elem, err := i.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.IndicatorDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'indicator' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.IndicatorDomainObjectsSTIX{IndicatorDomainObjectsSTIX: e},
			})

		case "infrastructure":
			var i datamodels.InfrastructureDomainObjectsSTIX
			elem, err := i.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.InfrastructureDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'infrastructure' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.InfrastructureDomainObjectsSTIX{InfrastructureDomainObjectsSTIX: e},
			})

		case "intrusion-set":
			var is datamodels.IntrusionSetDomainObjectsSTIX
			elem, err := is.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.IntrusionSetDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'intrusion-set' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.IntrusionSetDomainObjectsSTIX{IntrusionSetDomainObjectsSTIX: e},
			})

		case "location":
			var l datamodels.LocationDomainObjectsSTIX
			elem, err := l.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.LocationDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'location' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.LocationDomainObjectsSTIX{LocationDomainObjectsSTIX: e},
			})

		case "malware":
			var m datamodels.MalwareDomainObjectsSTIX
			elem, err := m.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.MalwareDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'malware' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.MalwareDomainObjectsSTIX{MalwareDomainObjectsSTIX: e},
			})

		case "malware-analysis":
			var ma datamodels.MalwareAnalysisDomainObjectsSTIX
			elem, err := ma.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.MalwareAnalysisDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'malware-analysis' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.MalwareAnalysisDomainObjectsSTIX{MalwareAnalysisDomainObjectsSTIX: e},
			})

		case "note":
			var n datamodels.NoteDomainObjectsSTIX
			elem, err := n.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.NoteDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'note' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.NoteDomainObjectsSTIX{NoteDomainObjectsSTIX: e},
			})

		case "observed-data":
			var od datamodels.ObservedDataDomainObjectsSTIX
			elem, err := od.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ObservedDataDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'observed-data' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.ObservedDataDomainObjectsSTIX{ObservedDataDomainObjectsSTIX: e},
			})

		case "opinion":
			var o datamodels.OpinionDomainObjectsSTIX
			elem, err := o.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.OpinionDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'opinion' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.OpinionDomainObjectsSTIX{OpinionDomainObjectsSTIX: e},
			})

		case "report":
			var r datamodels.ReportDomainObjectsSTIX
			elem, err := r.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ReportDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'report' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     changeReportTypes(datamodels.ReportDomainObjectsSTIX{ReportDomainObjectsSTIX: e}),
			})

		case "threat-actor":
			var ta datamodels.ThreatActorDomainObjectsSTIX
			elem, err := ta.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ThreatActorDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'threat-actor' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.ThreatActorDomainObjectsSTIX{ThreatActorDomainObjectsSTIX: e},
			})

		case "tool":
			var t datamodels.ToolDomainObjectsSTIX
			elem, err := t.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ToolDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'tool' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.ToolDomainObjectsSTIX{ToolDomainObjectsSTIX: e},
			})

		case "vulnerability":
			var v datamodels.VulnerabilityDomainObjectsSTIX
			elem, err := v.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.VulnerabilityDomainObjectsSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'vulnerability' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.VulnerabilityDomainObjectsSTIX{VulnerabilityDomainObjectsSTIX: e},
			})

		/* *** Relationship Objects *** */
		case "relationship":
			var r datamodels.RelationshipObjectSTIX
			elem, err := r.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.RelationshipObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'relationship' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.RelationshipObjectSTIX{RelationshipObjectSTIX: e},
			})

		case "sighting":
			var s datamodels.SightingObjectSTIX
			elem, err := s.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.SightingObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'sighting' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.SightingObjectSTIX{SightingObjectSTIX: e},
			})

		/* *** Cyber-observable Objects STIX *** */
		case "artifact":
			var a datamodels.ArtifactCyberObservableObjectSTIX
			elem, err := a.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ArtifactCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'artifact' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.ArtifactCyberObservableObjectSTIX{ArtifactCyberObservableObjectSTIX: e},
			})

		case "autonomous-system":
			var as datamodels.AutonomousSystemCyberObservableObjectSTIX
			elem, err := as.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.AutonomousSystemCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'autonomous-system' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.AutonomousSystemCyberObservableObjectSTIX{AutonomousSystemCyberObservableObjectSTIX: e},
			})

		case "directory":
			var d datamodels.DirectoryCyberObservableObjectSTIX
			elem, err := d.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.DirectoryCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'directory' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.DirectoryCyberObservableObjectSTIX{DirectoryCyberObservableObjectSTIX: e},
			})

		case "domain-name":
			var dn datamodels.DomainNameCyberObservableObjectSTIX
			elem, err := dn.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.DomainNameCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'domain-name' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.DomainNameCyberObservableObjectSTIX{DomainNameCyberObservableObjectSTIX: e},
			})

		case "email-addr":
			var ea datamodels.EmailAddressCyberObservableObjectSTIX
			elem, err := ea.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.EmailAddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'email-addr' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.EmailAddressCyberObservableObjectSTIX{EmailAddressCyberObservableObjectSTIX: e},
			})

		case "email-message":
			var em datamodels.EmailMessageCyberObservableObjectSTIX
			elem, err := em.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.EmailMessageCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'email-message' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.EmailMessageCyberObservableObjectSTIX{EmailMessageCyberObservableObjectSTIX: e},
			})

		case "file":
			var f datamodels.FileCyberObservableObjectSTIX
			elem, err := f.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.FileCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'file' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.FileCyberObservableObjectSTIX{FileCyberObservableObjectSTIX: e},
			})

		case "ipv4-addr":
			var ip4 datamodels.IPv4AddressCyberObservableObjectSTIX
			elem, err := ip4.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.IPv4AddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'ipv4-addr' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.IPv4AddressCyberObservableObjectSTIX{IPv4AddressCyberObservableObjectSTIX: e},
			})

		case "ipv6-addr":
			var ip6 datamodels.IPv6AddressCyberObservableObjectSTIX
			elem, err := ip6.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.IPv6AddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'ipv6-addr' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.IPv6AddressCyberObservableObjectSTIX{IPv6AddressCyberObservableObjectSTIX: e},
			})

		case "mac-addr":
			var mac datamodels.MACAddressCyberObservableObjectSTIX
			elem, err := mac.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.MACAddressCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'mac-addr' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.MACAddressCyberObservableObjectSTIX{MACAddressCyberObservableObjectSTIX: e},
			})

		case "mutex":
			var m datamodels.MutexCyberObservableObjectSTIX
			elem, err := m.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.MutexCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'mutex' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.MutexCyberObservableObjectSTIX{MutexCyberObservableObjectSTIX: e},
			})

		case "network-traffic":
			var nt datamodels.NetworkTrafficCyberObservableObjectSTIX
			elem, err := nt.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.NetworkTrafficCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'network-traffic' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.NetworkTrafficCyberObservableObjectSTIX{NetworkTrafficCyberObservableObjectSTIX: e},
			})

		case "process":
			var p datamodels.ProcessCyberObservableObjectSTIX
			elem, err := p.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.ProcessCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'process' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.ProcessCyberObservableObjectSTIX{ProcessCyberObservableObjectSTIX: e},
			})

		case "software":
			var s datamodels.SoftwareCyberObservableObjectSTIX
			elem, err := s.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.SoftwareCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'software' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.SoftwareCyberObservableObjectSTIX{SoftwareCyberObservableObjectSTIX: e},
			})

		case "url":
			var url datamodels.URLCyberObservableObjectSTIX
			elem, err := url.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.URLCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'url' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.URLCyberObservableObjectSTIX{URLCyberObservableObjectSTIX: e},
			})

		case "user-account":
			var ua datamodels.UserAccountCyberObservableObjectSTIX
			elem, err := ua.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.UserAccountCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'user-account' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.UserAccountCyberObservableObjectSTIX{UserAccountCyberObservableObjectSTIX: e},
			})

		case "windows-registry-key":
			var wrk datamodels.WindowsRegistryKeyCyberObservableObjectSTIX
			elem, err := wrk.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.WindowsRegistryKeyCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'windows-registry-key' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.WindowsRegistryKeyCyberObservableObjectSTIX{WindowsRegistryKeyCyberObservableObjectSTIX: e},
			})

		case "x509-certificate":
			var x509 datamodels.X509CertificateCyberObservableObjectSTIX
			elem, err := x509.DecodeJSON(item)
			if err != nil {
				return result, err
			}

			e, ok := elem.(mstixo.X509CertificateCyberObservableObjectSTIX)
			if !ok {
				return result, fmt.Errorf("error: type 'x509-certificate' conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     datamodels.X509CertificateCyberObservableObjectSTIX{X509CertificateCyberObservableObjectSTIX: e},
			})

		}
	}

	return result, nil
}

func changeReportTypes(rdostix datamodels.ReportDomainObjectsSTIX) datamodels.ReportDomainObjectsSTIX {
	var reportTypeOV = []string{
		"attack-pattern",
		"campaign",
		"identity",
		"indicator",
		"intrusion-set",
		"malware",
		"observed-data",
		"threat-actor",
		"threat-report",
		"tool",
		"vulnerability",
	}

	newrdostix := datamodels.ReportDomainObjectsSTIX{
		ReportDomainObjectsSTIX: mstixo.ReportDomainObjectsSTIX{
			CommonPropertiesObjectSTIX:       rdostix.CommonPropertiesObjectSTIX,
			CommonPropertiesDomainObjectSTIX: rdostix.CommonPropertiesDomainObjectSTIX,
			Name:                             rdostix.Name,
			Description:                      rdostix.Description,
			Published:                        rdostix.Published,
			OutsideSpecification:             rdostix.OutsideSpecification,
		},
	}
	newrdostix.ObjectRefs = append(newrdostix.ObjectRefs, rdostix.ObjectRefs...)

	for _, v := range newrdostix.ObjectRefs {
		typeName := strings.Split(string(v), "--")[0]

		for _, tn := range reportTypeOV {
			if typeName == tn {
				var isExist bool
				for _, nrt := range newrdostix.ReportTypes {
					if string(nrt) == typeName {
						isExist = true
					}
				}

				if !isExist {
					newrdostix.ReportTypes = append(newrdostix.ReportTypes, mstixo.OpenVocabTypeSTIX(typeName))
				}
			}
		}
	}

	return newrdostix
}
