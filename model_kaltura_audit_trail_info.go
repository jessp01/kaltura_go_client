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

// KalturaAuditTrailInfo `abstract`
type KalturaAuditTrailInfo struct {
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaAuditTrailInfo instantiates a new KalturaAuditTrailInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAuditTrailInfo() *KalturaAuditTrailInfo {
	this := KalturaAuditTrailInfo{}
	return &this
}

// NewKalturaAuditTrailInfoWithDefaults instantiates a new KalturaAuditTrailInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAuditTrailInfoWithDefaults() *KalturaAuditTrailInfo {
	this := KalturaAuditTrailInfo{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaAuditTrailInfo) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAuditTrailInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaAuditTrailInfo) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaAuditTrailInfo) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaAuditTrailInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAuditTrailInfo struct {
	value *KalturaAuditTrailInfo
	isSet bool
}

func (v NullableKalturaAuditTrailInfo) Get() *KalturaAuditTrailInfo {
	return v.value
}

func (v *NullableKalturaAuditTrailInfo) Set(val *KalturaAuditTrailInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAuditTrailInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAuditTrailInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAuditTrailInfo(val *KalturaAuditTrailInfo) *NullableKalturaAuditTrailInfo {
	return &NullableKalturaAuditTrailInfo{value: val, isSet: true}
}

func (v NullableKalturaAuditTrailInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAuditTrailInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


