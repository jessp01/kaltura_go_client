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

// ReportGetTableRequest struct for ReportGetTableRequest
type ReportGetTableRequest struct {
	ObjectIds *string `json:"objectIds,omitempty"`
	Order *string `json:"order,omitempty"`
	Pager KalturaFilterPager `json:"pager"`
	ReportInputFilter KalturaReportInputFilter `json:"reportInputFilter"`
	ReportType string `json:"reportType"`
	ResponseOptions *KalturaReportResponseOptions `json:"responseOptions,omitempty"`
}

// NewReportGetTableRequest instantiates a new ReportGetTableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportGetTableRequest(pager KalturaFilterPager, reportInputFilter KalturaReportInputFilter, reportType string) *ReportGetTableRequest {
	this := ReportGetTableRequest{}
	this.Pager = pager
	this.ReportInputFilter = reportInputFilter
	this.ReportType = reportType
	return &this
}

// NewReportGetTableRequestWithDefaults instantiates a new ReportGetTableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportGetTableRequestWithDefaults() *ReportGetTableRequest {
	this := ReportGetTableRequest{}
	return &this
}

// GetObjectIds returns the ObjectIds field value if set, zero value otherwise.
func (o *ReportGetTableRequest) GetObjectIds() string {
	if o == nil || o.ObjectIds == nil {
		var ret string
		return ret
	}
	return *o.ObjectIds
}

// GetObjectIdsOk returns a tuple with the ObjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetObjectIdsOk() (*string, bool) {
	if o == nil || o.ObjectIds == nil {
		return nil, false
	}
	return o.ObjectIds, true
}

// HasObjectIds returns a boolean if a field has been set.
func (o *ReportGetTableRequest) HasObjectIds() bool {
	if o != nil && o.ObjectIds != nil {
		return true
	}

	return false
}

// SetObjectIds gets a reference to the given string and assigns it to the ObjectIds field.
func (o *ReportGetTableRequest) SetObjectIds(v string) {
	o.ObjectIds = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ReportGetTableRequest) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ReportGetTableRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ReportGetTableRequest) SetOrder(v string) {
	o.Order = &v
}

// GetPager returns the Pager field value
func (o *ReportGetTableRequest) GetPager() KalturaFilterPager {
	if o == nil {
		var ret KalturaFilterPager
		return ret
	}

	return o.Pager
}

// GetPagerOk returns a tuple with the Pager field value
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pager, true
}

// SetPager sets field value
func (o *ReportGetTableRequest) SetPager(v KalturaFilterPager) {
	o.Pager = v
}

// GetReportInputFilter returns the ReportInputFilter field value
func (o *ReportGetTableRequest) GetReportInputFilter() KalturaReportInputFilter {
	if o == nil {
		var ret KalturaReportInputFilter
		return ret
	}

	return o.ReportInputFilter
}

// GetReportInputFilterOk returns a tuple with the ReportInputFilter field value
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetReportInputFilterOk() (*KalturaReportInputFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportInputFilter, true
}

// SetReportInputFilter sets field value
func (o *ReportGetTableRequest) SetReportInputFilter(v KalturaReportInputFilter) {
	o.ReportInputFilter = v
}

// GetReportType returns the ReportType field value
func (o *ReportGetTableRequest) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *ReportGetTableRequest) SetReportType(v string) {
	o.ReportType = v
}

// GetResponseOptions returns the ResponseOptions field value if set, zero value otherwise.
func (o *ReportGetTableRequest) GetResponseOptions() KalturaReportResponseOptions {
	if o == nil || o.ResponseOptions == nil {
		var ret KalturaReportResponseOptions
		return ret
	}
	return *o.ResponseOptions
}

// GetResponseOptionsOk returns a tuple with the ResponseOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportGetTableRequest) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool) {
	if o == nil || o.ResponseOptions == nil {
		return nil, false
	}
	return o.ResponseOptions, true
}

// HasResponseOptions returns a boolean if a field has been set.
func (o *ReportGetTableRequest) HasResponseOptions() bool {
	if o != nil && o.ResponseOptions != nil {
		return true
	}

	return false
}

// SetResponseOptions gets a reference to the given KalturaReportResponseOptions and assigns it to the ResponseOptions field.
func (o *ReportGetTableRequest) SetResponseOptions(v KalturaReportResponseOptions) {
	o.ResponseOptions = &v
}

func (o ReportGetTableRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectIds != nil {
		toSerialize["objectIds"] = o.ObjectIds
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if true {
		toSerialize["pager"] = o.Pager
	}
	if true {
		toSerialize["reportInputFilter"] = o.ReportInputFilter
	}
	if true {
		toSerialize["reportType"] = o.ReportType
	}
	if o.ResponseOptions != nil {
		toSerialize["responseOptions"] = o.ResponseOptions
	}
	return json.Marshal(toSerialize)
}

type NullableReportGetTableRequest struct {
	value *ReportGetTableRequest
	isSet bool
}

func (v NullableReportGetTableRequest) Get() *ReportGetTableRequest {
	return v.value
}

func (v *NullableReportGetTableRequest) Set(val *ReportGetTableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportGetTableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportGetTableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportGetTableRequest(val *ReportGetTableRequest) *NullableReportGetTableRequest {
	return &NullableReportGetTableRequest{value: val, isSet: true}
}

func (v NullableReportGetTableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportGetTableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


