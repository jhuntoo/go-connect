package config

import (
	"reflect"
	"fmt"
)
type Loader interface {
	Load() AppConfig
}

type FileLoader struct {
	
}

func(l *FileLoader) Load() AppConfig {
	appConfig := AppConfig{}
	sections := GetSections(appConfig)
	
	println(sections)
	
	return appConfig
}

func GetSections(x interface{}) map[string]reflect.Type {
	sections := make(map[string]reflect.Type)
	typ := reflect.TypeOf(x)
	
	for i := 0; i < typ.NumField(); i++ {
    p := typ.Field(i)
      if !p.Anonymous {
        sections[p.Name] = p.Type
      }
     }
	fmt.Printf("%v", sections)
	return sections
}

