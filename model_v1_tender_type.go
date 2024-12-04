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

type V1TenderType string

// List of V1TenderType
const (
	CREDIT_CARD_V1TenderType V1TenderType = "CREDIT_CARD"
	CASH_V1TenderType V1TenderType = "CASH"
	THIRD_PARTY_CARD_V1TenderType V1TenderType = "THIRD_PARTY_CARD"
	NO_SALE_V1TenderType V1TenderType = "NO_SALE"
	SQUARE_WALLET_V1TenderType V1TenderType = "SQUARE_WALLET"
	SQUARE_GIFT_CARD_V1TenderType V1TenderType = "SQUARE_GIFT_CARD"
	UNKNOWN_V1TenderType V1TenderType = "UNKNOWN"
	OTHER_V1TenderType V1TenderType = "OTHER"
)
