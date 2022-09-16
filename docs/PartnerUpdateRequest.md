# PartnerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEmpty** | Pointer to **bool** |  | [optional] [default to false]
**Partner** | [**KalturaPartner**](KalturaPartner.md) |  | 

## Methods

### NewPartnerUpdateRequest

`func NewPartnerUpdateRequest(partner KalturaPartner, ) *PartnerUpdateRequest`

NewPartnerUpdateRequest instantiates a new PartnerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerUpdateRequestWithDefaults

`func NewPartnerUpdateRequestWithDefaults() *PartnerUpdateRequest`

NewPartnerUpdateRequestWithDefaults instantiates a new PartnerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEmpty

`func (o *PartnerUpdateRequest) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *PartnerUpdateRequest) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *PartnerUpdateRequest) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *PartnerUpdateRequest) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetPartner

`func (o *PartnerUpdateRequest) GetPartner() KalturaPartner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *PartnerUpdateRequest) GetPartnerOk() (*KalturaPartner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *PartnerUpdateRequest) SetPartner(v KalturaPartner)`

SetPartner sets Partner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


