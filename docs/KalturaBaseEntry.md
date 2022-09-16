# KalturaBaseEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlId** | Pointer to **int32** | The Access Control ID assigned to this entry (null when not set, send -1 to remove) | [optional] 
**AdminTags** | Pointer to **string** | Entry admin tags can be updated only by administrators | [optional] 
**Application** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaEntryApplication&#x60;  Entry application | [optional] 
**ApplicationVersion** | Pointer to **string** | &#x60;insertOnly&#x60;  Entry application version | [optional] 
**BlockAutoTranscript** | Pointer to **bool** | Block auto transcript on Entry | [optional] 
**Capabilities** | Pointer to **string** | &#x60;readOnly&#x60;  Comma seperated string of the capabilities of the entry. Any capability needed can be added to this list. | [optional] [readonly] 
**Categories** | Pointer to **string** | Comma separated list of full names of categories to which this entry belongs. Only categories that don&#39;t have entitlement (privacy context) are listed, to retrieve the full list of categories, use the categoryEntry.list action. | [optional] 
**CategoriesIds** | Pointer to **string** | Comma separated list of ids of categories to which this entry belongs. Only categories that don&#39;t have entitlement (privacy context) are listed, to retrieve the full list of categories, use the categoryEntry.list action. | [optional] 
**ConversionProfileId** | Pointer to **int32** | Override the default ingestion profile | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CreatorId** | Pointer to **string** | &#x60;insertOnly&#x60;  The ID of the user who created this entry | [optional] 
**Description** | Pointer to **string** | Entry description | [optional] 
**DisplayInSearch** | Pointer to **int32** | Enum Type: &#x60;KalturaEntryDisplayInSearchType&#x60;  should we display this entry in search | [optional] 
**DownloadUrl** | Pointer to **string** | &#x60;readOnly&#x60;  Download URL for the entry | [optional] [readonly] 
**EndDate** | Pointer to **int32** | Entry scheduling end date (null when not set, send -1 to remove) | [optional] 
**EntitledUsersEdit** | Pointer to **string** | list of user ids that are entitled to edit the entry (no server enforcement) The difference between entitledUsersEdit, entitledUsersPublish and entitledUsersView is applicative only | [optional] 
**EntitledUsersPublish** | Pointer to **string** | list of user ids that are entitled to publish the entry (no server enforcement) The difference between entitledUsersEdit, entitledUsersPublish and entitledUsersView is applicative only | [optional] 
**EntitledUsersView** | Pointer to **string** | list of user ids that are entitled to view the entry (no server enforcement) The difference between entitledUsersEdit, entitledUsersPublish and entitledUsersView is applicative only | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60;  Auto generated 10 characters alphanumeric string | [optional] [readonly] 
**LicenseType** | Pointer to **int32** | Enum Type: &#x60;KalturaLicenseType&#x60;  License type used for this entry | [optional] 
**ModerationCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of moderation requests waiting for this entry | [optional] [readonly] 
**ModerationStatus** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryModerationStatus&#x60;  Entry moderation status | [optional] [readonly] 
**Name** | Pointer to **string** | Entry name (Min 1 chars) | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**OperationAttributes** | Pointer to [**[]KalturaOperationAttributes**](KalturaOperationAttributes.md) |  | [optional] 
**ParentEntryId** | Pointer to **string** | ID of source root entry, used for defining entires association | [optional] 
**PartnerData** | Pointer to **string** | Can be used to store various partner related data as a string | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerSortValue** | Pointer to **int32** | Can be used to store various partner related data as a numeric value | [optional] 
**Rank** | Pointer to **float32** | &#x60;readOnly&#x60;  The calculated average rank. rank &#x3D; totalRank / votes | [optional] [readonly] 
**RedirectEntryId** | Pointer to **string** | IF not empty, points to an entry ID the should replace this current entry&#39;s id. | [optional] 
**ReferenceId** | Pointer to **string** | Entry external reference id | [optional] 
**ReplacedEntryId** | Pointer to **string** | &#x60;readOnly&#x60;  ID of the entry that will be replaced when the replacement approved and this entry is ready | [optional] [readonly] 
**ReplacementStatus** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryReplacementStatus&#x60;  Status of the replacement readiness and approval | [optional] [readonly] 
**ReplacingEntryId** | Pointer to **string** | &#x60;readOnly&#x60;  ID of temporary entry that will replace this entry when it&#39;s approved and ready for replacement | [optional] [readonly] 
**RootEntryId** | Pointer to **string** | &#x60;readOnly&#x60;  ID of source root entry, used for clipped, skipped and cropped entries that created from another entry | [optional] [readonly] 
**SearchText** | Pointer to **string** | &#x60;readOnly&#x60;  Indexed search text for full text search | [optional] [readonly] 
**StartDate** | Pointer to **int32** | Entry scheduling start date (null when not set, send -1 to remove) | [optional] 
**Status** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEntryStatus&#x60; | [optional] [readonly] 
**Tags** | Pointer to **string** | Entry tags | [optional] 
**TemplateEntryId** | Pointer to **string** | &#x60;insertOnly&#x60;  Template entry id | [optional] 
**ThumbnailUrl** | Pointer to **string** | &#x60;readOnly&#x60;  Thumbnail URL | [optional] [readonly] 
**TotalRank** | Pointer to **int32** | &#x60;readOnly&#x60;  The sum of all rank values submitted to the baseEntry.anonymousRank action | [optional] [readonly] 
**Type** | Pointer to **string** | Enum Type: &#x60;KalturaEntryType&#x60;  The type of the entry, this is auto filled by the derived entry object | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UserId** | Pointer to **string** | The ID of the user who is the owner of this entry | [optional] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60;  Version of the entry data | [optional] [readonly] 
**Votes** | Pointer to **int32** | &#x60;readOnly&#x60;  A count of all requests made to the baseEntry.anonymousRank action | [optional] [readonly] 

