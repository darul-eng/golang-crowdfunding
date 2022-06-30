package transaction

import "golang-vue-nuxtjs-crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
