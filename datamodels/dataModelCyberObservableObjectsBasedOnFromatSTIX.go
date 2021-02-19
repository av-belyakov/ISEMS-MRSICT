package datamodels

import "time"

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

//ArtifactObjectSTIX объект "Artifact", по терминологии STIX, позволяет захватывать массив байтов (8 бит) в виде строки в кодировке base64
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
