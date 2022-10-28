package gen

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/twirphp/twirp/protoc-gen-twirp_symfony/internal/php"
	"github.com/twirphp/twirp/protoc-gen-twirp_symfony/templates/service"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const twirpVersion = "v8.1.0"

var serviceTemplates = template.Must(template.New("").Funcs(TxtFuncMap()).ParseFS(service.FS(), "*.php.tmpl"))

type serviceFileData struct {
	File           *protogen.File
	Service        *protogen.Service
	TwirpVersion   string
	TwirphpVersion string
	Version        string
}

// Generate is the main code generator.
func Generate(plugin *protogen.Plugin, version string) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	namespaces := map[string]bool{}

	for _, file := range plugin.Files {
		for _, svc := range file.Services {
			// Namespaces are only relevant when there are services defined in them
			namespaces[php.Namespace(file)] = true

			for _, tpl := range serviceTemplates.Templates() {
				fileName := fmt.Sprintf(
					"%s/%s%s",
					php.Path(file),
					php.ServiceName(file, svc),
					tpl.Name(),
				)
				fileName = strings.TrimSuffix(fileName, ".tmpl")

				generatedFile := plugin.NewGeneratedFile(fileName, "")

				data := &serviceFileData{
					File:           file,
					Service:        svc,
					TwirpVersion:   twirpVersion,
					TwirphpVersion: version,
					Version:        version,
				}

				err := tpl.Execute(generatedFile, data)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
