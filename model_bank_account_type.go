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
// BankAccountType : Indicates the financial purpose of the bank account.
type BankAccountType string

// List of BankAccountType
const (
	CHECKING_BankAccountType BankAccountType = "CHECKING"
	SAVINGS_BankAccountType BankAccountType = "SAVINGS"
	INVESTMENT_BankAccountType BankAccountType = "INVESTMENT"
	OTHER_BankAccountType BankAccountType = "OTHER"
	BUSINESS_CHECKING_BankAccountType BankAccountType = "BUSINESS_CHECKING"
)
