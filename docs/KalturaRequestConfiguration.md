# KalturaRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ks** | Pointer to **string** | Kaltura API session | [optional] 
**PartnerId** | Pointer to **int32** | Impersonated partner id | [optional] 
**ResponseProfile** | Pointer to [**KalturaBaseResponseProfile**](KalturaBaseResponseProfile.md) |  | [optional] 

## Methods

### NewKalturaRequestConfiguration

`func NewKalturaRequestConfiguration() *KalturaRequestConfiguration`

NewKalturaRequestConfiguration instantiates a new KalturaRequestConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaRequestConfigurationWithDefaults

`func NewKalturaRequestConfigurationWithDefaults() *KalturaRequestConfiguration`

NewKalturaRequestConfigurationWithDefaults instantiates a new KalturaRequestConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKs

`func (o *KalturaRequestConfiguration) GetKs() string`

GetKs returns the Ks field if non-nil, zero value otherwise.

### GetKsOk

`func (o *KalturaRequestConfiguration) GetKsOk() (*string, bool)`

GetKsOk returns a tuple with the Ks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKs

`func (o *KalturaRequestConfiguration) SetKs(v string)`

SetKs sets Ks field to given value.

### HasKs

`func (o *KalturaRequestConfiguration) HasKs() bool`

HasKs returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaRequestConfiguration) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaRequestConfiguration) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaRequestConfiguration) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaRequestConfiguration) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetResponseProfile

`func (o *KalturaRequestConfiguration) GetResponseProfile() KalturaBaseResponseProfile`

GetResponseProfile returns the ResponseProfile field if non-nil, zero value otherwise.

### GetResponseProfileOk

`func (o *KalturaRequestConfiguration) GetResponseProfileOk() (*KalturaBaseResponseProfile, bool)`

GetResponseProfileOk returns a tuple with the ResponseProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfile

`func (o *KalturaRequestConfiguration) SetResponseProfile(v KalturaBaseResponseProfile)`

SetResponseProfile sets ResponseProfile field to given value.

### HasResponseProfile

`func (o *KalturaRequestConfiguration) HasResponseProfile() bool`

HasResponseProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


