# KalturaPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalParams** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**AdminEmail** | Pointer to **string** |  | [optional] 
**AdminLoginUsersQuota** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**AdminName** | Pointer to **string** | deprecated - lastName and firstName replaces this field | [optional] 
**AdminSecret** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**AdminUserId** | Pointer to **string** |  | [optional] 
**AdultContent** | Pointer to **bool** |  | [optional] 
**AllowMultiNotification** | Pointer to **int32** |  | [optional] 
**AllowQuickEdit** | Pointer to **int32** |  | [optional] 
**AllowedDomains** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**AllowedFromEmailWhiteList** | Pointer to **string** |  | [optional] 
**AppearInSearch** | Pointer to **int32** |  | [optional] 
**AuthenticationType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaPartnerAuthenticationType&#x60; | [optional] [readonly] 
**BlockDirectLogin** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CdnHost** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CmsPassword** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CommercialUse** | Pointer to **int32** | Enum Type: &#x60;KalturaCommercialUseType&#x60; | [optional] 
**ContentCategories** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | country code (2char) - this field is optional | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CrmId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefConversionProfileType** | Pointer to **string** |  | [optional] 
**DefaultDeliveryType** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefaultEmbedCodeType** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefaultEntitlementEnforcement** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DeliveryTypes** | Pointer to [**[]KalturaPlayerDeliveryType**](KalturaPlayerDeliveryType.md) |  | [optional] 
**DescribeYourself** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ESearchLanguages** | Pointer to [**[]KalturaESearchLanguageItem**](KalturaESearchLanguageItem.md) |  | [optional] 
**EightyPercentWarning** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EmbedCodeTypes** | Pointer to [**[]KalturaPlayerEmbedCodeType**](KalturaPlayerEmbedCodeType.md) |  | [optional] 
**ExcludedAdminRoleName** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ExtendedFreeTrail** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ExtendedFreeTrailEndsWarning** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ExtendedFreeTrailExpiryDate** | Pointer to **int32** | &#x60;readOnly&#x60;  Unix timestamp (In seconds) | [optional] [readonly] 
**ExtendedFreeTrailExpiryReason** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FirstName** | Pointer to **string** | firstName and lastName replace the old (deprecated) adminName | [optional] 
**Host** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IgnoreSeoLinks** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IsFirstLogin** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IsSelfServe** | Pointer to **bool** |  | [optional] 
**LandingPage** | Pointer to **string** |  | [optional] 
**LastFreeTrialNotificationDay** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LastName** | Pointer to **string** | lastName and firstName replace the old (deprecated) adminName | [optional] 
**LoginBlockPeriod** | Pointer to **int32** |  | [optional] 
**LogoutUrl** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MaxLoginAttempts** | Pointer to **int32** |  | [optional] 
**MaxUploadSize** | Pointer to **int32** |  | [optional] 
**MergeEntryLists** | Pointer to **int32** |  | [optional] 
**MonitorUsage** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**NotificationUrl** | Pointer to **string** |  | [optional] 
**NotificationsConfig** | Pointer to **string** |  | [optional] 
**Notify** | Pointer to **int32** |  | [optional] 
**NumPrevPassToKeep** | Pointer to **int32** |  | [optional] 
**OttEnvironmentUrl** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**OvpEnvironmentUrl** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerGroupType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaPartnerGroupType&#x60; | [optional] [readonly] 
**PartnerPackage** | Pointer to **int32** |  | [optional] 
**PartnerParentId** | Pointer to **int32** |  | [optional] 
**PassReplaceFreq** | Pointer to **int32** |  | [optional] 
**PasswordStructureValidations** | Pointer to [**[]KalturaRegexItem**](KalturaRegexItem.md) |  | [optional] 
**PasswordStructureValidationsDescription** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**PublisherEnvironmentType** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PublishersQuota** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**State** | Pointer to **string** | state code (2char) - this field is optional | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaPartnerStatus&#x60; | [optional] [readonly] 
**TemplatePartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**TimeAlignedRenditions** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**TwoFactorAuthenticationMode** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaTwoFactorAuthenticationMode&#x60; | [optional] [readonly] 
**Type** | Pointer to **int32** | Enum Type: &#x60;KalturaPartnerType&#x60; | [optional] 
**UsageLimitWarning** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserLandingPage** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaPartner

`func NewKalturaPartner() *KalturaPartner`

NewKalturaPartner instantiates a new KalturaPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaPartnerWithDefaults

`func NewKalturaPartnerWithDefaults() *KalturaPartner`

NewKalturaPartnerWithDefaults instantiates a new KalturaPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalParams

`func (o *KalturaPartner) GetAdditionalParams() []KalturaKeyValue`

GetAdditionalParams returns the AdditionalParams field if non-nil, zero value otherwise.

### GetAdditionalParamsOk

`func (o *KalturaPartner) GetAdditionalParamsOk() (*[]KalturaKeyValue, bool)`

GetAdditionalParamsOk returns a tuple with the AdditionalParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParams

`func (o *KalturaPartner) SetAdditionalParams(v []KalturaKeyValue)`

SetAdditionalParams sets AdditionalParams field to given value.

### HasAdditionalParams

`func (o *KalturaPartner) HasAdditionalParams() bool`

HasAdditionalParams returns a boolean if a field has been set.

### GetAdminEmail

`func (o *KalturaPartner) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *KalturaPartner) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *KalturaPartner) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *KalturaPartner) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminLoginUsersQuota

`func (o *KalturaPartner) GetAdminLoginUsersQuota() int32`

GetAdminLoginUsersQuota returns the AdminLoginUsersQuota field if non-nil, zero value otherwise.

### GetAdminLoginUsersQuotaOk

`func (o *KalturaPartner) GetAdminLoginUsersQuotaOk() (*int32, bool)`

GetAdminLoginUsersQuotaOk returns a tuple with the AdminLoginUsersQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminLoginUsersQuota

`func (o *KalturaPartner) SetAdminLoginUsersQuota(v int32)`

SetAdminLoginUsersQuota sets AdminLoginUsersQuota field to given value.

### HasAdminLoginUsersQuota

`func (o *KalturaPartner) HasAdminLoginUsersQuota() bool`

HasAdminLoginUsersQuota returns a boolean if a field has been set.

### GetAdminName

`func (o *KalturaPartner) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *KalturaPartner) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *KalturaPartner) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *KalturaPartner) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAdminSecret

`func (o *KalturaPartner) GetAdminSecret() string`

GetAdminSecret returns the AdminSecret field if non-nil, zero value otherwise.

### GetAdminSecretOk

`func (o *KalturaPartner) GetAdminSecretOk() (*string, bool)`

GetAdminSecretOk returns a tuple with the AdminSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSecret

`func (o *KalturaPartner) SetAdminSecret(v string)`

SetAdminSecret sets AdminSecret field to given value.

### HasAdminSecret

`func (o *KalturaPartner) HasAdminSecret() bool`

HasAdminSecret returns a boolean if a field has been set.

### GetAdminUserId

`func (o *KalturaPartner) GetAdminUserId() string`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *KalturaPartner) GetAdminUserIdOk() (*string, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *KalturaPartner) SetAdminUserId(v string)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *KalturaPartner) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetAdultContent

`func (o *KalturaPartner) GetAdultContent() bool`

GetAdultContent returns the AdultContent field if non-nil, zero value otherwise.

### GetAdultContentOk

`func (o *KalturaPartner) GetAdultContentOk() (*bool, bool)`

GetAdultContentOk returns a tuple with the AdultContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultContent

`func (o *KalturaPartner) SetAdultContent(v bool)`

SetAdultContent sets AdultContent field to given value.

### HasAdultContent

`func (o *KalturaPartner) HasAdultContent() bool`

HasAdultContent returns a boolean if a field has been set.

### GetAllowMultiNotification

`func (o *KalturaPartner) GetAllowMultiNotification() int32`

GetAllowMultiNotification returns the AllowMultiNotification field if non-nil, zero value otherwise.

### GetAllowMultiNotificationOk

`func (o *KalturaPartner) GetAllowMultiNotificationOk() (*int32, bool)`

GetAllowMultiNotificationOk returns a tuple with the AllowMultiNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiNotification

`func (o *KalturaPartner) SetAllowMultiNotification(v int32)`

SetAllowMultiNotification sets AllowMultiNotification field to given value.

### HasAllowMultiNotification

`func (o *KalturaPartner) HasAllowMultiNotification() bool`

HasAllowMultiNotification returns a boolean if a field has been set.

### GetAllowQuickEdit

`func (o *KalturaPartner) GetAllowQuickEdit() int32`

GetAllowQuickEdit returns the AllowQuickEdit field if non-nil, zero value otherwise.

### GetAllowQuickEditOk

`func (o *KalturaPartner) GetAllowQuickEditOk() (*int32, bool)`

GetAllowQuickEditOk returns a tuple with the AllowQuickEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuickEdit

`func (o *KalturaPartner) SetAllowQuickEdit(v int32)`

SetAllowQuickEdit sets AllowQuickEdit field to given value.

### HasAllowQuickEdit

`func (o *KalturaPartner) HasAllowQuickEdit() bool`

HasAllowQuickEdit returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *KalturaPartner) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *KalturaPartner) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *KalturaPartner) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *KalturaPartner) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedFromEmailWhiteList

`func (o *KalturaPartner) GetAllowedFromEmailWhiteList() string`

GetAllowedFromEmailWhiteList returns the AllowedFromEmailWhiteList field if non-nil, zero value otherwise.

### GetAllowedFromEmailWhiteListOk

`func (o *KalturaPartner) GetAllowedFromEmailWhiteListOk() (*string, bool)`

GetAllowedFromEmailWhiteListOk returns a tuple with the AllowedFromEmailWhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedFromEmailWhiteList

`func (o *KalturaPartner) SetAllowedFromEmailWhiteList(v string)`

SetAllowedFromEmailWhiteList sets AllowedFromEmailWhiteList field to given value.

### HasAllowedFromEmailWhiteList

`func (o *KalturaPartner) HasAllowedFromEmailWhiteList() bool`

HasAllowedFromEmailWhiteList returns a boolean if a field has been set.

### GetAppearInSearch

`func (o *KalturaPartner) GetAppearInSearch() int32`

GetAppearInSearch returns the AppearInSearch field if non-nil, zero value otherwise.

### GetAppearInSearchOk

`func (o *KalturaPartner) GetAppearInSearchOk() (*int32, bool)`

GetAppearInSearchOk returns a tuple with the AppearInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearInSearch

`func (o *KalturaPartner) SetAppearInSearch(v int32)`

SetAppearInSearch sets AppearInSearch field to given value.

### HasAppearInSearch

`func (o *KalturaPartner) HasAppearInSearch() bool`