## Methods

### NewKalturaBaseEntry

`func NewKalturaBaseEntry() *KalturaBaseEntry`

NewKalturaBaseEntry instantiates a new KalturaBaseEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBaseEntryWithDefaults

`func NewKalturaBaseEntryWithDefaults() *KalturaBaseEntry`

NewKalturaBaseEntryWithDefaults instantiates a new KalturaBaseEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlId

`func (o *KalturaBaseEntry) GetAccessControlId() int32`

GetAccessControlId returns the AccessControlId field if non-nil, zero value otherwise.

### GetAccessControlIdOk

`func (o *KalturaBaseEntry) GetAccessControlIdOk() (*int32, bool)`

GetAccessControlIdOk returns a tuple with the AccessControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlId

`func (o *KalturaBaseEntry) SetAccessControlId(v int32)`

SetAccessControlId sets AccessControlId field to given value.

### HasAccessControlId

`func (o *KalturaBaseEntry) HasAccessControlId() bool`

HasAccessControlId returns a boolean if a field has been set.

### GetAdminTags

`func (o *KalturaBaseEntry) GetAdminTags() string`

GetAdminTags returns the AdminTags field if non-nil, zero value otherwise.

### GetAdminTagsOk

`func (o *KalturaBaseEntry) GetAdminTagsOk() (*string, bool)`

GetAdminTagsOk returns a tuple with the AdminTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTags

`func (o *KalturaBaseEntry) SetAdminTags(v string)`

SetAdminTags sets AdminTags field to given value.

### HasAdminTags

`func (o *KalturaBaseEntry) HasAdminTags() bool`

HasAdminTags returns a boolean if a field has been set.

### GetApplication

