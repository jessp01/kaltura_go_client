# KalturaScheduleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassificationType** | Pointer to **int32** | Enum Type: &#x60;KalturaScheduleEventClassificationType&#x60; | [optional] 
**Comment** | Pointer to **string** | Specifies non-processing information intended to provide a comment to the calendar user. | [optional] 
**Contact** | Pointer to **string** | Used to represent contact information or alternately a reference to contact information. | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | Duration in seconds | [optional] 
**EndDate** | Pointer to **int32** |  | [optional] 
**GeoLatitude** | Pointer to **float32** | Specifies the global position for the activity | [optional] 
**GeoLongitude** | Pointer to **float32** | Specifies the global position for the activity | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto-generated unique identifier | [optional] [readonly] 
**LinkedBy** | Pointer to **string** | An array of Schedule Event Ids that their start time depends on the end of the current. | [optional] 
**LinkedTo** | Pointer to [**KalturaLinkedScheduleEvent**](KalturaLinkedScheduleEvent.md) |  | [optional] 
**Location** | Pointer to **string** | Defines the intended venue for the activity | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Organizer** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Priority** | Pointer to **int32** | The value for the priority field. | [optional] 
**Recurrence** | Pointer to [**KalturaScheduleEventRecurrence**](KalturaScheduleEventRecurrence.md) |  | [optional] 
**RecurrenceType** | Pointer to **int32** | Enum Type: &#x60;KalturaScheduleEventRecurrenceType&#x60; | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Sequence** | Pointer to **int32** | Defines the revision sequence number. | [optional] 
**StartDate** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaScheduleEventStatus&#x60; | [optional] [readonly] 
**Summary** | Pointer to **string** | Defines a short summary or subject for the event | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaScheduleEvent

`func NewKalturaScheduleEvent() *KalturaScheduleEvent`

NewKalturaScheduleEvent instantiates a new KalturaScheduleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaScheduleEventWithDefaults

`func NewKalturaScheduleEventWithDefaults() *KalturaScheduleEvent`

NewKalturaScheduleEventWithDefaults instantiates a new KalturaScheduleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassificationType

`func (o *KalturaScheduleEvent) GetClassificationType() int32`

GetClassificationType returns the ClassificationType field if non-nil, zero value otherwise.

### GetClassificationTypeOk

`func (o *KalturaScheduleEvent) GetClassificationTypeOk() (*int32, bool)`

GetClassificationTypeOk returns a tuple with the ClassificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationType

`func (o *KalturaScheduleEvent) SetClassificationType(v int32)`

SetClassificationType sets ClassificationType field to given value.

### HasClassificationType

`func (o *KalturaScheduleEvent) HasClassificationType() bool`

HasClassificationType returns a boolean if a field has been set.

### GetComment

`func (o *KalturaScheduleEvent) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *KalturaScheduleEvent) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *KalturaScheduleEvent) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *KalturaScheduleEvent) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContact

`func (o *KalturaScheduleEvent) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *KalturaScheduleEvent) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *KalturaScheduleEvent) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *KalturaScheduleEvent) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaScheduleEvent) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaScheduleEvent) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaScheduleEvent) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaScheduleEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaScheduleEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaScheduleEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaScheduleEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaScheduleEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *KalturaScheduleEvent) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KalturaScheduleEvent) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KalturaScheduleEvent) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KalturaScheduleEvent) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndDate

