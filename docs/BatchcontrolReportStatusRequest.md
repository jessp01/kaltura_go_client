# BatchcontrolReportStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scheduler** | [**KalturaScheduler**](KalturaScheduler.md) |  | 
**SchedulerStatuses** | [**[]KalturaSchedulerStatus**](KalturaSchedulerStatus.md) |  | 
**WorkerQueueFilters** | [**[]KalturaWorkerQueueFilter**](KalturaWorkerQueueFilter.md) |  | 

## Methods

### NewBatchcontrolReportStatusRequest

`func NewBatchcontrolReportStatusRequest(scheduler KalturaScheduler, schedulerStatuses []KalturaSchedulerStatus, workerQueueFilters []KalturaWorkerQueueFilter, ) *BatchcontrolReportStatusRequest`

NewBatchcontrolReportStatusRequest instantiates a new BatchcontrolReportStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchcontrolReportStatusRequestWithDefaults

`func NewBatchcontrolReportStatusRequestWithDefaults() *BatchcontrolReportStatusRequest`

NewBatchcontrolReportStatusRequestWithDefaults instantiates a new BatchcontrolReportStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduler

`func (o *BatchcontrolReportStatusRequest) GetScheduler() KalturaScheduler`

GetScheduler returns the Scheduler field if non-nil, zero value otherwise.

### GetSchedulerOk

`func (o *BatchcontrolReportStatusRequest) GetSchedulerOk() (*KalturaScheduler, bool)`

GetSchedulerOk returns a tuple with the Scheduler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduler

`func (o *BatchcontrolReportStatusRequest) SetScheduler(v KalturaScheduler)`

SetScheduler sets Scheduler field to given value.


### GetSchedulerStatuses

`func (o *BatchcontrolReportStatusRequest) GetSchedulerStatuses() []KalturaSchedulerStatus`

GetSchedulerStatuses returns the SchedulerStatuses field if non-nil, zero value otherwise.

### GetSchedulerStatusesOk

`func (o *BatchcontrolReportStatusRequest) GetSchedulerStatusesOk() (*[]KalturaSchedulerStatus, bool)`

GetSchedulerStatusesOk returns a tuple with the SchedulerStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerStatuses

`func (o *BatchcontrolReportStatusRequest) SetSchedulerStatuses(v []KalturaSchedulerStatus)`

SetSchedulerStatuses sets SchedulerStatuses field to given value.


### GetWorkerQueueFilters

`func (o *BatchcontrolReportStatusRequest) GetWorkerQueueFilters() []KalturaWorkerQueueFilter`

GetWorkerQueueFilters returns the WorkerQueueFilters field if non-nil, zero value otherwise.

### GetWorkerQueueFiltersOk

`func (o *BatchcontrolReportStatusRequest) GetWorkerQueueFiltersOk() (*[]KalturaWorkerQueueFilter, bool)`

GetWorkerQueueFiltersOk returns a tuple with the WorkerQueueFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerQueueFilters

`func (o *BatchcontrolReportStatusRequest) SetWorkerQueueFilters(v []KalturaWorkerQueueFilter)`

SetWorkerQueueFilters sets WorkerQueueFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