`func (o *KalturaBaseEntry) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *KalturaBaseEntry) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *KalturaBaseEntry) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *KalturaBaseEntry) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *KalturaBaseEntry) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *KalturaBaseEntry) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *KalturaBaseEntry) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *KalturaBaseEntry) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetBlockAutoTranscript

`func (o *KalturaBaseEntry) GetBlockAutoTranscript() bool`

GetBlockAutoTranscript returns the BlockAutoTranscript field if non-nil, zero value otherwise.

### GetBlockAutoTranscriptOk

`func (o *KalturaBaseEntry) GetBlockAutoTranscriptOk() (*bool, bool)`

GetBlockAutoTranscriptOk returns a tuple with the BlockAutoTranscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAutoTranscript

`func (o *KalturaBaseEntry) SetBlockAutoTranscript(v bool)`

SetBlockAutoTranscript sets BlockAutoTranscript field to given value.

### HasBlockAutoTranscript

`func (o *KalturaBaseEntry) HasBlockAutoTranscript() bool`

HasBlockAutoTranscript returns a boolean if a field has been set.

### GetCapabilities

`func (o *KalturaBaseEntry) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *KalturaBaseEntry) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *KalturaBaseEntry) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *KalturaBaseEntry) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCategories

`func (o *KalturaBaseEntry) GetCategories() string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *KalturaBaseEntry) GetCategoriesOk() (*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *KalturaBaseEntry) SetCategories(v string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *KalturaBaseEntry) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoriesIds

`func (o *KalturaBaseEntry) GetCategoriesIds() string`

GetCategoriesIds returns the CategoriesIds field if non-nil, zero value otherwise.

### GetCategoriesIdsOk

`func (o *KalturaBaseEntry) GetCategoriesIdsOk() (*string, bool)`

GetCategoriesIdsOk returns a tuple with the CategoriesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesIds

`func (o *KalturaBaseEntry) SetCategoriesIds(v string)`

SetCategoriesIds sets CategoriesIds field to given value.

### HasCategoriesIds

`func (o *KalturaBaseEntry) HasCategoriesIds() bool`

HasCategoriesIds returns a boolean if a field has been set.

### GetConversionProfileId

`func (o *KalturaBaseEntry) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *KalturaBaseEntry) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *KalturaBaseEntry) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *KalturaBaseEntry) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaBaseEntry) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaBaseEntry) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaBaseEntry) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaBaseEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *KalturaBaseEntry) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *KalturaBaseEntry) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *KalturaBaseEntry) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *KalturaBaseEntry) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaBaseEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaBaseEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaBaseEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaBaseEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayInSearch

`func (o *KalturaBaseEntry) GetDisplayInSearch() int32`

GetDisplayInSearch returns the DisplayInSearch field if non-nil, zero value otherwise.

### GetDisplayInSearchOk

`func (o *KalturaBaseEntry) GetDisplayInSearchOk() (*int32, bool)`

GetDisplayInSearchOk returns a tuple with the DisplayInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInSearch

`func (o *KalturaBaseEntry) SetDisplayInSearch(v int32)`

SetDisplayInSearch sets DisplayInSearch field to given value.

### HasDisplayInSearch

`func (o *KalturaBaseEntry) HasDisplayInSearch() bool`

HasDisplayInSearch returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *KalturaBaseEntry) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *KalturaBaseEntry) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *KalturaBaseEntry) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *KalturaBaseEntry) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetEndDate

