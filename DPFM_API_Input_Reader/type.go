package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string          `json:"connection_key"`
	Result              bool            `json:"result"`
	RedisKey            string          `json:"redis_key"`
	Filepath            string          `json:"filepath"`
	APIStatusCode       int             `json:"api_status_code"`
	RuntimeSessionID    string          `json:"runtime_session_id"`
	BusinessPartnerID   *int            `json:"business_partner"`
	ServiceLabel        string          `json:"service_label"`
	ReservationDocument     ReservationDocument `json:"ReservationDocument"`
	APISchema           string          `json:"api_schema"`
	Accepter            []string        `json:"accepter"`
	OrderID             *int            `json:"order_id"`
	Deleted             bool            `json:"deleted"`
	SQLUpdateResult     *bool           `json:"sql_update_result"`
	SQLUpdateError      string          `json:"sql_update_error"`
	SubfuncResult       *bool           `json:"subfunc_result"`
	SubfuncError        string          `json:"subfunc_error"`
	ExconfResult        *bool           `json:"exconf_result"`
	ExconfError         string          `json:"exconf_error"`
	APIProcessingResult *bool           `json:"api_processing_result"`
	APIProcessingError  string          `json:"api_processing_error"`
}

type ReservationDocument struct {
	Reservation                     *int        `json:"Reservation"`
	OrderID                         *int        `json:"OrderID"`
	GoodsMovementType               *string     `json:"GoodsMovementType"`
	CostCenter                      *string     `json:"CostCenter"`
	GoodsRecipientName              *string     `json:"GoodsRecipientName"`
	ReservationDate                 *string     `json:"ReservationDate"`
	Customer                        *bool        `json:"Customer"`
	WBSElement                      *string     `json:"WBSElement"`
	ControllingArea                 *string     `json:"ControllingArea"`
	SalesOrder                      *string     `json:"SalesOrder"`
	SalesOrderItem                  *string     `json:"SalesOrderItem"`
	SalesOrderScheduleLine          *string     `json:"SalesOrderScheduleLine"`
	AssetNumber                     *string     `json:"AssetNumber"`
	AssetSubNumber                  *string     `json:"AssetSubNumber"`
	NetworkNumberForAcctAssgmt      *string     `json:"NetworkNumberForAcctAssgmt"`
	IssuingOrReceivingPlant         *string     `json:"IssuingOrReceivingPlant"`
	IssuingOrReceivingStorageLoc    *string     `json:"IssuingOrReceivingStorageLoc"`
	ReservationDocumentItem ReservationDocumentItem `json:"ReservationDocumentItem"`
}

type ReservationDocumentItem struct {
	Reservation                       *int    `json:"Reservation"`
	ReservationItem                   *int    `json:"ReservationItem"`
	RecordType                        *string `json:"RecordType"`
	Product                           *string `json:"Product"`
	RequirementType                   *string `json:"RequirementType"`
	MatlCompRequirementDate           *string `json:"MatlCompRequirementDate"`
	Plant                             *string `json:"Plant"`
	ManufacturingOrderOperation       *string `json:"ManufacturingOrderOperation"`
	GoodsMovementIsAllowed            *string `json:"GoodsMovementIsAllowed"`
	StorageLocation                   *string `json:"StorageLocation"`
	Batch                             *string `json:"Batch"`
	DebitCreditCode                   *string `json:"DebitCreditCode"`
	BaseUnit                          *string `json:"BaseUnit"`
	GLAccount                         *string `json:"GLAccount"`
	GoodsMovementType                 *string `json:"GoodsMovementType"`
	EntryUnit                         *string `json:"EntryUnit"`
	QuantityIsFixed                   *string `json:"QuantityIsFixed"`
	CompanyCodeCurrency               *string `json:"CompanyCodeCurrency"`
	IssuingOrReceivingPlant           *string `json:"IssuingOrReceivingPlant"`
	IssuingOrReceivingStorageLoc      *string `json:"IssuingOrReceivingStorageLoc"`
	PurchasingDocument                *string `json:"PurchasingDocument"`
	PurchasingDocumentItem            *string `json:"PurchasingDocumentItem"`
	Supplier                          *string `json:"Supplier"`
	ResvnItmRequiredQtyInBaseUnit     *string `json:"ResvnItmRequiredQtyInBaseUnit"`
	ReservationItemIsFinallyIssued    *string `json:"ReservationItemIsFinallyIssued"`
	ReservationItmIsMarkedForDeltn    *bool   `json:"ReservationItmIsMarkedForDeltn"`
	ResvnItmRequiredQtyInEntryUnit    *string `json:"ResvnItmRequiredQtyInEntryUnit"`
	ResvnItmWithdrawnQtyInBaseUnit    *string `json:"ResvnItmWithdrawnQtyInBaseUnit"`
	ResvnItmWithdrawnAmtInCCCrcy      *string `json:"ResvnItmWithdrawnAmtInCCCrcy"`
	GoodsRecipientName                *string `json:"GoodsRecipientName"`
	UnloadingPointName                *string `json:"UnloadingPointName"`
	ReservationItemText               *string `json:"ReservationItemText"`
}
