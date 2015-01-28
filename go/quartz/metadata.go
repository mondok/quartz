package quartz

import "reflect"

type structMetadata struct {
	NameToMethodMetadata map[string]*methodMetadata
	TargetStruct         interface{} `json:"-"`
}

func newStructMetadata(targetStruct interface{}) *structMetadata {
	return &structMetadata{
		make(map[string]*methodMetadata),
		targetStruct,
	}
}

type registry map[string]*structMetadata

func newRegistry() registry {
	return make(map[string]*structMetadata)
}

type methodMetadata struct {
	Method         reflect.Method `json:"-"`
	ArgumentToType map[string]string
}
