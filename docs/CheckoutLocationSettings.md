# CheckoutLocationSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The ID of the location that these settings apply to. | [optional] [default to null]
**CustomerNotesEnabled** | **bool** | Indicates whether customers are allowed to leave notes at checkout. | [optional] [default to null]
**Policies** | [**[]CheckoutLocationSettingsPolicy**](CheckoutLocationSettingsPolicy.md) | Policy information is displayed at the bottom of the checkout pages. You can set a maximum of two policies. | [optional] [default to null]
**Branding** | [***CheckoutLocationSettingsBranding**](CheckoutLocationSettingsBranding.md) |  | [optional] [default to null]
**Tipping** | [***CheckoutLocationSettingsTipping**](CheckoutLocationSettingsTipping.md) |  | [optional] [default to null]
**Coupons** | [***CheckoutLocationSettingsCoupons**](CheckoutLocationSettingsCoupons.md) |  | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the settings were last updated, in RFC 3339 format. Examples for January 25th, 2020 6:25:34pm Pacific Standard Time: UTC: 2020-01-26T02:25:34Z Pacific Standard Time with UTC offset: 2020-01-25T18:25:34-08:00 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

