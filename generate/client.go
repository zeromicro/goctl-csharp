package generate

import (
	"fmt"
	"strings"

	_ "embed"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/goctl-csharp/template"
	"github.com/zeromicro/goctl-csharp/util"
)

func genClient(dir string, ns string, api *spec.ApiSpec) error {
	data := template.CSharpTemplateData{
		Namespace: ns,
	}
	if err := template.WriteFile(dir, "ApiAttribute", template.ApiAttribute, data); err != nil {
		return err
	}
	if err := template.WriteFile(dir, "ApiBodyJsonContent", template.ApiBodyJsonContent, data); err != nil {
		return err
	}
	if err := template.WriteFile(dir, "ApiException", template.ApiException, data); err != nil {
		return err
	}
	if err := template.WriteFile(dir, "ApiBaseClient", template.ApiBaseClient, data); err != nil {
		return err
	}

	return writeClient(dir, ns, api)
}

func writeClient(dir string, ns string, api *spec.ApiSpec) error {
	name := lo.PascalCase(api.Service.Name)

	data := template.CSharpApiClientTemplateData{
		CSharpTemplateData: template.CSharpTemplateData{Namespace: ns},
		ClientName:         name,
		Routes:             []template.CSharpApiClientRouteTemplateData{},
	}

	// 组
	for _, g := range api.Service.Groups {
		prefix := g.GetAnnotation("prefix")
		p := lo.PascalCase(prefix)

		// 路由
		for _, r := range g.Routes {
			an := lo.PascalCase(r.Path)
			method := util.UpperHead(strings.ToLower(r.Method), 1)

			route := template.CSharpApiClientRouteTemplateData{
				HttpMethod:   method,
				Prefix:       prefix,
				ActionPrefix: p,
				ActionName:   an,
				UrlPath:      r.Path,
			}

			if r.ResponseType != nil {
				rn := r.ResponseType.Name()
				route.ResponseType = &rn
			}

			if r.RequestType != nil {
				rn := r.RequestType.Name()
				route.RequestType = &rn
			}

			data.Routes = append(data.Routes, route)
		}
	}

	return template.WriteFile(dir, fmt.Sprintf("%sClient", name), template.ApiClient, data)
}
