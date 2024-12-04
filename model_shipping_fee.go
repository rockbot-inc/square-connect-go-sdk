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

type ShippingFee struct {
	// The name for the shipping fee.
	Name string `json:"name,omitempty"`
	Charge *Money `json:"charge"`
}
