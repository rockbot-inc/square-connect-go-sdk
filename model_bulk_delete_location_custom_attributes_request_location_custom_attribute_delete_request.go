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

// Represents an individual delete request in a [BulkDeleteLocationCustomAttributes]($e/LocationCustomAttributes/BulkDeleteLocationCustomAttributes) request. An individual request contains an optional ID of the associated custom attribute definition and optional key of the associated custom attribute definition.
type BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest struct {
	// The key of the associated custom attribute definition. Represented as a qualified key if the requesting app is not the definition owner.
	Key string `json:"key,omitempty"`
}