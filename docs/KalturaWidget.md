# KalturaWidget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddEmbedHtml5Support** | Pointer to **bool** | Addes the HTML5 script line to the widget&#39;s embed code | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EnforceEntitlement** | Pointer to **bool** | Should enforce entitlement on feed entries | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerData** | Pointer to **string** | Can be used to store various partner related data as a string | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PrivacyContext** | Pointer to **string** | Set privacy context for search entries that assiged to private and public categories within a category privacy context. | [optional] 
**Privileges** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **string** |  | [optional] 
**RootWidgetId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**SecurityPolicy** | Pointer to **int32** |  | [optional] 
**SecurityType** | Pointer to **int32** | Enum Type: &#x60;KalturaWidgetSecurityType&#x60; | [optional] 
**SourceWidgetId** | Pointer to **string** |  | [optional] 
**UiConfId** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**WidgetHTML** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaWidget

`func NewKalturaWidget() *KalturaWidget`

NewKalturaWidget instantiates a new KalturaWidget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaWidgetWithDefaults

`func NewKalturaWidgetWithDefaults() *KalturaWidget`

NewKalturaWidgetWithDefaults instantiates a new KalturaWidget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddEmbedHtml5Support

`func (o *KalturaWidget) GetAddEmbedHtml5Support() bool`

GetAddEmbedHtml5Support returns the AddEmbedHtml5Support field if non-nil, zero value otherwise.

### GetAddEmbedHtml5SupportOk

`func (o *KalturaWidget) GetAddEmbedHtml5SupportOk() (*bool, bool)`

GetAddEmbedHtml5SupportOk returns a tuple with the AddEmbedHtml5Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddEmbedHtml5Support

`func (o *KalturaWidget) SetAddEmbedHtml5Support(v bool)`

SetAddEmbedHtml5Support sets AddEmbedHtml5Support field to given value.

### HasAddEmbedHtml5Support

`func (o *KalturaWidget) HasAddEmbedHtml5Support() bool`

HasAddEmbedHtml5Support returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaWidget) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaWidget) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaWidget) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaWidget) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnforceEntitlement

`func (o *KalturaWidget) GetEnforceEntitlement() bool`

GetEnforceEntitlement returns the EnforceEntitlement field if non-nil, zero value otherwise.

### GetEnforceEntitlementOk

`func (o *KalturaWidget) GetEnforceEntitlementOk() (*bool, bool)`

GetEnforceEntitlementOk returns a tuple with the EnforceEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceEntitlement

`func (o *KalturaWidget) SetEnforceEntitlement(v bool)`

SetEnforceEntitlement sets EnforceEntitlement field to given value.

### HasEnforceEntitlement

`func (o *KalturaWidget) HasEnforceEntitlement() bool`

HasEnforceEntitlement returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaWidget) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaWidget) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaWidget) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaWidget) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetId

`func (o *KalturaWidget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaWidget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaWidget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaWidget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaWidget) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaWidget) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaWidget) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaWidget) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaWidget) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaWidget) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaWidget) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaWidget) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPrivacyContext

`func (o *KalturaWidget) GetPrivacyContext() string`

GetPrivacyContext returns the PrivacyContext field if non-nil, zero value otherwise.

### GetPrivacyContextOk

`func (o *KalturaWidget) GetPrivacyContextOk() (*string, bool)`

GetPrivacyContextOk returns a tuple with the PrivacyContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyContext

`func (o *KalturaWidget) SetPrivacyContext(v string)`

SetPrivacyContext sets PrivacyContext field to given value.

### HasPrivacyContext

`func (o *KalturaWidget) HasPrivacyContext() bool`

HasPrivacyContext returns a boolean if a field has been set.

### GetPrivileges

`func (o *KalturaWidget) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *KalturaWidget) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *KalturaWidget) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *KalturaWidget) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetRoles

