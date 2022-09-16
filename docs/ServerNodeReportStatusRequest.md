# ServerNodeReportStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** |  | 
**ServerNode** | Pointer to [**KalturaServerNode**](KalturaServerNode.md) |  | [optional] 
**ServerNodeStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewServerNodeReportStatusRequest

`func NewServerNodeReportStatusRequest(hostName string, ) *ServerNodeReportStatusRequest`

NewServerNodeReportStatusRequest instantiates a new ServerNodeReportStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerNodeReportStatusRequestWithDefaults

`func NewServerNodeReportStatusRequestWithDefaults() *ServerNodeReportStatusRequest`

NewServerNodeReportStatusRequestWithDefaults instantiates a new ServerNodeReportStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *ServerNodeReportStatusRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ServerNodeReportStatusRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ServerNodeReportStatusRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetServerNode

`func (o *ServerNodeReportStatusRequest) GetServerNode() KalturaServerNode`

GetServerNode returns the ServerNode field if non-nil, zero value otherwise.

### GetServerNodeOk

`func (o *ServerNodeReportStatusRequest) GetServerNodeOk() (*KalturaServerNode, bool)`

GetServerNodeOk returns a tuple with the ServerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNode

`func (o *ServerNodeReportStatusRequest) SetServerNode(v KalturaServerNode)`

SetServerNode sets ServerNode field to given value.

### HasServerNode

`func (o *ServerNodeReportStatusRequest) HasServerNode() bool`

HasServerNode returns a boolean if a field has been set.

### GetServerNodeStatus

`func (o *ServerNodeReportStatusRequest) GetServerNodeStatus() int32`

GetServerNodeStatus returns the ServerNodeStatus field if non-nil, zero value otherwise.

### GetServerNodeStatusOk

`func (o *ServerNodeReportStatusRequest) GetServerNodeStatusOk() (*int32, bool)`

GetServerNodeStatusOk returns a tuple with the ServerNodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNodeStatus

`func (o *ServerNodeReportStatusRequest) SetServerNodeStatus(v int32)`

SetServerNodeStatus sets ServerNodeStatus field to given value.

### HasServerNodeStatus

`func (o *ServerNodeReportStatusRequest) HasServerNodeStatus() bool`

HasServerNodeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


