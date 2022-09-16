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

// GenericDistributionProviderActionAddRequest struct for GenericDistributionProviderActionAddRequest
type GenericDistributionProviderActionAddRequest struct {
	GenericDistributionProviderAction KalturaGenericDistributionProviderAction `json:"genericDistributionProviderAction"`
}

// NewGenericDistributionProviderActionAddRequest instantiates a new GenericDistributionProviderActionAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDistributionProviderActionAddRequest(genericDistributionProviderAction KalturaGenericDistributionProviderAction) *GenericDistributionProviderActionAddRequest {
	this := GenericDistributionProviderActionAddRequest{}
	this.GenericDistributionProviderAction = genericDistributionProviderAction
	return &this
}

// NewGenericDistributionProviderActionAddRequestWithDefaults instantiates a new GenericDistributionProviderActionAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDistributionProviderActionAddRequestWithDefaults() *GenericDistributionProviderActionAddRequest {
	this := GenericDistributionProviderActionAddRequest{}
	return &this
}

// GetGenericDistributionProviderAction returns the GenericDistributionProviderAction field value
func (o *GenericDistributionProviderActionAddRequest) GetGenericDistributionProviderAction() KalturaGenericDistributionProviderAction {
	if o == nil {
		var ret KalturaGenericDistributionProviderAction
		return ret
	}

	return o.GenericDistributionProviderAction
}

// GetGenericDistributionProviderActionOk returns a tuple with the GenericDistributionProviderAction field value
// and a boolean to check if the value has been set.
func (o *GenericDistributionProviderActionAddRequest) GetGenericDistributionProviderActionOk() (*KalturaGenericDistributionProviderAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenericDistributionProviderAction, true
}

// SetGenericDistributionProviderAction sets field value
func (o *GenericDistributionProviderActionAddRequest) SetGenericDistributionProviderAction(v KalturaGenericDistributionProviderAction) {
	o.GenericDistributionProviderAction = v
}

func (o GenericDistributionProviderActionAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["genericDistributionProviderAction"] = o.GenericDistributionProviderAction
	}
	return json.Marshal(toSerialize)
}

type NullableGenericDistributionProviderActionAddRequest struct {
	value *GenericDistributionProviderActionAddRequest
	isSet bool
}

func (v NullableGenericDistributionProviderActionAddRequest) Get() *GenericDistributionProviderActionAddRequest {
	return v.value
}

func (v *NullableGenericDistributionProviderActionAddRequest) Set(val *GenericDistributionProviderActionAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDistributionProviderActionAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDistributionProviderActionAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDistributionProviderActionAddRequest(val *GenericDistributionProviderActionAddRequest) *NullableGenericDistributionProviderActionAddRequest {
	return &NullableGenericDistributionProviderActionAddRequest{value: val, isSet: true}
}

func (v NullableGenericDistributionProviderActionAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDistributionProviderActionAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

