// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

// Package ml provides access to the Cloud Machine Learning Engine.
//
// See https://cloud.google.com/ml/
//
// Usage example:
//
//   import "google.golang.org/api/ml/v1"
//   ...
//   mlService, err := ml.New(oauthHttpClient)
package ml // import "google.golang.org/api/ml/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "ml:v1"
const apiName = "ml"
const apiVersion = "v1"
const basePath = "https://ml.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Jobs = NewProjectsJobsService(s)
	rs.Locations = NewProjectsLocationsService(s)
	rs.Models = NewProjectsModelsService(s)
	rs.Operations = NewProjectsOperationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Jobs *ProjectsJobsService

	Locations *ProjectsLocationsService

	Models *ProjectsModelsService

	Operations *ProjectsOperationsService
}

func NewProjectsJobsService(s *Service) *ProjectsJobsService {
	rs := &ProjectsJobsService{s: s}
	return rs
}

type ProjectsJobsService struct {
	s *Service
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	return rs
}

type ProjectsLocationsService struct {
	s *Service
}

func NewProjectsModelsService(s *Service) *ProjectsModelsService {
	rs := &ProjectsModelsService{s: s}
	rs.Versions = NewProjectsModelsVersionsService(s)
	return rs
}

type ProjectsModelsService struct {
	s *Service

	Versions *ProjectsModelsVersionsService
}

func NewProjectsModelsVersionsService(s *Service) *ProjectsModelsVersionsService {
	rs := &ProjectsModelsVersionsService{s: s}
	return rs
}

type ProjectsModelsVersionsService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

// GoogleApi__HttpBody: Message that represents an arbitrary HTTP body.
// It should only be used for
// payload formats that can't be represented as JSON, such as raw binary
// or
// an HTML page.
//
//
// This message can be used both in streaming and non-streaming API
// methods in
// the request as well as the response.
//
// It can be used as a top-level request field, which is convenient if
// one
// wants to extract parameters from either the URL or HTTP template into
// the
// request fields and also want access to the raw HTTP body.
//
// Example:
//
//     message GetResourceRequest {
//       // A unique request id.
//       string request_id = 1;
//
//       // The raw HTTP body is bound to this field.
//       google.api.HttpBody http_body = 2;
//     }
//
//     service ResourceService {
//       rpc GetResource(GetResourceRequest) returns
// (google.api.HttpBody);
//       rpc UpdateResource(google.api.HttpBody) returns
// (google.protobuf.Empty);
//     }
//
// Example with streaming methods:
//
//     service CaldavService {
//       rpc GetCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//       rpc UpdateCalendar(stream google.api.HttpBody)
//         returns (stream google.api.HttpBody);
//     }
//
// Use of this type only changes how the request and response bodies
// are
// handled, all other features will continue to work unchanged.
type GoogleApi__HttpBody struct {
	// ContentType: The HTTP Content-Type string representing the content
	// type of the body.
	ContentType string `json:"contentType,omitempty"`

	// Data: HTTP body binary data.
	Data string `json:"data,omitempty"`

	// Extensions: Application specific response metadata. Must be set in
	// the first response
	// for streaming APIs.
	Extensions []googleapi.RawMessage `json:"extensions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ContentType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentType") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleApi__HttpBody) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleApi__HttpBody
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1HyperparameterOutputHyperparameterMetric: An observed
// value of a metric.
type GoogleCloudMlV1HyperparameterOutputHyperparameterMetric struct {
	// ObjectiveValue: The objective value at this training step.
	ObjectiveValue float64 `json:"objectiveValue,omitempty"`

	// TrainingStep: The global training step for this metric.
	TrainingStep int64 `json:"trainingStep,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ObjectiveValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ObjectiveValue") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1HyperparameterOutputHyperparameterMetric) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1HyperparameterOutputHyperparameterMetric
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudMlV1HyperparameterOutputHyperparameterMetric) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudMlV1HyperparameterOutputHyperparameterMetric
	var s1 struct {
		ObjectiveValue gensupport.JSONFloat64 `json:"objectiveValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ObjectiveValue = float64(s1.ObjectiveValue)
	return nil
}

// GoogleCloudMlV1__AcceleratorConfig: Represents a hardware accelerator
// request config.
type GoogleCloudMlV1__AcceleratorConfig struct {
	// Count: The number of accelerators to attach to each machine running
	// the job.
	Count int64 `json:"count,omitempty,string"`

	// Type: The available types of accelerators.
	//
	// Possible values:
	//   "ACCELERATOR_TYPE_UNSPECIFIED" - Unspecified accelerator type.
	// Default to no GPU.
	//   "NVIDIA_TESLA_K80" - Nvidia Tesla K80 GPU.
	//   "NVIDIA_TESLA_P100" - Nvidia Tesla P100 GPU.
	//   "NVIDIA_TESLA_V100" - Nvidia Tesla V100 GPU.
	//   "NVIDIA_TESLA_P4" - Nvidia Tesla P4 GPU.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__AcceleratorConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__AcceleratorConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__AutoScaling: Options for automatically scaling a
// model.
type GoogleCloudMlV1__AutoScaling struct {
	// MinNodes: Optional. The minimum number of nodes to allocate for this
	// model. These
	// nodes are always up, starting from the time the model is
	// deployed.
	// Therefore, the cost of operating this model will be at least
	// `rate` * `min_nodes` * number of hours since last billing
	// cycle,
	// where `rate` is the cost per node-hour as documented in the
	// [pricing guide](/ml-engine/docs/pricing),
	// even if no predictions are performed. There is additional cost for
	// each
	// prediction performed.
	//
	// Unlike manual scaling, if the load gets too heavy for the nodes
	// that are up, the service will automatically add nodes to handle
	// the
	// increased load as well as scale back as traffic drops, always
	// maintaining
	// at least `min_nodes`. You will be charged for the time in which
	// additional
	// nodes are used.
	//
	// If not specified, `min_nodes` defaults to 0, in which case, when
	// traffic
	// to a model stops (and after a cool-down period), nodes will be shut
	// down
	// and no charges will be incurred until traffic to the model
	// resumes.
	//
	// You can set `min_nodes` when creating the model version, and you can
	// also
	// update `min_nodes` for an existing
	// version:
	// <pre>
	// update_body.json:
	// {
	//   'autoScaling': {
	//     'minNodes': 5
	//   }
	// }
	// </pre>
	// HTTP request:
	// <pre>
	// PATCH
	// https://ml.googleapis.com/v1/{name=projects/*/models/*/versions/*}?update_mask=autoScaling.minNodes -d
	// @./update_body.json
	// </pre>
	MinNodes int64 `json:"minNodes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MinNodes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MinNodes") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__AutoScaling) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__AutoScaling
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__CancelJobRequest: Request message for the CancelJob
// method.
type GoogleCloudMlV1__CancelJobRequest struct {
}

