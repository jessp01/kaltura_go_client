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

// BusinessProcessCaseAbortRequest struct for BusinessProcessCaseAbortRequest
type BusinessProcessCaseAbortRequest struct {
	BusinessProcessStartNotificationTemplateId int32 `json:"businessProcessStartNotificationTemplateId"`
	ObjectId string `json:"objectId"`
	ObjectType string `json:"objectType"`
}

// NewBusinessProcessCaseAbortRequest instantiates a new BusinessProcessCaseAbortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessProcessCaseAbortRequest(businessProcessStartNotificationTemplateId int32, objectId string, objectType string) *BusinessProcessCaseAbortRequest {
	this := BusinessProcessCaseAbortRequest{}
	this.BusinessProcessStartNotificationTemplateId = businessProcessStartNotificationTemplateId
	this.ObjectId = objectId
	this.ObjectType = objectType
	return &this
}

// NewBusinessProcessCaseAbortRequestWithDefaults instantiates a new BusinessProcessCaseAbortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessProcessCaseAbortRequestWithDefaults() *BusinessProcessCaseAbortRequest {
	this := BusinessProcessCaseAbortRequest{}
	return &this
}

// GetBusinessProcessStartNotificationTemplateId returns the BusinessProcessStartNotificationTemplateId field value
func (o *BusinessProcessCaseAbortRequest) GetBusinessProcessStartNotificationTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BusinessProcessStartNotificationTemplateId
}

// GetBusinessProcessStartNotificationTemplateIdOk returns a tuple with the BusinessProcessStartNotificationTemplateId field value
// and a boolean to check if the value has been set.
func (o *BusinessProcessCaseAbortRequest) GetBusinessProcessStartNotificationTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessProcessStartNotificationTemplateId, true
}

// SetBusinessProcessStartNotificationTemplateId sets field value
func (o *BusinessProcessCaseAbortRequest) SetBusinessProcessStartNotificationTemplateId(v int32) {
	o.BusinessProcessStartNotificationTemplateId = v
}

// GetObjectId returns the ObjectId field value
func (o *BusinessProcessCaseAbortRequest) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *BusinessProcessCaseAbortRequest) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *BusinessProcessCaseAbortRequest) SetObjectId(v string) {
	o.ObjectId = v
}

// GetObjectType returns the ObjectType field value
func (o *BusinessProcessCaseAbortRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BusinessProcessCaseAbortRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BusinessProcessCaseAbortRequest) SetObjectType(v string) {
	o.ObjectType = v
}

func (o BusinessProcessCaseAbortRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["businessProcessStartNotificationTemplateId"] = o.BusinessProcessStartNotificationTemplateId
	}
	if true {
		toSerialize["objectId"] = o.ObjectId
	}
	if true {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableBusinessProcessCaseAbortRequest struct {
	value *BusinessProcessCaseAbortRequest
	isSet bool
}

func (v NullableBusinessProcessCaseAbortRequest) Get() *BusinessProcessCaseAbortRequest {
	return v.value
}

func (v *NullableBusinessProcessCaseAbortRequest) Set(val *BusinessProcessCaseAbortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessProcessCaseAbortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessProcessCaseAbortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessProcessCaseAbortRequest(val *BusinessProcessCaseAbortRequest) *NullableBusinessProcessCaseAbortRequest {
	return &NullableBusinessProcessCaseAbortRequest{value: val, isSet: true}
}

func (v NullableBusinessProcessCaseAbortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessProcessCaseAbortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


