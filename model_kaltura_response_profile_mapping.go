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

// KalturaResponseProfileMapping struct for KalturaResponseProfileMapping
type KalturaResponseProfileMapping struct {
	AllowNull *bool `json:"allowNull,omitempty"`
	FilterProperty *string `json:"filterProperty,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	ParentProperty *string `json:"parentProperty,omitempty"`
}

// NewKalturaResponseProfileMapping instantiates a new KalturaResponseProfileMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaResponseProfileMapping() *KalturaResponseProfileMapping {
	this := KalturaResponseProfileMapping{}
	return &this
}

// NewKalturaResponseProfileMappingWithDefaults instantiates a new KalturaResponseProfileMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaResponseProfileMappingWithDefaults() *KalturaResponseProfileMapping {
	this := KalturaResponseProfileMapping{}
	return &this
}

// GetAllowNull returns the AllowNull field value if set, zero value otherwise.
func (o *KalturaResponseProfileMapping) GetAllowNull() bool {
	if o == nil || o.AllowNull == nil {
		var ret bool
		return ret
	}
	return *o.AllowNull
}

// GetAllowNullOk returns a tuple with the AllowNull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileMapping) GetAllowNullOk() (*bool, bool) {
	if o == nil || o.AllowNull == nil {
		return nil, false
	}
	return o.AllowNull, true
}

// HasAllowNull returns a boolean if a field has been set.
func (o *KalturaResponseProfileMapping) HasAllowNull() bool {
	if o != nil && o.AllowNull != nil {
		return true
	}

	return false
}

// SetAllowNull gets a reference to the given bool and assigns it to the AllowNull field.
func (o *KalturaResponseProfileMapping) SetAllowNull(v bool) {
	o.AllowNull = &v
}

// GetFilterProperty returns the FilterProperty field value if set, zero value otherwise.
func (o *KalturaResponseProfileMapping) GetFilterProperty() string {
	if o == nil || o.FilterProperty == nil {
		var ret string
		return ret
	}
	return *o.FilterProperty
}

// GetFilterPropertyOk returns a tuple with the FilterProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileMapping) GetFilterPropertyOk() (*string, bool) {
	if o == nil || o.FilterProperty == nil {
		return nil, false
	}
	return o.FilterProperty, true
}

// HasFilterProperty returns a boolean if a field has been set.
func (o *KalturaResponseProfileMapping) HasFilterProperty() bool {
	if o != nil && o.FilterProperty != nil {
		return true
	}

	return false
}

// SetFilterProperty gets a reference to the given string and assigns it to the FilterProperty field.
func (o *KalturaResponseProfileMapping) SetFilterProperty(v string) {
	o.FilterProperty = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaResponseProfileMapping) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileMapping) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaResponseProfileMapping) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaResponseProfileMapping) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetParentProperty returns the ParentProperty field value if set, zero value otherwise.
func (o *KalturaResponseProfileMapping) GetParentProperty() string {
	if o == nil || o.ParentProperty == nil {
		var ret string
		return ret
	}
	return *o.ParentProperty
}

// GetParentPropertyOk returns a tuple with the ParentProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaResponseProfileMapping) GetParentPropertyOk() (*string, bool) {
	if o == nil || o.ParentProperty == nil {
		return nil, false
	}
	return o.ParentProperty, true
}

// HasParentProperty returns a boolean if a field has been set.
func (o *KalturaResponseProfileMapping) HasParentProperty() bool {
	if o != nil && o.ParentProperty != nil {
		return true
	}

	return false
}

// SetParentProperty gets a reference to the given string and assigns it to the ParentProperty field.
func (o *KalturaResponseProfileMapping) SetParentProperty(v string) {
	o.ParentProperty = &v
}

func (o KalturaResponseProfileMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowNull != nil {
		toSerialize["allowNull"] = o.AllowNull
	}
	if o.FilterProperty != nil {
		toSerialize["filterProperty"] = o.FilterProperty
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.ParentProperty != nil {
		toSerialize["parentProperty"] = o.ParentProperty
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaResponseProfileMapping struct {
	value *KalturaResponseProfileMapping
	isSet bool
}

func (v NullableKalturaResponseProfileMapping) Get() *KalturaResponseProfileMapping {
	return v.value
}

func (v *NullableKalturaResponseProfileMapping) Set(val *KalturaResponseProfileMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaResponseProfileMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaResponseProfileMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaResponseProfileMapping(val *KalturaResponseProfileMapping) *NullableKalturaResponseProfileMapping {
	return &NullableKalturaResponseProfileMapping{value: val, isSet: true}
}

func (v NullableKalturaResponseProfileMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaResponseProfileMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

