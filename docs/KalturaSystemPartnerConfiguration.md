# KalturaSystemPartnerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmail** | Pointer to **string** |  | [optional] 
**AdminName** | Pointer to **string** |  | [optional] 
**AdminSessionRoleId** | Pointer to **int32** |  | [optional] 
**AllowMultiNotification** | Pointer to **bool** |  | [optional] 
**AllowedDomains** | Pointer to **string** |  | [optional] 
**AllowedFromEmailWhiteList** | Pointer to **string** |  | [optional] 
**AlwaysAllowedPermissionNames** | Pointer to **string** |  | [optional] 
**ApiAccessControlId** | Pointer to **int32** |  | [optional] 
**AudioThumbEntryId** | Pointer to **string** |  | [optional] 
**AutoModerateEntryFilter** | Pointer to [**KalturaBaseEntryFilter**](KalturaBaseEntryFilter.md) |  | [optional] 
**AvoidIndexingSearchHistory** | Pointer to **bool** |  | [optional] 
**BlockDirectLogin** | Pointer to **bool** |  | [optional] 
**BulkUploadNotificationsEmail** | Pointer to **string** |  | [optional] 
**CacheFlavorVersion** | Pointer to **int32** |  | [optional] 
**CdnHost** | Pointer to **string** |  | [optional] 
**CdnHostWhiteList** | Pointer to **string** |  | [optional] 
**CrmId** | Pointer to **string** |  | [optional] 
**CrmLink** | Pointer to **string** |  | [optional] 
**CustomDeliveryTypes** | Pointer to [**[]KalturaKeyBooleanValue**](KalturaKeyBooleanValue.md) |  | [optional] 
**DefThumbDensity** | Pointer to **int32** |  | [optional] 
**DefThumbOffset** | Pointer to **int32** |  | [optional] 
**DefaultDeliveryType** | Pointer to **string** |  | [optional] 
**DefaultEmbedCodeType** | Pointer to **string** |  | [optional] 
**DefaultEntitlementEnforcement** | Pointer to **bool** |  | [optional] 
**DefaultLiveStreamEntrySourceType** | Pointer to **string** | Enum Type: &#x60;KalturaSourceType&#x60; | [optional] 
**DefaultLiveStreamSegmentDuration** | Pointer to **int32** |  | [optional] 
**DefaultRecordingConversionProfile** | Pointer to **int32** |  | [optional] 
**DeliveryProfileIds** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ESearchLanguages** | Pointer to **string** |  | [optional] 
**EightyPercentWarning** | Pointer to **int32** |  | [optional] 
**EnableBulkUploadNotificationsEmails** | Pointer to **bool** |  | [optional] 
**EnableSelfServe** | Pointer to **bool** |  | [optional] 
**EnforceDelivery** | Pointer to **bool** |  | [optional] 
**EnforceHttpsApi** | Pointer to **bool** |  | [optional] 
**ExcludedAdminRoleName** | Pointer to **string** |  | [optional] 
**ExtendedFreeTrail** | Pointer to **int32** |  | [optional] 
**ExtendedFreeTrailEndsWarning** | Pointer to **bool** |  | [optional] 
**ExtendedFreeTrailExpiryDate** | Pointer to **int32** | Unix timestamp (In seconds) | [optional] 
**ExtendedFreeTrailExpiryReason** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**HtmlPurifierBaseListUsage** | Pointer to **bool** |  | [optional] 
**HtmlPurifierBehaviour** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IgnoreSynonymEsearch** | Pointer to **bool** |  | [optional] 
**ImportRemoteSourceForConvert** | Pointer to **bool** |  | [optional] 
**InternalUse** | Pointer to **bool** |  | [optional] 
**IsFirstLogin** | Pointer to **bool** |  | [optional] 
**IsSelfServe** | Pointer to **bool** |  | [optional] 
**KmcVersion** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to **string** | Enum Type: &#x60;KalturaLanguageCode&#x60; | [optional] 
**LastFreeTrialNotificationDay** | Pointer to **int32** |  | [optional] 
**Limits** | Pointer to [**[]KalturaSystemPartnerLimit**](KalturaSystemPartnerLimit.md) |  | [optional] 
**LiveDeliveryProfileIds** | Pointer to **string** |  | [optional] 
**LiveStreamProvisionParams** | Pointer to **string** |  | [optional] 
**LiveThumbEntryId** | Pointer to **string** |  | [optional] 
**LoginBlockPeriod** | Pointer to **int32** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**MaxWordForNgram** | Pointer to **int32** |  | [optional] 
**MediaProtocol** | Pointer to **string** | http/https, rtmp/rtmpe | [optional] 
**ModerateContent** | Pointer to **bool** |  | [optional] 
**MonitorUsage** | Pointer to **int32** |  | [optional] 
**NotificationsConfig** | Pointer to **string** |  | [optional] 
**NumPrevPassToKeep** | Pointer to **int32** |  | [optional] 
**OttEnvironmentUrl** | Pointer to **string** |  | [optional] 
**OvpEnvironmentUrl** | Pointer to **string** |  | [optional] 
**PartnerGroupType** | Pointer to **int32** | Enum Type: &#x60;KalturaPartnerGroupType&#x60; | [optional] 
**PartnerName** | Pointer to **string** |  | [optional] 
**PartnerPackage** | Pointer to **int32** |  | [optional] 
**PartnerPackageClassOfService** | Pointer to **string** |  | [optional] 
**PartnerParentId** | Pointer to **int32** |  | [optional] 
**PassReplaceFreq** | Pointer to **int32** |  | [optional] 
**PasswordStructureValidations** | Pointer to **string** |  | [optional] 
**PasswordStructureValidationsDescription** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]KalturaPermission**](KalturaPermission.md) |  | [optional] 
**PublisherEnvironmentType** | Pointer to **int32** |  | [optional] 
**PurifyImageContent** | Pointer to **bool** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**RestrictEntryByMetadata** | Pointer to **bool** |  | [optional] 
**RestrictThumbnailByKs** | Pointer to **int32** |  | [optional] 
**SecondarySecretRoleId** | Pointer to **int32** |  | [optional] 
**StorageDeleteFromKaltura** | Pointer to **bool** |  | [optional] 
**StorageServePriority** | Pointer to **int32** | Enum Type: &#x60;KalturaStorageServePriority&#x60; | [optional] 
**StreamerType** | Pointer to **string** | http/rtmp/hdnetwork | [optional] 
**SupportAnimatedThumbnails** | Pointer to **bool** |  | [optional] 
**ThumbnailHost** | Pointer to **string** |  | [optional] 
**TimeAlignedRenditions** | Pointer to **bool** |  | [optional] 
**TrigramPercentage** | Pointer to **int32** |  | [optional] 
**TwoFactorAuthenticationMode** | Pointer to **int32** | Enum Type: &#x60;KalturaTwoFactorAuthenticationMode&#x60; | [optional] 
**UsageLimitWarning** | Pointer to **int32** |  | [optional] 
**UsagePercent** | Pointer to **int32** |  | [optional] 
**UseSso** | Pointer to **bool** |  | [optional] 
**UseTwoFactorAuthentication** | Pointer to **bool** |  | [optional] 
**UserSessionRoleId** | Pointer to **int32** |  | [optional] 
**VerticalClasiffication** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaSystemPartnerConfiguration

`func NewKalturaSystemPartnerConfiguration() *KalturaSystemPartnerConfiguration`

NewKalturaSystemPartnerConfiguration instantiates a new KalturaSystemPartnerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSystemPartnerConfigurationWithDefaults

`func NewKalturaSystemPartnerConfigurationWithDefaults() *KalturaSystemPartnerConfiguration`

NewKalturaSystemPartnerConfigurationWithDefaults instantiates a new KalturaSystemPartnerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmail

`func (o *KalturaSystemPartnerConfiguration) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *KalturaSystemPartnerConfiguration) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *KalturaSystemPartnerConfiguration) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *KalturaSystemPartnerConfiguration) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminName

`func (o *KalturaSystemPartnerConfiguration) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *KalturaSystemPartnerConfiguration) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *KalturaSystemPartnerConfiguration) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *KalturaSystemPartnerConfiguration) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) GetAdminSessionRoleId() int32`

GetAdminSessionRoleId returns the AdminSessionRoleId field if non-nil, zero value otherwise.

### GetAdminSessionRoleIdOk

`func (o *KalturaSystemPartnerConfiguration) GetAdminSessionRoleIdOk() (*int32, bool)`

GetAdminSessionRoleIdOk returns a tuple with the AdminSessionRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) SetAdminSessionRoleId(v int32)`

SetAdminSessionRoleId sets AdminSessionRoleId field to given value.

### HasAdminSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) HasAdminSessionRoleId() bool`

HasAdminSessionRoleId returns a boolean if a field has been set.

### GetAllowMultiNotification

`func (o *KalturaSystemPartnerConfiguration) GetAllowMultiNotification() bool`

GetAllowMultiNotification returns the AllowMultiNotification field if non-nil, zero value otherwise.

### GetAllowMultiNotificationOk

`func (o *KalturaSystemPartnerConfiguration) GetAllowMultiNotificationOk() (*bool, bool)`

GetAllowMultiNotificationOk returns a tuple with the AllowMultiNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiNotification

`func (o *KalturaSystemPartnerConfiguration) SetAllowMultiNotification(v bool)`

SetAllowMultiNotification sets AllowMultiNotification field to given value.

### HasAllowMultiNotification

`func (o *KalturaSystemPartnerConfiguration) HasAllowMultiNotification() bool`

HasAllowMultiNotification returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *KalturaSystemPartnerConfiguration) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *KalturaSystemPartnerConfiguration) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *KalturaSystemPartnerConfiguration) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *KalturaSystemPartnerConfiguration) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedFromEmailWhiteList

`func (o *KalturaSystemPartnerConfiguration) GetAllowedFromEmailWhiteList() string`

GetAllowedFromEmailWhiteList returns the AllowedFromEmailWhiteList field if non-nil, zero value otherwise.

