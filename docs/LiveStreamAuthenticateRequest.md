# LiveStreamAuthenticateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** |  | [optional] 
**EntryId** | **string** |  | 
**Hostname** | Pointer to **string** |  | [optional] 
**MediaServerIndex** | Pointer to **string** |  | [optional] 
**Token** | **string** |  | 

## Methods

### NewLiveStreamAuthenticateRequest

`func NewLiveStreamAuthenticateRequest(entryId string, token string, ) *LiveStreamAuthenticateRequest`

NewLiveStreamAuthenticateRequest instantiates a new LiveStreamAuthenticateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveStreamAuthenticateRequestWithDefaults

`func NewLiveStreamAuthenticateRequestWithDefaults() *LiveStreamAuthenticateRequest`

NewLiveStreamAuthenticateRequestWithDefaults instantiates a new LiveStreamAuthenticateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *LiveStreamAuthenticateRequest) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *LiveStreamAuthenticateRequest) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *LiveStreamAuthenticateRequest) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *LiveStreamAuthenticateRequest) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetEntryId

`func (o *LiveStreamAuthenticateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *LiveStreamAuthenticateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *LiveStreamAuthenticateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetHostname

`func (o *LiveStreamAuthenticateRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LiveStreamAuthenticateRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LiveStreamAuthenticateRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LiveStreamAuthenticateRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMediaServerIndex

`func (o *LiveStreamAuthenticateRequest) GetMediaServerIndex() string`

GetMediaServerIndex returns the MediaServerIndex field if non-nil, zero value otherwise.

### GetMediaServerIndexOk

`func (o *LiveStreamAuthenticateRequest) GetMediaServerIndexOk() (*string, bool)`

GetMediaServerIndexOk returns a tuple with the MediaServerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaServerIndex

`func (o *LiveStreamAuthenticateRequest) SetMediaServerIndex(v string)`

SetMediaServerIndex sets MediaServerIndex field to given value.

### HasMediaServerIndex

`func (o *LiveStreamAuthenticateRequest) HasMediaServerIndex() bool`

HasMediaServerIndex returns a boolean if a field has been set.

### GetToken

`func (o *LiveStreamAuthenticateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LiveStreamAuthenticateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LiveStreamAuthenticateRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


