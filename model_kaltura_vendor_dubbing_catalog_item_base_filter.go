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

// KalturaVendorDubbingCatalogItemBaseFilter `abstract`
type KalturaVendorDubbingCatalogItemBaseFilter struct {
	KalturaVendorCatalogItemFilter
}

// NewKalturaVendorDubbingCatalogItemBaseFilter instantiates a new KalturaVendorDubbingCatalogItemBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVendorDubbingCatalogItemBaseFilter() *KalturaVendorDubbingCatalogItemBaseFilter {
	this := KalturaVendorDubbingCatalogItemBaseFilter{}
	return &this
}

// NewKalturaVendorDubbingCatalogItemBaseFilterWithDefaults instantiates a new KalturaVendorDubbingCatalogItemBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVendorDubbingCatalogItemBaseFilterWithDefaults() *KalturaVendorDubbingCatalogItemBaseFilter {
	this := KalturaVendorDubbingCatalogItemBaseFilter{}
	return &this
}

func (o KalturaVendorDubbingCatalogItemBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVendorCatalogItemFilter, errKalturaVendorCatalogItemFilter := json.Marshal(o.KalturaVendorCatalogItemFilter)
	if errKalturaVendorCatalogItemFilter != nil {
		return []byte{}, errKalturaVendorCatalogItemFilter
	}
	errKalturaVendorCatalogItemFilter = json.Unmarshal([]byte(serializedKalturaVendorCatalogItemFilter), &toSerialize)
	if errKalturaVendorCatalogItemFilter != nil {
		return []byte{}, errKalturaVendorCatalogItemFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVendorDubbingCatalogItemBaseFilter struct {
	value *KalturaVendorDubbingCatalogItemBaseFilter
	isSet bool
}

func (v NullableKalturaVendorDubbingCatalogItemBaseFilter) Get() *KalturaVendorDubbingCatalogItemBaseFilter {
	return v.value
}

func (v *NullableKalturaVendorDubbingCatalogItemBaseFilter) Set(val *KalturaVendorDubbingCatalogItemBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVendorDubbingCatalogItemBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVendorDubbingCatalogItemBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVendorDubbingCatalogItemBaseFilter(val *KalturaVendorDubbingCatalogItemBaseFilter) *NullableKalturaVendorDubbingCatalogItemBaseFilter {
	return &NullableKalturaVendorDubbingCatalogItemBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaVendorDubbingCatalogItemBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVendorDubbingCatalogItemBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