type GoogleCloudMlV1__Capability struct {
	// AvailableAccelerators: Available accelerators for the capability.
	//
	// Possible values:
	//   "ACCELERATOR_TYPE_UNSPECIFIED" - Unspecified accelerator type.
	// Default to no GPU.
	//   "NVIDIA_TESLA_K80" - Nvidia Tesla K80 GPU.
	//   "NVIDIA_TESLA_P100" - Nvidia Tesla P100 GPU.
	//   "NVIDIA_TESLA_V100" - Nvidia Tesla V100 GPU.
	//   "NVIDIA_TESLA_P4" - Nvidia Tesla P4 GPU.
	AvailableAccelerators []string `json:"availableAccelerators,omitempty"`

	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "TRAINING"
	//   "BATCH_PREDICTION"
	//   "ONLINE_PREDICTION"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AvailableAccelerators") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailableAccelerators") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Capability) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Capability
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GoogleCloudMlV1__Config struct {
	// TpuServiceAccount: The service account Cloud ML uses to run on TPU
	// node.
	TpuServiceAccount string `json:"tpuServiceAccount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TpuServiceAccount")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TpuServiceAccount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Config) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Config
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__GetConfigResponse: Returns service account
// information associated with a project.
type GoogleCloudMlV1__GetConfigResponse struct {
	Config *GoogleCloudMlV1__Config `json:"config,omitempty"`

	// ServiceAccount: The service account Cloud ML uses to access resources
	// in the project.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// ServiceAccountProject: The project number for `service_account`.
	ServiceAccountProject int64 `json:"serviceAccountProject,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Config") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Config") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__GetConfigResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__GetConfigResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__HyperparameterOutput: Represents the result of a
// single hyperparameter tuning trial from a
// training job. The TrainingOutput object that is returned on
// successful
// completion of a training job with hyperparameter tuning includes a
// list
// of HyperparameterOutput objects, one for each successful trial.
type GoogleCloudMlV1__HyperparameterOutput struct {
	// AllMetrics: All recorded object metrics for this trial. This field is
	// not currently
	// populated.
	AllMetrics []*GoogleCloudMlV1HyperparameterOutputHyperparameterMetric `json:"allMetrics,omitempty"`

	// FinalMetric: The final objective metric seen for this trial.
	FinalMetric *GoogleCloudMlV1HyperparameterOutputHyperparameterMetric `json:"finalMetric,omitempty"`

	// Hyperparameters: The hyperparameters given to this trial.
	Hyperparameters map[string]string `json:"hyperparameters,omitempty"`

	// IsTrialStoppedEarly: True if the trial is stopped early.
	IsTrialStoppedEarly bool `json:"isTrialStoppedEarly,omitempty"`

	// TrialId: The trial id for these results.
	TrialId string `json:"trialId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllMetrics") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AllMetrics") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__HyperparameterOutput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__HyperparameterOutput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__HyperparameterSpec: Represents a set of
// hyperparameters to optimize.
type GoogleCloudMlV1__HyperparameterSpec struct {
	// Algorithm: Optional. The search algorithm specified for the
	// hyperparameter
	// tuning job.
	// Uses the default CloudML Engine hyperparameter tuning
	// algorithm if unspecified.
	//
	// Possible values:
	//   "ALGORITHM_UNSPECIFIED" - The default algorithm used by
	// hyperparameter tuning service.
	//   "GRID_SEARCH" - Simple grid search within the feasible space. To
	// use grid search,
	// all parameters must be `INTEGER`, `CATEGORICAL`, or `DISCRETE`.
	//   "RANDOM_SEARCH" - Simple random search within the feasible space.
	Algorithm string `json:"algorithm,omitempty"`

	// EnableTrialEarlyStopping: Optional. Indicates if the hyperparameter
	// tuning job enables auto trial
	// early stopping.
	EnableTrialEarlyStopping bool `json:"enableTrialEarlyStopping,omitempty"`

	// Goal: Required. The type of goal to use for tuning. Available types
	// are
	// `MAXIMIZE` and `MINIMIZE`.
	//
	// Defaults to `MAXIMIZE`.
	//
	// Possible values:
	//   "GOAL_TYPE_UNSPECIFIED" - Goal Type will default to maximize.
	//   "MAXIMIZE" - Maximize the goal metric.
	//   "MINIMIZE" - Minimize the goal metric.
	Goal string `json:"goal,omitempty"`

	// HyperparameterMetricTag: Optional. The Tensorflow summary tag name to
	// use for optimizing trials. For
	// current versions of Tensorflow, this tag name should exactly match
	// what is
	// shown in Tensorboard, including all scopes.  For versions of
	// Tensorflow
	// prior to 0.12, this should be only the tag passed to tf.Summary.
	// By default, "training/hptuning/metric" will be used.
	HyperparameterMetricTag string `json:"hyperparameterMetricTag,omitempty"`

	// MaxParallelTrials: Optional. The number of training trials to run
	// concurrently.
	// You can reduce the time it takes to perform hyperparameter tuning by
	// adding
	// trials in parallel. However, each trail only benefits from the
	// information
	// gained in completed trials. That means that a trial does not get
	// access to
	// the results of trials running at the same time, which could reduce
	// the
	// quality of the overall optimization.
	//
	// Each trial will use the same scale tier and machine types.
	//
	// Defaults to one.
	MaxParallelTrials int64 `json:"maxParallelTrials,omitempty"`

	// MaxTrials: Optional. How many training trials should be attempted to
	// optimize
	// the specified hyperparameters.
	//
	// Defaults to one.
	MaxTrials int64 `json:"maxTrials,omitempty"`

	// Params: Required. The set of parameters to tune.
	Params []*GoogleCloudMlV1__ParameterSpec `json:"params,omitempty"`

	// ResumePreviousJobId: Optional. The prior hyperparameter tuning job id
	// that users hope to
	// continue with. The job id will be used to find the corresponding
	// vizier
	// study guid and resume the study.
	ResumePreviousJobId string `json:"resumePreviousJobId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Algorithm") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Algorithm") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__HyperparameterSpec) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__HyperparameterSpec
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__Job: Represents a training or prediction job.
type GoogleCloudMlV1__Job struct {
	// CreateTime: Output only. When the job was created.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: Output only. When the job processing was completed.
	EndTime string `json:"endTime,omitempty"`

	// ErrorMessage: Output only. The details of a failure or a
	// cancellation.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a job from overwriting each other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform job updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `GetJob`,
	// and
	// systems are expected to put that etag in the request to `UpdateJob`
	// to
	// ensure that their change will be applied to the same version of the
	// job.
	Etag string `json:"etag,omitempty"`

	// JobId: Required. The user-specified id of the job.
	JobId string `json:"jobId,omitempty"`

	// Labels: Optional. One or more labels that you can add, to organize
	// your jobs.
	// Each label is a key-value pair, where both the key and the value
	// are
	// arbitrary strings that you supply.
	// For more information, see the documentation on
	// <a href="/ml-engine/docs/tensorflow/resource-labels">using
	// labels</a>.
	Labels map[string]string `json:"labels,omitempty"`

	// PredictionInput: Input parameters to create a prediction job.
	PredictionInput *GoogleCloudMlV1__PredictionInput `json:"predictionInput,omitempty"`

	// PredictionOutput: The current prediction job result.
	PredictionOutput *GoogleCloudMlV1__PredictionOutput `json:"predictionOutput,omitempty"`

	// StartTime: Output only. When the job processing was started.
	StartTime string `json:"startTime,omitempty"`

	// State: Output only. The detailed state of a job.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The job state is unspecified.
	//   "QUEUED" - The job has been just created and processing has not yet
	// begun.
	//   "PREPARING" - The service is preparing to run the job.
	//   "RUNNING" - The job is in progress.
	//   "SUCCEEDED" - The job completed successfully.
	//   "FAILED" - The job failed.
	// `error_message` should contain the details of the failure.
	//   "CANCELLING" - The job is being cancelled.
	// `error_message` should describe the reason for the cancellation.
	//   "CANCELLED" - The job has been cancelled.
	// `error_message` should describe the reason for the cancellation.
	State string `json:"state,omitempty"`

	// TrainingInput: Input parameters to create a training job.
	TrainingInput *GoogleCloudMlV1__TrainingInput `json:"trainingInput,omitempty"`

	// TrainingOutput: The current training job result.
	TrainingOutput *GoogleCloudMlV1__TrainingOutput `json:"trainingOutput,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Job) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Job
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__ListJobsResponse: Response message for the ListJobs
// method.
type GoogleCloudMlV1__ListJobsResponse struct {
	// Jobs: The list of jobs.
	Jobs []*GoogleCloudMlV1__Job `json:"jobs,omitempty"`

	// NextPageToken: Optional. Pass this token as the `page_token` field of
	// the request for a
	// subsequent call.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Jobs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Jobs") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ListJobsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ListJobsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GoogleCloudMlV1__ListLocationsResponse struct {
	// Locations: Locations where at least one type of CMLE capability is
	// available.
	Locations []*GoogleCloudMlV1__Location `json:"locations,omitempty"`

	// NextPageToken: Optional. Pass this token as the `page_token` field of
	// the request for a
	// subsequent call.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Locations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locations") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ListLocationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ListLocationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__ListModelsResponse: Response message for the
// ListModels method.
type GoogleCloudMlV1__ListModelsResponse struct {
	// Models: The list of models.
	Models []*GoogleCloudMlV1__Model `json:"models,omitempty"`

	// NextPageToken: Optional. Pass this token as the `page_token` field of
	// the request for a
	// subsequent call.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Models") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Models") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ListModelsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ListModelsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__ListVersionsResponse: Response message for the
// ListVersions method.
type GoogleCloudMlV1__ListVersionsResponse struct {
	// NextPageToken: Optional. Pass this token as the `page_token` field of
	// the request for a
	// subsequent call.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Versions: The list of versions.
	Versions []*GoogleCloudMlV1__Version `json:"versions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ListVersionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ListVersionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GoogleCloudMlV1__Location struct {
	// Capabilities: Capabilities available in the location.
	Capabilities []*GoogleCloudMlV1__Capability `json:"capabilities,omitempty"`

	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Capabilities") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Capabilities") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Location) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Location
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__ManualScaling: Options for manually scaling a model.
type GoogleCloudMlV1__ManualScaling struct {
	// Nodes: The number of nodes to allocate for this model. These nodes
	// are always up,
	// starting from the time the model is deployed, so the cost of
	// operating
	// this model will be proportional to `nodes` * number of hours
	// since
	// last billing cycle plus the cost for each prediction performed.
	Nodes int64 `json:"nodes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Nodes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Nodes") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ManualScaling) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ManualScaling
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__Model: Represents a machine learning solution.
//
// A model can have multiple versions, each of which is a deployed,
// trained
// model ready to receive prediction requests. The model itself is just
// a
// container.
type GoogleCloudMlV1__Model struct {
	// DefaultVersion: Output only. The default version of the model. This
	// version will be used to
	// handle prediction requests that do not specify a version.
	//
	// You can change the default version by
	// calling
	// [projects.methods.versions.setDefault](/ml-engine/reference/re
	// st/v1/projects.models.versions/setDefault).
	DefaultVersion *GoogleCloudMlV1__Version `json:"defaultVersion,omitempty"`

	// Description: Optional. The description specified for the model when
	// it was created.
	Description string `json:"description,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a model from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform model updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `GetModel`,
	// and
	// systems are expected to put that etag in the request to `UpdateModel`
	// to
	// ensure that their change will be applied to the model as intended.
	Etag string `json:"etag,omitempty"`

	// Labels: Optional. One or more labels that you can add, to organize
	// your models.
	// Each label is a key-value pair, where both the key and the value
	// are
	// arbitrary strings that you supply.
	// For more information, see the documentation on
	// <a href="/ml-engine/docs/tensorflow/resource-labels">using
	// labels</a>.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: Required. The name specified for the model when it was
	// created.
	//
	// The model name must be unique within the project it is created in.
	Name string `json:"name,omitempty"`

	// OnlinePredictionLogging: Optional. If true, enables StackDriver
	// Logging for online prediction.
	// Default is false.
	OnlinePredictionLogging bool `json:"onlinePredictionLogging,omitempty"`

	// Regions: Optional. The list of regions where the model is going to be
	// deployed.
	// Currently only one region per model is supported.
	// Defaults to 'us-central1' if nothing is set.
	// See the <a href="/ml-engine/docs/tensorflow/regions">available
	// regions</a>
	// for ML Engine services.
	// Note:
	// *   No matter where a model is deployed, it can always be accessed
	// by
	//     users from anywhere, both for online and batch prediction.
	// *   The region for a batch prediction job is set by the region field
	// when
	//     submitting the batch prediction job and does not take its value
	// from
	//     this field.
	Regions []string `json:"regions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DefaultVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DefaultVersion") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Model) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Model
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__OperationMetadata: Represents the metadata of the
// long-running operation.
type GoogleCloudMlV1__OperationMetadata struct {
	// CreateTime: The time the operation was submitted.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: The time operation processing completed.
	EndTime string `json:"endTime,omitempty"`

	// IsCancellationRequested: Indicates whether a request to cancel this
	// operation has been made.
	IsCancellationRequested bool `json:"isCancellationRequested,omitempty"`

	// Labels: The user labels, inherited from the model or the model
	// version being
	// operated on.
	Labels map[string]string `json:"labels,omitempty"`

	// ModelName: Contains the name of the model associated with the
	// operation.
	ModelName string `json:"modelName,omitempty"`

	// OperationType: The operation type.
	//
	// Possible values:
	//   "OPERATION_TYPE_UNSPECIFIED" - Unspecified operation type.
	//   "CREATE_VERSION" - An operation to create a new version.
	//   "DELETE_VERSION" - An operation to delete an existing version.
	//   "DELETE_MODEL" - An operation to delete an existing model.
	//   "UPDATE_MODEL" - An operation to update an existing model.
	//   "UPDATE_VERSION" - An operation to update an existing version.
	//   "UPDATE_CONFIG" - An operation to update project configuration.
	OperationType string `json:"operationType,omitempty"`

	// ProjectNumber: Contains the project number associated with the
	// operation.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// StartTime: The time operation processing started.
	StartTime string `json:"startTime,omitempty"`

	// Version: Contains the version associated with the operation.
	Version *GoogleCloudMlV1__Version `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CreateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__OperationMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__OperationMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__ParameterSpec: Represents a single hyperparameter to
// optimize.
type GoogleCloudMlV1__ParameterSpec struct {
	// CategoricalValues: Required if type is `CATEGORICAL`. The list of
	// possible categories.
	CategoricalValues []string `json:"categoricalValues,omitempty"`

	// DiscreteValues: Required if type is `DISCRETE`.
	// A list of feasible points.
	// The list should be in strictly increasing order. For instance,
	// this
	// parameter might have possible settings of 1.5, 2.5, and 4.0. This
	// list
	// should not contain more than 1,000 values.
	DiscreteValues []float64 `json:"discreteValues,omitempty"`

	// MaxValue: Required if type is `DOUBLE` or `INTEGER`. This
	// field
	// should be unset if type is `CATEGORICAL`. This value should be
	// integers if
	// type is `INTEGER`.
	MaxValue float64 `json:"maxValue,omitempty"`

	// MinValue: Required if type is `DOUBLE` or `INTEGER`. This
	// field
	// should be unset if type is `CATEGORICAL`. This value should be
	// integers if
	// type is INTEGER.
	MinValue float64 `json:"minValue,omitempty"`

	// ParameterName: Required. The parameter name must be unique amongst
	// all ParameterConfigs in
	// a HyperparameterSpec message. E.g., "learning_rate".
	ParameterName string `json:"parameterName,omitempty"`

	// ScaleType: Optional. How the parameter should be scaled to the
	// hypercube.
	// Leave unset for categorical parameters.
	// Some kind of scaling is strongly recommended for real or
	// integral
	// parameters (e.g., `UNIT_LINEAR_SCALE`).
	//
	// Possible values:
	//   "NONE" - By default, no scaling is applied.
	//   "UNIT_LINEAR_SCALE" - Scales the feasible space to (0, 1) linearly.
	//   "UNIT_LOG_SCALE" - Scales the feasible space logarithmically to (0,
	// 1). The entire feasible
	// space must be strictly positive.
	//   "UNIT_REVERSE_LOG_SCALE" - Scales the feasible space "reverse"
	// logarithmically to (0, 1). The result
	// is that values close to the top of the feasible space are spread out
	// more
	// than points near the bottom. The entire feasible space must be
	// strictly
	// positive.
	ScaleType string `json:"scaleType,omitempty"`

	// Type: Required. The type of the parameter.
	//
	// Possible values:
	//   "PARAMETER_TYPE_UNSPECIFIED" - You must specify a valid type. Using
	// this unspecified type will result in
	// an error.
	//   "DOUBLE" - Type for real-valued parameters.
	//   "INTEGER" - Type for integral parameters.
	//   "CATEGORICAL" - The parameter is categorical, with a value chosen
	// from the categories
	// field.
	//   "DISCRETE" - The parameter is real valued, with a fixed set of
	// feasible points. If
	// `type==DISCRETE`, feasible_points must be provided, and
	// {`min_value`, `max_value`} will be ignored.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CategoricalValues")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CategoricalValues") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__ParameterSpec) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__ParameterSpec
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudMlV1__ParameterSpec) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudMlV1__ParameterSpec
	var s1 struct {
		MaxValue gensupport.JSONFloat64 `json:"maxValue"`
		MinValue gensupport.JSONFloat64 `json:"minValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.MaxValue = float64(s1.MaxValue)
	s.MinValue = float64(s1.MinValue)
	return nil
}

// GoogleCloudMlV1__PredictRequest: Request for predictions to be issued
// against a trained model.
type GoogleCloudMlV1__PredictRequest struct {
	// HttpBody:
	// Required. The prediction request body.
	HttpBody *GoogleApi__HttpBody `json:"httpBody,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HttpBody") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "HttpBody") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__PredictRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__PredictRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__PredictionInput: Represents input parameters for a
