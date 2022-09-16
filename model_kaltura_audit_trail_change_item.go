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

// KalturaAuditTrailChangeItem struct for KalturaAuditTrailChangeItem
type KalturaAuditTrailChangeItem struct {
	Descriptor *string `json:"descriptor,omitempty"`
	NewValue *string `json:"newValue,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	OldValue *string `json:"oldValue,omitempty"`
}

// NewKalturaAuditTrailChangeItem instantiates a new KalturaAuditTrailChangeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAuditTrailChangeItem() *KalturaAuditTrailChangeItem {
	this := KalturaAuditTrailChangeItem{}
	return &this
}

// NewKalturaAuditTrailChangeItemWithDefaults instantiates a new KalturaAuditTrailChangeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAuditTrailChangeItemWithDefaults() *KalturaAuditTrailChangeItem {
	this := KalturaAuditTrailChangeItem{}
	return &this
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *KalturaAuditTrailChangeItem) GetDescriptor() string {
	if o == nil || o.Descriptor == nil {
		var ret string
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAuditTrailChangeItem) GetDescriptorOk() (*string, bool) {
	if o == nil || o.Descriptor == nil {
		return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *KalturaAuditTrailChangeItem) HasDescriptor() bool {
	if o != nil && o.Descriptor != nil {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given string and assigns it to the Descriptor field.
func (o *KalturaAuditTrailChangeItem) SetDescriptor(v string) {
	o.Descriptor = &v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise.
func (o *KalturaAuditTrailChangeItem) GetNewValue() string {
	if o == nil || o.NewValue == nil {
		var ret string
		return ret
	}
	return *o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAuditTrailChangeItem) GetNewValueOk() (*string, bool) {
	if o == nil || o.NewValue == nil {
		return nil, false
	}
	return o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *KalturaAuditTrailChangeItem) HasNewValue() bool {
	if o != nil && o.NewValue != nil {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given string and assigns it to the NewValue field.
func (o *KalturaAuditTrailChangeItem) SetNewValue(v string) {
	o.NewValue = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaAuditTrailChangeItem) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAuditTrailChangeItem) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaAuditTrailChangeItem) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaAuditTrailChangeItem) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *KalturaAuditTrailChangeItem) GetOldValue() string {
	if o == nil || o.OldValue == nil {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAuditTrailChangeItem) GetOldValueOk() (*string, bool) {
	if o == nil || o.OldValue == nil {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *KalturaAuditTrailChangeItem) HasOldValue() bool {
	if o != nil && o.OldValue != nil {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *KalturaAuditTrailChangeItem) SetOldValue(v string) {
	o.OldValue = &v
}

func (o KalturaAuditTrailChangeItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Descriptor != nil {
		toSerialize["descriptor"] = o.Descriptor
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.OldValue != nil {
		toSerialize["oldValue"] = o.OldValue
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAuditTrailChangeItem struct {
	value *KalturaAuditTrailChangeItem
	isSet bool
}

func (v NullableKalturaAuditTrailChangeItem) Get() *KalturaAuditTrailChangeItem {
	return v.value
}

func (v *NullableKalturaAuditTrailChangeItem) Set(val *KalturaAuditTrailChangeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAuditTrailChangeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAuditTrailChangeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAuditTrailChangeItem(val *KalturaAuditTrailChangeItem) *NullableKalturaAuditTrailChangeItem {
	return &NullableKalturaAuditTrailChangeItem{value: val, isSet: true}
}

func (v NullableKalturaAuditTrailChangeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAuditTrailChangeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

