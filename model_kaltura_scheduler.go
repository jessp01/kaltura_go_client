/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KalturaScheduler struct for KalturaScheduler
type KalturaScheduler struct {
	Configs []KalturaSchedulerConfig `json:"configs,omitempty"`
	// The id as configured in the batch config
	ConfiguredId *int32 `json:"configuredId,omitempty"`
	// `readOnly`  creation time
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// The host name
	Host *string `json:"host,omitempty"`
	// `readOnly`  The id of the Scheduler
	Id *int32 `json:"id,omitempty"`
	// `readOnly`  last status time
	LastStatus *int32 `json:"lastStatus,omitempty"`
	// `readOnly`  last status formated
	LastStatusStr *string `json:"lastStatusStr,omitempty"`
	// The scheduler name
	Name *string `json:"name,omitempty"`
	Statuses []KalturaSchedulerStatus `json:"statuses,omitempty"`
	Workers []KalturaSchedulerWorker `json:"workers,omitempty"`
}

// NewKalturaScheduler instantiates a new KalturaScheduler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduler() *KalturaScheduler {
	this := KalturaScheduler{}
	return &this
}

// NewKalturaSchedulerWithDefaults instantiates a new KalturaScheduler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSchedulerWithDefaults() *KalturaScheduler {
	this := KalturaScheduler{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *KalturaScheduler) GetConfigs() []KalturaSchedulerConfig {
	if o == nil || o.Configs == nil {
		var ret []KalturaSchedulerConfig
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetConfigsOk() ([]KalturaSchedulerConfig, bool) {
	if o == nil || o.Configs == nil {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *KalturaScheduler) HasConfigs() bool {
	if o != nil && o.Configs != nil {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []KalturaSchedulerConfig and assigns it to the Configs field.
func (o *KalturaScheduler) SetConfigs(v []KalturaSchedulerConfig) {
	o.Configs = v
}

// GetConfiguredId returns the ConfiguredId field value if set, zero value otherwise.
func (o *KalturaScheduler) GetConfiguredId() int32 {
	if o == nil || o.ConfiguredId == nil {
		var ret int32
		return ret
	}
	return *o.ConfiguredId
}

// GetConfiguredIdOk returns a tuple with the ConfiguredId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetConfiguredIdOk() (*int32, bool) {
	if o == nil || o.ConfiguredId == nil {
		return nil, false
	}
	return o.ConfiguredId, true
}

// HasConfiguredId returns a boolean if a field has been set.
func (o *KalturaScheduler) HasConfiguredId() bool {
	if o != nil && o.ConfiguredId != nil {
		return true
	}

	return false
}

// SetConfiguredId gets a reference to the given int32 and assigns it to the ConfiguredId field.
func (o *KalturaScheduler) SetConfiguredId(v int32) {
	o.ConfiguredId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaScheduler) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaScheduler) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaScheduler) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *KalturaScheduler) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *KalturaScheduler) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *KalturaScheduler) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaScheduler) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaScheduler) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaScheduler) SetId(v int32) {
	o.Id = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *KalturaScheduler) GetLastStatus() int32 {
	if o == nil || o.LastStatus == nil {
		var ret int32
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetLastStatusOk() (*int32, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *KalturaScheduler) HasLastStatus() bool {
	if o != nil && o.LastStatus != nil {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given int32 and assigns it to the LastStatus field.
func (o *KalturaScheduler) SetLastStatus(v int32) {
	o.LastStatus = &v
}

// GetLastStatusStr returns the LastStatusStr field value if set, zero value otherwise.
func (o *KalturaScheduler) GetLastStatusStr() string {
	if o == nil || o.LastStatusStr == nil {
		var ret string
		return ret
	}
	return *o.LastStatusStr
}

// GetLastStatusStrOk returns a tuple with the LastStatusStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetLastStatusStrOk() (*string, bool) {
	if o == nil || o.LastStatusStr == nil {
		return nil, false
	}
	return o.LastStatusStr, true
}

// HasLastStatusStr returns a boolean if a field has been set.
func (o *KalturaScheduler) HasLastStatusStr() bool {
	if o != nil && o.LastStatusStr != nil {
		return true
	}

	return false
}

// SetLastStatusStr gets a reference to the given string and assigns it to the LastStatusStr field.
func (o *KalturaScheduler) SetLastStatusStr(v string) {
	o.LastStatusStr = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaScheduler) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaScheduler) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaScheduler) SetName(v string) {
	o.Name = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *KalturaScheduler) GetStatuses() []KalturaSchedulerStatus {
	if o == nil || o.Statuses == nil {
		var ret []KalturaSchedulerStatus
		return ret
	}
	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetStatusesOk() ([]KalturaSchedulerStatus, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *KalturaScheduler) HasStatuses() bool {
	if o != nil && o.Statuses != nil {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given []KalturaSchedulerStatus and assigns it to the Statuses field.
func (o *KalturaScheduler) SetStatuses(v []KalturaSchedulerStatus) {
	o.Statuses = v
}

// GetWorkers returns the Workers field value if set, zero value otherwise.
func (o *KalturaScheduler) GetWorkers() []KalturaSchedulerWorker {
	if o == nil || o.Workers == nil {
		var ret []KalturaSchedulerWorker
		return ret
	}
	return o.Workers
}

// GetWorkersOk returns a tuple with the Workers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduler) GetWorkersOk() ([]KalturaSchedulerWorker, bool) {
	if o == nil || o.Workers == nil {
		return nil, false
	}
	return o.Workers, true
}

// HasWorkers returns a boolean if a field has been set.
func (o *KalturaScheduler) HasWorkers() bool {
	if o != nil && o.Workers != nil {
		return true
	}

	return false
}

// SetWorkers gets a reference to the given []KalturaSchedulerWorker and assigns it to the Workers field.
func (o *KalturaScheduler) SetWorkers(v []KalturaSchedulerWorker) {
	o.Workers = v
}

func (o KalturaScheduler) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configs != nil {
		toSerialize["configs"] = o.Configs
	}
	if o.ConfiguredId != nil {
		toSerialize["configuredId"] = o.ConfiguredId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastStatus != nil {
		toSerialize["lastStatus"] = o.LastStatus
	}
	if o.LastStatusStr != nil {
		toSerialize["lastStatusStr"] = o.LastStatusStr
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.Workers != nil {
		toSerialize["workers"] = o.Workers
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduler struct {
	value *KalturaScheduler
	isSet bool
}

func (v NullableKalturaScheduler) Get() *KalturaScheduler {
	return v.value
}

func (v *NullableKalturaScheduler) Set(val *KalturaScheduler) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduler) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduler(val *KalturaScheduler) *NullableKalturaScheduler {
	return &NullableKalturaScheduler{value: val, isSet: true}
}

func (v NullableKalturaScheduler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


