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

// The parameters of a `Shift` search query, which includes filter and sort options.
type ShiftQuery struct {
	Filter *ShiftFilter `json:"filter,omitempty"`
	Sort *ShiftSort `json:"sort,omitempty"`
}
