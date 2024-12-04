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

// The query filter to return the items containing the specified tax IDs.
type CatalogQueryItemsForTax struct {
	// A set of `CatalogTax` IDs to be used to find associated `CatalogItem`s.
	TaxIds []string `json:"tax_ids"`
}
