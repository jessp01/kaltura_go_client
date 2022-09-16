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

// KalturaCuePoint `abstract`
type KalturaCuePoint struct {
	// `readOnly`
	CopiedFrom *string `json:"copiedFrom,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`  Enum Type: `KalturaCuePointType`
	CuePointType *string `json:"cuePointType,omitempty"`
	// `insertOnly`
	EntryId *string `json:"entryId,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	ForceStop *int32 `json:"forceStop,omitempty"`
	// `readOnly`
	Id *string `json:"id,omitempty"`
	// `readOnly`
	IntId *int32 `json:"intId,omitempty"`
	// `readOnly`
	IsMomentary *bool `json:"isMomentary,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	PartnerData *string `json:"partnerData,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	PartnerSortValue *int32 `json:"partnerSortValue,omitempty"`
	// Start time in milliseconds
	StartTime *int32 `json:"startTime,omitempty"`
	// `readOnly`  Enum Type: `KalturaCuePointStatus`
	Status *int32 `json:"status,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	Tags *string `json:"tags,omitempty"`
	ThumbOffset *int32 `json:"thumbOffset,omitempty"`
	TriggeredAt *int32 `json:"triggeredAt,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewKalturaCuePoint instantiates a new KalturaCuePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCuePoint() *KalturaCuePoint {
	this := KalturaCuePoint{}
	return &this
}

// NewKalturaCuePointWithDefaults instantiates a new KalturaCuePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCuePointWithDefaults() *KalturaCuePoint {
	this := KalturaCuePoint{}
	return &this
}

// GetCopiedFrom returns the CopiedFrom field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetCopiedFrom() string {
	if o == nil || o.CopiedFrom == nil {
		var ret string
		return ret
	}
	return *o.CopiedFrom
}

// GetCopiedFromOk returns a tuple with the CopiedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetCopiedFromOk() (*string, bool) {
	if o == nil || o.CopiedFrom == nil {
		return nil, false
	}
	return o.CopiedFrom, true
}

// HasCopiedFrom returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasCopiedFrom() bool {
	if o != nil && o.CopiedFrom != nil {
		return true
	}

	return false
}

// SetCopiedFrom gets a reference to the given string and assigns it to the CopiedFrom field.
func (o *KalturaCuePoint) SetCopiedFrom(v string) {
	o.CopiedFrom = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaCuePoint) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCuePointType returns the CuePointType field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetCuePointType() string {
	if o == nil || o.CuePointType == nil {
		var ret string
		return ret
	}
	return *o.CuePointType
}

// GetCuePointTypeOk returns a tuple with the CuePointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetCuePointTypeOk() (*string, bool) {
	if o == nil || o.CuePointType == nil {
		return nil, false
	}
	return o.CuePointType, true
}

// HasCuePointType returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasCuePointType() bool {
	if o != nil && o.CuePointType != nil {
		return true
	}

	return false
}

// SetCuePointType gets a reference to the given string and assigns it to the CuePointType field.
func (o *KalturaCuePoint) SetCuePointType(v string) {
	o.CuePointType = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaCuePoint) SetEntryId(v string) {
	o.EntryId = &v
}

// GetForceStop returns the ForceStop field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetForceStop() int32 {
	if o == nil || o.ForceStop == nil {
		var ret int32
		return ret
	}
	return *o.ForceStop
}

// GetForceStopOk returns a tuple with the ForceStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetForceStopOk() (*int32, bool) {
	if o == nil || o.ForceStop == nil {
		return nil, false
	}
	return o.ForceStop, true
}

// HasForceStop returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasForceStop() bool {
	if o != nil && o.ForceStop != nil {
		return true
	}

	return false
}

// SetForceStop gets a reference to the given int32 and assigns it to the ForceStop field.
func (o *KalturaCuePoint) SetForceStop(v int32) {
	o.ForceStop = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaCuePoint) SetId(v string) {
	o.Id = &v
}

// GetIntId returns the IntId field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetIntId() int32 {
	if o == nil || o.IntId == nil {
		var ret int32
		return ret
	}
	return *o.IntId
}

// GetIntIdOk returns a tuple with the IntId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetIntIdOk() (*int32, bool) {
	if o == nil || o.IntId == nil {
		return nil, false
	}
	return o.IntId, true
}

// HasIntId returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasIntId() bool {
	if o != nil && o.IntId != nil {
		return true
	}

	return false
}

// SetIntId gets a reference to the given int32 and assigns it to the IntId field.
func (o *KalturaCuePoint) SetIntId(v int32) {
	o.IntId = &v
}

// GetIsMomentary returns the IsMomentary field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetIsMomentary() bool {
	if o == nil || o.IsMomentary == nil {
		var ret bool
		return ret
	}
	return *o.IsMomentary
}

// GetIsMomentaryOk returns a tuple with the IsMomentary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetIsMomentaryOk() (*bool, bool) {
	if o == nil || o.IsMomentary == nil {
		return nil, false
	}
	return o.IsMomentary, true
}

// HasIsMomentary returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasIsMomentary() bool {
	if o != nil && o.IsMomentary != nil {
		return true
	}

	return false
}

