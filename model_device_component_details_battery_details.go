/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DeviceComponentDetailsBatteryDetails struct {
	// The battery charge percentage as displayed on the device.
	VisiblePercent int32 `json:"visible_percent,omitempty"`
	ExternalPower *DeviceComponentDetailsExternalPower `json:"external_power,omitempty"`
}
