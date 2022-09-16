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

// KalturaLiveClusterMediaServerNodeFilter struct for KalturaLiveClusterMediaServerNodeFilter
type KalturaLiveClusterMediaServerNodeFilter struct {
	KalturaLiveClusterMediaServerNodeBaseFilter
}

// NewKalturaLiveClusterMediaServerNodeFilter instantiates a new KalturaLiveClusterMediaServerNodeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveClusterMediaServerNodeFilter() *KalturaLiveClusterMediaServerNodeFilter {
	this := KalturaLiveClusterMediaServerNodeFilter{}
	return &this
}

// NewKalturaLiveClusterMediaServerNodeFilterWithDefaults instantiates a new KalturaLiveClusterMediaServerNodeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveClusterMediaServerNodeFilterWithDefaults() *KalturaLiveClusterMediaServerNodeFilter {
	this := KalturaLiveClusterMediaServerNodeFilter{}
	return &this
}

func (o KalturaLiveClusterMediaServerNodeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaLiveClusterMediaServerNodeBaseFilter, errKalturaLiveClusterMediaServerNodeBaseFilter := json.Marshal(o.KalturaLiveClusterMediaServerNodeBaseFilter)
	if errKalturaLiveClusterMediaServerNodeBaseFilter != nil {
		return []byte{}, errKalturaLiveClusterMediaServerNodeBaseFilter
	}
	errKalturaLiveClusterMediaServerNodeBaseFilter = json.Unmarshal([]byte(serializedKalturaLiveClusterMediaServerNodeBaseFilter), &toSerialize)
	if errKalturaLiveClusterMediaServerNodeBaseFilter != nil {
		return []byte{}, errKalturaLiveClusterMediaServerNodeBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveClusterMediaServerNodeFilter struct {
	value *KalturaLiveClusterMediaServerNodeFilter
	isSet bool
}

func (v NullableKalturaLiveClusterMediaServerNodeFilter) Get() *KalturaLiveClusterMediaServerNodeFilter {
	return v.value
}

func (v *NullableKalturaLiveClusterMediaServerNodeFilter) Set(val *KalturaLiveClusterMediaServerNodeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveClusterMediaServerNodeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveClusterMediaServerNodeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveClusterMediaServerNodeFilter(val *KalturaLiveClusterMediaServerNodeFilter) *NullableKalturaLiveClusterMediaServerNodeFilter {
	return &NullableKalturaLiveClusterMediaServerNodeFilter{value: val, isSet: true}
}

func (v NullableKalturaLiveClusterMediaServerNodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveClusterMediaServerNodeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


