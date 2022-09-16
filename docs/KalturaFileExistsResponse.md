# KalturaFileExistsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** | Indicates if the file exists | [optional] 
**SizeOk** | Pointer to **bool** | Indicates if the file size is right | [optional] 

## Methods

### NewKalturaFileExistsResponse

`func NewKalturaFileExistsResponse() *KalturaFileExistsResponse`

NewKalturaFileExistsResponse instantiates a new KalturaFileExistsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFileExistsResponseWithDefaults

`func NewKalturaFileExistsResponseWithDefaults() *KalturaFileExistsResponse`

NewKalturaFileExistsResponseWithDefaults instantiates a new KalturaFileExistsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *KalturaFileExistsResponse) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *KalturaFileExistsResponse) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *KalturaFileExistsResponse) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *KalturaFileExistsResponse) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetSizeOk

`func (o *KalturaFileExistsResponse) GetSizeOk() bool`

GetSizeOk returns the SizeOk field if non-nil, zero value otherwise.

### GetSizeOkOk

`func (o *KalturaFileExistsResponse) GetSizeOkOk() (*bool, bool)`

GetSizeOkOk returns a tuple with the SizeOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeOk

`func (o *KalturaFileExistsResponse) SetSizeOk(v bool)`

SetSizeOk sets SizeOk field to given value.

### HasSizeOk

`func (o *KalturaFileExistsResponse) HasSizeOk() bool`

HasSizeOk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