### GetAllowedFromEmailWhiteListOk

`func (o *KalturaSystemPartnerConfiguration) GetAllowedFromEmailWhiteListOk() (*string, bool)`

GetAllowedFromEmailWhiteListOk returns a tuple with the AllowedFromEmailWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFromEmailWhiteList

`func (o *KalturaSystemPartnerConfiguration) SetAllowedFromEmailWhiteList(v string)`

SetAllowedFromEmailWhiteList sets AllowedFromEmailWhiteList field to given value.

### HasAllowedFromEmailWhiteList

`func (o *KalturaSystemPartnerConfiguration) HasAllowedFromEmailWhiteList() bool`

HasAllowedFromEmailWhiteList returns a boolean if a field has been set.

### GetAlwaysAllowedPermissionNames

`func (o *KalturaSystemPartnerConfiguration) GetAlwaysAllowedPermissionNames() string`

GetAlwaysAllowedPermissionNames returns the AlwaysAllowedPermissionNames field if non-nil, zero value otherwise.

### GetAlwaysAllowedPermissionNamesOk

`func (o *KalturaSystemPartnerConfiguration) GetAlwaysAllowedPermissionNamesOk() (*string, bool)`

GetAlwaysAllowedPermissionNamesOk returns a tuple with the AlwaysAllowedPermissionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysAllowedPermissionNames

`func (o *KalturaSystemPartnerConfiguration) SetAlwaysAllowedPermissionNames(v string)`

SetAlwaysAllowedPermissionNames sets AlwaysAllowedPermissionNames field to given value.

### HasAlwaysAllowedPermissionNames

`func (o *KalturaSystemPartnerConfiguration) HasAlwaysAllowedPermissionNames() bool`

HasAlwaysAllowedPermissionNames returns a boolean if a field has been set.

### GetApiAccessControlId

`func (o *KalturaSystemPartnerConfiguration) GetApiAccessControlId() int32`

GetApiAccessControlId returns the ApiAccessControlId field if non-nil, zero value otherwise.

### GetApiAccessControlIdOk

`func (o *KalturaSystemPartnerConfiguration) GetApiAccessControlIdOk() (*int32, bool)`

GetApiAccessControlIdOk returns a tuple with the ApiAccessControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessControlId

`func (o *KalturaSystemPartnerConfiguration) SetApiAccessControlId(v int32)`

SetApiAccessControlId sets ApiAccessControlId field to given value.

### HasApiAccessControlId

`func (o *KalturaSystemPartnerConfiguration) HasApiAccessControlId() bool`

HasApiAccessControlId returns a boolean if a field has been set.

### GetAudioThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) GetAudioThumbEntryId() string`

GetAudioThumbEntryId returns the AudioThumbEntryId field if non-nil, zero value otherwise.

### GetAudioThumbEntryIdOk

`func (o *KalturaSystemPartnerConfiguration) GetAudioThumbEntryIdOk() (*string, bool)`

GetAudioThumbEntryIdOk returns a tuple with the AudioThumbEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) SetAudioThumbEntryId(v string)`

SetAudioThumbEntryId sets AudioThumbEntryId field to given value.

### HasAudioThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) HasAudioThumbEntryId() bool`

HasAudioThumbEntryId returns a boolean if a field has been set.

### GetAutoModerateEntryFilter

`func (o *KalturaSystemPartnerConfiguration) GetAutoModerateEntryFilter() KalturaBaseEntryFilter`

GetAutoModerateEntryFilter returns the AutoModerateEntryFilter field if non-nil, zero value otherwise.

### GetAutoModerateEntryFilterOk

`func (o *KalturaSystemPartnerConfiguration) GetAutoModerateEntryFilterOk() (*KalturaBaseEntryFilter, bool)`

GetAutoModerateEntryFilterOk returns a tuple with the AutoModerateEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoModerateEntryFilter

`func (o *KalturaSystemPartnerConfiguration) SetAutoModerateEntryFilter(v KalturaBaseEntryFilter)`

SetAutoModerateEntryFilter sets AutoModerateEntryFilter field to given value.

### HasAutoModerateEntryFilter

`func (o *KalturaSystemPartnerConfiguration) HasAutoModerateEntryFilter() bool`

HasAutoModerateEntryFilter returns a boolean if a field has been set.

### GetAvoidIndexingSearchHistory

`func (o *KalturaSystemPartnerConfiguration) GetAvoidIndexingSearchHistory() bool`

GetAvoidIndexingSearchHistory returns the AvoidIndexingSearchHistory field if non-nil, zero value otherwise.

### GetAvoidIndexingSearchHistoryOk

`func (o *KalturaSystemPartnerConfiguration) GetAvoidIndexingSearchHistoryOk() (*bool, bool)`

GetAvoidIndexingSearchHistoryOk returns a tuple with the AvoidIndexingSearchHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidIndexingSearchHistory

`func (o *KalturaSystemPartnerConfiguration) SetAvoidIndexingSearchHistory(v bool)`

SetAvoidIndexingSearchHistory sets AvoidIndexingSearchHistory field to given value.

### HasAvoidIndexingSearchHistory

`func (o *KalturaSystemPartnerConfiguration) HasAvoidIndexingSearchHistory() bool`

HasAvoidIndexingSearchHistory returns a boolean if a field has been set.

### GetBlockDirectLogin

`func (o *KalturaSystemPartnerConfiguration) GetBlockDirectLogin() bool`

GetBlockDirectLogin returns the BlockDirectLogin field if non-nil, zero value otherwise.

### GetBlockDirectLoginOk

`func (o *KalturaSystemPartnerConfiguration) GetBlockDirectLoginOk() (*bool, bool)`

GetBlockDirectLoginOk returns a tuple with the BlockDirectLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDirectLogin

`func (o *KalturaSystemPartnerConfiguration) SetBlockDirectLogin(v bool)`

SetBlockDirectLogin sets BlockDirectLogin field to given value.

### HasBlockDirectLogin

`func (o *KalturaSystemPartnerConfiguration) HasBlockDirectLogin() bool`

HasBlockDirectLogin returns a boolean if a field has been set.

### GetBulkUploadNotificationsEmail

`func (o *KalturaSystemPartnerConfiguration) GetBulkUploadNotificationsEmail() string`

GetBulkUploadNotificationsEmail returns the BulkUploadNotificationsEmail field if non-nil, zero value otherwise.

### GetBulkUploadNotificationsEmailOk

`func (o *KalturaSystemPartnerConfiguration) GetBulkUploadNotificationsEmailOk() (*string, bool)`

GetBulkUploadNotificationsEmailOk returns a tuple with the BulkUploadNotificationsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadNotificationsEmail

`func (o *KalturaSystemPartnerConfiguration) SetBulkUploadNotificationsEmail(v string)`

SetBulkUploadNotificationsEmail sets BulkUploadNotificationsEmail field to given value.

### HasBulkUploadNotificationsEmail

`func (o *KalturaSystemPartnerConfiguration) HasBulkUploadNotificationsEmail() bool`

HasBulkUploadNotificationsEmail returns a boolean if a field has been set.

### GetCacheFlavorVersion

`func (o *KalturaSystemPartnerConfiguration) GetCacheFlavorVersion() int32`

GetCacheFlavorVersion returns the CacheFlavorVersion field if non-nil, zero value otherwise.

### GetCacheFlavorVersionOk

`func (o *KalturaSystemPartnerConfiguration) GetCacheFlavorVersionOk() (*int32, bool)`

GetCacheFlavorVersionOk returns a tuple with the CacheFlavorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlavorVersion

`func (o *KalturaSystemPartnerConfiguration) SetCacheFlavorVersion(v int32)`

SetCacheFlavorVersion sets CacheFlavorVersion field to given value.

### HasCacheFlavorVersion

`func (o *KalturaSystemPartnerConfiguration) HasCacheFlavorVersion() bool`

HasCacheFlavorVersion returns a boolean if a field has been set.

### GetCdnHost

`func (o *KalturaSystemPartnerConfiguration) GetCdnHost() string`

GetCdnHost returns the CdnHost field if non-nil, zero value otherwise.

### GetCdnHostOk

`func (o *KalturaSystemPartnerConfiguration) GetCdnHostOk() (*string, bool)`

GetCdnHostOk returns a tuple with the CdnHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnHost

`func (o *KalturaSystemPartnerConfiguration) SetCdnHost(v string)`

SetCdnHost sets CdnHost field to given value.

### HasCdnHost

`func (o *KalturaSystemPartnerConfiguration) HasCdnHost() bool`

HasCdnHost returns a boolean if a field has been set.

### GetCdnHostWhiteList

`func (o *KalturaSystemPartnerConfiguration) GetCdnHostWhiteList() string`

GetCdnHostWhiteList returns the CdnHostWhiteList field if non-nil, zero value otherwise.

### GetCdnHostWhiteListOk

`func (o *KalturaSystemPartnerConfiguration) GetCdnHostWhiteListOk() (*string, bool)`

GetCdnHostWhiteListOk returns a tuple with the CdnHostWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnHostWhiteList

`func (o *KalturaSystemPartnerConfiguration) SetCdnHostWhiteList(v string)`

SetCdnHostWhiteList sets CdnHostWhiteList field to given value.

### HasCdnHostWhiteList

`func (o *KalturaSystemPartnerConfiguration) HasCdnHostWhiteList() bool`

HasCdnHostWhiteList returns a boolean if a field has been set.

### GetCrmId

`func (o *KalturaSystemPartnerConfiguration) GetCrmId() string`

GetCrmId returns the CrmId field if non-nil, zero value otherwise.

### GetCrmIdOk

`func (o *KalturaSystemPartnerConfiguration) GetCrmIdOk() (*string, bool)`

GetCrmIdOk returns a tuple with the CrmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmId

`func (o *KalturaSystemPartnerConfiguration) SetCrmId(v string)`

SetCrmId sets CrmId field to given value.

### HasCrmId

`func (o *KalturaSystemPartnerConfiguration) HasCrmId() bool`

