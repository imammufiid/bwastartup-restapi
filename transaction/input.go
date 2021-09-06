package transaction

import "bwastartup/user"

type GetCampaignTransactionsInput struct {
	campaignID int `uri:"id" binding:"required"`
	User       user.User
}
