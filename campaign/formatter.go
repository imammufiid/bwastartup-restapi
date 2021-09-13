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
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	BackerCount      int                      `json:"backer_count"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             FormatterUser            `json:"user"`
	Images           []FormatterCampaignImage `json:"images"`
}

type FormatterUser struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
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
	campaignDetailFormatter.BackerCount = campaign.BackerCount
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

	userFormatter := FormatterUser{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.Avatar,
	}
	campaignDetailFormatter.User = userFormatter

	images := []FormatterCampaignImage{}
	for _, image := range campaign.CampaignImages {
		imageFormatter := FormatCampaignImage(image)
		images = append(images, imageFormatter)
	}
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}

func FormatCampaignImage(campaign CampaignImage) FormatterCampaignImage {
	formatCampaignImage := FormatterCampaignImage{}

	formatCampaignImage.ImageUrl = campaign.FileName
	isPrimary := false
	if campaign.IsPrimary == 1 {
		isPrimary = true
	}
	formatCampaignImage.IsPrimary = isPrimary

	return formatCampaignImage
}