`func (o *KalturaScheduleEvent) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *KalturaScheduleEvent) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *KalturaScheduleEvent) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *KalturaScheduleEvent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGeoLatitude

`func (o *KalturaScheduleEvent) GetGeoLatitude() float32`

GetGeoLatitude returns the GeoLatitude field if non-nil, zero value otherwise.

### GetGeoLatitudeOk

`func (o *KalturaScheduleEvent) GetGeoLatitudeOk() (*float32, bool)`

GetGeoLatitudeOk returns a tuple with the GeoLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLatitude

`func (o *KalturaScheduleEvent) SetGeoLatitude(v float32)`

SetGeoLatitude sets GeoLatitude field to given value.

### HasGeoLatitude

`func (o *KalturaScheduleEvent) HasGeoLatitude() bool`

HasGeoLatitude returns a boolean if a field has been set.

### GetGeoLongitude

`func (o *KalturaScheduleEvent) GetGeoLongitude() float32`

GetGeoLongitude returns the GeoLongitude field if non-nil, zero value otherwise.

### GetGeoLongitudeOk

`func (o *KalturaScheduleEvent) GetGeoLongitudeOk() (*float32, bool)`

GetGeoLongitudeOk returns a tuple with the GeoLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLongitude

`func (o *KalturaScheduleEvent) SetGeoLongitude(v float32)`

SetGeoLongitude sets GeoLongitude field to given value.

### HasGeoLongitude

`func (o *KalturaScheduleEvent) HasGeoLongitude() bool`

HasGeoLongitude returns a boolean if a field has been set.

### GetId

`func (o *KalturaScheduleEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaScheduleEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaScheduleEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaScheduleEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinkedBy

`func (o *KalturaScheduleEvent) GetLinkedBy() string`

GetLinkedBy returns the LinkedBy field if non-nil, zero value otherwise.

### GetLinkedByOk

`func (o *KalturaScheduleEvent) GetLinkedByOk() (*string, bool)`

GetLinkedByOk returns a tuple with the LinkedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedBy

`func (o *KalturaScheduleEvent) SetLinkedBy(v string)`

SetLinkedBy sets LinkedBy field to given value.

### HasLinkedBy

`func (o *KalturaScheduleEvent) HasLinkedBy() bool`

HasLinkedBy returns a boolean if a field has been set.

### GetLinkedTo

`func (o *KalturaScheduleEvent) GetLinkedTo() KalturaLinkedScheduleEvent`

GetLinkedTo returns the LinkedTo field if non-nil, zero value otherwise.

### GetLinkedToOk

`func (o *KalturaScheduleEvent) GetLinkedToOk() (*KalturaLinkedScheduleEvent, bool)`

GetLinkedToOk returns a tuple with the LinkedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTo

`func (o *KalturaScheduleEvent) SetLinkedTo(v KalturaLinkedScheduleEvent)`

SetLinkedTo sets LinkedTo field to given value.

### HasLinkedTo

`func (o *KalturaScheduleEvent) HasLinkedTo() bool`

HasLinkedTo returns a boolean if a field has been set.

### GetLocation

`func (o *KalturaScheduleEvent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *KalturaScheduleEvent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *KalturaScheduleEvent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *KalturaScheduleEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaScheduleEvent) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaScheduleEvent) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaScheduleEvent) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaScheduleEvent) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOrganizer

`func (o *KalturaScheduleEvent) GetOrganizer() string`

GetOrganizer returns the Organizer field if non-nil, zero value otherwise.

### GetOrganizerOk

`func (o *KalturaScheduleEvent) GetOrganizerOk() (*string, bool)`

GetOrganizerOk returns a tuple with the Organizer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizer

`func (o *KalturaScheduleEvent) SetOrganizer(v string)`

SetOrganizer sets Organizer field to given value.

### HasOrganizer

`func (o *KalturaScheduleEvent) HasOrganizer() bool`

HasOrganizer returns a boolean if a field has been set.

### GetOwnerId

`func (o *KalturaScheduleEvent) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *KalturaScheduleEvent) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *KalturaScheduleEvent) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *KalturaScheduleEvent) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetParentId

`func (o *KalturaScheduleEvent) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *KalturaScheduleEvent) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *KalturaScheduleEvent) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *KalturaScheduleEvent) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaScheduleEvent) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaScheduleEvent) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaScheduleEvent) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaScheduleEvent) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPriority

`func (o *KalturaScheduleEvent) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *KalturaScheduleEvent) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *KalturaScheduleEvent) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *KalturaScheduleEvent) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRecurrence

`func (o *KalturaScheduleEvent) GetRecurrence() KalturaScheduleEventRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *KalturaScheduleEvent) GetRecurrenceOk() (*KalturaScheduleEventRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *KalturaScheduleEvent) SetRecurrence(v KalturaScheduleEventRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *KalturaScheduleEvent) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetRecurrenceType

`func (o *KalturaScheduleEvent) GetRecurrenceType() int32`

GetRecurrenceType returns the RecurrenceType field if non-nil, zero value otherwise.

### GetRecurrenceTypeOk

`func (o *KalturaScheduleEvent) GetRecurrenceTypeOk() (*int32, bool)`

GetRecurrenceTypeOk returns a tuple with the RecurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceType

`func (o *KalturaScheduleEvent) SetRecurrenceType(v int32)`

SetRecurrenceType sets RecurrenceType field to given value.

### HasRecurrenceType

`func (o *KalturaScheduleEvent) HasRecurrenceType() bool`

HasRecurrenceType returns a boolean if a field has been set.

### GetReferenceId

`func (o *KalturaScheduleEvent) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KalturaScheduleEvent) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KalturaScheduleEvent) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KalturaScheduleEvent) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSequence

`func (o *KalturaScheduleEvent) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *KalturaScheduleEvent) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *KalturaScheduleEvent) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *KalturaScheduleEvent) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetStartDate

`func (o *KalturaScheduleEvent) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *KalturaScheduleEvent) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *KalturaScheduleEvent) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *KalturaScheduleEvent) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaScheduleEvent) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaScheduleEvent) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaScheduleEvent) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaScheduleEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *KalturaScheduleEvent) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *KalturaScheduleEvent) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *KalturaScheduleEvent) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *KalturaScheduleEvent) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTags

`func (o *KalturaScheduleEvent) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaScheduleEvent) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaScheduleEvent) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaScheduleEvent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaScheduleEvent) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaScheduleEvent) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaScheduleEvent) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaScheduleEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


