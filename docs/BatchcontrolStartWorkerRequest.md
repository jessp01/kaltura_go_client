# BatchcontrolStartWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**Cause** | Pointer to **string** |  | [optional] 
**WorkerId** | **int32** |  | 

## Methods

### NewBatchcontrolStartWorkerRequest

`func NewBatchcontrolStartWorkerRequest(adminId int32, workerId int32, ) *BatchcontrolStartWorkerRequest`

NewBatchcontrolStartWorkerRequest instantiates a new BatchcontrolStartWorkerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolStartWorkerRequestWithDefaults

`func NewBatchcontrolStartWorkerRequestWithDefaults() *BatchcontrolStartWorkerRequest`

NewBatchcontrolStartWorkerRequestWithDefaults instantiates a new BatchcontrolStartWorkerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolStartWorkerRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolStartWorkerRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolStartWorkerRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetCause

`func (o *BatchcontrolStartWorkerRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolStartWorkerRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolStartWorkerRequest) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *BatchcontrolStartWorkerRequest) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetWorkerId

`func (o *BatchcontrolStartWorkerRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *BatchcontrolStartWorkerRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *BatchcontrolStartWorkerRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


