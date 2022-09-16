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

// FileSyncDeleteLocalFileSyncsRequest struct for FileSyncDeleteLocalFileSyncsRequest
type FileSyncDeleteLocalFileSyncsRequest struct {
	Filter KalturaFileSyncFilter `json:"filter"`
	LockExpiryTimeout int32 `json:"lockExpiryTimeout"`
	RelativeTimeDeletionLimit int32 `json:"relativeTimeDeletionLimit"`
	RelativeTimeRange int32 `json:"relativeTimeRange"`
	WorkerId int32 `json:"workerId"`
}

// NewFileSyncDeleteLocalFileSyncsRequest instantiates a new FileSyncDeleteLocalFileSyncsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSyncDeleteLocalFileSyncsRequest(filter KalturaFileSyncFilter, lockExpiryTimeout int32, relativeTimeDeletionLimit int32, relativeTimeRange int32, workerId int32) *FileSyncDeleteLocalFileSyncsRequest {
	this := FileSyncDeleteLocalFileSyncsRequest{}
	this.Filter = filter
	this.LockExpiryTimeout = lockExpiryTimeout
	this.RelativeTimeDeletionLimit = relativeTimeDeletionLimit
	this.RelativeTimeRange = relativeTimeRange
	this.WorkerId = workerId
	return &this
}

// NewFileSyncDeleteLocalFileSyncsRequestWithDefaults instantiates a new FileSyncDeleteLocalFileSyncsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSyncDeleteLocalFileSyncsRequestWithDefaults() *FileSyncDeleteLocalFileSyncsRequest {
	this := FileSyncDeleteLocalFileSyncsRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *FileSyncDeleteLocalFileSyncsRequest) GetFilter() KalturaFileSyncFilter {
	if o == nil {
		var ret KalturaFileSyncFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *FileSyncDeleteLocalFileSyncsRequest) GetFilterOk() (*KalturaFileSyncFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *FileSyncDeleteLocalFileSyncsRequest) SetFilter(v KalturaFileSyncFilter) {
	o.Filter = v
}

// GetLockExpiryTimeout returns the LockExpiryTimeout field value
func (o *FileSyncDeleteLocalFileSyncsRequest) GetLockExpiryTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockExpiryTimeout
}

// GetLockExpiryTimeoutOk returns a tuple with the LockExpiryTimeout field value
// and a boolean to check if the value has been set.
func (o *FileSyncDeleteLocalFileSyncsRequest) GetLockExpiryTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockExpiryTimeout, true
}

// SetLockExpiryTimeout sets field value
func (o *FileSyncDeleteLocalFileSyncsRequest) SetLockExpiryTimeout(v int32) {
	o.LockExpiryTimeout = v
}

// GetRelativeTimeDeletionLimit returns the RelativeTimeDeletionLimit field value
func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeDeletionLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelativeTimeDeletionLimit
}

// GetRelativeTimeDeletionLimitOk returns a tuple with the RelativeTimeDeletionLimit field value
// and a boolean to check if the value has been set.
func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeDeletionLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativeTimeDeletionLimit, true
}

// SetRelativeTimeDeletionLimit sets field value
func (o *FileSyncDeleteLocalFileSyncsRequest) SetRelativeTimeDeletionLimit(v int32) {
	o.RelativeTimeDeletionLimit = v
}

// GetRelativeTimeRange returns the RelativeTimeRange field value
func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeRange() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelativeTimeRange
}

// GetRelativeTimeRangeOk returns a tuple with the RelativeTimeRange field value
// and a boolean to check if the value has been set.
func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeRangeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativeTimeRange, true
}

// SetRelativeTimeRange sets field value
func (o *FileSyncDeleteLocalFileSyncsRequest) SetRelativeTimeRange(v int32) {
	o.RelativeTimeRange = v
}

// GetWorkerId returns the WorkerId field value
func (o *FileSyncDeleteLocalFileSyncsRequest) GetWorkerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *FileSyncDeleteLocalFileSyncsRequest) GetWorkerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *FileSyncDeleteLocalFileSyncsRequest) SetWorkerId(v int32) {
	o.WorkerId = v
}

func (o FileSyncDeleteLocalFileSyncsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["lockExpiryTimeout"] = o.LockExpiryTimeout
	}
	if true {
		toSerialize["relativeTimeDeletionLimit"] = o.RelativeTimeDeletionLimit
	}
	if true {
		toSerialize["relativeTimeRange"] = o.RelativeTimeRange
	}
	if true {
		toSerialize["workerId"] = o.WorkerId
	}
	return json.Marshal(toSerialize)
}

type NullableFileSyncDeleteLocalFileSyncsRequest struct {
	value *FileSyncDeleteLocalFileSyncsRequest
	isSet bool
}

func (v NullableFileSyncDeleteLocalFileSyncsRequest) Get() *FileSyncDeleteLocalFileSyncsRequest {
	return v.value
}

func (v *NullableFileSyncDeleteLocalFileSyncsRequest) Set(val *FileSyncDeleteLocalFileSyncsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSyncDeleteLocalFileSyncsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSyncDeleteLocalFileSyncsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSyncDeleteLocalFileSyncsRequest(val *FileSyncDeleteLocalFileSyncsRequest) *NullableFileSyncDeleteLocalFileSyncsRequest {
	return &NullableFileSyncDeleteLocalFileSyncsRequest{value: val, isSet: true}
}

func (v NullableFileSyncDeleteLocalFileSyncsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSyncDeleteLocalFileSyncsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


