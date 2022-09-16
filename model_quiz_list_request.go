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

// QuizListRequest struct for QuizListRequest
type QuizListRequest struct {
	Filter *KalturaQuizFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewQuizListRequest instantiates a new QuizListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuizListRequest() *QuizListRequest {
	this := QuizListRequest{}
	return &this
}

// NewQuizListRequestWithDefaults instantiates a new QuizListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuizListRequestWithDefaults() *QuizListRequest {
	this := QuizListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *QuizListRequest) GetFilter() KalturaQuizFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaQuizFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizListRequest) GetFilterOk() (*KalturaQuizFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *QuizListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaQuizFilter and assigns it to the Filter field.
func (o *QuizListRequest) SetFilter(v KalturaQuizFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *QuizListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuizListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *QuizListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *QuizListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o QuizListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableQuizListRequest struct {
	value *QuizListRequest
	isSet bool
}

func (v NullableQuizListRequest) Get() *QuizListRequest {
	return v.value
}

func (v *NullableQuizListRequest) Set(val *QuizListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuizListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuizListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuizListRequest(val *QuizListRequest) *NullableQuizListRequest {
	return &NullableQuizListRequest{value: val, isSet: true}
}

func (v NullableQuizListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuizListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