// prediction job. Next field: 19
type GoogleCloudMlV1__PredictionInput struct {
	// Accelerator: Optional. The type and number of accelerators to be
	// attached to each
	// machine running the job.
	Accelerator *GoogleCloudMlV1__AcceleratorConfig `json:"accelerator,omitempty"`

	// BatchSize: Optional. Number of records per batch, defaults to 64.
	// The service will buffer batch_size number of records in memory
	// before
	// invoking one Tensorflow prediction call internally. So take the
	// record
	// size and memory available into consideration when setting this
	// parameter.
	BatchSize int64 `json:"batchSize,omitempty,string"`

	// DataFormat: Required. The format of the input data files.
	//
	// Possible values:
	//   "DATA_FORMAT_UNSPECIFIED" - Unspecified format.
	//   "JSON" - Each line of the file is a JSON dictionary representing
	// one record.
	//   "TEXT" - Deprecated. Use JSON instead.
	//   "TF_RECORD" - INPUT ONLY. The source file is a TFRecord file.
	//   "TF_RECORD_GZIP" - INPUT ONLY. The source file is a GZIP-compressed
	// TFRecord file.
	//   "CSV" - OUTPUT ONLY. Output values will be in comma-separated rows,
	// with keys
	// in a separate file.
	DataFormat string `json:"dataFormat,omitempty"`

	// InputPaths: Required. The Google Cloud Storage location of the input
	// data files.
	// May contain wildcards.
	InputPaths []string `json:"inputPaths,omitempty"`

	// MaxWorkerCount: Optional. The maximum number of workers to be used
	// for parallel processing.
	// Defaults to 10 if not specified.
	MaxWorkerCount int64 `json:"maxWorkerCount,omitempty,string"`

	// ModelName: Use this field if you want to use the default version for
	// the specified
	// model. The string must use the following
	// format:
	//
	// "projects/YOUR_PROJECT/models/YOUR_MODEL"
	ModelName string `json:"modelName,omitempty"`

	// OutputDataFormat: Optional. Format of the output data files, defaults
	// to JSON.
	//
	// Possible values:
	//   "DATA_FORMAT_UNSPECIFIED" - Unspecified format.
	//   "JSON" - Each line of the file is a JSON dictionary representing
	// one record.
	//   "TEXT" - Deprecated. Use JSON instead.
	//   "TF_RECORD" - INPUT ONLY. The source file is a TFRecord file.
	//   "TF_RECORD_GZIP" - INPUT ONLY. The source file is a GZIP-compressed
	// TFRecord file.
	//   "CSV" - OUTPUT ONLY. Output values will be in comma-separated rows,
	// with keys
	// in a separate file.
	OutputDataFormat string `json:"outputDataFormat,omitempty"`

	// OutputPath: Required. The output Google Cloud Storage location.
	OutputPath string `json:"outputPath,omitempty"`

	// Region: Required. The Google Compute Engine region to run the
	// prediction job in.
	// See the <a href="/ml-engine/docs/tensorflow/regions">available
	// regions</a>
	// for ML Engine services.
	Region string `json:"region,omitempty"`

	// RuntimeVersion: Optional. The Google Cloud ML runtime version to use
	// for this batch
	// prediction. If not set, Google Cloud ML will pick the runtime version
	// used
	// during the CreateVersion request for this model version, or choose
	// the
	// latest stable version when model version information is not
	// available
	// such as when the model is specified by uri.
	RuntimeVersion string `json:"runtimeVersion,omitempty"`

	// SignatureName: Optional. The name of the signature defined in the
	// SavedModel to use for
	// this job. Please refer
	// to
	// [SavedModel](https://tensorflow.github.io/serving/serving_basic.htm
	// l)
	// for information about how to use signatures.
	//
	// Defaults
	// to
	// [DEFAULT_SERVING_SIGNATURE_DEF_KEY](https://www.tensorflow.org/api_
	// docs/python/tf/saved_model/signature_constants)
	// , which is "serving_default".
	SignatureName string `json:"signatureName,omitempty"`

	// Uri: Use this field if you want to specify a Google Cloud Storage
	// path for
	// the model to use.
	Uri string `json:"uri,omitempty"`

	// VersionName: Use this field if you want to specify a version of the
	// model to use. The
	// string is formatted the same way as `model_version`, with the
	// addition
	// of the version
	// information:
	//
	// "projects/YOUR_PROJECT/models/YOUR_MODEL/versions/YOUR_
	// VERSION"
	VersionName string `json:"versionName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Accelerator") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Accelerator") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__PredictionInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__PredictionInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__PredictionOutput: Represents results of a prediction
// job.
type GoogleCloudMlV1__PredictionOutput struct {
	// ErrorCount: The number of data instances which resulted in errors.
	ErrorCount int64 `json:"errorCount,omitempty,string"`

	// NodeHours: Node hours used by the batch prediction job.
	NodeHours float64 `json:"nodeHours,omitempty"`

	// OutputPath: The output Google Cloud Storage location provided at the
	// job creation time.
	OutputPath string `json:"outputPath,omitempty"`

	// PredictionCount: The number of generated predictions.
	PredictionCount int64 `json:"predictionCount,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ErrorCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ErrorCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__PredictionOutput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__PredictionOutput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudMlV1__PredictionOutput) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudMlV1__PredictionOutput
	var s1 struct {
		NodeHours gensupport.JSONFloat64 `json:"nodeHours"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.NodeHours = float64(s1.NodeHours)
	return nil
}

// GoogleCloudMlV1__SetDefaultVersionRequest: Request message for the
// SetDefaultVersion request.
type GoogleCloudMlV1__SetDefaultVersionRequest struct {
}

// GoogleCloudMlV1__TrainingInput: Represents input parameters for a
// training job. When using the
// gcloud command to submit your training job, you can specify
// the input parameters as command-line arguments and/or in a YAML
// configuration
// file referenced from the --config command-line argument. For
// details, see the guide to
// <a href="/ml-engine/docs/tensorflow/training-jobs">submitting a
// training
// job</a>.
type GoogleCloudMlV1__TrainingInput struct {
	// Args: Optional. Command line arguments to pass to the program.
	Args []string `json:"args,omitempty"`

	// Hyperparameters: Optional. The set of Hyperparameters to tune.
	Hyperparameters *GoogleCloudMlV1__HyperparameterSpec `json:"hyperparameters,omitempty"`

	// JobDir: Optional. A Google Cloud Storage path in which to store
	// training outputs
	// and other data needed for training. This path is passed to your
	// TensorFlow
	// program as the '--job-dir' command-line argument. The benefit of
	// specifying
	// this field is that Cloud ML validates the path for use in training.
	JobDir string `json:"jobDir,omitempty"`

	// MasterType: Optional. Specifies the type of virtual machine to use
	// for your training
	// job's master worker.
	//
	// The following types are supported:
	//
	// <dl>
	//   <dt>standard</dt>
	//   <dd>
	//   A basic machine configuration suitable for training simple models
	// with
	//   small to moderate datasets.
	//   </dd>
	//   <dt>large_model</dt>
	//   <dd>
	//   A machine with a lot of memory, specially suited for parameter
	// servers
	//   when your model is large (having many hidden layers or layers with
	// very
	//   large numbers of nodes).
	//   </dd>
	//   <dt>complex_model_s</dt>
	//   <dd>
	//   A machine suitable for the master and workers of the cluster when
	// your
	//   model requires more computation than the standard machine can
	// handle
	//   satisfactorily.
	//   </dd>
	//   <dt>complex_model_m</dt>
	//   <dd>
	//   A machine with roughly twice the number of cores and roughly double
	// the
	//   memory of <i>complex_model_s</i>.
	//   </dd>
	//   <dt>complex_model_l</dt>
	//   <dd>
	//   A machine with roughly twice the number of cores and roughly double
	// the
	//   memory of <i>complex_model_m</i>.
	//   </dd>
	//   <dt>standard_gpu</dt>
	//   <dd>
	//   A machine equivalent to <i>standard</i> that
	//   also includes a single NVIDIA Tesla K80 GPU. See more about
	//   <a href="/ml-engine/docs/tensorflow/using-gpus">using GPUs to
	//   train your model</a>.
	//   </dd>
	//   <dt>complex_model_m_gpu</dt>
	//   <dd>
	//   A machine equivalent to <i>complex_model_m</i> that also includes
	//   four NVIDIA Tesla K80 GPUs.
	//   </dd>
	//   <dt>complex_model_l_gpu</dt>
	//   <dd>
	//   A machine equivalent to <i>complex_model_l</i> that also includes
	//   eight NVIDIA Tesla K80 GPUs.
	//   </dd>
	//   <dt>standard_p100</dt>
	//   <dd>
	//   A machine equivalent to <i>standard</i> that
	//   also includes a single NVIDIA Tesla P100 GPU.
	//   </dd>
	//   <dt>complex_model_m_p100</dt>
	//   <dd>
	//   A machine equivalent to <i>complex_model_m</i> that also includes
	//   four NVIDIA Tesla P100 GPUs.
	//   </dd>
	//   <dt>standard_v100</dt>
	//   <dd>
	//   A machine equivalent to <i>standard</i> that
	//   also includes a single NVIDIA Tesla V100 GPU. The availability of
	// these
	//   GPUs is in the <i>Beta</i> launch stage.
	//   </dd>
	//   <dt>large_model_v100</dt>
	//   <dd>
	//   A machine equivalent to <i>large_model</i> that
	//   also includes a single NVIDIA Tesla V100 GPU. The availability of
	// these
	//   GPUs is in the <i>Beta</i> launch stage.
	//   </dd>
	//   <dt>complex_model_m_v100</dt>
	//   <dd>
	//   A machine equivalent to <i>complex_model_m</i> that
	//   also includes four NVIDIA Tesla V100 GPUs. The availability of
	// these
	//   GPUs is in the <i>Beta</i> launch stage.
	//   </dd>
	//   <dt>complex_model_l_v100</dt>
	//   <dd>
	//   A machine equivalent to <i>complex_model_l</i> that
	//   also includes eight NVIDIA Tesla V100 GPUs. The availability of
	// these
	//   GPUs is in the <i>Beta</i> launch stage.
	//   </dd>
	//   <dt>cloud_tpu</dt>
	//   <dd>
	//   A TPU VM including one Cloud TPU. See more about
	//   <a href="/ml-engine/docs/tensorflow/using-tpus">using TPUs to
	// train
	//   your model</a>.
	//   </dd>
	// </dl>
	//
	// You must set this value when `scaleTier` is set to `CUSTOM`.
	MasterType string `json:"masterType,omitempty"`

	// PackageUris: Required. The Google Cloud Storage location of the
	// packages with
	// the training program and any additional dependencies.
	// The maximum number of package URIs is 100.
	PackageUris []string `json:"packageUris,omitempty"`

	// ParameterServerCount: Optional. The number of parameter server
	// replicas to use for the training
	// job. Each replica in the cluster will be of the type specified
	// in
	// `parameter_server_type`.
	//
	// This value can only be used when `scale_tier` is set to `CUSTOM`.If
	// you
	// set this value, you must also set `parameter_server_type`.
	ParameterServerCount int64 `json:"parameterServerCount,omitempty,string"`

	// ParameterServerType: Optional. Specifies the type of virtual machine
	// to use for your training
	// job's parameter server.
	//
	// The supported values are the same as those described in the entry
	// for
	// `master_type`.
	//
	// This value must be present when `scaleTier` is set to `CUSTOM`
	// and
	// `parameter_server_count` is greater than zero.
	ParameterServerType string `json:"parameterServerType,omitempty"`

	// PythonModule: Required. The Python module name to run after
	// installing the packages.
	PythonModule string `json:"pythonModule,omitempty"`

	// PythonVersion: Optional. The version of Python used in training. If
	// not set, the default
	// version is '2.7'. Python '3.5' is available when `runtime_version` is
	// set
	// to '1.4' and above. Python '2.7' works with all supported runtime
	// versions.
	PythonVersion string `json:"pythonVersion,omitempty"`

	// Region: Required. The Google Compute Engine region to run the
	// training job in.
	// See the <a href="/ml-engine/docs/tensorflow/regions">available
	// regions</a>
	// for ML Engine services.
	Region string `json:"region,omitempty"`

	// RuntimeVersion: Optional. The Google Cloud ML runtime version to use
	// for training.  If not
	// set, Google Cloud ML will choose a stable version, which is defined
	// in the
	// documentation of runtime version list.
	RuntimeVersion string `json:"runtimeVersion,omitempty"`

	// ScaleTier: Required. Specifies the machine types, the number of
	// replicas for workers
	// and parameter servers.
	//
	// Possible values:
	//   "BASIC" - A single worker instance. This tier is suitable for
	// learning how to use
	// Cloud ML, and for experimenting with new models using small datasets.
	//   "STANDARD_1" - Many workers and a few parameter servers.
	//   "PREMIUM_1" - A large number of workers with many parameter
	// servers.
	//   "BASIC_GPU" - A single worker instance [with
	// a
	// GPU](/ml-engine/docs/tensorflow/using-gpus).
	//   "BASIC_TPU" - A single worker instance with a
	// [Cloud TPU](/ml-engine/docs/tensorflow/using-tpus).
	//   "CUSTOM" - The CUSTOM tier is not a set tier, but rather enables
	// you to use your
	// own cluster specification. When you use this tier, set values
	// to
	// configure your processing cluster according to these guidelines:
	//
	// *   You _must_ set `TrainingInput.masterType` to specify the type
	//     of machine to use for your master node. This is the only
	// required
	//     setting.
	//
	// *   You _may_ set `TrainingInput.workerCount` to specify the number
	// of
	//     workers to use. If you specify one or more workers, you _must_
	// also
	//     set `TrainingInput.workerType` to specify the type of machine to
	// use
	//     for your worker nodes.
	//
	// *   You _may_ set `TrainingInput.parameterServerCount` to specify
	// the
	//     number of parameter servers to use. If you specify one or more
	//     parameter servers, you _must_ also set
	//     `TrainingInput.parameterServerType` to specify the type of
	// machine to
	//     use for your parameter servers.
	//
	// Note that all of your workers must use the same machine type, which
	// can
	// be different from your parameter server type and master type.
	// Your
	// parameter servers must likewise use the same machine type, which can
	// be
	// different from your worker type and master type.
	ScaleTier string `json:"scaleTier,omitempty"`

	// WorkerCount: Optional. The number of worker replicas to use for the
	// training job. Each
	// replica in the cluster will be of the type specified in
	// `worker_type`.
	//
	// This value can only be used when `scale_tier` is set to `CUSTOM`. If
	// you
	// set this value, you must also set `worker_type`.
	WorkerCount int64 `json:"workerCount,omitempty,string"`

	// WorkerType: Optional. Specifies the type of virtual machine to use
	// for your training
	// job's worker nodes.
	//
	// The supported values are the same as those described in the entry
	// for
	// `masterType`.
	//
	// This value must be present when `scaleTier` is set to `CUSTOM`
	// and
	// `workerCount` is greater than zero.
	WorkerType string `json:"workerType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Args") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Args") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__TrainingInput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__TrainingInput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleCloudMlV1__TrainingOutput: Represents results of a training
// job. Output only.
type GoogleCloudMlV1__TrainingOutput struct {
	// CompletedTrialCount: The number of hyperparameter tuning trials that
	// completed successfully.
	// Only set for hyperparameter tuning jobs.
	CompletedTrialCount int64 `json:"completedTrialCount,omitempty,string"`

	// ConsumedMLUnits: The amount of ML units consumed by the job.
	ConsumedMLUnits float64 `json:"consumedMLUnits,omitempty"`

	// IsHyperparameterTuningJob: Whether this job is a hyperparameter
	// tuning job.
	IsHyperparameterTuningJob bool `json:"isHyperparameterTuningJob,omitempty"`

	// Trials: Results for individual Hyperparameter trials.
	// Only set for hyperparameter tuning jobs.
	Trials []*GoogleCloudMlV1__HyperparameterOutput `json:"trials,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CompletedTrialCount")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CompletedTrialCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__TrainingOutput) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__TrainingOutput
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *GoogleCloudMlV1__TrainingOutput) UnmarshalJSON(data []byte) error {
	type NoMethod GoogleCloudMlV1__TrainingOutput
	var s1 struct {
		ConsumedMLUnits gensupport.JSONFloat64 `json:"consumedMLUnits"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ConsumedMLUnits = float64(s1.ConsumedMLUnits)
	return nil
}

// GoogleCloudMlV1__Version: Represents a version of the model.
//
// Each version is a trained model deployed in the cloud, ready to
// handle
// prediction requests. A model can have multiple versions. You can
// get
// information about all of the versions of a given model by
// calling
// [projects.models.versions.list](/ml-engine/reference/rest/v1/p
// rojects.models.versions/list).
type GoogleCloudMlV1__Version struct {
	// AutoScaling: Automatically scale the number of nodes used to serve
	// the model in
	// response to increases and decreases in traffic. Care should be
	// taken to ramp up traffic according to the model's ability to scale
	// or you will start seeing increases in latency and 429 response codes.
	AutoScaling *GoogleCloudMlV1__AutoScaling `json:"autoScaling,omitempty"`

	// CreateTime: Output only. The time the version was created.
	CreateTime string `json:"createTime,omitempty"`

	// DeploymentUri: Required. The Google Cloud Storage location of the
	// trained model used to
	// create the version. See the
	// [guide to
	// model
	// deployment](/ml-engine/docs/tensorflow/deploying-models) for
	// more
	// information.
	//
	// When passing Version
	// to
	// [projects.models.versions.create](/ml-engine/reference/rest/v1/proj
	// ects.models.versions/create)
	// the model service uses the specified location as the source of the
	// model.
	// Once deployed, the model version is hosted by the prediction service,
	// so
	// this location is useful only as a historical record.
	// The total number of model files can't exceed 1000.
	DeploymentUri string `json:"deploymentUri,omitempty"`

	// Description: Optional. The description specified for the version when
	// it was created.
	Description string `json:"description,omitempty"`

	// ErrorMessage: Output only. The details of a failure or a
	// cancellation.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a model from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform model updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `GetVersion`,
	// and
	// systems are expected to put that etag in the request to
	// `UpdateVersion` to
	// ensure that their change will be applied to the model as intended.
	Etag string `json:"etag,omitempty"`

	// Framework: Optional. The machine learning framework Cloud ML Engine
	// uses to train
	// this version of the model. Valid values are `TENSORFLOW`,
	// `SCIKIT_LEARN`,
	// `XGBOOST`. If you do not specify a framework, Cloud ML Engine
	// will analyze files in the deployment_uri to determine a framework. If
	// you
	// choose `SCIKIT_LEARN` or `XGBOOST`, you must also set the runtime
	// version
	// of the model to 1.4 or greater.
	//
	// Possible values:
	//   "FRAMEWORK_UNSPECIFIED" - Unspecified framework. Defaults to
	// TensorFlow.
	//   "TENSORFLOW" - Tensorflow framework.
	//   "SCIKIT_LEARN" - Scikit-learn framework.
	//   "XGBOOST" - XGBoost framework.
	Framework string `json:"framework,omitempty"`

	// IsDefault: Output only. If true, this version will be used to handle
	// prediction
	// requests that do not specify a version.
	//
	// You can change the default version by
	// calling
	// [projects.methods.versions.setDefault](/ml-engine/reference/re
	// st/v1/projects.models.versions/setDefault).
	IsDefault bool `json:"isDefault,omitempty"`

	// Labels: Optional. One or more labels that you can add, to organize
	// your model
	// versions. Each label is a key-value pair, where both the key and the
	// value
	// are arbitrary strings that you supply.
	// For more information, see the documentation on
	// <a href="/ml-engine/docs/tensorflow/resource-labels">using
	// labels</a>.
	Labels map[string]string `json:"labels,omitempty"`

	// LastUseTime: Output only. The time the version was last used for
	// prediction.
	LastUseTime string `json:"lastUseTime,omitempty"`

	// MachineType: Optional. The type of machine on which to serve the
	// model. Currently only
	// applies to online prediction service.
	// The following are currently supported and will be deprecated in
	// Beta release.
	//   mls1-highmem-1    1 core    2 Gb RAM
	//   mls1-highcpu-4    4 core    2 Gb RAM
	// The following are available in Beta:
	//   mls1-c1-m2        1 core    2 Gb RAM   Default
	//   mls1-c4-m2        4 core    2 Gb RAM
	MachineType string `json:"machineType,omitempty"`

	// ManualScaling: Manually select the number of nodes to use for serving
	// the
	// model. You should generally use `auto_scaling` with an
	// appropriate
	// `min_nodes` instead, but this option is available if you want
	// more
	// predictable billing. Beware that latency and error rates will
	// increase
	// if the traffic exceeds that capability of the system to serve it
	// based
	// on the selected number of nodes.
	ManualScaling *GoogleCloudMlV1__ManualScaling `json:"manualScaling,omitempty"`

	// Name: Required.The name specified for the version when it was
	// created.
	//
	// The version name must be unique within the model it is created in.
	Name string `json:"name,omitempty"`

	// PythonVersion: Optional. The version of Python used in prediction. If
	// not set, the default
	// version is '2.7'. Python '3.5' is available when `runtime_version` is
	// set
	// to '1.4' and above. Python '2.7' works with all supported runtime
	// versions.
	PythonVersion string `json:"pythonVersion,omitempty"`

	// RuntimeVersion: Optional. The Google Cloud ML runtime version to use
	// for this deployment.
	// If not set, Google Cloud ML will choose a version.
	RuntimeVersion string `json:"runtimeVersion,omitempty"`

	// State: Output only. The state of a version.
	//
	// Possible values:
	//   "UNKNOWN" - The version state is unspecified.
	//   "READY" - The version is ready for prediction.
	//   "CREATING" - The version is being created. New UpdateVersion and
	// DeleteVersion
	// requests will fail if a version is in the CREATING state.
	//   "FAILED" - The version failed to be created, possibly
	// cancelled.
	// `error_message` should contain the details of the failure.
	//   "DELETING" - The version is being deleted. New UpdateVersion and
	// DeleteVersion
	// requests will fail if a version is in the DELETING state.
	//   "UPDATING" - The version is being updated. New UpdateVersion and
	// DeleteVersion
	// requests will fail if a version is in the UPDATING state.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AutoScaling") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AutoScaling") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleCloudMlV1__Version) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleCloudMlV1__Version
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__AuditConfig: Specifies the audit configuration for a
// service.
// The configuration determines which permission types are logged, and
// what
// identities, if any, are exempted from logging.
// An AuditConfig must have one or more AuditLogConfigs.
//
// If there are AuditConfigs for both `allServices` and a specific
// service,
// the union of the two AuditConfigs is used for that service: the
// log_types
// specified in each AuditConfig are enabled, and the exempted_members
// in each
// AuditLogConfig are exempted.
//
// Example Policy with multiple AuditConfigs:
//
//     {
//       "audit_configs": [
//         {
//           "service": "allServices"
//           "audit_log_configs": [
//             {
//               "log_type": "DATA_READ",
//               "exempted_members": [
//                 "user:foo@gmail.com"
//               ]
//             },
//             {
//               "log_type": "DATA_WRITE",
//             },
//             {
//               "log_type": "ADMIN_READ",
//             }
//           ]
//         },
//         {
//           "service": "fooservice.googleapis.com"
//           "audit_log_configs": [
//             {
//               "log_type": "DATA_READ",
//             },
//             {
//               "log_type": "DATA_WRITE",
//               "exempted_members": [
//                 "user:bar@gmail.com"
//               ]
//             }
//           ]
//         }
//       ]
//     }
//
// For fooservice, this policy enables DATA_READ, DATA_WRITE and
// ADMIN_READ
// logging. It also exempts foo@gmail.com from DATA_READ logging,
// and
// bar@gmail.com from DATA_WRITE logging.
type GoogleIamV1__AuditConfig struct {
	// AuditLogConfigs: The configuration for logging of each type of
	// permission.
	AuditLogConfigs []*GoogleIamV1__AuditLogConfig `json:"auditLogConfigs,omitempty"`

	// Service: Specifies a service that will be enabled for audit
	// logging.
	// For example, `storage.googleapis.com`,
	// `cloudsql.googleapis.com`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditLogConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditLogConfigs") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__AuditConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__AuditConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__AuditLogConfig: Provides the configuration for logging a
// type of permissions.
// Example:
//
//     {
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         }
//       ]
//     }
//
// This enables 'DATA_READ' and 'DATA_WRITE' logging, while
// exempting
// foo@gmail.com from DATA_READ logging.
type GoogleIamV1__AuditLogConfig struct {
	// ExemptedMembers: Specifies the identities that do not cause logging
	// for this type of
	// permission.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// LogType: The log type that this config enables.
	//
	// Possible values:
	//   "LOG_TYPE_UNSPECIFIED" - Default case. Should never be this.
	//   "ADMIN_READ" - Admin reads. Example: CloudIAM getIamPolicy
	//   "DATA_WRITE" - Data writes. Example: CloudSQL Users create
	//   "DATA_READ" - Data reads. Example: CloudSQL Users list
	LogType string `json:"logType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExemptedMembers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__AuditLogConfig) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__AuditLogConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__Binding: Associates `members` with a `role`.
type GoogleIamV1__Binding struct {
	// Condition: Unimplemented. The condition that is associated with this
	// binding.
	// NOTE: an unsatisfied condition will not allow user access via
	// current
	// binding. Different bindings, including their conditions, are
	// examined
	// independently.
	Condition *GoogleType__Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` .
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__Binding) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__Policy: Defines an Identity and Access Management (IAM)
// policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **JSON Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com"
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// **YAML Example**
//
//     bindings:
//     - members:
//       - user:mike@example.com
//       - group:admins@example.com
//       - domain:google.com
//       - serviceAccount:my-other-app@appspot.gserviceaccount.com
//       role: roles/owner
//     - members:
//       - user:sean@example.com
//       role: roles/viewer
//
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type GoogleIamV1__Policy struct {
	// AuditConfigs: Specifies cloud audit logging configuration for this
	// policy.
	AuditConfigs []*GoogleIamV1__AuditConfig `json:"auditConfigs,omitempty"`

	// Bindings: Associates a list of `members` to a `role`.
	// `bindings` with no members will result in an error.
	Bindings []*GoogleIamV1__Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuditConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuditConfigs") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__Policy) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__SetIamPolicyRequest: Request message for `SetIamPolicy`
// method.
type GoogleIamV1__SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *GoogleIamV1__Policy `json:"policy,omitempty"`

	// UpdateMask: OPTIONAL: A FieldMask specifying which fields of the
	// policy to modify. Only
	// the fields in the mask will be modified. If no mask is provided,
	// the
	// following default mask is used:
	// paths: "bindings, etag"
	// This field is only used by Cloud IAM.
	UpdateMask string `json:"updateMask,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsRequest: Request message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsResponse: Response message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleLongrunning__ListOperationsResponse: The response message for
// Operations.ListOperations.
type GoogleLongrunning__ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*GoogleLongrunning__Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunning__ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunning__ListOperationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleLongrunning__Operation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunning__Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If `true`, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *GoogleRpc__Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response googleapi.RawMessage `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Done") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleLongrunning__Operation) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleLongrunning__Operation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleProtobuf__Empty: A generic empty message that you can re-use to
// avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type GoogleProtobuf__Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleRpc__Status: The `Status` type defines a logical error model
// that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type GoogleRpc__Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleRpc__Status) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleRpc__Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleType__Expr: Represents an expression text. Example:
//
//     title: "User account presence"
//     description: "Determines whether the request has a user account"
//     expression: "size(request.user) > 0"
type GoogleType__Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which
	// describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in
	// Common Expression Language syntax.
	//
	// The application context of the containing message determines
	// which
	// well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error
	// reporting, e.g. a file name and a position in the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing
	// its purpose. This can be used e.g. in UIs which allow to enter
	// the
	// expression.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleType__Expr) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleType__Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "ml.projects.getConfig":

type ProjectsGetConfigCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetConfig: Get the service account information associated with your
// project. You need
// this information in order to grant the service account permissions
// for
// the Google Cloud Storage location where you put your model training
// code
// for training the model with Google Cloud Machine Learning.
func (r *ProjectsService) GetConfig(name string) *ProjectsGetConfigCall {
	c := &ProjectsGetConfigCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetConfigCall) Fields(s ...googleapi.Field) *ProjectsGetConfigCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetConfigCall) IfNoneMatch(entityTag string) *ProjectsGetConfigCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetConfigCall) Context(ctx context.Context) *ProjectsGetConfigCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGetConfigCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGetConfigCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:getConfig")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.getConfig" call.
// Exactly one of *GoogleCloudMlV1__GetConfigResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1__GetConfigResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsGetConfigCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__GetConfigResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__GetConfigResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the service account information associated with your project. You need\nthis information in order to grant the service account permissions for\nthe Google Cloud Storage location where you put your model training code\nfor training the model with Google Cloud Machine Learning.",
	//   "flatPath": "v1/projects/{projectsId}:getConfig",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.getConfig",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The project name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:getConfig",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__GetConfigResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.predict":

type ProjectsPredictCall struct {
	s                               *Service
	name                            string
	googlecloudmlv1__predictrequest *GoogleCloudMlV1__PredictRequest
	urlParams_                      gensupport.URLParams
	ctx_                            context.Context
	header_                         http.Header
}

// Predict: Performs prediction on the data in the request.
// Cloud ML Engine implements a custom `predict` verb on top of an HTTP
// POST
// method. <p>For details of the request and response format, see the
// **guide
// to the [predict request
// format](/ml-engine/docs/v1/predict-request)**.
func (r *ProjectsService) Predict(name string, googlecloudmlv1__predictrequest *GoogleCloudMlV1__PredictRequest) *ProjectsPredictCall {
	c := &ProjectsPredictCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__predictrequest = googlecloudmlv1__predictrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPredictCall) Fields(s ...googleapi.Field) *ProjectsPredictCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPredictCall) Context(ctx context.Context) *ProjectsPredictCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsPredictCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsPredictCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body = strings.NewReader(c.googlecloudmlv1__predictrequest.HttpBody.Data)
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:predict")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.predict" call.
// Exactly one of *GoogleApi__HttpBody or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleApi__HttpBody.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPredictCall) Do(opts ...googleapi.CallOption) (*GoogleApi__HttpBody, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleApi__HttpBody{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	var b bytes.Buffer
	if _, err := io.Copy(&b, res.Body); err != nil {
		return nil, err
	}
	if err := res.Body.Close(); err != nil {
		return nil, err
	}
	if err := json.NewDecoder(bytes.NewReader(b.Bytes())).Decode(target); err != nil {
		return nil, err
	}
	ret.Data = b.String()
	return ret, nil
	// {
	//   "description": "Performs prediction on the data in the request.\nCloud ML Engine implements a custom `predict` verb on top of an HTTP POST\nmethod. \u003cp\u003eFor details of the request and response format, see the **guide\nto the [predict request format](/ml-engine/docs/v1/predict-request)**.",
	//   "flatPath": "v1/projects/{projectsId}:predict",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.predict",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of a model or a version.\n\nAuthorization: requires the `predict` permission on the specified resource.",
	//       "location": "path",
	//       "pattern": "^projects/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:predict",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__PredictRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleApi__HttpBody"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.cancel":

