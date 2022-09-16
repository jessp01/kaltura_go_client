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

// ServerNodeReportStatusRequest struct for ServerNodeReportStatusRequest
type ServerNodeReportStatusRequest struct {
	HostName string `json:"hostName"`
	ServerNode *KalturaServerNode `json:"serverNode,omitempty"`
	ServerNodeStatus *int32 `json:"serverNodeStatus,omitempty"`
}

// NewServerNodeReportStatusRequest instantiates a new ServerNodeReportStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerNodeReportStatusRequest(hostName string) *ServerNodeReportStatusRequest {
	this := ServerNodeReportStatusRequest{}
	this.HostName = hostName
	return &this
}

// NewServerNodeReportStatusRequestWithDefaults instantiates a new ServerNodeReportStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerNodeReportStatusRequestWithDefaults() *ServerNodeReportStatusRequest {
	this := ServerNodeReportStatusRequest{}
	return &this
}

// GetHostName returns the HostName field value
func (o *ServerNodeReportStatusRequest) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ServerNodeReportStatusRequest) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ServerNodeReportStatusRequest) SetHostName(v string) {
	o.HostName = v
}

// GetServerNode returns the ServerNode field value if set, zero value otherwise.
func (o *ServerNodeReportStatusRequest) GetServerNode() KalturaServerNode {
	if o == nil || o.ServerNode == nil {
		var ret KalturaServerNode
		return ret
	}
	return *o.ServerNode
}

// GetServerNodeOk returns a tuple with the ServerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerNodeReportStatusRequest) GetServerNodeOk() (*KalturaServerNode, bool) {
	if o == nil || o.ServerNode == nil {
		return nil, false
	}
	return o.ServerNode, true
}

// HasServerNode returns a boolean if a field has been set.
func (o *ServerNodeReportStatusRequest) HasServerNode() bool {
	if o != nil && o.ServerNode != nil {
		return true
	}

	return false
}

// SetServerNode gets a reference to the given KalturaServerNode and assigns it to the ServerNode field.
func (o *ServerNodeReportStatusRequest) SetServerNode(v KalturaServerNode) {
	o.ServerNode = &v
}

// GetServerNodeStatus returns the ServerNodeStatus field value if set, zero value otherwise.
func (o *ServerNodeReportStatusRequest) GetServerNodeStatus() int32 {
	if o == nil || o.ServerNodeStatus == nil {
		var ret int32
		return ret
	}
	return *o.ServerNodeStatus
}

// GetServerNodeStatusOk returns a tuple with the ServerNodeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerNodeReportStatusRequest) GetServerNodeStatusOk() (*int32, bool) {
	if o == nil || o.ServerNodeStatus == nil {
		return nil, false
	}
	return o.ServerNodeStatus, true
}

// HasServerNodeStatus returns a boolean if a field has been set.
func (o *ServerNodeReportStatusRequest) HasServerNodeStatus() bool {
	if o != nil && o.ServerNodeStatus != nil {
		return true
	}

	return false
}

// SetServerNodeStatus gets a reference to the given int32 and assigns it to the ServerNodeStatus field.
func (o *ServerNodeReportStatusRequest) SetServerNodeStatus(v int32) {
	o.ServerNodeStatus = &v
}

func (o ServerNodeReportStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostName"] = o.HostName
	}
	if o.ServerNode != nil {
		toSerialize["serverNode"] = o.ServerNode
	}
	if o.ServerNodeStatus != nil {
		toSerialize["serverNodeStatus"] = o.ServerNodeStatus
	}
	return json.Marshal(toSerialize)
}

type NullableServerNodeReportStatusRequest struct {
	value *ServerNodeReportStatusRequest
	isSet bool
}

func (v NullableServerNodeReportStatusRequest) Get() *ServerNodeReportStatusRequest {
	return v.value
}

func (v *NullableServerNodeReportStatusRequest) Set(val *ServerNodeReportStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServerNodeReportStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServerNodeReportStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerNodeReportStatusRequest(val *ServerNodeReportStatusRequest) *NullableServerNodeReportStatusRequest {
	return &NullableServerNodeReportStatusRequest{value: val, isSet: true}
}

func (v NullableServerNodeReportStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerNodeReportStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


