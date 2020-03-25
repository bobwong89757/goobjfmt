package main

import (
	"fmt"
	"github.com/bobwong89757/protoplus/codegen"
)

// 报错行号+7
const goCodeTemplate = `// Generated by github.com/bobwong89757/goobjfmt/objfmtgen
// DO NOT EDIT!

package {{.PackageName}}

import (
	"fmt"
	"reflect"
	_ "github.com/bobwong89757/cellnet/codec/binary"
	"github.com/bobwong89757/cellnet"
	"github.com/bobwong89757/cellnet/codec"
)
{{range $a, $enumobj := .Enums}}
type {{.Name}} int32
const (	{{range .Fields}}
	{{$enumobj.Name}}_{{.Name}} {{.Type}} = {{TagNumber $enumobj .}} {{end}}
)

var (
{{$enumobj.Name}}MapperValueByName = map[string]int32{ {{range .Fields}}
	"{{.Name}}": {{TagNumber $enumobj .}}, {{end}}
}

{{$enumobj.Name}}MapperNameByValue = map[int32]string{ {{range .Fields}}
	{{TagNumber $enumobj .}}: "{{.Name}}" , {{end}}
}
)

func (self {{$enumobj.Name}}) String() string {
	return {{$enumobj.Name}}MapperNameByValue[int32(self)]
}
{{end}}

{{range .Structs}}
{{ObjectLeadingComment .}}
type {{.Name}} struct{	{{range .Fields}}
	{{GoFieldName .}} {{GoTypeName .}} {{GoStructTag .}}{{FieldTrailingComment .}} {{end}}
}
{{end}}
{{range .Structs}}
func (self *{{.Name}}) String() string { return fmt.Sprintf("%+v",*self) } {{end}}

func init() {
	{{range .Structs}} {{ if IsMessage . }}
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("binary"),	
		Type:  reflect.TypeOf((*{{.Name}})(nil)).Elem(),
		ID:    {{StructMsgID .}},
	}) {{end}} {{end}}
}

`

func GenGo(ctx *Context) error {

	gen := codegen.NewCodeGen("go").
		RegisterTemplateFunc(codegen.UsefulFunc).
		ParseTemplate(goCodeTemplate, ctx).
		FormatGoCode()

	if gen.Error() != nil {
		fmt.Println(string(gen.Code()))
		return gen.Error()
	}

	return gen.WriteOutputFile(ctx.OutputFileName).Error()
}
