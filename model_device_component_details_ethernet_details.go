/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

type DeviceComponentDetailsEthernetDetails struct {
	// A boolean to represent whether the Ethernet interface is currently active.
	Active bool `json:"active,omitempty"`
	// The string representation of the device’s IPv4 address.
	IpAddressV4 string `json:"ip_address_v4,omitempty"`
}
