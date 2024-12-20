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

type DeviceMetadata struct {
	// The Terminal’s remaining battery percentage, between 1-100.
	BatteryPercentage string `json:"battery_percentage,omitempty"`
	// The current charging state of the Terminal. Options: `CHARGING`, `NOT_CHARGING`
	ChargingState string `json:"charging_state,omitempty"`
	// The ID of the Square seller business location associated with the Terminal.
	LocationId string `json:"location_id,omitempty"`
	// The ID of the Square merchant account that is currently signed-in to the Terminal.
	MerchantId string `json:"merchant_id,omitempty"`
	// The Terminal’s current network connection type. Options: `WIFI`, `ETHERNET`
	NetworkConnectionType string `json:"network_connection_type,omitempty"`
	// The country in which the Terminal is authorized to take payments.
	PaymentRegion string `json:"payment_region,omitempty"`
	// The unique identifier assigned to the Terminal, which can be found on the lower back of the device.
	SerialNumber string `json:"serial_number,omitempty"`
	// The current version of the Terminal’s operating system.
	OsVersion string `json:"os_version,omitempty"`
	// The current version of the application running on the Terminal.
	AppVersion string `json:"app_version,omitempty"`
	// The name of the Wi-Fi network to which the Terminal is connected.
	WifiNetworkName string `json:"wifi_network_name,omitempty"`
	// The signal strength of the Wi-FI network connection. Options: `POOR`, `FAIR`, `GOOD`, `EXCELLENT`
	WifiNetworkStrength string `json:"wifi_network_strength,omitempty"`
	// The IP address of the Terminal.
	IpAddress string `json:"ip_address,omitempty"`
}
