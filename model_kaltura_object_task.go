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

// KalturaObjectTask `abstract`
type KalturaObjectTask struct {
	ObjectType *string `json:"objectType,omitempty"`
	StopProcessingOnError *bool `json:"stopProcessingOnError,omitempty"`
	// `readOnly`  Enum Type: `KalturaObjectTaskType`
	Type *string `json:"type,omitempty"`
}

// NewKalturaObjectTask instantiates a new KalturaObjectTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaObjectTask() *KalturaObjectTask {
	this := KalturaObjectTask{}
	return &this
}

// NewKalturaObjectTaskWithDefaults instantiates a new KalturaObjectTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaObjectTaskWithDefaults() *KalturaObjectTask {
	this := KalturaObjectTask{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaObjectTask) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaObjectTask) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaObjectTask) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaObjectTask) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetStopProcessingOnError returns the StopProcessingOnError field value if set, zero value otherwise.
func (o *KalturaObjectTask) GetStopProcessingOnError() bool {
	if o == nil || o.StopProcessingOnError == nil {
		var ret bool
		return ret
	}
	return *o.StopProcessingOnError
}

// GetStopProcessingOnErrorOk returns a tuple with the StopProcessingOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaObjectTask) GetStopProcessingOnErrorOk() (*bool, bool) {
	if o == nil || o.StopProcessingOnError == nil {
		return nil, false
	}
	return o.StopProcessingOnError, true
}

// HasStopProcessingOnError returns a boolean if a field has been set.
func (o *KalturaObjectTask) HasStopProcessingOnError() bool {
	if o != nil && o.StopProcessingOnError != nil {
		return true
	}

	return false
}

// SetStopProcessingOnError gets a reference to the given bool and assigns it to the StopProcessingOnError field.
func (o *KalturaObjectTask) SetStopProcessingOnError(v bool) {
	o.StopProcessingOnError = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaObjectTask) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaObjectTask) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaObjectTask) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaObjectTask) SetType(v string) {
	o.Type = &v
}

func (o KalturaObjectTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.StopProcessingOnError != nil {
		toSerialize["stopProcessingOnError"] = o.StopProcessingOnError
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaObjectTask struct {
	value *KalturaObjectTask
	isSet bool
}

func (v NullableKalturaObjectTask) Get() *KalturaObjectTask {
	return v.value
}

func (v *NullableKalturaObjectTask) Set(val *KalturaObjectTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaObjectTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaObjectTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaObjectTask(val *KalturaObjectTask) *NullableKalturaObjectTask {
	return &NullableKalturaObjectTask{value: val, isSet: true}
}

func (v NullableKalturaObjectTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaObjectTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

