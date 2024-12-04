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

// Used *internally* to encapsulate pagination details. The resulting proto will be base62 encoded in order to produce a cursor that can be used externally.
type PaginationCursor struct {
	// The ID of the last resource in the current page. The page can be in an ascending or descending order
	OrderValue string `json:"order_value,omitempty"`
}
