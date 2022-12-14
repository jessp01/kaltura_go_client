/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KalturaBatchJob struct for KalturaBatchJob
type KalturaBatchJob struct {
	Abort *int32 `json:"abort,omitempty"`
	BatchIndex *int32 `json:"batchIndex,omitempty"`
	BatchVersion *int32 `json:"batchVersion,omitempty"`
	// The id of the bulk upload job that initiated this job
	BulkJobId *int32 `json:"bulkJobId,omitempty"`
	CheckAgainTimeout *int32 `json:"checkAgainTimeout,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	Data *KalturaJobData `json:"data,omitempty"`
	Dc *int32 `json:"dc,omitempty"`
	// `readOnly`
	DeletedAt *int32 `json:"deletedAt,omitempty"`
	Description *string `json:"description,omitempty"`
	EntryId *string `json:"entryId,omitempty"`
	EntryName *string `json:"entryName,omitempty"`
	ErrNumber *int32 `json:"errNumber,omitempty"`
	// Enum Type: `KalturaBatchJobErrorTypes`
	ErrType *int32 `json:"errType,omitempty"`
	EstimatedEffort *int32 `json:"estimatedEffort,omitempty"`
	// `readOnly`
	ExecutionAttempts *int32 `json:"executionAttempts,omitempty"`
	// The time that the job was finished or closed as failed
	FinishTime *int32 `json:"finishTime,omitempty"`
	History []KalturaBatchHistoryData `json:"history,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	JobObjectId *string `json:"jobObjectId,omitempty"`
	JobObjectType *int32 `json:"jobObjectType,omitempty"`
	JobSubType *int32 `json:"jobSubType,omitempty"`
	// `readOnly`  Enum Type: `KalturaBatchJobType`
	JobType *string `json:"jobType,omitempty"`
	LastSchedulerId *int32 `json:"lastSchedulerId,omitempty"`
	LastWorkerId *int32 `json:"lastWorkerId,omitempty"`
	// `readOnly`
	LockExpiration *int32 `json:"lockExpiration,omitempty"`
	// `readOnly`
	LockVersion *int32 `json:"lockVersion,omitempty"`
	Message *string `json:"message,omitempty"`
	// When one job creates another - the parent should set this parentJobId to be its own id.
	ParentJobId *int32 `json:"parentJobId,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	// The time that the job was pulled from the queue
	QueueTime *int32 `json:"queueTime,omitempty"`
	// The id of the root parent job
	RootJobId *int32 `json:"rootJobId,omitempty"`
	SchedulerId *int32 `json:"schedulerId,omitempty"`
	// Enum Type: `KalturaBatchJobStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	Urgency *int32 `json:"urgency,omitempty"`
	WorkerId *int32 `json:"workerId,omitempty"`
}

// NewKalturaBatchJob instantiates a new KalturaBatchJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBatchJob() *KalturaBatchJob {
	this := KalturaBatchJob{}
	return &this
}

// NewKalturaBatchJobWithDefaults instantiates a new KalturaBatchJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBatchJobWithDefaults() *KalturaBatchJob {
	this := KalturaBatchJob{}
	return &this
}

// GetAbort returns the Abort field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetAbort() int32 {
	if o == nil || o.Abort == nil {
		var ret int32
		return ret
	}
	return *o.Abort
}

// GetAbortOk returns a tuple with the Abort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetAbortOk() (*int32, bool) {
	if o == nil || o.Abort == nil {
		return nil, false
	}
	return o.Abort, true
}

// HasAbort returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasAbort() bool {
	if o != nil && o.Abort != nil {
		return true
	}

	return false
}

// SetAbort gets a reference to the given int32 and assigns it to the Abort field.
func (o *KalturaBatchJob) SetAbort(v int32) {
	o.Abort = &v
}

// GetBatchIndex returns the BatchIndex field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetBatchIndex() int32 {
	if o == nil || o.BatchIndex == nil {
		var ret int32
		return ret
	}
	return *o.BatchIndex
}

// GetBatchIndexOk returns a tuple with the BatchIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetBatchIndexOk() (*int32, bool) {
	if o == nil || o.BatchIndex == nil {
		return nil, false
	}
	return o.BatchIndex, true
}

// HasBatchIndex returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasBatchIndex() bool {
	if o != nil && o.BatchIndex != nil {
		return true
	}

	return false
}

