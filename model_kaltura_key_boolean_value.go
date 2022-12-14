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

// KalturaKeyBooleanValue A key (boolean) value pair representation to return an array of key-(boolean)value pairs (associative array)
type KalturaKeyBooleanValue struct {
	Key *string `json:"key,omitempty"`
	Value *bool `json:"value,omitempty"`
}

// NewKalturaKeyBooleanValue instantiates a new KalturaKeyBooleanValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaKeyBooleanValue() *KalturaKeyBooleanValue {
	this := KalturaKeyBooleanValue{}
	return &this
}

// NewKalturaKeyBooleanValueWithDefaults instantiates a new KalturaKeyBooleanValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaKeyBooleanValueWithDefaults() *KalturaKeyBooleanValue {
	this := KalturaKeyBooleanValue{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *KalturaKeyBooleanValue) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaKeyBooleanValue) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *KalturaKeyBooleanValue) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *KalturaKeyBooleanValue) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KalturaKeyBooleanValue) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaKeyBooleanValue) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KalturaKeyBooleanValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *KalturaKeyBooleanValue) SetValue(v bool) {
	o.Value = &v
}

func (o KalturaKeyBooleanValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaKeyBooleanValue struct {
	value *KalturaKeyBooleanValue
	isSet bool
}

func (v NullableKalturaKeyBooleanValue) Get() *KalturaKeyBooleanValue {
	return v.value
}

func (v *NullableKalturaKeyBooleanValue) Set(val *KalturaKeyBooleanValue) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaKeyBooleanValue) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaKeyBooleanValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaKeyBooleanValue(val *KalturaKeyBooleanValue) *NullableKalturaKeyBooleanValue {
	return &NullableKalturaKeyBooleanValue{value: val, isSet: true}
}

func (v NullableKalturaKeyBooleanValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaKeyBooleanValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


