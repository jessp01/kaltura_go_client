# KalturaEmailIngestionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionProfile2Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefaultAdminTags** | Pointer to **string** |  | [optional] 
**DefaultCategory** | Pointer to **string** |  | [optional] 
**DefaultTags** | Pointer to **string** |  | [optional] 
**DefaultUserId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MailboxId** | Pointer to **string** |  | [optional] 
**MaxAttachmentSizeKbytes** | Pointer to **int32** |  | [optional] 
**MaxAttachmentsPerMail** | Pointer to **int32** |  | [optional] 
**ModerationStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaEntryModerationStatus&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEmailIngestionProfileStatus&#x60; | [optional] [readonly] 

## Methods

### NewKalturaEmailIngestionProfile

`func NewKalturaEmailIngestionProfile() *KalturaEmailIngestionProfile`

NewKalturaEmailIngestionProfile instantiates a new KalturaEmailIngestionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEmailIngestionProfileWithDefaults

`func NewKalturaEmailIngestionProfileWithDefaults() *KalturaEmailIngestionProfile`

NewKalturaEmailIngestionProfileWithDefaults instantiates a new KalturaEmailIngestionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionProfile2Id

`func (o *KalturaEmailIngestionProfile) GetConversionProfile2Id() int32`

GetConversionProfile2Id returns the ConversionProfile2Id field if non-nil, zero value otherwise.

### GetConversionProfile2IdOk

`func (o *KalturaEmailIngestionProfile) GetConversionProfile2IdOk() (*int32, bool)`

GetConversionProfile2IdOk returns a tuple with the ConversionProfile2Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfile2Id

`func (o *KalturaEmailIngestionProfile) SetConversionProfile2Id(v int32)`

SetConversionProfile2Id sets ConversionProfile2Id field to given value.

### HasConversionProfile2Id

`func (o *KalturaEmailIngestionProfile) HasConversionProfile2Id() bool`

HasConversionProfile2Id returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaEmailIngestionProfile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaEmailIngestionProfile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaEmailIngestionProfile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaEmailIngestionProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultAdminTags

`func (o *KalturaEmailIngestionProfile) GetDefaultAdminTags() string`

GetDefaultAdminTags returns the DefaultAdminTags field if non-nil, zero value otherwise.

### GetDefaultAdminTagsOk

`func (o *KalturaEmailIngestionProfile) GetDefaultAdminTagsOk() (*string, bool)`

GetDefaultAdminTagsOk returns a tuple with the DefaultAdminTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAdminTags

`func (o *KalturaEmailIngestionProfile) SetDefaultAdminTags(v string)`

SetDefaultAdminTags sets DefaultAdminTags field to given value.

### HasDefaultAdminTags

`func (o *KalturaEmailIngestionProfile) HasDefaultAdminTags() bool`

HasDefaultAdminTags returns a boolean if a field has been set.

### GetDefaultCategory

`func (o *KalturaEmailIngestionProfile) GetDefaultCategory() string`

GetDefaultCategory returns the DefaultCategory field if non-nil, zero value otherwise.

### GetDefaultCategoryOk

`func (o *KalturaEmailIngestionProfile) GetDefaultCategoryOk() (*string, bool)`

GetDefaultCategoryOk returns a tuple with the DefaultCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCategory

`func (o *KalturaEmailIngestionProfile) SetDefaultCategory(v string)`

SetDefaultCategory sets DefaultCategory field to given value.

### HasDefaultCategory

`func (o *KalturaEmailIngestionProfile) HasDefaultCategory() bool`

HasDefaultCategory returns a boolean if a field has been set.

### GetDefaultTags

`func (o *KalturaEmailIngestionProfile) GetDefaultTags() string`

GetDefaultTags returns the DefaultTags field if non-nil, zero value otherwise.

### GetDefaultTagsOk

`func (o *KalturaEmailIngestionProfile) GetDefaultTagsOk() (*string, bool)`

GetDefaultTagsOk returns a tuple with the DefaultTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTags

`func (o *KalturaEmailIngestionProfile) SetDefaultTags(v string)`

SetDefaultTags sets DefaultTags field to given value.

### HasDefaultTags

`func (o *KalturaEmailIngestionProfile) HasDefaultTags() bool`

HasDefaultTags returns a boolean if a field has been set.

### GetDefaultUserId

