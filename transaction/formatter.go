package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	ID       int               `json:"id"`
	Amount   int               `json:"amount"`
	Status   string            `json:"status"`
	CratedAt time.Time         `json:"created_at"`
	Campaign CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatCampaignTransaction(trs Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID:        trs.ID,
		Name:      trs.User.Name,
		Amount:    trs.Amount,
		CreatedAt: trs.CreatedAt,
	}
	return formatter
}

func FormatCampaignTransactions(campaignTrsFormatter []Transaction) []CampaignTransactionFormatter {
	campaignTransactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range campaignTrsFormatter {
		formatCampaignTrs := FormatCampaignTransaction(transaction)
		campaignTransactionsFormatter = append(campaignTransactionsFormatter, formatCampaignTrs)
	}

	return campaignTransactionsFormatter
}

func FormatUserTransaction(trx Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{
		ID:       trx.ID,
		Amount:   trx.Amount,
		Status:   trx.Status,
		CratedAt: trx.CreatedAt,
	}

	formatCampaign := CampaignFormatter{
		Name:     trx.Campaign.Name,
		ImageUrl: "",
	}

	if len(trx.Campaign.CampaignImages) > 0 {
		formatCampaign.ImageUrl = trx.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = formatCampaign

	return formatter
}

func FormatUserTransactions(trxs []Transaction) []UserTransactionFormatter {
	userTrxsFormatter := []UserTransactionFormatter{}

	for _, trx := range trxs {
		formatUserTrx := FormatUserTransaction(trx)
		userTrxsFormatter = append(userTrxsFormatter, formatUserTrx)
	}
	return userTrxsFormatter
}

func FormatTransaction(trs Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         trs.ID,
		CampaignID: trs.CampaignID,
		Amount:     trs.Amount,
		Status:     trs.Status,
		UserID:     trs.User.ID,
		Code:       trs.Code,
		PaymentURL: trs.PaymentURL,
	}
	return formatter
}
