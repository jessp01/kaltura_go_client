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

// MetadataAddFromBulkRequest struct for MetadataAddFromBulkRequest
type MetadataAddFromBulkRequest struct {
	MetadataProfileId int32 `json:"metadataProfileId"`
	ObjectId string `json:"objectId"`
	ObjectType string `json:"objectType"`
	Url string `json:"url"`
}

// NewMetadataAddFromBulkRequest instantiates a new MetadataAddFromBulkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataAddFromBulkRequest(metadataProfileId int32, objectId string, objectType string, url string) *MetadataAddFromBulkRequest {
	this := MetadataAddFromBulkRequest{}
	this.MetadataProfileId = metadataProfileId
	this.ObjectId = objectId
	this.ObjectType = objectType
	this.Url = url
	return &this
}

// NewMetadataAddFromBulkRequestWithDefaults instantiates a new MetadataAddFromBulkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataAddFromBulkRequestWithDefaults() *MetadataAddFromBulkRequest {
	this := MetadataAddFromBulkRequest{}
	return &this
}

// GetMetadataProfileId returns the MetadataProfileId field value
func (o *MetadataAddFromBulkRequest) GetMetadataProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MetadataProfileId
}

// GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field value
// and a boolean to check if the value has been set.
func (o *MetadataAddFromBulkRequest) GetMetadataProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataProfileId, true
}

// SetMetadataProfileId sets field value
func (o *MetadataAddFromBulkRequest) SetMetadataProfileId(v int32) {
	o.MetadataProfileId = v
}

// GetObjectId returns the ObjectId field value
func (o *MetadataAddFromBulkRequest) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *MetadataAddFromBulkRequest) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *MetadataAddFromBulkRequest) SetObjectId(v string) {
	o.ObjectId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetadataAddFromBulkRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetadataAddFromBulkRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetadataAddFromBulkRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUrl returns the Url field value
func (o *MetadataAddFromBulkRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MetadataAddFromBulkRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *MetadataAddFromBulkRequest) SetUrl(v string) {
	o.Url = v
}

func (o MetadataAddFromBulkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadataProfileId"] = o.MetadataProfileId
	}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataAddFromBulkRequest struct {
	value *MetadataAddFromBulkRequest
	isSet bool
}

func (v NullableMetadataAddFromBulkRequest) Get() *MetadataAddFromBulkRequest {
	return v.value
}

func (v *NullableMetadataAddFromBulkRequest) Set(val *MetadataAddFromBulkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataAddFromBulkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataAddFromBulkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataAddFromBulkRequest(val *MetadataAddFromBulkRequest) *NullableMetadataAddFromBulkRequest {
	return &NullableMetadataAddFromBulkRequest{value: val, isSet: true}
}

func (v NullableMetadataAddFromBulkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataAddFromBulkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


