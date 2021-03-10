package datamodels

import "time"

=======
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
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системы
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
// ProcessObjectSTIX объект "Process Object", по терминологии STIX.
// ProcessObjectSTIX представляет основные свойства экземпляра приложения запущенного в контексте операционой системы.
// В ProcessObjectSTIX должно быть поределено  хотя бы одно из его свойств помимо type .
// Extentions - Объект Process определяет следующие расширения (windows-process-ext, windows-service-ext).
// 				В дополнение к ним производители МОГУТ создавать свои собственные. Ключи словаря ДОЛЖНЫ идентифицировать тип расширения по имени.
// ISHidden - Указывает, является ли процесс скрытым.
// PID - Указывает идентификатор или PID процесса.
// CreatedTime - Указывает дату и время создания процесса.
// CWD - Указывает текущий рабочий каталог процесса.
// CommandLine - Указывает полную командную строку, используемую при выполнении процесса, включая имя процесса (которое может быть
//			     указано индивидуально через image_ref.name  cвойство) и любые аргументы.
// EnvironmentVariables - Задает список переменных среды, связанных с процессом в виде словаря. Каждый ключ в словаре ДОЛЖЕН быть сохраненной в регистре версией
//                        имени переменной окружения, и каждое соответствующее значение ДОЛЖНО быть значением переменной окружения в виде строки.
// OpenedConnectionRefs - Указывает список сетевых подключений, открытых процессом, в качестве ссылки на один или несколько объектов сетевого трафика.
//                        Объекты, на который в этом списке хранятся ссылки должены быть типа NetworkTrafficObjectSTIX.
// CreatorUserRef - Указывает пользователя, создавшего процесс, в качестве ссылки на объект учетной записи пользователя.
//                  Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип UserAccountObjectSTIX.
// ImageRef - Указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект.
//            Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип FileObjectSTIX.
// ParentRef - Указывает другой процесс, который породил (т. е. является родителем) этот процесс,
//             как ссылку на объект процесса. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип ProcessObjectSTIX.
// ChildRefs - Указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других объектов процесса.
//             Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип ProcessObjectSTIX.
type ProcessObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Extentions           ProcessObjectExtensionsSTIX `json:"extensions" bson:"extensions"`
	ISHidden             bool                        `json:"is_hidden" bson:"is_hidden"`
	PID                  int                         `json:"pid" bson:"pid"`
	CreatedTime          time.Time                   `json:"created_time" bson:"created_time"`
	CWD                  string                      `json:"cwd" bson:"cwd"`
	CommandLine          string                      `json:"command_line" bson:"command_line"`
	EnvironmentVariables map[string]string           `json:"environment_variables" bson:"environment_variables"`
	OpenedConnectionRefs []string                    `json:"opened_connection_refs" bson:"opened_connection_refs"`
	CreatorUserRef       string                      `json:"creator_user_ref" bson:"creator_user_ref"`
	ImageRef             string                      `json:"image_ref" bson:"image_ref"`
	ParentRef            string                      `json:"parent_ref" bson:"parent_ref"`
	ChildRefs            string                      `json:"child_refs" bson:"child_refs"`
}

type ProcessObjectExtensionsSTIX struct {
	WindowsProcessExtSTIX
	WindowsServiceExtSTIX
}

