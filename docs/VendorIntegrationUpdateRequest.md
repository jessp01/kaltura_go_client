# VendorIntegrationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**IntegrationSetting** | [**KalturaIntegrationSetting**](KalturaIntegrationSetting.md) |  | 

## Methods

### NewVendorIntegrationUpdateRequest

`func NewVendorIntegrationUpdateRequest(id int32, integrationSetting KalturaIntegrationSetting, ) *VendorIntegrationUpdateRequest`

NewVendorIntegrationUpdateRequest instantiates a new VendorIntegrationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorIntegrationUpdateRequestWithDefaults

`func NewVendorIntegrationUpdateRequestWithDefaults() *VendorIntegrationUpdateRequest`

NewVendorIntegrationUpdateRequestWithDefaults instantiates a new VendorIntegrationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorIntegrationUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorIntegrationUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorIntegrationUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetIntegrationSetting

`func (o *VendorIntegrationUpdateRequest) GetIntegrationSetting() KalturaIntegrationSetting`

GetIntegrationSetting returns the IntegrationSetting field if non-nil, zero value otherwise.

### GetIntegrationSettingOk

`func (o *VendorIntegrationUpdateRequest) GetIntegrationSettingOk() (*KalturaIntegrationSetting, bool)`

GetIntegrationSettingOk returns a tuple with the IntegrationSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationSetting

`func (o *VendorIntegrationUpdateRequest) SetIntegrationSetting(v KalturaIntegrationSetting)`

SetIntegrationSetting sets IntegrationSetting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


