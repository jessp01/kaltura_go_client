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

// KalturaJobData struct for KalturaJobData
type KalturaJobData struct {
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaJobData instantiates a new KalturaJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaJobData() *KalturaJobData {
	this := KalturaJobData{}
	return &this
}

// NewKalturaJobDataWithDefaults instantiates a new KalturaJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaJobDataWithDefaults() *KalturaJobData {
	this := KalturaJobData{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaJobData) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaJobData) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaJobData) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaJobData) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaJobData struct {
	value *KalturaJobData
	isSet bool
}

func (v NullableKalturaJobData) Get() *KalturaJobData {
	return v.value
}

func (v *NullableKalturaJobData) Set(val *KalturaJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaJobData(val *KalturaJobData) *NullableKalturaJobData {
	return &NullableKalturaJobData{value: val, isSet: true}
}

func (v NullableKalturaJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


