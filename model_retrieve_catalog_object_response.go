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

type RetrieveCatalogObjectResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	Object *CatalogObject `json:"object,omitempty"`
	// A list of `CatalogObject`s referenced by the object in the `object` field.
	RelatedObjects []CatalogObject `json:"related_objects,omitempty"`
}
