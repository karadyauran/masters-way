/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasToUserMentoringRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasToUserMentoringRequestResponse{}

// SchemasToUserMentoringRequestResponse struct for SchemasToUserMentoringRequestResponse
type SchemasToUserMentoringRequestResponse struct {
	UserId string `json:"userId"`
	WayId string `json:"wayId"`
}

type _SchemasToUserMentoringRequestResponse SchemasToUserMentoringRequestResponse

// NewSchemasToUserMentoringRequestResponse instantiates a new SchemasToUserMentoringRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasToUserMentoringRequestResponse(userId string, wayId string) *SchemasToUserMentoringRequestResponse {
	this := SchemasToUserMentoringRequestResponse{}
	this.UserId = userId
	this.WayId = wayId
	return &this
}

// NewSchemasToUserMentoringRequestResponseWithDefaults instantiates a new SchemasToUserMentoringRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasToUserMentoringRequestResponseWithDefaults() *SchemasToUserMentoringRequestResponse {
	this := SchemasToUserMentoringRequestResponse{}
	return &this
}

// GetUserId returns the UserId field value
func (o *SchemasToUserMentoringRequestResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SchemasToUserMentoringRequestResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SchemasToUserMentoringRequestResponse) SetUserId(v string) {
	o.UserId = v
}

// GetWayId returns the WayId field value
func (o *SchemasToUserMentoringRequestResponse) GetWayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayId
}

// GetWayIdOk returns a tuple with the WayId field value
// and a boolean to check if the value has been set.
func (o *SchemasToUserMentoringRequestResponse) GetWayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayId, true
}

// SetWayId sets field value
func (o *SchemasToUserMentoringRequestResponse) SetWayId(v string) {
	o.WayId = v
}

func (o SchemasToUserMentoringRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasToUserMentoringRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["wayId"] = o.WayId
	return toSerialize, nil
}

func (o *SchemasToUserMentoringRequestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"wayId",
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

	varSchemasToUserMentoringRequestResponse := _SchemasToUserMentoringRequestResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasToUserMentoringRequestResponse)

	if err != nil {
		return err
	}

	*o = SchemasToUserMentoringRequestResponse(varSchemasToUserMentoringRequestResponse)

	return err
}

type NullableSchemasToUserMentoringRequestResponse struct {
	value *SchemasToUserMentoringRequestResponse
	isSet bool
}

func (v NullableSchemasToUserMentoringRequestResponse) Get() *SchemasToUserMentoringRequestResponse {
	return v.value
}

func (v *NullableSchemasToUserMentoringRequestResponse) Set(val *SchemasToUserMentoringRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasToUserMentoringRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasToUserMentoringRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasToUserMentoringRequestResponse(val *SchemasToUserMentoringRequestResponse) *NullableSchemasToUserMentoringRequestResponse {
	return &NullableSchemasToUserMentoringRequestResponse{value: val, isSet: true}
}

func (v NullableSchemasToUserMentoringRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasToUserMentoringRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


