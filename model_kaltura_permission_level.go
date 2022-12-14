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

// KalturaPermissionLevel struct for KalturaPermissionLevel
type KalturaPermissionLevel struct {
	// Enum Type: `KalturaUserEntryPermissionLevel`  Permission Level
	PermissionLevel *int32 `json:"permissionLevel,omitempty"`
}

// NewKalturaPermissionLevel instantiates a new KalturaPermissionLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPermissionLevel() *KalturaPermissionLevel {
	this := KalturaPermissionLevel{}
	return &this
}

// NewKalturaPermissionLevelWithDefaults instantiates a new KalturaPermissionLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPermissionLevelWithDefaults() *KalturaPermissionLevel {
	this := KalturaPermissionLevel{}
	return &this
}

// GetPermissionLevel returns the PermissionLevel field value if set, zero value otherwise.
func (o *KalturaPermissionLevel) GetPermissionLevel() int32 {
	if o == nil || o.PermissionLevel == nil {
		var ret int32
		return ret
	}
	return *o.PermissionLevel
}

// GetPermissionLevelOk returns a tuple with the PermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPermissionLevel) GetPermissionLevelOk() (*int32, bool) {
	if o == nil || o.PermissionLevel == nil {
		return nil, false
	}
	return o.PermissionLevel, true
}

// HasPermissionLevel returns a boolean if a field has been set.
func (o *KalturaPermissionLevel) HasPermissionLevel() bool {
	if o != nil && o.PermissionLevel != nil {
		return true
	}

	return false
}

// SetPermissionLevel gets a reference to the given int32 and assigns it to the PermissionLevel field.
func (o *KalturaPermissionLevel) SetPermissionLevel(v int32) {
	o.PermissionLevel = &v
}

func (o KalturaPermissionLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PermissionLevel != nil {
		toSerialize["permissionLevel"] = o.PermissionLevel
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPermissionLevel struct {
	value *KalturaPermissionLevel
	isSet bool
}

func (v NullableKalturaPermissionLevel) Get() *KalturaPermissionLevel {
	return v.value
}

func (v *NullableKalturaPermissionLevel) Set(val *KalturaPermissionLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPermissionLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPermissionLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPermissionLevel(val *KalturaPermissionLevel) *NullableKalturaPermissionLevel {
	return &NullableKalturaPermissionLevel{value: val, isSet: true}
}

func (v NullableKalturaPermissionLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPermissionLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