HasCrmId returns a boolean if a field has been set.

### GetCrmLink

`func (o *KalturaSystemPartnerConfiguration) GetCrmLink() string`

GetCrmLink returns the CrmLink field if non-nil, zero value otherwise.

### GetCrmLinkOk

`func (o *KalturaSystemPartnerConfiguration) GetCrmLinkOk() (*string, bool)`

GetCrmLinkOk returns a tuple with the CrmLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmLink

`func (o *KalturaSystemPartnerConfiguration) SetCrmLink(v string)`

SetCrmLink sets CrmLink field to given value.

### HasCrmLink

`func (o *KalturaSystemPartnerConfiguration) HasCrmLink() bool`

HasCrmLink returns a boolean if a field has been set.

### GetCustomDeliveryTypes

`func (o *KalturaSystemPartnerConfiguration) GetCustomDeliveryTypes() []KalturaKeyBooleanValue`

GetCustomDeliveryTypes returns the CustomDeliveryTypes field if non-nil, zero value otherwise.

### GetCustomDeliveryTypesOk

`func (o *KalturaSystemPartnerConfiguration) GetCustomDeliveryTypesOk() (*[]KalturaKeyBooleanValue, bool)`

GetCustomDeliveryTypesOk returns a tuple with the CustomDeliveryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDeliveryTypes

`func (o *KalturaSystemPartnerConfiguration) SetCustomDeliveryTypes(v []KalturaKeyBooleanValue)`

SetCustomDeliveryTypes sets CustomDeliveryTypes field to given value.

### HasCustomDeliveryTypes

`func (o *KalturaSystemPartnerConfiguration) HasCustomDeliveryTypes() bool`

HasCustomDeliveryTypes returns a boolean if a field has been set.

### GetDefThumbDensity

`func (o *KalturaSystemPartnerConfiguration) GetDefThumbDensity() int32`

GetDefThumbDensity returns the DefThumbDensity field if non-nil, zero value otherwise.

### GetDefThumbDensityOk

`func (o *KalturaSystemPartnerConfiguration) GetDefThumbDensityOk() (*int32, bool)`

GetDefThumbDensityOk returns a tuple with the DefThumbDensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefThumbDensity

`func (o *KalturaSystemPartnerConfiguration) SetDefThumbDensity(v int32)`

SetDefThumbDensity sets DefThumbDensity field to given value.

### HasDefThumbDensity

`func (o *KalturaSystemPartnerConfiguration) HasDefThumbDensity() bool`

HasDefThumbDensity returns a boolean if a field has been set.

### GetDefThumbOffset

`func (o *KalturaSystemPartnerConfiguration) GetDefThumbOffset() int32`

GetDefThumbOffset returns the DefThumbOffset field if non-nil, zero value otherwise.

### GetDefThumbOffsetOk

`func (o *KalturaSystemPartnerConfiguration) GetDefThumbOffsetOk() (*int32, bool)`

GetDefThumbOffsetOk returns a tuple with the DefThumbOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefThumbOffset

`func (o *KalturaSystemPartnerConfiguration) SetDefThumbOffset(v int32)`

SetDefThumbOffset sets DefThumbOffset field to given value.

### HasDefThumbOffset

`func (o *KalturaSystemPartnerConfiguration) HasDefThumbOffset() bool`

HasDefThumbOffset returns a boolean if a field has been set.

### GetDefaultDeliveryType

`func (o *KalturaSystemPartnerConfiguration) GetDefaultDeliveryType() string`

GetDefaultDeliveryType returns the DefaultDeliveryType field if non-nil, zero value otherwise.

### GetDefaultDeliveryTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultDeliveryTypeOk() (*string, bool)`

GetDefaultDeliveryTypeOk returns a tuple with the DefaultDeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeliveryType

`func (o *KalturaSystemPartnerConfiguration) SetDefaultDeliveryType(v string)`

SetDefaultDeliveryType sets DefaultDeliveryType field to given value.

### HasDefaultDeliveryType

`func (o *KalturaSystemPartnerConfiguration) HasDefaultDeliveryType() bool`

HasDefaultDeliveryType returns a boolean if a field has been set.

### GetDefaultEmbedCodeType

`func (o *KalturaSystemPartnerConfiguration) GetDefaultEmbedCodeType() string`

GetDefaultEmbedCodeType returns the DefaultEmbedCodeType field if non-nil, zero value otherwise.

### GetDefaultEmbedCodeTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultEmbedCodeTypeOk() (*string, bool)`

GetDefaultEmbedCodeTypeOk returns a tuple with the DefaultEmbedCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmbedCodeType

`func (o *KalturaSystemPartnerConfiguration) SetDefaultEmbedCodeType(v string)`

SetDefaultEmbedCodeType sets DefaultEmbedCodeType field to given value.

### HasDefaultEmbedCodeType

`func (o *KalturaSystemPartnerConfiguration) HasDefaultEmbedCodeType() bool`

HasDefaultEmbedCodeType returns a boolean if a field has been set.

### GetDefaultEntitlementEnforcement

`func (o *KalturaSystemPartnerConfiguration) GetDefaultEntitlementEnforcement() bool`

GetDefaultEntitlementEnforcement returns the DefaultEntitlementEnforcement field if non-nil, zero value otherwise.

### GetDefaultEntitlementEnforcementOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultEntitlementEnforcementOk() (*bool, bool)`

GetDefaultEntitlementEnforcementOk returns a tuple with the DefaultEntitlementEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntitlementEnforcement

`func (o *KalturaSystemPartnerConfiguration) SetDefaultEntitlementEnforcement(v bool)`

SetDefaultEntitlementEnforcement sets DefaultEntitlementEnforcement field to given value.

### HasDefaultEntitlementEnforcement

`func (o *KalturaSystemPartnerConfiguration) HasDefaultEntitlementEnforcement() bool`

HasDefaultEntitlementEnforcement returns a boolean if a field has been set.

### GetDefaultLiveStreamEntrySourceType

`func (o *KalturaSystemPartnerConfiguration) GetDefaultLiveStreamEntrySourceType() string`

GetDefaultLiveStreamEntrySourceType returns the DefaultLiveStreamEntrySourceType field if non-nil, zero value otherwise.

### GetDefaultLiveStreamEntrySourceTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultLiveStreamEntrySourceTypeOk() (*string, bool)`

GetDefaultLiveStreamEntrySourceTypeOk returns a tuple with the DefaultLiveStreamEntrySourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLiveStreamEntrySourceType

`func (o *KalturaSystemPartnerConfiguration) SetDefaultLiveStreamEntrySourceType(v string)`

SetDefaultLiveStreamEntrySourceType sets DefaultLiveStreamEntrySourceType field to given value.

### HasDefaultLiveStreamEntrySourceType

`func (o *KalturaSystemPartnerConfiguration) HasDefaultLiveStreamEntrySourceType() bool`

HasDefaultLiveStreamEntrySourceType returns a boolean if a field has been set.

### GetDefaultLiveStreamSegmentDuration

`func (o *KalturaSystemPartnerConfiguration) GetDefaultLiveStreamSegmentDuration() int32`

GetDefaultLiveStreamSegmentDuration returns the DefaultLiveStreamSegmentDuration field if non-nil, zero value otherwise.

### GetDefaultLiveStreamSegmentDurationOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultLiveStreamSegmentDurationOk() (*int32, bool)`

GetDefaultLiveStreamSegmentDurationOk returns a tuple with the DefaultLiveStreamSegmentDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLiveStreamSegmentDuration

`func (o *KalturaSystemPartnerConfiguration) SetDefaultLiveStreamSegmentDuration(v int32)`

SetDefaultLiveStreamSegmentDuration sets DefaultLiveStreamSegmentDuration field to given value.

### HasDefaultLiveStreamSegmentDuration

`func (o *KalturaSystemPartnerConfiguration) HasDefaultLiveStreamSegmentDuration() bool`

HasDefaultLiveStreamSegmentDuration returns a boolean if a field has been set.

### GetDefaultRecordingConversionProfile

`func (o *KalturaSystemPartnerConfiguration) GetDefaultRecordingConversionProfile() int32`

GetDefaultRecordingConversionProfile returns the DefaultRecordingConversionProfile field if non-nil, zero value otherwise.

### GetDefaultRecordingConversionProfileOk

`func (o *KalturaSystemPartnerConfiguration) GetDefaultRecordingConversionProfileOk() (*int32, bool)`

GetDefaultRecordingConversionProfileOk returns a tuple with the DefaultRecordingConversionProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRecordingConversionProfile

`func (o *KalturaSystemPartnerConfiguration) SetDefaultRecordingConversionProfile(v int32)`

SetDefaultRecordingConversionProfile sets DefaultRecordingConversionProfile field to given value.

### HasDefaultRecordingConversionProfile

`func (o *KalturaSystemPartnerConfiguration) HasDefaultRecordingConversionProfile() bool`

HasDefaultRecordingConversionProfile returns a boolean if a field has been set.

### GetDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) GetDeliveryProfileIds() string`

GetDeliveryProfileIds returns the DeliveryProfileIds field if non-nil, zero value otherwise.

### GetDeliveryProfileIdsOk

`func (o *KalturaSystemPartnerConfiguration) GetDeliveryProfileIdsOk() (*string, bool)`

GetDeliveryProfileIdsOk returns a tuple with the DeliveryProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) SetDeliveryProfileIds(v string)`

SetDeliveryProfileIds sets DeliveryProfileIds field to given value.

### HasDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) HasDeliveryProfileIds() bool`

HasDeliveryProfileIds returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaSystemPartnerConfiguration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaSystemPartnerConfiguration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaSystemPartnerConfiguration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaSystemPartnerConfiguration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetESearchLanguages

`func (o *KalturaSystemPartnerConfiguration) GetESearchLanguages() string`

GetESearchLanguages returns the ESearchLanguages field if non-nil, zero value otherwise.

### GetESearchLanguagesOk

