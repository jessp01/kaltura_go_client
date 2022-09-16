# BatchcontrolSetCommandResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | **int32** |  | 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**Status** | **int32** |  | 

## Methods

### NewBatchcontrolSetCommandResultRequest

`func NewBatchcontrolSetCommandResultRequest(commandId int32, status int32, ) *BatchcontrolSetCommandResultRequest`

NewBatchcontrolSetCommandResultRequest instantiates a new BatchcontrolSetCommandResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolSetCommandResultRequestWithDefaults

`func NewBatchcontrolSetCommandResultRequestWithDefaults() *BatchcontrolSetCommandResultRequest`

NewBatchcontrolSetCommandResultRequestWithDefaults instantiates a new BatchcontrolSetCommandResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *BatchcontrolSetCommandResultRequest) GetCommandId() int32`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *BatchcontrolSetCommandResultRequest) GetCommandIdOk() (*int32, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *BatchcontrolSetCommandResultRequest) SetCommandId(v int32)`

SetCommandId sets CommandId field to given value.


### GetErrorDescription

`func (o *BatchcontrolSetCommandResultRequest) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BatchcontrolSetCommandResultRequest) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BatchcontrolSetCommandResultRequest) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BatchcontrolSetCommandResultRequest) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetStatus

`func (o *BatchcontrolSetCommandResultRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchcontrolSetCommandResultRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchcontrolSetCommandResultRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


