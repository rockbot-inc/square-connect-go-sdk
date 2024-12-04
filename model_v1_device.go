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

type V1Device struct {
	// The device's Square-issued ID.
	Id string `json:"id,omitempty"`
	// The device's merchant-specified name.
	Name string `json:"name,omitempty"`
}
