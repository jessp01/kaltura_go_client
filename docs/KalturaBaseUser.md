# KalturaBaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminTags** | Pointer to **string** | Admin tags can be updated only by using an admin session | [optional] 
**AllowedPartnerIds** | Pointer to **string** |  | [optional] 
**AllowedPartnerPackages** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**DeletedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IndexedPartnerDataInt** | Pointer to **int32** |  | [optional] 
**IndexedPartnerDataString** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | Enum Type: &#x60;KalturaLanguageCode&#x60; | [optional] 
**LastLoginTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerData** | Pointer to **string** | Can be used to store various partner related data as a string | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ScreenName** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaUserStatus&#x60; | [optional] 
**StatusUpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**StorageSize** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Tags** | Pointer to **string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UserMode** | Pointer to **int32** | Enum Type: &#x60;KalturaUserMode&#x60; | [optional] 
**Zip** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaBaseUser

`func NewKalturaBaseUser() *KalturaBaseUser`

NewKalturaBaseUser instantiates a new KalturaBaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBaseUserWithDefaults

`func NewKalturaBaseUserWithDefaults() *KalturaBaseUser`

NewKalturaBaseUserWithDefaults instantiates a new KalturaBaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminTags

`func (o *KalturaBaseUser) GetAdminTags() string`

GetAdminTags returns the AdminTags field if non-nil, zero value otherwise.

### GetAdminTagsOk

`func (o *KalturaBaseUser) GetAdminTagsOk() (*string, bool)`

GetAdminTagsOk returns a tuple with the AdminTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTags

`func (o *KalturaBaseUser) SetAdminTags(v string)`

SetAdminTags sets AdminTags field to given value.

### HasAdminTags

`func (o *KalturaBaseUser) HasAdminTags() bool`

HasAdminTags returns a boolean if a field has been set.

### GetAllowedPartnerIds

`func (o *KalturaBaseUser) GetAllowedPartnerIds() string`

GetAllowedPartnerIds returns the AllowedPartnerIds field if non-nil, zero value otherwise.

### GetAllowedPartnerIdsOk

`func (o *KalturaBaseUser) GetAllowedPartnerIdsOk() (*string, bool)`

GetAllowedPartnerIdsOk returns a tuple with the AllowedPartnerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPartnerIds

`func (o *KalturaBaseUser) SetAllowedPartnerIds(v string)`

SetAllowedPartnerIds sets AllowedPartnerIds field to given value.

### HasAllowedPartnerIds

`func (o *KalturaBaseUser) HasAllowedPartnerIds() bool`

HasAllowedPartnerIds returns a boolean if a field has been set.

### GetAllowedPartnerPackages

`func (o *KalturaBaseUser) GetAllowedPartnerPackages() string`

GetAllowedPartnerPackages returns the AllowedPartnerPackages field if non-nil, zero value otherwise.

### GetAllowedPartnerPackagesOk

`func (o *KalturaBaseUser) GetAllowedPartnerPackagesOk() (*string, bool)`

GetAllowedPartnerPackagesOk returns a tuple with the AllowedPartnerPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPartnerPackages

`func (o *KalturaBaseUser) SetAllowedPartnerPackages(v string)`

SetAllowedPartnerPackages sets AllowedPartnerPackages field to given value.

### HasAllowedPartnerPackages

`func (o *KalturaBaseUser) HasAllowedPartnerPackages() bool`

HasAllowedPartnerPackages returns a boolean if a field has been set.

### GetCity

