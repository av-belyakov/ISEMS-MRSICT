package datamodels

import (
	"time"
)

/*********************************************************************************/
/********** 			Cyber-observable Objects STIX (МЕТОДЫ)			**********/
/*********************************************************************************/

/*func (ocpcstix *OptionalCommonPropertiesCyberObservableObjectSTIX) validateStructCommonFields() bool {
	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(ocpcstix.SpecVersion)) {
		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(ocpcstix.ObjectMarkingRefs) > 0 {
		for _, value := range ocpcstix.ObjectMarkingRefs {
			if !value.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(ocpcstix.GranularMarkings) > 0 {
		for _, value := range ocpcstix.GranularMarkings {
			//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
			if !value.CheckGranularMarkingsTypeSTIX() {
				return false
			}
		}
	}

	return true
}*/

/* --- ArtifactCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (astix ArtifactCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return astix.ArtifactCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ArtifactCyberObservableObjectSTIX
func (astix ArtifactCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             astix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- AutonomousSystemCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (asstix AutonomousSystemCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return asstix.AutonomousSystemCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа AutonomousSystemCyberObservableObjectSTIX
func (asstix AutonomousSystemCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             asstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- DirectoryCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (dstix DirectoryCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return dstix.DirectoryCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа DirectoryCyberObservableObjectSTIX
func (dstix DirectoryCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             dstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- DomainNameCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (dnstix DomainNameCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return dnstix.DomainNameCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа DomainNameCyberObservableObjectSTIX
func (dnstix DomainNameCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             dnstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- EmailAddressCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (eastix EmailAddressCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return eastix.EmailAddressCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа EmailAddressCyberObservableObjectSTIX
func (eastix EmailAddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             eastix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- EmailMessageCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (emstix EmailMessageCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return emstix.EmailMessageCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа EmailMessageCyberObservableObjectSTIX
func (emstix EmailMessageCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             emstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- FileCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (fstix FileCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return fstix.FileCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа FileCyberObservableObjectSTIX
func (fstix FileCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             fstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- IPv4AddressCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (ipv4stix IPv4AddressCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return ipv4stix.IPv4AddressCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа IPv4AddressCyberObservableObjectSTIX
func (ipv4stix IPv4AddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ipv4stix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- IPv6AddressCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (ipv6stix IPv6AddressCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return ipv6stix.IPv6AddressCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа IPv6AddressCyberObservableObjectSTIX
func (ipv6stix IPv6AddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ipv6stix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- MACAddressCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (macstix MACAddressCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return macstix.MACAddressCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа MACAddressCyberObservableObjectSTIX
func (macstix MACAddressCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             macstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- MutexCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (mstix MutexCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return mstix.MutexCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа MutexCyberObservableObjectSTIX
func (mstix MutexCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             mstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- NetworkTrafficCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (ntstix NetworkTrafficCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return ntstix.NetworkTrafficCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа NetworkTrafficCyberObservableObjectSTIX
func (ntstix NetworkTrafficCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             ntstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- ProcessCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (pstix ProcessCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return pstix.ProcessCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа ProcessCyberObservableObjectSTIX
func (pstix ProcessCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             pstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- SoftwareCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (sstix SoftwareCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return sstix.SoftwareCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа SoftwareCyberObservableObjectSTIX
func (sstix SoftwareCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             sstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- URLCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (urlstix URLCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return urlstix.URLCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа URLCyberObservableObjectSTIX
func (urlstix URLCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             urlstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- UserAccountCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (uastix UserAccountCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return uastix.UserAccountCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа UserAccountCyberObservableObjectSTIX
func (uastix UserAccountCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             uastix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- WindowsRegistryKeyCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (wrstix WindowsRegistryKeyCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return wrstix.WindowsRegistryKeyCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа WindowsRegistryKeyCyberObservableObjectSTIX
func (wrstix WindowsRegistryKeyCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             wrstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}

/* --- X509CertificateCyberObservableObjectSTIX --- */

// GetCurrentObject возвращает текущий объект
func (x509sstix X509CertificateCyberObservableObjectSTIX) GetCurrentObject() interface{} {
	return x509sstix.X509CertificateCyberObservableObjectSTIX
}

// ComparisonTypeCommonFields выполняет сравнение двух объектов типа X509CertificateCyberObservableObjectSTIX
func (x509sstix X509CertificateCyberObservableObjectSTIX) ComparisonTypeCommonFields(newObj interface{}, src string) (bool, DifferentObjectType, error) {
	var (
		isEqual bool = true
		cot          = DifferentObjectType{
			SourceReceivingChanges: src,
			ModifiedTime:           time.Now(),
			CollectionName:         "stix_object_collection",
			DocumentID:             x509sstix.ID,
		}
	)

	/*
				!!! ПОКА ЭТО ТОЛЬКО ЗАГЛУШКА необходимая для совместимости !!!
		vNew, ok := newObj.(*)
		if !ok {
			return isEqual, cot, fmt.Errorf("type conversion error")
		}
	*/

	return isEqual, cot, nil
}
