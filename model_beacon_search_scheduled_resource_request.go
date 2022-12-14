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

// BeaconSearchScheduledResourceRequest struct for BeaconSearchScheduledResourceRequest
type BeaconSearchScheduledResourceRequest struct {
	Pager *KalturaPager `json:"pager,omitempty"`
	SearchParams KalturaBeaconSearchParams `json:"searchParams"`
}

// NewBeaconSearchScheduledResourceRequest instantiates a new BeaconSearchScheduledResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeaconSearchScheduledResourceRequest(searchParams KalturaBeaconSearchParams) *BeaconSearchScheduledResourceRequest {
	this := BeaconSearchScheduledResourceRequest{}
	this.SearchParams = searchParams
	return &this
}

// NewBeaconSearchScheduledResourceRequestWithDefaults instantiates a new BeaconSearchScheduledResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeaconSearchScheduledResourceRequestWithDefaults() *BeaconSearchScheduledResourceRequest {
	this := BeaconSearchScheduledResourceRequest{}
	return &this
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *BeaconSearchScheduledResourceRequest) GetPager() KalturaPager {
	if o == nil || o.Pager == nil {
		var ret KalturaPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeaconSearchScheduledResourceRequest) GetPagerOk() (*KalturaPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *BeaconSearchScheduledResourceRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaPager and assigns it to the Pager field.
func (o *BeaconSearchScheduledResourceRequest) SetPager(v KalturaPager) {
	o.Pager = &v
}

// GetSearchParams returns the SearchParams field value
func (o *BeaconSearchScheduledResourceRequest) GetSearchParams() KalturaBeaconSearchParams {
	if o == nil {
		var ret KalturaBeaconSearchParams
		return ret
	}

	return o.SearchParams
}

// GetSearchParamsOk returns a tuple with the SearchParams field value
// and a boolean to check if the value has been set.
func (o *BeaconSearchScheduledResourceRequest) GetSearchParamsOk() (*KalturaBeaconSearchParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchParams, true
}

// SetSearchParams sets field value
func (o *BeaconSearchScheduledResourceRequest) SetSearchParams(v KalturaBeaconSearchParams) {
	o.SearchParams = v
}

func (o BeaconSearchScheduledResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	if true {
		toSerialize["searchParams"] = o.SearchParams
	}
	return json.Marshal(toSerialize)
}

type NullableBeaconSearchScheduledResourceRequest struct {
	value *BeaconSearchScheduledResourceRequest
	isSet bool
}

func (v NullableBeaconSearchScheduledResourceRequest) Get() *BeaconSearchScheduledResourceRequest {
	return v.value
}

func (v *NullableBeaconSearchScheduledResourceRequest) Set(val *BeaconSearchScheduledResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBeaconSearchScheduledResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBeaconSearchScheduledResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeaconSearchScheduledResourceRequest(val *BeaconSearchScheduledResourceRequest) *NullableBeaconSearchScheduledResourceRequest {
	return &NullableBeaconSearchScheduledResourceRequest{value: val, isSet: true}
}

func (v NullableBeaconSearchScheduledResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeaconSearchScheduledResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