`func (o *KalturaEmailIngestionProfile) GetDefaultUserId() string`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *KalturaEmailIngestionProfile) GetDefaultUserIdOk() (*string, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *KalturaEmailIngestionProfile) SetDefaultUserId(v string)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *KalturaEmailIngestionProfile) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaEmailIngestionProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaEmailIngestionProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaEmailIngestionProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaEmailIngestionProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmailAddress

`func (o *KalturaEmailIngestionProfile) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *KalturaEmailIngestionProfile) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *KalturaEmailIngestionProfile) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *KalturaEmailIngestionProfile) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetId

`func (o *KalturaEmailIngestionProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaEmailIngestionProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaEmailIngestionProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaEmailIngestionProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMailboxId

`func (o *KalturaEmailIngestionProfile) GetMailboxId() string`

GetMailboxId returns the MailboxId field if non-nil, zero value otherwise.

### GetMailboxIdOk

`func (o *KalturaEmailIngestionProfile) GetMailboxIdOk() (*string, bool)`

GetMailboxIdOk returns a tuple with the MailboxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailboxId

`func (o *KalturaEmailIngestionProfile) SetMailboxId(v string)`

SetMailboxId sets MailboxId field to given value.

### HasMailboxId

`func (o *KalturaEmailIngestionProfile) HasMailboxId() bool`

HasMailboxId returns a boolean if a field has been set.

### GetMaxAttachmentSizeKbytes

`func (o *KalturaEmailIngestionProfile) GetMaxAttachmentSizeKbytes() int32`

GetMaxAttachmentSizeKbytes returns the MaxAttachmentSizeKbytes field if non-nil, zero value otherwise.

### GetMaxAttachmentSizeKbytesOk

`func (o *KalturaEmailIngestionProfile) GetMaxAttachmentSizeKbytesOk() (*int32, bool)`

GetMaxAttachmentSizeKbytesOk returns a tuple with the MaxAttachmentSizeKbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttachmentSizeKbytes

`func (o *KalturaEmailIngestionProfile) SetMaxAttachmentSizeKbytes(v int32)`

SetMaxAttachmentSizeKbytes sets MaxAttachmentSizeKbytes field to given value.

### HasMaxAttachmentSizeKbytes

`func (o *KalturaEmailIngestionProfile) HasMaxAttachmentSizeKbytes() bool`

HasMaxAttachmentSizeKbytes returns a boolean if a field has been set.

### GetMaxAttachmentsPerMail

`func (o *KalturaEmailIngestionProfile) GetMaxAttachmentsPerMail() int32`

GetMaxAttachmentsPerMail returns the MaxAttachmentsPerMail field if non-nil, zero value otherwise.

### GetMaxAttachmentsPerMailOk

`func (o *KalturaEmailIngestionProfile) GetMaxAttachmentsPerMailOk() (*int32, bool)`

GetMaxAttachmentsPerMailOk returns a tuple with the MaxAttachmentsPerMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttachmentsPerMail

`func (o *KalturaEmailIngestionProfile) SetMaxAttachmentsPerMail(v int32)`

SetMaxAttachmentsPerMail sets MaxAttachmentsPerMail field to given value.

### HasMaxAttachmentsPerMail

`func (o *KalturaEmailIngestionProfile) HasMaxAttachmentsPerMail() bool`

HasMaxAttachmentsPerMail returns a boolean if a field has been set.

### GetModerationStatus

`func (o *KalturaEmailIngestionProfile) GetModerationStatus() int32`

GetModerationStatus returns the ModerationStatus field if non-nil, zero value otherwise.

### GetModerationStatusOk

`func (o *KalturaEmailIngestionProfile) GetModerationStatusOk() (*int32, bool)`

GetModerationStatusOk returns a tuple with the ModerationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationStatus

`func (o *KalturaEmailIngestionProfile) SetModerationStatus(v int32)`

SetModerationStatus sets ModerationStatus field to given value.

### HasModerationStatus

`func (o *KalturaEmailIngestionProfile) HasModerationStatus() bool`

HasModerationStatus returns a boolean if a field has been set.

### GetName

`func (o *KalturaEmailIngestionProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaEmailIngestionProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaEmailIngestionProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaEmailIngestionProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaEmailIngestionProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaEmailIngestionProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaEmailIngestionProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaEmailIngestionProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaEmailIngestionProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaEmailIngestionProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaEmailIngestionProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaEmailIngestionProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


