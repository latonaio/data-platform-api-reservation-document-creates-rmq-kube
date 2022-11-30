package requests

type ReservationDocument struct {
	Reservation                     *int        `json:"Reservation"`
	OrderID                         *int        `json:"OrderID"`
	GoodsMovementType               *string     `json:"GoodsMovementType"`
	CostCenter                      *string     `json:"CostCenter"`
	GoodsRecipientName              *string     `json:"GoodsRecipientName"`
	ReservationDate                 *string     `json:"ReservationDate"`
	Customer                        *bool       `json:"Customer"`
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