// SetBatchIndex gets a reference to the given int32 and assigns it to the BatchIndex field.
func (o *KalturaBatchJob) SetBatchIndex(v int32) {
	o.BatchIndex = &v
}

// GetBatchVersion returns the BatchVersion field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetBatchVersion() int32 {
	if o == nil || o.BatchVersion == nil {
		var ret int32
		return ret
	}
	return *o.BatchVersion
}

// GetBatchVersionOk returns a tuple with the BatchVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetBatchVersionOk() (*int32, bool) {
	if o == nil || o.BatchVersion == nil {
		return nil, false
	}
	return o.BatchVersion, true
}

// HasBatchVersion returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasBatchVersion() bool {
	if o != nil && o.BatchVersion != nil {
		return true
	}

	return false
}

// SetBatchVersion gets a reference to the given int32 and assigns it to the BatchVersion field.
func (o *KalturaBatchJob) SetBatchVersion(v int32) {
	o.BatchVersion = &v
}

// GetBulkJobId returns the BulkJobId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetBulkJobId() int32 {
	if o == nil || o.BulkJobId == nil {
		var ret int32
		return ret
	}
	return *o.BulkJobId
}

// GetBulkJobIdOk returns a tuple with the BulkJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetBulkJobIdOk() (*int32, bool) {
	if o == nil || o.BulkJobId == nil {
		return nil, false
	}
	return o.BulkJobId, true
}

// HasBulkJobId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasBulkJobId() bool {
	if o != nil && o.BulkJobId != nil {
		return true
	}

	return false
}

// SetBulkJobId gets a reference to the given int32 and assigns it to the BulkJobId field.
func (o *KalturaBatchJob) SetBulkJobId(v int32) {
	o.BulkJobId = &v
}

// GetCheckAgainTimeout returns the CheckAgainTimeout field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetCheckAgainTimeout() int32 {
	if o == nil || o.CheckAgainTimeout == nil {
		var ret int32
		return ret
	}
	return *o.CheckAgainTimeout
}

// GetCheckAgainTimeoutOk returns a tuple with the CheckAgainTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetCheckAgainTimeoutOk() (*int32, bool) {
	if o == nil || o.CheckAgainTimeout == nil {
		return nil, false
	}
	return o.CheckAgainTimeout, true
}

// HasCheckAgainTimeout returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasCheckAgainTimeout() bool {
	if o != nil && o.CheckAgainTimeout != nil {
		return true
	}

	return false
}

// SetCheckAgainTimeout gets a reference to the given int32 and assigns it to the CheckAgainTimeout field.
func (o *KalturaBatchJob) SetCheckAgainTimeout(v int32) {
	o.CheckAgainTimeout = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaBatchJob) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetData() KalturaJobData {
	if o == nil || o.Data == nil {
		var ret KalturaJobData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetDataOk() (*KalturaJobData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given KalturaJobData and assigns it to the Data field.
func (o *KalturaBatchJob) SetData(v KalturaJobData) {
	o.Data = &v
}

// GetDc returns the Dc field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetDc() int32 {
	if o == nil || o.Dc == nil {
		var ret int32
		return ret
	}
	return *o.Dc
}

// GetDcOk returns a tuple with the Dc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetDcOk() (*int32, bool) {
	if o == nil || o.Dc == nil {
		return nil, false
	}
	return o.Dc, true
}

// HasDc returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasDc() bool {
	if o != nil && o.Dc != nil {
		return true
	}

	return false
}

// SetDc gets a reference to the given int32 and assigns it to the Dc field.
func (o *KalturaBatchJob) SetDc(v int32) {
	o.Dc = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetDeletedAt() int32 {
	if o == nil || o.DeletedAt == nil {
		var ret int32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetDeletedAtOk() (*int32, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given int32 and assigns it to the DeletedAt field.
func (o *KalturaBatchJob) SetDeletedAt(v int32) {
	o.DeletedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaBatchJob) SetDescription(v string) {
	o.Description = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaBatchJob) SetEntryId(v string) {
	o.EntryId = &v
}

// GetEntryName returns the EntryName field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetEntryName() string {
	if o == nil || o.EntryName == nil {
		var ret string
		return ret
	}
	return *o.EntryName
}

// GetEntryNameOk returns a tuple with the EntryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetEntryNameOk() (*string, bool) {
	if o == nil || o.EntryName == nil {
		return nil, false
	}
	return o.EntryName, true
}

// HasEntryName returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasEntryName() bool {
	if o != nil && o.EntryName != nil {
		return true
	}

	return false
}

// SetEntryName gets a reference to the given string and assigns it to the EntryName field.
func (o *KalturaBatchJob) SetEntryName(v string) {
	o.EntryName = &v
}

// GetErrNumber returns the ErrNumber field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetErrNumber() int32 {
	if o == nil || o.ErrNumber == nil {
		var ret int32
		return ret
	}
	return *o.ErrNumber
}

