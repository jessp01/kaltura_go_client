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

// DrmProfileListRequest struct for DrmProfileListRequest
type DrmProfileListRequest struct {
	Filter *KalturaDrmProfileFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewDrmProfileListRequest instantiates a new DrmProfileListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrmProfileListRequest() *DrmProfileListRequest {
	this := DrmProfileListRequest{}
	return &this
}

// NewDrmProfileListRequestWithDefaults instantiates a new DrmProfileListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrmProfileListRequestWithDefaults() *DrmProfileListRequest {
	this := DrmProfileListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DrmProfileListRequest) GetFilter() KalturaDrmProfileFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaDrmProfileFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrmProfileListRequest) GetFilterOk() (*KalturaDrmProfileFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DrmProfileListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaDrmProfileFilter and assigns it to the Filter field.
func (o *DrmProfileListRequest) SetFilter(v KalturaDrmProfileFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DrmProfileListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrmProfileListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DrmProfileListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *DrmProfileListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o DrmProfileListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableDrmProfileListRequest struct {
	value *DrmProfileListRequest
	isSet bool
}

func (v NullableDrmProfileListRequest) Get() *DrmProfileListRequest {
	return v.value
}

func (v *NullableDrmProfileListRequest) Set(val *DrmProfileListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDrmProfileListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDrmProfileListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrmProfileListRequest(val *DrmProfileListRequest) *NullableDrmProfileListRequest {
	return &NullableDrmProfileListRequest{value: val, isSet: true}
}

func (v NullableDrmProfileListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrmProfileListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


