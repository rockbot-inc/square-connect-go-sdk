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
// VisibilityFilter : Enumeration of visibility-filter values used to set the ability to view custom attributes or custom attribute definitions.
type VisibilityFilter string

// List of VisibilityFilter
const (
	ALL_VisibilityFilter VisibilityFilter = "ALL"
	READ_VisibilityFilter VisibilityFilter = "READ"
	READ_WRITE_VisibilityFilter VisibilityFilter = "READ_WRITE"
)
