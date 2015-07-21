package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func i18n_resources_de_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x60, 0xad, 0x89, 0x20, 0xad, 0x65, 0xf9, 0x79, 0x30, 0x9d, 0x3e,
		0x20, 0x5d, 0xa5, 0xd8, 0x74, 0xe6, 0xe6, 0xa7, 0x64, 0xa6, 0x65, 0xa6,
		0x82, 0x6c, 0x4c, 0x4b, 0xcc, 0x29, 0x4e, 0x05, 0x89, 0xd7, 0xea, 0x60,
		0x71, 0x4e, 0x58, 0x6a, 0x51, 0x31, 0xd0, 0x78, 0x85, 0xea, 0x6a, 0x3d,
		0xc7, 0xa2, 0x74, 0x83, 0xda, 0xda, 0x98, 0x3c, 0x9c, 0xf6, 0x63, 0xa8,
		0x25, 0x60, 0x1f, 0x57, 0x2c, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xed,
		0xb5, 0x19, 0x01, 0x01, 0x00, 0x00,
	},
		"i18n/resources/de.all.json",
	)
}

func i18n_resources_de_de_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x60, 0xad, 0x89, 0x20, 0xad, 0x65, 0xf9, 0x79, 0x30, 0x9d, 0x3e,
		0x20, 0x5d, 0xa5, 0xd8, 0x74, 0xe6, 0xe6, 0xa7, 0x64, 0xa6, 0x65, 0xa6,
		0x82, 0x6c, 0x4c, 0x4b, 0xcc, 0x29, 0x4e, 0x05, 0x89, 0xd7, 0xea, 0x60,
		0x71, 0x4e, 0x58, 0x6a, 0x51, 0x31, 0xd0, 0x78, 0x85, 0xea, 0x6a, 0x3d,
		0xc7, 0xa2, 0x74, 0x83, 0xda, 0xda, 0x98, 0x3c, 0x9c, 0xf6, 0x63, 0xa8,
		0x25, 0x60, 0x1f, 0x57, 0x2c, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xed,
		0xb5, 0x19, 0x01, 0x01, 0x00, 0x00,
	},
		"i18n/resources/de_DE.all.json",
	)
}

func i18n_resources_en_us_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0xa4, 0x6a, 0xcd, 0xcd, 0x4f, 0xc9, 0x4c, 0xcb, 0x4c, 0x05, 0x59,
		0x99, 0x96, 0x98, 0x53, 0x9c, 0x0a, 0x12, 0xaf, 0xd5, 0xc1, 0xe2, 0x9e,
		0xb0, 0xd4, 0xa2, 0x62, 0xa0, 0xf9, 0x0a, 0xd5, 0xd5, 0x7a, 0x8e, 0x45,
		0xe9, 0x06, 0xb5, 0xb5, 0x31, 0x79, 0x38, 0x1d, 0x80, 0x57, 0x2d, 0x56,
		0x1b, 0xb9, 0x62, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x13, 0xa3, 0x6a,
		0xed, 0x04, 0x01, 0x00, 0x00,
	},
		"i18n/resources/en_US.all.json",
	)
}

func i18n_resources_es_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x20, 0xad, 0xc1, 0x89, 0x39, 0xa5, 0x29, 0xf9, 0xc5, 0x0a, 0x29,
		0xa9, 0xc5, 0x29, 0xa9, 0x40, 0x6d, 0x0a, 0x25, 0x99, 0xa9, 0x45, 0x45,
		0x89, 0x30, 0x73, 0x2a, 0x31, 0x0d, 0xc8, 0xcd, 0x4f, 0xc9, 0x4c, 0xcb,
		0x4c, 0x05, 0x59, 0x9c, 0x96, 0x98, 0x53, 0x9c, 0x0a, 0x12, 0xaf, 0xd5,
		0xc1, 0xe2, 0xaa, 0xb0, 0xd4, 0xa2, 0x62, 0xa0, 0x2d, 0x0a, 0xd5, 0xd5,
		0x7a, 0x8e, 0x45, 0xe9, 0x06, 0xb5, 0xb5, 0x31, 0x79, 0x38, 0x9d, 0x01,
		0x56, 0x7b, 0x78, 0x33, 0x92, 0x62, 0x02, 0x16, 0x72, 0xc5, 0x02, 0x02,
		0x00, 0x00, 0xff, 0xff, 0x04, 0x26, 0xb8, 0x64, 0x09, 0x01, 0x00, 0x00,
	},
		"i18n/resources/es.all.json",
	)
}

