package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	SalesInquiry  struct {
		SalesInquiry    string `json:"document_no"`
		Plant           string `json:"deliver_to"`
		InquiryQuantity string `json:"quantity"`
		PickedQuantity  string `json:"picked_quantity"`
		NetPriceAmount  string `json:"price"`
		Batch           string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema        string `json:"api_schema"`
	MaterialCode     string `json:"material_code"`
	Supplier         string `json:"plant/supplier"`
	Stock            string `json:"stock"`
	SalesInquiryType string `json:"document_type"`
	SINumber         string `json:"document_no"`
	PlannedDate      string `json:"planned_date"`
	ValidatedDate    string `json:"validated_date"`
	Deleted          bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	SalesInquiry  struct {
		SalesInquiry                   string `json:"SalesInquiry"`
		SalesInquiryType               string `json:"SalesInquiryType"`
		SalesOrganization              string `json:"SalesOrganization"`
		DistributionChannel            string `json:"DistributionChannel"`
		OrganizationDivision           string `json:"OrganizationDivision"`
		SalesGroup                     string `json:"SalesGroup"`
		SalesOffice                    string `json:"SalesOffice"`
		SalesDistrict                  string `json:"SalesDistrict"`
		SoldToParty                    string `json:"SoldToParty"`
		CreationDate                   string `json:"CreationDate"`
		LastChangeDate                 string `json:"LastChangeDate"`
		PurchaseOrderByCustomer        string `json:"PurchaseOrderByCustomer"`
		CustomerPurchaseOrderType      string `json:"CustomerPurchaseOrderType"`
		CustomerPurchaseOrderDate      string `json:"CustomerPurchaseOrderDate"`
		SalesInquiryDate               string `json:"SalesInquiryDate"`
		TotalNetAmount                 string `json:"TotalNetAmount"`
		TransactionCurrency            string `json:"TransactionCurrency"`
		SDDocumentReason               string `json:"SDDocumentReason"`
		PricingDate                    string `json:"PricingDate"`
		HeaderBillingBlockReason       string `json:"HeaderBillingBlockReason"`
		BindingPeriodValidityStartDate string `json:"BindingPeriodValidityStartDate"`
		BindingPeriodValidityEndDate   string `json:"BindingPeriodValidityEndDate"`
		HdrOrderProbabilityInPercent   string `json:"HdrOrderProbabilityInPercent"`
		ExpectedOrderNetAmount         string `json:"ExpectedOrderNetAmount"`
		IncotermsClassification        string `json:"IncotermsClassification"`
		CustomerPaymentTerms           string `json:"CustomerPaymentTerms"`
		PaymentMethod                  string `json:"PaymentMethod"`
		OverallSDProcessStatus         string `json:"OverallSDProcessStatus"`
		TotalCreditCheckStatus         string `json:"TotalCreditCheckStatus"`
		OverallSDDocumentRejectionSts  string `json:"OverallSDDocumentRejectionSts"`
		HeaderPartner                  struct {
			PartnerFunction string `json:"PartnerFunction"`
			Customer        string `json:"Customer"`
			Supplier        string `json:"Supplier"`
		} `json:"HeaderPartner"`
		SalesInquiryItem struct {
			SalesInquiryItem              string `json:"SalesInquiryItem"`
			SalesInquiryItemCategory      string `json:"SalesInquiryItemCategory"`
			SalesInquiryItemText          string `json:"SalesInquiryItemText"`
			PurchaseOrderByCustomer       string `json:"PurchaseOrderByCustomer"`
			Material                      string `json:"Material"`
			MaterialByCustomer            string `json:"MaterialByCustomer"`
			RequestedQuantity             string `json:"RequestedQuantity"`
			RequestedQuantityUnit         string `json:"RequestedQuantityUnit"`
			ItemOrderProbabilityInPercent string `json:"ItemOrderProbabilityInPercent"`
			ItemGrossWeight               string `json:"ItemGrossWeight"`
			ItemNetWeight                 string `json:"ItemNetWeight"`
			ItemWeightUnit                string `json:"ItemWeightUnit"`
			ItemVolume                    string `json:"ItemVolume"`
			ItemVolumeUnit                string `json:"ItemVolumeUnit"`
			TransactionCurrency           string `json:"TransactionCurrency"`
			NetAmount                     string `json:"NetAmount"`
			MaterialGroup                 string `json:"MaterialGroup"`
			Batch                         string `json:"Batch"`
			IncotermsClassification       string `json:"IncotermsClassification"`
			CustomerPaymentTerms          string `json:"CustomerPaymentTerms"`
			SalesDocumentRjcnReason       string `json:"SalesDocumentRjcnReason"`
			WBSElement                    string `json:"WBSElement"`
			SDProcessStatus               string `json:"SDProcessStatus"`
			ItemPricingElement            struct {
				PricingProcedureStep        string `json:"PricingProcedureStep"`
				PricingProcedureCounter     string `json:"PricingProcedureCounter"`
				ConditionApplication        string `json:"ConditionApplication"`
				ConditionType               string `json:"ConditionType"`
				PricingDateTime             string `json:"PricingDateTime"`
				ConditionCalculationType    string `json:"ConditionCalculationType"`
				ConditionBaseValue          string `json:"ConditionBaseValue"`
				ConditionRateValue          string `json:"ConditionRateValue"`
				ConditionCurrency           string `json:"ConditionCurrency"`
				ConditionQuantity           string `json:"ConditionQuantity"`
				ConditionQuantityUnit       string `json:"ConditionQuantityUnit"`
				ConditionCategory           string `json:"ConditionCategory"`
				PricingScaleType            string `json:"PricingScaleType"`
				ConditionRecord             string `json:"ConditionRecord"`
				ConditionSequentialNumber   string `json:"ConditionSequentialNumber"`
				TaxCode                     string `json:"TaxCode"`
				ConditionAmount             string `json:"ConditionAmount"`
				TransactionCurrency         string `json:"TransactionCurrency"`
				PricingScaleBasis           string `json:"PricingScaleBasis"`
				ConditionScaleBasisValue    string `json:"ConditionScaleBasisValue"`
				ConditionScaleBasisUnit     string `json:"ConditionScaleBasisUnit"`
				ConditionScaleBasisCurrency string `json:"ConditionScaleBasisCurrency"`
				ConditionIsManuallyChanged  string `json:"ConditionIsManuallyChanged"`
			} `json:"ItemPricingElement"`
		} `json:"SalesInquiryItem"`
	} `json:"SalesInquiry"`
	APISchema      string   `json:"api_schema"`
	Accepter       []string `json:"accepter"`
	SalesInquiryNo string   `json:"sales_inquiry"`
	Deleted        bool     `json:"deleted"`
}
