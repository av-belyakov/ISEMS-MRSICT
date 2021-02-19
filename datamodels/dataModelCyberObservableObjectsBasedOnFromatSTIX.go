package datamodels

import "time"

/********** 			Некоторые 'сложные' типы относящиеся к Cyber-observable Objects STIX 			**********/

//EmailMIMEPartTypeSTIX тип "email-mime-part-type", по терминалогии STIX, содержит один компонент тела email из нескольких частей
// Body - содержит содержимое части MIME, если content_type не указан или начинается с text/ (например, в случае обычного текста или HTML-письма)
// BodyRawRef - содержит содержимое нетекстовых частей MIME, то есть тех, чей content_type не начинается с text/, в качестве
//  ссылки на объект артефакта или Файловый объект
// ContentType - содержимое поля 'Content-Type' заголовка MIME части email
// ContentDisposition - содержимое поля 'Content-Disposition' заголовка MIME части email
type EmailMIMEPartTypeSTIX struct {
	Body               string             `json:"body" bson:"body"`
	BodyRawRef         IdentifierTypeSTIX `json:"body_raw_ref" bson:"body_raw_ref"`
	ContentType        string             `json:"content_type" bson:"content_type"`
	ContentDisposition string             `json:"content_disposition" bson:"content_disposition"`
}

/********** 			Cyber-observable Objects STIX 			**********/

//CommonPropertiesCyberObservableObjectSTIX содержит общие свойства для объекта Cyber-observable Objects STIX
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
//  или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
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
// Number - содержит номер присвоенный Автономной системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - название Автономной системы
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системыs
type AutonomousSystemObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Number int    `json:"number" bson:"number"`
	Name   string `json:"name" bson:"name"`
	RIR    string `json:"rir" bson:"rir"`
}

//DirectoryObjectSTIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
// Path - указывает путь, как было первоначально замечено, к каталогу в файловой системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PathEnc - указывает наблюдаемую кодировку для пути. Значение ДОЛЖНО быть указано, если путь хранится в кодировке, отличной от Unicode.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания директории
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации или записи в директорию
// Atime - время, в формате "2016-05-12T08:17:27.000Z", последнего обращения к директории
// ContainsRefs - содержит список файловых объектов или директорий содержащихся внутри директории
type DirectoryObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Path         string                `json:"path" bson:"path"`
	PathEnc      string                `json:"path_enc" bson:"path_enc"`
	Ctime        time.Time             `json:"ctime" bson:"ctime"`
	Mtime        time.Time             `json:"mtime" bson:"mtime"`
	Atime        time.Time             `json:"atime" bson:"atime"`
	ContainsRefs []*IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
}

//DomainNameObjectSTIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Value - сетевое доменное имя (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - список ссылок на один или несколько IP-адресов или доменных имен, на которые разрешается доменное имя
type DomainNameObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Value          string                `json:"value" bson:"value"`
	ResolvesToRefs []*IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
}

//EmailAddressObjectSTIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Value - email адрес (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// DisplayName - содержит единственное почтовое имя которое видит человек при просмотре письма
// BelongsToRef - учетная запись пользователя, которой принадлежит адрес электронной почты, в качестве ссылки на объект учетной записи пользователя
type EmailAddressObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Value        string             `json:"value" bson:"value"`
	DisplayName  string             `json:"display_name" bson:"display_name"`
	BelongsToRef IdentifierTypeSTIX `json:"belongs_to_ref" bson:"belongs_to_ref"`
}

//EmailMessageObjectSTIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
// IsMultipart - информирует содержит ли 'тело' email множественные MIME части (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Date - время, в формате "2016-05-12T08:17:27.000Z", когда email сообщение было отправлено
// ContentType - содержит содержимое 'Content-Type' заголовка email сообщения
// FromRef - содержит содержимое 'From:' заголовка email сообщения
// SenderRef - содержит содержимое поля 'Sender:' email сообщения
// ToRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'To:'
// CcRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'CC:'
// BccRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'BCC:'
// MessageID - содержимое Message-ID email сообщения
// Subject - содержит тему сообщения
// ReceivedLines - содержит одно или несколько полей заголовка 'Received', которые могут быть включены в заголовки email
// AdditionalHeaderFields - содержит любые другие поля заголовка (за исключением date, received_lines, content_type, from_ref,
//  sender_ref, to_ref, cc_ref, bcc_refs и subject), найденные в сообщении электронной почты в виде словаря
// Body - содержит тело сообщения
// BodyMultipart - содержит адает список MIME-части, которые составляют тело email. Это свойство НЕ ДОЛЖНО использоваться, если
//  is_multipart имеет значение false
// RawEmailRef - содержит 'сырое' бинарное содержимое email сообщения
type EmailMessageObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	IsMultipart            bool                     `json:"is_multipart" bson:"is_multipart"`
	Date                   time.Time                `json:"date" bson:"date"`
	ContentType            string                   `json:"content_type" bson:"content_type"`
	FromRef                IdentifierTypeSTIX       `json:"from_ref" bson:"from_ref"`
	SenderRef              IdentifierTypeSTIX       `json:"sender_ref" bson:"sender_ref"`
	ToRefs                 []*IdentifierTypeSTIX    `json:"to_refs" bson:"to_refs"`
	CcRefs                 []*IdentifierTypeSTIX    `json:"cc_refs" bson:"cc_refs"`
	BccRefs                []*IdentifierTypeSTIX    `json:"bcc_refs" bson:"bcc_refs"`
	MessageID              string                   `json:"message_id" bson:"message_id"`
	Subject                string                   `json:"subject" bson:"subject"`
	ReceivedLines          []string                 `json:"received_lines" bson:"received_lines"`
	AdditionalHeaderFields map[string]string        `json:"additional_header_fields" bson:"additional_header_fields"`
	Body                   string                   `json:"body" bson:"body"`
	BodyMultipart          []*EmailMIMEPartTypeSTIX `json:"body_multipart" bson:"body_multipart"`
	RawEmailRef            IdentifierTypeSTIX       `json:"raw_email_ref" bson:"raw_email_ref"`
}

/*
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
	`json:"" bson:""`
*/
