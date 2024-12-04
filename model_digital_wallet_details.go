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

// Additional details about `WALLET` type payments. Contains only non-confidential information.
type DigitalWalletDetails struct {
	// The status of the `WALLET` payment. The status can be `AUTHORIZED`, `CAPTURED`, `VOIDED`, or `FAILED`.
	Status string `json:"status,omitempty"`
	// The brand used for the `WALLET` payment. The brand can be `CASH_APP`, `PAYPAY`, `ALIPAY`, `RAKUTEN_PAY`, `AU_PAY`, `D_BARAI`, `MERPAY`, `WECHAT_PAY` or `UNKNOWN`.
	Brand string `json:"brand,omitempty"`
	CashAppDetails *CashAppDetails `json:"cash_app_details,omitempty"`
}
