// Copyright 2022 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package types

import (
	"google.golang.org/genproto/googleapis/api"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/configchange"
	"google.golang.org/genproto/googleapis/api/distribution"
	_ "google.golang.org/genproto/googleapis/api/error_reason"
	_ "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/genproto/googleapis/api/label"
	"google.golang.org/genproto/googleapis/api/metric"
	"google.golang.org/genproto/googleapis/api/monitoredres"
	"google.golang.org/genproto/googleapis/api/serviceconfig"
	_ "google.golang.org/genproto/googleapis/api/servicecontrol/v2"
	_ "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
	_ "google.golang.org/genproto/googleapis/api/serviceusage/v1"
	_ "google.golang.org/genproto/googleapis/api/visibility"

	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	Authentication              = "google.api.Authentication"
	AuthenticationRule          = "google.api.AuthenticationRule"
	AuthProvider                = "google.api.AuthProvider"
	OAuthRequirements           = "google.api.OAuthRequirements"
	AuthRequirement             = "google.api.AuthRequirement"
	Backend                     = "google.api.Backend"
	BackendRule                 = "google.api.BackendRule"
	Billing                     = "google.api.Billing"
	MethodSignatureExtension    = "google.api.method_signature"
	DefaultHostExtension        = "google.api.default_host"
	OauthScopesExtension        = "google.api.oauth_scopes"
	ConfigChange                = "google.api.ConfigChange"
	Advice                      = "google.api.Advice"
	ChangeType                  = "google.api.ChangeType"
	ProjectProperties           = "google.api.ProjectProperties"
	Property                    = "google.api.Property"
	Context                     = "google.api.Context"
	ContextRule                 = "google.api.ContextRule"
	Control                     = "google.api.Control"
	Distribution                = "google.api.Distribution"
	Documentation               = "google.api.Documentation"
	DocumentationRule           = "google.api.DocumentationRule"
	Page                        = "google.api.Page"
	Endpoint                    = "google.api.Endpoint"
	FieldBehaviorExtension      = "google.api.field_behavior"
	Http                        = "google.api.Http"
	HttpRule                    = "google.api.HttpRule"
	CustomHttpPattern           = "google.api.CustomHttpPattern"
	HttpBody                    = "google.api.HttpBody"
	LabelDescriptor             = "google.api.LabelDescriptor"
	LaunchStage                 = "google.api.LaunchStage"
	LogDescriptor               = "google.api.LogDescriptor"
	Logging                     = "google.api.Logging"
	MetricDescriptor            = "google.api.MetricDescriptor"
	Metric                      = "google.api.Metric"
	MonitoredResourceDescriptor = "google.api.MonitoredResourceDescriptor"
	MonitoredResource           = "google.api.MonitoredResource"
	MonitoredResourceMetadata   = "google.api.MonitoredResourceMetadata"
	Monitoring                  = "google.api.Monitoring"
	Quota                       = "google.api.Quota"
	MetricRule                  = "google.api.MetricRule"
	QuotaLimit                  = "google.api.QuotaLimit"
	ResourceReferenceExtension  = "google.api.resource_reference"
	ResourceDefinitionExtension = "google.api.resource_definition"
	ResourceExtension           = "google.api.resource"
	ResourceDescriptor          = "google.api.ResourceDescriptor"
	ResourceReference           = "google.api.ResourceReference"
)

const (
	AuthProto              = "google/api/auth.proto"
	BackendProto           = "google/api/backend.proto"
	BillingProto           = "google/api/billing.proto"
	ClientProto            = "google/api/client.proto"
	ConfigChangeProto      = "google/api/config_change.proto"
	ConsumerProto          = "google/api/consumer.proto"
	ContextProto           = "google/api/context.proto"
	ControlProto           = "google/api/control.proto"
	DistributionProto      = "google/api/distribution.proto"
	DocumentationProto     = "google/api/documentation.proto"
	EndpointProto          = "google/api/endpoint.proto"
	FieldBehaviorProto     = "google/api/field_behavior.proto"
	HttpProto              = "google/api/http.proto"
	HttpBodyProto          = "google/api/httpbody.proto"
	LabelProto             = "google/api/label.proto"
	LaunchStageProto       = "google/api/launch_stage.proto"
	LogProto               = "google/api/log.proto"
	LoggingProto           = "google/api/logging.proto"
	MetricProto            = "google/api/metric.proto"
	MonitoredResourceProto = "google/api/monitored_resource.proto"
	MonitoringProto        = "google/api/monitoring.proto"
	QuotaProto             = "google/api/quota.proto"
	ResourceProto          = "google/api/resource.proto"
	// TODO(zchee): add
	// service.proto
	// source_info.proto
	// system_parameter.proto
	// usage.proto
)

