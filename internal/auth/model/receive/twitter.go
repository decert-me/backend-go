package receive

type TweetsGenerated struct {
	Data []struct {
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		ID                  string   `json:"id"`
		Text                string   `json:"text"`
	} `json:"data"`
}
