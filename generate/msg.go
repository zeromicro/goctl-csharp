package generate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/goctl-csharp/template"
	"github.com/zeromicro/goctl-csharp/util"
)

func genMessages(dir string, ns string, api *spec.ApiSpec) error {
	for _, t := range api.Types {
		tn := t.Name()
		definedType, ok := t.(spec.DefineStruct)
		if !ok {
			return fmt.Errorf("type %s not supported", tn)
		}

		cn := lo.PascalCase(tn)

		data := template.CSharpApiMessageTemplateData{
			CSharpTemplateData: template.CSharpTemplateData{Namespace: ns},
			MessageName:        cn,
			Fields:             []template.CSharpApiMessageFieldTemplateData{},
		}

		for _, m := range definedType.Members {
			k, err := m.GetPropertyName()
			if err != nil {
				return err
			}
			tn, err := apiTypeToCsTypeName(m.Type)
			if err != nil {
				return err
			}

			tag, _, ok := strings.Cut(strings.Trim(m.Tag, "`"), ":")
			if !ok {
				return fmt.Errorf("type %s not tag: %s", tn, m.Tag)
			}

			f := template.CSharpApiMessageFieldTemplateData{
				FieldName:  lo.PascalCase(m.Name),
				KeyName:    k,
				TypeName:   tn,
				IsOptional: util.IsOptionalOrOmitEmpty(m),
				Tag:        tag,
			}
			data.Fields = append(data.Fields, f)
		}

		if err := template.WriteFile(dir, cn, template.ApiMessage, data); err != nil {
			return err
		}
	}
	return nil
}

func apiTypeToCsTypeName(t spec.Type) (string, error) {
	switch tt := t.(type) {
	case spec.PrimitiveType:
		r, ok := primitiveType(t.Name())
		if !ok {
			return "", errors.New("unsupported primitive type " + t.Name())
		}

		return r, nil
	case spec.ArrayType:
		et, err := apiTypeToCsTypeName(tt.Value)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("List<%s>", et), nil
	case spec.MapType:
		vt, err := apiTypeToCsTypeName(tt.Value)
		if err != nil {
			return "", err
		}
		kt, ok := primitiveType(tt.Key)
		if !ok {
			return "", errors.New("unsupported key is not primitive type " + t.Name())
		}
		return fmt.Sprintf("Dictionary<%s,%s>", kt, vt), nil
	case spec.DefineStruct:
		return t.Name(), nil
	}

	return "", errors.New("unsupported type " + t.Name())
}

func primitiveType(tp string) (string, bool) {
	switch tp {
	case "string", "int", "uint", "bool", "byte":
		return tp, true
	case "int8":
		return "SByte", true
	case "uint8":
		return "byte", true
	case "int16", "int32", "int64":
		return util.UpperHead(tp, 1), true
	case "uint16", "uint32", "uint64":
		return util.UpperHead(tp, 2), true
	case "float", "float32":
		return "float", true
	case "float64":
		return "double", true
	}
	return "", false
}