// WindowsProcessExtSTIX  объект "Windows Process Extension", по терминологии STIX.
// WindowsProcessExtSTIX задает расширение по умолчанию для захвата свойств, специфичных для процессов Windows.
// Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-process-ext.
// Объект, использующий расширение процесса Windows, должен содержать хотя бы одно свойство из этого расширения.
// AslrEnabled - Указывает если для этого процесса включен, Address Space Layout Randomization.
// DepEnabled - Указывает, включено ли для процесса Data Execution Prevention (DEP).
// Priority - Указывает текущий класс приоритета процесса в Windows. Это значение должно быть строкой, которая заканчивается на "_CLASS"
// OwnerSid - Указывает значение идентификатора безопасности (SID) владельца процесса.
// WindowTitle - Задает заголовок главного окна процесса.
// StartupInfo - Указывает структуру STARTUP_INFO, используемую процессом в качестве словаря. Каждая пара ключ/значение в STARTUP_INFO ДОЛЖНА быть представлена как пара ключ/значение в словаре, где каждый ключ ДОЛЖЕН быть сохраненной в регистре версией исходного имени.
//               Например, учитывая имя "lpDesktop", соответствующий ключ будет lpDesktop. Тип занчений в оригенале данной структуры разнороден ...
// IntegrityLevel - Указывает уровень целостности Windows или надежность процесса. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-integrity-level-enum.
type WindowsProcessExtSTIX struct {
	AslrEnabled bool   `json:"aslr_enabled" bson:"aslr_enabled"`
	DepEnabled  bool   `json:"dep_enabled" bson:"dep_enabled"`
	Priority    string `json:"priority" bson:"priority"`
	OwnerSid    string `json:"owner_sid" bson:"owner_sid"`
	WindowTitle string `json:"window_title" bson:"window_title"`
	//StartupInfo dictionary `json:"startup_info" bson:"startup_info"`
	IntegrityLevel EnumTypeSTIX `json:"integrity_level" bson:"integrity_level"`
}

// WindowsServiceExtSTIX  объект "Windows Service Extension", по терминологии STIX.
// WindowsServiceExtSTIX расширяет объект описанием свойств, специфичных для служб Windows.
// Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-service-ext.
// Т.к. все свойства данного типа необязательные то должно быть определено хотябы одно из них.
// ServiceName - указывает название службы
// Descriptions - указывает список описаний для сулжбы
// DisplayName - указывает имя службы отображаемое в элементах управления графического интерфейса Windows (Windows GUI controls).
// GroupName - указывает имя группы порядка загрузки, членом которой является служба.
// StartType - указывает параметры запуска, определенные для службы.
//             Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-service-start-type-enum
// ServiceDllRefs - Указывает загруженные службой библиотеки DLL, как ссылки на один или несколько объектов,
//                  типа данных объектов должен быть FileObjectSTIX.
// ServiceType - Указывает тип службы. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-service-type-enum.
// ServiceStatus - Указывает текущее состояние службы.Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-service-status-enum.
type WindowsServiceExtSTIX struct {
	ServiceName    string       `json:"service_name" bson:"service_name"`
	Descriptions   []string     `json:"descriptions" bson:"descriptions"`
	DisplayName    string       `json:"display_name" bson:"display_name"`
	GroupName      string       `json:"group_name" bson:"group_name"`
	StartType      EnumTypeSTIX `json:"start_type" bson:"start_type"`
	ServiceDllRefs string       `json:"service_dll_refs" bson:"service_dll_refs"`
	ServiceType    EnumTypeSTIX `json:"service_type" bson:"service_type"`
	ServiceStatus  EnumTypeSTIX `json:"service_status" bson:"service_status"`
}

//ivi
// SoftwareObject объект "Software Object", по терминологии STIX.
// SoftwareObject представляет собой высокоуровневые свойства, связанные с программным обеспечением, включая программные продукты.
// Name - Указывает назвыание программного обеспечения.
// CPE - Указывает запись Common Platform Enumeration (CPE) для программного обеспечения, если она доступна.
//       Значение этого свойства должно быть значением CPE v2.3 из официального словаря NVD CPE [NVD] .
// SwID - Указывает запись Тегов Software Identification ID (SWID) [SWID] для программного обеспечения, если таковая имеется.
//        swID помеченный tagId, является глобально уникальным идентификатором и ДОЛЖЕН использоваться как полномочие
//        для идентификации помеченного продукта.
// Languages -Указывает языки, поддерживаемые программным обеспечением.
//            Значение каждого члена списка ДОЛЖНО быть кодом языка ISO 639-2 [ISO639 -2].
// Vendor - Указывает название производителя программного обеспечения.
// Version - Указывает версию ПО

