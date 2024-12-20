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
// Product : Indicates the Square product used to generate a change.
type Product string

// List of Product
const (
	SQUARE_POS_Product Product = "SQUARE_POS"
	EXTERNAL_API_Product Product = "EXTERNAL_API"
	BILLING_Product Product = "BILLING"
	APPOINTMENTS_Product Product = "APPOINTMENTS"
	INVOICES_Product Product = "INVOICES"
	ONLINE_STORE_Product Product = "ONLINE_STORE"
	PAYROLL_Product Product = "PAYROLL"
	DASHBOARD_Product Product = "DASHBOARD"
	ITEM_LIBRARY_IMPORT_Product Product = "ITEM_LIBRARY_IMPORT"
	OTHER_Product Product = "OTHER"
)
