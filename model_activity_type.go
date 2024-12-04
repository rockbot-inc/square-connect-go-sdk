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

type ActivityType string

// List of ActivityType
const (
	ADJUSTMENT_ActivityType ActivityType = "ADJUSTMENT"
	APP_FEE_REFUND_ActivityType ActivityType = "APP_FEE_REFUND"
	APP_FEE_REVENUE_ActivityType ActivityType = "APP_FEE_REVENUE"
	AUTOMATIC_SAVINGS_ActivityType ActivityType = "AUTOMATIC_SAVINGS"
	AUTOMATIC_SAVINGS_REVERSED_ActivityType ActivityType = "AUTOMATIC_SAVINGS_REVERSED"
	CHARGE_ActivityType ActivityType = "CHARGE"
	DEPOSIT_FEE_ActivityType ActivityType = "DEPOSIT_FEE"
	DEPOSIT_FEE_REVERSED_ActivityType ActivityType = "DEPOSIT_FEE_REVERSED"
	DISPUTE_ActivityType ActivityType = "DISPUTE"
	ESCHEATMENT_ActivityType ActivityType = "ESCHEATMENT"
	FEE_ActivityType ActivityType = "FEE"
	FREE_PROCESSING_ActivityType ActivityType = "FREE_PROCESSING"
	HOLD_ADJUSTMENT_ActivityType ActivityType = "HOLD_ADJUSTMENT"
	INITIAL_BALANCE_CHANGE_ActivityType ActivityType = "INITIAL_BALANCE_CHANGE"
	MONEY_TRANSFER_ActivityType ActivityType = "MONEY_TRANSFER"
	MONEY_TRANSFER_REVERSAL_ActivityType ActivityType = "MONEY_TRANSFER_REVERSAL"
	OPEN_DISPUTE_ActivityType ActivityType = "OPEN_DISPUTE"
	OTHER_ActivityType ActivityType = "OTHER"
	OTHER_ADJUSTMENT_ActivityType ActivityType = "OTHER_ADJUSTMENT"
	PAID_SERVICE_FEE_ActivityType ActivityType = "PAID_SERVICE_FEE"
	PAID_SERVICE_FEE_REFUND_ActivityType ActivityType = "PAID_SERVICE_FEE_REFUND"
	REDEMPTION_CODE_ActivityType ActivityType = "REDEMPTION_CODE"
	REFUND_ActivityType ActivityType = "REFUND"
	RELEASE_ADJUSTMENT_ActivityType ActivityType = "RELEASE_ADJUSTMENT"
	RESERVE_HOLD_ActivityType ActivityType = "RESERVE_HOLD"
	RESERVE_RELEASE_ActivityType ActivityType = "RESERVE_RELEASE"
	RETURNED_PAYOUT_ActivityType ActivityType = "RETURNED_PAYOUT"
	SQUARE_CAPITAL_PAYMENT_ActivityType ActivityType = "SQUARE_CAPITAL_PAYMENT"
	SQUARE_CAPITAL_REVERSED_PAYMENT_ActivityType ActivityType = "SQUARE_CAPITAL_REVERSED_PAYMENT"
	SUBSCRIPTION_FEE_ActivityType ActivityType = "SUBSCRIPTION_FEE"
	SUBSCRIPTION_FEE_PAID_REFUND_ActivityType ActivityType = "SUBSCRIPTION_FEE_PAID_REFUND"
	SUBSCRIPTION_FEE_REFUND_ActivityType ActivityType = "SUBSCRIPTION_FEE_REFUND"
	TAX_ON_FEE_ActivityType ActivityType = "TAX_ON_FEE"
	THIRD_PARTY_FEE_ActivityType ActivityType = "THIRD_PARTY_FEE"
	THIRD_PARTY_FEE_REFUND_ActivityType ActivityType = "THIRD_PARTY_FEE_REFUND"
	PAYOUT_ActivityType ActivityType = "PAYOUT"
	AUTOMATIC_BITCOIN_CONVERSIONS_ActivityType ActivityType = "AUTOMATIC_BITCOIN_CONVERSIONS"
	AUTOMATIC_BITCOIN_CONVERSIONS_REVERSED_ActivityType ActivityType = "AUTOMATIC_BITCOIN_CONVERSIONS_REVERSED"
	CREDIT_CARD_REPAYMENT_ActivityType ActivityType = "CREDIT_CARD_REPAYMENT"
	CREDIT_CARD_REPAYMENT_REVERSED_ActivityType ActivityType = "CREDIT_CARD_REPAYMENT_REVERSED"
	LOCAL_OFFERS_CASHBACK_ActivityType ActivityType = "LOCAL_OFFERS_CASHBACK"
	LOCAL_OFFERS_FEE_ActivityType ActivityType = "LOCAL_OFFERS_FEE"
	PERCENTAGE_PROCESSING_ENROLLMENT_ActivityType ActivityType = "PERCENTAGE_PROCESSING_ENROLLMENT"
	PERCENTAGE_PROCESSING_DEACTIVATION_ActivityType ActivityType = "PERCENTAGE_PROCESSING_DEACTIVATION"
	PERCENTAGE_PROCESSING_REPAYMENT_ActivityType ActivityType = "PERCENTAGE_PROCESSING_REPAYMENT"
	PERCENTAGE_PROCESSING_REPAYMENT_REVERSED_ActivityType ActivityType = "PERCENTAGE_PROCESSING_REPAYMENT_REVERSED"
	PROCESSING_FEE_ActivityType ActivityType = "PROCESSING_FEE"
	PROCESSING_FEE_REFUND_ActivityType ActivityType = "PROCESSING_FEE_REFUND"
	UNDO_PROCESSING_FEE_REFUND_ActivityType ActivityType = "UNDO_PROCESSING_FEE_REFUND"
	GIFT_CARD_LOAD_FEE_ActivityType ActivityType = "GIFT_CARD_LOAD_FEE"
	GIFT_CARD_LOAD_FEE_REFUND_ActivityType ActivityType = "GIFT_CARD_LOAD_FEE_REFUND"
	UNDO_GIFT_CARD_LOAD_FEE_REFUND_ActivityType ActivityType = "UNDO_GIFT_CARD_LOAD_FEE_REFUND"
	BALANCE_FOLDERS_TRANSFER_ActivityType ActivityType = "BALANCE_FOLDERS_TRANSFER"
	BALANCE_FOLDERS_TRANSFER_REVERSED_ActivityType ActivityType = "BALANCE_FOLDERS_TRANSFER_REVERSED"
	GIFT_CARD_POOL_TRANSFER_ActivityType ActivityType = "GIFT_CARD_POOL_TRANSFER"
	GIFT_CARD_POOL_TRANSFER_REVERSED_ActivityType ActivityType = "GIFT_CARD_POOL_TRANSFER_REVERSED"
	SQUARE_PAYROLL_TRANSFER_ActivityType ActivityType = "SQUARE_PAYROLL_TRANSFER"
	SQUARE_PAYROLL_TRANSFER_REVERSED_ActivityType ActivityType = "SQUARE_PAYROLL_TRANSFER_REVERSED"
)
