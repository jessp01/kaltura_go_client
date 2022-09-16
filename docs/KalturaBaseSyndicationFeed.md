# KalturaBaseSyndicationFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddToDefaultConversionProfile** | Pointer to **bool** |  | [optional] 
**AllowEmbed** | Pointer to **bool** | allow_embed tells google OR yahoo weather to allow embedding the video on google OR yahoo video results  or just to provide a link to the landing page.  it is applied on the video-player_loc property in the XML (google)  and addes media-player tag (yahoo) | [optional] 
**Categories** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**EnforceEntitlement** | Pointer to **bool** | Should enforce entitlement on feed entries | [optional] 
**EntriesOrderBy** | Pointer to **string** | Enum Type: &#x60;KalturaSyndicationFeedEntriesOrderBy&#x60; | [optional] 
**FeedContentTypeHeader** | Pointer to **string** | Feed content-type header value | [optional] 
**FeedUrl** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FlavorParamId** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LandingPage** | Pointer to **string** | Base URL for each video, on the partners site  This is required by all syndication types. | [optional] 
**Name** | Pointer to **string** | feed name | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PlayerUiconfId** | Pointer to **int32** | Select a uiconf ID as player skin to include in the kwidget url | [optional] 
**PlaylistId** | Pointer to **string** | link a playlist that will set what content the feed will include  if empty, all content will be included in feed | [optional] 
**PrivacyContext** | Pointer to **string** | Set privacy context for search entries that assiged to private and public categories within a category privacy context. | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaSyndicationFeedStatus&#x60;  feed status | [optional] [readonly] 
**StorageId** | Pointer to **int32** |  | [optional] 
**TranscodeExistingContent** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **int32** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaSyndicationFeedType&#x60;  feed type | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UseCategoryEntries** | Pointer to **bool** |  | [optional] 

## Methods

### NewKalturaBaseSyndicationFeed

`func NewKalturaBaseSyndicationFeed() *KalturaBaseSyndicationFeed`

NewKalturaBaseSyndicationFeed instantiates a new KalturaBaseSyndicationFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBaseSyndicationFeedWithDefaults

`func NewKalturaBaseSyndicationFeedWithDefaults() *KalturaBaseSyndicationFeed`

NewKalturaBaseSyndicationFeedWithDefaults instantiates a new KalturaBaseSyndicationFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddToDefaultConversionProfile

`func (o *KalturaBaseSyndicationFeed) GetAddToDefaultConversionProfile() bool`

GetAddToDefaultConversionProfile returns the AddToDefaultConversionProfile field if non-nil, zero value otherwise.

### GetAddToDefaultConversionProfileOk

`func (o *KalturaBaseSyndicationFeed) GetAddToDefaultConversionProfileOk() (*bool, bool)`

GetAddToDefaultConversionProfileOk returns a tuple with the AddToDefaultConversionProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToDefaultConversionProfile

`func (o *KalturaBaseSyndicationFeed) SetAddToDefaultConversionProfile(v bool)`

SetAddToDefaultConversionProfile sets AddToDefaultConversionProfile field to given value.

### HasAddToDefaultConversionProfile

`func (o *KalturaBaseSyndicationFeed) HasAddToDefaultConversionProfile() bool`

HasAddToDefaultConversionProfile returns a boolean if a field has been set.

### GetAllowEmbed

`func (o *KalturaBaseSyndicationFeed) GetAllowEmbed() bool`

GetAllowEmbed returns the AllowEmbed field if non-nil, zero value otherwise.

### GetAllowEmbedOk

`func (o *KalturaBaseSyndicationFeed) GetAllowEmbedOk() (*bool, bool)`

GetAllowEmbedOk returns a tuple with the AllowEmbed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmbed

`func (o *KalturaBaseSyndicationFeed) SetAllowEmbed(v bool)`

SetAllowEmbed sets AllowEmbed field to given value.

### HasAllowEmbed

`func (o *KalturaBaseSyndicationFeed) HasAllowEmbed() bool`

HasAllowEmbed returns a boolean if a field has been set.

### GetCategories

