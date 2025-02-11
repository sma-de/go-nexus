# BowerGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 

## Methods

### NewBowerGroupRepositoryApiRequest

`func NewBowerGroupRepositoryApiRequest(name string, online bool, storage StorageAttributes, group GroupAttributes, ) *BowerGroupRepositoryApiRequest`

NewBowerGroupRepositoryApiRequest instantiates a new BowerGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBowerGroupRepositoryApiRequestWithDefaults

`func NewBowerGroupRepositoryApiRequestWithDefaults() *BowerGroupRepositoryApiRequest`

NewBowerGroupRepositoryApiRequestWithDefaults instantiates a new BowerGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BowerGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BowerGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BowerGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *BowerGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *BowerGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *BowerGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *BowerGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BowerGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BowerGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetGroup

`func (o *BowerGroupRepositoryApiRequest) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BowerGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BowerGroupRepositoryApiRequest) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


