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

// A category that can be assigned to an item or a parent category that can be assigned  to another category. For example, a clothing category can be assigned to a t-shirt item or  be made as the parent category to the pants category.
type CatalogObjectCategory struct {
	// The ID of the object's category.
	Id string `json:"id,omitempty"`
	// The order of the object within the context of the category.
	Ordinal int64 `json:"ordinal,omitempty"`
}
