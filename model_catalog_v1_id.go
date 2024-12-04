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

// A Square API V1 identifier of an item, including the object ID and its associated location ID.
type CatalogV1Id struct {
	// The ID for an object used in the Square API V1, if the object ID differs from the Square API V2 object ID.
	CatalogV1Id string `json:"catalog_v1_id,omitempty"`
	// The ID of the `Location` this Connect V1 ID is associated with.
	LocationId string `json:"location_id,omitempty"`
}