type ProjectsJobsCancelCall struct {
	s                                 *Service
	name                              string
	googlecloudmlv1__canceljobrequest *GoogleCloudMlV1__CancelJobRequest
	urlParams_                        gensupport.URLParams
	ctx_                              context.Context
	header_                           http.Header
}

// Cancel: Cancels a running job.
func (r *ProjectsJobsService) Cancel(name string, googlecloudmlv1__canceljobrequest *GoogleCloudMlV1__CancelJobRequest) *ProjectsJobsCancelCall {
	c := &ProjectsJobsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__canceljobrequest = googlecloudmlv1__canceljobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsCancelCall) Fields(s ...googleapi.Field) *ProjectsJobsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsCancelCall) Context(ctx context.Context) *ProjectsJobsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__canceljobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.cancel" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsCancelCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleProtobuf__Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels a running job.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.jobs.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the job to cancel.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__CancelJobRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.create":

type ProjectsJobsCreateCall struct {
	s                    *Service
	parent               string
	googlecloudmlv1__job *GoogleCloudMlV1__Job
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: Creates a training or a batch prediction job.
func (r *ProjectsJobsService) Create(parent string, googlecloudmlv1__job *GoogleCloudMlV1__Job) *ProjectsJobsCreateCall {
	c := &ProjectsJobsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlecloudmlv1__job = googlecloudmlv1__job
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsCreateCall) Fields(s ...googleapi.Field) *ProjectsJobsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsCreateCall) Context(ctx context.Context) *ProjectsJobsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__job)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.create" call.
// Exactly one of *GoogleCloudMlV1__Job or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Job.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a training or a batch prediction job.",
	//   "flatPath": "v1/projects/{projectsId}/jobs",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.jobs.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/jobs",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Job"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.get":

type ProjectsJobsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Describes a job.
func (r *ProjectsJobsService) Get(name string) *ProjectsJobsGetCall {
	c := &ProjectsJobsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsGetCall) Fields(s ...googleapi.Field) *ProjectsJobsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsJobsGetCall) IfNoneMatch(entityTag string) *ProjectsJobsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsGetCall) Context(ctx context.Context) *ProjectsJobsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.get" call.
// Exactly one of *GoogleCloudMlV1__Job or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Job.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Describes a job.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.jobs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the job to get the description of.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.getIamPolicy":

type ProjectsJobsGetIamPolicyCall struct {
	s            *Service
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a resource.
// Returns an empty policy if the resource exists and does not have a
// policy
// set.
func (r *ProjectsJobsService) GetIamPolicy(resource string) *ProjectsJobsGetIamPolicyCall {
	c := &ProjectsJobsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsJobsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsJobsGetIamPolicyCall) IfNoneMatch(entityTag string) *ProjectsJobsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsGetIamPolicyCall) Context(ctx context.Context) *ProjectsJobsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.getIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource.\nReturns an empty policy if the resource exists and does not have a policy\nset.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.jobs.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.list":

type ProjectsJobsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the jobs in the project.
//
// If there are no jobs that match the request parameters, the
// list
// request returns an empty response body: {}.
func (r *ProjectsJobsService) List(parent string) *ProjectsJobsListCall {
	c := &ProjectsJobsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": Specifies the subset of
// jobs to retrieve.
// You can filter on the value of one or more attributes of the job
// object.
// For example, retrieve jobs with a job identifier that starts with
// 'census':
// <p><code>gcloud ml-engine jobs list
// --filter='jobId:census*'</code>
// <p>List all failed jobs with names that start with
// 'rnn':
// <p><code>gcloud ml-engine jobs list --filter='jobId:rnn*
// AND state:FAILED'</code>
// <p>For more examples, see the guide to
// <a href="/ml-engine/docs/tensorflow/monitor-training">monitoring
// jobs</a>.
func (c *ProjectsJobsListCall) Filter(filter string) *ProjectsJobsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of jobs
// to retrieve per "page" of results. If there
// are more remaining results than this number, the response message
// will
// contain a valid value in the `next_page_token` field.
//
// The default value is 20, and the maximum page size is 100.
func (c *ProjectsJobsListCall) PageSize(pageSize int64) *ProjectsJobsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token to
// request the next page of results.
//
// You get the token from the `next_page_token` field of the response
// from
// the previous call.
func (c *ProjectsJobsListCall) PageToken(pageToken string) *ProjectsJobsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsListCall) Fields(s ...googleapi.Field) *ProjectsJobsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsJobsListCall) IfNoneMatch(entityTag string) *ProjectsJobsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsListCall) Context(ctx context.Context) *ProjectsJobsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/jobs")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.list" call.
// Exactly one of *GoogleCloudMlV1__ListJobsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1__ListJobsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsJobsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__ListJobsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__ListJobsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the jobs in the project.\n\nIf there are no jobs that match the request parameters, the list\nrequest returns an empty response body: {}.",
	//   "flatPath": "v1/projects/{projectsId}/jobs",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.jobs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Specifies the subset of jobs to retrieve.\nYou can filter on the value of one or more attributes of the job object.\nFor example, retrieve jobs with a job identifier that starts with 'census':\n\u003cp\u003e\u003ccode\u003egcloud ml-engine jobs list --filter='jobId:census*'\u003c/code\u003e\n\u003cp\u003eList all failed jobs with names that start with 'rnn':\n\u003cp\u003e\u003ccode\u003egcloud ml-engine jobs list --filter='jobId:rnn*\nAND state:FAILED'\u003c/code\u003e\n\u003cp\u003eFor more examples, see the guide to\n\u003ca href=\"/ml-engine/docs/tensorflow/monitor-training\"\u003emonitoring jobs\u003c/a\u003e.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The number of jobs to retrieve per \"page\" of results. If there\nare more remaining results than this number, the response message will\ncontain a valid value in the `next_page_token` field.\n\nThe default value is 20, and the maximum page size is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A page token to request the next page of results.\n\nYou get the token from the `next_page_token` field of the response from\nthe previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the project for which to list jobs.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/jobs",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsJobsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1__ListJobsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "ml.projects.jobs.patch":

