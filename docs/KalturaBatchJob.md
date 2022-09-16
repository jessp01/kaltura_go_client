# KalturaBatchJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abort** | Pointer to **int32** |  | [optional] 
**BatchIndex** | Pointer to **int32** |  | [optional] 
**BatchVersion** | Pointer to **int32** |  | [optional] 
**BulkJobId** | Pointer to **int32** | The id of the bulk upload job that initiated this job | [optional] 
**CheckAgainTimeout** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Data** | Pointer to [**KalturaJobData**](KalturaJobData.md) |  | [optional] 
**Dc** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EntryName** | Pointer to **string** |  | [optional] 
**ErrNumber** | Pointer to **int32** |  | [optional] 
**ErrType** | Pointer to **int32** | Enum Type: &#x60;KalturaBatchJobErrorTypes&#x60; | [optional] 
**EstimatedEffort** | Pointer to **int32** |  | [optional] 
**ExecutionAttempts** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FinishTime** | Pointer to **int32** | The time that the job was finished or closed as failed | [optional] 
**History** | Pointer to [**[]KalturaBatchHistoryData**](KalturaBatchHistoryData.md) |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**JobObjectId** | Pointer to **string** |  | [optional] 
**JobObjectType** | Pointer to **int32** |  | [optional] 
**JobSubType** | Pointer to **int32** |  | [optional] 
**JobType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaBatchJobType&#x60; | [optional] [readonly] 
**LastSchedulerId** | Pointer to **int32** |  | [optional] 
**LastWorkerId** | Pointer to **int32** |  | [optional] 
**LockExpiration** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LockVersion** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Message** | Pointer to **string** |  | [optional] 
**ParentJobId** | Pointer to **int32** | When one job creates another - the parent should set this parentJobId to be its own id. | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Priority** | Pointer to **int32** |  | [optional] 
**QueueTime** | Pointer to **int32** | The time that the job was pulled from the queue | [optional] 
**RootJobId** | Pointer to **int32** | The id of the root parent job | [optional] 
**SchedulerId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaBatchJobStatus&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Urgency** | Pointer to **int32** |  | [optional] 
**WorkerId** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaBatchJob

`func NewKalturaBatchJob() *KalturaBatchJob`

NewKalturaBatchJob instantiates a new KalturaBatchJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBatchJobWithDefaults

`func NewKalturaBatchJobWithDefaults() *KalturaBatchJob`

NewKalturaBatchJobWithDefaults instantiates a new KalturaBatchJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbort

`func (o *KalturaBatchJob) GetAbort() int32`

GetAbort returns the Abort field if non-nil, zero value otherwise.

### GetAbortOk

`func (o *KalturaBatchJob) GetAbortOk() (*int32, bool)`

GetAbortOk returns a tuple with the Abort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbort

`func (o *KalturaBatchJob) SetAbort(v int32)`

SetAbort sets Abort field to given value.

### HasAbort

`func (o *KalturaBatchJob) HasAbort() bool`

HasAbort returns a boolean if a field has been set.

### GetBatchIndex

`func (o *KalturaBatchJob) GetBatchIndex() int32`

GetBatchIndex returns the BatchIndex field if non-nil, zero value otherwise.

### GetBatchIndexOk

`func (o *KalturaBatchJob) GetBatchIndexOk() (*int32, bool)`

GetBatchIndexOk returns a tuple with the BatchIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchIndex

`func (o *KalturaBatchJob) SetBatchIndex(v int32)`

SetBatchIndex sets BatchIndex field to given value.

### HasBatchIndex

`func (o *KalturaBatchJob) HasBatchIndex() bool`

HasBatchIndex returns a boolean if a field has been set.

### GetBatchVersion

`func (o *KalturaBatchJob) GetBatchVersion() int32`

GetBatchVersion returns the BatchVersion field if non-nil, zero value otherwise.

### GetBatchVersionOk

`func (o *KalturaBatchJob) GetBatchVersionOk() (*int32, bool)`

GetBatchVersionOk returns a tuple with the BatchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchVersion

`func (o *KalturaBatchJob) SetBatchVersion(v int32)`

SetBatchVersion sets BatchVersion field to given value.

### HasBatchVersion

`func (o *KalturaBatchJob) HasBatchVersion() bool`

HasBatchVersion returns a boolean if a field has been set.

### GetBulkJobId

`func (o *KalturaBatchJob) GetBulkJobId() int32`

GetBulkJobId returns the BulkJobId field if non-nil, zero value otherwise.

### GetBulkJobIdOk

`func (o *KalturaBatchJob) GetBulkJobIdOk() (*int32, bool)`

GetBulkJobIdOk returns a tuple with the BulkJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkJobId

