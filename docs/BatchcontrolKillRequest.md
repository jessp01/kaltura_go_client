# BatchcontrolKillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**BatchIndex** | **int32** |  | 
**Cause** | **string** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewBatchcontrolKillRequest

`func NewBatchcontrolKillRequest(adminId int32, batchIndex int32, cause string, workerId int32, ) *BatchcontrolKillRequest`

NewBatchcontrolKillRequest instantiates a new BatchcontrolKillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolKillRequestWithDefaults

`func NewBatchcontrolKillRequestWithDefaults() *BatchcontrolKillRequest`

NewBatchcontrolKillRequestWithDefaults instantiates a new BatchcontrolKillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolKillRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolKillRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolKillRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetBatchIndex

`func (o *BatchcontrolKillRequest) GetBatchIndex() int32`

GetBatchIndex returns the BatchIndex field if non-nil, zero value otherwise.

### GetBatchIndexOk

`func (o *BatchcontrolKillRequest) GetBatchIndexOk() (*int32, bool)`

GetBatchIndexOk returns a tuple with the BatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIndex

`func (o *BatchcontrolKillRequest) SetBatchIndex(v int32)`

SetBatchIndex sets BatchIndex field to given value.


### GetCause

`func (o *BatchcontrolKillRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolKillRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolKillRequest) SetCause(v string)`

SetCause sets Cause field to given value.


### GetWorkerId

`func (o *BatchcontrolKillRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *BatchcontrolKillRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *BatchcontrolKillRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