type ProjectsJobsPatchCall struct {
	s                    *Service
	name                 string
	googlecloudmlv1__job *GoogleCloudMlV1__Job
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Patch: Updates a specific job resource.
//
// Currently the only supported fields to update are `labels`.
func (r *ProjectsJobsService) Patch(name string, googlecloudmlv1__job *GoogleCloudMlV1__Job) *ProjectsJobsPatchCall {
	c := &ProjectsJobsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__job = googlecloudmlv1__job
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// Specifies the path, relative to `Job`, of the field to update.
// To adopt etag mechanism, include `etag` field in the mask, and
// include the
// `etag` value in your job resource.
//
// For example, to change the labels of a job, the `update_mask`
// parameter
// would be specified as `labels`, `etag`, and the
// `PATCH` request body would specify the new value, as follows:
//     {
//       "labels": {
//          "owner": "Google",
//          "color": "Blue"
//       }
//       "etag": "33a64df551425fcc55e4d42a148795d9f25f89d4"
//     }
// If `etag` matches the one on the server, the labels of the job will
// be
// replaced with the given ones, and the server end `etag` will
// be
// recalculated.
//
// Currently the only supported update masks are `labels` and `etag`.
func (c *ProjectsJobsPatchCall) UpdateMask(updateMask string) *ProjectsJobsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsPatchCall) Fields(s ...googleapi.Field) *ProjectsJobsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsPatchCall) Context(ctx context.Context) *ProjectsJobsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__job)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.patch" call.
// Exactly one of *GoogleCloudMlV1__Job or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Job.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Job, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Job{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a specific job resource.\n\nCurrently the only supported fields to update are `labels`.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}",
	//   "httpMethod": "PATCH",
	//   "id": "ml.projects.jobs.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The job name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Specifies the path, relative to `Job`, of the field to update.\nTo adopt etag mechanism, include `etag` field in the mask, and include the\n`etag` value in your job resource.\n\nFor example, to change the labels of a job, the `update_mask` parameter\nwould be specified as `labels`, `etag`, and the\n`PATCH` request body would specify the new value, as follows:\n    {\n      \"labels\": {\n         \"owner\": \"Google\",\n         \"color\": \"Blue\"\n      }\n      \"etag\": \"33a64df551425fcc55e4d42a148795d9f25f89d4\"\n    }\nIf `etag` matches the one on the server, the labels of the job will be\nreplaced with the given ones, and the server end `etag` will be\nrecalculated.\n\nCurrently the only supported update masks are `labels` and `etag`.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Job"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.setIamPolicy":

