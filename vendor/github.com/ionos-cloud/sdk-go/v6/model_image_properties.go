/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ImageProperties struct for ImageProperties
type ImageProperties struct {
	// Cloud init compatibility.
	CloudInit *string `json:"cloudInit,omitempty"`
	// Hot-plug capable CPU (no reboot required).
	CpuHotPlug *bool `json:"cpuHotPlug,omitempty"`
	// Hot-unplug capable CPU (no reboot required).
	CpuHotUnplug *bool `json:"cpuHotUnplug,omitempty"`
	// Human-readable description.
	Description *string `json:"description,omitempty"`
	// Hot-plug capable SCSI drive (no reboot required).
	DiscScsiHotPlug *bool `json:"discScsiHotPlug,omitempty"`
	// Hot-unplug capable SCSI drive (no reboot required). Not supported with Windows VMs.
	DiscScsiHotUnplug *bool `json:"discScsiHotUnplug,omitempty"`
	// Hot-plug capable Virt-IO drive (no reboot required).
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty"`
	// Hot-unplug capable Virt-IO drive (no reboot required). Not supported with Windows VMs.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty"`
	// List of image aliases mapped for this image
	ImageAliases *[]string `json:"imageAliases,omitempty"`
	// The image type.
	ImageType *string `json:"imageType,omitempty"`
	// The OS type of this image.
	LicenceType *string `json:"licenceType"`
	// The location of this image/snapshot.
	Location *string `json:"location,omitempty"`
	// The resource name.
	Name *string `json:"name,omitempty"`
	// Hot-plug capable NIC (no reboot required).
	NicHotPlug *bool `json:"nicHotPlug,omitempty"`
	// Hot-unplug capable NIC (no reboot required).
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty"`
	// Indicates whether the image is part of a public repository.
	Public *bool `json:"public,omitempty"`
	// Hot-plug capable RAM (no reboot required).
	RamHotPlug *bool `json:"ramHotPlug,omitempty"`
	// Hot-unplug capable RAM (no reboot required).
	RamHotUnplug *bool `json:"ramHotUnplug,omitempty"`
	// The image size in GB.
	Size *float32 `json:"size,omitempty"`
}

// NewImageProperties instantiates a new ImageProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageProperties(licenceType string) *ImageProperties {
	this := ImageProperties{}

	this.LicenceType = &licenceType

	return &this
}

// NewImagePropertiesWithDefaults instantiates a new ImageProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagePropertiesWithDefaults() *ImageProperties {
	this := ImageProperties{}
	return &this
}

// GetCloudInit returns the CloudInit field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetCloudInit() *string {
	if o == nil {
		return nil
	}

	return o.CloudInit

}

// GetCloudInitOk returns a tuple with the CloudInit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetCloudInitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CloudInit, true
}

// SetCloudInit sets field value
func (o *ImageProperties) SetCloudInit(v string) {

	o.CloudInit = &v

}

// HasCloudInit returns a boolean if a field has been set.
func (o *ImageProperties) HasCloudInit() bool {
	if o != nil && o.CloudInit != nil {
		return true
	}

	return false
}

// GetCpuHotPlug returns the CpuHotPlug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetCpuHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.CpuHotPlug

}

// GetCpuHotPlugOk returns a tuple with the CpuHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetCpuHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.CpuHotPlug, true
}

// SetCpuHotPlug sets field value
func (o *ImageProperties) SetCpuHotPlug(v bool) {

	o.CpuHotPlug = &v

}

// HasCpuHotPlug returns a boolean if a field has been set.
func (o *ImageProperties) HasCpuHotPlug() bool {
	if o != nil && o.CpuHotPlug != nil {
		return true
	}

	return false
}

// GetCpuHotUnplug returns the CpuHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetCpuHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.CpuHotUnplug

}

// GetCpuHotUnplugOk returns a tuple with the CpuHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetCpuHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.CpuHotUnplug, true
}

// SetCpuHotUnplug sets field value
func (o *ImageProperties) SetCpuHotUnplug(v bool) {

	o.CpuHotUnplug = &v

}

// HasCpuHotUnplug returns a boolean if a field has been set.
func (o *ImageProperties) HasCpuHotUnplug() bool {
	if o != nil && o.CpuHotUnplug != nil {
		return true
	}

	return false
}

// GetDescription returns the Description field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetDescription() *string {
	if o == nil {
		return nil
	}

	return o.Description

}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Description, true
}

// SetDescription sets field value
func (o *ImageProperties) SetDescription(v string) {

	o.Description = &v

}

// HasDescription returns a boolean if a field has been set.
func (o *ImageProperties) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// GetDiscScsiHotPlug returns the DiscScsiHotPlug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetDiscScsiHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscScsiHotPlug

}

// GetDiscScsiHotPlugOk returns a tuple with the DiscScsiHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetDiscScsiHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscScsiHotPlug, true
}

