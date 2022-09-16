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

// ContentDistributionBatchGetAssetUrlRequest struct for ContentDistributionBatchGetAssetUrlRequest
type ContentDistributionBatchGetAssetUrlRequest struct {
	AssetId string `json:"assetId"`
}

// NewContentDistributionBatchGetAssetUrlRequest instantiates a new ContentDistributionBatchGetAssetUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentDistributionBatchGetAssetUrlRequest(assetId string) *ContentDistributionBatchGetAssetUrlRequest {
	this := ContentDistributionBatchGetAssetUrlRequest{}
	this.AssetId = assetId
	return &this
}

// NewContentDistributionBatchGetAssetUrlRequestWithDefaults instantiates a new ContentDistributionBatchGetAssetUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentDistributionBatchGetAssetUrlRequestWithDefaults() *ContentDistributionBatchGetAssetUrlRequest {
	this := ContentDistributionBatchGetAssetUrlRequest{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *ContentDistributionBatchGetAssetUrlRequest) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *ContentDistributionBatchGetAssetUrlRequest) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *ContentDistributionBatchGetAssetUrlRequest) SetAssetId(v string) {
	o.AssetId = v
}

func (o ContentDistributionBatchGetAssetUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assetId"] = o.AssetId
	}
	return json.Marshal(toSerialize)
}

type NullableContentDistributionBatchGetAssetUrlRequest struct {
	value *ContentDistributionBatchGetAssetUrlRequest
	isSet bool
}

func (v NullableContentDistributionBatchGetAssetUrlRequest) Get() *ContentDistributionBatchGetAssetUrlRequest {
	return v.value
}

func (v *NullableContentDistributionBatchGetAssetUrlRequest) Set(val *ContentDistributionBatchGetAssetUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContentDistributionBatchGetAssetUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContentDistributionBatchGetAssetUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentDistributionBatchGetAssetUrlRequest(val *ContentDistributionBatchGetAssetUrlRequest) *NullableContentDistributionBatchGetAssetUrlRequest {
	return &NullableContentDistributionBatchGetAssetUrlRequest{value: val, isSet: true}
}

func (v NullableContentDistributionBatchGetAssetUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentDistributionBatchGetAssetUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