`func (o *KalturaWidget) GetRoles() string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *KalturaWidget) GetRolesOk() (*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *KalturaWidget) SetRoles(v string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *KalturaWidget) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRootWidgetId

`func (o *KalturaWidget) GetRootWidgetId() string`

GetRootWidgetId returns the RootWidgetId field if non-nil, zero value otherwise.

### GetRootWidgetIdOk

`func (o *KalturaWidget) GetRootWidgetIdOk() (*string, bool)`

GetRootWidgetIdOk returns a tuple with the RootWidgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWidgetId

`func (o *KalturaWidget) SetRootWidgetId(v string)`

SetRootWidgetId sets RootWidgetId field to given value.

### HasRootWidgetId

`func (o *KalturaWidget) HasRootWidgetId() bool`

HasRootWidgetId returns a boolean if a field has been set.

### GetSecurityPolicy

`func (o *KalturaWidget) GetSecurityPolicy() int32`

GetSecurityPolicy returns the SecurityPolicy field if non-nil, zero value otherwise.

### GetSecurityPolicyOk

`func (o *KalturaWidget) GetSecurityPolicyOk() (*int32, bool)`

GetSecurityPolicyOk returns a tuple with the SecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicy

`func (o *KalturaWidget) SetSecurityPolicy(v int32)`

SetSecurityPolicy sets SecurityPolicy field to given value.

### HasSecurityPolicy

`func (o *KalturaWidget) HasSecurityPolicy() bool`

HasSecurityPolicy returns a boolean if a field has been set.

### GetSecurityType

`func (o *KalturaWidget) GetSecurityType() int32`

GetSecurityType returns the SecurityType field if non-nil, zero value otherwise.

### GetSecurityTypeOk

`func (o *KalturaWidget) GetSecurityTypeOk() (*int32, bool)`

GetSecurityTypeOk returns a tuple with the SecurityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityType

`func (o *KalturaWidget) SetSecurityType(v int32)`

SetSecurityType sets SecurityType field to given value.

### HasSecurityType

`func (o *KalturaWidget) HasSecurityType() bool`

HasSecurityType returns a boolean if a field has been set.

### GetSourceWidgetId

`func (o *KalturaWidget) GetSourceWidgetId() string`

GetSourceWidgetId returns the SourceWidgetId field if non-nil, zero value otherwise.

### GetSourceWidgetIdOk

`func (o *KalturaWidget) GetSourceWidgetIdOk() (*string, bool)`

GetSourceWidgetIdOk returns a tuple with the SourceWidgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceWidgetId

`func (o *KalturaWidget) SetSourceWidgetId(v string)`

SetSourceWidgetId sets SourceWidgetId field to given value.

### HasSourceWidgetId

`func (o *KalturaWidget) HasSourceWidgetId() bool`

HasSourceWidgetId returns a boolean if a field has been set.

### GetUiConfId

`func (o *KalturaWidget) GetUiConfId() int32`

GetUiConfId returns the UiConfId field if non-nil, zero value otherwise.

### GetUiConfIdOk

`func (o *KalturaWidget) GetUiConfIdOk() (*int32, bool)`

GetUiConfIdOk returns a tuple with the UiConfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConfId

`func (o *KalturaWidget) SetUiConfId(v int32)`

SetUiConfId sets UiConfId field to given value.

### HasUiConfId

`func (o *KalturaWidget) HasUiConfId() bool`

HasUiConfId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaWidget) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaWidget) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaWidget) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaWidget) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWidgetHTML

`func (o *KalturaWidget) GetWidgetHTML() string`

GetWidgetHTML returns the WidgetHTML field if non-nil, zero value otherwise.

### GetWidgetHTMLOk

`func (o *KalturaWidget) GetWidgetHTMLOk() (*string, bool)`

GetWidgetHTMLOk returns a tuple with the WidgetHTML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetHTML

`func (o *KalturaWidget) SetWidgetHTML(v string)`

SetWidgetHTML sets WidgetHTML field to given value.

### HasWidgetHTML

`func (o *KalturaWidget) HasWidgetHTML() bool`

HasWidgetHTML returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