`func (o *KalturaBaseEntry) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *KalturaBaseEntry) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *KalturaBaseEntry) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *KalturaBaseEntry) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEntitledUsersEdit

`func (o *KalturaBaseEntry) GetEntitledUsersEdit() string`

GetEntitledUsersEdit returns the EntitledUsersEdit field if non-nil, zero value otherwise.

### GetEntitledUsersEditOk

`func (o *KalturaBaseEntry) GetEntitledUsersEditOk() (*string, bool)`

GetEntitledUsersEditOk returns a tuple with the EntitledUsersEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledUsersEdit

`func (o *KalturaBaseEntry) SetEntitledUsersEdit(v string)`

SetEntitledUsersEdit sets EntitledUsersEdit field to given value.

### HasEntitledUsersEdit

`func (o *KalturaBaseEntry) HasEntitledUsersEdit() bool`

HasEntitledUsersEdit returns a boolean if a field has been set.

### GetEntitledUsersPublish

`func (o *KalturaBaseEntry) GetEntitledUsersPublish() string`

GetEntitledUsersPublish returns the EntitledUsersPublish field if non-nil, zero value otherwise.

### GetEntitledUsersPublishOk

`func (o *KalturaBaseEntry) GetEntitledUsersPublishOk() (*string, bool)`

GetEntitledUsersPublishOk returns a tuple with the EntitledUsersPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledUsersPublish

`func (o *KalturaBaseEntry) SetEntitledUsersPublish(v string)`

SetEntitledUsersPublish sets EntitledUsersPublish field to given value.

### HasEntitledUsersPublish

`func (o *KalturaBaseEntry) HasEntitledUsersPublish() bool`

HasEntitledUsersPublish returns a boolean if a field has been set.

### GetEntitledUsersView

`func (o *KalturaBaseEntry) GetEntitledUsersView() string`

GetEntitledUsersView returns the EntitledUsersView field if non-nil, zero value otherwise.

### GetEntitledUsersViewOk

`func (o *KalturaBaseEntry) GetEntitledUsersViewOk() (*string, bool)`

GetEntitledUsersViewOk returns a tuple with the EntitledUsersView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledUsersView

`func (o *KalturaBaseEntry) SetEntitledUsersView(v string)`

SetEntitledUsersView sets EntitledUsersView field to given value.

### HasEntitledUsersView

`func (o *KalturaBaseEntry) HasEntitledUsersView() bool`

HasEntitledUsersView returns a boolean if a field has been set.

### GetGroupId

`func (o *KalturaBaseEntry) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *KalturaBaseEntry) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *KalturaBaseEntry) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *KalturaBaseEntry) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *KalturaBaseEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBaseEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBaseEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBaseEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseType

`func (o *KalturaBaseEntry) GetLicenseType() int32`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *KalturaBaseEntry) GetLicenseTypeOk() (*int32, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *KalturaBaseEntry) SetLicenseType(v int32)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *KalturaBaseEntry) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetModerationCount

`func (o *KalturaBaseEntry) GetModerationCount() int32`

GetModerationCount returns the ModerationCount field if non-nil, zero value otherwise.

### GetModerationCountOk

`func (o *KalturaBaseEntry) GetModerationCountOk() (*int32, bool)`

GetModerationCountOk returns a tuple with the ModerationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationCount

`func (o *KalturaBaseEntry) SetModerationCount(v int32)`

SetModerationCount sets ModerationCount field to given value.

### HasModerationCount

`func (o *KalturaBaseEntry) HasModerationCount() bool`

HasModerationCount returns a boolean if a field has been set.

### GetModerationStatus

`func (o *KalturaBaseEntry) GetModerationStatus() int32`

GetModerationStatus returns the ModerationStatus field if non-nil, zero value otherwise.

### GetModerationStatusOk

`func (o *KalturaBaseEntry) GetModerationStatusOk() (*int32, bool)`

GetModerationStatusOk returns a tuple with the ModerationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationStatus

`func (o *KalturaBaseEntry) SetModerationStatus(v int32)`

SetModerationStatus sets ModerationStatus field to given value.

### HasModerationStatus

`func (o *KalturaBaseEntry) HasModerationStatus() bool`

HasModerationStatus returns a boolean if a field has been set.

### GetName