var KnownCommonImports = map[string]string{
	Authentication:              AuthProto,
	AuthenticationRule:          AuthProto,
	AuthProvider:                AuthProto,
	OAuthRequirements:           AuthProto,
	AuthRequirement:             AuthProto,
	Backend:                     BackendProto,
	BackendRule:                 BackendProto,
	Billing:                     BillingProto,
	MethodSignatureExtension:    ClientProto,
	DefaultHostExtension:        ClientProto,
	OauthScopesExtension:        ClientProto,
	ConfigChange:                ConfigChangeProto,
	Advice:                      ConfigChangeProto,
	ChangeType:                  ConfigChangeProto,
	ProjectProperties:           ConsumerProto,
	Property:                    ConsumerProto,
	Context:                     ContextProto,
	ContextRule:                 ContextProto,
	Control:                     ControlProto,
	Distribution:                DistributionProto,
	Documentation:               DocumentationProto,
	DocumentationRule:           DocumentationProto,
	Page:                        DocumentationProto,
	Endpoint:                    EndpointProto,
	FieldBehaviorExtension:      FieldBehaviorProto,
	Http:                        HttpProto,
	HttpRule:                    HttpProto,
	CustomHttpPattern:           HttpProto,
	HttpBody:                    HttpBodyProto,
	LabelDescriptor:             LabelProto,
	LaunchStage:                 LaunchStageProto,
	LogDescriptor:               LogProto,
	Logging:                     LoggingProto,
	MetricDescriptor:            MetricProto,
	Metric:                      MetricProto,
	MonitoredResourceDescriptor: MonitoredResourceProto,
	MonitoredResource:           MonitoredResourceProto,
	MonitoredResourceMetadata:   MonitoredResourceProto,
	Monitoring:                  MonitoringProto,
	Quota:                       QuotaProto,
	MetricRule:                  QuotaProto,
	QuotaLimit:                  QuotaProto,
	ResourceReferenceExtension:  ResourceProto,
	ResourceDefinitionExtension: ResourceProto,
	ResourceExtension:           ResourceProto,
	ResourceDescriptor:          ResourceProto,
	ResourceReference:           ResourceProto,
}

func AuthDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_auth_proto)
}

func BackendDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_backend_proto)
}

func BillingDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_billing_proto)
}

func ClientDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(annotations.File_google_api_client_proto)
}

func ConfigChangeDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(configchange.File_google_api_config_change_proto)
}

func ConsumerDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_consumer_proto)
}

func ContextDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_context_proto)
}

func DistributionDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(distribution.File_google_api_distribution_proto)
}

func DocumentationDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_documentation_proto)
}

func EndpointDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_endpoint_proto)
}

func FieldBehaviorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(annotations.File_google_api_field_behavior_proto)
}

func HttpDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(annotations.File_google_api_http_proto)
}

func HttpBodyDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(httpbody.File_google_api_httpbody_proto)
}

func LabelDescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(label.File_google_api_label_proto)
}

func LaunchStageDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(api.File_google_api_launch_stage_proto)
}

func LogDescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_log_proto)
}

func LoggingDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_logging_proto)
}

func MetricDescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(metric.File_google_api_metric_proto)
}

func MonitoredResourceDescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(monitoredres.File_google_api_monitored_resource_proto)
}

func MonitoringDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_monitoring_proto)
}

func QuotaDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(serviceconfig.File_google_api_quota_proto)
}

func ResourceDescriptorDescriptor() *descriptorpb.FileDescriptorProto {
	return protodesc.ToFileDescriptorProto(annotations.File_google_api_resource_proto)
}

// func WrappersDescriptor() *descriptorpb.FileDescriptorProto {
// 	return protodesc.ToFileDescriptorProto(wrapperspb.File_google_protobuf_wrappers_proto)
// }

// func WrappersDescriptor() *descriptorpb.FileDescriptorProto {
// 	return protodesc.ToFileDescriptorProto(wrapperspb.File_google_protobuf_wrappers_proto)
// }

var KnownCommonDescriptor = map[string]*descriptorpb.FileDescriptorProto{
	AuthProto:              AuthDescriptor(),
	BackendProto:           BackendDescriptor(),
	BillingProto:           BillingDescriptor(),
	ClientProto:            ClientDescriptor(),
	ConfigChangeProto:      ConfigChangeDescriptor(),
	ConsumerProto:          ConsumerDescriptor(),
	ContextProto:           ContextDescriptor(),
	DistributionProto:      DistributionDescriptor(),
	DocumentationProto:     DocumentationDescriptor(),
	EndpointProto:          EndpointDescriptor(),
	FieldBehaviorProto:     FieldBehaviorDescriptor(),
	HttpProto:              HttpDescriptor(),
	HttpBodyProto:          HttpBodyDescriptor(),
	LabelProto:             LabelDescriptorDescriptor(),
	LaunchStageProto:       LaunchStageDescriptor(),
	LogProto:               LogDescriptorDescriptor(),
	LoggingProto:           LoggingDescriptor(),
	MetricProto:            MetricDescriptorDescriptor(),
	MonitoredResourceProto: MonitoredResourceDescriptorDescriptor(),
	MonitoringProto:        MonitoringDescriptor(),
	QuotaProto:             QuotaDescriptor(),
	ResourceProto:          ResourceDescriptorDescriptor(),
}
