package handler

import "bwastartup/campaign"

type CampaignHandler interface {
}

type campaignHandler struct {
	campaignService campaign.Service
}

func InstanceCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService: campaignService}
}