`func (o *KalturaBaseSyndicationFeed) GetCategories() string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *KalturaBaseSyndicationFeed) GetCategoriesOk() (*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *KalturaBaseSyndicationFeed) SetCategories(v string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *KalturaBaseSyndicationFeed) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaBaseSyndicationFeed) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaBaseSyndicationFeed) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaBaseSyndicationFeed) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaBaseSyndicationFeed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnforceEntitlement

`func (o *KalturaBaseSyndicationFeed) GetEnforceEntitlement() bool`

GetEnforceEntitlement returns the EnforceEntitlement field if non-nil, zero value otherwise.

### GetEnforceEntitlementOk

`func (o *KalturaBaseSyndicationFeed) GetEnforceEntitlementOk() (*bool, bool)`

GetEnforceEntitlementOk returns a tuple with the EnforceEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceEntitlement

`func (o *KalturaBaseSyndicationFeed) SetEnforceEntitlement(v bool)`

SetEnforceEntitlement sets EnforceEntitlement field to given value.

### HasEnforceEntitlement

`func (o *KalturaBaseSyndicationFeed) HasEnforceEntitlement() bool`

HasEnforceEntitlement returns a boolean if a field has been set.

### GetEntriesOrderBy

`func (o *KalturaBaseSyndicationFeed) GetEntriesOrderBy() string`

GetEntriesOrderBy returns the EntriesOrderBy field if non-nil, zero value otherwise.

### GetEntriesOrderByOk

`func (o *KalturaBaseSyndicationFeed) GetEntriesOrderByOk() (*string, bool)`

GetEntriesOrderByOk returns a tuple with the EntriesOrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesOrderBy

`func (o *KalturaBaseSyndicationFeed) SetEntriesOrderBy(v string)`

SetEntriesOrderBy sets EntriesOrderBy field to given value.

### HasEntriesOrderBy

`func (o *KalturaBaseSyndicationFeed) HasEntriesOrderBy() bool`

HasEntriesOrderBy returns a boolean if a field has been set.

### GetFeedContentTypeHeader

`func (o *KalturaBaseSyndicationFeed) GetFeedContentTypeHeader() string`

GetFeedContentTypeHeader returns the FeedContentTypeHeader field if non-nil, zero value otherwise.

### GetFeedContentTypeHeaderOk

`func (o *KalturaBaseSyndicationFeed) GetFeedContentTypeHeaderOk() (*string, bool)`

GetFeedContentTypeHeaderOk returns a tuple with the FeedContentTypeHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedContentTypeHeader

`func (o *KalturaBaseSyndicationFeed) SetFeedContentTypeHeader(v string)`

SetFeedContentTypeHeader sets FeedContentTypeHeader field to given value.

### HasFeedContentTypeHeader

`func (o *KalturaBaseSyndicationFeed) HasFeedContentTypeHeader() bool`

HasFeedContentTypeHeader returns a boolean if a field has been set.

### GetFeedUrl

`func (o *KalturaBaseSyndicationFeed) GetFeedUrl() string`

GetFeedUrl returns the FeedUrl field if non-nil, zero value otherwise.

### GetFeedUrlOk

`func (o *KalturaBaseSyndicationFeed) GetFeedUrlOk() (*string, bool)`

GetFeedUrlOk returns a tuple with the FeedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedUrl

`func (o *KalturaBaseSyndicationFeed) SetFeedUrl(v string)`

SetFeedUrl sets FeedUrl field to given value.

### HasFeedUrl

`func (o *KalturaBaseSyndicationFeed) HasFeedUrl() bool`

HasFeedUrl returns a boolean if a field has been set.

### GetFlavorParamId

`func (o *KalturaBaseSyndicationFeed) GetFlavorParamId() int32`

GetFlavorParamId returns the FlavorParamId field if non-nil, zero value otherwise.

### GetFlavorParamIdOk

`func (o *KalturaBaseSyndicationFeed) GetFlavorParamIdOk() (*int32, bool)`

GetFlavorParamIdOk returns a tuple with the FlavorParamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamId

`func (o *KalturaBaseSyndicationFeed) SetFlavorParamId(v int32)`

SetFlavorParamId sets FlavorParamId field to given value.

### HasFlavorParamId

`func (o *KalturaBaseSyndicationFeed) HasFlavorParamId() bool`

HasFlavorParamId returns a boolean if a field has been set.

### GetId

`func (o *KalturaBaseSyndicationFeed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBaseSyndicationFeed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBaseSyndicationFeed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBaseSyndicationFeed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLandingPage

`func (o *KalturaBaseSyndicationFeed) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *KalturaBaseSyndicationFeed) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *KalturaBaseSyndicationFeed) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *KalturaBaseSyndicationFeed) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### GetName

`func (o *KalturaBaseSyndicationFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaBaseSyndicationFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaBaseSyndicationFeed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaBaseSyndicationFeed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaBaseSyndicationFeed) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaBaseSyndicationFeed) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaBaseSyndicationFeed) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaBaseSyndicationFeed) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBaseSyndicationFeed) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBaseSyndicationFeed) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBaseSyndicationFeed) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBaseSyndicationFeed) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPlayerUiconfId

