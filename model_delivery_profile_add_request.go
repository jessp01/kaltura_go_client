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

// DeliveryProfileAddRequest struct for DeliveryProfileAddRequest
type DeliveryProfileAddRequest struct {
	Delivery KalturaDeliveryProfile `json:"delivery"`
}

// NewDeliveryProfileAddRequest instantiates a new DeliveryProfileAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryProfileAddRequest(delivery KalturaDeliveryProfile) *DeliveryProfileAddRequest {
	this := DeliveryProfileAddRequest{}
	this.Delivery = delivery
	return &this
}

// NewDeliveryProfileAddRequestWithDefaults instantiates a new DeliveryProfileAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryProfileAddRequestWithDefaults() *DeliveryProfileAddRequest {
	this := DeliveryProfileAddRequest{}
	return &this
}

// GetDelivery returns the Delivery field value
func (o *DeliveryProfileAddRequest) GetDelivery() KalturaDeliveryProfile {
	if o == nil {
		var ret KalturaDeliveryProfile
		return ret
	}

	return o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value
// and a boolean to check if the value has been set.
func (o *DeliveryProfileAddRequest) GetDeliveryOk() (*KalturaDeliveryProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delivery, true
}

// SetDelivery sets field value
func (o *DeliveryProfileAddRequest) SetDelivery(v KalturaDeliveryProfile) {
	o.Delivery = v
}

func (o DeliveryProfileAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["delivery"] = o.Delivery
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryProfileAddRequest struct {
	value *DeliveryProfileAddRequest
	isSet bool
}

func (v NullableDeliveryProfileAddRequest) Get() *DeliveryProfileAddRequest {
	return v.value
}

func (v *NullableDeliveryProfileAddRequest) Set(val *DeliveryProfileAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryProfileAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryProfileAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryProfileAddRequest(val *DeliveryProfileAddRequest) *NullableDeliveryProfileAddRequest {
	return &NullableDeliveryProfileAddRequest{value: val, isSet: true}
}

func (v NullableDeliveryProfileAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryProfileAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


