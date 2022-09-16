# DrmPolicyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrmPolicy** | [**KalturaDrmPolicy**](KalturaDrmPolicy.md) |  | 
**DrmPolicyId** | **int32** |  | 

## Methods

### NewDrmPolicyUpdateRequest

`func NewDrmPolicyUpdateRequest(drmPolicy KalturaDrmPolicy, drmPolicyId int32, ) *DrmPolicyUpdateRequest`

NewDrmPolicyUpdateRequest instantiates a new DrmPolicyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrmPolicyUpdateRequestWithDefaults

`func NewDrmPolicyUpdateRequestWithDefaults() *DrmPolicyUpdateRequest`

NewDrmPolicyUpdateRequestWithDefaults instantiates a new DrmPolicyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrmPolicy

`func (o *DrmPolicyUpdateRequest) GetDrmPolicy() KalturaDrmPolicy`

GetDrmPolicy returns the DrmPolicy field if non-nil, zero value otherwise.

### GetDrmPolicyOk

`func (o *DrmPolicyUpdateRequest) GetDrmPolicyOk() (*KalturaDrmPolicy, bool)`

GetDrmPolicyOk returns a tuple with the DrmPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmPolicy

`func (o *DrmPolicyUpdateRequest) SetDrmPolicy(v KalturaDrmPolicy)`

SetDrmPolicy sets DrmPolicy field to given value.


### GetDrmPolicyId

`func (o *DrmPolicyUpdateRequest) GetDrmPolicyId() int32`

GetDrmPolicyId returns the DrmPolicyId field if non-nil, zero value otherwise.

### GetDrmPolicyIdOk

`func (o *DrmPolicyUpdateRequest) GetDrmPolicyIdOk() (*int32, bool)`

GetDrmPolicyIdOk returns a tuple with the DrmPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmPolicyId

`func (o *DrmPolicyUpdateRequest) SetDrmPolicyId(v int32)`

SetDrmPolicyId sets DrmPolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


