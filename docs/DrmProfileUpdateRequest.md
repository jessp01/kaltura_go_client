# DrmProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrmProfile** | [**KalturaDrmProfile**](KalturaDrmProfile.md) |  | 
**DrmProfileId** | **int32** |  | 

## Methods

### NewDrmProfileUpdateRequest

`func NewDrmProfileUpdateRequest(drmProfile KalturaDrmProfile, drmProfileId int32, ) *DrmProfileUpdateRequest`

NewDrmProfileUpdateRequest instantiates a new DrmProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrmProfileUpdateRequestWithDefaults

`func NewDrmProfileUpdateRequestWithDefaults() *DrmProfileUpdateRequest`

NewDrmProfileUpdateRequestWithDefaults instantiates a new DrmProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrmProfile

`func (o *DrmProfileUpdateRequest) GetDrmProfile() KalturaDrmProfile`

GetDrmProfile returns the DrmProfile field if non-nil, zero value otherwise.

### GetDrmProfileOk

`func (o *DrmProfileUpdateRequest) GetDrmProfileOk() (*KalturaDrmProfile, bool)`

GetDrmProfileOk returns a tuple with the DrmProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmProfile

`func (o *DrmProfileUpdateRequest) SetDrmProfile(v KalturaDrmProfile)`

SetDrmProfile sets DrmProfile field to given value.


### GetDrmProfileId

`func (o *DrmProfileUpdateRequest) GetDrmProfileId() int32`

GetDrmProfileId returns the DrmProfileId field if non-nil, zero value otherwise.

### GetDrmProfileIdOk

`func (o *DrmProfileUpdateRequest) GetDrmProfileIdOk() (*int32, bool)`

GetDrmProfileIdOk returns a tuple with the DrmProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmProfileId

`func (o *DrmProfileUpdateRequest) SetDrmProfileId(v int32)`

SetDrmProfileId sets DrmProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


