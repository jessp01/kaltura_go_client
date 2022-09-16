# KalturaLiveChannelSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | Pointer to **string** | Live channel id | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Segment creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | Segment description | [optional] 
**Duration** | Pointer to **float32** | Segment play duration time, in mili-seconds | [optional] 
**EntryId** | Pointer to **string** | Entry id to be played | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Unique identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Segment name | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**StartTime** | Pointer to **float32** | Segment play start time, in mili-seconds, according to trigger type | [optional] 
**Status** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaLiveChannelSegmentStatus&#x60; | [optional] [readonly] 
**Tags** | Pointer to **string** | Segment tags | [optional] 
**TriggerSegmentId** | Pointer to **int32** | Live channel segment that the trigger relates to | [optional] 
**TriggerType** | Pointer to **string** | Enum Type: &#x60;KalturaLiveChannelSegmentTriggerType&#x60;  Segment start time trigger type | [optional] 
**Type** | Pointer to **string** | Enum Type: &#x60;KalturaLiveChannelSegmentType&#x60;  Segment could be associated with the main stream, as additional stream or as overlay | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Segment update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaLiveChannelSegment

`func NewKalturaLiveChannelSegment() *KalturaLiveChannelSegment`

NewKalturaLiveChannelSegment instantiates a new KalturaLiveChannelSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLiveChannelSegmentWithDefaults

`func NewKalturaLiveChannelSegmentWithDefaults() *KalturaLiveChannelSegment`

NewKalturaLiveChannelSegmentWithDefaults instantiates a new KalturaLiveChannelSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *KalturaLiveChannelSegment) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *KalturaLiveChannelSegment) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *KalturaLiveChannelSegment) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *KalturaLiveChannelSegment) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaLiveChannelSegment) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaLiveChannelSegment) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaLiveChannelSegment) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaLiveChannelSegment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaLiveChannelSegment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaLiveChannelSegment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaLiveChannelSegment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaLiveChannelSegment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *KalturaLiveChannelSegment) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KalturaLiveChannelSegment) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KalturaLiveChannelSegment) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KalturaLiveChannelSegment) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaLiveChannelSegment) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaLiveChannelSegment) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaLiveChannelSegment) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaLiveChannelSegment) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetId

`func (o *KalturaLiveChannelSegment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaLiveChannelSegment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaLiveChannelSegment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaLiveChannelSegment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaLiveChannelSegment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaLiveChannelSegment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaLiveChannelSegment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaLiveChannelSegment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaLiveChannelSegment) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaLiveChannelSegment) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaLiveChannelSegment) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaLiveChannelSegment) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStartTime

`func (o *KalturaLiveChannelSegment) GetStartTime() float32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KalturaLiveChannelSegment) GetStartTimeOk() (*float32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KalturaLiveChannelSegment) SetStartTime(v float32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *KalturaLiveChannelSegment) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaLiveChannelSegment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaLiveChannelSegment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaLiveChannelSegment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaLiveChannelSegment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaLiveChannelSegment) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaLiveChannelSegment) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaLiveChannelSegment) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaLiveChannelSegment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTriggerSegmentId

`func (o *KalturaLiveChannelSegment) GetTriggerSegmentId() int32`

GetTriggerSegmentId returns the TriggerSegmentId field if non-nil, zero value otherwise.

### GetTriggerSegmentIdOk

`func (o *KalturaLiveChannelSegment) GetTriggerSegmentIdOk() (*int32, bool)`

GetTriggerSegmentIdOk returns a tuple with the TriggerSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSegmentId

`func (o *KalturaLiveChannelSegment) SetTriggerSegmentId(v int32)`

SetTriggerSegmentId sets TriggerSegmentId field to given value.

### HasTriggerSegmentId

`func (o *KalturaLiveChannelSegment) HasTriggerSegmentId() bool`

HasTriggerSegmentId returns a boolean if a field has been set.

### GetTriggerType

`func (o *KalturaLiveChannelSegment) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *KalturaLiveChannelSegment) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *KalturaLiveChannelSegment) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *KalturaLiveChannelSegment) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### GetType

`func (o *KalturaLiveChannelSegment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaLiveChannelSegment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaLiveChannelSegment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaLiveChannelSegment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaLiveChannelSegment) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaLiveChannelSegment) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaLiveChannelSegment) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaLiveChannelSegment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


