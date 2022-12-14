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

// KalturaGenericDistributionProviderListResponse struct for KalturaGenericDistributionProviderListResponse
type KalturaGenericDistributionProviderListResponse struct {
	KalturaListResponse
}

// NewKalturaGenericDistributionProviderListResponse instantiates a new KalturaGenericDistributionProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGenericDistributionProviderListResponse() *KalturaGenericDistributionProviderListResponse {
	this := KalturaGenericDistributionProviderListResponse{}
	return &this
}

// NewKalturaGenericDistributionProviderListResponseWithDefaults instantiates a new KalturaGenericDistributionProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGenericDistributionProviderListResponseWithDefaults() *KalturaGenericDistributionProviderListResponse {
	this := KalturaGenericDistributionProviderListResponse{}
	return &this
}

func (o KalturaGenericDistributionProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaListResponse, errKalturaListResponse := json.Marshal(o.KalturaListResponse)
	if errKalturaListResponse != nil {
		return []byte{}, errKalturaListResponse
	}
	errKalturaListResponse = json.Unmarshal([]byte(serializedKalturaListResponse), &toSerialize)
	if errKalturaListResponse != nil {
		return []byte{}, errKalturaListResponse
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaGenericDistributionProviderListResponse struct {
	value *KalturaGenericDistributionProviderListResponse
	isSet bool
}

func (v NullableKalturaGenericDistributionProviderListResponse) Get() *KalturaGenericDistributionProviderListResponse {
	return v.value
}

func (v *NullableKalturaGenericDistributionProviderListResponse) Set(val *KalturaGenericDistributionProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGenericDistributionProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGenericDistributionProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGenericDistributionProviderListResponse(val *KalturaGenericDistributionProviderListResponse) *NullableKalturaGenericDistributionProviderListResponse {
	return &NullableKalturaGenericDistributionProviderListResponse{value: val, isSet: true}
}

func (v NullableKalturaGenericDistributionProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGenericDistributionProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


