package config_test

import (
	. "github.com/jhuntoo/go-connect/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load", func() {
	Describe("FileLoader", func() {
		
		loader := &FileLoader{}
		
		_ = loader.Load();
		
		It("x", func() {
			Expect(1).To(Equal(1))
		})
	})

})
