# JobsAbortJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int32** |  | 
**JobType** | **string** |  | 

## Methods

### NewJobsAbortJobRequest

`func NewJobsAbortJobRequest(jobId int32, jobType string, ) *JobsAbortJobRequest`

NewJobsAbortJobRequest instantiates a new JobsAbortJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsAbortJobRequestWithDefaults

`func NewJobsAbortJobRequestWithDefaults() *JobsAbortJobRequest`

NewJobsAbortJobRequestWithDefaults instantiates a new JobsAbortJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobsAbortJobRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobsAbortJobRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobsAbortJobRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetJobType

`func (o *JobsAbortJobRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *JobsAbortJobRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *JobsAbortJobRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


