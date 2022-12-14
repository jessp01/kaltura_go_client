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

// KalturaCaptionAssetListResponse struct for KalturaCaptionAssetListResponse
type KalturaCaptionAssetListResponse struct {
	KalturaListResponse
}

// NewKalturaCaptionAssetListResponse instantiates a new KalturaCaptionAssetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCaptionAssetListResponse() *KalturaCaptionAssetListResponse {
	this := KalturaCaptionAssetListResponse{}
	return &this
}

// NewKalturaCaptionAssetListResponseWithDefaults instantiates a new KalturaCaptionAssetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCaptionAssetListResponseWithDefaults() *KalturaCaptionAssetListResponse {
	this := KalturaCaptionAssetListResponse{}
	return &this
}

func (o KalturaCaptionAssetListResponse) MarshalJSON() ([]byte, error) {
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

type NullableKalturaCaptionAssetListResponse struct {
	value *KalturaCaptionAssetListResponse
	isSet bool
}

func (v NullableKalturaCaptionAssetListResponse) Get() *KalturaCaptionAssetListResponse {
	return v.value
}

func (v *NullableKalturaCaptionAssetListResponse) Set(val *KalturaCaptionAssetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCaptionAssetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCaptionAssetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCaptionAssetListResponse(val *KalturaCaptionAssetListResponse) *NullableKalturaCaptionAssetListResponse {
	return &NullableKalturaCaptionAssetListResponse{value: val, isSet: true}
}

func (v NullableKalturaCaptionAssetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCaptionAssetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