HasAppearInSearch returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *KalturaPartner) GetAuthenticationType() int32`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *KalturaPartner) GetAuthenticationTypeOk() (*int32, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *KalturaPartner) SetAuthenticationType(v int32)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *KalturaPartner) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetBlockDirectLogin

`func (o *KalturaPartner) GetBlockDirectLogin() bool`

GetBlockDirectLogin returns the BlockDirectLogin field if non-nil, zero value otherwise.

### GetBlockDirectLoginOk

`func (o *KalturaPartner) GetBlockDirectLoginOk() (*bool, bool)`

GetBlockDirectLoginOk returns a tuple with the BlockDirectLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDirectLogin

`func (o *KalturaPartner) SetBlockDirectLogin(v bool)`

SetBlockDirectLogin sets BlockDirectLogin field to given value.

### HasBlockDirectLogin

`func (o *KalturaPartner) HasBlockDirectLogin() bool`

HasBlockDirectLogin returns a boolean if a field has been set.

### GetCdnHost

`func (o *KalturaPartner) GetCdnHost() string`

GetCdnHost returns the CdnHost field if non-nil, zero value otherwise.

### GetCdnHostOk

`func (o *KalturaPartner) GetCdnHostOk() (*string, bool)`

GetCdnHostOk returns a tuple with the CdnHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnHost

`func (o *KalturaPartner) SetCdnHost(v string)`

SetCdnHost sets CdnHost field to given value.

### HasCdnHost

`func (o *KalturaPartner) HasCdnHost() bool`

HasCdnHost returns a boolean if a field has been set.

### GetCmsPassword

`func (o *KalturaPartner) GetCmsPassword() string`

GetCmsPassword returns the CmsPassword field if non-nil, zero value otherwise.

### GetCmsPasswordOk

`func (o *KalturaPartner) GetCmsPasswordOk() (*string, bool)`

GetCmsPasswordOk returns a tuple with the CmsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmsPassword

`func (o *KalturaPartner) SetCmsPassword(v string)`

SetCmsPassword sets CmsPassword field to given value.

### HasCmsPassword

`func (o *KalturaPartner) HasCmsPassword() bool`

HasCmsPassword returns a boolean if a field has been set.

### GetCommercialUse

`func (o *KalturaPartner) GetCommercialUse() int32`

GetCommercialUse returns the CommercialUse field if non-nil, zero value otherwise.

### GetCommercialUseOk

`func (o *KalturaPartner) GetCommercialUseOk() (*int32, bool)`

GetCommercialUseOk returns a tuple with the CommercialUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialUse

`func (o *KalturaPartner) SetCommercialUse(v int32)`

SetCommercialUse sets CommercialUse field to given value.

### HasCommercialUse

`func (o *KalturaPartner) HasCommercialUse() bool`

HasCommercialUse returns a boolean if a field has been set.

### GetContentCategories

`func (o *KalturaPartner) GetContentCategories() string`

GetContentCategories returns the ContentCategories field if non-nil, zero value otherwise.

### GetContentCategoriesOk

`func (o *KalturaPartner) GetContentCategoriesOk() (*string, bool)`

GetContentCategoriesOk returns a tuple with the ContentCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCategories

`func (o *KalturaPartner) SetContentCategories(v string)`

SetContentCategories sets ContentCategories field to given value.

### HasContentCategories

`func (o *KalturaPartner) HasContentCategories() bool`

HasContentCategories returns a boolean if a field has been set.

### GetCountry

`func (o *KalturaPartner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KalturaPartner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KalturaPartner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *KalturaPartner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaPartner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaPartner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaPartner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaPartner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCrmId

`func (o *KalturaPartner) GetCrmId() string`

GetCrmId returns the CrmId field if non-nil, zero value otherwise.

### GetCrmIdOk

`func (o *KalturaPartner) GetCrmIdOk() (*string, bool)`

GetCrmIdOk returns a tuple with the CrmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmId

`func (o *KalturaPartner) SetCrmId(v string)`

SetCrmId sets CrmId field to given value.

### HasCrmId

`func (o *KalturaPartner) HasCrmId() bool`

HasCrmId returns a boolean if a field has been set.

### GetDefConversionProfileType

`func (o *KalturaPartner) GetDefConversionProfileType() string`

GetDefConversionProfileType returns the DefConversionProfileType field if non-nil, zero value otherwise.

### GetDefConversionProfileTypeOk

`func (o *KalturaPartner) GetDefConversionProfileTypeOk() (*string, bool)`

GetDefConversionProfileTypeOk returns a tuple with the DefConversionProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefConversionProfileType

`func (o *KalturaPartner) SetDefConversionProfileType(v string)`

SetDefConversionProfileType sets DefConversionProfileType field to given value.

### HasDefConversionProfileType

`func (o *KalturaPartner) HasDefConversionProfileType() bool`

HasDefConversionProfileType returns a boolean if a field has been set.

### GetDefaultDeliveryType

`func (o *KalturaPartner) GetDefaultDeliveryType() string`

GetDefaultDeliveryType returns the DefaultDeliveryType field if non-nil, zero value otherwise.

### GetDefaultDeliveryTypeOk

`func (o *KalturaPartner) GetDefaultDeliveryTypeOk() (*string, bool)`

GetDefaultDeliveryTypeOk returns a tuple with the DefaultDeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeliveryType

`func (o *KalturaPartner) SetDefaultDeliveryType(v string)`

SetDefaultDeliveryType sets DefaultDeliveryType field to given value.

### HasDefaultDeliveryType

`func (o *KalturaPartner) HasDefaultDeliveryType() bool`

HasDefaultDeliveryType returns a boolean if a field has been set.

### GetDefaultEmbedCodeType

`func (o *KalturaPartner) GetDefaultEmbedCodeType() string`

GetDefaultEmbedCodeType returns the DefaultEmbedCodeType field if non-nil, zero value otherwise.

### GetDefaultEmbedCodeTypeOk

`func (o *KalturaPartner) GetDefaultEmbedCodeTypeOk() (*string, bool)`

GetDefaultEmbedCodeTypeOk returns a tuple with the DefaultEmbedCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmbedCodeType

`func (o *KalturaPartner) SetDefaultEmbedCodeType(v string)`

SetDefaultEmbedCodeType sets DefaultEmbedCodeType field to given value.

### HasDefaultEmbedCodeType

`func (o *KalturaPartner) HasDefaultEmbedCodeType() bool`

HasDefaultEmbedCodeType returns a boolean if a field has been set.

### GetDefaultEntitlementEnforcement

`func (o *KalturaPartner) GetDefaultEntitlementEnforcement() bool`

GetDefaultEntitlementEnforcement returns the DefaultEntitlementEnforcement field if non-nil, zero value otherwise.

### GetDefaultEntitlementEnforcementOk

`func (o *KalturaPartner) GetDefaultEntitlementEnforcementOk() (*bool, bool)`

GetDefaultEntitlementEnforcementOk returns a tuple with the DefaultEntitlementEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntitlementEnforcement

`func (o *KalturaPartner) SetDefaultEntitlementEnforcement(v bool)`

SetDefaultEntitlementEnforcement sets DefaultEntitlementEnforcement field to given value.

### HasDefaultEntitlementEnforcement

`func (o *KalturaPartner) HasDefaultEntitlementEnforcement() bool`

HasDefaultEntitlementEnforcement returns a boolean if a field has been set.

### GetDeliveryTypes

`func (o *KalturaPartner) GetDeliveryTypes() []KalturaPlayerDeliveryType`

GetDeliveryTypes returns the DeliveryTypes field if non-nil, zero value otherwise.

### GetDeliveryTypesOk

`func (o *KalturaPartner) GetDeliveryTypesOk() (*[]KalturaPlayerDeliveryType, bool)`

GetDeliveryTypesOk returns a tuple with the DeliveryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTypes

`func (o *KalturaPartner) SetDeliveryTypes(v []KalturaPlayerDeliveryType)`

SetDeliveryTypes sets DeliveryTypes field to given value.

### HasDeliveryTypes

`func (o *KalturaPartner) HasDeliveryTypes() bool`

HasDeliveryTypes returns a boolean if a field has been set.

### GetDescribeYourself

`func (o *KalturaPartner) GetDescribeYourself() string`

GetDescribeYourself returns the DescribeYourself field if non-nil, zero value otherwise.

### GetDescribeYourselfOk

`func (o *KalturaPartner) GetDescribeYourselfOk() (*string, bool)`

GetDescribeYourselfOk returns a tuple with the DescribeYourself field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribeYourself

`func (o *KalturaPartner) SetDescribeYourself(v string)`

SetDescribeYourself sets DescribeYourself field to given value.

### HasDescribeYourself

`func (o *KalturaPartner) HasDescribeYourself() bool`

HasDescribeYourself returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaPartner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaPartner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaPartner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaPartner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetESearchLanguages

`func (o *KalturaPartner) GetESearchLanguages() []KalturaESearchLanguageItem`

GetESearchLanguages returns the ESearchLanguages field if non-nil, zero value otherwise.

### GetESearchLanguagesOk

`func (o *KalturaPartner) GetESearchLanguagesOk() (*[]KalturaESearchLanguageItem, bool)`

GetESearchLanguagesOk returns a tuple with the ESearchLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESearchLanguages

`func (o *KalturaPartner) SetESearchLanguages(v []KalturaESearchLanguageItem)`

SetESearchLanguages sets ESearchLanguages field to given value.

### HasESearchLanguages

`func (o *KalturaPartner) HasESearchLanguages() bool`

HasESearchLanguages returns a boolean if a field has been set.

### GetEightyPercentWarning

`func (o *KalturaPartner) GetEightyPercentWarning() int32`

GetEightyPercentWarning returns the EightyPercentWarning field if non-nil, zero value otherwise.

### GetEightyPercentWarningOk

`func (o *KalturaPartner) GetEightyPercentWarningOk() (*int32, bool)`

GetEightyPercentWarningOk returns a tuple with the EightyPercentWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEightyPercentWarning

`func (o *KalturaPartner) SetEightyPercentWarning(v int32)`

SetEightyPercentWarning sets EightyPercentWarning field to given value.

### HasEightyPercentWarning

`func (o *KalturaPartner) HasEightyPercentWarning() bool`

HasEightyPercentWarning returns a boolean if a field has been set.

### GetEmbedCodeTypes

`func (o *KalturaPartner) GetEmbedCodeTypes() []KalturaPlayerEmbedCodeType`

GetEmbedCodeTypes returns the EmbedCodeTypes field if non-nil, zero value otherwise.

### GetEmbedCodeTypesOk

`func (o *KalturaPartner) GetEmbedCodeTypesOk() (*[]KalturaPlayerEmbedCodeType, bool)`

GetEmbedCodeTypesOk returns a tuple with the EmbedCodeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedCodeTypes

`func (o *KalturaPartner) SetEmbedCodeTypes(v []KalturaPlayerEmbedCodeType)`

SetEmbedCodeTypes sets EmbedCodeTypes field to given value.

### HasEmbedCodeTypes

`func (o *KalturaPartner) HasEmbedCodeTypes() bool`

HasEmbedCodeTypes returns a boolean if a field has been set.

### GetExcludedAdminRoleName

`func (o *KalturaPartner) GetExcludedAdminRoleName() string`

GetExcludedAdminRoleName returns the ExcludedAdminRoleName field if non-nil, zero value otherwise.

### GetExcludedAdminRoleNameOk

`func (o *KalturaPartner) GetExcludedAdminRoleNameOk() (*string, bool)`

GetExcludedAdminRoleNameOk returns a tuple with the ExcludedAdminRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAdminRoleName

`func (o *KalturaPartner) SetExcludedAdminRoleName(v string)`

SetExcludedAdminRoleName sets ExcludedAdminRoleName field to given value.

### HasExcludedAdminRoleName

`func (o *KalturaPartner) HasExcludedAdminRoleName() bool`

HasExcludedAdminRoleName returns a boolean if a field has been set.

### GetExtendedFreeTrail

`func (o *KalturaPartner) GetExtendedFreeTrail() int32`

GetExtendedFreeTrail returns the ExtendedFreeTrail field if non-nil, zero value otherwise.

### GetExtendedFreeTrailOk

`func (o *KalturaPartner) GetExtendedFreeTrailOk() (*int32, bool)`

GetExtendedFreeTrailOk returns a tuple with the ExtendedFreeTrail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrail

`func (o *KalturaPartner) SetExtendedFreeTrail(v int32)`

SetExtendedFreeTrail sets ExtendedFreeTrail field to given value.

### HasExtendedFreeTrail

`func (o *KalturaPartner) HasExtendedFreeTrail() bool`

HasExtendedFreeTrail returns a boolean if a field has been set.

### GetExtendedFreeTrailEndsWarning

`func (o *KalturaPartner) GetExtendedFreeTrailEndsWarning() bool`

GetExtendedFreeTrailEndsWarning returns the ExtendedFreeTrailEndsWarning field if non-nil, zero value otherwise.

### GetExtendedFreeTrailEndsWarningOk

`func (o *KalturaPartner) GetExtendedFreeTrailEndsWarningOk() (*bool, bool)`

GetExtendedFreeTrailEndsWarningOk returns a tuple with the ExtendedFreeTrailEndsWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailEndsWarning

`func (o *KalturaPartner) SetExtendedFreeTrailEndsWarning(v bool)`

SetExtendedFreeTrailEndsWarning sets ExtendedFreeTrailEndsWarning field to given value.

### HasExtendedFreeTrailEndsWarning

`func (o *KalturaPartner) HasExtendedFreeTrailEndsWarning() bool`

HasExtendedFreeTrailEndsWarning returns a boolean if a field has been set.

### GetExtendedFreeTrailExpiryDate

`func (o *KalturaPartner) GetExtendedFreeTrailExpiryDate() int32`

GetExtendedFreeTrailExpiryDate returns the ExtendedFreeTrailExpiryDate field if non-nil, zero value otherwise.

### GetExtendedFreeTrailExpiryDateOk

`func (o *KalturaPartner) GetExtendedFreeTrailExpiryDateOk() (*int32, bool)`

GetExtendedFreeTrailExpiryDateOk returns a tuple with the ExtendedFreeTrailExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailExpiryDate

`func (o *KalturaPartner) SetExtendedFreeTrailExpiryDate(v int32)`

SetExtendedFreeTrailExpiryDate sets ExtendedFreeTrailExpiryDate field to given value.

### HasExtendedFreeTrailExpiryDate

`func (o *KalturaPartner) HasExtendedFreeTrailExpiryDate() bool`

HasExtendedFreeTrailExpiryDate returns a boolean if a field has been set.

### GetExtendedFreeTrailExpiryReason

`func (o *KalturaPartner) GetExtendedFreeTrailExpiryReason() string`

GetExtendedFreeTrailExpiryReason returns the ExtendedFreeTrailExpiryReason field if non-nil, zero value otherwise.

### GetExtendedFreeTrailExpiryReasonOk

`func (o *KalturaPartner) GetExtendedFreeTrailExpiryReasonOk() (*string, bool)`

GetExtendedFreeTrailExpiryReasonOk returns a tuple with the ExtendedFreeTrailExpiryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFreeTrailExpiryReason

`func (o *KalturaPartner) SetExtendedFreeTrailExpiryReason(v string)`

SetExtendedFreeTrailExpiryReason sets ExtendedFreeTrailExpiryReason field to given value.

### HasExtendedFreeTrailExpiryReason

`func (o *KalturaPartner) HasExtendedFreeTrailExpiryReason() bool`

HasExtendedFreeTrailExpiryReason returns a boolean if a field has been set.

### GetFirstName

`func (o *KalturaPartner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *KalturaPartner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *KalturaPartner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *KalturaPartner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHost

`func (o *KalturaPartner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KalturaPartner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KalturaPartner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KalturaPartner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *KalturaPartner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaPartner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaPartner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaPartner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreSeoLinks

`func (o *KalturaPartner) GetIgnoreSeoLinks() bool`

GetIgnoreSeoLinks returns the IgnoreSeoLinks field if non-nil, zero value otherwise.

### GetIgnoreSeoLinksOk

`func (o *KalturaPartner) GetIgnoreSeoLinksOk() (*bool, bool)`

GetIgnoreSeoLinksOk returns a tuple with the IgnoreSeoLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSeoLinks

`func (o *KalturaPartner) SetIgnoreSeoLinks(v bool)`

SetIgnoreSeoLinks sets IgnoreSeoLinks field to given value.

### HasIgnoreSeoLinks

`func (o *KalturaPartner) HasIgnoreSeoLinks() bool`

HasIgnoreSeoLinks returns a boolean if a field has been set.

### GetIsFirstLogin

`func (o *KalturaPartner) GetIsFirstLogin() bool`

GetIsFirstLogin returns the IsFirstLogin field if non-nil, zero value otherwise.

### GetIsFirstLoginOk

`func (o *KalturaPartner) GetIsFirstLoginOk() (*bool, bool)`

GetIsFirstLoginOk returns a tuple with the IsFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstLogin

`func (o *KalturaPartner) SetIsFirstLogin(v bool)`

SetIsFirstLogin sets IsFirstLogin field to given value.

### HasIsFirstLogin

`func (o *KalturaPartner) HasIsFirstLogin() bool`

HasIsFirstLogin returns a boolean if a field has been set.

### GetIsSelfServe

`func (o *KalturaPartner) GetIsSelfServe() bool`

GetIsSelfServe returns the IsSelfServe field if non-nil, zero value otherwise.

### GetIsSelfServeOk

`func (o *KalturaPartner) GetIsSelfServeOk() (*bool, bool)`

GetIsSelfServeOk returns a tuple with the IsSelfServe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfServe

`func (o *KalturaPartner) SetIsSelfServe(v bool)`

SetIsSelfServe sets IsSelfServe field to given value.

### HasIsSelfServe

`func (o *KalturaPartner) HasIsSelfServe() bool`

HasIsSelfServe returns a boolean if a field has been set.

### GetLandingPage

`func (o *KalturaPartner) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *KalturaPartner) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *KalturaPartner) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *KalturaPartner) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### GetLastFreeTrialNotificationDay

`func (o *KalturaPartner) GetLastFreeTrialNotificationDay() int32`

GetLastFreeTrialNotificationDay returns the LastFreeTrialNotificationDay field if non-nil, zero value otherwise.

### GetLastFreeTrialNotificationDayOk

`func (o *KalturaPartner) GetLastFreeTrialNotificationDayOk() (*int32, bool)`

GetLastFreeTrialNotificationDayOk returns a tuple with the LastFreeTrialNotificationDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFreeTrialNotificationDay

`func (o *KalturaPartner) SetLastFreeTrialNotificationDay(v int32)`

SetLastFreeTrialNotificationDay sets LastFreeTrialNotificationDay field to given value.

### HasLastFreeTrialNotificationDay

`func (o *KalturaPartner) HasLastFreeTrialNotificationDay() bool`

HasLastFreeTrialNotificationDay returns a boolean if a field has been set.

### GetLastName

`func (o *KalturaPartner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *KalturaPartner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *KalturaPartner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *KalturaPartner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLoginBlockPeriod

`func (o *KalturaPartner) GetLoginBlockPeriod() int32`

GetLoginBlockPeriod returns the LoginBlockPeriod field if non-nil, zero value otherwise.

### GetLoginBlockPeriodOk

`func (o *KalturaPartner) GetLoginBlockPeriodOk() (*int32, bool)`

GetLoginBlockPeriodOk returns a tuple with the LoginBlockPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBlockPeriod

`func (o *KalturaPartner) SetLoginBlockPeriod(v int32)`

SetLoginBlockPeriod sets LoginBlockPeriod field to given value.

### HasLoginBlockPeriod

`func (o *KalturaPartner) HasLoginBlockPeriod() bool`

HasLoginBlockPeriod returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *KalturaPartner) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *KalturaPartner) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *KalturaPartner) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *KalturaPartner) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetMaxLoginAttempts

`func (o *KalturaPartner) GetMaxLoginAttempts() int32`

GetMaxLoginAttempts returns the MaxLoginAttempts field if non-nil, zero value otherwise.

### GetMaxLoginAttemptsOk

`func (o *KalturaPartner) GetMaxLoginAttemptsOk() (*int32, bool)`

GetMaxLoginAttemptsOk returns a tuple with the MaxLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoginAttempts

`func (o *KalturaPartner) SetMaxLoginAttempts(v int32)`

SetMaxLoginAttempts sets MaxLoginAttempts field to given value.

### HasMaxLoginAttempts

`func (o *KalturaPartner) HasMaxLoginAttempts() bool`

HasMaxLoginAttempts returns a boolean if a field has been set.

### GetMaxUploadSize

`func (o *KalturaPartner) GetMaxUploadSize() int32`

GetMaxUploadSize returns the MaxUploadSize field if non-nil, zero value otherwise.

### GetMaxUploadSizeOk

`func (o *KalturaPartner) GetMaxUploadSizeOk() (*int32, bool)`

GetMaxUploadSizeOk returns a tuple with the MaxUploadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUploadSize

`func (o *KalturaPartner) SetMaxUploadSize(v int32)`

SetMaxUploadSize sets MaxUploadSize field to given value.