// GetErrNumberOk returns a tuple with the ErrNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetErrNumberOk() (*int32, bool) {
	if o == nil || o.ErrNumber == nil {
		return nil, false
	}
	return o.ErrNumber, true
}

// HasErrNumber returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasErrNumber() bool {
	if o != nil && o.ErrNumber != nil {
		return true
	}

	return false
}

// SetErrNumber gets a reference to the given int32 and assigns it to the ErrNumber field.
func (o *KalturaBatchJob) SetErrNumber(v int32) {
	o.ErrNumber = &v
}

// GetErrType returns the ErrType field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetErrType() int32 {
	if o == nil || o.ErrType == nil {
		var ret int32
		return ret
	}
	return *o.ErrType
}

// GetErrTypeOk returns a tuple with the ErrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetErrTypeOk() (*int32, bool) {
	if o == nil || o.ErrType == nil {
		return nil, false
	}
	return o.ErrType, true
}

// HasErrType returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasErrType() bool {
	if o != nil && o.ErrType != nil {
		return true
	}

	return false
}

// SetErrType gets a reference to the given int32 and assigns it to the ErrType field.
func (o *KalturaBatchJob) SetErrType(v int32) {
	o.ErrType = &v
}

// GetEstimatedEffort returns the EstimatedEffort field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetEstimatedEffort() int32 {
	if o == nil || o.EstimatedEffort == nil {
		var ret int32
		return ret
	}
	return *o.EstimatedEffort
}

// GetEstimatedEffortOk returns a tuple with the EstimatedEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetEstimatedEffortOk() (*int32, bool) {
	if o == nil || o.EstimatedEffort == nil {
		return nil, false
	}
	return o.EstimatedEffort, true
}

// HasEstimatedEffort returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasEstimatedEffort() bool {
	if o != nil && o.EstimatedEffort != nil {
		return true
	}

	return false
}

// SetEstimatedEffort gets a reference to the given int32 and assigns it to the EstimatedEffort field.
func (o *KalturaBatchJob) SetEstimatedEffort(v int32) {
	o.EstimatedEffort = &v
}

// GetExecutionAttempts returns the ExecutionAttempts field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetExecutionAttempts() int32 {
	if o == nil || o.ExecutionAttempts == nil {
		var ret int32
		return ret
	}
	return *o.ExecutionAttempts
}

// GetExecutionAttemptsOk returns a tuple with the ExecutionAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetExecutionAttemptsOk() (*int32, bool) {
	if o == nil || o.ExecutionAttempts == nil {
		return nil, false
	}
	return o.ExecutionAttempts, true
}

// HasExecutionAttempts returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasExecutionAttempts() bool {
	if o != nil && o.ExecutionAttempts != nil {
		return true
	}

	return false
}

// SetExecutionAttempts gets a reference to the given int32 and assigns it to the ExecutionAttempts field.
func (o *KalturaBatchJob) SetExecutionAttempts(v int32) {
	o.ExecutionAttempts = &v
}

// GetFinishTime returns the FinishTime field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetFinishTime() int32 {
	if o == nil || o.FinishTime == nil {
		var ret int32
		return ret
	}
	return *o.FinishTime
}

// GetFinishTimeOk returns a tuple with the FinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetFinishTimeOk() (*int32, bool) {
	if o == nil || o.FinishTime == nil {
		return nil, false
	}
	return o.FinishTime, true
}

// HasFinishTime returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasFinishTime() bool {
	if o != nil && o.FinishTime != nil {
		return true
	}

	return false
}