`func (o *KalturaBaseSyndicationFeed) GetPlayerUiconfId() int32`

GetPlayerUiconfId returns the PlayerUiconfId field if non-nil, zero value otherwise.

### GetPlayerUiconfIdOk

`func (o *KalturaBaseSyndicationFeed) GetPlayerUiconfIdOk() (*int32, bool)`

GetPlayerUiconfIdOk returns a tuple with the PlayerUiconfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerUiconfId

`func (o *KalturaBaseSyndicationFeed) SetPlayerUiconfId(v int32)`

SetPlayerUiconfId sets PlayerUiconfId field to given value.

### HasPlayerUiconfId

`func (o *KalturaBaseSyndicationFeed) HasPlayerUiconfId() bool`

HasPlayerUiconfId returns a boolean if a field has been set.

### GetPlaylistId

`func (o *KalturaBaseSyndicationFeed) GetPlaylistId() string`

GetPlaylistId returns the PlaylistId field if non-nil, zero value otherwise.

### GetPlaylistIdOk

`func (o *KalturaBaseSyndicationFeed) GetPlaylistIdOk() (*string, bool)`

GetPlaylistIdOk returns a tuple with the PlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistId

`func (o *KalturaBaseSyndicationFeed) SetPlaylistId(v string)`

SetPlaylistId sets PlaylistId field to given value.

### HasPlaylistId

`func (o *KalturaBaseSyndicationFeed) HasPlaylistId() bool`

HasPlaylistId returns a boolean if a field has been set.

### GetPrivacyContext

`func (o *KalturaBaseSyndicationFeed) GetPrivacyContext() string`

GetPrivacyContext returns the PrivacyContext field if non-nil, zero value otherwise.

### GetPrivacyContextOk

`func (o *KalturaBaseSyndicationFeed) GetPrivacyContextOk() (*string, bool)`

GetPrivacyContextOk returns a tuple with the PrivacyContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyContext

`func (o *KalturaBaseSyndicationFeed) SetPrivacyContext(v string)`

SetPrivacyContext sets PrivacyContext field to given value.

### HasPrivacyContext

`func (o *KalturaBaseSyndicationFeed) HasPrivacyContext() bool`

HasPrivacyContext returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBaseSyndicationFeed) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBaseSyndicationFeed) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBaseSyndicationFeed) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBaseSyndicationFeed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageId

`func (o *KalturaBaseSyndicationFeed) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *KalturaBaseSyndicationFeed) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *KalturaBaseSyndicationFeed) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *KalturaBaseSyndicationFeed) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetTranscodeExistingContent

`func (o *KalturaBaseSyndicationFeed) GetTranscodeExistingContent() bool`

GetTranscodeExistingContent returns the TranscodeExistingContent field if non-nil, zero value otherwise.

### GetTranscodeExistingContentOk

`func (o *KalturaBaseSyndicationFeed) GetTranscodeExistingContentOk() (*bool, bool)`

GetTranscodeExistingContentOk returns a tuple with the TranscodeExistingContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeExistingContent

`func (o *KalturaBaseSyndicationFeed) SetTranscodeExistingContent(v bool)`

SetTranscodeExistingContent sets TranscodeExistingContent field to given value.

### HasTranscodeExistingContent

`func (o *KalturaBaseSyndicationFeed) HasTranscodeExistingContent() bool`

HasTranscodeExistingContent returns a boolean if a field has been set.

### GetType

`func (o *KalturaBaseSyndicationFeed) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaBaseSyndicationFeed) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaBaseSyndicationFeed) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaBaseSyndicationFeed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBaseSyndicationFeed) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBaseSyndicationFeed) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBaseSyndicationFeed) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBaseSyndicationFeed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseCategoryEntries

`func (o *KalturaBaseSyndicationFeed) GetUseCategoryEntries() bool`

GetUseCategoryEntries returns the UseCategoryEntries field if non-nil, zero value otherwise.

### GetUseCategoryEntriesOk

`func (o *KalturaBaseSyndicationFeed) GetUseCategoryEntriesOk() (*bool, bool)`

GetUseCategoryEntriesOk returns a tuple with the UseCategoryEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCategoryEntries

`func (o *KalturaBaseSyndicationFeed) SetUseCategoryEntries(v bool)`

SetUseCategoryEntries sets UseCategoryEntries field to given value.

### HasUseCategoryEntries

`func (o *KalturaBaseSyndicationFeed) HasUseCategoryEntries() bool`

HasUseCategoryEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


