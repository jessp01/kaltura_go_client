# BatchcontrolStopSchedulerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **int32** |  | 
**Cause** | **string** |  | 
**SchedulerId** | **int32** |  | 

## Methods

### NewBatchcontrolStopSchedulerRequest

`func NewBatchcontrolStopSchedulerRequest(adminId int32, cause string, schedulerId int32, ) *BatchcontrolStopSchedulerRequest`

NewBatchcontrolStopSchedulerRequest instantiates a new BatchcontrolStopSchedulerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolStopSchedulerRequestWithDefaults

`func NewBatchcontrolStopSchedulerRequestWithDefaults() *BatchcontrolStopSchedulerRequest`

NewBatchcontrolStopSchedulerRequestWithDefaults instantiates a new BatchcontrolStopSchedulerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminId

`func (o *BatchcontrolStopSchedulerRequest) GetAdminId() int32`

GetAdminId returns the AdminId field if non-nil, zero value otherwise.

### GetAdminIdOk

`func (o *BatchcontrolStopSchedulerRequest) GetAdminIdOk() (*int32, bool)`

GetAdminIdOk returns a tuple with the AdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminId

`func (o *BatchcontrolStopSchedulerRequest) SetAdminId(v int32)`

SetAdminId sets AdminId field to given value.


### GetCause

`func (o *BatchcontrolStopSchedulerRequest) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *BatchcontrolStopSchedulerRequest) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *BatchcontrolStopSchedulerRequest) SetCause(v string)`

SetCause sets Cause field to given value.


### GetSchedulerId

`func (o *BatchcontrolStopSchedulerRequest) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *BatchcontrolStopSchedulerRequest) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *BatchcontrolStopSchedulerRequest) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


