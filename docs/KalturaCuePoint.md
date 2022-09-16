# KalturaCuePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopiedFrom** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CuePointType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaCuePointType&#x60; | [optional] [readonly] 
**EntryId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**ForceStop** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IntId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IsMomentary** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerData** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerSortValue** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **int32** | Start time in milliseconds | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaCuePointStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**ThumbOffset** | Pointer to **int32** |  | [optional] 
**TriggeredAt** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaCuePoint

`func NewKalturaCuePoint() *KalturaCuePoint`

NewKalturaCuePoint instantiates a new KalturaCuePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaCuePointWithDefaults

`func NewKalturaCuePointWithDefaults() *KalturaCuePoint`

NewKalturaCuePointWithDefaults instantiates a new KalturaCuePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopiedFrom

`func (o *KalturaCuePoint) GetCopiedFrom() string`

GetCopiedFrom returns the CopiedFrom field if non-nil, zero value otherwise.

### GetCopiedFromOk

`func (o *KalturaCuePoint) GetCopiedFromOk() (*string, bool)`

GetCopiedFromOk returns a tuple with the CopiedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiedFrom

`func (o *KalturaCuePoint) SetCopiedFrom(v string)`

SetCopiedFrom sets CopiedFrom field to given value.

### HasCopiedFrom

`func (o *KalturaCuePoint) HasCopiedFrom() bool`

HasCopiedFrom returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaCuePoint) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaCuePoint) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaCuePoint) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaCuePoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCuePointType

`func (o *KalturaCuePoint) GetCuePointType() string`

GetCuePointType returns the CuePointType field if non-nil, zero value otherwise.

### GetCuePointTypeOk

`func (o *KalturaCuePoint) GetCuePointTypeOk() (*string, bool)`

GetCuePointTypeOk returns a tuple with the CuePointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuePointType

`func (o *KalturaCuePoint) SetCuePointType(v string)`

SetCuePointType sets CuePointType field to given value.

### HasCuePointType

`func (o *KalturaCuePoint) HasCuePointType() bool`

HasCuePointType returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaCuePoint) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaCuePoint) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaCuePoint) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaCuePoint) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetForceStop

`func (o *KalturaCuePoint) GetForceStop() int32`

GetForceStop returns the ForceStop field if non-nil, zero value otherwise.

### GetForceStopOk

`func (o *KalturaCuePoint) GetForceStopOk() (*int32, bool)`

GetForceStopOk returns a tuple with the ForceStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceStop

`func (o *KalturaCuePoint) SetForceStop(v int32)`

SetForceStop sets ForceStop field to given value.

### HasForceStop

`func (o *KalturaCuePoint) HasForceStop() bool`

HasForceStop returns a boolean if a field has been set.

### GetId

`func (o *KalturaCuePoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaCuePoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaCuePoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaCuePoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntId

`func (o *KalturaCuePoint) GetIntId() int32`

GetIntId returns the IntId field if non-nil, zero value otherwise.

### GetIntIdOk

`func (o *KalturaCuePoint) GetIntIdOk() (*int32, bool)`

GetIntIdOk returns a tuple with the IntId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntId

`func (o *KalturaCuePoint) SetIntId(v int32)`

SetIntId sets IntId field to given value.

### HasIntId

`func (o *KalturaCuePoint) HasIntId() bool`

HasIntId returns a boolean if a field has been set.

### GetIsMomentary

`func (o *KalturaCuePoint) GetIsMomentary() bool`

GetIsMomentary returns the IsMomentary field if non-nil, zero value otherwise.

### GetIsMomentaryOk

`func (o *KalturaCuePoint) GetIsMomentaryOk() (*bool, bool)`

GetIsMomentaryOk returns a tuple with the IsMomentary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMomentary

`func (o *KalturaCuePoint) SetIsMomentary(v bool)`

SetIsMomentary sets IsMomentary field to given value.

### HasIsMomentary

`func (o *KalturaCuePoint) HasIsMomentary() bool`

HasIsMomentary returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaCuePoint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaCuePoint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaCuePoint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaCuePoint) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaCuePoint) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaCuePoint) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaCuePoint) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaCuePoint) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaCuePoint) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaCuePoint) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaCuePoint) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaCuePoint) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerSortValue

`func (o *KalturaCuePoint) GetPartnerSortValue() int32`

GetPartnerSortValue returns the PartnerSortValue field if non-nil, zero value otherwise.

### GetPartnerSortValueOk

`func (o *KalturaCuePoint) GetPartnerSortValueOk() (*int32, bool)`

GetPartnerSortValueOk returns a tuple with the PartnerSortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSortValue

`func (o *KalturaCuePoint) SetPartnerSortValue(v int32)`

SetPartnerSortValue sets PartnerSortValue field to given value.

### HasPartnerSortValue

`func (o *KalturaCuePoint) HasPartnerSortValue() bool`

HasPartnerSortValue returns a boolean if a field has been set.

### GetStartTime

`func (o *KalturaCuePoint) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KalturaCuePoint) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KalturaCuePoint) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *KalturaCuePoint) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaCuePoint) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaCuePoint) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaCuePoint) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaCuePoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaCuePoint) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaCuePoint) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaCuePoint) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaCuePoint) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaCuePoint) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaCuePoint) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaCuePoint) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaCuePoint) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbOffset

`func (o *KalturaCuePoint) GetThumbOffset() int32`

GetThumbOffset returns the ThumbOffset field if non-nil, zero value otherwise.

### GetThumbOffsetOk

`func (o *KalturaCuePoint) GetThumbOffsetOk() (*int32, bool)`

GetThumbOffsetOk returns a tuple with the ThumbOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbOffset

`func (o *KalturaCuePoint) SetThumbOffset(v int32)`

SetThumbOffset sets ThumbOffset field to given value.

### HasThumbOffset

`func (o *KalturaCuePoint) HasThumbOffset() bool`

HasThumbOffset returns a boolean if a field has been set.

### GetTriggeredAt

`func (o *KalturaCuePoint) GetTriggeredAt() int32`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *KalturaCuePoint) GetTriggeredAtOk() (*int32, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *KalturaCuePoint) SetTriggeredAt(v int32)`

SetTriggeredAt sets TriggeredAt field to given value.

### HasTriggeredAt

`func (o *KalturaCuePoint) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaCuePoint) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaCuePoint) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaCuePoint) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaCuePoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaCuePoint) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaCuePoint) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaCuePoint) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaCuePoint) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


