# JobsRetryJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** |  | [optional] [default to false]
**JobId** | **int32** |  | 
**JobType** | **string** |  | 

## Methods

### NewJobsRetryJobRequest

`func NewJobsRetryJobRequest(jobId int32, jobType string, ) *JobsRetryJobRequest`

NewJobsRetryJobRequest instantiates a new JobsRetryJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsRetryJobRequestWithDefaults

`func NewJobsRetryJobRequestWithDefaults() *JobsRetryJobRequest`

NewJobsRetryJobRequestWithDefaults instantiates a new JobsRetryJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *JobsRetryJobRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *JobsRetryJobRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *JobsRetryJobRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *JobsRetryJobRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetJobId

`func (o *JobsRetryJobRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobsRetryJobRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobsRetryJobRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobsRetryJobRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobsRetryJobRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobsRetryJobRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


