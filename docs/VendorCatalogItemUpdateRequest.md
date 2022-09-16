# VendorCatalogItemUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**VendorCatalogItem** | [**KalturaVendorCatalogItem**](KalturaVendorCatalogItem.md) |  | 

## Methods

### NewVendorCatalogItemUpdateRequest

`func NewVendorCatalogItemUpdateRequest(id int32, vendorCatalogItem KalturaVendorCatalogItem, ) *VendorCatalogItemUpdateRequest`

NewVendorCatalogItemUpdateRequest instantiates a new VendorCatalogItemUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorCatalogItemUpdateRequestWithDefaults

`func NewVendorCatalogItemUpdateRequestWithDefaults() *VendorCatalogItemUpdateRequest`

NewVendorCatalogItemUpdateRequestWithDefaults instantiates a new VendorCatalogItemUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorCatalogItemUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorCatalogItemUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorCatalogItemUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVendorCatalogItem

`func (o *VendorCatalogItemUpdateRequest) GetVendorCatalogItem() KalturaVendorCatalogItem`

GetVendorCatalogItem returns the VendorCatalogItem field if non-nil, zero value otherwise.

### GetVendorCatalogItemOk

`func (o *VendorCatalogItemUpdateRequest) GetVendorCatalogItemOk() (*KalturaVendorCatalogItem, bool)`

GetVendorCatalogItemOk returns a tuple with the VendorCatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCatalogItem

`func (o *VendorCatalogItemUpdateRequest) SetVendorCatalogItem(v KalturaVendorCatalogItem)`

SetVendorCatalogItem sets VendorCatalogItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


