# VendorIntegrationAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Integration** | [**KalturaIntegrationSetting**](KalturaIntegrationSetting.md) |  | 
**RemoteId** | **string** |  | 

## Methods

### NewVendorIntegrationAddRequest

`func NewVendorIntegrationAddRequest(integration KalturaIntegrationSetting, remoteId string, ) *VendorIntegrationAddRequest`

NewVendorIntegrationAddRequest instantiates a new VendorIntegrationAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorIntegrationAddRequestWithDefaults

`func NewVendorIntegrationAddRequestWithDefaults() *VendorIntegrationAddRequest`

NewVendorIntegrationAddRequestWithDefaults instantiates a new VendorIntegrationAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegration

`func (o *VendorIntegrationAddRequest) GetIntegration() KalturaIntegrationSetting`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *VendorIntegrationAddRequest) GetIntegrationOk() (*KalturaIntegrationSetting, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *VendorIntegrationAddRequest) SetIntegration(v KalturaIntegrationSetting)`

SetIntegration sets Integration field to given value.


### GetRemoteId

`func (o *VendorIntegrationAddRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *VendorIntegrationAddRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *VendorIntegrationAddRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