type SoftwareObject struct {
	CommonPropertiesCyberObservableObjectSTIX
	Name      string   `json:"name" bson:"name"`
	CPE       string   `json:"cpe" bson:"cpe"`
	SwID      string   `json:"swid" bson:"swid"`
	Languages []string `json:"languages" bson:"languages"`
	Vendor    string   `json:"vendor" bson:"vendor"`
	Version   string   `json:"version" bson:"version"`
}

//ivi
// UserAccountObjectSTIX объект "User Account Object", по терминологии STIX.
// Представляет экземпляр любого типа учетной записи пользователя, включая, учетные записи операционной системы,
// устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей.
// Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств, определенных ниже,
// ДОЛЖНО быть инициализировано при использовании этого объекта.
// Extensions - Определяет словарь расширяющий тип "User Account Object" одно из расширений "unix-account-ext", реализуется описанным ниже типом, UNIXTMAccountExtensionSTIX
//              кроме этого производитель может созавать свои собственные типы расширений.
//  		    Ключи данного словаря идентифицируют тип расширения по имени, значения являются содержимым экземпляра расширения.
// UserId - Указывает идентификатор учетной записи. Формат идентификатора зависит от системы в которой находится данная учетная запись пользователя,
//          и может быть числовым идентификатором, идентификатором GUID, именем учетной записи, адресом электронной почты и т. Д.
//          Свойство  UserId должно быть заполнено любым значанием, являющимся уникальным идентификатором системы, членом которой является учетная запись.
//          Например, в системах UNIX он будет заполнено значением UID.
// Credential - Указывает учетные данные пользователя в открытом виде. Предназначено только для закрытого применения при изучении метаданных вредоносных программ
//				при их исследовании (например, жестко закодированный пароль администратора домена, который вредоносная
//              программа пытается использовать реализации тактики для бокового (латерального) перемещения)
//              и не должно применяться для совместного пользования PII.
// AccountLogin - Указывает логин пользователя используется в тех случаях,когда свойство UserId указывает другие данные,
//                чем то, что пользователь вводить при входе в систему.
// AccountType - Указывает тип аккаунта. Значения этого свойства берутся из множества закрепленного в открытом словаре, account-type-ov.
// DisplayName - Указывает отображаемое имя учетной записи, которое будет отображаться в пользовательских интерфейсах. \
//               В Unix, это равносильно полю gecos (gecos это поле учётной записи пользователя в файле /etc/passwd ).
// IsServiceAccount - Указывает, что учетная запись связана с сетевой службой или системным процессом (демоном), а не с конкретным человеком. (системный пользователь)
// IsPrivileged - Указывает, что учетная запись имеет повышенные привилегии (например, в случае root в Unix или учетной записи администратора Windows).
// CanEscalatePrivs  - Указывает, что учетная запись имеет возможность повышать привилегии (например, в случае sudo в Unix или учетной записи администратора домена Windows)
// IsDisabled  - Указывает, отключена ли учетная запись.
// AccountCreated - Указывает дату создания аккаунта
// AccountExpires - Указывает дату истечения срока действия учетной записи.
// CredentialLastChanged - Указывает lдату, когда учетные данные учетной записи были изменены в последний раз.
// AccountFirstLogin - Указывает время первого доступа к учетной записи
// AccountLastLogin - Указывает, когда к учетной записи был последний доступ.

type UserAccountObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	//Extensions            map[string]Dictiona `json:"" bson:""`
	UserId                string            `json:"user_id" bson:"user_id"`
	Credential            string            `json:"credential" bson:"credential"`
	AccountLogin          string            `json:"account_login" bson:"account_login"`
	AccountType           OpenVocabTypeSTIX `json:"account_type" bson:"account_type"`
	DisplayName           string            `json:"display_name" bson:"display_name"`
	IsServiceAccount      bool              `json:"is_service_account" bson:"is_service_account"`
	IsPrivileged          bool              `json:"is_privileged" bson:"is_privileged"`
	CanEscalatePrivs      bool              `json:"can_escalate_privs" bson:"can_escalate_privs"`
	IsDisabled            bool              `json:"is_disabled" bson:"is_disabled"`
	AccountCreated        time.Time         `json:"account_created" bson:"account_created"`
	AccountExpires        time.Time         `json:"account_expires" bson:"account_expires"`
	CredentialLastChanged time.Time         `json:"credential_last_changed" bson:"credential_last_changed"`
	AccountFirstLogin     time.Time         `json:"account_first_login" bson:"account_first_login"`
	AccountLastLogin      time.Time         `json:"account_last_login" bson:"account_last_login"`
}

//ivi
// UNIXTMAccountExtension объект "UNIXTM Account Extension", по терминологии STIX.
// Указывает расширение по умолчанию для сбора дополнительной информации об учетной записи UNIX-подобных ОС.
// Объект, использующий расширение учетной записи UNIX, должен содержать хотя бы одно свойство из этого расширения.
// Gid - указывает идентификатор основной группы пользователей в которую входит учетная запись
// Groups - указывает список названий групп, членом которых является учетная запись
// HomeDir - Указывает домашний каталог учетной записи.
// Shell - Указывает командную оболочку которая связана с учетной записью.

type UNIXTMAccountExtensionSTIX struct {
	Gid     int      `json:"gid" bson:"gid"`
	Groups  []string `json:"groups" bson:"groups"`
	HomeDir string   `json:"home_dir" bson:"home_dir"`
	Shell   string   `json:"shell" bson:"shell"`
}

//ivi
// WindowsTMRegistryKeyObject составной тип данных  "WindowsTM Registry Key Object", по терминологии STIX.
// Oписывает значение полей раздела реестра Windows. Поскольку все свойства этого объекта являются необязательными,
// по крайней мере одно из свойств,определенных ниже, должно быть инициализировано при использовании этого объекта.
// Key - Указывает полный путь к разделу реестра. Значение ключа,должно быть сохранено в регистре. В название елюча все сокращения должны быть раскрыты например, вместо HKLM следует использовать HKEY_LOCAL_MACHINE.
// Values - Указывает значения, найденные в разделе реестра.
// ModifiedTime - Указывает последнюю дату и время изменения раздела реестра.
// CreatorUserRef - Указывает ссылку на учетную запись пользователя, из под которой создан раздел реестра. Объект, на который ссылается это свойство, должен иметь тип user-account.
// NumberOfSubkeys - Указывает количество подразделов, содержащихся в разделе реестра.

type WindowsTMRegistryKeyObjectSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	Key             string                         `json:"key" bson:"key"`
	Values          []WindowsRegistryValueTypeSTIX `json:"values" bson:"values"`
	ModifiedTime    time.Time                      `json:"modified_time" bson:"modified_time"`
	CreatorUserRef  IdentifierTypeSTIX             `json:"creator_user_ref" bson:"creator_user_ref"`
	NumberOfSubkeys int                            `json:"number_of_subkeys" bson:"number_of_subkeys"`
}

//ivi
// WindowsTMRegistryValueType составной тип данных "windows-registry-value-type", по терминологии STIX. Данный тип фиксирует
// значения свойств находящихся в разделе реестра Windows.  Поскольку все свойства этого типа являются необязательными,
// по крайней мере одно из свойств, определенных ниже, должно быть инициализировано при использовании этого типа.
// Name - Указывает название параметра реестра. Для указания значения ключа реестра по умолчанию необходимо использовать пустую строку.
// Data - Указывает данные, содержащиеся в значении реестра.
// DataType - Указывает тип данных реестра (REG_*), используемый в значении реестра. Значения этого свойства должны быть получены из перечисления windows-registry-datatype enum.
type WindowsRegistryValueTypeSTIX struct {
	Name     string       `json:"name" bson:"name"`
	Data     string       `json:"data" bson:"data"`
	DataType EnumTypeSTIX `json:"data_type" bson:"data_type"`
}

