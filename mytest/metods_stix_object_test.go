package mytest_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ISEMS-MRSICT/datamodels"
)

var _ = Describe("MetodsStixObject", func() {

	Context("Тест 1. Выполняем проверку типа 'attack-pattern' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа AttackPatternDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			apbyte := json.RawMessage([]byte(`{
	"type": "attack-pattern",
	"spec_version": "2.1",
	"id": "attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5",
	"created": "2016-05-12T08:17:27.000Z",
	"modified": "2016-05-12T08:17:27.000Z",
	"name": "Spear Phishing as Practiced $< by Adversary\n' X",
	"description": "A particular form of spear phishing where the attacker claims that $ the target <> had won a contest, including personal details, to get them to click on a link.",
	"external_references": [{ "source_name": "capec", "external_id": "CAPEC-163" }],
	"labels": ["__lable 1", "__lable 2", "__lable$ 3 <.'"],
	"aliases": ["new ali&ase", "aliase <>", "Ali12$$"]
}`))
			var apo datamodels.AttackPatternDomainObjectsSTIX
			apoTmp, err := apo.DecoderJSON(&apbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newapo, ok := apoTmp.(datamodels.AttackPatternDomainObjectsSTIX)
			apoIsTrue := newapo.CheckingTypeFields()
			newapo = newapo.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newapo.ToStringBeautiful())

			Expect(apoIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 2. Выполняем проверку типа 'campaign' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа CampaignDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			cobyte := json.RawMessage([]byte(`{
				"type": "campaign",
				"spec_version": "2.1",
				"id": "campaign--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"modified": "2016-04-06T20:03:00.000Z",
				"name": "Green Group Attacks 'Against' Finance",
				"description": "Campaign by Green Group against a series of targets in the financialservices $sector.",
				"aliases": ["aa12$", "asd_3'", "\nerrx"],
				"first_seen": "2016-04-12T20:12:05.000Z",
				"last_seen": "2016-04-06T20:22:10.000Z",
				"objective": "Example objectiv<e!"
}`))
			var co datamodels.CampaignDomainObjectsSTIX
			coTmp, err := co.DecoderJSON(&cobyte)

			Expect(err).ShouldNot(HaveOccurred())

			newco, ok := coTmp.(datamodels.CampaignDomainObjectsSTIX)
			coIsTrue := newco.CheckingTypeFields()
			newco = newco.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newco.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 3. Выполняем проверку типа 'course-of-action' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа CourseOfActionDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			cabyte := json.RawMessage([]byte(`{
				"type": "course-of-action",
				"spec_version": "2.1",
				"id": "course-of-action--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"name": "Add TCP port 80 Filter Rule to the existing Block UDP 1434 Filter",
				"description": "This is how to add a filter rule to block inbound access to TCP port 80 tothe existing UDP 1434 filter ..."
}`))
			var ca datamodels.CourseOfActionDomainObjectsSTIX
			caTmp, err := ca.DecoderJSON(&cabyte)

			Expect(err).ShouldNot(HaveOccurred())

			newca, ok := caTmp.(datamodels.CourseOfActionDomainObjectsSTIX)
			coIsTrue := newca.CheckingTypeFields()
			newca = newca.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newca.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 4. Выполняем проверку типа 'grouping' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа GroupingDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			gdbyte := json.RawMessage([]byte(`{
				"type": "grouping",
				"spec_version": "2.1",
				"id": "grouping--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3",
				"created_by_ref": "identity--a463ffb3-1bd9-4d94-b02d-74e4f1658283",
				"created": "2015-12-21T19:59:11.000Z",
				"modified": "2015-12-21T19:59:11.000Z",
				"name": "The Black Vine Cyberespionage Group",
				"description": "A simple collection of Black Vine Cyberespionage Group attributed intel",
				"context": "suspicious-activity",
				"object_refs": [
					"indicator--26ffb872-1dd9-446e-b6f5-d58527e5b5d2",
					"campaign--83422c77-904c-4dc1-aff5-5c38f3a2c55c",
					"relationship--f82356ae-fe6c-437c-9c24-6b64314ae68a",
					"file--0203b5c8-f8b6-4ddb-9ad0-527d727f968b"
					]
}`))
			var gd datamodels.GroupingDomainObjectsSTIX
			gdTmp, err := gd.DecoderJSON(&gdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newgd, ok := gdTmp.(datamodels.GroupingDomainObjectsSTIX)
			coIsTrue := newgd.CheckingTypeFields()
			newgd = newgd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newgd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 5. Выполняем проверку типа 'identity' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IdentityDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
				"type": "identity",
				"spec_version": "2.1",
				"id": "identity--023d105b-752e-4e3c-941c-7d3f3cb15e9e",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:00.000Z",
				"modified": "2016-04-06T20:03:00.000Z",
				"name": "John Smith",
				"identity_class": "individual"
}`))
			var id datamodels.IdentityDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.IdentityDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 6. Выполняем проверку типа 'indicator' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IndicatorDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
				"type": "indicator",
				"spec_version": "2.1",
				"id": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"indicator_types": ["malicious-activity"],
				"name": "Poison Ivy Malware",
				"description": "This file is part of Poison Ivy",
				"pattern": "[ file:hashes.'SHA-256' = '4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877' ]",
				"pattern_type": "stix",
				"valid_from": "2016-01-01T00:00:00Z"
}`))
			var id datamodels.IndicatorDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.IndicatorDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 7. Выполняем проверку типа 'infrastructure' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа InfrastructureDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			idbyte := json.RawMessage([]byte(`{
				"type":"infrastructure",
				"spec_version": "2.1",
				"id":"infrastructure--38c47d93-d984-4fd9-b87b-d69d0841628d",
				"created":"2016-05-07T11:22:30.000Z",
				"modified":"2016-05-07T11:22:30.000Z",
				"name":"Poison Ivy C2",
				"infrastructure_types": ["command-and-control"]
}`))
			var id datamodels.InfrastructureDomainObjectsSTIX
			idTmp, err := id.DecoderJSON(&idbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newid, ok := idTmp.(datamodels.InfrastructureDomainObjectsSTIX)
			coIsTrue := newid.CheckingTypeFields()
			newid = newid.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newid.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 8. Выполняем проверку типа 'infrastructure' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа IntrusionSetDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			isdbyte := json.RawMessage([]byte(`{
				"type": "intrusion-set",
				"spec_version": "2.1",
				"id": "intrusion-set--4e78f46f-a023-4e5f-bc24-71b3ca22ec29",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"name": "Bobcat Breakin",
				"description": "Incidents usually feature a shared TTP of a bobcat being released within the building containing network access, scaring users to leave their computers without locking them first. Still determining where the threat actors are getting the bobcats.",
				"aliases": ["Zookeeper"],
				"goals": ["acquisition-theft", "harassment", "damage"]
}`))
			var isd datamodels.IntrusionSetDomainObjectsSTIX
			isdTmp, err := isd.DecoderJSON(&isdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newisd, ok := isdTmp.(datamodels.IntrusionSetDomainObjectsSTIX)
			coIsTrue := newisd.CheckingTypeFields()
			newisd = newisd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newisd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 9. Выполняем проверку типа 'location' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа LocationDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			ldbyte := json.RawMessage([]byte(`{
				"type": "location",
				"spec_version": "2.1",
				"id": "location--a6e9345f-5a15-4c29-8bb3-7dcc5d168d64",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:00.000Z",
				"modified": "2016-04-06T20:03:00.000Z",
				"latitude": 48.8566,
				"longitude": 2.3522,
				"region": "south-eastern-asia",
				"country": "th",
				"administrative_area": "Tak",
				"postal_code": "63170"
}`))
			var ld datamodels.LocationDomainObjectsSTIX
			ldTmp, err := ld.DecoderJSON(&ldbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newld, ok := ldTmp.(datamodels.LocationDomainObjectsSTIX)
			coIsTrue := newld.CheckingTypeFields()
			newld = newld.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newld.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 10. Выполняем проверку типа 'malware' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа MalwareDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "malware",
				"spec_version": "2.1",
				"id": "malware--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"name": "Cryptolocker",
				"description": "A variant of the cryptolocker family",
				"malware_types": ["ransomware"],
				"is_family": false    
}`))
			var md datamodels.MalwareDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.MalwareDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Println(newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 11. Выполняем проверку типа 'malware-analysis' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа MalwareAnalysisDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "malware-analysis",
				"spec_version": "2.1",
				"id": "malware-analysis--d25167b7-fed0-4068-9ccd-a73dd2c5b07c",
				"created": "2020-01-16T18:52:24.277Z",
				"modified": "2020-01-16T18:52:24.277Z",
				"product": "microsoft",
				"analysis_engine_version": "5.1.0",
				"analysis_definition_version": "053514-0062",
				"analysis_started": "2012-02-11T08:36:14Z",
				"analysis_ended": "2012-02-11T08:36:14Z",
				"av_result": "malicious"
}`))
			var md datamodels.MalwareAnalysisDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.MalwareAnalysisDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 12. Выполняем проверку типа 'note' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа NoteDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "note",
				"spec_version": "2.1",
				"id": "note--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"external_references": [
					{
						"source_name": "job-tracker",
						"external_id": "job-id-1234"
					}
				],
				"abstract": "Tracking Team Note#1",
				"content": "This note indicates the various steps taken by the threat analyst team to investigate this specific campaign. Step 1) Do a scan 2) Review scanned results for identified hosts not known by external intel...etc.",
				"authors": ["John Doe"],
				"object_refs": ["campaign--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f"]
}`))
			var md datamodels.NoteDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.NoteDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 13. Выполняем проверку типа 'observed-data' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа ObservedDataDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "observed-data",
				"spec_version": "2.1",
				"id": "observed-data--b67d30ff-02ac-498a-92f9-32f845f448cf",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T19:58:16.000Z",
				"modified": "2016-04-06T19:58:16.000Z",
				"first_observed": "2015-12-21T19:00:00Z",
				"last_observed": "2015-12-21T19:00:00Z",
				"number_observed": 50,
				"object_refs": [
					"ipv4-address--efcd5e80-570d-4131-b213-62cb18eaa6a8",
					"domain-name--ecb120bf-2694-4902-a737-62b74539a41b"
				]
}`))
			var md datamodels.ObservedDataDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ObservedDataDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 14. Выполняем проверку типа 'opinion' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа OpinionDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "opinion",
				"spec_version": "2.1",
				"id": "opinion--b01efc25-77b4-4003-b18b-f6e24b5cd9f7",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"object_refs": ["relationship--16d2358f-3b0d-4c88-b047-0da2f7ed4471"],
				"opinion": "strongly-disagree",
				"explanation": "This doesn't seem like it is feasible. We've seen how PandaCat has attacked Spanish infrastructure over the last 3 years, so this change in targeting seems too great to be viable. The methods used are more commonly associated with the FlameDragonCrew."	
}`))
			var md datamodels.OpinionDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.OpinionDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 15. Выполняем проверку типа 'report' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа ReportDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "report",
				"spec_version": "2.1",
				"id": "report--84e4d88f-44ea-4bcd-bbf3-b2c1c320bcb3",
				"created_by_ref": "identity--a463ffb3-1bd9-4d94-b02d-74e4f1658283",
				"created": "2015-12-21T19:59:11.000Z",
				"modified": "2015-12-21T19:59:11.000Z",
				"name": "The Black Vine Cyberespionage Group",
				"description": "A simple report with an indicator and campaign",
				"published": "2016-01-20T17:00:00.000Z",
				"report_types": ["campaign"],
				"object_refs": [
					"indicator--26ffb872-1dd9-446e-b6f5-d58527e5b5d2",
					"campaign--83422c77-904c-4dc1-aff5-5c38f3a2c55c",
					"relationship--f82356ae-fe6c-437c-9c24-6b64314ae68a"
				]
}`))
			var md datamodels.ReportDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ReportDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 16. Выполняем проверку типа 'threat-actor' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа ThreatActorDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "threat-actor",
				"spec_version": "2.1",
				"id": "threat-actor--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"threat_actor_types": [ "crime-syndicate"],
				"name": "Evil Org",
				"description": "The Evil Org threat actor group",
				"aliases": ["Syndicate 1", "Evil Syndicate 99"],
				"roles": ["director"],
				"goals": ["Steal bank money", "Steal credit cards"],
				"sophistication": "advanced",
				"resource_level": "team",
				"primary_motivation": "organizational-gain"
}`))
			var md datamodels.ThreatActorDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ThreatActorDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 17. Выполняем проверку типа 'tool' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа ToolDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "tool",
				"spec_version": "2.1",
				"id": "tool--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:03:48.000Z",
				"modified": "2016-04-06T20:03:48.000Z",
				"tool_types": [ "remote-access"],
				"name": "VNC"  
}`))
			var md datamodels.ToolDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ToolDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 18. Выполняем проверку типа 'vulnerability' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа VulnerabilityDomainObjectsSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "vulnerability",
				"spec_version": "2.1",
				"id": "vulnerability--0c7b5b88-8ff7-4a4d-aa9d-feb398cd0061",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"name": "CVE-2016-1234",
				"external_references": [
					{
						"source_name": "cve",
						"external_id": "CVE-2016-1234"
					}
				]
}`))
			var md datamodels.VulnerabilityDomainObjectsSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.VulnerabilityDomainObjectsSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 19. Выполняем проверку типа 'relationship' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа RelationshipObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "relationship",
				"spec_version": "2.1",
				"id": "relationship--57b56a43-b8b0-4cba-9deb-34e3e1faed9e",
				"created": "2016-05-12T08:17:27.000Z",
				"modified": "2016-05-12T08:17:27.000Z",
				"relationship_type": "uses",
				"source_ref": "intrusion-set--0c7e22ad-b099-4dc3-b0df-2ea3f49ae2e6",
				"target_ref": "attack-pattern--7e33a43e-e34b-40ec-89da-36c9bb2cacd5"
}`))
			var md datamodels.RelationshipObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.RelationshipObjectSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 20. Выполняем проверку типа 'sighting' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа SightingObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "sighting",
				"spec_version": "2.1",
				"id": "sighting--ee20065d-2555-424f-ad9e-0f8428623c75",
				"created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
				"created": "2016-04-06T20:08:31.000Z",
				"modified": "2016-04-06T20:08:31.000Z",
				"first_seen": "2015-12-21T19:00:00Z",
				"last_seen": "2015-12-21T19:00:00Z",
				"count": 50,
				"sighting_of_ref": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
				"observed_data_refs": ["observed-data--b67d30ff-02ac-498a-92f9-32f845f448cf"],
				"where_sighted_refs": ["identity--b67d30ff-02ac-498a-92f9-32f845f448ff"]
}`))
			var md datamodels.SightingObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.SightingObjectSTIX)
			coIsTrue := newmd.CheckingTypeFields()
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(coIsTrue).Should(BeTrue())
		})
	})

	Context("Тест 21. Выполняем проверку типа 'artifact' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа ArtifactCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "artifact",
				"spec_version": "2.1",
				"id": "artifact--6f437177-6e48-5cf8-9d9e-872a2bddd641",
				"mime_type": "application/zip",
				"payload_bin": "J25mIG50Y25qZGZ6IGNuaGZyZiBkIGJhc2U2NCE=",
				"url": "https://pkg.go.dev/github.com/asaskevich/govalidator#IsURL",
				"encryption_algorithm": "mime-type-indicated",
				"decryption_key": "My voice is my passport"
}`))
			var md datamodels.ArtifactCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.ArtifactCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 21. Выполняем проверку типа 'autonomous-system' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа AutonomousSystemCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "autonomous-system",
				"spec_version": "2.1",
				"id": "autonomous-system--f720c34b-98ae-597f-ade5-27dc241e8c74",
				"number": 15139,
				"name": "Slime Industries",
				"rir": "ARIN"           
}`))
			var md datamodels.AutonomousSystemCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.AutonomousSystemCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 22. Выполняем проверку типа 'directory' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа DirectoryCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "directory",
				"spec_version": "2.1",
				"id": "directory--93c0a9b0-520d-545d-9094-1a08ddf46b05",
				"path": "C:\\Windows\\System32"         
}`))
			var md datamodels.DirectoryCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.DirectoryCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 23. Выполняем проверку типа 'domain-name' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа DomainNameCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "domain-name",
				"spec_version": "2.1",
				"id": "domain-name--3c10e93f-798e-5a26-a0c1-08156efab7f5",
				"value": "example.com",
				"resolves_to_refs": ["ipv4-addr--ff26c055-6336-5bc5-b98d-13d6226742dd"]
}`))
			var md datamodels.DomainNameCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.DomainNameCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 24. Выполняем проверку типа 'email-addr' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа EmailAddressCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "email-addr",
				"spec_version": "2.1",
				"id": "email-addr--2d77a846-6264-5d51-b586-e43822ea1ea3",
				"value": "john@example.com",
				"display_name": "John Doe"           
}`))
			var md datamodels.EmailAddressCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.EmailAddressCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 25. Выполняем проверку типа 'email-message' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа EmailMessageCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "email-message",
				"spec_version": "2.1",
				"id": "email-message--cf9b4b7f-14c8-5955-8065-020e0316b559",
				"is_multipart": true,
				"received_lines": [
					"from mail.example.com ([198.51.100.3]) by smtp.gmail.com with ESMTPSA id q23sm23309939wme.17.2016.07.19.07.20.32 (version=TLS1_2 cipher=ECDHE-RSA-AES128-GCM-SHA256 bits=128/128); Tue, 19 Jul 2016 07:20:40 -0700 (PDT)"
				],
				"content_type": "multipart/mixed",
				"date": "2016-06-19T14:20:40.000Z",
				"from_ref": "email-addr--89f52ea8-d6ef-51e9-8fce-6a29236436ed",
				"to_refs": ["email-addr--d1b3bf0c-f02a-51a1-8102-11aba7959868"],
				"cc_refs": ["email-addr--e4ee5301-b52d-59cd-a8fa-8036738c7194"],
				"subject": "Check out this picture of a cat!",
				"additional_header_fields": {
					"Content-Disposition": "inline",
					"X-Mailer": "Mutt/1.5.23",
					"X-Originating-IP": "198.51.100.3"
				},
				"body_multipart": [
					{
						"content_type": "text/plain; charset=utf-8",
						"content_disposition": "inline",
						"body": "Cats are funny!"
					},
					{
						"content_type": "image/png",
						"content_disposition": "attachment; filename=\"tabby.png\"",
						"body_raw_ref": "artifact--4cce66f8-6eaa-53cb-85d5-3a85fca3a6c5"
					},
					{
						"content_type": "application/zip",
						"content_disposition": "attachment; filename=\"tabby_pics.zip\"",
						"body_raw_ref": "file--6ce09d9c-0ad3-5ebf-900c-e3cb288955b5"
					}
				]
}`))
			var md datamodels.EmailMessageCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.EmailMessageCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			//fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})

	Context("Тест 25. Выполняем проверку типа 'file' методом CheckingTypeFields", func() {
		It("На валидное содержимое типа FileCyberObservableObjectSTIX должно быть TRUE, ошибки при декодировании быть не должно", func() {
			mdbyte := json.RawMessage([]byte(`{
				"type": "file",
				"spec_version": "2.1",
				"id": "file--fb0419a8-f09c-57f8-be64-71a80417591c",
				"extensions": {
					"windows-pebinary-ext": {
						"pe_type": "exe",
						"machine_hex": "014c",
						"number_of_sections": 4,
						"time_date_stamp": "2016-01-22T12:31:12Z",
						"pointer_to_symbol_table_hex": "74726144",
						"number_of_symbols": 4542568,
						"size_of_optional_header": 224,
						"characteristics_hex": "818f",
						"optional_header": {
							"magic_hex": "010b",
							"major_linker_version": 2,
							"minor_linker_version": 25,
							"size_of_code": 512,
							"size_of_initialized_data": 283648,
							"size_of_uninitialized_data": 0,
							"address_of_entry_point": 4096,
							"base_of_code": 4096,
							"base_of_data": 8192,
							"image_base": 14548992,
							"section_alignment": 4096,
							"file_alignment": 4096,
							"major_os_version": 1,
							"minor_os_version": 0,
							"major_image_version": 0,
							"minor_image_version": 0,
							"major_subsystem_version": 4,
							"minor_subsystem_version": 0,
							"win32_version_value_hex": "00",
							"size_of_image": 299008,
							"size_of_headers": 4096,
							"checksum_hex": "00",
							"subsystem_hex": "03",
							"dll_characteristics_hex": "00",
							"size_of_stack_reserve": 100000,
							"size_of_stack_commit": 8192,
							"size_of_heap_reserve": 100000,
							"size_of_heap_commit": 4096
						},
						"loader_flags_hex": "abdbffde",
						"number_of_rva_and_sizes": 3758087646,
						"sections": [
							{
								"name": "CODE",
								"entropy": 0.061089
							},
							{
								"name": "DATA",
								"entropy": 7.980693
							},
							{
								"name": "NicolasB",
								"entropy": 0.607433
							},
							{
								"name": ".idata",
								"entropy": 0.607433
							}
						]
					}
				}
}`))
			var md datamodels.FileCyberObservableObjectSTIX
			mdTmp, err := md.DecoderJSON(&mdbyte)

			Expect(err).ShouldNot(HaveOccurred())

			newmd, ok := mdTmp.(datamodels.FileCyberObservableObjectSTIX)
			newmd = newmd.SanitizeStruct()

			Expect(ok).Should(BeTrue())

			fmt.Printf("\n%v\n", newmd.ToStringBeautiful())

			Expect(newmd.CheckingTypeFields()).Should(BeTrue())
		})
	})
})
