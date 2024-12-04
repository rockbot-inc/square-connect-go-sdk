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

// Represents one delete within the bulk operation.
type BulkDeleteOrderCustomAttributesRequestDeleteCustomAttribute struct {
	// The key of the custom attribute to delete.  This key must match the key  of an existing custom attribute definition.
	Key string `json:"key,omitempty"`
	// The ID of the target [order](entity:Order).
	OrderId string `json:"order_id"`
}