// SetFinishTime gets a reference to the given int32 and assigns it to the FinishTime field.
func (o *KalturaBatchJob) SetFinishTime(v int32) {
	o.FinishTime = &v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetHistory() []KalturaBatchHistoryData {
	if o == nil || o.History == nil {
		var ret []KalturaBatchHistoryData
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetHistoryOk() ([]KalturaBatchHistoryData, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []KalturaBatchHistoryData and assigns it to the History field.
func (o *KalturaBatchJob) SetHistory(v []KalturaBatchHistoryData) {
	o.History = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaBatchJob) SetId(v int32) {
	o.Id = &v
}

// GetJobObjectId returns the JobObjectId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetJobObjectId() string {
	if o == nil || o.JobObjectId == nil {
		var ret string
		return ret
	}
	return *o.JobObjectId
}

// GetJobObjectIdOk returns a tuple with the JobObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetJobObjectIdOk() (*string, bool) {
	if o == nil || o.JobObjectId == nil {
		return nil, false
	}
	return o.JobObjectId, true
}

// HasJobObjectId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasJobObjectId() bool {
	if o != nil && o.JobObjectId != nil {
		return true
	}

	return false
}

// SetJobObjectId gets a reference to the given string and assigns it to the JobObjectId field.
func (o *KalturaBatchJob) SetJobObjectId(v string) {
	o.JobObjectId = &v
}

// GetJobObjectType returns the JobObjectType field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetJobObjectType() int32 {
	if o == nil || o.JobObjectType == nil {
		var ret int32
		return ret
	}
	return *o.JobObjectType
}

// GetJobObjectTypeOk returns a tuple with the JobObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetJobObjectTypeOk() (*int32, bool) {
	if o == nil || o.JobObjectType == nil {
		return nil, false
	}
	return o.JobObjectType, true
}

// HasJobObjectType returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasJobObjectType() bool {
	if o != nil && o.JobObjectType != nil {
		return true
	}

	return false
}

// SetJobObjectType gets a reference to the given int32 and assigns it to the JobObjectType field.
func (o *KalturaBatchJob) SetJobObjectType(v int32) {
	o.JobObjectType = &v
}

// GetJobSubType returns the JobSubType field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetJobSubType() int32 {
	if o == nil || o.JobSubType == nil {
		var ret int32
		return ret
	}
	return *o.JobSubType
}

// GetJobSubTypeOk returns a tuple with the JobSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetJobSubTypeOk() (*int32, bool) {
	if o == nil || o.JobSubType == nil {
		return nil, false
	}
	return o.JobSubType, true
}

// HasJobSubType returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasJobSubType() bool {
	if o != nil && o.JobSubType != nil {
		return true
	}

	return false
}

// SetJobSubType gets a reference to the given int32 and assigns it to the JobSubType field.
func (o *KalturaBatchJob) SetJobSubType(v int32) {
	o.JobSubType = &v
}

// GetJobType returns the JobType field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetJobType() string {
	if o == nil || o.JobType == nil {
		var ret string
		return ret
	}
	return *o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetJobTypeOk() (*string, bool) {
	if o == nil || o.JobType == nil {
		return nil, false
	}
	return o.JobType, true
}

// HasJobType returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasJobType() bool {
	if o != nil && o.JobType != nil {
		return true
	}

	return false
}

// SetJobType gets a reference to the given string and assigns it to the JobType field.
func (o *KalturaBatchJob) SetJobType(v string) {
	o.JobType = &v
}

// GetLastSchedulerId returns the LastSchedulerId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetLastSchedulerId() int32 {
	if o == nil || o.LastSchedulerId == nil {
		var ret int32
		return ret
	}
	return *o.LastSchedulerId
}

// GetLastSchedulerIdOk returns a tuple with the LastSchedulerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetLastSchedulerIdOk() (*int32, bool) {
	if o == nil || o.LastSchedulerId == nil {
		return nil, false
	}
	return o.LastSchedulerId, true
}

// HasLastSchedulerId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasLastSchedulerId() bool {
	if o != nil && o.LastSchedulerId != nil {
		return true
	}

	return false
}

// SetLastSchedulerId gets a reference to the given int32 and assigns it to the LastSchedulerId field.
func (o *KalturaBatchJob) SetLastSchedulerId(v int32) {
	o.LastSchedulerId = &v
}

// GetLastWorkerId returns the LastWorkerId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetLastWorkerId() int32 {
	if o == nil || o.LastWorkerId == nil {
		var ret int32
		return ret
	}
	return *o.LastWorkerId
}

// GetLastWorkerIdOk returns a tuple with the LastWorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetLastWorkerIdOk() (*int32, bool) {
	if o == nil || o.LastWorkerId == nil {
		return nil, false
	}
	return o.LastWorkerId, true
}

// HasLastWorkerId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasLastWorkerId() bool {
	if o != nil && o.LastWorkerId != nil {
		return true
	}

	return false
}

