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

// KalturaEntryServerNode `abstract`
type KalturaEntryServerNode struct {
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`
	EntryId *string `json:"entryId,omitempty"`
	// `readOnly`  unique auto-generated identifier
	Id *int32 `json:"id,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`
	ServerNodeId *int32 `json:"serverNodeId,omitempty"`
	// `readOnly`  Enum Type: `KalturaEntryServerNodeType`
	ServerType *string `json:"serverType,omitempty"`
	// `readOnly`  Enum Type: `KalturaEntryServerNodeStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaEntryServerNode instantiates a new KalturaEntryServerNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryServerNode() *KalturaEntryServerNode {
	this := KalturaEntryServerNode{}
	return &this
}

// NewKalturaEntryServerNodeWithDefaults instantiates a new KalturaEntryServerNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryServerNodeWithDefaults() *KalturaEntryServerNode {
	this := KalturaEntryServerNode{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaEntryServerNode) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaEntryServerNode) SetEntryId(v string) {
	o.EntryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaEntryServerNode) SetId(v int32) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaEntryServerNode) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaEntryServerNode) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetServerNodeId returns the ServerNodeId field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetServerNodeId() int32 {
	if o == nil || o.ServerNodeId == nil {
		var ret int32
		return ret
	}
	return *o.ServerNodeId
}

// GetServerNodeIdOk returns a tuple with the ServerNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetServerNodeIdOk() (*int32, bool) {
	if o == nil || o.ServerNodeId == nil {
		return nil, false
	}
	return o.ServerNodeId, true
}

// HasServerNodeId returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasServerNodeId() bool {
	if o != nil && o.ServerNodeId != nil {
		return true
	}

	return false
}

// SetServerNodeId gets a reference to the given int32 and assigns it to the ServerNodeId field.
func (o *KalturaEntryServerNode) SetServerNodeId(v int32) {
	o.ServerNodeId = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetServerType() string {
	if o == nil || o.ServerType == nil {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetServerTypeOk() (*string, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *KalturaEntryServerNode) SetServerType(v string) {
	o.ServerType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaEntryServerNode) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaEntryServerNode) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryServerNode) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaEntryServerNode) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaEntryServerNode) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaEntryServerNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.ServerNodeId != nil {
		toSerialize["serverNodeId"] = o.ServerNodeId
	}
	if o.ServerType != nil {
		toSerialize["serverType"] = o.ServerType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryServerNode struct {
	value *KalturaEntryServerNode
	isSet bool
}

func (v NullableKalturaEntryServerNode) Get() *KalturaEntryServerNode {
	return v.value
}

func (v *NullableKalturaEntryServerNode) Set(val *KalturaEntryServerNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryServerNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryServerNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryServerNode(val *KalturaEntryServerNode) *NullableKalturaEntryServerNode {
	return &NullableKalturaEntryServerNode{value: val, isSet: true}
}

func (v NullableKalturaEntryServerNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryServerNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


