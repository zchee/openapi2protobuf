module go.lsp.dev/openapi2protobuf

go 1.18

require (
	github.com/getkin/kin-openapi v0.94.1-0.20220403132454-ebcbb7269761
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/golang/protobuf v1.5.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jhump/protoreflect v1.12.1-0.20220417024638-438db461d753
	google.golang.org/protobuf v1.28.1-0.20220524200550-784c48255455
)

// fix for CVE-2022-28948
replace gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.38.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
