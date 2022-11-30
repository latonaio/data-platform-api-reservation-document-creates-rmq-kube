package dpfm_api_input_reader

import (
	"data-platform-api-reservation-document-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToReservationDocument() *requests.ReservationDocument {
	data := sdc.ReservationDocument
	return &requests.ReservationDocument{
		Reservation:                  data.Reservation,
		OrderID:                      data.OrderID,
		GoodsMovementType:            data.GoodsMovementType,
		CostCenter:                   data.CostCenter,
		GoodsRecipientName:           data.GoodsRecipientName,
		ReservationDate:              data.ReservationDate,
		Customer:                     data.Customer,
		WBSElement:                   data.WBSElement,
		ControllingArea:              data.ControllingArea,
		SalesOrder:                   data.SalesOrder,
		SalesOrderItem:               data.SalesOrderItem,
		SalesOrderScheduleLine:       data.SalesOrderScheduleLine,
		AssetNumber:                  data.AssetNumber,
		AssetSubNumber:               data.AssetSubNumber,
		NetworkNumberForAcctAssgmt:   data.NetworkNumberForAcctAssgmt,
		IssuingOrReceivingPlant:      data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc: data.IssuingOrReceivingStorageLoc,
	}
}

func (sdc *SDC) ConvertToReservationDocumentItem() *requests.ReservationDocumentItem {
	dataReservationDocument := sdc.ReservationDocument
	data := sdc.ReservationDocument.ReservationDocumentItem
	return &requests.ReservationDocumentItem{
		Reservation:                    dataReservationDocument.Reservation,
		ReservationItem:                data.ReservationItem,
		RecordType:                     data.RecordType,
		Product:                        data.Product,
		RequirementType:                data.RequirementType,
		MatlCompRequirementDate:        data.MatlCompRequirementDate,
		Plant:                          data.Plant,
		ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
		GoodsMovementIsAllowed:         data.GoodsMovementIsAllowed,
		StorageLocation:                data.StorageLocation,
		Batch:                          data.Batch,
		DebitCreditCode:                data.DebitCreditCode,
		BaseUnit:                       data.BaseUnit,
		GLAccount:                      data.GLAccount,
		GoodsMovementType:              data.GoodsMovementType,
		EntryUnit:                      data.EntryUnit,
		QuantityIsFixed:                data.QuantityIsFixed,
		CompanyCodeCurrency:            data.CompanyCodeCurrency,
		IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
		IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
		PurchasingDocument:             data.PurchasingDocument,
		PurchasingDocumentItem:         data.PurchasingDocumentItem,
		Supplier:                       data.Supplier,
		ResvnItmRequiredQtyInBaseUnit:  data.ResvnItmRequiredQtyInBaseUnit,
		ReservationItemIsFinallyIssued: data.ReservationItemIsFinallyIssued,
		ReservationItmIsMarkedForDeltn: data.ReservationItmIsMarkedForDeltn,
		ResvnItmRequiredQtyInEntryUnit: data.ResvnItmRequiredQtyInEntryUnit,
		ResvnItmWithdrawnQtyInBaseUnit: data.ResvnItmWithdrawnQtyInBaseUnit,
		ResvnItmWithdrawnAmtInCCCrcy:   data.ResvnItmWithdrawnAmtInCCCrcy,
		GoodsRecipientName:             data.GoodsRecipientName,
		UnloadingPointName:             data.UnloadingPointName,
		ReservationItemText:            data.ReservationItemText,
	}
}
