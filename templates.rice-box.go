package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `edit.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x20, 0x2e, 0x20, 0x7d, 0x7d, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x3f, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x65, 0x64, 0x69, 0x74, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x73, 0x70, 0x65, 0x6c, 0x6c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3d, 0x22, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x22, 0x20, 0x72, 0x6f, 0x77, 0x73, 0x3d, 0x22, 0x31, 0x35, 0x22, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3d, 0x22, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x20, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x20, 0x68, 0x65, 0x72, 0x65, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x7b, 0x7b, 0x20, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x61, 0x72, 0x65, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x63, 0x6f, 0x6c, 0x2d, 0x6d, 0x64, 0x2d, 0x38, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x20, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3d, 0x22, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x7b, 0x7b, 0x20, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x20, 0x7d, 0x7d, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x63, 0x6f, 0x6c, 0x2d, 0x6d, 0x64, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3d, 0x22, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x7b, 0x7b, 0x20, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x20, 0x7d, 0x7d, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x63, 0x6f, 0x6c, 0x2d, 0x6d, 0x64, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x20, 0x62, 0x74, 0x6e, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x20, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x66, 0x6c, 0x6f, 0x70, 0x70, 0x79, 0x2d, 0x64, 0x69, 0x73, 0x6b, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x20, 0x53, 0x61, 0x76, 0x65, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x7b, 0x7b, 0x20, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x22, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x22, 0x20, 0x2e, 0x20, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `footer.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x22, 0x7d, 0x7d, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x68, 0x72, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0xa, 0x9, 0x9, 0x3c, 0x70, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x61, 0x6d, 0x2d, 0x70, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2d, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x2d, 0x43, 0x68, 0x65, 0x61, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x3e, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x20, 0x43, 0x68, 0x65, 0x61, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x3c, 0x2f, 0x61, 0x3e, 0x20, 0x7c, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x70, 0x78, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x3e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x21, 0x2d, 0x2d, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x20, 0x2d, 0x2d, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `header.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x7d, 0x7d, 0xa, 0x3c, 0x21, 0x64, 0x6f, 0x63, 0x74, 0x79, 0x70, 0x65, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0xa, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0x9, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x3d, 0x22, 0x55, 0x54, 0x46, 0x2d, 0x38, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x7b, 0x7b, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x9, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x77, 0x69, 0x6b, 0x69, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x76, 0x69, 0x65, 0x77, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x2c, 0x20, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x2d, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3d, 0x31, 0x22, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0xa, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x6f, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x24, 0x64, 0x69, 0x72, 0x20, 0x3a, 0x3d, 0x20, 0x2e, 0x44, 0x69, 0x72, 0x73, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x69, 0x66, 0x20, 0x24, 0x64, 0x69, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3e, 0x7b, 0x7b, 0x24, 0x64, 0x69, 0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6c, 0x73, 0x65, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x7b, 0x7b, 0x20, 0x24, 0x64, 0x69, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x20, 0x7d, 0x7d, 0x22, 0x3e, 0x7b, 0x7b, 0x24, 0x64, 0x69, 0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6e, 0x64, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6e, 0x6f, 0x2d, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x3e, 0x7b, 0x7b, 0x20, 0x69, 0x66, 0x20, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x7d, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x3f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3d, 0x7b, 0x7b, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x7d, 0x26, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3d, 0x31, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3e, 0x7b, 0x7b, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x7d, 0x3c, 0x2f, 0x61, 0x3e, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x69, 0x66, 0x20, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x20, 0x7c, 0x20, 0x6f, 0x72, 0x20, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x65, 0x64, 0x69, 0x74, 0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x3f, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x20, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x20, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x20, 0x65, 0x6c, 0x73, 0x65, 0x20, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x6c, 0x69, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x65, 0x64, 0x69, 0x74, 0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x3f, 0x65, 0x64, 0x69, 0x74, 0x3d, 0x31, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x20, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x65, 0x64, 0x69, 0x74, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x20, 0x45, 0x64, 0x69, 0x74, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x6f, 0x6c, 0x3e, 0xa, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    `node.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x7d, 0x7d, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x7b, 0x7b, 0x20, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x64, 0x6f, 0x77, 0x6e, 0x20, 0x7d, 0x7d, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    `revision.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x20, 0x7d, 0x7d, 0xa, 0x3c, 0x21, 0x2d, 0x2d, 0x20, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x28, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x2c, 0x20, 0x64, 0x69, 0x66, 0x66, 0x20, 0x65, 0x74, 0x63, 0x29, 0x20, 0x2d, 0x2d, 0x3e, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x74, 0x6e, 0x20, 0x62, 0x74, 0x6e, 0x2d, 0x64, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x20, 0x62, 0x74, 0x6e, 0x2d, 0x78, 0x73, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x9, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x20, 0x67, 0x6c, 0x79, 0x70, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x74, 0x65, 0x70, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0x20, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x72, 0x65, 0x76, 0x65, 0x72, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x7b, 0x7b, 0x20, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x7d, 0x7d, 0x22, 0x20, 0x2f, 0x3e, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `revisions.tpl`,
		FileModTime: time.Unix(1447182441, 0),
		Content:     string([]byte{0x7b, 0x7b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x20, 0x22, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7d, 0x7d, 0xa, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x72, 0x6f, 0x77, 0x20, 0x63, 0x6f, 0x6c, 0x22, 0x3e, 0xa, 0x9, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x24, 0x6c, 0x6f, 0x67, 0x20, 0x3a, 0x3d, 0x20, 0x2e, 0x4c, 0x6f, 0x67, 0x7d, 0x7d, 0x20, 0x7b, 0x7b, 0x69, 0x66, 0x20, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x3f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3d, 0x7b, 0x7b, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x7d, 0x7d, 0x26, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3d, 0x31, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x65, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x3f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x3d, 0x7b, 0x7b, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x7d, 0x7d, 0x26, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3d, 0x31, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x6b, 0x62, 0x64, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3e, 0x7b, 0x7b, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x7d, 0x7d, 0x3c, 0x2f, 0x6b, 0x62, 0x64, 0x3e, 0x20, 0x7b, 0x7b, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x7d, 0x7d, 0x20, 0x28, 0x7b, 0x7b, 0x24, 0x6c, 0x6f, 0x67, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x7d, 0x7d, 0x29, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x9, 0x9, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x9, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x9, 0x3c, 0x68, 0x72, 0x20, 0x2f, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1447182441, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // edit.tpl
			file3, // footer.tpl
			file4, // header.tpl
			file5, // node.tpl
			file6, // revision.tpl
			file7, // revisions.tpl

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1447185350, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"edit.tpl":      file2,
			"footer.tpl":    file3,
			"header.tpl":    file4,
			"node.tpl":      file5,
			"revision.tpl":  file6,
			"revisions.tpl": file7,
		},
	})
}
