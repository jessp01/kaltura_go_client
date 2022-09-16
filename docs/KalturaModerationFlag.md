# KalturaModerationFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** | The comment that was added to the flag | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FlagType** | Pointer to **int32** | Enum Type: &#x60;KalturaModerationFlagType&#x60; | [optional] 
**FlaggedEntryId** | Pointer to **string** | If moderation flag is set for entry, this is the flagged entry id | [optional] 
**FlaggedUserId** | Pointer to **string** | If moderation flag is set for user, this is the flagged user id | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Moderation flag id | [optional] [readonly] 
**ModerationObjectType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaModerationObjectType&#x60;  The type of the moderation flag (entry or user) | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaModerationFlagStatus&#x60;  The moderation flag status | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;readOnly&#x60;  The user id that added the moderation flag | [optional] [readonly] 

## Methods

### NewKalturaModerationFlag

`func NewKalturaModerationFlag() *KalturaModerationFlag`

NewKalturaModerationFlag instantiates a new KalturaModerationFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaModerationFlagWithDefaults

`func NewKalturaModerationFlagWithDefaults() *KalturaModerationFlag`

NewKalturaModerationFlagWithDefaults instantiates a new KalturaModerationFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *KalturaModerationFlag) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *KalturaModerationFlag) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *KalturaModerationFlag) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *KalturaModerationFlag) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaModerationFlag) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaModerationFlag) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaModerationFlag) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaModerationFlag) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFlagType

`func (o *KalturaModerationFlag) GetFlagType() int32`

GetFlagType returns the FlagType field if non-nil, zero value otherwise.

### GetFlagTypeOk

`func (o *KalturaModerationFlag) GetFlagTypeOk() (*int32, bool)`

GetFlagTypeOk returns a tuple with the FlagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagType

`func (o *KalturaModerationFlag) SetFlagType(v int32)`

SetFlagType sets FlagType field to given value.

### HasFlagType

`func (o *KalturaModerationFlag) HasFlagType() bool`

HasFlagType returns a boolean if a field has been set.

### GetFlaggedEntryId

`func (o *KalturaModerationFlag) GetFlaggedEntryId() string`

GetFlaggedEntryId returns the FlaggedEntryId field if non-nil, zero value otherwise.

### GetFlaggedEntryIdOk

`func (o *KalturaModerationFlag) GetFlaggedEntryIdOk() (*string, bool)`

GetFlaggedEntryIdOk returns a tuple with the FlaggedEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedEntryId

`func (o *KalturaModerationFlag) SetFlaggedEntryId(v string)`

SetFlaggedEntryId sets FlaggedEntryId field to given value.

### HasFlaggedEntryId

`func (o *KalturaModerationFlag) HasFlaggedEntryId() bool`

HasFlaggedEntryId returns a boolean if a field has been set.

### GetFlaggedUserId

`func (o *KalturaModerationFlag) GetFlaggedUserId() string`

GetFlaggedUserId returns the FlaggedUserId field if non-nil, zero value otherwise.

### GetFlaggedUserIdOk

`func (o *KalturaModerationFlag) GetFlaggedUserIdOk() (*string, bool)`

GetFlaggedUserIdOk returns a tuple with the FlaggedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedUserId

`func (o *KalturaModerationFlag) SetFlaggedUserId(v string)`

SetFlaggedUserId sets FlaggedUserId field to given value.

### HasFlaggedUserId

`func (o *KalturaModerationFlag) HasFlaggedUserId() bool`

HasFlaggedUserId returns a boolean if a field has been set.

### GetId

`func (o *KalturaModerationFlag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaModerationFlag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaModerationFlag) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaModerationFlag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModerationObjectType

`func (o *KalturaModerationFlag) GetModerationObjectType() string`

GetModerationObjectType returns the ModerationObjectType field if non-nil, zero value otherwise.

### GetModerationObjectTypeOk

`func (o *KalturaModerationFlag) GetModerationObjectTypeOk() (*string, bool)`

GetModerationObjectTypeOk returns a tuple with the ModerationObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationObjectType

`func (o *KalturaModerationFlag) SetModerationObjectType(v string)`

SetModerationObjectType sets ModerationObjectType field to given value.

### HasModerationObjectType

`func (o *KalturaModerationFlag) HasModerationObjectType() bool`

HasModerationObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaModerationFlag) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaModerationFlag) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaModerationFlag) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaModerationFlag) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaModerationFlag) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaModerationFlag) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaModerationFlag) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaModerationFlag) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaModerationFlag) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaModerationFlag) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaModerationFlag) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaModerationFlag) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaModerationFlag) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaModerationFlag) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaModerationFlag) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaModerationFlag) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