`func (o *KalturaBaseEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaBaseEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaBaseEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaBaseEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaBaseEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaBaseEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaBaseEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaBaseEntry) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOperationAttributes

`func (o *KalturaBaseEntry) GetOperationAttributes() []KalturaOperationAttributes`

GetOperationAttributes returns the OperationAttributes field if non-nil, zero value otherwise.

### GetOperationAttributesOk

`func (o *KalturaBaseEntry) GetOperationAttributesOk() (*[]KalturaOperationAttributes, bool)`

GetOperationAttributesOk returns a tuple with the OperationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAttributes

`func (o *KalturaBaseEntry) SetOperationAttributes(v []KalturaOperationAttributes)`

SetOperationAttributes sets OperationAttributes field to given value.

### HasOperationAttributes

`func (o *KalturaBaseEntry) HasOperationAttributes() bool`

HasOperationAttributes returns a boolean if a field has been set.

### GetParentEntryId

`func (o *KalturaBaseEntry) GetParentEntryId() string`

GetParentEntryId returns the ParentEntryId field if non-nil, zero value otherwise.

### GetParentEntryIdOk

`func (o *KalturaBaseEntry) GetParentEntryIdOk() (*string, bool)`

GetParentEntryIdOk returns a tuple with the ParentEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntryId

`func (o *KalturaBaseEntry) SetParentEntryId(v string)`

SetParentEntryId sets ParentEntryId field to given value.

### HasParentEntryId

`func (o *KalturaBaseEntry) HasParentEntryId() bool`

HasParentEntryId returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaBaseEntry) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaBaseEntry) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaBaseEntry) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaBaseEntry) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBaseEntry) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBaseEntry) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBaseEntry) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBaseEntry) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerSortValue

`func (o *KalturaBaseEntry) GetPartnerSortValue() int32`

GetPartnerSortValue returns the PartnerSortValue field if non-nil, zero value otherwise.

### GetPartnerSortValueOk

`func (o *KalturaBaseEntry) GetPartnerSortValueOk() (*int32, bool)`

GetPartnerSortValueOk returns a tuple with the PartnerSortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSortValue

`func (o *KalturaBaseEntry) SetPartnerSortValue(v int32)`

SetPartnerSortValue sets PartnerSortValue field to given value.

### HasPartnerSortValue

`func (o *KalturaBaseEntry) HasPartnerSortValue() bool`

HasPartnerSortValue returns a boolean if a field has been set.

### GetRank

`func (o *KalturaBaseEntry) GetRank() float32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *KalturaBaseEntry) GetRankOk() (*float32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *KalturaBaseEntry) SetRank(v float32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *KalturaBaseEntry) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRedirectEntryId

`func (o *KalturaBaseEntry) GetRedirectEntryId() string`

GetRedirectEntryId returns the RedirectEntryId field if non-nil, zero value otherwise.

### GetRedirectEntryIdOk

`func (o *KalturaBaseEntry) GetRedirectEntryIdOk() (*string, bool)`

GetRedirectEntryIdOk returns a tuple with the RedirectEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectEntryId

`func (o *KalturaBaseEntry) SetRedirectEntryId(v string)`

SetRedirectEntryId sets RedirectEntryId field to given value.

### HasRedirectEntryId

`func (o *KalturaBaseEntry) HasRedirectEntryId() bool`

HasRedirectEntryId returns a boolean if a field has been set.

### GetReferenceId