// SetLastWorkerId gets a reference to the given int32 and assigns it to the LastWorkerId field.
func (o *KalturaBatchJob) SetLastWorkerId(v int32) {
	o.LastWorkerId = &v
}

// GetLockExpiration returns the LockExpiration field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetLockExpiration() int32 {
	if o == nil || o.LockExpiration == nil {
		var ret int32
		return ret
	}
	return *o.LockExpiration
}

// GetLockExpirationOk returns a tuple with the LockExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetLockExpirationOk() (*int32, bool) {
	if o == nil || o.LockExpiration == nil {
		return nil, false
	}
	return o.LockExpiration, true
}

// HasLockExpiration returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasLockExpiration() bool {
	if o != nil && o.LockExpiration != nil {
		return true
	}

	return false
}

// SetLockExpiration gets a reference to the given int32 and assigns it to the LockExpiration field.
func (o *KalturaBatchJob) SetLockExpiration(v int32) {
	o.LockExpiration = &v
}

// GetLockVersion returns the LockVersion field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetLockVersion() int32 {
	if o == nil || o.LockVersion == nil {
		var ret int32
		return ret
	}
	return *o.LockVersion
}

// GetLockVersionOk returns a tuple with the LockVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetLockVersionOk() (*int32, bool) {
	if o == nil || o.LockVersion == nil {
		return nil, false
	}
	return o.LockVersion, true
}

// HasLockVersion returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasLockVersion() bool {
	if o != nil && o.LockVersion != nil {
		return true
	}

	return false
}

// SetLockVersion gets a reference to the given int32 and assigns it to the LockVersion field.
func (o *KalturaBatchJob) SetLockVersion(v int32) {
	o.LockVersion = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KalturaBatchJob) SetMessage(v string) {
	o.Message = &v
}

// GetParentJobId returns the ParentJobId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetParentJobId() int32 {
	if o == nil || o.ParentJobId == nil {
		var ret int32
		return ret
	}
	return *o.ParentJobId
}

// GetParentJobIdOk returns a tuple with the ParentJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetParentJobIdOk() (*int32, bool) {
	if o == nil || o.ParentJobId == nil {
		return nil, false
	}
	return o.ParentJobId, true
}

// HasParentJobId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasParentJobId() bool {
	if o != nil && o.ParentJobId != nil {
		return true
	}

	return false
}

// SetParentJobId gets a reference to the given int32 and assigns it to the ParentJobId field.
func (o *KalturaBatchJob) SetParentJobId(v int32) {
	o.ParentJobId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaBatchJob) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *KalturaBatchJob) SetPriority(v int32) {
	o.Priority = &v
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetQueueTime() int32 {
	if o == nil || o.QueueTime == nil {
		var ret int32
		return ret
	}
	return *o.QueueTime
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetQueueTimeOk() (*int32, bool) {
	if o == nil || o.QueueTime == nil {
		return nil, false
	}
	return o.QueueTime, true
}

// HasQueueTime returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasQueueTime() bool {
	if o != nil && o.QueueTime != nil {
		return true
	}

	return false
}

// SetQueueTime gets a reference to the given int32 and assigns it to the QueueTime field.
func (o *KalturaBatchJob) SetQueueTime(v int32) {
	o.QueueTime = &v
}

// GetRootJobId returns the RootJobId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetRootJobId() int32 {
	if o == nil || o.RootJobId == nil {
		var ret int32
		return ret
	}
	return *o.RootJobId
}

// GetRootJobIdOk returns a tuple with the RootJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetRootJobIdOk() (*int32, bool) {
	if o == nil || o.RootJobId == nil {
		return nil, false
	}
	return o.RootJobId, true
}

// HasRootJobId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasRootJobId() bool {
	if o != nil && o.RootJobId != nil {
		return true
	}

	return false
}

// SetRootJobId gets a reference to the given int32 and assigns it to the RootJobId field.
func (o *KalturaBatchJob) SetRootJobId(v int32) {
	o.RootJobId = &v
}

// GetSchedulerId returns the SchedulerId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetSchedulerId() int32 {
	if o == nil || o.SchedulerId == nil {
		var ret int32
		return ret
	}
	return *o.SchedulerId
}

// GetSchedulerIdOk returns a tuple with the SchedulerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetSchedulerIdOk() (*int32, bool) {
	if o == nil || o.SchedulerId == nil {
		return nil, false
	}
	return o.SchedulerId, true
}

