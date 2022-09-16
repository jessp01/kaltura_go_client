# KalturaEntryVendorTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | &#x60;readOnly&#x60;  Access key generated by Kaltura to allow vendors to ingest the end result to the destination | [optional] [readonly] 
**Accuracy** | Pointer to **int32** | Task result accuracy percentage | [optional] 
**CatalogItemId** | Pointer to **int32** | &#x60;insertOnly&#x60;  The catalog item Id containing the task description | [optional] 
**Context** | Pointer to **string** | Task context | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CreationMode** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryVendorTaskCreationMode&#x60;  Task creation mode | [optional] [readonly] 
**Dictionary** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EntryId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**ErrDescription** | Pointer to **string** | Err description provided by provider in case job execution has failed | [optional] 
**ExpectedFinishTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ExternalTaskId** | Pointer to **string** | The vendor&#39;s task internal Id | [optional] 
**FinishTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ModeratingUser** | Pointer to **string** | &#x60;readOnly&#x60;  The user ID that approved this task for execution (in case moderation is requested) | [optional] [readonly] 
**Notes** | Pointer to **string** | User generated notes that should be taken into account by the vendor while executing the task | [optional] 
**OutputObjectId** | Pointer to **string** | Task main object generated by executing the task | [optional] 
**PartnerData** | Pointer to **string** | Json object containing extra task data required by the requester | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Price** | Pointer to **float32** | &#x60;readOnly&#x60;  The charged price to execute this task | [optional] [readonly] 
**QueueTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ReachProfileId** | Pointer to **int32** | &#x60;insertOnly&#x60;  The profile id from which this task base config is taken from | [optional] 
**ServiceFeature** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorServiceFeature&#x60; | [optional] [readonly] 
**ServiceType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorServiceType&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaEntryVendorTaskStatus&#x60; | [optional] 
**TaskJobData** | Pointer to [**KalturaVendorTaskData**](KalturaVendorTaskData.md) |  | [optional] 
**TurnAroundTime** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorServiceTurnAroundTime&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;readOnly&#x60;  The ID of the user who created this task | [optional] [readonly] 
**VendorPartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **string** | &#x60;readOnly&#x60;  Vendor generated by Kaltura representing the entry vendor task version correlated to the entry version | [optional] [readonly] 

## Methods

### NewKalturaEntryVendorTask

`func NewKalturaEntryVendorTask() *KalturaEntryVendorTask`

NewKalturaEntryVendorTask instantiates a new KalturaEntryVendorTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEntryVendorTaskWithDefaults

`func NewKalturaEntryVendorTaskWithDefaults() *KalturaEntryVendorTask`

NewKalturaEntryVendorTaskWithDefaults instantiates a new KalturaEntryVendorTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *KalturaEntryVendorTask) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *KalturaEntryVendorTask) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *KalturaEntryVendorTask) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *KalturaEntryVendorTask) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccuracy

`func (o *KalturaEntryVendorTask) GetAccuracy() int32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *KalturaEntryVendorTask) GetAccuracyOk() (*int32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *KalturaEntryVendorTask) SetAccuracy(v int32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *KalturaEntryVendorTask) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetCatalogItemId

`func (o *KalturaEntryVendorTask) GetCatalogItemId() int32`

GetCatalogItemId returns the CatalogItemId field if non-nil, zero value otherwise.

### GetCatalogItemIdOk

`func (o *KalturaEntryVendorTask) GetCatalogItemIdOk() (*int32, bool)`

GetCatalogItemIdOk returns a tuple with the CatalogItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemId

`func (o *KalturaEntryVendorTask) SetCatalogItemId(v int32)`

SetCatalogItemId sets CatalogItemId field to given value.

### HasCatalogItemId

`func (o *KalturaEntryVendorTask) HasCatalogItemId() bool`

HasCatalogItemId returns a boolean if a field has been set.

### GetContext

`func (o *KalturaEntryVendorTask) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *KalturaEntryVendorTask) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *KalturaEntryVendorTask) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *KalturaEntryVendorTask) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaEntryVendorTask) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaEntryVendorTask) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaEntryVendorTask) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaEntryVendorTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreationMode

`func (o *KalturaEntryVendorTask) GetCreationMode() int32`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *KalturaEntryVendorTask) GetCreationModeOk() (*int32, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *KalturaEntryVendorTask) SetCreationMode(v int32)`

SetCreationMode sets CreationMode field to given value.

### HasCreationMode

`func (o *KalturaEntryVendorTask) HasCreationMode() bool`

HasCreationMode returns a boolean if a field has been set.

### GetDictionary

`func (o *KalturaEntryVendorTask) GetDictionary() string`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *KalturaEntryVendorTask) GetDictionaryOk() (*string, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *KalturaEntryVendorTask) SetDictionary(v string)`

