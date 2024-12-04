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

// SEO data for for a seller's Square Online store.
type CatalogEcomSeoData struct {
	// The SEO title used for the Square Online store.
	PageTitle string `json:"page_title,omitempty"`
	// The SEO description used for the Square Online store.
	PageDescription string `json:"page_description,omitempty"`
	// The SEO permalink used for the Square Online store.
	Permalink string `json:"permalink,omitempty"`
}
