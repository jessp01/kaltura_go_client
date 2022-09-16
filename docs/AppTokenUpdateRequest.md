# AppTokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppToken** | [**KalturaAppToken**](KalturaAppToken.md) |  | 
**Id** | **string** |  | 

## Methods

### NewAppTokenUpdateRequest

`func NewAppTokenUpdateRequest(appToken KalturaAppToken, id string, ) *AppTokenUpdateRequest`

NewAppTokenUpdateRequest instantiates a new AppTokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTokenUpdateRequestWithDefaults

`func NewAppTokenUpdateRequestWithDefaults() *AppTokenUpdateRequest`

NewAppTokenUpdateRequestWithDefaults instantiates a new AppTokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppToken

`func (o *AppTokenUpdateRequest) GetAppToken() KalturaAppToken`

GetAppToken returns the AppToken field if non-nil, zero value otherwise.

### GetAppTokenOk

`func (o *AppTokenUpdateRequest) GetAppTokenOk() (*KalturaAppToken, bool)`

GetAppTokenOk returns a tuple with the AppToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppToken

`func (o *AppTokenUpdateRequest) SetAppToken(v KalturaAppToken)`

SetAppToken sets AppToken field to given value.


### GetId

`func (o *AppTokenUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppTokenUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppTokenUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


