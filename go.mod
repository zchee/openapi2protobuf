module go.lsp.dev/openapi2protobuf

go 1.18

require (
	github.com/Code-Hex/dd v1.1.0
	github.com/alecthomas/chroma v0.10.0
	github.com/bytedance/sonic v1.3.0
	github.com/getkin/kin-openapi v0.96.1-0.20220601174204-648d6b9a170b
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/golang/protobuf v1.5.2
	github.com/iancoleman/strcase v0.2.0
	github.com/jhump/protoreflect v1.12.1-0.20220417024638-438db461d753
	google.golang.org/genproto v0.0.0-20220601144221-27df5f98adab
	google.golang.org/protobuf v1.28.1-0.20220524200550-784c48255455
)

// fix for CVE-2022-28948
replace gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1

require (
	github.com/chenzhuoyu/base64x v0.0.0-20220526154910-8bf9453eb81a // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.0.0-20190626092158-b2ccc519800e // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/grpc v1.46.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
