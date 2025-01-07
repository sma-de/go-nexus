/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.75.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiCreateUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCreateUser{}

// ApiCreateUser struct for ApiCreateUser
type ApiCreateUser struct {
	// The userid which is required for login. This value cannot be changed.
	UserId *string `json:"userId,omitempty"`
	// The first name of the user.
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the user.
	LastName *string `json:"lastName,omitempty"`
	// The email address associated with the user.
	EmailAddress *string `json:"emailAddress,omitempty"`
	// The password for the new user.
	Password *string `json:"password,omitempty"`
	// The user's status, e.g. active or disabled.
	Status string `json:"status"`
	// The roles which the user has been assigned within Nexus.
	Roles []string `json:"roles,omitempty"`
}

type _ApiCreateUser ApiCreateUser

// NewApiCreateUser instantiates a new ApiCreateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCreateUser(status string) *ApiCreateUser {
	this := ApiCreateUser{}
	this.Status = status
	return &this
}

// NewApiCreateUserWithDefaults instantiates a new ApiCreateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCreateUserWithDefaults() *ApiCreateUser {
	this := ApiCreateUser{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiCreateUser) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiCreateUser) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiCreateUser) SetUserId(v string) {
	o.UserId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ApiCreateUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ApiCreateUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ApiCreateUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ApiCreateUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ApiCreateUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ApiCreateUser) SetLastName(v string) {
	o.LastName = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ApiCreateUser) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ApiCreateUser) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ApiCreateUser) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApiCreateUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApiCreateUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApiCreateUser) SetPassword(v string) {
	o.Password = &v
}

// GetStatus returns the Status field value
func (o *ApiCreateUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiCreateUser) SetStatus(v string) {
	o.Status = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApiCreateUser) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCreateUser) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApiCreateUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ApiCreateUser) SetRoles(v []string) {
	o.Roles = v
}

func (o ApiCreateUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCreateUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *ApiCreateUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiCreateUser := _ApiCreateUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiCreateUser)

	if err != nil {
		return err
	}

	*o = ApiCreateUser(varApiCreateUser)

	return err
}

type NullableApiCreateUser struct {
	value *ApiCreateUser
	isSet bool
}

func (v NullableApiCreateUser) Get() *ApiCreateUser {
	return v.value
}

func (v *NullableApiCreateUser) Set(val *ApiCreateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCreateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCreateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCreateUser(val *ApiCreateUser) *NullableApiCreateUser {
	return &NullableApiCreateUser{value: val, isSet: true}
}

func (v NullableApiCreateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCreateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