### HasMaxUploadSize

`func (o *KalturaPartner) HasMaxUploadSize() bool`

HasMaxUploadSize returns a boolean if a field has been set.

### GetMergeEntryLists

`func (o *KalturaPartner) GetMergeEntryLists() int32`

GetMergeEntryLists returns the MergeEntryLists field if non-nil, zero value otherwise.

### GetMergeEntryListsOk

`func (o *KalturaPartner) GetMergeEntryListsOk() (*int32, bool)`

GetMergeEntryListsOk returns a tuple with the MergeEntryLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeEntryLists

`func (o *KalturaPartner) SetMergeEntryLists(v int32)`

SetMergeEntryLists sets MergeEntryLists field to given value.

### HasMergeEntryLists

`func (o *KalturaPartner) HasMergeEntryLists() bool`

HasMergeEntryLists returns a boolean if a field has been set.

### GetMonitorUsage

`func (o *KalturaPartner) GetMonitorUsage() int32`

GetMonitorUsage returns the MonitorUsage field if non-nil, zero value otherwise.

### GetMonitorUsageOk

`func (o *KalturaPartner) GetMonitorUsageOk() (*int32, bool)`

GetMonitorUsageOk returns a tuple with the MonitorUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsage

`func (o *KalturaPartner) SetMonitorUsage(v int32)`