// ivi
// X509CertificateObjectSTIX объект "X.509 Certificate Object", по терминологии STIX, представлет свойства сертификата X. 509,
// определенные в рекомендациях ITU X. 509 [X. 509].  X. 509  Certificate объект должен содержать по крайней мере одно
// cвойство специфичное для этого объекта (помимо type)
// IsSelfSigned - Указывает, является ли сертификат самозаверяющим, то есть подписан ли он тем же субъектом, личность которого он удостоверяет.
// Hashes - Указывает любые хэши, которые были вычислены для всего содержимого сертификата. Является типом данных словар, значения ключей которого должны быть из открытого словаря hash-algorithm-ov.
// Version- указывает версию закодированного сертификата
// SerialNumber - Указывает уникальный идентификатор сертификата, выданного конкретным Центр сертификации.
// SignatureAlgorithm - Указывает имя алгоритма, используемого для подписи сертификата.
// Issuer - содержит название удостоверяющего центра выдавшего сертификат
// ValidityNotBefore -дата начала действия сертификата.
// ValidityNotAfter - дата окончания действия сертификата.
// Subject - Указывает имя сущности, связанной с открытым ключом, хранящимся в поле "subject public key" открого ключа сертификата.
// SubjectPublicKeyAlgorithm - содержит название алгоритма применяемого для шифрования данных, отправляемых субъекту.
// SubjectPublicKeyModulus - Указывает модульную часть открытого ключа RSA.
// SubjectPublicKeyExponent - Указывает экспоненциальную часть открытого ключа RSA субъекта в виде целого числа.
// X509V3Extension - Указывает любые стандартные расширения X. 509 v3, которые могут использоваться в сертификате.
type X509CertificateSTIX struct {
	CommonPropertiesCyberObservableObjectSTIX
	IsSelfSigned              bool                     `json:"is_self_signed" bson:"is_self_signed"`
	Hashes                    HashesTypeSTIX           `json:"hashes" bson:"hashes"`
	Version                   string                   `json:"version" bson:"version"`
	SerialNumber              string                   `json:"serial_number" bson:"serial_number"`
	SignatureAlgorithm        string                   `json:"signature_algorithm" bson:"signature_algorithm"`
	Issuer                    string                   `json:"issuer" bson:"issuer"`
	ValidityNotBefore         time.Time                `json:"validity_not_before" bson:"validity_not_before"`
	ValidityNotAfter          time.Time                `json:"validity_not_after" bson:"validity_not_after"`
	Subject                   string                   `json:"subject" bson:"subject"`
	SubjectPublicKeyAlgorithm string                   `json:"subject_public_key_algorithm" bson:"subject_public_key_algorithm"`
	SubjectPublicKeyModulus   string                   `json:"subject_public_key_modulus" bson:"subject_public_key_modulus"`
	SubjectPublicKeyExponent  int                      `json:"subject_public_key_exponent" bson:"subject_public_key_exponent"`
	X509V3Extensions          X509V3ExtensionsTypeSTIX `json:"x509_v3_extensions" bson:"x509_v3_extensions"`
}

