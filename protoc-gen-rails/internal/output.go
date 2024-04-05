package internal

import (
  "bytes"
  "fmt"
  "github.com/iancoleman/strcase"
  "google.golang.org/protobuf/types/descriptorpb"
	"strings"
	"text/template"
)

type FileResult struct {
	Name string
	Content string
}

type controller struct {
	ControllerName string
	Methods []method
	ServiceName string
	FullServiceName string
	MethodName string
}

type method struct {
	Name string
	RequestType string
	Path string
	PathInfo []PathInfo
	Body string
	HttpMethod string
}

type Route struct {
	MethodName string
	Path string
	Controller string
	HttpMethod string
}

var controllerTemplate = `
####### THIS FILE IS AUTOMATICALLY GENERATED BY protoc-gen-rails. DO NOT MODIFY. #######

require 'grpc_rest'
class {{.ControllerName}}Controller < ActionController::Base
  protect_from_forgery with: :null_session

	rescue_from StandardError do |e|
		render json: GrpcRest.error_msg(e)
	end
  METHOD_PARAM_MAP = {
{{range .Methods }}
    "{{.Name}}" => [
       {{range .PathInfo -}}
			   {name: "{{.Name}}", val: {{if .HasValPattern}}"{{.ValPattern}}"{{else}}nil{{end}}, split_name:{{.SplitName}}},
			 {{end -}}
    ],
{{end -}}
  }.freeze
{{$fullServiceName := .FullServiceName -}}
{{range .Methods }}
	def {{.Name}}
    fields = {{.RequestType}}.descriptor.to_a.map(&:name)
    grpc_request = GrpcRest.init_request({{.RequestType}}, request.parameters.to_h.slice(*fields))
    GrpcRest.assign_params(grpc_request, METHOD_PARAM_MAP["{{.Name}}"], "{{.Body}}", request.parameters)
    render json: GrpcRest.send_request("{{$fullServiceName}}", "{{.Name}}", grpc_request)
  end
{{end}}
end
`

func ProcessService(service *descriptorpb.ServiceDescriptorProto, pkg string) (FileResult, []Route, error) {
  var routes []Route
	data := controller{
		Methods: []method{},
		ServiceName: Classify(service.GetName()),
		ControllerName: Demodulize(service.GetName()),
		FullServiceName: Classify(pkg + "." + service.GetName()),
	}
	for _, m := range service.GetMethod() {
		opts, err := ExtractAPIOptions(m)
		if err != nil {
			return FileResult{}, routes, err
		}
		httpMethod, path, err := MethodAndPath(opts.Pattern)
		pathInfo, err := ParsedPath(path)
		if err != nil {
			return FileResult{}, routes, err
		}
		controllerMethod := method{
			Name: strcase.ToSnake(m.GetName()),
			RequestType: Classify(m.GetInputType()),
			Path: path,
			HttpMethod: httpMethod,
			Body: opts.Body,
			PathInfo: pathInfo,
		}
		data.Methods = append(data.Methods, controllerMethod)
		routes = append(routes, Route{
			HttpMethod: strings.ToLower(httpMethod),
			Path: SanitizePath(path),
			Controller: strcase.ToSnake(data.ControllerName),
			MethodName: strcase.ToSnake(m.GetName()),
		})
	}
	resultTemplate, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		return FileResult{}, routes, fmt.Errorf("can't parse controller template: %w", err)
	}
	var resultContent bytes.Buffer
	err = resultTemplate.Execute(&resultContent, data)
	if err != nil {
		return FileResult{}, routes, fmt.Errorf("can't execute controller template: %w", err)
	}
	return FileResult{
		Content: resultContent.String(),
		Name: fmt.Sprintf("app/controllers/%s_controller.rb", strcase.ToSnake(data.ControllerName)),
	}, routes, nil
}

var routeTemplate = `# frozen_string_literal: true

####### THIS FILE IS AUTOMATICALLY GENERATED BY protoc-gen-rails. DO NOT MODIFY. #######

{{range . -}}
{{.HttpMethod}} '{{.Path}}' => '{{.Controller}}#{{.MethodName}}'
{{end -}}
`

func OutputRoutes(routes []Route) (string, error) {
	resultTemplate, err := template.New("routes").Parse(routeTemplate)
	if err != nil {
		return "", err
	}
	var resultContent bytes.Buffer
	err = resultTemplate.Execute(&resultContent, routes)
	if err != nil {
		return "", err
	}
	return resultContent.String(), nil
}