`func (o *KalturaSystemPartnerConfiguration) GetESearchLanguagesOk() (*string, bool)`

GetESearchLanguagesOk returns a tuple with the ESearchLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESearchLanguages

`func (o *KalturaSystemPartnerConfiguration) SetESearchLanguages(v string)`

SetESearchLanguages sets ESearchLanguages field to given value.

### HasESearchLanguages

`func (o *KalturaSystemPartnerConfiguration) HasESearchLanguages() bool`

HasESearchLanguages returns a boolean if a field has been set.

### GetEightyPercentWarning

`func (o *KalturaSystemPartnerConfiguration) GetEightyPercentWarning() int32`

GetEightyPercentWarning returns the EightyPercentWarning field if non-nil, zero value otherwise.

### GetEightyPercentWarningOk

`func (o *KalturaSystemPartnerConfiguration) GetEightyPercentWarningOk() (*int32, bool)`

GetEightyPercentWarningOk returns a tuple with the EightyPercentWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEightyPercentWarning

`func (o *KalturaSystemPartnerConfiguration) SetEightyPercentWarning(v int32)`

SetEightyPercentWarning sets EightyPercentWarning field to given value.

### HasEightyPercentWarning

`func (o *KalturaSystemPartnerConfiguration) HasEightyPercentWarning() bool`

HasEightyPercentWarning returns a boolean if a field has been set.

### GetEnableBulkUploadNotificationsEmails

`func (o *KalturaSystemPartnerConfiguration) GetEnableBulkUploadNotificationsEmails() bool`

GetEnableBulkUploadNotificationsEmails returns the EnableBulkUploadNotificationsEmails field if non-nil, zero value otherwise.

### GetEnableBulkUploadNotificationsEmailsOk

`func (o *KalturaSystemPartnerConfiguration) GetEnableBulkUploadNotificationsEmailsOk() (*bool, bool)`

GetEnableBulkUploadNotificationsEmailsOk returns a tuple with the EnableBulkUploadNotificationsEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBulkUploadNotificationsEmails

`func (o *KalturaSystemPartnerConfiguration) SetEnableBulkUploadNotificationsEmails(v bool)`

SetEnableBulkUploadNotificationsEmails sets EnableBulkUploadNotificationsEmails field to given value.

### HasEnableBulkUploadNotificationsEmails

`func (o *KalturaSystemPartnerConfiguration) HasEnableBulkUploadNotificationsEmails() bool`

HasEnableBulkUploadNotificationsEmails returns a boolean if a field has been set.

### GetEnableSelfServe

`func (o *KalturaSystemPartnerConfiguration) GetEnableSelfServe() bool`

GetEnableSelfServe returns the EnableSelfServe field if non-nil, zero value otherwise.

### GetEnableSelfServeOk

`func (o *KalturaSystemPartnerConfiguration) GetEnableSelfServeOk() (*bool, bool)`

GetEnableSelfServeOk returns a tuple with the EnableSelfServe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSelfServe

`func (o *KalturaSystemPartnerConfiguration) SetEnableSelfServe(v bool)`

SetEnableSelfServe sets EnableSelfServe field to given value.

### HasEnableSelfServe

`func (o *KalturaSystemPartnerConfiguration) HasEnableSelfServe() bool`

HasEnableSelfServe returns a boolean if a field has been set.

### GetEnforceDelivery

`func (o *KalturaSystemPartnerConfiguration) GetEnforceDelivery() bool`

GetEnforceDelivery returns the EnforceDelivery field if non-nil, zero value otherwise.

### GetEnforceDeliveryOk

`func (o *KalturaSystemPartnerConfiguration) GetEnforceDeliveryOk() (*bool, bool)`

GetEnforceDeliveryOk returns a tuple with the EnforceDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDelivery

`func (o *KalturaSystemPartnerConfiguration) SetEnforceDelivery(v bool)`

SetEnforceDelivery sets EnforceDelivery field to given value.

### HasEnforceDelivery

`func (o *KalturaSystemPartnerConfiguration) HasEnforceDelivery() bool`

HasEnforceDelivery returns a boolean if a field has been set.

### GetEnforceHttpsApi

`func (o *KalturaSystemPartnerConfiguration) GetEnforceHttpsApi() bool`

GetEnforceHttpsApi returns the EnforceHttpsApi field if non-nil, zero value otherwise.

### GetEnforceHttpsApiOk

`func (o *KalturaSystemPartnerConfiguration) GetEnforceHttpsApiOk() (*bool, bool)`

GetEnforceHttpsApiOk returns a tuple with the EnforceHttpsApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceHttpsApi

`func (o *KalturaSystemPartnerConfiguration) SetEnforceHttpsApi(v bool)`

SetEnforceHttpsApi sets EnforceHttpsApi field to given value.

### HasEnforceHttpsApi

`func (o *KalturaSystemPartnerConfiguration) HasEnforceHttpsApi() bool`

HasEnforceHttpsApi returns a boolean if a field has been set.

### GetExcludedAdminRoleName

`func (o *KalturaSystemPartnerConfiguration) GetExcludedAdminRoleName() string`

GetExcludedAdminRoleName returns the ExcludedAdminRoleName field if non-nil, zero value otherwise.

### GetExcludedAdminRoleNameOk

`func (o *KalturaSystemPartnerConfiguration) GetExcludedAdminRoleNameOk() (*string, bool)`

GetExcludedAdminRoleNameOk returns a tuple with the ExcludedAdminRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAdminRoleName

`func (o *KalturaSystemPartnerConfiguration) SetExcludedAdminRoleName(v string)`

SetExcludedAdminRoleName sets ExcludedAdminRoleName field to given value.

### HasExcludedAdminRoleName

`func (o *KalturaSystemPartnerConfiguration) HasExcludedAdminRoleName() bool`

HasExcludedAdminRoleName returns a boolean if a field has been set.

### GetExtendedFreeTrail

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrail() int32`

GetExtendedFreeTrail returns the ExtendedFreeTrail field if non-nil, zero value otherwise.

### GetExtendedFreeTrailOk

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailOk() (*int32, bool)`

GetExtendedFreeTrailOk returns a tuple with the ExtendedFreeTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrail

`func (o *KalturaSystemPartnerConfiguration) SetExtendedFreeTrail(v int32)`

SetExtendedFreeTrail sets ExtendedFreeTrail field to given value.

### HasExtendedFreeTrail

`func (o *KalturaSystemPartnerConfiguration) HasExtendedFreeTrail() bool`

HasExtendedFreeTrail returns a boolean if a field has been set.

### GetExtendedFreeTrailEndsWarning

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailEndsWarning() bool`

GetExtendedFreeTrailEndsWarning returns the ExtendedFreeTrailEndsWarning field if non-nil, zero value otherwise.

### GetExtendedFreeTrailEndsWarningOk

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailEndsWarningOk() (*bool, bool)`

GetExtendedFreeTrailEndsWarningOk returns a tuple with the ExtendedFreeTrailEndsWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailEndsWarning

`func (o *KalturaSystemPartnerConfiguration) SetExtendedFreeTrailEndsWarning(v bool)`

SetExtendedFreeTrailEndsWarning sets ExtendedFreeTrailEndsWarning field to given value.

### HasExtendedFreeTrailEndsWarning

`func (o *KalturaSystemPartnerConfiguration) HasExtendedFreeTrailEndsWarning() bool`

HasExtendedFreeTrailEndsWarning returns a boolean if a field has been set.

### GetExtendedFreeTrailExpiryDate

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailExpiryDate() int32`

GetExtendedFreeTrailExpiryDate returns the ExtendedFreeTrailExpiryDate field if non-nil, zero value otherwise.

### GetExtendedFreeTrailExpiryDateOk

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailExpiryDateOk() (*int32, bool)`

GetExtendedFreeTrailExpiryDateOk returns a tuple with the ExtendedFreeTrailExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailExpiryDate

`func (o *KalturaSystemPartnerConfiguration) SetExtendedFreeTrailExpiryDate(v int32)`

SetExtendedFreeTrailExpiryDate sets ExtendedFreeTrailExpiryDate field to given value.

### HasExtendedFreeTrailExpiryDate

`func (o *KalturaSystemPartnerConfiguration) HasExtendedFreeTrailExpiryDate() bool`

HasExtendedFreeTrailExpiryDate returns a boolean if a field has been set.

