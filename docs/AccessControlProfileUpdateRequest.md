# AccessControlProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControlProfile** | [**KalturaAccessControlProfile**](KalturaAccessControlProfile.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewAccessControlProfileUpdateRequest

`func NewAccessControlProfileUpdateRequest(accessControlProfile KalturaAccessControlProfile, id int32, ) *AccessControlProfileUpdateRequest`

NewAccessControlProfileUpdateRequest instantiates a new AccessControlProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessControlProfileUpdateRequestWithDefaults

`func NewAccessControlProfileUpdateRequestWithDefaults() *AccessControlProfileUpdateRequest`

NewAccessControlProfileUpdateRequestWithDefaults instantiates a new AccessControlProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControlProfile

`func (o *AccessControlProfileUpdateRequest) GetAccessControlProfile() KalturaAccessControlProfile`

GetAccessControlProfile returns the AccessControlProfile field if non-nil, zero value otherwise.

### GetAccessControlProfileOk

`func (o *AccessControlProfileUpdateRequest) GetAccessControlProfileOk() (*KalturaAccessControlProfile, bool)`

GetAccessControlProfileOk returns a tuple with the AccessControlProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlProfile

`func (o *AccessControlProfileUpdateRequest) SetAccessControlProfile(v KalturaAccessControlProfile)`

SetAccessControlProfile sets AccessControlProfile field to given value.


### GetId

`func (o *AccessControlProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessControlProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessControlProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