func i18n_resources_es_es_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x20, 0xad, 0xc1, 0x89, 0x39, 0xa5, 0x29, 0xf9, 0xc5, 0x0a, 0x29,
		0xa9, 0xc5, 0x29, 0xa9, 0x40, 0x6d, 0x0a, 0x25, 0x99, 0xa9, 0x45, 0x45,
		0x89, 0x30, 0x73, 0x2a, 0x31, 0x0d, 0xc8, 0xcd, 0x4f, 0xc9, 0x4c, 0xcb,
		0x4c, 0x05, 0x59, 0x9c, 0x96, 0x98, 0x53, 0x9c, 0x0a, 0x12, 0xaf, 0xd5,
		0xc1, 0xe2, 0xaa, 0xb0, 0xd4, 0xa2, 0x62, 0xa0, 0x2d, 0x0a, 0xd5, 0xd5,
		0x7a, 0x8e, 0x45, 0xe9, 0x06, 0xb5, 0xb5, 0x31, 0x79, 0x38, 0x9d, 0x01,
		0x56, 0x7b, 0x78, 0x33, 0x92, 0x62, 0x02, 0x16, 0x72, 0xc5, 0x02, 0x02,
		0x00, 0x00, 0xff, 0xff, 0x04, 0x26, 0xb8, 0x64, 0x09, 0x01, 0x00, 0x00,
	},
		"i18n/resources/es_ES.all.json",
	)
}

func i18n_resources_fr_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x84, 0x8e,
		0xb1, 0x0a, 0x02, 0x31, 0x0c, 0x86, 0xf7, 0x7b, 0x8a, 0x9f, 0xce, 0x87,
		0x28, 0x38, 0x88, 0x9b, 0x2e, 0xfa, 0x04, 0x2e, 0xea, 0x50, 0x68, 0x7a,
		0x54, 0x7a, 0x09, 0xa4, 0x75, 0x2a, 0x7d, 0x77, 0xdb, 0xe1, 0xa6, 0x53,
		0x0c, 0x24, 0x43, 0xc8, 0x97, 0xef, 0xbf, 0x0f, 0x00, 0x4a, 0x1f, 0xad,
		0x4c, 0x70, 0xe6, 0x08, 0x73, 0xa5, 0x18, 0x05, 0x5e, 0x65, 0xc6, 0x45,
		0xbc, 0x27, 0x45, 0xb4, 0xec, 0xd0, 0x3b, 0xec, 0x0e, 0xbc, 0x9f, 0xc4,
		0x8c, 0x0b, 0x91, 0xd5, 0x72, 0x8a, 0x36, 0x07, 0xe1, 0x8e, 0x9e, 0x85,
		0x5f, 0xf2, 0x56, 0x38, 0x6a, 0x0c, 0x32, 0xa9, 0x12, 0x28, 0x2f, 0x6f,
		0x56, 0xf4, 0x2c, 0x2e, 0xf8, 0x40, 0xdd, 0xea, 0x6d, 0x4c, 0xd4, 0xf7,
		0x75, 0xfc, 0x12, 0xe9, 0x46, 0x9a, 0x9a, 0x02, 0xa5, 0x6c, 0x4e, 0x3a,
		0x6d, 0x6b, 0x7d, 0xf0, 0xcf, 0x0c, 0xab, 0xdb, 0x3f, 0xbe, 0xe1, 0xf9,
		0x09, 0x00, 0x00, 0xff, 0xff, 0x33, 0xcb, 0xf9, 0x00, 0x05, 0x01, 0x00,
		0x00,
	},
		"i18n/resources/fr.all.json",
	)
}

func i18n_resources_fr_fr_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x8e,
		0x31, 0x0b, 0x02, 0x31, 0x0c, 0x85, 0xf7, 0xfb, 0x15, 0x8f, 0xce, 0x87,
		0x28, 0x38, 0x88, 0x9b, 0x2e, 0xfa, 0x0b, 0x5c, 0xd4, 0xa1, 0xd0, 0xf4,
		0xa8, 0xf4, 0x12, 0x48, 0xeb, 0x74, 0xdc, 0x7f, 0xb7, 0x19, 0x1c, 0xc4,
		0xd3, 0x42, 0x43, 0x48, 0xf2, 0xde, 0xfb, 0xae, 0x1d, 0x80, 0xc9, 0x4a,
		0x7b, 0x2e, 0x05, 0xb7, 0x87, 0x3b, 0x53, 0xce, 0x82, 0xa8, 0x32, 0xe2,
		0x24, 0x31, 0x92, 0x22, 0x7b, 0x0e, 0xb0, 0x9f, 0x36, 0x3b, 0xde, 0x0e,
		0xe2, 0xfa, 0xb7, 0xa2, 0xaa, 0xe7, 0x92, 0x7d, 0x4d, 0xc2, 0x26, 0x3d,
		0x0a, 0x3f, 0xe4, 0xa9, 0x08, 0xd4, 0x34, 0xa8, 0xa4, 0x4a, 0xad, 0x2f,
		0x1f, 0x3e, 0x54, 0x6d, 0xfd, 0xe5, 0x34, 0x4a, 0x48, 0x31, 0x91, 0x11,
		0x44, 0x9f, 0x0b, 0xd9, 0x7c, 0xee, 0x17, 0xf0, 0x2e, 0xa4, 0xa5, 0xc5,
		0x61, 0x9a, 0x56, 0x07, 0x1d, 0xd6, 0xf3, 0x7c, 0xe3, 0x9f, 0x3c, 0x7f,
		0x6f, 0x17, 0x13, 0xbb, 0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb8,
		0x6a, 0xd8, 0x13, 0x01, 0x00, 0x00,
	},
		"i18n/resources/fr_FR.all.json",
	)
}

