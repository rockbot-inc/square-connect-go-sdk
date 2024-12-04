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
// TransactionType : The transaction type used in the disputed payment.
type TransactionType string

// List of TransactionType
const (
	DEBIT_TransactionType TransactionType = "DEBIT"
	CREDIT_TransactionType TransactionType = "CREDIT"
)
