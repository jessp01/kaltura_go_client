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

// KalturaLiveClusterMediaServerNodeBaseFilter struct for KalturaLiveClusterMediaServerNodeBaseFilter
type KalturaLiveClusterMediaServerNodeBaseFilter struct {
	KalturaMediaServerNodeFilter
}

// NewKalturaLiveClusterMediaServerNodeBaseFilter instantiates a new KalturaLiveClusterMediaServerNodeBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveClusterMediaServerNodeBaseFilter() *KalturaLiveClusterMediaServerNodeBaseFilter {
	this := KalturaLiveClusterMediaServerNodeBaseFilter{}
	return &this
}

// NewKalturaLiveClusterMediaServerNodeBaseFilterWithDefaults instantiates a new KalturaLiveClusterMediaServerNodeBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveClusterMediaServerNodeBaseFilterWithDefaults() *KalturaLiveClusterMediaServerNodeBaseFilter {
	this := KalturaLiveClusterMediaServerNodeBaseFilter{}
	return &this
}

func (o KalturaLiveClusterMediaServerNodeBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaServerNodeFilter, errKalturaMediaServerNodeFilter := json.Marshal(o.KalturaMediaServerNodeFilter)
	if errKalturaMediaServerNodeFilter != nil {
		return []byte{}, errKalturaMediaServerNodeFilter
	}
	errKalturaMediaServerNodeFilter = json.Unmarshal([]byte(serializedKalturaMediaServerNodeFilter), &toSerialize)
	if errKalturaMediaServerNodeFilter != nil {
		return []byte{}, errKalturaMediaServerNodeFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveClusterMediaServerNodeBaseFilter struct {
	value *KalturaLiveClusterMediaServerNodeBaseFilter
	isSet bool
}

func (v NullableKalturaLiveClusterMediaServerNodeBaseFilter) Get() *KalturaLiveClusterMediaServerNodeBaseFilter {
	return v.value
}

func (v *NullableKalturaLiveClusterMediaServerNodeBaseFilter) Set(val *KalturaLiveClusterMediaServerNodeBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveClusterMediaServerNodeBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveClusterMediaServerNodeBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveClusterMediaServerNodeBaseFilter(val *KalturaLiveClusterMediaServerNodeBaseFilter) *NullableKalturaLiveClusterMediaServerNodeBaseFilter {
	return &NullableKalturaLiveClusterMediaServerNodeBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaLiveClusterMediaServerNodeBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveClusterMediaServerNodeBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


