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

// KalturaBeacon struct for KalturaBeacon
type KalturaBeacon struct {
	EventType *string `json:"eventType,omitempty"`
	// `readOnly`  Beacon id
	Id *string `json:"id,omitempty"`
	// `readOnly`  Beacon indexType
	IndexType *string `json:"indexType,omitempty"`
	ObjectId *string `json:"objectId,omitempty"`
	PrivateData *string `json:"privateData,omitempty"`
	RawData *string `json:"rawData,omitempty"`
	// Enum Type: `KalturaBeaconObjectTypes`  The object which this beacon belongs to
	RelatedObjectType *string `json:"relatedObjectType,omitempty"`
	// `readOnly`  Beacon update date as Unix timestamp (In seconds)
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaBeacon instantiates a new KalturaBeacon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBeacon() *KalturaBeacon {
	this := KalturaBeacon{}
	return &this
}

// NewKalturaBeaconWithDefaults instantiates a new KalturaBeacon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBeaconWithDefaults() *KalturaBeacon {
	this := KalturaBeacon{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *KalturaBeacon) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *KalturaBeacon) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *KalturaBeacon) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaBeacon) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaBeacon) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaBeacon) SetId(v string) {
	o.Id = &v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *KalturaBeacon) GetIndexType() string {
	if o == nil || o.IndexType == nil {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetIndexTypeOk() (*string, bool) {
	if o == nil || o.IndexType == nil {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *KalturaBeacon) HasIndexType() bool {
	if o != nil && o.IndexType != nil {
		return true
	}

	return false
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *KalturaBeacon) SetIndexType(v string) {
	o.IndexType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *KalturaBeacon) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *KalturaBeacon) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *KalturaBeacon) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPrivateData returns the PrivateData field value if set, zero value otherwise.
func (o *KalturaBeacon) GetPrivateData() string {
	if o == nil || o.PrivateData == nil {
		var ret string
		return ret
	}
	return *o.PrivateData
}

// GetPrivateDataOk returns a tuple with the PrivateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetPrivateDataOk() (*string, bool) {
	if o == nil || o.PrivateData == nil {
		return nil, false
	}
	return o.PrivateData, true
}

// HasPrivateData returns a boolean if a field has been set.
func (o *KalturaBeacon) HasPrivateData() bool {
	if o != nil && o.PrivateData != nil {
		return true
	}

	return false
}

// SetPrivateData gets a reference to the given string and assigns it to the PrivateData field.
func (o *KalturaBeacon) SetPrivateData(v string) {
	o.PrivateData = &v
}

// GetRawData returns the RawData field value if set, zero value otherwise.
func (o *KalturaBeacon) GetRawData() string {
	if o == nil || o.RawData == nil {
		var ret string
		return ret
	}
	return *o.RawData
}

// GetRawDataOk returns a tuple with the RawData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetRawDataOk() (*string, bool) {
	if o == nil || o.RawData == nil {
		return nil, false
	}
	return o.RawData, true
}

// HasRawData returns a boolean if a field has been set.
func (o *KalturaBeacon) HasRawData() bool {
	if o != nil && o.RawData != nil {
		return true
	}

	return false
}

// SetRawData gets a reference to the given string and assigns it to the RawData field.
func (o *KalturaBeacon) SetRawData(v string) {
	o.RawData = &v
}

// GetRelatedObjectType returns the RelatedObjectType field value if set, zero value otherwise.
func (o *KalturaBeacon) GetRelatedObjectType() string {
	if o == nil || o.RelatedObjectType == nil {
		var ret string
		return ret
	}
	return *o.RelatedObjectType
}

// GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetRelatedObjectTypeOk() (*string, bool) {
	if o == nil || o.RelatedObjectType == nil {
		return nil, false
	}
	return o.RelatedObjectType, true
}

// HasRelatedObjectType returns a boolean if a field has been set.
func (o *KalturaBeacon) HasRelatedObjectType() bool {
	if o != nil && o.RelatedObjectType != nil {
		return true
	}

	return false
}

// SetRelatedObjectType gets a reference to the given string and assigns it to the RelatedObjectType field.
func (o *KalturaBeacon) SetRelatedObjectType(v string) {
	o.RelatedObjectType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaBeacon) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBeacon) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaBeacon) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaBeacon) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaBeacon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IndexType != nil {
		toSerialize["indexType"] = o.IndexType
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.PrivateData != nil {
		toSerialize["privateData"] = o.PrivateData
	}
	if o.RawData != nil {
		toSerialize["rawData"] = o.RawData
	}
	if o.RelatedObjectType != nil {
		toSerialize["relatedObjectType"] = o.RelatedObjectType
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBeacon struct {
	value *KalturaBeacon
	isSet bool
}

func (v NullableKalturaBeacon) Get() *KalturaBeacon {
	return v.value
}

func (v *NullableKalturaBeacon) Set(val *KalturaBeacon) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBeacon) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBeacon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBeacon(val *KalturaBeacon) *NullableKalturaBeacon {
	return &NullableKalturaBeacon{value: val, isSet: true}
}

func (v NullableKalturaBeacon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBeacon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