### GetExtendedFreeTrailExpiryReason

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailExpiryReason() string`

GetExtendedFreeTrailExpiryReason returns the ExtendedFreeTrailExpiryReason field if non-nil, zero value otherwise.

### GetExtendedFreeTrailExpiryReasonOk

`func (o *KalturaSystemPartnerConfiguration) GetExtendedFreeTrailExpiryReasonOk() (*string, bool)`

GetExtendedFreeTrailExpiryReasonOk returns a tuple with the ExtendedFreeTrailExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailExpiryReason

`func (o *KalturaSystemPartnerConfiguration) SetExtendedFreeTrailExpiryReason(v string)`

SetExtendedFreeTrailExpiryReason sets ExtendedFreeTrailExpiryReason field to given value.

### HasExtendedFreeTrailExpiryReason

`func (o *KalturaSystemPartnerConfiguration) HasExtendedFreeTrailExpiryReason() bool`

HasExtendedFreeTrailExpiryReason returns a boolean if a field has been set.

### GetHost

`func (o *KalturaSystemPartnerConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KalturaSystemPartnerConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KalturaSystemPartnerConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KalturaSystemPartnerConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHtmlPurifierBaseListUsage

`func (o *KalturaSystemPartnerConfiguration) GetHtmlPurifierBaseListUsage() bool`

GetHtmlPurifierBaseListUsage returns the HtmlPurifierBaseListUsage field if non-nil, zero value otherwise.

### GetHtmlPurifierBaseListUsageOk

`func (o *KalturaSystemPartnerConfiguration) GetHtmlPurifierBaseListUsageOk() (*bool, bool)`

GetHtmlPurifierBaseListUsageOk returns a tuple with the HtmlPurifierBaseListUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlPurifierBaseListUsage

`func (o *KalturaSystemPartnerConfiguration) SetHtmlPurifierBaseListUsage(v bool)`

SetHtmlPurifierBaseListUsage sets HtmlPurifierBaseListUsage field to given value.

### HasHtmlPurifierBaseListUsage

`func (o *KalturaSystemPartnerConfiguration) HasHtmlPurifierBaseListUsage() bool`

HasHtmlPurifierBaseListUsage returns a boolean if a field has been set.

### GetHtmlPurifierBehaviour

`func (o *KalturaSystemPartnerConfiguration) GetHtmlPurifierBehaviour() int32`

GetHtmlPurifierBehaviour returns the HtmlPurifierBehaviour field if non-nil, zero value otherwise.

### GetHtmlPurifierBehaviourOk

`func (o *KalturaSystemPartnerConfiguration) GetHtmlPurifierBehaviourOk() (*int32, bool)`

GetHtmlPurifierBehaviourOk returns a tuple with the HtmlPurifierBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlPurifierBehaviour

`func (o *KalturaSystemPartnerConfiguration) SetHtmlPurifierBehaviour(v int32)`

SetHtmlPurifierBehaviour sets HtmlPurifierBehaviour field to given value.

### HasHtmlPurifierBehaviour

`func (o *KalturaSystemPartnerConfiguration) HasHtmlPurifierBehaviour() bool`

HasHtmlPurifierBehaviour returns a boolean if a field has been set.

### GetId

`func (o *KalturaSystemPartnerConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaSystemPartnerConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaSystemPartnerConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaSystemPartnerConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreSynonymEsearch

`func (o *KalturaSystemPartnerConfiguration) GetIgnoreSynonymEsearch() bool`

GetIgnoreSynonymEsearch returns the IgnoreSynonymEsearch field if non-nil, zero value otherwise.

### GetIgnoreSynonymEsearchOk

`func (o *KalturaSystemPartnerConfiguration) GetIgnoreSynonymEsearchOk() (*bool, bool)`

GetIgnoreSynonymEsearchOk returns a tuple with the IgnoreSynonymEsearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSynonymEsearch

`func (o *KalturaSystemPartnerConfiguration) SetIgnoreSynonymEsearch(v bool)`

SetIgnoreSynonymEsearch sets IgnoreSynonymEsearch field to given value.

### HasIgnoreSynonymEsearch

`func (o *KalturaSystemPartnerConfiguration) HasIgnoreSynonymEsearch() bool`

HasIgnoreSynonymEsearch returns a boolean if a field has been set.

### GetImportRemoteSourceForConvert

`func (o *KalturaSystemPartnerConfiguration) GetImportRemoteSourceForConvert() bool`

GetImportRemoteSourceForConvert returns the ImportRemoteSourceForConvert field if non-nil, zero value otherwise.

### GetImportRemoteSourceForConvertOk

`func (o *KalturaSystemPartnerConfiguration) GetImportRemoteSourceForConvertOk() (*bool, bool)`

GetImportRemoteSourceForConvertOk returns a tuple with the ImportRemoteSourceForConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportRemoteSourceForConvert

`func (o *KalturaSystemPartnerConfiguration) SetImportRemoteSourceForConvert(v bool)`

SetImportRemoteSourceForConvert sets ImportRemoteSourceForConvert field to given value.

### HasImportRemoteSourceForConvert

`func (o *KalturaSystemPartnerConfiguration) HasImportRemoteSourceForConvert() bool`

HasImportRemoteSourceForConvert returns a boolean if a field has been set.

### GetInternalUse

`func (o *KalturaSystemPartnerConfiguration) GetInternalUse() bool`

GetInternalUse returns the InternalUse field if non-nil, zero value otherwise.

### GetInternalUseOk

`func (o *KalturaSystemPartnerConfiguration) GetInternalUseOk() (*bool, bool)`

GetInternalUseOk returns a tuple with the InternalUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalUse

`func (o *KalturaSystemPartnerConfiguration) SetInternalUse(v bool)`

SetInternalUse sets InternalUse field to given value.

### HasInternalUse

`func (o *KalturaSystemPartnerConfiguration) HasInternalUse() bool`

HasInternalUse returns a boolean if a field has been set.

### GetIsFirstLogin

`func (o *KalturaSystemPartnerConfiguration) GetIsFirstLogin() bool`

GetIsFirstLogin returns the IsFirstLogin field if non-nil, zero value otherwise.

### GetIsFirstLoginOk

`func (o *KalturaSystemPartnerConfiguration) GetIsFirstLoginOk() (*bool, bool)`

GetIsFirstLoginOk returns a tuple with the IsFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstLogin

`func (o *KalturaSystemPartnerConfiguration) SetIsFirstLogin(v bool)`

SetIsFirstLogin sets IsFirstLogin field to given value.

### HasIsFirstLogin

`func (o *KalturaSystemPartnerConfiguration) HasIsFirstLogin() bool`

HasIsFirstLogin returns a boolean if a field has been set.

### GetIsSelfServe

`func (o *KalturaSystemPartnerConfiguration) GetIsSelfServe() bool`

GetIsSelfServe returns the IsSelfServe field if non-nil, zero value otherwise.

### GetIsSelfServeOk

`func (o *KalturaSystemPartnerConfiguration) GetIsSelfServeOk() (*bool, bool)`

GetIsSelfServeOk returns a tuple with the IsSelfServe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfServe

`func (o *KalturaSystemPartnerConfiguration) SetIsSelfServe(v bool)`

SetIsSelfServe sets IsSelfServe field to given value.

### HasIsSelfServe

`func (o *KalturaSystemPartnerConfiguration) HasIsSelfServe() bool`

HasIsSelfServe returns a boolean if a field has been set.

### GetKmcVersion

`func (o *KalturaSystemPartnerConfiguration) GetKmcVersion() int32`

GetKmcVersion returns the KmcVersion field if non-nil, zero value otherwise.

### GetKmcVersionOk

`func (o *KalturaSystemPartnerConfiguration) GetKmcVersionOk() (*int32, bool)`

GetKmcVersionOk returns a tuple with the KmcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmcVersion

`func (o *KalturaSystemPartnerConfiguration) SetKmcVersion(v int32)`

SetKmcVersion sets KmcVersion field to given value.

### HasKmcVersion

`func (o *KalturaSystemPartnerConfiguration) HasKmcVersion() bool`

HasKmcVersion returns a boolean if a field has been set.

### GetLanguage

`func (o *KalturaSystemPartnerConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *KalturaSystemPartnerConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *KalturaSystemPartnerConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *KalturaSystemPartnerConfiguration) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastFreeTrialNotificationDay

`func (o *KalturaSystemPartnerConfiguration) GetLastFreeTrialNotificationDay() int32`

GetLastFreeTrialNotificationDay returns the LastFreeTrialNotificationDay field if non-nil, zero value otherwise.

### GetLastFreeTrialNotificationDayOk

`func (o *KalturaSystemPartnerConfiguration) GetLastFreeTrialNotificationDayOk() (*int32, bool)`

GetLastFreeTrialNotificationDayOk returns a tuple with the LastFreeTrialNotificationDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFreeTrialNotificationDay

`func (o *KalturaSystemPartnerConfiguration) SetLastFreeTrialNotificationDay(v int32)`

SetLastFreeTrialNotificationDay sets LastFreeTrialNotificationDay field to given value.

### HasLastFreeTrialNotificationDay

`func (o *KalturaSystemPartnerConfiguration) HasLastFreeTrialNotificationDay() bool`

HasLastFreeTrialNotificationDay returns a boolean if a field has been set.

### GetLimits

`func (o *KalturaSystemPartnerConfiguration) GetLimits() []KalturaSystemPartnerLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KalturaSystemPartnerConfiguration) GetLimitsOk() (*[]KalturaSystemPartnerLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KalturaSystemPartnerConfiguration) SetLimits(v []KalturaSystemPartnerLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KalturaSystemPartnerConfiguration) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetLiveDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) GetLiveDeliveryProfileIds() string`

GetLiveDeliveryProfileIds returns the LiveDeliveryProfileIds field if non-nil, zero value otherwise.

### GetLiveDeliveryProfileIdsOk

`func (o *KalturaSystemPartnerConfiguration) GetLiveDeliveryProfileIdsOk() (*string, bool)`

GetLiveDeliveryProfileIdsOk returns a tuple with the LiveDeliveryProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) SetLiveDeliveryProfileIds(v string)`

SetLiveDeliveryProfileIds sets LiveDeliveryProfileIds field to given value.

### HasLiveDeliveryProfileIds

`func (o *KalturaSystemPartnerConfiguration) HasLiveDeliveryProfileIds() bool`

HasLiveDeliveryProfileIds returns a boolean if a field has been set.

### GetLiveStreamProvisionParams

`func (o *KalturaSystemPartnerConfiguration) GetLiveStreamProvisionParams() string`

GetLiveStreamProvisionParams returns the LiveStreamProvisionParams field if non-nil, zero value otherwise.

### GetLiveStreamProvisionParamsOk

`func (o *KalturaSystemPartnerConfiguration) GetLiveStreamProvisionParamsOk() (*string, bool)`

GetLiveStreamProvisionParamsOk returns a tuple with the LiveStreamProvisionParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamProvisionParams

`func (o *KalturaSystemPartnerConfiguration) SetLiveStreamProvisionParams(v string)`

SetLiveStreamProvisionParams sets LiveStreamProvisionParams field to given value.

### HasLiveStreamProvisionParams

`func (o *KalturaSystemPartnerConfiguration) HasLiveStreamProvisionParams() bool`

HasLiveStreamProvisionParams returns a boolean if a field has been set.

### GetLiveThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) GetLiveThumbEntryId() string`

GetLiveThumbEntryId returns the LiveThumbEntryId field if non-nil, zero value otherwise.

### GetLiveThumbEntryIdOk