`func (o *KalturaBatchJob) SetBulkJobId(v int32)`

SetBulkJobId sets BulkJobId field to given value.

### HasBulkJobId

`func (o *KalturaBatchJob) HasBulkJobId() bool`

HasBulkJobId returns a boolean if a field has been set.

### GetCheckAgainTimeout

`func (o *KalturaBatchJob) GetCheckAgainTimeout() int32`

GetCheckAgainTimeout returns the CheckAgainTimeout field if non-nil, zero value otherwise.

### GetCheckAgainTimeoutOk

`func (o *KalturaBatchJob) GetCheckAgainTimeoutOk() (*int32, bool)`

GetCheckAgainTimeoutOk returns a tuple with the CheckAgainTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAgainTimeout

`func (o *KalturaBatchJob) SetCheckAgainTimeout(v int32)`

SetCheckAgainTimeout sets CheckAgainTimeout field to given value.

### HasCheckAgainTimeout

`func (o *KalturaBatchJob) HasCheckAgainTimeout() bool`

HasCheckAgainTimeout returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaBatchJob) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaBatchJob) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaBatchJob) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaBatchJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *KalturaBatchJob) GetData() KalturaJobData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KalturaBatchJob) GetDataOk() (*KalturaJobData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KalturaBatchJob) SetData(v KalturaJobData)`

SetData sets Data field to given value.

### HasData

`func (o *KalturaBatchJob) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDc

`func (o *KalturaBatchJob) GetDc() int32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *KalturaBatchJob) GetDcOk() (*int32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *KalturaBatchJob) SetDc(v int32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *KalturaBatchJob) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDeletedAt

`func (o *KalturaBatchJob) GetDeletedAt() int32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *KalturaBatchJob) GetDeletedAtOk() (*int32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *KalturaBatchJob) SetDeletedAt(v int32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *KalturaBatchJob) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaBatchJob) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaBatchJob) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaBatchJob) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaBatchJob) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaBatchJob) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaBatchJob) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaBatchJob) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaBatchJob) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryName

`func (o *KalturaBatchJob) GetEntryName() string`

GetEntryName returns the EntryName field if non-nil, zero value otherwise.

### GetEntryNameOk

`func (o *KalturaBatchJob) GetEntryNameOk() (*string, bool)`

GetEntryNameOk returns a tuple with the EntryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryName

`func (o *KalturaBatchJob) SetEntryName(v string)`

SetEntryName sets EntryName field to given value.

### HasEntryName

`func (o *KalturaBatchJob) HasEntryName() bool`

HasEntryName returns a boolean if a field has been set.

### GetErrNumber

`func (o *KalturaBatchJob) GetErrNumber() int32`

GetErrNumber returns the ErrNumber field if non-nil, zero value otherwise.

### GetErrNumberOk

`func (o *KalturaBatchJob) GetErrNumberOk() (*int32, bool)`

GetErrNumberOk returns a tuple with the ErrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrNumber

`func (o *KalturaBatchJob) SetErrNumber(v int32)`

SetErrNumber sets ErrNumber field to given value.

### HasErrNumber

`func (o *KalturaBatchJob) HasErrNumber() bool`

HasErrNumber returns a boolean if a field has been set.

### GetErrType

`func (o *KalturaBatchJob) GetErrType() int32`

GetErrType returns the ErrType field if non-nil, zero value otherwise.

### GetErrTypeOk

`func (o *KalturaBatchJob) GetErrTypeOk() (*int32, bool)`

GetErrTypeOk returns a tuple with the ErrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrType

`func (o *KalturaBatchJob) SetErrType(v int32)`

SetErrType sets ErrType field to given value.

### HasErrType

`func (o *KalturaBatchJob) HasErrType() bool`

HasErrType returns a boolean if a field has been set.

### GetEstimatedEffort

`func (o *KalturaBatchJob) GetEstimatedEffort() int32`

GetEstimatedEffort returns the EstimatedEffort field if non-nil, zero value otherwise.

### GetEstimatedEffortOk

`func (o *KalturaBatchJob) GetEstimatedEffortOk() (*int32, bool)`

GetEstimatedEffortOk returns a tuple with the EstimatedEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEffort

`func (o *KalturaBatchJob) SetEstimatedEffort(v int32)`

SetEstimatedEffort sets EstimatedEffort field to given value.

### HasEstimatedEffort

`func (o *KalturaBatchJob) HasEstimatedEffort() bool`

HasEstimatedEffort returns a boolean if a field has been set.

### GetExecutionAttempts

`func (o *KalturaBatchJob) GetExecutionAttempts() int32`

GetExecutionAttempts returns the ExecutionAttempts field if non-nil, zero value otherwise.

