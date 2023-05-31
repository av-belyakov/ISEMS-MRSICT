package datamodels

// decodingExtensionsSTIX декодирует следующие типы STIX расширений:
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

	//fmt.Printf("func 'decodingExtensionsSTIX', extType = %s, \n", extType)

	switch extType {
	case "archive-ext":
		var archiveExt mstixo.ArchiveFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &archiveExt)

		return archiveExt, err
	case "ntfs-ext":
		var ntfsExt mstixo.NTFSFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &ntfsExt)

		return ntfsExt, err
	case "pdf-ext":
		var pdfExt mstixo.PDFFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &pdfExt)

		return pdfExt, err
	case "raster-image-ext":
		var rasterImageExt mstixo.RasterImageFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &rasterImageExt)

		//fmt.Printf("func 'decodingExtensionsSTIX', rasterImageExt = %v, \n", rasterImageExt)

		return rasterImageExt, err
	case "windows-pebinary-ext":
		var windowsPebinaryExt mstixo.WindowsPEBinaryFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &windowsPebinaryExt)

		return windowsPebinaryExt, err
	case "http-request-ext":
		var httpRequestExt mstixo.HTTPRequestExtensionSTIX
		err = json.Unmarshal(*rawMsg, &httpRequestExt)

		return httpRequestExt, err
	case "icmp-ext":
		var icmpExt mstixo.ICMPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &icmpExt)

		return icmpExt, err
	case "socket-ext":
		var socketExt mstixo.NetworkSocketExtensionSTIX
		err := json.Unmarshal(*rawMsg, &socketExt)

		return socketExt, err
	case "tcp-ext":
		var tcpExt mstixo.TCPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &tcpExt)

		return tcpExt, err
	case "windows-process-ext":
		var windowsProcessExt mstixo.WindowsProcessExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsProcessExt)

		return windowsProcessExt, err
	case "windows-service-ext":
		var windowsServiceExt mstixo.WindowsServiceExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsServiceExt)

		return windowsServiceExt, err
	case "unix-account-ext":
		var unixAccountExt mstixo.UNIXAccountExtensionSTIX
		err := json.Unmarshal(*rawMsg, &unixAccountExt)

		return unixAccountExt, err
	default:
		return struct{}{}, nil
	}
}

