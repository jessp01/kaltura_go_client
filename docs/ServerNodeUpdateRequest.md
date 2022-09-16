# ServerNodeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerNode** | [**KalturaServerNode**](KalturaServerNode.md) |  | 
**ServerNodeId** | **int32** |  | 

## Methods

### NewServerNodeUpdateRequest

`func NewServerNodeUpdateRequest(serverNode KalturaServerNode, serverNodeId int32, ) *ServerNodeUpdateRequest`

NewServerNodeUpdateRequest instantiates a new ServerNodeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerNodeUpdateRequestWithDefaults

`func NewServerNodeUpdateRequestWithDefaults() *ServerNodeUpdateRequest`

NewServerNodeUpdateRequestWithDefaults instantiates a new ServerNodeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerNode

`func (o *ServerNodeUpdateRequest) GetServerNode() KalturaServerNode`

GetServerNode returns the ServerNode field if non-nil, zero value otherwise.

### GetServerNodeOk

`func (o *ServerNodeUpdateRequest) GetServerNodeOk() (*KalturaServerNode, bool)`

GetServerNodeOk returns a tuple with the ServerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNode

`func (o *ServerNodeUpdateRequest) SetServerNode(v KalturaServerNode)`

SetServerNode sets ServerNode field to given value.


### GetServerNodeId

`func (o *ServerNodeUpdateRequest) GetServerNodeId() int32`

GetServerNodeId returns the ServerNodeId field if non-nil, zero value otherwise.

### GetServerNodeIdOk

`func (o *ServerNodeUpdateRequest) GetServerNodeIdOk() (*int32, bool)`

GetServerNodeIdOk returns a tuple with the ServerNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeId

`func (o *ServerNodeUpdateRequest) SetServerNodeId(v int32)`

SetServerNodeId sets ServerNodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


