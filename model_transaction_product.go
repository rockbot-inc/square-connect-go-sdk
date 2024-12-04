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
// TransactionProduct : Indicates the Square product used to process a transaction.
type TransactionProduct string

// List of TransactionProduct
const (
	REGISTER_TransactionProduct TransactionProduct = "REGISTER"
	EXTERNAL_API_TransactionProduct TransactionProduct = "EXTERNAL_API"
	BILLING_TransactionProduct TransactionProduct = "BILLING"
	APPOINTMENTS_TransactionProduct TransactionProduct = "APPOINTMENTS"
	INVOICES_TransactionProduct TransactionProduct = "INVOICES"
	ONLINE_STORE_TransactionProduct TransactionProduct = "ONLINE_STORE"
	PAYROLL_TransactionProduct TransactionProduct = "PAYROLL"
	OTHER_TransactionProduct TransactionProduct = "OTHER"
)
