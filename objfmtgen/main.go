package main

import (
	"flag"
	"fmt"
	"github.com/bobwong89757/protoplus/gen"
	"github.com/bobwong89757/protoplus/model"
	"github.com/bobwong89757/protoplus/msgidutil"
	"github.com/bobwong89757/protoplus/util"
	"os"
)

var (
	flagPackage         = flag.String("package", "", "package name in source files")
	flagGenSuggestMsgID = flag.Bool("GenSuggestMsgID", false, "Generate suggest msgid, default is false")
)

func genAdaptor(ctx *Context, f func(*Context) error) gen.GenFunc {

	return func(dset *model.DescriptorSet, fileName string) error {
		ctx.OutputFileName = fileName
		ctx.DescriptorSet = dset
		ctx.PackageName = *flagPackage

		return f(ctx)
	}
}

func main() {

	flag.Parse()

	var ctx Context
	ctx.DescriptorSet = new(model.DescriptorSet)

	if err := util.ParseFileList(ctx.DescriptorSet); err != nil {
		fmt.Println("ParseFileList error: ", err)
		os.Exit(1)
	}

	if *flagGenSuggestMsgID {
		msgidutil.GenSuggestMsgID(ctx.DescriptorSet)
		return
	}

	if err := GenGo(&ctx); err != nil {
		fmt.Println("Generate error: ", err)
		os.Exit(1)
	}

}
