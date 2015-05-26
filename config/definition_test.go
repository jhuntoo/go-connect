package config_test

import (
	. "github.com/jhuntoo/go-connect/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Definition", func() {
	
	Describe("Section: softwarestatement", func() {
		var softwarestatement = NewAppConfigDefinition().SoftwareStatement
		
		It("name should be correct", func() {
			Expect(softwarestatement.SectionName).To(Equal("softwarestatement"))
		})
		It("path should be correct", func() {
			Expect(softwarestatement.SectionPath).To(Equal("auth2"))
		})
		
		Context("Field: SigningMethod", func() {
		    It("key should be correct", func() {
				Expect(softwarestatement.SigningMethod.Key).To(Equal("signing_method"))
			})
			It("default value should be correct", func() {
				Expect(softwarestatement.SigningMethod.DefaultValue).To(Equal("RS256"))
			})
			It("allowed values should be correct", func() {
				Expect(softwarestatement.SigningMethod.AllowedValues).To(Equal([]string{"RS256"}))
			})
		})
	})
})
