// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackendRequestDataMaskingHeadersInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendRequestDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendRequestDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendRequestDataMaskingInitParameters struct {

	// A headers block as defined below.
	Headers []BackendRequestDataMaskingHeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []BackendRequestDataMaskingQueryParamsInitParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendRequestDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []BackendRequestDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []BackendRequestDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendRequestDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []BackendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []BackendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type BackendRequestDataMaskingQueryParamsInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendRequestDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendRequestDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingHeadersInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendResponseDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendResponseDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackendResponseDataMaskingQueryParamsInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendResponseDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type BackendResponseDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticBackendRequestInitParameters struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *BackendRequestDataMaskingInitParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendRequestObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *BackendRequestDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking *BackendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendResponseDataMaskingInitParameters struct {

	// A headers block as defined below.
	Headers []BackendResponseDataMaskingHeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []BackendResponseDataMaskingQueryParamsInitParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticBackendResponseDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []BackendResponseDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []BackendResponseDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticBackendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []BackendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []BackendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticBackendResponseInitParameters struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticBackendResponseDataMaskingInitParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendResponseObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticBackendResponseDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticBackendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking *DiagnosticBackendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingHeadersInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingInitParameters struct {

	// A headers block as defined below.
	Headers []DiagnosticFrontendRequestDataMaskingHeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []DiagnosticFrontendRequestDataMaskingQueryParamsInitParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []DiagnosticFrontendRequestDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []DiagnosticFrontendRequestDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendRequestDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendRequestDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingQueryParamsInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendRequestInitParameters struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticFrontendRequestDataMaskingInitParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendRequestObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticFrontendRequestDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendRequestParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking *DiagnosticFrontendRequestDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingHeadersInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingHeadersObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingHeadersParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingInitParameters struct {

	// A headers block as defined below.
	Headers []DiagnosticFrontendResponseDataMaskingHeadersInitParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []DiagnosticFrontendResponseDataMaskingQueryParamsInitParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingObservation struct {

	// A headers block as defined below.
	Headers []DiagnosticFrontendResponseDataMaskingHeadersObservation `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	QueryParams []DiagnosticFrontendResponseDataMaskingQueryParamsObservation `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingParameters struct {

	// A headers block as defined below.
	// +kubebuilder:validation:Optional
	Headers []DiagnosticFrontendResponseDataMaskingHeadersParameters `json:"headers,omitempty" tf:"headers,omitempty"`

	// A query_params block as defined below.
	// +kubebuilder:validation:Optional
	QueryParams []DiagnosticFrontendResponseDataMaskingQueryParamsParameters `json:"queryParams,omitempty" tf:"query_params,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingQueryParamsInitParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingQueryParamsObservation struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseDataMaskingQueryParamsParameters struct {

	// The data masking mode. Possible values are Mask and Hide for query_params. The only possible value is Mask for headers.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The name of the header or the query parameter to mask.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type DiagnosticFrontendResponseInitParameters struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticFrontendResponseDataMaskingInitParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendResponseObservation struct {

	// Number of payload bytes to log (up to 8192).
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	DataMasking *DiagnosticFrontendResponseDataMaskingObservation `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticFrontendResponseParameters struct {

	// Number of payload bytes to log (up to 8192).
	// +kubebuilder:validation:Optional
	BodyBytes *float64 `json:"bodyBytes,omitempty" tf:"body_bytes,omitempty"`

	// A data_masking block as defined below.
	// +kubebuilder:validation:Optional
	DataMasking *DiagnosticFrontendResponseDataMaskingParameters `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies a list of headers to log.
	// +kubebuilder:validation:Optional
	// +listType=set
	HeadersToLog []*string `json:"headersToLog,omitempty" tf:"headers_to_log,omitempty"`
}

type DiagnosticInitParameters struct {

	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Logger
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// Reference to a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDRef *v1.Reference `json:"apiManagementLoggerIdRef,omitempty" tf:"-"`

	// Selector for a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDSelector *v1.Selector `json:"apiManagementLoggerIdSelector,omitempty" tf:"-"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	BackendRequest *DiagnosticBackendRequestInitParameters `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	BackendResponse *DiagnosticBackendResponseInitParameters `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	FrontendRequest *DiagnosticFrontendRequestInitParameters `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	FrontendResponse *DiagnosticFrontendResponseInitParameters `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// Log client IP address.
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url.
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

type DiagnosticObservation struct {

	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	BackendRequest *DiagnosticBackendRequestObservation `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	BackendResponse *DiagnosticBackendResponseObservation `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	FrontendRequest *DiagnosticFrontendRequestObservation `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	FrontendResponse *DiagnosticFrontendResponseObservation `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// The ID of the API Management Diagnostic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Log client IP address.
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url.
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

type DiagnosticParameters struct {

	// The id of the target API Management Logger where the API Management Diagnostic should be saved.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Logger
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementLoggerID *string `json:"apiManagementLoggerId,omitempty" tf:"api_management_logger_id,omitempty"`

	// Reference to a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDRef *v1.Reference `json:"apiManagementLoggerIdRef,omitempty" tf:"-"`

	// Selector for a Logger in apimanagement to populate apiManagementLoggerId.
	// +kubebuilder:validation:Optional
	APIManagementLoggerIDSelector *v1.Selector `json:"apiManagementLoggerIdSelector,omitempty" tf:"-"`

	// The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta2.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	// +kubebuilder:validation:Optional
	AlwaysLogErrors *bool `json:"alwaysLogErrors,omitempty" tf:"always_log_errors,omitempty"`

	// A backend_request block as defined below.
	// +kubebuilder:validation:Optional
	BackendRequest *DiagnosticBackendRequestParameters `json:"backendRequest,omitempty" tf:"backend_request,omitempty"`

	// A backend_response block as defined below.
	// +kubebuilder:validation:Optional
	BackendResponse *DiagnosticBackendResponseParameters `json:"backendResponse,omitempty" tf:"backend_response,omitempty"`

	// A frontend_request block as defined below.
	// +kubebuilder:validation:Optional
	FrontendRequest *DiagnosticFrontendRequestParameters `json:"frontendRequest,omitempty" tf:"frontend_request,omitempty"`

	// A frontend_response block as defined below.
	// +kubebuilder:validation:Optional
	FrontendResponse *DiagnosticFrontendResponseParameters `json:"frontendResponse,omitempty" tf:"frontend_response,omitempty"`

	// The HTTP Correlation Protocol to use. Possible values are None, Legacy or W3C.
	// +kubebuilder:validation:Optional
	HTTPCorrelationProtocol *string `json:"httpCorrelationProtocol,omitempty" tf:"http_correlation_protocol,omitempty"`

	// Log client IP address.
	// +kubebuilder:validation:Optional
	LogClientIP *bool `json:"logClientIp,omitempty" tf:"log_client_ip,omitempty"`

	// The format of the Operation Name for Application Insights telemetries. Possible values are Name, and Url.
	// +kubebuilder:validation:Optional
	OperationNameFormat *string `json:"operationNameFormat,omitempty" tf:"operation_name_format,omitempty"`

	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Sampling (%). For high traffic APIs, please read this documentation to understand performance implications and log sampling. Valid values are between 0.0 and 100.0.
	// +kubebuilder:validation:Optional
	SamplingPercentage *float64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage,omitempty"`

	// Logging verbosity. Possible values are verbose, information or error.
	// +kubebuilder:validation:Optional
	Verbosity *string `json:"verbosity,omitempty" tf:"verbosity,omitempty"`
}

// DiagnosticSpec defines the desired state of Diagnostic
type DiagnosticSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiagnosticParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DiagnosticInitParameters `json:"initProvider,omitempty"`
}

// DiagnosticStatus defines the observed state of Diagnostic.
type DiagnosticStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiagnosticObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Diagnostic is the Schema for the Diagnostics API. Manages an API Management Service Diagnostic.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Diagnostic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiagnosticSpec   `json:"spec"`
	Status            DiagnosticStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiagnosticList contains a list of Diagnostics
type DiagnosticList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Diagnostic `json:"items"`
}

// Repository type metadata.
var (
	Diagnostic_Kind             = "Diagnostic"
	Diagnostic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Diagnostic_Kind}.String()
	Diagnostic_KindAPIVersion   = Diagnostic_Kind + "." + CRDGroupVersion.String()
	Diagnostic_GroupVersionKind = CRDGroupVersion.WithKind(Diagnostic_Kind)
)

func init() {
	SchemeBuilder.Register(&Diagnostic{}, &DiagnosticList{})
}
