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

// KalturaFlavorAssetWithParams struct for KalturaFlavorAssetWithParams
type KalturaFlavorAssetWithParams struct {
	// The entry id
	EntryId *string `json:"entryId,omitempty"`
	FlavorAsset *KalturaFlavorAsset `json:"flavorAsset,omitempty"`
	FlavorParams *KalturaFlavorParams `json:"flavorParams,omitempty"`
}

// NewKalturaFlavorAssetWithParams instantiates a new KalturaFlavorAssetWithParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFlavorAssetWithParams() *KalturaFlavorAssetWithParams {
	this := KalturaFlavorAssetWithParams{}
	return &this
}

// NewKalturaFlavorAssetWithParamsWithDefaults instantiates a new KalturaFlavorAssetWithParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFlavorAssetWithParamsWithDefaults() *KalturaFlavorAssetWithParams {
	this := KalturaFlavorAssetWithParams{}
	return &this
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaFlavorAssetWithParams) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFlavorAssetWithParams) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaFlavorAssetWithParams) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaFlavorAssetWithParams) SetEntryId(v string) {
	o.EntryId = &v
}

// GetFlavorAsset returns the FlavorAsset field value if set, zero value otherwise.
func (o *KalturaFlavorAssetWithParams) GetFlavorAsset() KalturaFlavorAsset {
	if o == nil || o.FlavorAsset == nil {
		var ret KalturaFlavorAsset
		return ret
	}
	return *o.FlavorAsset
}

// GetFlavorAssetOk returns a tuple with the FlavorAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFlavorAssetWithParams) GetFlavorAssetOk() (*KalturaFlavorAsset, bool) {
	if o == nil || o.FlavorAsset == nil {
		return nil, false
	}
	return o.FlavorAsset, true
}

// HasFlavorAsset returns a boolean if a field has been set.
func (o *KalturaFlavorAssetWithParams) HasFlavorAsset() bool {
	if o != nil && o.FlavorAsset != nil {
		return true
	}

	return false
}

// SetFlavorAsset gets a reference to the given KalturaFlavorAsset and assigns it to the FlavorAsset field.
func (o *KalturaFlavorAssetWithParams) SetFlavorAsset(v KalturaFlavorAsset) {
	o.FlavorAsset = &v
}

// GetFlavorParams returns the FlavorParams field value if set, zero value otherwise.
func (o *KalturaFlavorAssetWithParams) GetFlavorParams() KalturaFlavorParams {
	if o == nil || o.FlavorParams == nil {
		var ret KalturaFlavorParams
		return ret
	}
	return *o.FlavorParams
}

// GetFlavorParamsOk returns a tuple with the FlavorParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFlavorAssetWithParams) GetFlavorParamsOk() (*KalturaFlavorParams, bool) {
	if o == nil || o.FlavorParams == nil {
		return nil, false
	}
	return o.FlavorParams, true
}

// HasFlavorParams returns a boolean if a field has been set.
func (o *KalturaFlavorAssetWithParams) HasFlavorParams() bool {
	if o != nil && o.FlavorParams != nil {
		return true
	}

	return false
}

// SetFlavorParams gets a reference to the given KalturaFlavorParams and assigns it to the FlavorParams field.
func (o *KalturaFlavorAssetWithParams) SetFlavorParams(v KalturaFlavorParams) {
	o.FlavorParams = &v
}

func (o KalturaFlavorAssetWithParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.FlavorAsset != nil {
		toSerialize["flavorAsset"] = o.FlavorAsset
	}
	if o.FlavorParams != nil {
		toSerialize["flavorParams"] = o.FlavorParams
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFlavorAssetWithParams struct {
	value *KalturaFlavorAssetWithParams
	isSet bool
}

func (v NullableKalturaFlavorAssetWithParams) Get() *KalturaFlavorAssetWithParams {
	return v.value
}

func (v *NullableKalturaFlavorAssetWithParams) Set(val *KalturaFlavorAssetWithParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFlavorAssetWithParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFlavorAssetWithParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFlavorAssetWithParams(val *KalturaFlavorAssetWithParams) *NullableKalturaFlavorAssetWithParams {
	return &NullableKalturaFlavorAssetWithParams{value: val, isSet: true}
}

func (v NullableKalturaFlavorAssetWithParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFlavorAssetWithParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


