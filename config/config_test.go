package config_test

import (
	"reflect"
	. "github.com/jhuntoo/go-connect/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	
	Describe("Tags: SoftwareStatementConfig", func() {
		var config = SoftwareStatementConfig{}
		
		sectionTagGetter := func(c interface{}, fieldName string) string {
			field, _ := reflect.TypeOf(c).FieldByName(fieldName)
			return field.Tag.Get("section")
		}
		keyTagGetter := func(c interface{}, fieldName string) string {
			field, _ := reflect.TypeOf(c).FieldByName(fieldName)
			return field.Tag.Get("key")
		}
				
		
		Context("Field: SigningMethod", func() {
		    It("section tag should be correct", func() {
				Expect(sectionTagGetter(config,"SigningMethod")).To(Equal("softwarestatement"))
			})
			It("key tag should be correct", func() {
				Expect(keyTagGetter(config,"SigningMethod")).To(Equal("signing_method"))
			})
		})
	})
})
