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

// UiConfListRequest struct for UiConfListRequest
type UiConfListRequest struct {
	Filter *KalturaUiConfFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewUiConfListRequest instantiates a new UiConfListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiConfListRequest() *UiConfListRequest {
	this := UiConfListRequest{}
	return &this
}

// NewUiConfListRequestWithDefaults instantiates a new UiConfListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiConfListRequestWithDefaults() *UiConfListRequest {
	this := UiConfListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *UiConfListRequest) GetFilter() KalturaUiConfFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaUiConfFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfListRequest) GetFilterOk() (*KalturaUiConfFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *UiConfListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaUiConfFilter and assigns it to the Filter field.
func (o *UiConfListRequest) SetFilter(v KalturaUiConfFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *UiConfListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiConfListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *UiConfListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *UiConfListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o UiConfListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableUiConfListRequest struct {
	value *UiConfListRequest
	isSet bool
}

func (v NullableUiConfListRequest) Get() *UiConfListRequest {
	return v.value
}

func (v *NullableUiConfListRequest) Set(val *UiConfListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUiConfListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUiConfListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiConfListRequest(val *UiConfListRequest) *NullableUiConfListRequest {
	return &NullableUiConfListRequest{value: val, isSet: true}
}

func (v NullableUiConfListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiConfListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