`func (o *KalturaBaseUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *KalturaBaseUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *KalturaBaseUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *KalturaBaseUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *KalturaBaseUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KalturaBaseUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KalturaBaseUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *KalturaBaseUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaBaseUser) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaBaseUser) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaBaseUser) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaBaseUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *KalturaBaseUser) GetDeletedAt() int32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *KalturaBaseUser) GetDeletedAtOk() (*int32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *KalturaBaseUser) SetDeletedAt(v int32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *KalturaBaseUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaBaseUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaBaseUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaBaseUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaBaseUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *KalturaBaseUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *KalturaBaseUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *KalturaBaseUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *KalturaBaseUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *KalturaBaseUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *KalturaBaseUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *KalturaBaseUser) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *KalturaBaseUser) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *KalturaBaseUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBaseUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBaseUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBaseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexedPartnerDataInt

`func (o *KalturaBaseUser) GetIndexedPartnerDataInt() int32`

GetIndexedPartnerDataInt returns the IndexedPartnerDataInt field if non-nil, zero value otherwise.

### GetIndexedPartnerDataIntOk

`func (o *KalturaBaseUser) GetIndexedPartnerDataIntOk() (*int32, bool)`

GetIndexedPartnerDataIntOk returns a tuple with the IndexedPartnerDataInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedPartnerDataInt

`func (o *KalturaBaseUser) SetIndexedPartnerDataInt(v int32)`

SetIndexedPartnerDataInt sets IndexedPartnerDataInt field to given value.

### HasIndexedPartnerDataInt

`func (o *KalturaBaseUser) HasIndexedPartnerDataInt() bool`

HasIndexedPartnerDataInt returns a boolean if a field has been set.

### GetIndexedPartnerDataString

`func (o *KalturaBaseUser) GetIndexedPartnerDataString() string`

GetIndexedPartnerDataString returns the IndexedPartnerDataString field if non-nil, zero value otherwise.

### GetIndexedPartnerDataStringOk

`func (o *KalturaBaseUser) GetIndexedPartnerDataStringOk() (*string, bool)`

GetIndexedPartnerDataStringOk returns a tuple with the IndexedPartnerDataString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedPartnerDataString

`func (o *KalturaBaseUser) SetIndexedPartnerDataString(v string)`

SetIndexedPartnerDataString sets IndexedPartnerDataString field to given value.

### HasIndexedPartnerDataString

`func (o *KalturaBaseUser) HasIndexedPartnerDataString() bool`

HasIndexedPartnerDataString returns a boolean if a field has been set.

### GetLanguage

`func (o *KalturaBaseUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *KalturaBaseUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *KalturaBaseUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *KalturaBaseUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *KalturaBaseUser) GetLastLoginTime() int32`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *KalturaBaseUser) GetLastLoginTimeOk() (*int32, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *KalturaBaseUser) SetLastLoginTime(v int32)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *KalturaBaseUser) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaBaseUser) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaBaseUser) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaBaseUser) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaBaseUser) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaBaseUser) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaBaseUser) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaBaseUser) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaBaseUser) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaBaseUser) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaBaseUser) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaBaseUser) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaBaseUser) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetScreenName

`func (o *KalturaBaseUser) GetScreenName() string`

GetScreenName returns the ScreenName field if non-nil, zero value otherwise.

### GetScreenNameOk

`func (o *KalturaBaseUser) GetScreenNameOk() (*string, bool)`

GetScreenNameOk returns a tuple with the ScreenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenName

`func (o *KalturaBaseUser) SetScreenName(v string)`

SetScreenName sets ScreenName field to given value.

### HasScreenName

`func (o *KalturaBaseUser) HasScreenName() bool`

HasScreenName returns a boolean if a field has been set.

### GetState

`func (o *KalturaBaseUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KalturaBaseUser) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KalturaBaseUser) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KalturaBaseUser) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaBaseUser) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaBaseUser) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaBaseUser) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaBaseUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *KalturaBaseUser) GetStatusUpdatedAt() int32`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *KalturaBaseUser) GetStatusUpdatedAtOk() (*int32, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *KalturaBaseUser) SetStatusUpdatedAt(v int32)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *KalturaBaseUser) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetStorageSize

`func (o *KalturaBaseUser) GetStorageSize() int32`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *KalturaBaseUser) GetStorageSizeOk() (*int32, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *KalturaBaseUser) SetStorageSize(v int32)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *KalturaBaseUser) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetTags

`func (o *KalturaBaseUser) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaBaseUser) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaBaseUser) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaBaseUser) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *KalturaBaseUser) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *KalturaBaseUser) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *KalturaBaseUser) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *KalturaBaseUser) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBaseUser) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBaseUser) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBaseUser) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBaseUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserMode

`func (o *KalturaBaseUser) GetUserMode() int32`

GetUserMode returns the UserMode field if non-nil, zero value otherwise.

### GetUserModeOk

`func (o *KalturaBaseUser) GetUserModeOk() (*int32, bool)`

GetUserModeOk returns a tuple with the UserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMode

`func (o *KalturaBaseUser) SetUserMode(v int32)`

SetUserMode sets UserMode field to given value.

### HasUserMode

`func (o *KalturaBaseUser) HasUserMode() bool`

HasUserMode returns a boolean if a field has been set.

### GetZip

`func (o *KalturaBaseUser) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *KalturaBaseUser) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *KalturaBaseUser) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *KalturaBaseUser) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


