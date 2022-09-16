# SystemPartnerUpdateConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**KalturaSystemPartnerConfiguration**](KalturaSystemPartnerConfiguration.md) |  | 
**PId** | **int32** |  | 

## Methods

### NewSystemPartnerUpdateConfigurationRequest

`func NewSystemPartnerUpdateConfigurationRequest(configuration KalturaSystemPartnerConfiguration, pId int32, ) *SystemPartnerUpdateConfigurationRequest`

NewSystemPartnerUpdateConfigurationRequest instantiates a new SystemPartnerUpdateConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemPartnerUpdateConfigurationRequestWithDefaults

`func NewSystemPartnerUpdateConfigurationRequestWithDefaults() *SystemPartnerUpdateConfigurationRequest`

NewSystemPartnerUpdateConfigurationRequestWithDefaults instantiates a new SystemPartnerUpdateConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *SystemPartnerUpdateConfigurationRequest) GetConfiguration() KalturaSystemPartnerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SystemPartnerUpdateConfigurationRequest) GetConfigurationOk() (*KalturaSystemPartnerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SystemPartnerUpdateConfigurationRequest) SetConfiguration(v KalturaSystemPartnerConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetPId

`func (o *SystemPartnerUpdateConfigurationRequest) GetPId() int32`

GetPId returns the PId field if non-nil, zero value otherwise.

### GetPIdOk

`func (o *SystemPartnerUpdateConfigurationRequest) GetPIdOk() (*int32, bool)`

GetPIdOk returns a tuple with the PId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPId

`func (o *SystemPartnerUpdateConfigurationRequest) SetPId(v int32)`

SetPId sets PId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