type ProjectsJobsSetIamPolicyCall struct {
	s                                *Service
	resource                         string
	googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
func (r *ProjectsJobsService) SetIamPolicy(resource string, googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest) *ProjectsJobsSetIamPolicyCall {
	c := &ProjectsJobsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__setiampolicyrequest = googleiamv1__setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsJobsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsSetIamPolicyCall) Context(ctx context.Context) *ProjectsJobsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.setIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.jobs.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "GoogleIamV1__SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.jobs.testIamPermissions":

type ProjectsJobsTestIamPermissionsCall struct {
	s                                      *Service
	resource                               string
	googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
// If the resource does not exist, this will return an empty set
// of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware
// UIs and command-line tools, not for authorization checking. This
// operation
// may "fail open" without warning.
func (r *ProjectsJobsService) TestIamPermissions(resource string, googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest) *ProjectsJobsTestIamPermissionsCall {
	c := &ProjectsJobsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__testiampermissionsrequest = googleiamv1__testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsJobsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsJobsTestIamPermissionsCall) Context(ctx context.Context) *ProjectsJobsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsJobsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsJobsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.jobs.testIamPermissions" call.
// Exactly one of *GoogleIamV1__TestIamPermissionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleIamV1__TestIamPermissionsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsJobsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource.\nIf the resource does not exist, this will return an empty set of\npermissions, not a NOT_FOUND error.\n\nNote: This operation is designed to be used for building permission-aware\nUIs and command-line tools, not for authorization checking. This operation\nmay \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/jobs/{jobsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.jobs.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/jobs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.locations.get":

type ProjectsLocationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get the complete list of CMLE capabilities in a location, along
// with their
// location-specific properties.
func (r *ProjectsLocationsService) Get(name string) *ProjectsLocationsGetCall {
	c := &ProjectsLocationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsGetCall) Context(ctx context.Context) *ProjectsLocationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.locations.get" call.
// Exactly one of *GoogleCloudMlV1__Location or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Location.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Location, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Location{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get the complete list of CMLE capabilities in a location, along with their\nlocation-specific properties.",
	//   "flatPath": "v1/projects/{projectsId}/locations/{locationsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.locations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the location.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.locations.list":

type ProjectsLocationsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List all locations that provides at least one type of CMLE
// capability.
func (r *ProjectsLocationsService) List(parent string) *ProjectsLocationsListCall {
	c := &ProjectsLocationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// locations to retrieve per "page" of results. If there
// are more remaining results than this number, the response message
// will
// contain a valid value in the `next_page_token` field.
//
// The default value is 20, and the maximum page size is 100.
func (c *ProjectsLocationsListCall) PageSize(pageSize int64) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token to
// request the next page of results.
//
// You get the token from the `next_page_token` field of the response
// from
// the previous call.
func (c *ProjectsLocationsListCall) PageToken(pageToken string) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsListCall) Context(ctx context.Context) *ProjectsLocationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/locations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.locations.list" call.
// Exactly one of *GoogleCloudMlV1__ListLocationsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudMlV1__ListLocationsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__ListLocationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__ListLocationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all locations that provides at least one type of CMLE capability.",
	//   "flatPath": "v1/projects/{projectsId}/locations",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.locations.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The number of locations to retrieve per \"page\" of results. If there\nare more remaining results than this number, the response message will\ncontain a valid value in the `next_page_token` field.\n\nThe default value is 20, and the maximum page size is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A page token to request the next page of results.\n\nYou get the token from the `next_page_token` field of the response from\nthe previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the project for which available locations are to be\nlisted (since some locations might be whitelisted for specific projects).",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/locations",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__ListLocationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1__ListLocationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "ml.projects.models.create":

type ProjectsModelsCreateCall struct {
	s                      *Service
	parent                 string
	googlecloudmlv1__model *GoogleCloudMlV1__Model
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Create: Creates a model which will later contain one or more
// versions.
//
// You must add at least one version before you can request predictions
// from
// the model. Add versions by
// calling
// [projects.models.versions.create](/ml-engine/reference/rest/v1
// /projects.models.versions/create).
func (r *ProjectsModelsService) Create(parent string, googlecloudmlv1__model *GoogleCloudMlV1__Model) *ProjectsModelsCreateCall {
	c := &ProjectsModelsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlecloudmlv1__model = googlecloudmlv1__model
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsCreateCall) Fields(s ...googleapi.Field) *ProjectsModelsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsCreateCall) Context(ctx context.Context) *ProjectsModelsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__model)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/models")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.create" call.
// Exactly one of *GoogleCloudMlV1__Model or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Model.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Model, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Model{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a model which will later contain one or more versions.\n\nYou must add at least one version before you can request predictions from\nthe model. Add versions by calling\n[projects.models.versions.create](/ml-engine/reference/rest/v1/projects.models.versions/create).",
	//   "flatPath": "v1/projects/{projectsId}/models",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The project name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/models",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Model"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Model"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.delete":

type ProjectsModelsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a model.
//
// You can only delete a model if there are no versions in it. You can
// delete
// versions by
// calling
// [projects.models.versions.delete](/ml-engine/reference/rest/v1
// /projects.models.versions/delete).
func (r *ProjectsModelsService) Delete(name string) *ProjectsModelsDeleteCall {
	c := &ProjectsModelsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsDeleteCall) Fields(s ...googleapi.Field) *ProjectsModelsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsDeleteCall) Context(ctx context.Context) *ProjectsModelsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.delete" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a model.\n\nYou can only delete a model if there are no versions in it. You can delete\nversions by calling\n[projects.models.versions.delete](/ml-engine/reference/rest/v1/projects.models.versions/delete).",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.models.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the model.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.get":

type ProjectsModelsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a model, including its name, the
// description (if
// set), and the default version (if at least one version of the model
// has
// been deployed).
func (r *ProjectsModelsService) Get(name string) *ProjectsModelsGetCall {
	c := &ProjectsModelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsGetCall) Fields(s ...googleapi.Field) *ProjectsModelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsGetCall) IfNoneMatch(entityTag string) *ProjectsModelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsGetCall) Context(ctx context.Context) *ProjectsModelsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.get" call.
// Exactly one of *GoogleCloudMlV1__Model or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Model.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Model, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Model{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets information about a model, including its name, the description (if\nset), and the default version (if at least one version of the model has\nbeen deployed).",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the model.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Model"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.getIamPolicy":

type ProjectsModelsGetIamPolicyCall struct {
	s            *Service
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a resource.
// Returns an empty policy if the resource exists and does not have a
// policy
// set.
func (r *ProjectsModelsService) GetIamPolicy(resource string) *ProjectsModelsGetIamPolicyCall {
	c := &ProjectsModelsGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsModelsGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsGetIamPolicyCall) IfNoneMatch(entityTag string) *ProjectsModelsGetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsGetIamPolicyCall) Context(ctx context.Context) *ProjectsModelsGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.getIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource.\nReturns an empty policy if the resource exists and does not have a policy\nset.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.list":

type ProjectsModelsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the models in a project.
//
// Each project can contain multiple models, and each model can have
// multiple
// versions.
//
// If there are no models that match the request parameters, the list
// request
// returns an empty response body: {}.
func (r *ProjectsModelsService) List(parent string) *ProjectsModelsListCall {
	c := &ProjectsModelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": Specifies the subset of
// models to retrieve.
func (c *ProjectsModelsListCall) Filter(filter string) *ProjectsModelsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of models
// to retrieve per "page" of results. If there
// are more remaining results than this number, the response message
// will
// contain a valid value in the `next_page_token` field.
//
// The default value is 20, and the maximum page size is 100.
func (c *ProjectsModelsListCall) PageSize(pageSize int64) *ProjectsModelsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token to
// request the next page of results.
//
// You get the token from the `next_page_token` field of the response
// from
// the previous call.
func (c *ProjectsModelsListCall) PageToken(pageToken string) *ProjectsModelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsListCall) Fields(s ...googleapi.Field) *ProjectsModelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsListCall) IfNoneMatch(entityTag string) *ProjectsModelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsListCall) Context(ctx context.Context) *ProjectsModelsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/models")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.list" call.
// Exactly one of *GoogleCloudMlV1__ListModelsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1__ListModelsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsModelsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__ListModelsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__ListModelsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the models in a project.\n\nEach project can contain multiple models, and each model can have multiple\nversions.\n\nIf there are no models that match the request parameters, the list request\nreturns an empty response body: {}.",
	//   "flatPath": "v1/projects/{projectsId}/models",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Specifies the subset of models to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The number of models to retrieve per \"page\" of results. If there\nare more remaining results than this number, the response message will\ncontain a valid value in the `next_page_token` field.\n\nThe default value is 20, and the maximum page size is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A page token to request the next page of results.\n\nYou get the token from the `next_page_token` field of the response from\nthe previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the project whose models are to be listed.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/models",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__ListModelsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsModelsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1__ListModelsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "ml.projects.models.patch":

type ProjectsModelsPatchCall struct {
	s                      *Service
	name                   string
	googlecloudmlv1__model *GoogleCloudMlV1__Model
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Patch: Updates a specific model resource.
//
// Currently the only supported fields to update are `description`
// and
// `default_version.name`.
func (r *ProjectsModelsService) Patch(name string, googlecloudmlv1__model *GoogleCloudMlV1__Model) *ProjectsModelsPatchCall {
	c := &ProjectsModelsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__model = googlecloudmlv1__model
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// Specifies the path, relative to `Model`, of the field to update.
//
// For example, to change the description of a model to "foo" and set
// its
// default version to "version_1", the `update_mask` parameter would
// be
// specified as `description`, `default_version.name`, and the
// `PATCH`
// request body would specify the new value, as follows:
//     {
//       "description": "foo",
//       "defaultVersion": {
//         "name":"version_1"
//       }
//     }
//
// Currently the supported update masks are `description`
// and
// `default_version.name`.
func (c *ProjectsModelsPatchCall) UpdateMask(updateMask string) *ProjectsModelsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsPatchCall) Fields(s ...googleapi.Field) *ProjectsModelsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsPatchCall) Context(ctx context.Context) *ProjectsModelsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__model)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.patch" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a specific model resource.\n\nCurrently the only supported fields to update are `description` and\n`default_version.name`.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "PATCH",
	//   "id": "ml.projects.models.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The project name.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Specifies the path, relative to `Model`, of the field to update.\n\nFor example, to change the description of a model to \"foo\" and set its\ndefault version to \"version_1\", the `update_mask` parameter would be\nspecified as `description`, `default_version.name`, and the `PATCH`\nrequest body would specify the new value, as follows:\n    {\n      \"description\": \"foo\",\n      \"defaultVersion\": {\n        \"name\":\"version_1\"\n      }\n    }\n\nCurrently the supported update masks are `description` and\n`default_version.name`.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Model"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.setIamPolicy":

