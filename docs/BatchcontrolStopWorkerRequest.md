# BatchcontrolStopWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**Cause** | **string** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewBatchcontrolStopWorkerRequest

`func NewBatchcontrolStopWorkerRequest(adminId int32, cause string, workerId int32, ) *BatchcontrolStopWorkerRequest`

NewBatchcontrolStopWorkerRequest instantiates a new BatchcontrolStopWorkerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolStopWorkerRequestWithDefaults

`func NewBatchcontrolStopWorkerRequestWithDefaults() *BatchcontrolStopWorkerRequest`

NewBatchcontrolStopWorkerRequestWithDefaults instantiates a new BatchcontrolStopWorkerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolStopWorkerRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolStopWorkerRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolStopWorkerRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetCause

`func (o *BatchcontrolStopWorkerRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolStopWorkerRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolStopWorkerRequest) SetCause(v string)`

SetCause sets Cause field to given value.


### GetWorkerId

`func (o *BatchcontrolStopWorkerRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *BatchcontrolStopWorkerRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *BatchcontrolStopWorkerRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


