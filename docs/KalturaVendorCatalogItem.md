# KalturaVendorCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowResubmission** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**EngineType** | Pointer to **string** | Enum Type: &#x60;KalturaReachVendorEngineType&#x60;  Property showing the catalog item&#39;s engine type, in case a vendor can offer the same service via different engines. | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to [**KalturaVendorCatalogItemPricing**](KalturaVendorCatalogItemPricing.md) |  | [optional] 
**ServiceFeature** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorServiceFeature&#x60; | [optional] [readonly] 
**ServiceType** | Pointer to **int32** | Enum Type: &#x60;KalturaVendorServiceType&#x60; | [optional] 
**SourceLanguage** | Pointer to **string** | Enum Type: &#x60;KalturaCatalogItemLanguage&#x60; | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorCatalogItemStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**TurnAroundTime** | Pointer to **int32** | Enum Type: &#x60;KalturaVendorServiceTurnAroundTime&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**VendorPartnerId** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaVendorCatalogItem

`func NewKalturaVendorCatalogItem() *KalturaVendorCatalogItem`

NewKalturaVendorCatalogItem instantiates a new KalturaVendorCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaVendorCatalogItemWithDefaults

`func NewKalturaVendorCatalogItemWithDefaults() *KalturaVendorCatalogItem`

NewKalturaVendorCatalogItemWithDefaults instantiates a new KalturaVendorCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowResubmission

`func (o *KalturaVendorCatalogItem) GetAllowResubmission() bool`

GetAllowResubmission returns the AllowResubmission field if non-nil, zero value otherwise.

### GetAllowResubmissionOk

`func (o *KalturaVendorCatalogItem) GetAllowResubmissionOk() (*bool, bool)`

GetAllowResubmissionOk returns a tuple with the AllowResubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowResubmission

`func (o *KalturaVendorCatalogItem) SetAllowResubmission(v bool)`

SetAllowResubmission sets AllowResubmission field to given value.

### HasAllowResubmission

`func (o *KalturaVendorCatalogItem) HasAllowResubmission() bool`

HasAllowResubmission returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaVendorCatalogItem) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaVendorCatalogItem) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaVendorCatalogItem) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaVendorCatalogItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEngineType

`func (o *KalturaVendorCatalogItem) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *KalturaVendorCatalogItem) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *KalturaVendorCatalogItem) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *KalturaVendorCatalogItem) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### GetId

`func (o *KalturaVendorCatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaVendorCatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaVendorCatalogItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaVendorCatalogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaVendorCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaVendorCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaVendorCatalogItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaVendorCatalogItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaVendorCatalogItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaVendorCatalogItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaVendorCatalogItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaVendorCatalogItem) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPricing

`func (o *KalturaVendorCatalogItem) GetPricing() KalturaVendorCatalogItemPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *KalturaVendorCatalogItem) GetPricingOk() (*KalturaVendorCatalogItemPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *KalturaVendorCatalogItem) SetPricing(v KalturaVendorCatalogItemPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *KalturaVendorCatalogItem) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetServiceFeature

`func (o *KalturaVendorCatalogItem) GetServiceFeature() int32`

GetServiceFeature returns the ServiceFeature field if non-nil, zero value otherwise.

### GetServiceFeatureOk

`func (o *KalturaVendorCatalogItem) GetServiceFeatureOk() (*int32, bool)`

GetServiceFeatureOk returns a tuple with the ServiceFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFeature

`func (o *KalturaVendorCatalogItem) SetServiceFeature(v int32)`

SetServiceFeature sets ServiceFeature field to given value.

### HasServiceFeature

`func (o *KalturaVendorCatalogItem) HasServiceFeature() bool`

HasServiceFeature returns a boolean if a field has been set.

### GetServiceType

`func (o *KalturaVendorCatalogItem) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *KalturaVendorCatalogItem) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *KalturaVendorCatalogItem) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *KalturaVendorCatalogItem) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSourceLanguage

`func (o *KalturaVendorCatalogItem) GetSourceLanguage() string`

GetSourceLanguage returns the SourceLanguage field if non-nil, zero value otherwise.

### GetSourceLanguageOk

`func (o *KalturaVendorCatalogItem) GetSourceLanguageOk() (*string, bool)`

GetSourceLanguageOk returns a tuple with the SourceLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLanguage

`func (o *KalturaVendorCatalogItem) SetSourceLanguage(v string)`

SetSourceLanguage sets SourceLanguage field to given value.

### HasSourceLanguage

`func (o *KalturaVendorCatalogItem) HasSourceLanguage() bool`

HasSourceLanguage returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaVendorCatalogItem) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaVendorCatalogItem) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaVendorCatalogItem) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaVendorCatalogItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaVendorCatalogItem) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaVendorCatalogItem) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaVendorCatalogItem) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaVendorCatalogItem) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTurnAroundTime

`func (o *KalturaVendorCatalogItem) GetTurnAroundTime() int32`

GetTurnAroundTime returns the TurnAroundTime field if non-nil, zero value otherwise.

### GetTurnAroundTimeOk

`func (o *KalturaVendorCatalogItem) GetTurnAroundTimeOk() (*int32, bool)`

GetTurnAroundTimeOk returns a tuple with the TurnAroundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnAroundTime

`func (o *KalturaVendorCatalogItem) SetTurnAroundTime(v int32)`

SetTurnAroundTime sets TurnAroundTime field to given value.

### HasTurnAroundTime

`func (o *KalturaVendorCatalogItem) HasTurnAroundTime() bool`

HasTurnAroundTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaVendorCatalogItem) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaVendorCatalogItem) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaVendorCatalogItem) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaVendorCatalogItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVendorPartnerId

`func (o *KalturaVendorCatalogItem) GetVendorPartnerId() int32`

GetVendorPartnerId returns the VendorPartnerId field if non-nil, zero value otherwise.

### GetVendorPartnerIdOk

`func (o *KalturaVendorCatalogItem) GetVendorPartnerIdOk() (*int32, bool)`

GetVendorPartnerIdOk returns a tuple with the VendorPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartnerId

`func (o *KalturaVendorCatalogItem) SetVendorPartnerId(v int32)`

SetVendorPartnerId sets VendorPartnerId field to given value.

### HasVendorPartnerId

`func (o *KalturaVendorCatalogItem) HasVendorPartnerId() bool`

HasVendorPartnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


