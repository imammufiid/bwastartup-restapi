package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(trs Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID: trs.ID,
		Name: trs.User.Name,
		Amount: trs.Amount,
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