SetDictionary sets Dictionary field to given value.

### HasDictionary

`func (o *KalturaEntryVendorTask) HasDictionary() bool`

HasDictionary returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaEntryVendorTask) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaEntryVendorTask) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaEntryVendorTask) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaEntryVendorTask) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetErrDescription

`func (o *KalturaEntryVendorTask) GetErrDescription() string`

GetErrDescription returns the ErrDescription field if non-nil, zero value otherwise.

### GetErrDescriptionOk

`func (o *KalturaEntryVendorTask) GetErrDescriptionOk() (*string, bool)`

GetErrDescriptionOk returns a tuple with the ErrDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrDescription

`func (o *KalturaEntryVendorTask) SetErrDescription(v string)`

SetErrDescription sets ErrDescription field to given value.

### HasErrDescription

`func (o *KalturaEntryVendorTask) HasErrDescription() bool`

HasErrDescription returns a boolean if a field has been set.

### GetExpectedFinishTime

`func (o *KalturaEntryVendorTask) GetExpectedFinishTime() int32`

GetExpectedFinishTime returns the ExpectedFinishTime field if non-nil, zero value otherwise.

### GetExpectedFinishTimeOk

`func (o *KalturaEntryVendorTask) GetExpectedFinishTimeOk() (*int32, bool)`

GetExpectedFinishTimeOk returns a tuple with the ExpectedFinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedFinishTime

`func (o *KalturaEntryVendorTask) SetExpectedFinishTime(v int32)`

SetExpectedFinishTime sets ExpectedFinishTime field to given value.

### HasExpectedFinishTime

`func (o *KalturaEntryVendorTask) HasExpectedFinishTime() bool`

HasExpectedFinishTime returns a boolean if a field has been set.

### GetExternalTaskId

`func (o *KalturaEntryVendorTask) GetExternalTaskId() string`

GetExternalTaskId returns the ExternalTaskId field if non-nil, zero value otherwise.

### GetExternalTaskIdOk

`func (o *KalturaEntryVendorTask) GetExternalTaskIdOk() (*string, bool)`

GetExternalTaskIdOk returns a tuple with the ExternalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTaskId

`func (o *KalturaEntryVendorTask) SetExternalTaskId(v string)`

SetExternalTaskId sets ExternalTaskId field to given value.

### HasExternalTaskId

`func (o *KalturaEntryVendorTask) HasExternalTaskId() bool`

HasExternalTaskId returns a boolean if a field has been set.

### GetFinishTime

`func (o *KalturaEntryVendorTask) GetFinishTime() int32`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *KalturaEntryVendorTask) GetFinishTimeOk() (*int32, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *KalturaEntryVendorTask) SetFinishTime(v int32)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *KalturaEntryVendorTask) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetId

`func (o *KalturaEntryVendorTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaEntryVendorTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaEntryVendorTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaEntryVendorTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModeratingUser

`func (o *KalturaEntryVendorTask) GetModeratingUser() string`

GetModeratingUser returns the ModeratingUser field if non-nil, zero value otherwise.

### GetModeratingUserOk

`func (o *KalturaEntryVendorTask) GetModeratingUserOk() (*string, bool)`

GetModeratingUserOk returns a tuple with the ModeratingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratingUser

`func (o *KalturaEntryVendorTask) SetModeratingUser(v string)`

SetModeratingUser sets ModeratingUser field to given value.

### HasModeratingUser

`func (o *KalturaEntryVendorTask) HasModeratingUser() bool`

HasModeratingUser returns a boolean if a field has been set.

### GetNotes

`func (o *KalturaEntryVendorTask) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *KalturaEntryVendorTask) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *KalturaEntryVendorTask) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *KalturaEntryVendorTask) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOutputObjectId

`func (o *KalturaEntryVendorTask) GetOutputObjectId() string`

GetOutputObjectId returns the OutputObjectId field if non-nil, zero value otherwise.

### GetOutputObjectIdOk

`func (o *KalturaEntryVendorTask) GetOutputObjectIdOk() (*string, bool)`

GetOutputObjectIdOk returns a tuple with the OutputObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputObjectId

`func (o *KalturaEntryVendorTask) SetOutputObjectId(v string)`

SetOutputObjectId sets OutputObjectId field to given value.

### HasOutputObjectId

`func (o *KalturaEntryVendorTask) HasOutputObjectId() bool`

HasOutputObjectId returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaEntryVendorTask) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaEntryVendorTask) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaEntryVendorTask) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaEntryVendorTask) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaEntryVendorTask) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaEntryVendorTask) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaEntryVendorTask) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaEntryVendorTask) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPrice

