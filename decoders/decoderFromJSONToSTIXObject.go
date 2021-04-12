package decoders

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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
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
				return result, fmt.Errorf("error: type conversion error")
			}

			result = append(result, &datamodels.ElementSTIXObject{
				DataType: commonPropertiesObjectSTIX.Type,
				Data:     e,
			})

		}
	}

	return result, nil
}
