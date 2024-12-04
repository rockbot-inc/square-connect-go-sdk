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
// BankAccountStatus : Indicates the current verification status of a `BankAccount` object.
type BankAccountStatus string

// List of BankAccountStatus
const (
	VERIFICATION_IN_PROGRESS_BankAccountStatus BankAccountStatus = "VERIFICATION_IN_PROGRESS"
	VERIFIED_BankAccountStatus BankAccountStatus = "VERIFIED"
	DISABLED_BankAccountStatus BankAccountStatus = "DISABLED"
)