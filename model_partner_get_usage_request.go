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

// PartnerGetUsageRequest struct for PartnerGetUsageRequest
type PartnerGetUsageRequest struct {
	Month *int32 `json:"month,omitempty"`
	Resolution *string `json:"resolution,omitempty"`
	Year *int32 `json:"year,omitempty"`
}

// NewPartnerGetUsageRequest instantiates a new PartnerGetUsageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerGetUsageRequest() *PartnerGetUsageRequest {
	this := PartnerGetUsageRequest{}
	return &this
}

// NewPartnerGetUsageRequestWithDefaults instantiates a new PartnerGetUsageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerGetUsageRequestWithDefaults() *PartnerGetUsageRequest {
	this := PartnerGetUsageRequest{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *PartnerGetUsageRequest) GetMonth() int32 {
	if o == nil || o.Month == nil {
		var ret int32
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerGetUsageRequest) GetMonthOk() (*int32, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *PartnerGetUsageRequest) HasMonth() bool {
	if o != nil && o.Month != nil {
		return true
	}

	return false
}

// SetMonth gets a reference to the given int32 and assigns it to the Month field.
func (o *PartnerGetUsageRequest) SetMonth(v int32) {
	o.Month = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *PartnerGetUsageRequest) GetResolution() string {
	if o == nil || o.Resolution == nil {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerGetUsageRequest) GetResolutionOk() (*string, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *PartnerGetUsageRequest) HasResolution() bool {
	if o != nil && o.Resolution != nil {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *PartnerGetUsageRequest) SetResolution(v string) {
	o.Resolution = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *PartnerGetUsageRequest) GetYear() int32 {
	if o == nil || o.Year == nil {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerGetUsageRequest) GetYearOk() (*int32, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *PartnerGetUsageRequest) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *PartnerGetUsageRequest) SetYear(v int32) {
	o.Year = &v
}

func (o PartnerGetUsageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Month != nil {
		toSerialize["month"] = o.Month
	}
	if o.Resolution != nil {
		toSerialize["resolution"] = o.Resolution
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	return json.Marshal(toSerialize)
}

type NullablePartnerGetUsageRequest struct {
	value *PartnerGetUsageRequest
	isSet bool
}

func (v NullablePartnerGetUsageRequest) Get() *PartnerGetUsageRequest {
	return v.value
}

func (v *NullablePartnerGetUsageRequest) Set(val *PartnerGetUsageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerGetUsageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerGetUsageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerGetUsageRequest(val *PartnerGetUsageRequest) *NullablePartnerGetUsageRequest {
	return &NullablePartnerGetUsageRequest{value: val, isSet: true}
}

func (v NullablePartnerGetUsageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerGetUsageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