// SetIsMomentary gets a reference to the given bool and assigns it to the IsMomentary field.
func (o *KalturaCuePoint) SetIsMomentary(v bool) {
	o.IsMomentary = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaCuePoint) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerData returns the PartnerData field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetPartnerData() string {
	if o == nil || o.PartnerData == nil {
		var ret string
		return ret
	}
	return *o.PartnerData
}

// GetPartnerDataOk returns a tuple with the PartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetPartnerDataOk() (*string, bool) {
	if o == nil || o.PartnerData == nil {
		return nil, false
	}
	return o.PartnerData, true
}

// HasPartnerData returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasPartnerData() bool {
	if o != nil && o.PartnerData != nil {
		return true
	}

	return false
}

// SetPartnerData gets a reference to the given string and assigns it to the PartnerData field.
func (o *KalturaCuePoint) SetPartnerData(v string) {
	o.PartnerData = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaCuePoint) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPartnerSortValue returns the PartnerSortValue field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetPartnerSortValue() int32 {
	if o == nil || o.PartnerSortValue == nil {
		var ret int32
		return ret
	}
	return *o.PartnerSortValue
}

// GetPartnerSortValueOk returns a tuple with the PartnerSortValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetPartnerSortValueOk() (*int32, bool) {
	if o == nil || o.PartnerSortValue == nil {
		return nil, false
	}
	return o.PartnerSortValue, true
}

// HasPartnerSortValue returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasPartnerSortValue() bool {
	if o != nil && o.PartnerSortValue != nil {
		return true
	}

	return false
}

// SetPartnerSortValue gets a reference to the given int32 and assigns it to the PartnerSortValue field.
func (o *KalturaCuePoint) SetPartnerSortValue(v int32) {
	o.PartnerSortValue = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetStartTime() int32 {
	if o == nil || o.StartTime == nil {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetStartTimeOk() (*int32, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *KalturaCuePoint) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaCuePoint) SetStatus(v int32) {
	o.Status = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaCuePoint) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *KalturaCuePoint) SetTags(v string) {
	o.Tags = &v
}

// GetThumbOffset returns the ThumbOffset field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetThumbOffset() int32 {
	if o == nil || o.ThumbOffset == nil {
		var ret int32
		return ret
	}
	return *o.ThumbOffset
}

// GetThumbOffsetOk returns a tuple with the ThumbOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetThumbOffsetOk() (*int32, bool) {
	if o == nil || o.ThumbOffset == nil {
		return nil, false
	}
	return o.ThumbOffset, true
}

// HasThumbOffset returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasThumbOffset() bool {
	if o != nil && o.ThumbOffset != nil {
		return true
	}

	return false
}

// SetThumbOffset gets a reference to the given int32 and assigns it to the ThumbOffset field.
func (o *KalturaCuePoint) SetThumbOffset(v int32) {
	o.ThumbOffset = &v
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetTriggeredAt() int32 {
	if o == nil || o.TriggeredAt == nil {
		var ret int32
		return ret
	}
	return *o.TriggeredAt
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetTriggeredAtOk() (*int32, bool) {
	if o == nil || o.TriggeredAt == nil {
		return nil, false
	}
	return o.TriggeredAt, true
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasTriggeredAt() bool {
	if o != nil && o.TriggeredAt != nil {
		return true
	}

	return false
}

// SetTriggeredAt gets a reference to the given int32 and assigns it to the TriggeredAt field.
func (o *KalturaCuePoint) SetTriggeredAt(v int32) {
	o.TriggeredAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaCuePoint) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *KalturaCuePoint) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCuePoint) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *KalturaCuePoint) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *KalturaCuePoint) SetUserId(v string) {
	o.UserId = &v
}

func (o KalturaCuePoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CopiedFrom != nil {
		toSerialize["copiedFrom"] = o.CopiedFrom
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CuePointType != nil {
		toSerialize["cuePointType"] = o.CuePointType
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.ForceStop != nil {
		toSerialize["forceStop"] = o.ForceStop
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IntId != nil {
		toSerialize["intId"] = o.IntId
	}
	if o.IsMomentary != nil {
		toSerialize["isMomentary"] = o.IsMomentary
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerData != nil {
		toSerialize["partnerData"] = o.PartnerData
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PartnerSortValue != nil {
		toSerialize["partnerSortValue"] = o.PartnerSortValue
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ThumbOffset != nil {
		toSerialize["thumbOffset"] = o.ThumbOffset
	}
	if o.TriggeredAt != nil {
		toSerialize["triggeredAt"] = o.TriggeredAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCuePoint struct {
	value *KalturaCuePoint
	isSet bool
}

func (v NullableKalturaCuePoint) Get() *KalturaCuePoint {
	return v.value
}

func (v *NullableKalturaCuePoint) Set(val *KalturaCuePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCuePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCuePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCuePoint(val *KalturaCuePoint) *NullableKalturaCuePoint {
	return &NullableKalturaCuePoint{value: val, isSet: true}
}

func (v NullableKalturaCuePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCuePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

