package commonlibs

import (
	"encoding/json"

	"ISEMS-MRSICT/datamodels"
)

//DecoderFromJSONToSTIXObject декодирует сообщения формата JSON в STIX объект
func DecoderFromJSONToSTIXObject(objectType string, rawMessage *json.RawMessage) (interface{}, error) {
	ListFuncDecoderFromJSONToSTIXObject := map[string]func(*json.RawMessage) (interface{}, error){
		/* *** Domain Objects STIX *** */
		"attack-pattern": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.AttackPatternDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"campaign": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.CampaignDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"course-of-action": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.CourseOfActionDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"grouping": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.GroupingDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"identity": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.IdentityDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"indicator": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.IndicatorDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"infrastructure": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.InfrastructureDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"intrusion-set": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.IntrusionSetDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"location": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.LocationDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"malware": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.MalwareDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"malware-analysis": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.MalwareAnalysisDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"note": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.NoteDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"observed-data": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.ObservedDataDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"opinion": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.OpinionDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"report": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.ReportDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"threat-actor": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.ThreatActorDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"tool": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.ToolDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"vulnerability": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.VulnerabilityDomainObjectsSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
				/* *** Relationship Objects *** */
				/*"relationship": func(msg *json.RawMessage) (interface{}, error) {
					var object datamodels.
					if err := json.Unmarshal(*msg, &object); err != nil {
						return nil, err
					}
		
					return object, nil
				},
				"sighting": func(msg *json.RawMessage) (interface{}, error) {
					var object datamodels.
					if err := json.Unmarshal(*msg, &object); err != nil {
						return nil, err
					}
		
					return object, nil
				},*/
		/* *** Cyber-observable Objects STIX *** */


// - "file"
// - "archive-ext"
// - "ntfs-ext"
// - "alternate-data-stream-type"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "windows-pe-optional-header-type"
// - "windows-pe-section-type"
// - "ipv4-addr"
// - "ipv6-addr"
// - "mac-addr"
// - "mutex"
// - "network-traffic"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "process"
// - "windows-process-ext"
// - "windows-service-ext"
// - "software"
// - "url"
// - "user-account"
// - "unix-account-ext"
// - "windows-registry-key"
// - "windows-registry-value-type"
// - "x509-certificate"
// - "x509-v3-extensions-type"
		"artifact": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.ArtifactCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"autonomous-system": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.AutonomousSystemCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"directory": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.DirectoryCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"domain-name": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.DomainNameCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"email-addr": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.EmailAddressCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"email-message": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.EmailMessageCyberObservableObjectSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"email-mime-part-type": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.EmailMIMEPartTypeSTIX
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
		"": func(msg *json.RawMessage) (interface{}, error) {
			var object datamodels.
			if err := json.Unmarshal(*msg, &object); err != nil {
				return nil, err
			}

			return object, nil
		},
	}

	return ListFuncDecoderFromJSONToSTIXObject[objectType](rawMessage)
}
