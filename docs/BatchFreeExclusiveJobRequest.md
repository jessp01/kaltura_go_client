# BatchFreeExclusiveJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**JobType** | **string** |  | 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 
**ResetExecutionAttempts** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewBatchFreeExclusiveJobRequest

`func NewBatchFreeExclusiveJobRequest(id int32, jobType string, lockKey KalturaExclusiveLockKey, ) *BatchFreeExclusiveJobRequest`

NewBatchFreeExclusiveJobRequest instantiates a new BatchFreeExclusiveJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchFreeExclusiveJobRequestWithDefaults

`func NewBatchFreeExclusiveJobRequestWithDefaults() *BatchFreeExclusiveJobRequest`

NewBatchFreeExclusiveJobRequestWithDefaults instantiates a new BatchFreeExclusiveJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BatchFreeExclusiveJobRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BatchFreeExclusiveJobRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BatchFreeExclusiveJobRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetJobType

`func (o *BatchFreeExclusiveJobRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BatchFreeExclusiveJobRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BatchFreeExclusiveJobRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.


### GetLockKey

`func (o *BatchFreeExclusiveJobRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchFreeExclusiveJobRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchFreeExclusiveJobRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.


### GetResetExecutionAttempts

`func (o *BatchFreeExclusiveJobRequest) GetResetExecutionAttempts() bool`

GetResetExecutionAttempts returns the ResetExecutionAttempts field if non-nil, zero value otherwise.

### GetResetExecutionAttemptsOk

`func (o *BatchFreeExclusiveJobRequest) GetResetExecutionAttemptsOk() (*bool, bool)`

GetResetExecutionAttemptsOk returns a tuple with the ResetExecutionAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetExecutionAttempts

`func (o *BatchFreeExclusiveJobRequest) SetResetExecutionAttempts(v bool)`

SetResetExecutionAttempts sets ResetExecutionAttempts field to given value.

### HasResetExecutionAttempts

`func (o *BatchFreeExclusiveJobRequest) HasResetExecutionAttempts() bool`

HasResetExecutionAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


