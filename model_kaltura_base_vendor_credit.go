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

// KalturaBaseVendorCredit `abstract`
type KalturaBaseVendorCredit struct {
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaBaseVendorCredit instantiates a new KalturaBaseVendorCredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBaseVendorCredit() *KalturaBaseVendorCredit {
	this := KalturaBaseVendorCredit{}
	return &this
}

// NewKalturaBaseVendorCreditWithDefaults instantiates a new KalturaBaseVendorCredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBaseVendorCreditWithDefaults() *KalturaBaseVendorCredit {
	this := KalturaBaseVendorCredit{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaBaseVendorCredit) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBaseVendorCredit) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaBaseVendorCredit) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaBaseVendorCredit) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaBaseVendorCredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBaseVendorCredit struct {
	value *KalturaBaseVendorCredit
	isSet bool
}

func (v NullableKalturaBaseVendorCredit) Get() *KalturaBaseVendorCredit {
	return v.value
}

func (v *NullableKalturaBaseVendorCredit) Set(val *KalturaBaseVendorCredit) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBaseVendorCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBaseVendorCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBaseVendorCredit(val *KalturaBaseVendorCredit) *NullableKalturaBaseVendorCredit {
	return &NullableKalturaBaseVendorCredit{value: val, isSet: true}
}

func (v NullableKalturaBaseVendorCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBaseVendorCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