`func (o *KalturaEntryVendorTask) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *KalturaEntryVendorTask) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *KalturaEntryVendorTask) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *KalturaEntryVendorTask) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQueueTime

`func (o *KalturaEntryVendorTask) GetQueueTime() int32`

GetQueueTime returns the QueueTime field if non-nil, zero value otherwise.

### GetQueueTimeOk

`func (o *KalturaEntryVendorTask) GetQueueTimeOk() (*int32, bool)`

GetQueueTimeOk returns a tuple with the QueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueTime

`func (o *KalturaEntryVendorTask) SetQueueTime(v int32)`

SetQueueTime sets QueueTime field to given value.

### HasQueueTime

`func (o *KalturaEntryVendorTask) HasQueueTime() bool`

HasQueueTime returns a boolean if a field has been set.

### GetReachProfileId

`func (o *KalturaEntryVendorTask) GetReachProfileId() int32`

GetReachProfileId returns the ReachProfileId field if non-nil, zero value otherwise.

### GetReachProfileIdOk

`func (o *KalturaEntryVendorTask) GetReachProfileIdOk() (*int32, bool)`

GetReachProfileIdOk returns a tuple with the ReachProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachProfileId

`func (o *KalturaEntryVendorTask) SetReachProfileId(v int32)`

SetReachProfileId sets ReachProfileId field to given value.

### HasReachProfileId

`func (o *KalturaEntryVendorTask) HasReachProfileId() bool`

HasReachProfileId returns a boolean if a field has been set.

### GetServiceFeature

`func (o *KalturaEntryVendorTask) GetServiceFeature() int32`

GetServiceFeature returns the ServiceFeature field if non-nil, zero value otherwise.

### GetServiceFeatureOk

`func (o *KalturaEntryVendorTask) GetServiceFeatureOk() (*int32, bool)`

GetServiceFeatureOk returns a tuple with the ServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFeature

`func (o *KalturaEntryVendorTask) SetServiceFeature(v int32)`

SetServiceFeature sets ServiceFeature field to given value.

### HasServiceFeature

`func (o *KalturaEntryVendorTask) HasServiceFeature() bool`

HasServiceFeature returns a boolean if a field has been set.

### GetServiceType

`func (o *KalturaEntryVendorTask) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *KalturaEntryVendorTask) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *KalturaEntryVendorTask) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *KalturaEntryVendorTask) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaEntryVendorTask) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaEntryVendorTask) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaEntryVendorTask) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaEntryVendorTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskJobData

`func (o *KalturaEntryVendorTask) GetTaskJobData() KalturaVendorTaskData`

GetTaskJobData returns the TaskJobData field if non-nil, zero value otherwise.

### GetTaskJobDataOk

`func (o *KalturaEntryVendorTask) GetTaskJobDataOk() (*KalturaVendorTaskData, bool)`

GetTaskJobDataOk returns a tuple with the TaskJobData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskJobData

`func (o *KalturaEntryVendorTask) SetTaskJobData(v KalturaVendorTaskData)`

SetTaskJobData sets TaskJobData field to given value.

### HasTaskJobData

`func (o *KalturaEntryVendorTask) HasTaskJobData() bool`

HasTaskJobData returns a boolean if a field has been set.

### GetTurnAroundTime

`func (o *KalturaEntryVendorTask) GetTurnAroundTime() int32`

GetTurnAroundTime returns the TurnAroundTime field if non-nil, zero value otherwise.

### GetTurnAroundTimeOk

`func (o *KalturaEntryVendorTask) GetTurnAroundTimeOk() (*int32, bool)`

GetTurnAroundTimeOk returns a tuple with the TurnAroundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnAroundTime

`func (o *KalturaEntryVendorTask) SetTurnAroundTime(v int32)`

SetTurnAroundTime sets TurnAroundTime field to given value.

### HasTurnAroundTime

`func (o *KalturaEntryVendorTask) HasTurnAroundTime() bool`

HasTurnAroundTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaEntryVendorTask) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaEntryVendorTask) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaEntryVendorTask) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaEntryVendorTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaEntryVendorTask) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaEntryVendorTask) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaEntryVendorTask) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaEntryVendorTask) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVendorPartnerId

`func (o *KalturaEntryVendorTask) GetVendorPartnerId() int32`

GetVendorPartnerId returns the VendorPartnerId field if non-nil, zero value otherwise.

### GetVendorPartnerIdOk

`func (o *KalturaEntryVendorTask) GetVendorPartnerIdOk() (*int32, bool)`

GetVendorPartnerIdOk returns a tuple with the VendorPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartnerId

`func (o *KalturaEntryVendorTask) SetVendorPartnerId(v int32)`

SetVendorPartnerId sets VendorPartnerId field to given value.

### HasVendorPartnerId

`func (o *KalturaEntryVendorTask) HasVendorPartnerId() bool`

HasVendorPartnerId returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaEntryVendorTask) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaEntryVendorTask) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaEntryVendorTask) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaEntryVendorTask) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

