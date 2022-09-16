# PartnerGetSecretsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminEmail** | **string** |  | 
**CmsPassword** | **string** |  | 
**Otp** | Pointer to **string** |  | [optional] 
**PartnerId** | **int32** |  | 

## Methods

### NewPartnerGetSecretsRequest

`func NewPartnerGetSecretsRequest(adminEmail string, cmsPassword string, partnerId int32, ) *PartnerGetSecretsRequest`

NewPartnerGetSecretsRequest instantiates a new PartnerGetSecretsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerGetSecretsRequestWithDefaults

`func NewPartnerGetSecretsRequestWithDefaults() *PartnerGetSecretsRequest`

NewPartnerGetSecretsRequestWithDefaults instantiates a new PartnerGetSecretsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminEmail

`func (o *PartnerGetSecretsRequest) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *PartnerGetSecretsRequest) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *PartnerGetSecretsRequest) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.


### GetCmsPassword

`func (o *PartnerGetSecretsRequest) GetCmsPassword() string`

GetCmsPassword returns the CmsPassword field if non-nil, zero value otherwise.

### GetCmsPasswordOk

`func (o *PartnerGetSecretsRequest) GetCmsPasswordOk() (*string, bool)`

GetCmsPasswordOk returns a tuple with the CmsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmsPassword

`func (o *PartnerGetSecretsRequest) SetCmsPassword(v string)`

SetCmsPassword sets CmsPassword field to given value.


### GetOtp

`func (o *PartnerGetSecretsRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *PartnerGetSecretsRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *PartnerGetSecretsRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *PartnerGetSecretsRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPartnerId

`func (o *PartnerGetSecretsRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PartnerGetSecretsRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PartnerGetSecretsRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