// checkingExtensionsSTIX выполняет проверку полей следующих типов STIX расширений:
// - "archive-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "windows-service-ext"
func checkingExtensionsSTIX(extType interface{}) bool {
	switch et := extType.(type) {
	case mstixo.ArchiveFileExtensionSTIX:
		for _, v := range et.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}

	case mstixo.NTFSFileExtensionSTIX:
		if len(et.AlternateDataStreams) == 0 {
			return true
		}

		for _, v := range et.AlternateDataStreams {
			if !v.Hashes.CheckHashesTypeSTIX() {
				return false
			}
		}

	case mstixo.WindowsPEBinaryFileExtensionSTIX:
		if !et.FileHeaderHashes.CheckHashesTypeSTIX() {
			return false
		}

	case mstixo.HTTPRequestExtensionSTIX:
		if !et.MessageBodyDataRef.CheckIdentifierTypeSTIX() {
			return false
		}

	case mstixo.WindowsServiceExtensionSTIX:
		for _, v := range et.ServiceDllRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// sanitizeExtensionsSTIX для ряда полей следующих типов STIX расширений:
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
// выполняет замену некоторых специальных символов на их HTML код
func sanitizeExtensionsSTIX(extType interface{}) interface{} {
	sanitizeList := func(rh map[string]string) map[string]string {
		mapTmp := make(map[string]string, len(rh))
		for k, v := range rh {
			mapTmp[k] = commonlibs.StringSanitize(v)
		}

		return mapTmp
	}

	sanitizeExifTags := func(et mstixo.ExifTags) mstixo.ExifTags {
		et.Make = commonlibs.StringSanitize(et.Make)
		et.Model = commonlibs.StringSanitize(et.Model)

		return et
	}

	switch et := extType.(type) {
	case mstixo.ArchiveFileExtensionSTIX:
		return commonlibs.StringSanitize(et.Comment)

	case mstixo.NTFSFileExtensionSTIX:
		var tmpType mstixo.NTFSFileExtensionSTIX
		tmpType.SID = commonlibs.StringSanitize(et.SID)

		size := len(et.AlternateDataStreams)

		if size == 0 {
			return tmpType
		}

		ads := make([]mstixo.AlternateDataStreamTypeSTIX, 0, size)
		for _, v := range et.AlternateDataStreams {
			ads = append(ads, mstixo.AlternateDataStreamTypeSTIX{
				Name:   commonlibs.StringSanitize(v.Name),
				Hashes: v.Hashes,
				Size:   v.Size,
			})
		}
		tmpType.AlternateDataStreams = ads

		return tmpType

	case mstixo.PDFFileExtensionSTIX:
		return mstixo.PDFFileExtensionSTIX{
			Version:          commonlibs.StringSanitize(et.Version),
			IsOptimized:      et.IsOptimized,
			DocumentInfoDict: sanitizeList(et.DocumentInfoDict),
			Pdfid0:           commonlibs.StringSanitize(et.Pdfid0),
			Pdfid1:           commonlibs.StringSanitize(et.Pdfid1),
		}

	case mstixo.RasterImageFileExtensionSTIX:
		return mstixo.RasterImageFileExtensionSTIX{
			ImageHeight:  et.ImageHeight,
			ImageWidth:   et.ImageWidth,
			BitsPerPixel: et.BitsPerPixel,
			ExifTags:     sanitizeExifTags(et.ExifTags),
		}

	case mstixo.WindowsPEBinaryFileExtensionSTIX:
		ssize := len(et.Sections)
		ns := make([]mstixo.WindowsPESectionTypeSTIX, 0, ssize)
		if ssize > 0 {
			for _, v := range et.Sections {
				tmp := v.SanitizeStructWindowsPESectionTypeSTIX()
				ns = append(ns, tmp)
			}
			et.Sections = ns
		}

		return mstixo.WindowsPEBinaryFileExtensionSTIX{
			PeType:                  et.PeType.SanitizeStructOpenVocabTypeSTIX(),
			Imphash:                 commonlibs.StringSanitize(et.Imphash),
			MachineHex:              commonlibs.StringSanitize(et.MachineHex),
			NumberOfSections:        et.NumberOfSections,
			TimeDateStamp:           et.TimeDateStamp,
			PointerToSymbolTableHex: commonlibs.StringSanitize(et.PointerToSymbolTableHex),
			NumberOfSymbols:         et.NumberOfSymbols,
			SizeOfOptionalHeader:    et.SizeOfOptionalHeader,
			CharacteristicsHex:      commonlibs.StringSanitize(et.CharacteristicsHex),
			FileHeaderHashes: func(l mstixo.HashesTypeSTIX) mstixo.HashesTypeSTIX {
				fhh := make(mstixo.HashesTypeSTIX, len(l))

				for k, v := range l {
					fhh[k] = commonlibs.StringSanitize(string(v))
				}

				return fhh
			}(et.FileHeaderHashes),
			OptionalHeader: et.OptionalHeader.SanitizeStructWindowsPEOptionalHeaderTypeSTIX(),
			Sections:       ns,
		}

	case mstixo.HTTPRequestExtensionSTIX:
		return mstixo.HTTPRequestExtensionSTIX{
			RequestMethod:      commonlibs.StringSanitize(et.RequestMethod),
			RequestValue:       commonlibs.StringSanitize(et.RequestValue),
			RequestVersion:     commonlibs.StringSanitize(et.RequestVersion),
			RequestHeader:      sanitizeList(et.RequestHeader),
			MessageBodyLength:  et.MessageBodyLength,
			MessageBodyDataRef: et.MessageBodyDataRef,
		}

	case mstixo.ICMPExtensionSTIX:
		return mstixo.ICMPExtensionSTIX{
			ICMPTypeHex: commonlibs.StringSanitize(et.ICMPTypeHex),
			ICMPCodeHex: commonlibs.StringSanitize(et.ICMPCodeHex),
		}

	case mstixo.NetworkSocketExtensionSTIX:
		return mstixo.NetworkSocketExtensionSTIX{
			AddressFamily:    mstixo.EnumTypeSTIX((commonlibs.StringSanitize(string(et.AddressFamily)))),
			IsBlocking:       et.IsBlocking,
			IsListening:      et.IsListening,
			Options:          et.Options,
			SocketType:       mstixo.EnumTypeSTIX((commonlibs.StringSanitize(string(et.SocketType)))),
			SocketDescriptor: et.SocketDescriptor,
			SocketHandle:     et.SocketHandle,
		}

	case mstixo.TCPExtensionSTIX:
		return mstixo.TCPExtensionSTIX{
			SrcFlagsHex: commonlibs.StringSanitize(et.SrcFlagsHex),
			DstFlagsHex: commonlibs.StringSanitize(et.DstFlagsHex),
		}

	case mstixo.WindowsProcessExtensionSTIX:
		return mstixo.WindowsProcessExtensionSTIX{
			ASLREnabled:    et.ASLREnabled,
			DEPEnabled:     et.DEPEnabled,
			Priority:       commonlibs.StringSanitize(et.Priority),
			OwnerSID:       commonlibs.StringSanitize(et.OwnerSID),
			WindowTitle:    commonlibs.StringSanitize(et.WindowTitle),
			StartupInfo:    sanitizeList(et.StartupInfo),
			IntegrityLevel: mstixo.EnumTypeSTIX(commonlibs.StringSanitize(string(et.IntegrityLevel))),
		}

	case mstixo.WindowsServiceExtensionSTIX:
		return mstixo.WindowsServiceExtensionSTIX{
			ServiceName: commonlibs.StringSanitize(et.ServiceName),
			Descriptions: func(l []string) []string {
				size := len(l)
				nl := make([]string, 0, size)
				for _, v := range l {
					nl = append(nl, commonlibs.StringSanitize(v))
				}

				return nl
			}(et.Descriptions),
			DisplayName:    commonlibs.StringSanitize(et.DisplayName),
			GroupName:      commonlibs.StringSanitize(et.GroupName),
			StartType:      mstixo.EnumTypeSTIX(commonlibs.StringSanitize(string(et.StartType))),
			ServiceDllRefs: et.ServiceDllRefs,
			ServiceType:    mstixo.EnumTypeSTIX(commonlibs.StringSanitize(string(et.ServiceType))),
			ServiceStatus:  mstixo.EnumTypeSTIX(commonlibs.StringSanitize(string(et.ServiceStatus))),
		}

	case mstixo.UNIXAccountExtensionSTIX:
		return mstixo.UNIXAccountExtensionSTIX{
			GID: et.GID,
			Groups: func(l []string) []string {
				size := len(l)
				nl := make([]string, 0, size)
				for _, v := range l {
					nl = append(nl, commonlibs.StringSanitize(v))
				}

				return nl
			}(et.Groups),
			HomeDir: commonlibs.StringSanitize(et.HomeDir),
			Shell:   commonlibs.StringSanitize(et.Shell),
		}
	}

	return extType
}

// toStringBeautiful выполняет красивое представление информации содержащейся в следующих типах
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
func toStringBeautiful(extType interface{}) string {
	str := "\t\t"

	switch et := extType.(type) {
	case mstixo.ArchiveFileExtensionSTIX:
		str += fmt.Sprintln("contains_refs:")
		for k, v := range et.ContainsRefs {
			str += fmt.Sprintf("\t\t\tcontains_ref '%d': '%v'", k, v)
		}
		str += fmt.Sprintf("\t\tcomment: '%v'\n", et.Comment)

	case mstixo.NTFSFileExtensionSTIX:
		str += fmt.Sprintf("sid: '%s'\n", et.SID)
		str += fmt.Sprintln("\t\talternate_data_streams:")
		for k, v := range et.AlternateDataStreams {
			str += fmt.Sprintf("\t\t\talternate_data_stream '%d':\n", k)
			str += fmt.Sprintf("\t\t\t\tname: '%s'\n", v.Name)
			str += fmt.Sprintln("\t\t\t\thashes:")
			for a, b := range v.Hashes {
				str += fmt.Sprintf("\t\t\t\t\t'%s': '%v'\n", a, b)
			}
			str += fmt.Sprintf("\t\t\t\tsize: '%d'\n", v.Size)
		}

	case mstixo.PDFFileExtensionSTIX:
		str += fmt.Sprintf("version: '%s'\n", et.Version)
		str += fmt.Sprintf("\t\tis_optimized: '%v'\n", et.IsOptimized)
		str += fmt.Sprintln("\t\tdocument_info_dict:")
		for k, v := range et.DocumentInfoDict {
			str += fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v)
		}
		str += fmt.Sprintf("\t\tpdfid0: '%s'\n", et.Pdfid0)
		str += fmt.Sprintf("\t\tpdfid1: '%s'\n", et.Pdfid1)

	case mstixo.RasterImageFileExtensionSTIX:
		str += fmt.Sprintf("image_height: '%d'\n", et.ImageHeight)
		str += fmt.Sprintf("\t\timage_width: '%d'\n", et.ImageWidth)
		str += fmt.Sprintf("\t\tbits_per_pixel: '%d'\n", et.BitsPerPixel)
		str += fmt.Sprintln("\t\texif_tags:")
		str += fmt.Sprintf("\t\t\t'make': '%s'\n", et.ExifTags.Make)
		str += fmt.Sprintf("\t\t\t'model': '%s'\n", et.ExifTags.Model)
		str += fmt.Sprintf("\t\t\t'xResolution': '%d'\n", et.ExifTags.XResolution)
		str += fmt.Sprintf("\t\t\t'yResolution': '%d'\n", et.ExifTags.YResolution)

	case mstixo.WindowsPEBinaryFileExtensionSTIX:
		str += fmt.Sprintf("pe_type: '%v'\n", et.PeType)
		str += fmt.Sprintf("\t\timphash: '%s'\n", et.Imphash)
		str += fmt.Sprintf("\t\tmachine_hex: '%s'\n", et.MachineHex)
		str += fmt.Sprintf("\t\tnumber_of_sections: '%d'\n", et.NumberOfSections)
		str += fmt.Sprintf("\t\ttime_date_stamp: '%v'\n", et.TimeDateStamp)
		str += fmt.Sprintf("\t\tpointer_to_symbol_table_hex: '%s'\n", et.PointerToSymbolTableHex)
		str += fmt.Sprintf("\t\tnumber_of_symbols: '%d'\n", et.NumberOfSymbols)
		str += fmt.Sprintf("\t\tsize_of_optional_header: '%d'\n", et.SizeOfOptionalHeader)
		str += fmt.Sprintf("\t\tcharacteristics_hex: '%s'\n", et.CharacteristicsHex)
		str += fmt.Sprintln("\t\tfile_header_hashes:")
		for k, v := range et.FileHeaderHashes {
			str += fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v)
		}
		str += fmt.Sprintln("\t\toptional_header:")
		str += fmt.Sprintf("\t\t\tmagic_hex: '%s'\n", et.OptionalHeader.MagicHex)
		str += fmt.Sprintf("\t\t\tmajor_linker_version: '%d'\n", et.OptionalHeader.MajorLinkerVersion)
		str += fmt.Sprintf("\t\t\tminor_linker_version: '%d'\n", et.OptionalHeader.MinorLinkerVersion)
		str += fmt.Sprintf("\t\t\tsize_of_code: '%d'\n", et.OptionalHeader.SizeOfCode)
		str += fmt.Sprintf("\t\t\tsize_of_initialized_data: '%d'\n", et.OptionalHeader.SizeOfInitializedData)
		str += fmt.Sprintf("\t\t\tsize_of_uninitialized_data: '%d'\n", et.OptionalHeader.SizeOfUninitializedData)
		str += fmt.Sprintf("\t\t\taddress_of_entry_point: '%d'\n", et.OptionalHeader.AddressOfEntryPoint)
		str += fmt.Sprintf("\t\t\tbase_of_code: '%d'\n", et.OptionalHeader.BaseOfCode)
		str += fmt.Sprintf("\t\t\timage_base: '%d'\n", et.OptionalHeader.ImageBase)
		str += fmt.Sprintf("\t\t\tsection_alignment: '%d'\n", et.OptionalHeader.SectionAlignment)
		str += fmt.Sprintf("\t\t\tfile_alignment: '%d'\n", et.OptionalHeader.FileAlignment)
		str += fmt.Sprintf("\t\t\tmajor_os_version: '%d'\n", et.OptionalHeader.MajorOSVersion)
		str += fmt.Sprintf("\t\t\tminor_os_version: '%d'\n", et.OptionalHeader.MinorOSVersion)
		str += fmt.Sprintf("\t\t\tmajor_image_version: '%d'\n", et.OptionalHeader.MajorImageVersion)
		str += fmt.Sprintf("\t\t\tminor_image_version: '%d'\n", et.OptionalHeader.MinorImageVersion)
		str += fmt.Sprintf("\t\t\tmajor_subsystem_version: '%d'\n", et.OptionalHeader.MajorSubsystemVersion)
		str += fmt.Sprintf("\t\t\tminor_subsystem_version: '%d'\n", et.OptionalHeader.MinorSubsystemVersion)
		str += fmt.Sprintf("\t\t\twin32_version_value_hex: '%s'\n", et.OptionalHeader.Win32VersionValueHex)
		str += fmt.Sprintf("\t\t\tsize_of_image: '%d'\n", et.OptionalHeader.SizeOfImage)
		str += fmt.Sprintf("\t\t\tsize_of_headers: '%d'\n", et.OptionalHeader.SizeOfHeaders)
		str += fmt.Sprintf("\t\t\tchecksum_hex: '%s'\n", et.OptionalHeader.ChecksumHex)
		str += fmt.Sprintf("\t\t\tsubsystem_hex: '%s'\n", et.OptionalHeader.SubsystemHex)
		str += fmt.Sprintf("\t\t\tdll_characteristics_hex: '%s'\n", et.OptionalHeader.DllCharacteristicsHex)
		str += fmt.Sprintf("\t\t\tsize_of_stack_reserve: '%d'\n", et.OptionalHeader.SizeOfStackReserve)
		str += fmt.Sprintf("\t\t\tsize_of_stack_commit: '%d'\n", et.OptionalHeader.SizeOfStackCommit)
		str += fmt.Sprintf("\t\t\tsize_of_heap_reserve: '%d'\n", et.OptionalHeader.SizeOfHeapReserve)
		str += fmt.Sprintf("\t\t\tsize_of_heap_commit: '%d'\n", et.OptionalHeader.SizeOfHeapCommit)
		str += fmt.Sprintf("\t\t\tloader_flags_hex: '%s'\n", et.OptionalHeader.LoaderFlagsHex)
		str += fmt.Sprintf("\t\t\tnumber_of_rva_and_aizes: '%d'\n", et.OptionalHeader.NumberOfRvaAndSizes)
		str += fmt.Sprintln("\t\t\thashes:")
		for k, v := range et.OptionalHeader.Hashes {
			str += fmt.Sprintf("\t\t\t\thashe '%d': '%v'\n", k, v)
		}
		str += fmt.Sprintln("\t\tsections:")
		for k, v := range et.Sections {
			str += fmt.Sprintf("\t\t\tsection '%d':\n", k)
			str += fmt.Sprintf("\t\t\t\tname: '%s'\n", v.Name)
			str += fmt.Sprintf("\t\t\t\tsize: '%v'\n", v.Size)
			str += fmt.Sprintf("\t\t\t\tentropy: '%v'\n", v.Entropy)
			str += fmt.Sprintln("\t\t\t\thashes:", v.Hashes)
			for k, v := range v.Hashes {
				str += fmt.Sprintf("\t\t\t\t\t'%s': '%s'\n", k, v)
			}
		}

	case mstixo.HTTPRequestExtensionSTIX:
		str += fmt.Sprintf("request_method: '%s'\n", et.RequestMethod)
		str += fmt.Sprintf("\t\trequest_value: '%s'\n", et.RequestValue)
		str += fmt.Sprintf("\t\trequest_version: '%s'\n", et.RequestVersion)
		str += fmt.Sprintln("\t\trequest_header:")
		for k, v := range et.RequestHeader {
			str += fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v)
		}
		str += fmt.Sprintf("\t\tmessage_body_length: '%d'\n", et.MessageBodyLength)
		str += fmt.Sprintf("\t\tmessage_body_data_ref: '%v'\n", et.MessageBodyDataRef)

	case mstixo.ICMPExtensionSTIX:
		str += fmt.Sprintf("icmp_type_hex: '%s'\n", et.ICMPTypeHex)
		str += fmt.Sprintf("\t\ticmp_code_hex: '%s'\n", et.ICMPCodeHex)

	case mstixo.NetworkSocketExtensionSTIX:
		str += fmt.Sprintf("address_family: '%v'\n", et.AddressFamily)
		str += fmt.Sprintf("\t\tis_blocking: '%v'\n", et.IsBlocking)
		str += fmt.Sprintf("\t\tis_listening: '%v'\n", et.IsListening)
		str += fmt.Sprintln("\t\toptions:")
		for k, v := range et.Options {
			str += fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v)
		}
		str += fmt.Sprintf("\t\tsocket_type: '%v'\n", et.SocketType)
		str += fmt.Sprintf("\t\tsocket_descriptor: '%d'\n", et.SocketDescriptor)
		str += fmt.Sprintf("\t\tsocket_handle: '%d'\n", et.SocketHandle)

	case mstixo.TCPExtensionSTIX:
		str += fmt.Sprintf("src_flags_hex: '%v'\n", et.SrcFlagsHex)
		str += fmt.Sprintf("\t\tdst_flags_hex: '%v'\n", et.DstFlagsHex)

	case mstixo.WindowsProcessExtensionSTIX:
		str += fmt.Sprintf("aslr_enabled: '%v'\n", et.ASLREnabled)
		str += fmt.Sprintf("\t\tdep_enabled: '%v'\n", et.DEPEnabled)
		str += fmt.Sprintf("\t\tpriority: '%s'\n", et.Priority)
		str += fmt.Sprintf("\t\towner_sid: '%s'\n", et.OwnerSID)
		str += fmt.Sprintf("\t\twindow_title: '%s'\n", et.WindowTitle)
		str += fmt.Sprintln("\t\tstartup_info:")
		for k, v := range et.StartupInfo {
			str += fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v)
		}
		str += fmt.Sprintf("\t\tintegrity_level: '%v'\n", et.IntegrityLevel)

	case mstixo.WindowsServiceExtensionSTIX:
		str += fmt.Sprintf("service_name: '%s'\n", et.ServiceName)
		str += fmt.Sprintln("\t\tdescriptions:")
		for k, v := range et.Descriptions {
			str += fmt.Sprintf("\t\t\tdescription '%d': '%s'\n", k, v)
		}
		str += fmt.Sprintf("\t\tdisplay_name: '%s'\n", et.DisplayName)
		str += fmt.Sprintf("\t\tgroup_name: '%s'\n", et.GroupName)
		str += fmt.Sprintf("\t\tstart_type: '%v'\n", et.StartType)
		str += fmt.Sprintln("\t\tservice_dll_refs:")
		for k, v := range et.ServiceDllRefs {
			str += fmt.Sprintf("\t\t\tservice_dll_ref '%d': '%v'", k, v)
		}
		str += fmt.Sprintf("\t\tservice_type: '%v'\n", et.ServiceType)
		str += fmt.Sprintf("\t\tservice_status: '%v'\n", et.ServiceStatus)

	case mstixo.UNIXAccountExtensionSTIX:
		str += fmt.Sprintf("gid: '%d'\n", et.GID)
		str += fmt.Sprintln("\t\tgroups:")
		for k, v := range et.Groups {
			str += fmt.Sprintf("\t\t\tgroup '%d': '%s'\n", k, v)
		}
		str += fmt.Sprintf("\t\thome_dir: '%s'\n", et.HomeDir)
		str += fmt.Sprintf("\t\tshell: '%s'\n", et.Shell)

	default:
		str += fmt.Sprintf("%v\n", et)

	}

	return str
}*/