### GetExecutionAttemptsOk

`func (o *KalturaBatchJob) GetExecutionAttemptsOk() (*int32, bool)`

GetExecutionAttemptsOk returns a tuple with the ExecutionAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionAttempts

`func (o *KalturaBatchJob) SetExecutionAttempts(v int32)`

SetExecutionAttempts sets ExecutionAttempts field to given value.

### HasExecutionAttempts

`func (o *KalturaBatchJob) HasExecutionAttempts() bool`

HasExecutionAttempts returns a boolean if a field has been set.

### GetFinishTime

`func (o *KalturaBatchJob) GetFinishTime() int32`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *KalturaBatchJob) GetFinishTimeOk() (*int32, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *KalturaBatchJob) SetFinishTime(v int32)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *KalturaBatchJob) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetHistory

`func (o *KalturaBatchJob) GetHistory() []KalturaBatchHistoryData`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *KalturaBatchJob) GetHistoryOk() (*[]KalturaBatchHistoryData, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *KalturaBatchJob) SetHistory(v []KalturaBatchHistoryData)`

SetHistory sets History field to given value.

### HasHistory

`func (o *KalturaBatchJob) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *KalturaBatchJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBatchJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBatchJob) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBatchJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobObjectId

`func (o *KalturaBatchJob) GetJobObjectId() string`

GetJobObjectId returns the JobObjectId field if non-nil, zero value otherwise.

### GetJobObjectIdOk

`func (o *KalturaBatchJob) GetJobObjectIdOk() (*string, bool)`

GetJobObjectIdOk returns a tuple with the JobObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobObjectId

`func (o *KalturaBatchJob) SetJobObjectId(v string)`

SetJobObjectId sets JobObjectId field to given value.

### HasJobObjectId

`func (o *KalturaBatchJob) HasJobObjectId() bool`

HasJobObjectId returns a boolean if a field has been set.

### GetJobObjectType

`func (o *KalturaBatchJob) GetJobObjectType() int32`

GetJobObjectType returns the JobObjectType field if non-nil, zero value otherwise.

### GetJobObjectTypeOk

`func (o *KalturaBatchJob) GetJobObjectTypeOk() (*int32, bool)`

GetJobObjectTypeOk returns a tuple with the JobObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobObjectType

`func (o *KalturaBatchJob) SetJobObjectType(v int32)`

SetJobObjectType sets JobObjectType field to given value.

### HasJobObjectType

`func (o *KalturaBatchJob) HasJobObjectType() bool`

HasJobObjectType returns a boolean if a field has been set.

### GetJobSubType

`func (o *KalturaBatchJob) GetJobSubType() int32`

GetJobSubType returns the JobSubType field if non-nil, zero value otherwise.

### GetJobSubTypeOk

`func (o *KalturaBatchJob) GetJobSubTypeOk() (*int32, bool)`

GetJobSubTypeOk returns a tuple with the JobSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSubType

`func (o *KalturaBatchJob) SetJobSubType(v int32)`

SetJobSubType sets JobSubType field to given value.

### HasJobSubType

`func (o *KalturaBatchJob) HasJobSubType() bool`

HasJobSubType returns a boolean if a field has been set.

### GetJobType

`func (o *KalturaBatchJob) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KalturaBatchJob) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KalturaBatchJob) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KalturaBatchJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLastSchedulerId

`func (o *KalturaBatchJob) GetLastSchedulerId() int32`

GetLastSchedulerId returns the LastSchedulerId field if non-nil, zero value otherwise.

### GetLastSchedulerIdOk

`func (o *KalturaBatchJob) GetLastSchedulerIdOk() (*int32, bool)`

GetLastSchedulerIdOk returns a tuple with the LastSchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSchedulerId

`func (o *KalturaBatchJob) SetLastSchedulerId(v int32)`

SetLastSchedulerId sets LastSchedulerId field to given value.

### HasLastSchedulerId

`func (o *KalturaBatchJob) HasLastSchedulerId() bool`

HasLastSchedulerId returns a boolean if a field has been set.

### GetLastWorkerId

`func (o *KalturaBatchJob) GetLastWorkerId() int32`

GetLastWorkerId returns the LastWorkerId field if non-nil, zero value otherwise.

### GetLastWorkerIdOk

`func (o *KalturaBatchJob) GetLastWorkerIdOk() (*int32, bool)`

GetLastWorkerIdOk returns a tuple with the LastWorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWorkerId

`func (o *KalturaBatchJob) SetLastWorkerId(v int32)`

SetLastWorkerId sets LastWorkerId field to given value.

### HasLastWorkerId

`func (o *KalturaBatchJob) HasLastWorkerId() bool`

HasLastWorkerId returns a boolean if a field has been set.

### GetLockExpiration

`func (o *KalturaBatchJob) GetLockExpiration() int32`

GetLockExpiration returns the LockExpiration field if non-nil, zero value otherwise.

### GetLockExpirationOk

`func (o *KalturaBatchJob) GetLockExpirationOk() (*int32, bool)`

GetLockExpirationOk returns a tuple with the LockExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiration

`func (o *KalturaBatchJob) SetLockExpiration(v int32)`

SetLockExpiration sets LockExpiration field to given value.

### HasLockExpiration

`func (o *KalturaBatchJob) HasLockExpiration() bool`

HasLockExpiration returns a boolean if a field has been set.

### GetLockVersion

`func (o *KalturaBatchJob) GetLockVersion() int32`

GetLockVersion returns the LockVersion field if non-nil, zero value otherwise.

### GetLockVersionOk

`func (o *KalturaBatchJob) GetLockVersionOk() (*int32, bool)`

GetLockVersionOk returns a tuple with the LockVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockVersion

`func (o *KalturaBatchJob) SetLockVersion(v int32)`

SetLockVersion sets LockVersion field to given value.

### HasLockVersion

`func (o *KalturaBatchJob) HasLockVersion() bool`

HasLockVersion returns a boolean if a field has been set.

### GetMessage

`func (o *KalturaBatchJob) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KalturaBatchJob) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KalturaBatchJob) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KalturaBatchJob) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetParentJobId

