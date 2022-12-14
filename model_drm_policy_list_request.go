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

// DrmPolicyListRequest struct for DrmPolicyListRequest
type DrmPolicyListRequest struct {
	Filter *KalturaDrmPolicyFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewDrmPolicyListRequest instantiates a new DrmPolicyListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDrmPolicyListRequest() *DrmPolicyListRequest {
	this := DrmPolicyListRequest{}
	return &this
}

// NewDrmPolicyListRequestWithDefaults instantiates a new DrmPolicyListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDrmPolicyListRequestWithDefaults() *DrmPolicyListRequest {
	this := DrmPolicyListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DrmPolicyListRequest) GetFilter() KalturaDrmPolicyFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaDrmPolicyFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrmPolicyListRequest) GetFilterOk() (*KalturaDrmPolicyFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DrmPolicyListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaDrmPolicyFilter and assigns it to the Filter field.
func (o *DrmPolicyListRequest) SetFilter(v KalturaDrmPolicyFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *DrmPolicyListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DrmPolicyListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *DrmPolicyListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *DrmPolicyListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o DrmPolicyListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableDrmPolicyListRequest struct {
	value *DrmPolicyListRequest
	isSet bool
}

func (v NullableDrmPolicyListRequest) Get() *DrmPolicyListRequest {
	return v.value
}

func (v *NullableDrmPolicyListRequest) Set(val *DrmPolicyListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDrmPolicyListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDrmPolicyListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDrmPolicyListRequest(val *DrmPolicyListRequest) *NullableDrmPolicyListRequest {
	return &NullableDrmPolicyListRequest{value: val, isSet: true}
}

func (v NullableDrmPolicyListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDrmPolicyListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