// SetDiscScsiHotPlug sets field value
func (o *ImageProperties) SetDiscScsiHotPlug(v bool) {

	o.DiscScsiHotPlug = &v

}

// HasDiscScsiHotPlug returns a boolean if a field has been set.
func (o *ImageProperties) HasDiscScsiHotPlug() bool {
	if o != nil && o.DiscScsiHotPlug != nil {
		return true
	}

	return false
}

// GetDiscScsiHotUnplug returns the DiscScsiHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetDiscScsiHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscScsiHotUnplug

}

// GetDiscScsiHotUnplugOk returns a tuple with the DiscScsiHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetDiscScsiHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscScsiHotUnplug, true
}

// SetDiscScsiHotUnplug sets field value
func (o *ImageProperties) SetDiscScsiHotUnplug(v bool) {

	o.DiscScsiHotUnplug = &v

}

// HasDiscScsiHotUnplug returns a boolean if a field has been set.
func (o *ImageProperties) HasDiscScsiHotUnplug() bool {
	if o != nil && o.DiscScsiHotUnplug != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotPlug returns the DiscVirtioHotPlug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetDiscVirtioHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotPlug

}

// GetDiscVirtioHotPlugOk returns a tuple with the DiscVirtioHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetDiscVirtioHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscVirtioHotPlug, true
}

// SetDiscVirtioHotPlug sets field value
func (o *ImageProperties) SetDiscVirtioHotPlug(v bool) {

	o.DiscVirtioHotPlug = &v

}

// HasDiscVirtioHotPlug returns a boolean if a field has been set.
func (o *ImageProperties) HasDiscVirtioHotPlug() bool {
	if o != nil && o.DiscVirtioHotPlug != nil {
		return true
	}

	return false
}

// GetDiscVirtioHotUnplug returns the DiscVirtioHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetDiscVirtioHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.DiscVirtioHotUnplug

}

// GetDiscVirtioHotUnplugOk returns a tuple with the DiscVirtioHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetDiscVirtioHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.DiscVirtioHotUnplug, true
}

// SetDiscVirtioHotUnplug sets field value
func (o *ImageProperties) SetDiscVirtioHotUnplug(v bool) {

	o.DiscVirtioHotUnplug = &v

}

// HasDiscVirtioHotUnplug returns a boolean if a field has been set.
func (o *ImageProperties) HasDiscVirtioHotUnplug() bool {
	if o != nil && o.DiscVirtioHotUnplug != nil {
		return true
	}

	return false
}

// GetImageAliases returns the ImageAliases field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetImageAliases() *[]string {
	if o == nil {
		return nil
	}

	return o.ImageAliases

}

// GetImageAliasesOk returns a tuple with the ImageAliases field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetImageAliasesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ImageAliases, true
}

// SetImageAliases sets field value
func (o *ImageProperties) SetImageAliases(v []string) {

	o.ImageAliases = &v

}

// HasImageAliases returns a boolean if a field has been set.
func (o *ImageProperties) HasImageAliases() bool {
	if o != nil && o.ImageAliases != nil {
		return true
	}

	return false
}

// GetImageType returns the ImageType field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetImageType() *string {
	if o == nil {
		return nil
	}

	return o.ImageType

}

// GetImageTypeOk returns a tuple with the ImageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetImageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.ImageType, true
}

// SetImageType sets field value
func (o *ImageProperties) SetImageType(v string) {

	o.ImageType = &v

}

// HasImageType returns a boolean if a field has been set.
func (o *ImageProperties) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// GetLicenceType returns the LicenceType field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetLicenceType() *string {
	if o == nil {
		return nil
	}

	return o.LicenceType

}

// GetLicenceTypeOk returns a tuple with the LicenceType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetLicenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LicenceType, true
}

// SetLicenceType sets field value
func (o *ImageProperties) SetLicenceType(v string) {

	o.LicenceType = &v

}

// HasLicenceType returns a boolean if a field has been set.
func (o *ImageProperties) HasLicenceType() bool {
	if o != nil && o.LicenceType != nil {
		return true
	}

	return false
}

// GetLocation returns the Location field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Location, true
}

// SetLocation sets field value
func (o *ImageProperties) SetLocation(v string) {

	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *ImageProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *ImageProperties) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *ImageProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetNicHotPlug returns the NicHotPlug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetNicHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotPlug

}

// GetNicHotPlugOk returns a tuple with the NicHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetNicHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.NicHotPlug, true
}

// SetNicHotPlug sets field value
func (o *ImageProperties) SetNicHotPlug(v bool) {

	o.NicHotPlug = &v

}

// HasNicHotPlug returns a boolean if a field has been set.
func (o *ImageProperties) HasNicHotPlug() bool {
	if o != nil && o.NicHotPlug != nil {
		return true
	}

	return false
}

// GetNicHotUnplug returns the NicHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetNicHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.NicHotUnplug

}

