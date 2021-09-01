package campaign

import (
	"strings"
)

type FormatterCampaign struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type FormatterDetailCampaign struct {
	ID               int                       `json:"id"`
	UserID           int                       `json:"user_id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageUrl         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             FormatterUser             `json:"user"`
	Images           []FormatterCampaignImage `json:"images"`
}

type FormatterUser struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

type FormatterCampaignImage struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) FormatterCampaign {
	campaignFormatter := FormatterCampaign{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ImageUrl = ""
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []FormatterCampaign {
	campaignsFormatter := []FormatterCampaign{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

func FormatDetailCampaign(campaign Campaign) FormatterDetailCampaign {
	campaignDetailFormatter := FormatterDetailCampaign{}

	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, "-") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	campaignDetailFormatter.User.Name = campaign.User.Name
	campaignDetailFormatter.User.Avatar = campaign.User.Avatar

	// images := []FormatterCampaignImage{}
	// for _, image := range campaign.CampaignImages {
	// 	imageFormatter := FormatCampaignImage(image)
	// 	images = append(images, imageFormatter) 
	// }
	// campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}


// func FormatCampaignImage(campaign CampaignImage) FormatterCampaignImage{
// 	formatCampaignImage := FormatterCampaignImage{}

// 	formatCampaignImage.ImageUrl = campaign.FileName
// 	formatCampaignImage.IsPrimary = false 
// 	if campaign.IsPrimary == 1 {
// 		formatCampaignImage.IsPrimary = true
// 	}

// 	return formatCampaignImage
// }