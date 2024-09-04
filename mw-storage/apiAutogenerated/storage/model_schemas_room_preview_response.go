/*
Masters way storage API

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

// checks if the SchemasRoomPreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasRoomPreviewResponse{}

// SchemasRoomPreviewResponse struct for SchemasRoomPreviewResponse
type SchemasRoomPreviewResponse struct {
	IsBlocked bool `json:"isBlocked"`
	Name NullableString `json:"name"`
	RoomId string `json:"roomId"`
	RoomType string `json:"roomType"`
	Users []SchemasUserResponse `json:"users"`
}

type _SchemasRoomPreviewResponse SchemasRoomPreviewResponse

// NewSchemasRoomPreviewResponse instantiates a new SchemasRoomPreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasRoomPreviewResponse(isBlocked bool, name NullableString, roomId string, roomType string, users []SchemasUserResponse) *SchemasRoomPreviewResponse {
	this := SchemasRoomPreviewResponse{}
	this.IsBlocked = isBlocked
	this.Name = name
	this.RoomId = roomId
	this.RoomType = roomType
	this.Users = users
	return &this
}

// NewSchemasRoomPreviewResponseWithDefaults instantiates a new SchemasRoomPreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasRoomPreviewResponseWithDefaults() *SchemasRoomPreviewResponse {
	this := SchemasRoomPreviewResponse{}
	return &this
}

// GetIsBlocked returns the IsBlocked field value
func (o *SchemasRoomPreviewResponse) GetIsBlocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPreviewResponse) GetIsBlockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBlocked, true
}

// SetIsBlocked sets field value
func (o *SchemasRoomPreviewResponse) SetIsBlocked(v bool) {
	o.IsBlocked = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SchemasRoomPreviewResponse) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasRoomPreviewResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *SchemasRoomPreviewResponse) SetName(v string) {
	o.Name.Set(&v)
}

// GetRoomId returns the RoomId field value
func (o *SchemasRoomPreviewResponse) GetRoomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPreviewResponse) GetRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *SchemasRoomPreviewResponse) SetRoomId(v string) {
	o.RoomId = v
}

// GetRoomType returns the RoomType field value
func (o *SchemasRoomPreviewResponse) GetRoomType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPreviewResponse) GetRoomTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *SchemasRoomPreviewResponse) SetRoomType(v string) {
	o.RoomType = v
}

// GetUsers returns the Users field value
func (o *SchemasRoomPreviewResponse) GetUsers() []SchemasUserResponse {
	if o == nil {
		var ret []SchemasUserResponse
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SchemasRoomPreviewResponse) GetUsersOk() ([]SchemasUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SchemasRoomPreviewResponse) SetUsers(v []SchemasUserResponse) {
	o.Users = v
}

func (o SchemasRoomPreviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasRoomPreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isBlocked"] = o.IsBlocked
	toSerialize["name"] = o.Name.Get()
	toSerialize["roomId"] = o.RoomId
	toSerialize["roomType"] = o.RoomType
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *SchemasRoomPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isBlocked",
		"name",
		"roomId",
		"roomType",
		"users",
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

	varSchemasRoomPreviewResponse := _SchemasRoomPreviewResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasRoomPreviewResponse)

	if err != nil {
		return err
	}

	*o = SchemasRoomPreviewResponse(varSchemasRoomPreviewResponse)

	return err
}

type NullableSchemasRoomPreviewResponse struct {
	value *SchemasRoomPreviewResponse
	isSet bool
}

func (v NullableSchemasRoomPreviewResponse) Get() *SchemasRoomPreviewResponse {
	return v.value
}

func (v *NullableSchemasRoomPreviewResponse) Set(val *SchemasRoomPreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasRoomPreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasRoomPreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasRoomPreviewResponse(val *SchemasRoomPreviewResponse) *NullableSchemasRoomPreviewResponse {
	return &NullableSchemasRoomPreviewResponse{value: val, isSet: true}
}

func (v NullableSchemasRoomPreviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasRoomPreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