SetMonitorUsage sets MonitorUsage field to given value.

### HasMonitorUsage

`func (o *KalturaPartner) HasMonitorUsage() bool`

HasMonitorUsage returns a boolean if a field has been set.

### GetName

`func (o *KalturaPartner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaPartner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaPartner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaPartner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotificationUrl

`func (o *KalturaPartner) GetNotificationUrl() string`

GetNotificationUrl returns the NotificationUrl field if non-nil, zero value otherwise.

### GetNotificationUrlOk

`func (o *KalturaPartner) GetNotificationUrlOk() (*string, bool)`

GetNotificationUrlOk returns a tuple with the NotificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUrl

`func (o *KalturaPartner) SetNotificationUrl(v string)`

SetNotificationUrl sets NotificationUrl field to given value.

### HasNotificationUrl

`func (o *KalturaPartner) HasNotificationUrl() bool`

HasNotificationUrl returns a boolean if a field has been set.

### GetNotificationsConfig

`func (o *KalturaPartner) GetNotificationsConfig() string`

GetNotificationsConfig returns the NotificationsConfig field if non-nil, zero value otherwise.

### GetNotificationsConfigOk

`func (o *KalturaPartner) GetNotificationsConfigOk() (*string, bool)`

GetNotificationsConfigOk returns a tuple with the NotificationsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsConfig

`func (o *KalturaPartner) SetNotificationsConfig(v string)`

SetNotificationsConfig sets NotificationsConfig field to given value.

### HasNotificationsConfig

`func (o *KalturaPartner) HasNotificationsConfig() bool`

HasNotificationsConfig returns a boolean if a field has been set.

### GetNotify

`func (o *KalturaPartner) GetNotify() int32`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *KalturaPartner) GetNotifyOk() (*int32, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *KalturaPartner) SetNotify(v int32)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *KalturaPartner) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNumPrevPassToKeep

`func (o *KalturaPartner) GetNumPrevPassToKeep() int32`

GetNumPrevPassToKeep returns the NumPrevPassToKeep field if non-nil, zero value otherwise.

### GetNumPrevPassToKeepOk

`func (o *KalturaPartner) GetNumPrevPassToKeepOk() (*int32, bool)`

GetNumPrevPassToKeepOk returns a tuple with the NumPrevPassToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPrevPassToKeep

`func (o *KalturaPartner) SetNumPrevPassToKeep(v int32)`

SetNumPrevPassToKeep sets NumPrevPassToKeep field to given value.

### HasNumPrevPassToKeep

`func (o *KalturaPartner) HasNumPrevPassToKeep() bool`

HasNumPrevPassToKeep returns a boolean if a field has been set.

### GetOttEnvironmentUrl

`func (o *KalturaPartner) GetOttEnvironmentUrl() string`

GetOttEnvironmentUrl returns the OttEnvironmentUrl field if non-nil, zero value otherwise.

### GetOttEnvironmentUrlOk

`func (o *KalturaPartner) GetOttEnvironmentUrlOk() (*string, bool)`

GetOttEnvironmentUrlOk returns a tuple with the OttEnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttEnvironmentUrl

`func (o *KalturaPartner) SetOttEnvironmentUrl(v string)`

SetOttEnvironmentUrl sets OttEnvironmentUrl field to given value.

### HasOttEnvironmentUrl

`func (o *KalturaPartner) HasOttEnvironmentUrl() bool`

HasOttEnvironmentUrl returns a boolean if a field has been set.

### GetOvpEnvironmentUrl

`func (o *KalturaPartner) GetOvpEnvironmentUrl() string`

GetOvpEnvironmentUrl returns the OvpEnvironmentUrl field if non-nil, zero value otherwise.

### GetOvpEnvironmentUrlOk

`func (o *KalturaPartner) GetOvpEnvironmentUrlOk() (*string, bool)`

GetOvpEnvironmentUrlOk returns a tuple with the OvpEnvironmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvpEnvironmentUrl

`func (o *KalturaPartner) SetOvpEnvironmentUrl(v string)`

SetOvpEnvironmentUrl sets OvpEnvironmentUrl field to given value.

### HasOvpEnvironmentUrl

`func (o *KalturaPartner) HasOvpEnvironmentUrl() bool`

HasOvpEnvironmentUrl returns a boolean if a field has been set.

### GetPartnerGroupType

`func (o *KalturaPartner) GetPartnerGroupType() int32`

GetPartnerGroupType returns the PartnerGroupType field if non-nil, zero value otherwise.

### GetPartnerGroupTypeOk

`func (o *KalturaPartner) GetPartnerGroupTypeOk() (*int32, bool)`

GetPartnerGroupTypeOk returns a tuple with the PartnerGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerGroupType

`func (o *KalturaPartner) SetPartnerGroupType(v int32)`

SetPartnerGroupType sets PartnerGroupType field to given value.

### HasPartnerGroupType

`func (o *KalturaPartner) HasPartnerGroupType() bool`

HasPartnerGroupType returns a boolean if a field has been set.

### GetPartnerPackage

`func (o *KalturaPartner) GetPartnerPackage() int32`

GetPartnerPackage returns the PartnerPackage field if non-nil, zero value otherwise.

### GetPartnerPackageOk

`func (o *KalturaPartner) GetPartnerPackageOk() (*int32, bool)`

GetPartnerPackageOk returns a tuple with the PartnerPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPackage

`func (o *KalturaPartner) SetPartnerPackage(v int32)`

SetPartnerPackage sets PartnerPackage field to given value.

### HasPartnerPackage

`func (o *KalturaPartner) HasPartnerPackage() bool`

HasPartnerPackage returns a boolean if a field has been set.

### GetPartnerParentId

`func (o *KalturaPartner) GetPartnerParentId() int32`

GetPartnerParentId returns the PartnerParentId field if non-nil, zero value otherwise.

### GetPartnerParentIdOk

`func (o *KalturaPartner) GetPartnerParentIdOk() (*int32, bool)`

GetPartnerParentIdOk returns a tuple with the PartnerParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerParentId

`func (o *KalturaPartner) SetPartnerParentId(v int32)`

SetPartnerParentId sets PartnerParentId field to given value.

### HasPartnerParentId

`func (o *KalturaPartner) HasPartnerParentId() bool`

HasPartnerParentId returns a boolean if a field has been set.

### GetPassReplaceFreq

`func (o *KalturaPartner) GetPassReplaceFreq() int32`

GetPassReplaceFreq returns the PassReplaceFreq field if non-nil, zero value otherwise.

### GetPassReplaceFreqOk

`func (o *KalturaPartner) GetPassReplaceFreqOk() (*int32, bool)`

GetPassReplaceFreqOk returns a tuple with the PassReplaceFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassReplaceFreq

`func (o *KalturaPartner) SetPassReplaceFreq(v int32)`

SetPassReplaceFreq sets PassReplaceFreq field to given value.

### HasPassReplaceFreq

`func (o *KalturaPartner) HasPassReplaceFreq() bool`

HasPassReplaceFreq returns a boolean if a field has been set.

### GetPasswordStructureValidations

`func (o *KalturaPartner) GetPasswordStructureValidations() []KalturaRegexItem`

GetPasswordStructureValidations returns the PasswordStructureValidations field if non-nil, zero value otherwise.

### GetPasswordStructureValidationsOk

`func (o *KalturaPartner) GetPasswordStructureValidationsOk() (*[]KalturaRegexItem, bool)`

GetPasswordStructureValidationsOk returns a tuple with the PasswordStructureValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStructureValidations

`func (o *KalturaPartner) SetPasswordStructureValidations(v []KalturaRegexItem)`

SetPasswordStructureValidations sets PasswordStructureValidations field to given value.

### HasPasswordStructureValidations

`func (o *KalturaPartner) HasPasswordStructureValidations() bool`

HasPasswordStructureValidations returns a boolean if a field has been set.

### GetPasswordStructureValidationsDescription

`func (o *KalturaPartner) GetPasswordStructureValidationsDescription() string`

GetPasswordStructureValidationsDescription returns the PasswordStructureValidationsDescription field if non-nil, zero value otherwise.

### GetPasswordStructureValidationsDescriptionOk

`func (o *KalturaPartner) GetPasswordStructureValidationsDescriptionOk() (*string, bool)`

GetPasswordStructureValidationsDescriptionOk returns a tuple with the PasswordStructureValidationsDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordStructureValidationsDescription

`func (o *KalturaPartner) SetPasswordStructureValidationsDescription(v string)`

SetPasswordStructureValidationsDescription sets PasswordStructureValidationsDescription field to given value.

### HasPasswordStructureValidationsDescription

`func (o *KalturaPartner) HasPasswordStructureValidationsDescription() bool`

HasPasswordStructureValidationsDescription returns a boolean if a field has been set.

### GetPhone

`func (o *KalturaPartner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *KalturaPartner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *KalturaPartner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *KalturaPartner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPublisherEnvironmentType

`func (o *KalturaPartner) GetPublisherEnvironmentType() int32`

GetPublisherEnvironmentType returns the PublisherEnvironmentType field if non-nil, zero value otherwise.

### GetPublisherEnvironmentTypeOk

`func (o *KalturaPartner) GetPublisherEnvironmentTypeOk() (*int32, bool)`

GetPublisherEnvironmentTypeOk returns a tuple with the PublisherEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherEnvironmentType

`func (o *KalturaPartner) SetPublisherEnvironmentType(v int32)`

SetPublisherEnvironmentType sets PublisherEnvironmentType field to given value.

### HasPublisherEnvironmentType

`func (o *KalturaPartner) HasPublisherEnvironmentType() bool`

HasPublisherEnvironmentType returns a boolean if a field has been set.

### GetPublishersQuota

`func (o *KalturaPartner) GetPublishersQuota() int32`

GetPublishersQuota returns the PublishersQuota field if non-nil, zero value otherwise.

### GetPublishersQuotaOk

`func (o *KalturaPartner) GetPublishersQuotaOk() (*int32, bool)`

GetPublishersQuotaOk returns a tuple with the PublishersQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishersQuota

`func (o *KalturaPartner) SetPublishersQuota(v int32)`

SetPublishersQuota sets PublishersQuota field to given value.

### HasPublishersQuota

`func (o *KalturaPartner) HasPublishersQuota() bool`

HasPublishersQuota returns a boolean if a field has been set.

### GetReferenceId

`func (o *KalturaPartner) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KalturaPartner) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KalturaPartner) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KalturaPartner) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetSecret

`func (o *KalturaPartner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *KalturaPartner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *KalturaPartner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *KalturaPartner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetState

`func (o *KalturaPartner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KalturaPartner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KalturaPartner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KalturaPartner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaPartner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaPartner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaPartner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaPartner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplatePartnerId

`func (o *KalturaPartner) GetTemplatePartnerId() int32`

GetTemplatePartnerId returns the TemplatePartnerId field if non-nil, zero value otherwise.

### GetTemplatePartnerIdOk

`func (o *KalturaPartner) GetTemplatePartnerIdOk() (*int32, bool)`

GetTemplatePartnerIdOk returns a tuple with the TemplatePartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePartnerId

`func (o *KalturaPartner) SetTemplatePartnerId(v int32)`

SetTemplatePartnerId sets TemplatePartnerId field to given value.

### HasTemplatePartnerId

`func (o *KalturaPartner) HasTemplatePartnerId() bool`

HasTemplatePartnerId returns a boolean if a field has been set.

### GetTimeAlignedRenditions

`func (o *KalturaPartner) GetTimeAlignedRenditions() bool`

GetTimeAlignedRenditions returns the TimeAlignedRenditions field if non-nil, zero value otherwise.

### GetTimeAlignedRenditionsOk

`func (o *KalturaPartner) GetTimeAlignedRenditionsOk() (*bool, bool)`

GetTimeAlignedRenditionsOk returns a tuple with the TimeAlignedRenditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAlignedRenditions

`func (o *KalturaPartner) SetTimeAlignedRenditions(v bool)`

SetTimeAlignedRenditions sets TimeAlignedRenditions field to given value.

### HasTimeAlignedRenditions

`func (o *KalturaPartner) HasTimeAlignedRenditions() bool`

HasTimeAlignedRenditions returns a boolean if a field has been set.

### GetTwoFactorAuthenticationMode

`func (o *KalturaPartner) GetTwoFactorAuthenticationMode() int32`

GetTwoFactorAuthenticationMode returns the TwoFactorAuthenticationMode field if non-nil, zero value otherwise.

### GetTwoFactorAuthenticationModeOk

`func (o *KalturaPartner) GetTwoFactorAuthenticationModeOk() (*int32, bool)`

GetTwoFactorAuthenticationModeOk returns a tuple with the TwoFactorAuthenticationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthenticationMode

`func (o *KalturaPartner) SetTwoFactorAuthenticationMode(v int32)`

SetTwoFactorAuthenticationMode sets TwoFactorAuthenticationMode field to given value.

### HasTwoFactorAuthenticationMode

`func (o *KalturaPartner) HasTwoFactorAuthenticationMode() bool`

HasTwoFactorAuthenticationMode returns a boolean if a field has been set.

### GetType

`func (o *KalturaPartner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaPartner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaPartner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaPartner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsageLimitWarning

`func (o *KalturaPartner) GetUsageLimitWarning() int32`

GetUsageLimitWarning returns the UsageLimitWarning field if non-nil, zero value otherwise.

### GetUsageLimitWarningOk

`func (o *KalturaPartner) GetUsageLimitWarningOk() (*int32, bool)`

GetUsageLimitWarningOk returns a tuple with the UsageLimitWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimitWarning

`func (o *KalturaPartner) SetUsageLimitWarning(v int32)`

SetUsageLimitWarning sets UsageLimitWarning field to given value.

### HasUsageLimitWarning

`func (o *KalturaPartner) HasUsageLimitWarning() bool`

HasUsageLimitWarning returns a boolean if a field has been set.

### GetUserLandingPage

`func (o *KalturaPartner) GetUserLandingPage() string`

GetUserLandingPage returns the UserLandingPage field if non-nil, zero value otherwise.

### GetUserLandingPageOk

`func (o *KalturaPartner) GetUserLandingPageOk() (*string, bool)`

GetUserLandingPageOk returns a tuple with the UserLandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLandingPage

`func (o *KalturaPartner) SetUserLandingPage(v string)`

SetUserLandingPage sets UserLandingPage field to given value.

### HasUserLandingPage

`func (o *KalturaPartner) HasUserLandingPage() bool`

HasUserLandingPage returns a boolean if a field has been set.

### GetWebsite

`func (o *KalturaPartner) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *KalturaPartner) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *KalturaPartner) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *KalturaPartner) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


