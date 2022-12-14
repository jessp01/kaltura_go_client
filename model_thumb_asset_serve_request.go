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

// ThumbAssetServeRequest struct for ThumbAssetServeRequest
type ThumbAssetServeRequest struct {
	Options *KalturaThumbnailServeOptions `json:"options,omitempty"`
	ThumbAssetId string `json:"thumbAssetId"`
	ThumbParams *KalturaThumbParams `json:"thumbParams,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewThumbAssetServeRequest instantiates a new ThumbAssetServeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbAssetServeRequest(thumbAssetId string) *ThumbAssetServeRequest {
	this := ThumbAssetServeRequest{}
	this.ThumbAssetId = thumbAssetId
	return &this
}

// NewThumbAssetServeRequestWithDefaults instantiates a new ThumbAssetServeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbAssetServeRequestWithDefaults() *ThumbAssetServeRequest {
	this := ThumbAssetServeRequest{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ThumbAssetServeRequest) GetOptions() KalturaThumbnailServeOptions {
	if o == nil || o.Options == nil {
		var ret KalturaThumbnailServeOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeRequest) GetOptionsOk() (*KalturaThumbnailServeOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ThumbAssetServeRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given KalturaThumbnailServeOptions and assigns it to the Options field.
func (o *ThumbAssetServeRequest) SetOptions(v KalturaThumbnailServeOptions) {
	o.Options = &v
}

// GetThumbAssetId returns the ThumbAssetId field value
func (o *ThumbAssetServeRequest) GetThumbAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbAssetId
}

// GetThumbAssetIdOk returns a tuple with the ThumbAssetId field value
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeRequest) GetThumbAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbAssetId, true
}

// SetThumbAssetId sets field value
func (o *ThumbAssetServeRequest) SetThumbAssetId(v string) {
	o.ThumbAssetId = v
}

// GetThumbParams returns the ThumbParams field value if set, zero value otherwise.
func (o *ThumbAssetServeRequest) GetThumbParams() KalturaThumbParams {
	if o == nil || o.ThumbParams == nil {
		var ret KalturaThumbParams
		return ret
	}
	return *o.ThumbParams
}

// GetThumbParamsOk returns a tuple with the ThumbParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeRequest) GetThumbParamsOk() (*KalturaThumbParams, bool) {
	if o == nil || o.ThumbParams == nil {
		return nil, false
	}
	return o.ThumbParams, true
}

// HasThumbParams returns a boolean if a field has been set.
func (o *ThumbAssetServeRequest) HasThumbParams() bool {
	if o != nil && o.ThumbParams != nil {
		return true
	}

	return false
}

// SetThumbParams gets a reference to the given KalturaThumbParams and assigns it to the ThumbParams field.
func (o *ThumbAssetServeRequest) SetThumbParams(v KalturaThumbParams) {
	o.ThumbParams = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ThumbAssetServeRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ThumbAssetServeRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ThumbAssetServeRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o ThumbAssetServeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if true {
		toSerialize["thumbAssetId"] = o.ThumbAssetId
	}
	if o.ThumbParams != nil {
		toSerialize["thumbParams"] = o.ThumbParams
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableThumbAssetServeRequest struct {
	value *ThumbAssetServeRequest
	isSet bool
}

func (v NullableThumbAssetServeRequest) Get() *ThumbAssetServeRequest {
	return v.value
}

func (v *NullableThumbAssetServeRequest) Set(val *ThumbAssetServeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbAssetServeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbAssetServeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbAssetServeRequest(val *ThumbAssetServeRequest) *NullableThumbAssetServeRequest {
	return &NullableThumbAssetServeRequest{value: val, isSet: true}
}

func (v NullableThumbAssetServeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbAssetServeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