// GetNicHotUnplugOk returns a tuple with the NicHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetNicHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.NicHotUnplug, true
}

// SetNicHotUnplug sets field value
func (o *ImageProperties) SetNicHotUnplug(v bool) {

	o.NicHotUnplug = &v

}

// HasNicHotUnplug returns a boolean if a field has been set.
func (o *ImageProperties) HasNicHotUnplug() bool {
	if o != nil && o.NicHotUnplug != nil {
		return true
	}

	return false
}

// GetPublic returns the Public field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetPublic() *bool {
	if o == nil {
		return nil
	}

	return o.Public

}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.Public, true
}

// SetPublic sets field value
func (o *ImageProperties) SetPublic(v bool) {

	o.Public = &v

}

// HasPublic returns a boolean if a field has been set.
func (o *ImageProperties) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// GetRamHotPlug returns the RamHotPlug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetRamHotPlug() *bool {
	if o == nil {
		return nil
	}

	return o.RamHotPlug

}

// GetRamHotPlugOk returns a tuple with the RamHotPlug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetRamHotPlugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.RamHotPlug, true
}

// SetRamHotPlug sets field value
func (o *ImageProperties) SetRamHotPlug(v bool) {

	o.RamHotPlug = &v

}

// HasRamHotPlug returns a boolean if a field has been set.
func (o *ImageProperties) HasRamHotPlug() bool {
	if o != nil && o.RamHotPlug != nil {
		return true
	}

	return false
}

// GetRamHotUnplug returns the RamHotUnplug field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetRamHotUnplug() *bool {
	if o == nil {
		return nil
	}

	return o.RamHotUnplug

}

// GetRamHotUnplugOk returns a tuple with the RamHotUnplug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetRamHotUnplugOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}

	return o.RamHotUnplug, true
}

// SetRamHotUnplug sets field value
func (o *ImageProperties) SetRamHotUnplug(v bool) {

	o.RamHotUnplug = &v

}

// HasRamHotUnplug returns a boolean if a field has been set.
func (o *ImageProperties) HasRamHotUnplug() bool {
	if o != nil && o.RamHotUnplug != nil {
		return true
	}

	return false
}

// GetSize returns the Size field value
// If the value is explicit nil, nil is returned
func (o *ImageProperties) GetSize() *float32 {
	if o == nil {
		return nil
	}

	return o.Size

}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageProperties) GetSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Size, true
}

// SetSize sets field value
func (o *ImageProperties) SetSize(v float32) {

	o.Size = &v

}

// HasSize returns a boolean if a field has been set.
func (o *ImageProperties) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

func (o ImageProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudInit != nil {
		toSerialize["cloudInit"] = o.CloudInit
	}

	if o.CpuHotPlug != nil {
		toSerialize["cpuHotPlug"] = o.CpuHotPlug
	}

	if o.CpuHotUnplug != nil {
		toSerialize["cpuHotUnplug"] = o.CpuHotUnplug
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.DiscScsiHotPlug != nil {
		toSerialize["discScsiHotPlug"] = o.DiscScsiHotPlug
	}

	if o.DiscScsiHotUnplug != nil {
		toSerialize["discScsiHotUnplug"] = o.DiscScsiHotUnplug
	}

	if o.DiscVirtioHotPlug != nil {
		toSerialize["discVirtioHotPlug"] = o.DiscVirtioHotPlug
	}

	if o.DiscVirtioHotUnplug != nil {
		toSerialize["discVirtioHotUnplug"] = o.DiscVirtioHotUnplug
	}

	if o.ImageAliases != nil {
		toSerialize["imageAliases"] = o.ImageAliases
	}

	if o.ImageType != nil {
		toSerialize["imageType"] = o.ImageType
	}

	if o.LicenceType != nil {
		toSerialize["licenceType"] = o.LicenceType
	}

	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.NicHotPlug != nil {
		toSerialize["nicHotPlug"] = o.NicHotPlug
	}

	if o.NicHotUnplug != nil {
		toSerialize["nicHotUnplug"] = o.NicHotUnplug
	}

	if o.Public != nil {
		toSerialize["public"] = o.Public
	}

	if o.RamHotPlug != nil {
		toSerialize["ramHotPlug"] = o.RamHotPlug
	}

	if o.RamHotUnplug != nil {
		toSerialize["ramHotUnplug"] = o.RamHotUnplug
	}

	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	return json.Marshal(toSerialize)
}

type NullableImageProperties struct {
	value *ImageProperties
	isSet bool
}

func (v NullableImageProperties) Get() *ImageProperties {
	return v.value
}

func (v *NullableImageProperties) Set(val *ImageProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableImageProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableImageProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageProperties(val *ImageProperties) *NullableImageProperties {
	return &NullableImageProperties{value: val, isSet: true}
}

func (v NullableImageProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
