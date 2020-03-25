package main

import "github.com/bobwong89757/protoplus/model"

type Context struct {
	*model.DescriptorSet
	OutputFileName string
	PackageName    string
}