type ProjectsModelsSetIamPolicyCall struct {
	s                                *Service
	resource                         string
	googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
func (r *ProjectsModelsService) SetIamPolicy(resource string, googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest) *ProjectsModelsSetIamPolicyCall {
	c := &ProjectsModelsSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__setiampolicyrequest = googleiamv1__setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsModelsSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsSetIamPolicyCall) Context(ctx context.Context) *ProjectsModelsSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.setIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "GoogleIamV1__SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.testIamPermissions":

type ProjectsModelsTestIamPermissionsCall struct {
	s                                      *Service
	resource                               string
	googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
// If the resource does not exist, this will return an empty set
// of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware
// UIs and command-line tools, not for authorization checking. This
// operation
// may "fail open" without warning.
func (r *ProjectsModelsService) TestIamPermissions(resource string, googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest) *ProjectsModelsTestIamPermissionsCall {
	c := &ProjectsModelsTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__testiampermissionsrequest = googleiamv1__testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsModelsTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsTestIamPermissionsCall) Context(ctx context.Context) *ProjectsModelsTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.testIamPermissions" call.
// Exactly one of *GoogleIamV1__TestIamPermissionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleIamV1__TestIamPermissionsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource.\nIf the resource does not exist, this will return an empty set of\npermissions, not a NOT_FOUND error.\n\nNote: This operation is designed to be used for building permission-aware\nUIs and command-line tools, not for authorization checking. This operation\nmay \"fail open\" without warning.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.create":

type ProjectsModelsVersionsCreateCall struct {
	s                        *Service
	parent                   string
	googlecloudmlv1__version *GoogleCloudMlV1__Version
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// Create: Creates a new version of a model from a trained TensorFlow
// model.
//
// If the version created in the cloud by this call is the first
// deployed
// version of the specified model, it will be made the default version
// of the
// model. When you add a version to a model that already has one or
// more
// versions, the default version does not automatically change. If you
// want a
// new version to be the default, you must
// call
// [projects.models.versions.setDefault](/ml-engine/reference/rest/v
// 1/projects.models.versions/setDefault).
func (r *ProjectsModelsVersionsService) Create(parent string, googlecloudmlv1__version *GoogleCloudMlV1__Version) *ProjectsModelsVersionsCreateCall {
	c := &ProjectsModelsVersionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.googlecloudmlv1__version = googlecloudmlv1__version
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsCreateCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsCreateCall) Context(ctx context.Context) *ProjectsModelsVersionsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.create" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new version of a model from a trained TensorFlow model.\n\nIf the version created in the cloud by this call is the first deployed\nversion of the specified model, it will be made the default version of the\nmodel. When you add a version to a model that already has one or more\nversions, the default version does not automatically change. If you want a\nnew version to be the default, you must call\n[projects.models.versions.setDefault](/ml-engine/reference/rest/v1/projects.models.versions/setDefault).",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.versions.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required. The name of the model.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/versions",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Version"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.delete":

type ProjectsModelsVersionsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a model version.
//
// Each model can have multiple versions deployed and in use at any
// given
// time. Use this method to remove a single version.
//
// Note: You cannot delete the version that is set as the default
// version
// of the model unless it is the only remaining version.
func (r *ProjectsModelsVersionsService) Delete(name string) *ProjectsModelsVersionsDeleteCall {
	c := &ProjectsModelsVersionsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsDeleteCall) Context(ctx context.Context) *ProjectsModelsVersionsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.delete" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a model version.\n\nEach model can have multiple versions deployed and in use at any given\ntime. Use this method to remove a single version.\n\nNote: You cannot delete the version that is set as the default version\nof the model unless it is the only remaining version.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.models.versions.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the version. You can get the names of all the\nversions of a model by calling\n[projects.models.versions.list](/ml-engine/reference/rest/v1/projects.models.versions/list).",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.get":

type ProjectsModelsVersionsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets information about a model version.
//
// Models can have multiple versions. You can
// call
// [projects.models.versions.list](/ml-engine/reference/rest/v1/proj
// ects.models.versions/list)
// to get the same information that this method returns for all of
// the
// versions of a model.
func (r *ProjectsModelsVersionsService) Get(name string) *ProjectsModelsVersionsGetCall {
	c := &ProjectsModelsVersionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsGetCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsVersionsGetCall) IfNoneMatch(entityTag string) *ProjectsModelsVersionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsGetCall) Context(ctx context.Context) *ProjectsModelsVersionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.get" call.
// Exactly one of *GoogleCloudMlV1__Version or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Version.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets information about a model version.\n\nModels can have multiple versions. You can call\n[projects.models.versions.list](/ml-engine/reference/rest/v1/projects.models.versions/list)\nto get the same information that this method returns for all of the\nversions of a model.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.versions.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the version.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.list":

type ProjectsModelsVersionsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Gets basic information about all the versions of a model.
//
// If you expect that a model has many versions, or if you need to
// handle
// only a limited number of results at a time, you can request that the
// list
// be retrieved in batches (called pages).
//
// If there are no versions that match the request parameters, the
// list
// request returns an empty response body: {}.
func (r *ProjectsModelsVersionsService) List(parent string) *ProjectsModelsVersionsListCall {
	c := &ProjectsModelsVersionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": Specifies the subset of
// versions to retrieve.
func (c *ProjectsModelsVersionsListCall) Filter(filter string) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The number of
// versions to retrieve per "page" of results. If
// there are more remaining results than this number, the response
// message
// will contain a valid value in the `next_page_token` field.
//
// The default value is 20, and the maximum page size is 100.
func (c *ProjectsModelsVersionsListCall) PageSize(pageSize int64) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A page token to
// request the next page of results.
//
// You get the token from the `next_page_token` field of the response
// from
// the previous call.
func (c *ProjectsModelsVersionsListCall) PageToken(pageToken string) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsListCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsVersionsListCall) IfNoneMatch(entityTag string) *ProjectsModelsVersionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsListCall) Context(ctx context.Context) *ProjectsModelsVersionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+parent}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.list" call.
// Exactly one of *GoogleCloudMlV1__ListVersionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudMlV1__ListVersionsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__ListVersionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__ListVersionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets basic information about all the versions of a model.\n\nIf you expect that a model has many versions, or if you need to handle\nonly a limited number of results at a time, you can request that the list\nbe retrieved in batches (called pages).\n\nIf there are no versions that match the request parameters, the list\nrequest returns an empty response body: {}.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.versions.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Specifies the subset of versions to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The number of versions to retrieve per \"page\" of results. If\nthere are more remaining results than this number, the response message\nwill contain a valid value in the `next_page_token` field.\n\nThe default value is 20, and the maximum page size is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. A page token to request the next page of results.\n\nYou get the token from the `next_page_token` field of the response from\nthe previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required. The name of the model for which to list the version.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+parent}/versions",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__ListVersionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsModelsVersionsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1__ListVersionsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "ml.projects.models.versions.patch":

type ProjectsModelsVersionsPatchCall struct {
	s                        *Service
	name                     string
	googlecloudmlv1__version *GoogleCloudMlV1__Version
	urlParams_               gensupport.URLParams
	ctx_                     context.Context
	header_                  http.Header
}

// Patch: Updates the specified Version resource.
//
// Currently the only update-able fields are `description`
// and
// `autoScaling.minNodes`.
func (r *ProjectsModelsVersionsService) Patch(name string, googlecloudmlv1__version *GoogleCloudMlV1__Version) *ProjectsModelsVersionsPatchCall {
	c := &ProjectsModelsVersionsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__version = googlecloudmlv1__version
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required.
// Specifies the path, relative to `Version`, of the field to
// update. Must be present and non-empty.
//
// For example, to change the description of a version to "foo",
// the
// `update_mask` parameter would be specified as `description`, and
// the
// `PATCH` request body would specify the new value, as follows:
//     {
//       "description": "foo"
//     }
//
// Currently the only supported update mask fields are `description`
// and
// `autoScaling.minNodes`.
func (c *ProjectsModelsVersionsPatchCall) UpdateMask(updateMask string) *ProjectsModelsVersionsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsPatchCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsPatchCall) Context(ctx context.Context) *ProjectsModelsVersionsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.patch" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsPatchCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified Version resource.\n\nCurrently the only update-able fields are `description` and\n`autoScaling.minNodes`.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "httpMethod": "PATCH",
	//   "id": "ml.projects.models.versions.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the model.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. Specifies the path, relative to `Version`, of the field to\nupdate. Must be present and non-empty.\n\nFor example, to change the description of a version to \"foo\", the\n`update_mask` parameter would be specified as `description`, and the\n`PATCH` request body would specify the new value, as follows:\n    {\n      \"description\": \"foo\"\n    }\n\nCurrently the only supported update mask fields are `description` and\n`autoScaling.minNodes`.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__Version"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.setDefault":

type ProjectsModelsVersionsSetDefaultCall struct {
	s                                         *Service
	name                                      string
	googlecloudmlv1__setdefaultversionrequest *GoogleCloudMlV1__SetDefaultVersionRequest
	urlParams_                                gensupport.URLParams
	ctx_                                      context.Context
	header_                                   http.Header
}

// SetDefault: Designates a version to be the default for the
// model.
//
// The default version is used for prediction requests made against the
// model
// that don't specify a version.
//
// The first version to be created for a model is automatically set as
// the
// default. You must make any subsequent changes to the default
// version
// setting manually using this method.
func (r *ProjectsModelsVersionsService) SetDefault(name string, googlecloudmlv1__setdefaultversionrequest *GoogleCloudMlV1__SetDefaultVersionRequest) *ProjectsModelsVersionsSetDefaultCall {
	c := &ProjectsModelsVersionsSetDefaultCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.googlecloudmlv1__setdefaultversionrequest = googlecloudmlv1__setdefaultversionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsSetDefaultCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsSetDefaultCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsSetDefaultCall) Context(ctx context.Context) *ProjectsModelsVersionsSetDefaultCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsModelsVersionsSetDefaultCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsModelsVersionsSetDefaultCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1__setdefaultversionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:setDefault")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.models.versions.setDefault" call.
// Exactly one of *GoogleCloudMlV1__Version or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GoogleCloudMlV1__Version.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsSetDefaultCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1__Version, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleCloudMlV1__Version{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Designates a version to be the default for the model.\n\nThe default version is used for prediction requests made against the model\nthat don't specify a version.\n\nThe first version to be created for a model is automatically set as the\ndefault. You must make any subsequent changes to the default version\nsetting manually using this method.",
	//   "flatPath": "v1/projects/{projectsId}/models/{modelsId}/versions/{versionsId}:setDefault",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.versions.setDefault",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The name of the version to make the default for the model. You\ncan get the names of all the versions of a model by calling\n[projects.models.versions.list](/ml-engine/reference/rest/v1/projects.models.versions/list).",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/models/[^/]+/versions/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:setDefault",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1__SetDefaultVersionRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1__Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.cancel":

type ProjectsOperationsCancelCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
//  The server
// makes a best effort to cancel the operation, but success is
// not
// guaranteed.  If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.  Clients can
// use
// Operations.GetOperation or
// other methods to check whether the cancellation succeeded or whether
// the
// operation completed despite cancellation. On successful
// cancellation,
// the operation is not deleted; instead, it becomes an operation
// with
// an Operation.error value with a google.rpc.Status.code of
// 1,
// corresponding to `Code.CANCELLED`.
func (r *ProjectsOperationsService) Cancel(name string) *ProjectsOperationsCancelCall {
	c := &ProjectsOperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsOperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsCancelCall) Context(ctx context.Context) *ProjectsOperationsCancelCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsCancelCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.operations.cancel" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsCancelCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleProtobuf__Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation. On successful cancellation,\nthe operation is not deleted; instead, it becomes an operation with\nan Operation.error value with a google.rpc.Status.code of 1,\ncorresponding to `Code.CANCELLED`.",
	//   "flatPath": "v1/projects/{projectsId}/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.delete":

type ProjectsOperationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is
// no longer interested in the operation result. It does not cancel
// the
// operation. If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.
func (r *ProjectsOperationsService) Delete(name string) *ProjectsOperationsDeleteCall {
	c := &ProjectsOperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsOperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsDeleteCall) Context(ctx context.Context) *ProjectsOperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.operations.delete" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleProtobuf__Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v1/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.operations.get" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/operations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.list":

type ProjectsOperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding allows API services to override the
// binding
// to use different resource name schemes, such as `users/*/operations`.
// To
// override the binding, API services can add a binding such
// as
// "/v1/{name=users/*}/operations" to their service configuration.
// For backwards compatibility, the default name includes the
// operations
// collection id, however overriding users must ensure the name
// binding
// is the parent resource, without the operations collection id.
func (r *ProjectsOperationsService) List(name string) *ProjectsOperationsListCall {
	c := &ProjectsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsOperationsListCall) Filter(filter string) *ProjectsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsOperationsListCall) PageSize(pageSize int64) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsOperationsListCall) PageToken(pageToken string) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsListCall) Context(ctx context.Context) *ProjectsOperationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsOperationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/operations")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "ml.projects.operations.list" call.
// Exactly one of *GoogleLongrunning__ListOperationsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleLongrunning__ListOperationsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsOperationsListCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleLongrunning__ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`. To\noverride the binding, API services can add a binding such as\n`\"/v1/{name=users/*}/operations\"` to their service configuration.\nFor backwards compatibility, the default name includes the operations\ncollection id, however overriding users must ensure the name binding\nis the parent resource, without the operations collection id.",
	//   "flatPath": "v1/projects/{projectsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation's parent resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/operations",
	//   "response": {
	//     "$ref": "GoogleLongrunning__ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsOperationsListCall) Pages(ctx context.Context, f func(*GoogleLongrunning__ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
