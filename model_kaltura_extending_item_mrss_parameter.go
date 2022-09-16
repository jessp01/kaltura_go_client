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

// KalturaExtendingItemMrssParameter struct for KalturaExtendingItemMrssParameter
type KalturaExtendingItemMrssParameter struct {
	// Enum Type: `KalturaMrssExtensionMode`  Mode of extension - append to MRSS or replace the xpath content.
	ExtensionMode *int32 `json:"extensionMode,omitempty"`
	Identifier *KalturaObjectIdentifier `json:"identifier,omitempty"`
	// XPath for the extending item
	Xpath *string `json:"xpath,omitempty"`
}

// NewKalturaExtendingItemMrssParameter instantiates a new KalturaExtendingItemMrssParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaExtendingItemMrssParameter() *KalturaExtendingItemMrssParameter {
	this := KalturaExtendingItemMrssParameter{}
	return &this
}

// NewKalturaExtendingItemMrssParameterWithDefaults instantiates a new KalturaExtendingItemMrssParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaExtendingItemMrssParameterWithDefaults() *KalturaExtendingItemMrssParameter {
	this := KalturaExtendingItemMrssParameter{}
	return &this
}

// GetExtensionMode returns the ExtensionMode field value if set, zero value otherwise.
func (o *KalturaExtendingItemMrssParameter) GetExtensionMode() int32 {
	if o == nil || o.ExtensionMode == nil {
		var ret int32
		return ret
	}
	return *o.ExtensionMode
}

// GetExtensionModeOk returns a tuple with the ExtensionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaExtendingItemMrssParameter) GetExtensionModeOk() (*int32, bool) {
	if o == nil || o.ExtensionMode == nil {
		return nil, false
	}
	return o.ExtensionMode, true
}

// HasExtensionMode returns a boolean if a field has been set.
func (o *KalturaExtendingItemMrssParameter) HasExtensionMode() bool {
	if o != nil && o.ExtensionMode != nil {
		return true
	}

	return false
}

// SetExtensionMode gets a reference to the given int32 and assigns it to the ExtensionMode field.
func (o *KalturaExtendingItemMrssParameter) SetExtensionMode(v int32) {
	o.ExtensionMode = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *KalturaExtendingItemMrssParameter) GetIdentifier() KalturaObjectIdentifier {
	if o == nil || o.Identifier == nil {
		var ret KalturaObjectIdentifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaExtendingItemMrssParameter) GetIdentifierOk() (*KalturaObjectIdentifier, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *KalturaExtendingItemMrssParameter) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given KalturaObjectIdentifier and assigns it to the Identifier field.
func (o *KalturaExtendingItemMrssParameter) SetIdentifier(v KalturaObjectIdentifier) {
	o.Identifier = &v
}

// GetXpath returns the Xpath field value if set, zero value otherwise.
func (o *KalturaExtendingItemMrssParameter) GetXpath() string {
	if o == nil || o.Xpath == nil {
		var ret string
		return ret
	}
	return *o.Xpath
}

// GetXpathOk returns a tuple with the Xpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaExtendingItemMrssParameter) GetXpathOk() (*string, bool) {
	if o == nil || o.Xpath == nil {
		return nil, false
	}
	return o.Xpath, true
}

// HasXpath returns a boolean if a field has been set.
func (o *KalturaExtendingItemMrssParameter) HasXpath() bool {
	if o != nil && o.Xpath != nil {
		return true
	}

	return false
}

// SetXpath gets a reference to the given string and assigns it to the Xpath field.
func (o *KalturaExtendingItemMrssParameter) SetXpath(v string) {
	o.Xpath = &v
}

func (o KalturaExtendingItemMrssParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExtensionMode != nil {
		toSerialize["extensionMode"] = o.ExtensionMode
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Xpath != nil {
		toSerialize["xpath"] = o.Xpath
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaExtendingItemMrssParameter struct {
	value *KalturaExtendingItemMrssParameter
	isSet bool
}

func (v NullableKalturaExtendingItemMrssParameter) Get() *KalturaExtendingItemMrssParameter {
	return v.value
}

func (v *NullableKalturaExtendingItemMrssParameter) Set(val *KalturaExtendingItemMrssParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaExtendingItemMrssParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaExtendingItemMrssParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaExtendingItemMrssParameter(val *KalturaExtendingItemMrssParameter) *NullableKalturaExtendingItemMrssParameter {
	return &NullableKalturaExtendingItemMrssParameter{value: val, isSet: true}
}

func (v NullableKalturaExtendingItemMrssParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaExtendingItemMrssParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


