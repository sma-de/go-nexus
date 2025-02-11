# MavenHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Maven** | [**MavenAttributes**](MavenAttributes.md) |  | 

## Methods

### NewMavenHostedApiRepository

`func NewMavenHostedApiRepository(online bool, storage HostedStorageAttributes, maven MavenAttributes, ) *MavenHostedApiRepository`

NewMavenHostedApiRepository instantiates a new MavenHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenHostedApiRepositoryWithDefaults

`func NewMavenHostedApiRepositoryWithDefaults() *MavenHostedApiRepository`

NewMavenHostedApiRepositoryWithDefaults instantiates a new MavenHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MavenHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MavenHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MavenHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MavenHostedApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *MavenHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *MavenHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *MavenHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *MavenHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *MavenHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *MavenHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *MavenHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *MavenHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *MavenHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *MavenHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *MavenHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *MavenHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *MavenHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *MavenHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetMaven

`func (o *MavenHostedApiRepository) GetMaven() MavenAttributes`

GetMaven returns the Maven field if non-nil, zero value otherwise.

### GetMavenOk

`func (o *MavenHostedApiRepository) GetMavenOk() (*MavenAttributes, bool)`

GetMavenOk returns a tuple with the Maven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven

`func (o *MavenHostedApiRepository) SetMaven(v MavenAttributes)`

SetMaven sets Maven field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