func i18n_resources_zh_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x20, 0xad, 0x4f, 0xf6, 0x2e, 0x78, 0xba, 0x74, 0xef, 0xfb, 0x3d,
		0x3d, 0x4f, 0x76, 0xf7, 0xbd, 0x5c, 0x3d, 0xe3, 0xc9, 0xde, 0xd9, 0x4f,
		0x37, 0x4c, 0x79, 0x3e, 0xab, 0xe5, 0xe9, 0x9c, 0xf9, 0x4f, 0xe7, 0x6c,
		0x78, 0x3a, 0xa9, 0x07, 0xc3, 0x80, 0xdc, 0xfc, 0x94, 0xcc, 0xb4, 0xcc,
		0x54, 0x90, 0xc5, 0x69, 0x89, 0x39, 0xc5, 0xa9, 0x20, 0xf1, 0x5a, 0x1d,
		0x2c, 0xae, 0x0a, 0x4b, 0x2d, 0x2a, 0x06, 0xda, 0xa2, 0x50, 0x5d, 0xad,
		0xe7, 0x58, 0x94, 0x6e, 0x50, 0x5b, 0x1b, 0x93, 0x87, 0xd3, 0x19, 0xcf,
		0x3b, 0x3b, 0x9e, 0xcd, 0x59, 0x03, 0x57, 0x49, 0xc0, 0x36, 0xae, 0x58,
		0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x70, 0x38, 0xa1, 0x06, 0x01,
		0x00, 0x00,
	},
		"i18n/resources/zh.all.json",
	)
}

func i18n_resources_zh_hans_all_json() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8a, 0xe6,
		0x52, 0x50, 0x50, 0xa8, 0x06, 0x11, 0x40, 0xa0, 0x94, 0x99, 0xa2, 0x64,
		0xa5, 0xa0, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0xaf, 0x90, 0x56, 0x94, 0x9f,
		0xab, 0xe0, 0x9e, 0x9f, 0x96, 0x96, 0x5a, 0xa4, 0x90, 0x93, 0x98, 0x97,
		0xa2, 0x00, 0xc2, 0x99, 0x86, 0x16, 0x79, 0x26, 0xe9, 0xf9, 0x4a, 0x3a,
		0x30, 0x1d, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x39, 0x89, 0x25, 0x99, 0xf9,
		0x79, 0x20, 0xad, 0x4f, 0xf6, 0x2e, 0x78, 0xba, 0x74, 0xef, 0xfb, 0x3d,
		0x3d, 0x4f, 0x76, 0xf7, 0xbd, 0x5c, 0x3d, 0xe3, 0xc9, 0xde, 0xd9, 0x4f,
		0x37, 0x4c, 0x79, 0x3e, 0xab, 0xe5, 0xe9, 0x9c, 0xf9, 0x4f, 0xe7, 0x6c,
		0x78, 0x3a, 0xa9, 0x07, 0xc3, 0x80, 0xdc, 0xfc, 0x94, 0xcc, 0xb4, 0xcc,
		0x54, 0x90, 0xc5, 0x69, 0x89, 0x39, 0xc5, 0xa9, 0x20, 0xf1, 0x5a, 0x1d,
		0x2c, 0xae, 0x0a, 0x4b, 0x2d, 0x2a, 0x06, 0xda, 0xa2, 0x50, 0x5d, 0xad,
		0xe7, 0x58, 0x94, 0x6e, 0x50, 0x5b, 0x1b, 0x93, 0x87, 0xd3, 0x19, 0xcf,
		0x3b, 0x3b, 0x9e, 0xcd, 0x59, 0x03, 0x57, 0x49, 0xc0, 0x36, 0xae, 0x58,
		0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x70, 0x38, 0xa1, 0x06, 0x01,
		0x00, 0x00,
	},
		"i18n/resources/zh_Hans.all.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"i18n/resources/de.all.json":      i18n_resources_de_all_json,
	"i18n/resources/de_DE.all.json":   i18n_resources_de_de_all_json,
	"i18n/resources/en_US.all.json":   i18n_resources_en_us_all_json,
	"i18n/resources/es.all.json":      i18n_resources_es_all_json,
	"i18n/resources/es_ES.all.json":   i18n_resources_es_es_all_json,
	"i18n/resources/fr.all.json":      i18n_resources_fr_all_json,
	"i18n/resources/fr_FR.all.json":   i18n_resources_fr_fr_all_json,
	"i18n/resources/zh.all.json":      i18n_resources_zh_all_json,
	"i18n/resources/zh_Hans.all.json": i18n_resources_zh_hans_all_json,
}