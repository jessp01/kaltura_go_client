# PartnerRegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmsPassword** | Pointer to **string** |  | [optional] 
**Partner** | [**KalturaPartner**](KalturaPartner.md) |  | 
**Silent** | Pointer to **bool** |  | [optional] [default to false]
**TemplatePartnerId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPartnerRegisterRequest

`func NewPartnerRegisterRequest(partner KalturaPartner, ) *PartnerRegisterRequest`

NewPartnerRegisterRequest instantiates a new PartnerRegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerRegisterRequestWithDefaults

`func NewPartnerRegisterRequestWithDefaults() *PartnerRegisterRequest`

NewPartnerRegisterRequestWithDefaults instantiates a new PartnerRegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmsPassword

`func (o *PartnerRegisterRequest) GetCmsPassword() string`

GetCmsPassword returns the CmsPassword field if non-nil, zero value otherwise.

### GetCmsPasswordOk

`func (o *PartnerRegisterRequest) GetCmsPasswordOk() (*string, bool)`

GetCmsPasswordOk returns a tuple with the CmsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmsPassword

`func (o *PartnerRegisterRequest) SetCmsPassword(v string)`

SetCmsPassword sets CmsPassword field to given value.

### HasCmsPassword

`func (o *PartnerRegisterRequest) HasCmsPassword() bool`

HasCmsPassword returns a boolean if a field has been set.

### GetPartner

`func (o *PartnerRegisterRequest) GetPartner() KalturaPartner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *PartnerRegisterRequest) GetPartnerOk() (*KalturaPartner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *PartnerRegisterRequest) SetPartner(v KalturaPartner)`

SetPartner sets Partner field to given value.


### GetSilent

`func (o *PartnerRegisterRequest) GetSilent() bool`

GetSilent returns the Silent field if non-nil, zero value otherwise.

### GetSilentOk

`func (o *PartnerRegisterRequest) GetSilentOk() (*bool, bool)`

GetSilentOk returns a tuple with the Silent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilent

`func (o *PartnerRegisterRequest) SetSilent(v bool)`

SetSilent sets Silent field to given value.

### HasSilent

`func (o *PartnerRegisterRequest) HasSilent() bool`

HasSilent returns a boolean if a field has been set.

### GetTemplatePartnerId

`func (o *PartnerRegisterRequest) GetTemplatePartnerId() int32`

GetTemplatePartnerId returns the TemplatePartnerId field if non-nil, zero value otherwise.

### GetTemplatePartnerIdOk

`func (o *PartnerRegisterRequest) GetTemplatePartnerIdOk() (*int32, bool)`

GetTemplatePartnerIdOk returns a tuple with the TemplatePartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePartnerId

`func (o *PartnerRegisterRequest) SetTemplatePartnerId(v int32)`

SetTemplatePartnerId sets TemplatePartnerId field to given value.

### HasTemplatePartnerId

`func (o *PartnerRegisterRequest) HasTemplatePartnerId() bool`

HasTemplatePartnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


