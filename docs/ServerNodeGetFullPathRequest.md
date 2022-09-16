# ServerNodeGetFullPathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryFormat** | Pointer to **string** |  | [optional] 
**DeliveryType** | Pointer to **string** |  | [optional] 
**HostName** | **string** |  | 
**Protocol** | Pointer to **string** |  | [optional] [default to "http"]

## Methods

### NewServerNodeGetFullPathRequest

`func NewServerNodeGetFullPathRequest(hostName string, ) *ServerNodeGetFullPathRequest`

NewServerNodeGetFullPathRequest instantiates a new ServerNodeGetFullPathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerNodeGetFullPathRequestWithDefaults

`func NewServerNodeGetFullPathRequestWithDefaults() *ServerNodeGetFullPathRequest`

NewServerNodeGetFullPathRequestWithDefaults instantiates a new ServerNodeGetFullPathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryFormat

`func (o *ServerNodeGetFullPathRequest) GetDeliveryFormat() string`

GetDeliveryFormat returns the DeliveryFormat field if non-nil, zero value otherwise.

### GetDeliveryFormatOk

`func (o *ServerNodeGetFullPathRequest) GetDeliveryFormatOk() (*string, bool)`

GetDeliveryFormatOk returns a tuple with the DeliveryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFormat

`func (o *ServerNodeGetFullPathRequest) SetDeliveryFormat(v string)`

SetDeliveryFormat sets DeliveryFormat field to given value.

### HasDeliveryFormat

`func (o *ServerNodeGetFullPathRequest) HasDeliveryFormat() bool`

HasDeliveryFormat returns a boolean if a field has been set.

### GetDeliveryType

`func (o *ServerNodeGetFullPathRequest) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *ServerNodeGetFullPathRequest) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *ServerNodeGetFullPathRequest) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *ServerNodeGetFullPathRequest) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetHostName

`func (o *ServerNodeGetFullPathRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ServerNodeGetFullPathRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ServerNodeGetFullPathRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetProtocol

`func (o *ServerNodeGetFullPathRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServerNodeGetFullPathRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServerNodeGetFullPathRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServerNodeGetFullPathRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


