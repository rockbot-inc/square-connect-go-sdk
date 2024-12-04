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

// Identifies the sort field and sort order.
type InvoiceSort struct {
	Field *InvoiceSortField `json:"field"`
	Order *SortOrder `json:"order,omitempty"`
}
