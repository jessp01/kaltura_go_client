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

// KalturaObjectIdentifier `abstract`  Configuration for extended item in the Kaltura MRSS feeds
type KalturaObjectIdentifier struct {
	// Comma separated string of enum values denoting which features of the item need to be included in the MRSS
	ExtendedFeatures *string `json:"extendedFeatures,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaObjectIdentifier instantiates a new KalturaObjectIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaObjectIdentifier() *KalturaObjectIdentifier {
	this := KalturaObjectIdentifier{}
	return &this
}

// NewKalturaObjectIdentifierWithDefaults instantiates a new KalturaObjectIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaObjectIdentifierWithDefaults() *KalturaObjectIdentifier {
	this := KalturaObjectIdentifier{}
	return &this
}

// GetExtendedFeatures returns the ExtendedFeatures field value if set, zero value otherwise.
func (o *KalturaObjectIdentifier) GetExtendedFeatures() string {
	if o == nil || o.ExtendedFeatures == nil {
		var ret string
		return ret
	}
	return *o.ExtendedFeatures
}

// GetExtendedFeaturesOk returns a tuple with the ExtendedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaObjectIdentifier) GetExtendedFeaturesOk() (*string, bool) {
	if o == nil || o.ExtendedFeatures == nil {
		return nil, false
	}
	return o.ExtendedFeatures, true
}

// HasExtendedFeatures returns a boolean if a field has been set.
func (o *KalturaObjectIdentifier) HasExtendedFeatures() bool {
	if o != nil && o.ExtendedFeatures != nil {
		return true
	}

	return false
}

// SetExtendedFeatures gets a reference to the given string and assigns it to the ExtendedFeatures field.
func (o *KalturaObjectIdentifier) SetExtendedFeatures(v string) {
	o.ExtendedFeatures = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaObjectIdentifier) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaObjectIdentifier) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaObjectIdentifier) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaObjectIdentifier) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaObjectIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtendedFeatures != nil {
		toSerialize["extendedFeatures"] = o.ExtendedFeatures
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaObjectIdentifier struct {
	value *KalturaObjectIdentifier
	isSet bool
}

func (v NullableKalturaObjectIdentifier) Get() *KalturaObjectIdentifier {
	return v.value
}

func (v *NullableKalturaObjectIdentifier) Set(val *KalturaObjectIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaObjectIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaObjectIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaObjectIdentifier(val *KalturaObjectIdentifier) *NullableKalturaObjectIdentifier {
	return &NullableKalturaObjectIdentifier{value: val, isSet: true}
}

func (v NullableKalturaObjectIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaObjectIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