`func (o *KalturaSystemPartnerConfiguration) GetLiveThumbEntryIdOk() (*string, bool)`

GetLiveThumbEntryIdOk returns a tuple with the LiveThumbEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) SetLiveThumbEntryId(v string)`

SetLiveThumbEntryId sets LiveThumbEntryId field to given value.

### HasLiveThumbEntryId

`func (o *KalturaSystemPartnerConfiguration) HasLiveThumbEntryId() bool`

HasLiveThumbEntryId returns a boolean if a field has been set.

### GetLoginBlockPeriod

`func (o *KalturaSystemPartnerConfiguration) GetLoginBlockPeriod() int32`

GetLoginBlockPeriod returns the LoginBlockPeriod field if non-nil, zero value otherwise.

### GetLoginBlockPeriodOk

`func (o *KalturaSystemPartnerConfiguration) GetLoginBlockPeriodOk() (*int32, bool)`

GetLoginBlockPeriodOk returns a tuple with the LoginBlockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBlockPeriod

`func (o *KalturaSystemPartnerConfiguration) SetLoginBlockPeriod(v int32)`

SetLoginBlockPeriod sets LoginBlockPeriod field to given value.

### HasLoginBlockPeriod

`func (o *KalturaSystemPartnerConfiguration) HasLoginBlockPeriod() bool`

HasLoginBlockPeriod returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *KalturaSystemPartnerConfiguration) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *KalturaSystemPartnerConfiguration) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *KalturaSystemPartnerConfiguration) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *KalturaSystemPartnerConfiguration) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetMaxWordForNgram

`func (o *KalturaSystemPartnerConfiguration) GetMaxWordForNgram() int32`

GetMaxWordForNgram returns the MaxWordForNgram field if non-nil, zero value otherwise.

### GetMaxWordForNgramOk

`func (o *KalturaSystemPartnerConfiguration) GetMaxWordForNgramOk() (*int32, bool)`

GetMaxWordForNgramOk returns a tuple with the MaxWordForNgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWordForNgram

`func (o *KalturaSystemPartnerConfiguration) SetMaxWordForNgram(v int32)`

SetMaxWordForNgram sets MaxWordForNgram field to given value.

### HasMaxWordForNgram

`func (o *KalturaSystemPartnerConfiguration) HasMaxWordForNgram() bool`

HasMaxWordForNgram returns a boolean if a field has been set.

### GetMediaProtocol

`func (o *KalturaSystemPartnerConfiguration) GetMediaProtocol() string`

GetMediaProtocol returns the MediaProtocol field if non-nil, zero value otherwise.

### GetMediaProtocolOk

`func (o *KalturaSystemPartnerConfiguration) GetMediaProtocolOk() (*string, bool)`

GetMediaProtocolOk returns a tuple with the MediaProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaProtocol

`func (o *KalturaSystemPartnerConfiguration) SetMediaProtocol(v string)`

SetMediaProtocol sets MediaProtocol field to given value.

### HasMediaProtocol

`func (o *KalturaSystemPartnerConfiguration) HasMediaProtocol() bool`

HasMediaProtocol returns a boolean if a field has been set.

### GetModerateContent

`func (o *KalturaSystemPartnerConfiguration) GetModerateContent() bool`

GetModerateContent returns the ModerateContent field if non-nil, zero value otherwise.

### GetModerateContentOk

`func (o *KalturaSystemPartnerConfiguration) GetModerateContentOk() (*bool, bool)`

GetModerateContentOk returns a tuple with the ModerateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerateContent

`func (o *KalturaSystemPartnerConfiguration) SetModerateContent(v bool)`

SetModerateContent sets ModerateContent field to given value.

### HasModerateContent

`func (o *KalturaSystemPartnerConfiguration) HasModerateContent() bool`

HasModerateContent returns a boolean if a field has been set.

### GetMonitorUsage

`func (o *KalturaSystemPartnerConfiguration) GetMonitorUsage() int32`

GetMonitorUsage returns the MonitorUsage field if non-nil, zero value otherwise.

### GetMonitorUsageOk

`func (o *KalturaSystemPartnerConfiguration) GetMonitorUsageOk() (*int32, bool)`

GetMonitorUsageOk returns a tuple with the MonitorUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsage

`func (o *KalturaSystemPartnerConfiguration) SetMonitorUsage(v int32)`

SetMonitorUsage sets MonitorUsage field to given value.

### HasMonitorUsage

`func (o *KalturaSystemPartnerConfiguration) HasMonitorUsage() bool`

HasMonitorUsage returns a boolean if a field has been set.

### GetNotificationsConfig

`func (o *KalturaSystemPartnerConfiguration) GetNotificationsConfig() string`

GetNotificationsConfig returns the NotificationsConfig field if non-nil, zero value otherwise.

### GetNotificationsConfigOk

`func (o *KalturaSystemPartnerConfiguration) GetNotificationsConfigOk() (*string, bool)`

GetNotificationsConfigOk returns a tuple with the NotificationsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsConfig

`func (o *KalturaSystemPartnerConfiguration) SetNotificationsConfig(v string)`

SetNotificationsConfig sets NotificationsConfig field to given value.

### HasNotificationsConfig

`func (o *KalturaSystemPartnerConfiguration) HasNotificationsConfig() bool`

HasNotificationsConfig returns a boolean if a field has been set.

### GetNumPrevPassToKeep

`func (o *KalturaSystemPartnerConfiguration) GetNumPrevPassToKeep() int32`

GetNumPrevPassToKeep returns the NumPrevPassToKeep field if non-nil, zero value otherwise.

### GetNumPrevPassToKeepOk

`func (o *KalturaSystemPartnerConfiguration) GetNumPrevPassToKeepOk() (*int32, bool)`

GetNumPrevPassToKeepOk returns a tuple with the NumPrevPassToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPrevPassToKeep

`func (o *KalturaSystemPartnerConfiguration) SetNumPrevPassToKeep(v int32)`

SetNumPrevPassToKeep sets NumPrevPassToKeep field to given value.

### HasNumPrevPassToKeep

`func (o *KalturaSystemPartnerConfiguration) HasNumPrevPassToKeep() bool`

HasNumPrevPassToKeep returns a boolean if a field has been set.

### GetOttEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) GetOttEnvironmentUrl() string`

GetOttEnvironmentUrl returns the OttEnvironmentUrl field if non-nil, zero value otherwise.

### GetOttEnvironmentUrlOk

`func (o *KalturaSystemPartnerConfiguration) GetOttEnvironmentUrlOk() (*string, bool)`

GetOttEnvironmentUrlOk returns a tuple with the OttEnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) SetOttEnvironmentUrl(v string)`

SetOttEnvironmentUrl sets OttEnvironmentUrl field to given value.

### HasOttEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) HasOttEnvironmentUrl() bool`

HasOttEnvironmentUrl returns a boolean if a field has been set.

### GetOvpEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) GetOvpEnvironmentUrl() string`

GetOvpEnvironmentUrl returns the OvpEnvironmentUrl field if non-nil, zero value otherwise.

### GetOvpEnvironmentUrlOk

`func (o *KalturaSystemPartnerConfiguration) GetOvpEnvironmentUrlOk() (*string, bool)`

GetOvpEnvironmentUrlOk returns a tuple with the OvpEnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvpEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) SetOvpEnvironmentUrl(v string)`

SetOvpEnvironmentUrl sets OvpEnvironmentUrl field to given value.

### HasOvpEnvironmentUrl

`func (o *KalturaSystemPartnerConfiguration) HasOvpEnvironmentUrl() bool`

HasOvpEnvironmentUrl returns a boolean if a field has been set.

### GetPartnerGroupType

`func (o *KalturaSystemPartnerConfiguration) GetPartnerGroupType() int32`

GetPartnerGroupType returns the PartnerGroupType field if non-nil, zero value otherwise.

### GetPartnerGroupTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetPartnerGroupTypeOk() (*int32, bool)`

GetPartnerGroupTypeOk returns a tuple with the PartnerGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerGroupType

`func (o *KalturaSystemPartnerConfiguration) SetPartnerGroupType(v int32)`

SetPartnerGroupType sets PartnerGroupType field to given value.

### HasPartnerGroupType

`func (o *KalturaSystemPartnerConfiguration) HasPartnerGroupType() bool`

HasPartnerGroupType returns a boolean if a field has been set.

### GetPartnerName

