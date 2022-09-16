# KalturaSearchAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthData** | Pointer to **string** | The authentication data that further should be used for search | [optional] 
**LoginUrl** | Pointer to **string** | Login URL when user need to sign-in and authorize the search | [optional] 
**Message** | Pointer to **string** | Information when there was an error | [optional] 

## Methods

### NewKalturaSearchAuthData

`func NewKalturaSearchAuthData() *KalturaSearchAuthData`

NewKalturaSearchAuthData instantiates a new KalturaSearchAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSearchAuthDataWithDefaults

`func NewKalturaSearchAuthDataWithDefaults() *KalturaSearchAuthData`

NewKalturaSearchAuthDataWithDefaults instantiates a new KalturaSearchAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthData

`func (o *KalturaSearchAuthData) GetAuthData() string`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *KalturaSearchAuthData) GetAuthDataOk() (*string, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *KalturaSearchAuthData) SetAuthData(v string)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *KalturaSearchAuthData) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### GetLoginUrl

`func (o *KalturaSearchAuthData) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *KalturaSearchAuthData) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *KalturaSearchAuthData) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *KalturaSearchAuthData) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetMessage

`func (o *KalturaSearchAuthData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KalturaSearchAuthData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KalturaSearchAuthData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KalturaSearchAuthData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


