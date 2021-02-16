package datamodels

/********** 			Cyber-observable Objects STIX 			**********/

//CommonPropertiesCyberObservableObjectSTIX содержит общие свойства для объекта Cyber-observable Objects
// Type - наименование типа шаблона (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
//  Type должен содержать одно из следующих значений:
// - "artifact"
// - "autonomous-system"
// - "directory"
// - "domain-name"
// - "email-addr"
// - "email-message"
// - "email-mime-part-type"
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
// ID - уникальный идентификатор объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type CommonPropertiesCyberObservableObjectSTIX struct {
	Type string `json:"type" bson:"type"`
	ID   string `json:"id" bson:"id"`
}

//ArtifactObjectSTIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки в кодировке base64
// или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
// MimeType - по возможности это значение ДОЛЖНО быть одним из значений, определенных в реестре типов носителей IANA. В универсальном каталоге
//  всех существующих типов файлов.
// PayloadBin - бинарные данные в base64
// URL - унифицированный указатель ресурса (URL)
// Hashes - словарь хешей для URL или PayloadBin
// EncryptionAlgorithm - тип алгоритма шифрования для бинарных данных
// DecryptionKey - определяет ключ для дешифрования зашифрованных данных
type ArtifactObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	MimeType            string         `json:"mime_type" bson:"mime_type"`
	PayloadBin          string         `json:"payload_bin" bson:"payload_bin"`
	URL                 string         `json:"url" bson:"url"`
	Hashes              HashesTypeSTIX `json:"hashes" bson:"hashes"`
	EncryptionAlgorithm EnumTypeSTIX   `json:"encryption_algorithm" bson:"encryption_algorithm"`
	DecryptionKey       string         `json:"decryption_key" bson:"decryption_key"`
}

//AutonomousSystemObjectSTIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Number -
// Name - название Автономной системы
// RIR -
type AutonomousSystemObjectSTIX struct {
	Number int    `json:"number" bson:"number"`
	Name   string `json:"name" bson:"name"`
	RIR    string `json:"rir" bson:"rir"`
}

/*
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
*/