// Тип ExtentionX509v3 составной тип данных "x509-v3-extensions-type", по терминологии STIX. Описывает поля расширения X. 509 v3,
// фиксрует дополнительной информации, такую как альтернативные имена субъектов.Объект, использующий тип "x509-v3-extensions-type",
// должен определить хотя бы одно из этих полей в нем. Данный тип расширяет только объекты "X.509 Certificate Object".
// BasicConstraints - Задает многозначное расширение, которое указывает, является ли сертификат сертификатом Удостоверяющего центра (CA). Первое (обязательное) название CA сопровождается истинным или ложным. Если CA имеет значение TRUE, то может быть включено необязательное имя pathlen, за которым следует неотрицательное значение. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.19.
// NameConstraints - Указывает пространство имен, в котором должны располагаться все имена  применяемые в сертификатах указанных в пути сертификации. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.30.
// PolicyConstraints - Указывает любые ограничения на проверку сертификатов, выданных Удостоверяющим центром.  Также эквивалентно значению идентификатора объекта (OID) 2.5.29.36.
// KeyUsage - Задает многозначное расширение, состоящее из списка имен разрешенных для использования ключей. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.15.
// ExtendedkeyUsage - Указывает список целей, для которых может использоваться открытый ключ сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.37.
// Subjectkeyidentifier - Указывает идентификатор, который обеспечивает средство идентификации сертификатов, содержащих определенный открытый ключ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.14.
// AuthorityKeyIdentifier - Указывает идентификатор, который обеспечивает средство идентификации открытого ключа, соответствующего закрытому ключу, используемому для подписи сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.35.
// SubjectAlternativeName - Указывает дополнительные идентификаторы, которые должны быть привязаны к субъекту сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.17.
// IssuerAlternativeName - Указывает дополнительные идентификаторы, которые должны быть привязаны к эмитенту сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.18.
// SubjectDirectoryAttributes - Указывает идентификационные признаки (например, национальность) субъекта. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.9.
// CrlDistributionPoints - Указывает способ получения информации CRL. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.31.
// InhibitAnyPolicy - Указывает количество дополнительных сертификатов, которые могут появиться в пути до того, как любая Политика больше не будет разрешена. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.54.
// PrivateKeyUsagePeriodNotBefore - Указывает дату начала срока действия закрытого ключа, если он отличается от срока действия сертификата.
// PrivateKeyUsagePeriodNotAfter -  Указывает дату окончания срока действия закрытого ключа, если он отличается от срока действия сертификата.
// CertificatePolicies - Задает последовательность из одного или нескольких терминов информации о политике, каждый из которых состоит из идентификатора объекта (OID) и необязательных квалификаторов. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.32.
// PolicyMappings - Указывает одну или несколько пар OID; каждая пара включает issuerDomainPolicy и subjectDomainPolicy. Пары индикаторов указывают на то, считает ли выдающий УЦ свою issuerDomainPolicy эквивалентной subjectDomainPolicy субъекта УЦ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.33.

type X509V3ExtensionsTypeSTIX struct {
	BasicConstraints               string    `json:"basic_constraints" bson:"basic_constraints"`
	NameConstraints                string    `json:"name_constraints" bson:"name_constraints"`
	PolicyConstraints              string    `json:"policy_constraints" bson:"policy_constraints"`
	KeyUsage                       string    `json:"key_usage" bson:"key_usage"`
	ExtendedkeyUsage               string    `json:"extended_key_usage" bson:"extended_key_usage"`
	Subjectkeyidentifier           string    `json:"subject_key_identifier" bson:"subject_key_identifier"`
	AuthorityKeyIdentifier         string    `json:"authority_key_identifier" bson:"authority_key_identifier"`
	SubjectAlternativeName         string    `json:"subject_alternative_name" bson:"subject_alternative_name"`
	IssuerAlternativeName          string    `json:"issuer_alternative_name" bson:"issuer_alternative_name"`
	SubjectDirectoryAttributes     string    `json:"subject_directory_attributes" bson:"subject_directory_attributes"`
	CrlDistributionPoints          string    `json:"crl_distribution_points" bson:"crl_distribution_points"`
	InhibitAnyPolicy               string    `json:"inhibit_any_policy" bson:"inhibit_any_policy"`
	PrivateKeyUsagePeriodNotBefore time.Time `json:"private_key_usage_period_not_before" bson:"private_key_usage_period_not_before"`
	PrivateKeyUsagePeriodNotAfter  time.Time `json:"private_key_usage_period_not_after" bson:"private_key_usage_period_not_after"`
	CertificatePolicies            string    `json:"certificate_policies" bson:"certificate_policies"`
	PolicyMappings                 string    `json:"policy_mappings" bson:"policy_mappings"`
}