`func (o *KalturaBatchJob) GetParentJobId() int32`

GetParentJobId returns the ParentJobId field if non-nil, zero value otherwise.

### GetParentJobIdOk

`func (o *KalturaBatchJob) GetParentJobIdOk() (*int32, bool)`

GetParentJobIdOk returns a tuple with the ParentJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJobId

`func (o *KalturaBatchJob) SetParentJobId(v int32)`

SetParentJobId sets ParentJobId field to given value.

### HasParentJobId

`func (o *KalturaBatchJob) HasParentJobId() bool`

HasParentJobId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBatchJob) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBatchJob) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBatchJob) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBatchJob) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPriority

`func (o *KalturaBatchJob) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *KalturaBatchJob) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *KalturaBatchJob) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *KalturaBatchJob) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQueueTime

`func (o *KalturaBatchJob) GetQueueTime() int32`

GetQueueTime returns the QueueTime field if non-nil, zero value otherwise.

### GetQueueTimeOk

`func (o *KalturaBatchJob) GetQueueTimeOk() (*int32, bool)`

GetQueueTimeOk returns a tuple with the QueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTime

`func (o *KalturaBatchJob) SetQueueTime(v int32)`

SetQueueTime sets QueueTime field to given value.

### HasQueueTime

`func (o *KalturaBatchJob) HasQueueTime() bool`

HasQueueTime returns a boolean if a field has been set.

### GetRootJobId

`func (o *KalturaBatchJob) GetRootJobId() int32`

GetRootJobId returns the RootJobId field if non-nil, zero value otherwise.

### GetRootJobIdOk

`func (o *KalturaBatchJob) GetRootJobIdOk() (*int32, bool)`

GetRootJobIdOk returns a tuple with the RootJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootJobId

`func (o *KalturaBatchJob) SetRootJobId(v int32)`

SetRootJobId sets RootJobId field to given value.

### HasRootJobId

`func (o *KalturaBatchJob) HasRootJobId() bool`

HasRootJobId returns a boolean if a field has been set.

### GetSchedulerId

`func (o *KalturaBatchJob) GetSchedulerId() int32`

GetSchedulerId returns the SchedulerId field if non-nil, zero value otherwise.

### GetSchedulerIdOk

`func (o *KalturaBatchJob) GetSchedulerIdOk() (*int32, bool)`

GetSchedulerIdOk returns a tuple with the SchedulerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerId

`func (o *KalturaBatchJob) SetSchedulerId(v int32)`

SetSchedulerId sets SchedulerId field to given value.

### HasSchedulerId

`func (o *KalturaBatchJob) HasSchedulerId() bool`

HasSchedulerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBatchJob) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBatchJob) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBatchJob) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBatchJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBatchJob) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBatchJob) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBatchJob) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBatchJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrgency

`func (o *KalturaBatchJob) GetUrgency() int32`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *KalturaBatchJob) GetUrgencyOk() (*int32, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *KalturaBatchJob) SetUrgency(v int32)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *KalturaBatchJob) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.

### GetWorkerId

`func (o *KalturaBatchJob) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *KalturaBatchJob) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *KalturaBatchJob) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.

### HasWorkerId

`func (o *KalturaBatchJob) HasWorkerId() bool`

HasWorkerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


