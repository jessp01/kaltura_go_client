# BatchExtendLockExpirationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int32** |  | 
**MaxExecutionTime** | **int32** |  | 

## Methods

### NewBatchExtendLockExpirationRequest

`func NewBatchExtendLockExpirationRequest(jobId int32, maxExecutionTime int32, ) *BatchExtendLockExpirationRequest`

NewBatchExtendLockExpirationRequest instantiates a new BatchExtendLockExpirationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExtendLockExpirationRequestWithDefaults

`func NewBatchExtendLockExpirationRequestWithDefaults() *BatchExtendLockExpirationRequest`

NewBatchExtendLockExpirationRequestWithDefaults instantiates a new BatchExtendLockExpirationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *BatchExtendLockExpirationRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BatchExtendLockExpirationRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BatchExtendLockExpirationRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetMaxExecutionTime

`func (o *BatchExtendLockExpirationRequest) GetMaxExecutionTime() int32`

GetMaxExecutionTime returns the MaxExecutionTime field if non-nil, zero value otherwise.

### GetMaxExecutionTimeOk

`func (o *BatchExtendLockExpirationRequest) GetMaxExecutionTimeOk() (*int32, bool)`

GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionTime

`func (o *BatchExtendLockExpirationRequest) SetMaxExecutionTime(v int32)`

SetMaxExecutionTime sets MaxExecutionTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