`func (o *KalturaSystemPartnerConfiguration) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *KalturaSystemPartnerConfiguration) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *KalturaSystemPartnerConfiguration) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *KalturaSystemPartnerConfiguration) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### GetPartnerPackage

`func (o *KalturaSystemPartnerConfiguration) GetPartnerPackage() int32`

GetPartnerPackage returns the PartnerPackage field if non-nil, zero value otherwise.

### GetPartnerPackageOk

`func (o *KalturaSystemPartnerConfiguration) GetPartnerPackageOk() (*int32, bool)`

GetPartnerPackageOk returns a tuple with the PartnerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPackage

`func (o *KalturaSystemPartnerConfiguration) SetPartnerPackage(v int32)`

SetPartnerPackage sets PartnerPackage field to given value.

### HasPartnerPackage

`func (o *KalturaSystemPartnerConfiguration) HasPartnerPackage() bool`

HasPartnerPackage returns a boolean if a field has been set.

### GetPartnerPackageClassOfService

`func (o *KalturaSystemPartnerConfiguration) GetPartnerPackageClassOfService() string`

GetPartnerPackageClassOfService returns the PartnerPackageClassOfService field if non-nil, zero value otherwise.

### GetPartnerPackageClassOfServiceOk

`func (o *KalturaSystemPartnerConfiguration) GetPartnerPackageClassOfServiceOk() (*string, bool)`

GetPartnerPackageClassOfServiceOk returns a tuple with the PartnerPackageClassOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPackageClassOfService

`func (o *KalturaSystemPartnerConfiguration) SetPartnerPackageClassOfService(v string)`

SetPartnerPackageClassOfService sets PartnerPackageClassOfService field to given value.

### HasPartnerPackageClassOfService

`func (o *KalturaSystemPartnerConfiguration) HasPartnerPackageClassOfService() bool`

HasPartnerPackageClassOfService returns a boolean if a field has been set.

### GetPartnerParentId

`func (o *KalturaSystemPartnerConfiguration) GetPartnerParentId() int32`

GetPartnerParentId returns the PartnerParentId field if non-nil, zero value otherwise.

### GetPartnerParentIdOk

`func (o *KalturaSystemPartnerConfiguration) GetPartnerParentIdOk() (*int32, bool)`

GetPartnerParentIdOk returns a tuple with the PartnerParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerParentId

`func (o *KalturaSystemPartnerConfiguration) SetPartnerParentId(v int32)`

SetPartnerParentId sets PartnerParentId field to given value.

### HasPartnerParentId

`func (o *KalturaSystemPartnerConfiguration) HasPartnerParentId() bool`

HasPartnerParentId returns a boolean if a field has been set.

### GetPassReplaceFreq

`func (o *KalturaSystemPartnerConfiguration) GetPassReplaceFreq() int32`

GetPassReplaceFreq returns the PassReplaceFreq field if non-nil, zero value otherwise.

### GetPassReplaceFreqOk

`func (o *KalturaSystemPartnerConfiguration) GetPassReplaceFreqOk() (*int32, bool)`

GetPassReplaceFreqOk returns a tuple with the PassReplaceFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassReplaceFreq

`func (o *KalturaSystemPartnerConfiguration) SetPassReplaceFreq(v int32)`

SetPassReplaceFreq sets PassReplaceFreq field to given value.

### HasPassReplaceFreq

`func (o *KalturaSystemPartnerConfiguration) HasPassReplaceFreq() bool`

HasPassReplaceFreq returns a boolean if a field has been set.

### GetPasswordStructureValidations

`func (o *KalturaSystemPartnerConfiguration) GetPasswordStructureValidations() string`

GetPasswordStructureValidations returns the PasswordStructureValidations field if non-nil, zero value otherwise.

### GetPasswordStructureValidationsOk

`func (o *KalturaSystemPartnerConfiguration) GetPasswordStructureValidationsOk() (*string, bool)`

GetPasswordStructureValidationsOk returns a tuple with the PasswordStructureValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStructureValidations

`func (o *KalturaSystemPartnerConfiguration) SetPasswordStructureValidations(v string)`

SetPasswordStructureValidations sets PasswordStructureValidations field to given value.

### HasPasswordStructureValidations

`func (o *KalturaSystemPartnerConfiguration) HasPasswordStructureValidations() bool`

HasPasswordStructureValidations returns a boolean if a field has been set.

### GetPasswordStructureValidationsDescription

`func (o *KalturaSystemPartnerConfiguration) GetPasswordStructureValidationsDescription() string`

GetPasswordStructureValidationsDescription returns the PasswordStructureValidationsDescription field if non-nil, zero value otherwise.

### GetPasswordStructureValidationsDescriptionOk

`func (o *KalturaSystemPartnerConfiguration) GetPasswordStructureValidationsDescriptionOk() (*string, bool)`

GetPasswordStructureValidationsDescriptionOk returns a tuple with the PasswordStructureValidationsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStructureValidationsDescription

`func (o *KalturaSystemPartnerConfiguration) SetPasswordStructureValidationsDescription(v string)`

SetPasswordStructureValidationsDescription sets PasswordStructureValidationsDescription field to given value.

### HasPasswordStructureValidationsDescription

`func (o *KalturaSystemPartnerConfiguration) HasPasswordStructureValidationsDescription() bool`

HasPasswordStructureValidationsDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *KalturaSystemPartnerConfiguration) GetPermissions() []KalturaPermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KalturaSystemPartnerConfiguration) GetPermissionsOk() (*[]KalturaPermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KalturaSystemPartnerConfiguration) SetPermissions(v []KalturaPermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KalturaSystemPartnerConfiguration) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPublisherEnvironmentType

`func (o *KalturaSystemPartnerConfiguration) GetPublisherEnvironmentType() int32`

GetPublisherEnvironmentType returns the PublisherEnvironmentType field if non-nil, zero value otherwise.

### GetPublisherEnvironmentTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetPublisherEnvironmentTypeOk() (*int32, bool)`

GetPublisherEnvironmentTypeOk returns a tuple with the PublisherEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherEnvironmentType

`func (o *KalturaSystemPartnerConfiguration) SetPublisherEnvironmentType(v int32)`

SetPublisherEnvironmentType sets PublisherEnvironmentType field to given value.

### HasPublisherEnvironmentType

`func (o *KalturaSystemPartnerConfiguration) HasPublisherEnvironmentType() bool`

HasPublisherEnvironmentType returns a boolean if a field has been set.

### GetPurifyImageContent

`func (o *KalturaSystemPartnerConfiguration) GetPurifyImageContent() bool`

GetPurifyImageContent returns the PurifyImageContent field if non-nil, zero value otherwise.

### GetPurifyImageContentOk

`func (o *KalturaSystemPartnerConfiguration) GetPurifyImageContentOk() (*bool, bool)`

GetPurifyImageContentOk returns a tuple with the PurifyImageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurifyImageContent

`func (o *KalturaSystemPartnerConfiguration) SetPurifyImageContent(v bool)`

SetPurifyImageContent sets PurifyImageContent field to given value.

### HasPurifyImageContent

`func (o *KalturaSystemPartnerConfiguration) HasPurifyImageContent() bool`

HasPurifyImageContent returns a boolean if a field has been set.

### GetReferenceId