`func (o *KalturaBaseEntry) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KalturaBaseEntry) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KalturaBaseEntry) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KalturaBaseEntry) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReplacedEntryId

`func (o *KalturaBaseEntry) GetReplacedEntryId() string`

GetReplacedEntryId returns the ReplacedEntryId field if non-nil, zero value otherwise.

### GetReplacedEntryIdOk

`func (o *KalturaBaseEntry) GetReplacedEntryIdOk() (*string, bool)`

GetReplacedEntryIdOk returns a tuple with the ReplacedEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedEntryId

`func (o *KalturaBaseEntry) SetReplacedEntryId(v string)`

SetReplacedEntryId sets ReplacedEntryId field to given value.

### HasReplacedEntryId

`func (o *KalturaBaseEntry) HasReplacedEntryId() bool`

HasReplacedEntryId returns a boolean if a field has been set.

### GetReplacementStatus

`func (o *KalturaBaseEntry) GetReplacementStatus() string`

GetReplacementStatus returns the ReplacementStatus field if non-nil, zero value otherwise.

### GetReplacementStatusOk

`func (o *KalturaBaseEntry) GetReplacementStatusOk() (*string, bool)`

GetReplacementStatusOk returns a tuple with the ReplacementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementStatus

`func (o *KalturaBaseEntry) SetReplacementStatus(v string)`

SetReplacementStatus sets ReplacementStatus field to given value.

### HasReplacementStatus

`func (o *KalturaBaseEntry) HasReplacementStatus() bool`

HasReplacementStatus returns a boolean if a field has been set.

### GetReplacingEntryId

`func (o *KalturaBaseEntry) GetReplacingEntryId() string`

GetReplacingEntryId returns the ReplacingEntryId field if non-nil, zero value otherwise.

### GetReplacingEntryIdOk

`func (o *KalturaBaseEntry) GetReplacingEntryIdOk() (*string, bool)`

GetReplacingEntryIdOk returns a tuple with the ReplacingEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingEntryId

`func (o *KalturaBaseEntry) SetReplacingEntryId(v string)`

SetReplacingEntryId sets ReplacingEntryId field to given value.

### HasReplacingEntryId

`func (o *KalturaBaseEntry) HasReplacingEntryId() bool`

HasReplacingEntryId returns a boolean if a field has been set.

### GetRootEntryId

`func (o *KalturaBaseEntry) GetRootEntryId() string`

GetRootEntryId returns the RootEntryId field if non-nil, zero value otherwise.

### GetRootEntryIdOk

`func (o *KalturaBaseEntry) GetRootEntryIdOk() (*string, bool)`

GetRootEntryIdOk returns a tuple with the RootEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootEntryId

`func (o *KalturaBaseEntry) SetRootEntryId(v string)`

SetRootEntryId sets RootEntryId field to given value.

### HasRootEntryId

`func (o *KalturaBaseEntry) HasRootEntryId() bool`

HasRootEntryId returns a boolean if a field has been set.

### GetSearchText

`func (o *KalturaBaseEntry) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *KalturaBaseEntry) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *KalturaBaseEntry) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *KalturaBaseEntry) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetStartDate

`func (o *KalturaBaseEntry) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *KalturaBaseEntry) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *KalturaBaseEntry) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *KalturaBaseEntry) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBaseEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBaseEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBaseEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBaseEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaBaseEntry) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaBaseEntry) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaBaseEntry) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaBaseEntry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplateEntryId

`func (o *KalturaBaseEntry) GetTemplateEntryId() string`

GetTemplateEntryId returns the TemplateEntryId field if non-nil, zero value otherwise.

### GetTemplateEntryIdOk

`func (o *KalturaBaseEntry) GetTemplateEntryIdOk() (*string, bool)`

GetTemplateEntryIdOk returns a tuple with the TemplateEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEntryId

`func (o *KalturaBaseEntry) SetTemplateEntryId(v string)`

SetTemplateEntryId sets TemplateEntryId field to given value.

### HasTemplateEntryId

`func (o *KalturaBaseEntry) HasTemplateEntryId() bool`

HasTemplateEntryId returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *KalturaBaseEntry) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *KalturaBaseEntry) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *KalturaBaseEntry) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *KalturaBaseEntry) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetTotalRank

`func (o *KalturaBaseEntry) GetTotalRank() int32`

GetTotalRank returns the TotalRank field if non-nil, zero value otherwise.

### GetTotalRankOk

`func (o *KalturaBaseEntry) GetTotalRankOk() (*int32, bool)`

GetTotalRankOk returns a tuple with the TotalRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRank

`func (o *KalturaBaseEntry) SetTotalRank(v int32)`

SetTotalRank sets TotalRank field to given value.

### HasTotalRank

`func (o *KalturaBaseEntry) HasTotalRank() bool`

HasTotalRank returns a boolean if a field has been set.

### GetType

`func (o *KalturaBaseEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaBaseEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaBaseEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaBaseEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBaseEntry) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBaseEntry) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBaseEntry) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBaseEntry) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaBaseEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaBaseEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaBaseEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaBaseEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaBaseEntry) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaBaseEntry) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaBaseEntry) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaBaseEntry) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotes

`func (o *KalturaBaseEntry) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *KalturaBaseEntry) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *KalturaBaseEntry) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *KalturaBaseEntry) HasVotes() bool`

HasVotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