// HasSchedulerId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasSchedulerId() bool {
	if o != nil && o.SchedulerId != nil {
		return true
	}

	return false
}

// SetSchedulerId gets a reference to the given int32 and assigns it to the SchedulerId field.
func (o *KalturaBatchJob) SetSchedulerId(v int32) {
	o.SchedulerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaBatchJob) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaBatchJob) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetUrgency() int32 {
	if o == nil || o.Urgency == nil {
		var ret int32
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetUrgencyOk() (*int32, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasUrgency() bool {
	if o != nil && o.Urgency != nil {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given int32 and assigns it to the Urgency field.
func (o *KalturaBatchJob) SetUrgency(v int32) {
	o.Urgency = &v
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise.
func (o *KalturaBatchJob) GetWorkerId() int32 {
	if o == nil || o.WorkerId == nil {
		var ret int32
		return ret
	}
	return *o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaBatchJob) GetWorkerIdOk() (*int32, bool) {
	if o == nil || o.WorkerId == nil {
		return nil, false
	}
	return o.WorkerId, true
}

// HasWorkerId returns a boolean if a field has been set.
func (o *KalturaBatchJob) HasWorkerId() bool {
	if o != nil && o.WorkerId != nil {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given int32 and assigns it to the WorkerId field.
func (o *KalturaBatchJob) SetWorkerId(v int32) {
	o.WorkerId = &v
}

func (o KalturaBatchJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Abort != nil {
		toSerialize["abort"] = o.Abort
	}
	if o.BatchIndex != nil {
		toSerialize["batchIndex"] = o.BatchIndex
	}
	if o.BatchVersion != nil {
		toSerialize["batchVersion"] = o.BatchVersion
	}
	if o.BulkJobId != nil {
		toSerialize["bulkJobId"] = o.BulkJobId
	}
	if o.CheckAgainTimeout != nil {
		toSerialize["checkAgainTimeout"] = o.CheckAgainTimeout
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Dc != nil {
		toSerialize["dc"] = o.Dc
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.EntryName != nil {
		toSerialize["entryName"] = o.EntryName
	}
	if o.ErrNumber != nil {
		toSerialize["errNumber"] = o.ErrNumber
	}
	if o.ErrType != nil {
		toSerialize["errType"] = o.ErrType
	}
	if o.EstimatedEffort != nil {
		toSerialize["estimatedEffort"] = o.EstimatedEffort
	}
	if o.ExecutionAttempts != nil {
		toSerialize["executionAttempts"] = o.ExecutionAttempts
	}
	if o.FinishTime != nil {
		toSerialize["finishTime"] = o.FinishTime
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.JobObjectId != nil {
		toSerialize["jobObjectId"] = o.JobObjectId
	}
	if o.JobObjectType != nil {
		toSerialize["jobObjectType"] = o.JobObjectType
	}
	if o.JobSubType != nil {
		toSerialize["jobSubType"] = o.JobSubType
	}
	if o.JobType != nil {
		toSerialize["jobType"] = o.JobType
	}
	if o.LastSchedulerId != nil {
		toSerialize["lastSchedulerId"] = o.LastSchedulerId
	}
	if o.LastWorkerId != nil {
		toSerialize["lastWorkerId"] = o.LastWorkerId
	}
	if o.LockExpiration != nil {
		toSerialize["lockExpiration"] = o.LockExpiration
	}
	if o.LockVersion != nil {
		toSerialize["lockVersion"] = o.LockVersion
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ParentJobId != nil {
		toSerialize["parentJobId"] = o.ParentJobId
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.QueueTime != nil {
		toSerialize["queueTime"] = o.QueueTime
	}
	if o.RootJobId != nil {
		toSerialize["rootJobId"] = o.RootJobId
	}
	if o.SchedulerId != nil {
		toSerialize["schedulerId"] = o.SchedulerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}
	if o.WorkerId != nil {
		toSerialize["workerId"] = o.WorkerId
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBatchJob struct {
	value *KalturaBatchJob
	isSet bool
}

func (v NullableKalturaBatchJob) Get() *KalturaBatchJob {
	return v.value
}

func (v *NullableKalturaBatchJob) Set(val *KalturaBatchJob) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBatchJob) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBatchJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBatchJob(val *KalturaBatchJob) *NullableKalturaBatchJob {
	return &NullableKalturaBatchJob{value: val, isSet: true}
}

func (v NullableKalturaBatchJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBatchJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


