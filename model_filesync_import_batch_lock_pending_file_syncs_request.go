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

// FilesyncImportBatchLockPendingFileSyncsRequest struct for FilesyncImportBatchLockPendingFileSyncsRequest
type FilesyncImportBatchLockPendingFileSyncsRequest struct {
	Filter KalturaFileSyncFilter `json:"filter"`
	MaxCount int32 `json:"maxCount"`
	MaxSize *int32 `json:"maxSize,omitempty"`
	SourceDc int32 `json:"sourceDc"`
	WorkerId int32 `json:"workerId"`
}

// NewFilesyncImportBatchLockPendingFileSyncsRequest instantiates a new FilesyncImportBatchLockPendingFileSyncsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesyncImportBatchLockPendingFileSyncsRequest(filter KalturaFileSyncFilter, maxCount int32, sourceDc int32, workerId int32) *FilesyncImportBatchLockPendingFileSyncsRequest {
	this := FilesyncImportBatchLockPendingFileSyncsRequest{}
	this.Filter = filter
	this.MaxCount = maxCount
	this.SourceDc = sourceDc
	this.WorkerId = workerId
	return &this
}

// NewFilesyncImportBatchLockPendingFileSyncsRequestWithDefaults instantiates a new FilesyncImportBatchLockPendingFileSyncsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesyncImportBatchLockPendingFileSyncsRequestWithDefaults() *FilesyncImportBatchLockPendingFileSyncsRequest {
	this := FilesyncImportBatchLockPendingFileSyncsRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetFilter() KalturaFileSyncFilter {
	if o == nil {
		var ret KalturaFileSyncFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetFilterOk() (*KalturaFileSyncFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) SetFilter(v KalturaFileSyncFilter) {
	o.Filter = v
}

// GetMaxCount returns the MaxCount field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetMaxCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value
// and a boolean to check if the value has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetMaxCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCount, true
}

// SetMaxCount sets field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) SetMaxCount(v int32) {
	o.MaxCount = v
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetMaxSize() int32 {
	if o == nil || o.MaxSize == nil {
		var ret int32
		return ret
	}
	return *o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetMaxSizeOk() (*int32, bool) {
	if o == nil || o.MaxSize == nil {
		return nil, false
	}
	return o.MaxSize, true
}

// HasMaxSize returns a boolean if a field has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) HasMaxSize() bool {
	if o != nil && o.MaxSize != nil {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given int32 and assigns it to the MaxSize field.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) SetMaxSize(v int32) {
	o.MaxSize = &v
}

// GetSourceDc returns the SourceDc field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetSourceDc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SourceDc
}

// GetSourceDcOk returns a tuple with the SourceDc field value
// and a boolean to check if the value has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetSourceDcOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDc, true
}

// SetSourceDc sets field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) SetSourceDc(v int32) {
	o.SourceDc = v
}

// GetWorkerId returns the WorkerId field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetWorkerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) GetWorkerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *FilesyncImportBatchLockPendingFileSyncsRequest) SetWorkerId(v int32) {
	o.WorkerId = v
}

func (o FilesyncImportBatchLockPendingFileSyncsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["maxCount"] = o.MaxCount
	}
	if o.MaxSize != nil {
		toSerialize["maxSize"] = o.MaxSize
	}
	if true {
		toSerialize["sourceDc"] = o.SourceDc
	}
	if true {
		toSerialize["workerId"] = o.WorkerId
	}
	return json.Marshal(toSerialize)
}

type NullableFilesyncImportBatchLockPendingFileSyncsRequest struct {
	value *FilesyncImportBatchLockPendingFileSyncsRequest
	isSet bool
}

func (v NullableFilesyncImportBatchLockPendingFileSyncsRequest) Get() *FilesyncImportBatchLockPendingFileSyncsRequest {
	return v.value
}

func (v *NullableFilesyncImportBatchLockPendingFileSyncsRequest) Set(val *FilesyncImportBatchLockPendingFileSyncsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesyncImportBatchLockPendingFileSyncsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesyncImportBatchLockPendingFileSyncsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesyncImportBatchLockPendingFileSyncsRequest(val *FilesyncImportBatchLockPendingFileSyncsRequest) *NullableFilesyncImportBatchLockPendingFileSyncsRequest {
	return &NullableFilesyncImportBatchLockPendingFileSyncsRequest{value: val, isSet: true}
}

func (v NullableFilesyncImportBatchLockPendingFileSyncsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesyncImportBatchLockPendingFileSyncsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