`func (o *KalturaSystemPartnerConfiguration) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KalturaSystemPartnerConfiguration) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KalturaSystemPartnerConfiguration) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KalturaSystemPartnerConfiguration) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetRestrictEntryByMetadata

`func (o *KalturaSystemPartnerConfiguration) GetRestrictEntryByMetadata() bool`

GetRestrictEntryByMetadata returns the RestrictEntryByMetadata field if non-nil, zero value otherwise.

### GetRestrictEntryByMetadataOk

`func (o *KalturaSystemPartnerConfiguration) GetRestrictEntryByMetadataOk() (*bool, bool)`

GetRestrictEntryByMetadataOk returns a tuple with the RestrictEntryByMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictEntryByMetadata

`func (o *KalturaSystemPartnerConfiguration) SetRestrictEntryByMetadata(v bool)`

SetRestrictEntryByMetadata sets RestrictEntryByMetadata field to given value.

### HasRestrictEntryByMetadata

`func (o *KalturaSystemPartnerConfiguration) HasRestrictEntryByMetadata() bool`

HasRestrictEntryByMetadata returns a boolean if a field has been set.

### GetRestrictThumbnailByKs

`func (o *KalturaSystemPartnerConfiguration) GetRestrictThumbnailByKs() int32`

GetRestrictThumbnailByKs returns the RestrictThumbnailByKs field if non-nil, zero value otherwise.

### GetRestrictThumbnailByKsOk

`func (o *KalturaSystemPartnerConfiguration) GetRestrictThumbnailByKsOk() (*int32, bool)`

GetRestrictThumbnailByKsOk returns a tuple with the RestrictThumbnailByKs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictThumbnailByKs

`func (o *KalturaSystemPartnerConfiguration) SetRestrictThumbnailByKs(v int32)`

SetRestrictThumbnailByKs sets RestrictThumbnailByKs field to given value.

### HasRestrictThumbnailByKs

`func (o *KalturaSystemPartnerConfiguration) HasRestrictThumbnailByKs() bool`

HasRestrictThumbnailByKs returns a boolean if a field has been set.

### GetSecondarySecretRoleId

`func (o *KalturaSystemPartnerConfiguration) GetSecondarySecretRoleId() int32`

GetSecondarySecretRoleId returns the SecondarySecretRoleId field if non-nil, zero value otherwise.

### GetSecondarySecretRoleIdOk

`func (o *KalturaSystemPartnerConfiguration) GetSecondarySecretRoleIdOk() (*int32, bool)`

GetSecondarySecretRoleIdOk returns a tuple with the SecondarySecretRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySecretRoleId

`func (o *KalturaSystemPartnerConfiguration) SetSecondarySecretRoleId(v int32)`

SetSecondarySecretRoleId sets SecondarySecretRoleId field to given value.

### HasSecondarySecretRoleId

`func (o *KalturaSystemPartnerConfiguration) HasSecondarySecretRoleId() bool`

HasSecondarySecretRoleId returns a boolean if a field has been set.

### GetStorageDeleteFromKaltura

`func (o *KalturaSystemPartnerConfiguration) GetStorageDeleteFromKaltura() bool`

GetStorageDeleteFromKaltura returns the StorageDeleteFromKaltura field if non-nil, zero value otherwise.

### GetStorageDeleteFromKalturaOk

`func (o *KalturaSystemPartnerConfiguration) GetStorageDeleteFromKalturaOk() (*bool, bool)`

GetStorageDeleteFromKalturaOk returns a tuple with the StorageDeleteFromKaltura field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDeleteFromKaltura

`func (o *KalturaSystemPartnerConfiguration) SetStorageDeleteFromKaltura(v bool)`

SetStorageDeleteFromKaltura sets StorageDeleteFromKaltura field to given value.

### HasStorageDeleteFromKaltura

`func (o *KalturaSystemPartnerConfiguration) HasStorageDeleteFromKaltura() bool`

HasStorageDeleteFromKaltura returns a boolean if a field has been set.

### GetStorageServePriority

`func (o *KalturaSystemPartnerConfiguration) GetStorageServePriority() int32`

GetStorageServePriority returns the StorageServePriority field if non-nil, zero value otherwise.

### GetStorageServePriorityOk

`func (o *KalturaSystemPartnerConfiguration) GetStorageServePriorityOk() (*int32, bool)`

GetStorageServePriorityOk returns a tuple with the StorageServePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServePriority

`func (o *KalturaSystemPartnerConfiguration) SetStorageServePriority(v int32)`

SetStorageServePriority sets StorageServePriority field to given value.

### HasStorageServePriority

`func (o *KalturaSystemPartnerConfiguration) HasStorageServePriority() bool`

HasStorageServePriority returns a boolean if a field has been set.

### GetStreamerType

`func (o *KalturaSystemPartnerConfiguration) GetStreamerType() string`

GetStreamerType returns the StreamerType field if non-nil, zero value otherwise.

### GetStreamerTypeOk

`func (o *KalturaSystemPartnerConfiguration) GetStreamerTypeOk() (*string, bool)`

GetStreamerTypeOk returns a tuple with the StreamerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerType

`func (o *KalturaSystemPartnerConfiguration) SetStreamerType(v string)`

SetStreamerType sets StreamerType field to given value.

### HasStreamerType

`func (o *KalturaSystemPartnerConfiguration) HasStreamerType() bool`

HasStreamerType returns a boolean if a field has been set.

### GetSupportAnimatedThumbnails

`func (o *KalturaSystemPartnerConfiguration) GetSupportAnimatedThumbnails() bool`

GetSupportAnimatedThumbnails returns the SupportAnimatedThumbnails field if non-nil, zero value otherwise.

### GetSupportAnimatedThumbnailsOk

`func (o *KalturaSystemPartnerConfiguration) GetSupportAnimatedThumbnailsOk() (*bool, bool)`

GetSupportAnimatedThumbnailsOk returns a tuple with the SupportAnimatedThumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnimatedThumbnails

`func (o *KalturaSystemPartnerConfiguration) SetSupportAnimatedThumbnails(v bool)`

SetSupportAnimatedThumbnails sets SupportAnimatedThumbnails field to given value.

### HasSupportAnimatedThumbnails

`func (o *KalturaSystemPartnerConfiguration) HasSupportAnimatedThumbnails() bool`

HasSupportAnimatedThumbnails returns a boolean if a field has been set.

### GetThumbnailHost

`func (o *KalturaSystemPartnerConfiguration) GetThumbnailHost() string`

GetThumbnailHost returns the ThumbnailHost field if non-nil, zero value otherwise.

### GetThumbnailHostOk

`func (o *KalturaSystemPartnerConfiguration) GetThumbnailHostOk() (*string, bool)`

GetThumbnailHostOk returns a tuple with the ThumbnailHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailHost

`func (o *KalturaSystemPartnerConfiguration) SetThumbnailHost(v string)`

SetThumbnailHost sets ThumbnailHost field to given value.

### HasThumbnailHost

`func (o *KalturaSystemPartnerConfiguration) HasThumbnailHost() bool`

HasThumbnailHost returns a boolean if a field has been set.

### GetTimeAlignedRenditions

`func (o *KalturaSystemPartnerConfiguration) GetTimeAlignedRenditions() bool`

GetTimeAlignedRenditions returns the TimeAlignedRenditions field if non-nil, zero value otherwise.

### GetTimeAlignedRenditionsOk

`func (o *KalturaSystemPartnerConfiguration) GetTimeAlignedRenditionsOk() (*bool, bool)`

GetTimeAlignedRenditionsOk returns a tuple with the TimeAlignedRenditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAlignedRenditions

`func (o *KalturaSystemPartnerConfiguration) SetTimeAlignedRenditions(v bool)`

SetTimeAlignedRenditions sets TimeAlignedRenditions field to given value.

### HasTimeAlignedRenditions

`func (o *KalturaSystemPartnerConfiguration) HasTimeAlignedRenditions() bool`

HasTimeAlignedRenditions returns a boolean if a field has been set.

### GetTrigramPercentage

`func (o *KalturaSystemPartnerConfiguration) GetTrigramPercentage() int32`

GetTrigramPercentage returns the TrigramPercentage field if non-nil, zero value otherwise.

### GetTrigramPercentageOk

`func (o *KalturaSystemPartnerConfiguration) GetTrigramPercentageOk() (*int32, bool)`

GetTrigramPercentageOk returns a tuple with the TrigramPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigramPercentage

`func (o *KalturaSystemPartnerConfiguration) SetTrigramPercentage(v int32)`

SetTrigramPercentage sets TrigramPercentage field to given value.

### HasTrigramPercentage

`func (o *KalturaSystemPartnerConfiguration) HasTrigramPercentage() bool`

HasTrigramPercentage returns a boolean if a field has been set.

### GetTwoFactorAuthenticationMode

`func (o *KalturaSystemPartnerConfiguration) GetTwoFactorAuthenticationMode() int32`

GetTwoFactorAuthenticationMode returns the TwoFactorAuthenticationMode field if non-nil, zero value otherwise.

### GetTwoFactorAuthenticationModeOk

`func (o *KalturaSystemPartnerConfiguration) GetTwoFactorAuthenticationModeOk() (*int32, bool)`

GetTwoFactorAuthenticationModeOk returns a tuple with the TwoFactorAuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthenticationMode

`func (o *KalturaSystemPartnerConfiguration) SetTwoFactorAuthenticationMode(v int32)`

SetTwoFactorAuthenticationMode sets TwoFactorAuthenticationMode field to given value.

### HasTwoFactorAuthenticationMode

`func (o *KalturaSystemPartnerConfiguration) HasTwoFactorAuthenticationMode() bool`

HasTwoFactorAuthenticationMode returns a boolean if a field has been set.

### GetUsageLimitWarning

`func (o *KalturaSystemPartnerConfiguration) GetUsageLimitWarning() int32`

GetUsageLimitWarning returns the UsageLimitWarning field if non-nil, zero value otherwise.

### GetUsageLimitWarningOk

`func (o *KalturaSystemPartnerConfiguration) GetUsageLimitWarningOk() (*int32, bool)`

GetUsageLimitWarningOk returns a tuple with the UsageLimitWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimitWarning

`func (o *KalturaSystemPartnerConfiguration) SetUsageLimitWarning(v int32)`

SetUsageLimitWarning sets UsageLimitWarning field to given value.

### HasUsageLimitWarning

`func (o *KalturaSystemPartnerConfiguration) HasUsageLimitWarning() bool`

HasUsageLimitWarning returns a boolean if a field has been set.

### GetUsagePercent

`func (o *KalturaSystemPartnerConfiguration) GetUsagePercent() int32`

GetUsagePercent returns the UsagePercent field if non-nil, zero value otherwise.

### GetUsagePercentOk

`func (o *KalturaSystemPartnerConfiguration) GetUsagePercentOk() (*int32, bool)`

GetUsagePercentOk returns a tuple with the UsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercent

`func (o *KalturaSystemPartnerConfiguration) SetUsagePercent(v int32)`

SetUsagePercent sets UsagePercent field to given value.

### HasUsagePercent

`func (o *KalturaSystemPartnerConfiguration) HasUsagePercent() bool`

HasUsagePercent returns a boolean if a field has been set.

### GetUseSso

`func (o *KalturaSystemPartnerConfiguration) GetUseSso() bool`

GetUseSso returns the UseSso field if non-nil, zero value otherwise.

### GetUseSsoOk

`func (o *KalturaSystemPartnerConfiguration) GetUseSsoOk() (*bool, bool)`

GetUseSsoOk returns a tuple with the UseSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSso

`func (o *KalturaSystemPartnerConfiguration) SetUseSso(v bool)`

SetUseSso sets UseSso field to given value.

### HasUseSso

`func (o *KalturaSystemPartnerConfiguration) HasUseSso() bool`

HasUseSso returns a boolean if a field has been set.

### GetUseTwoFactorAuthentication

`func (o *KalturaSystemPartnerConfiguration) GetUseTwoFactorAuthentication() bool`

GetUseTwoFactorAuthentication returns the UseTwoFactorAuthentication field if non-nil, zero value otherwise.

### GetUseTwoFactorAuthenticationOk

`func (o *KalturaSystemPartnerConfiguration) GetUseTwoFactorAuthenticationOk() (*bool, bool)`

GetUseTwoFactorAuthenticationOk returns a tuple with the UseTwoFactorAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTwoFactorAuthentication

`func (o *KalturaSystemPartnerConfiguration) SetUseTwoFactorAuthentication(v bool)`

SetUseTwoFactorAuthentication sets UseTwoFactorAuthentication field to given value.

### HasUseTwoFactorAuthentication

`func (o *KalturaSystemPartnerConfiguration) HasUseTwoFactorAuthentication() bool`

HasUseTwoFactorAuthentication returns a boolean if a field has been set.

### GetUserSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) GetUserSessionRoleId() int32`

GetUserSessionRoleId returns the UserSessionRoleId field if non-nil, zero value otherwise.

### GetUserSessionRoleIdOk

`func (o *KalturaSystemPartnerConfiguration) GetUserSessionRoleIdOk() (*int32, bool)`

GetUserSessionRoleIdOk returns a tuple with the UserSessionRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) SetUserSessionRoleId(v int32)`

SetUserSessionRoleId sets UserSessionRoleId field to given value.

### HasUserSessionRoleId

`func (o *KalturaSystemPartnerConfiguration) HasUserSessionRoleId() bool`

HasUserSessionRoleId returns a boolean if a field has been set.

### GetVerticalClasiffication

`func (o *KalturaSystemPartnerConfiguration) GetVerticalClasiffication() string`

GetVerticalClasiffication returns the VerticalClasiffication field if non-nil, zero value otherwise.

### GetVerticalClasifficationOk

`func (o *KalturaSystemPartnerConfiguration) GetVerticalClasifficationOk() (*string, bool)`

GetVerticalClasifficationOk returns a tuple with the VerticalClasiffication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerticalClasiffication

`func (o *KalturaSystemPartnerConfiguration) SetVerticalClasiffication(v string)`

SetVerticalClasiffication sets VerticalClasiffication field to given value.

### HasVerticalClasiffication

`func (o *KalturaSystemPartnerConfiguration) HasVerticalClasiffication() bool`

HasVerticalClasiffication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


