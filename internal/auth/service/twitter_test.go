package service

import (
	"fmt"
	twitterClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"testing"
)

func TestService_TwitterUserTweet(t *testing.T) {
	token := oauth1.NewToken("2648158200-NrVrn7465pLkunXfzymUydlH94b9uH95sSsFACr", "KXeXarzrqJDVM7iTNeAxzlmsG14SISMSge2IRZ8qgfQyK")
	oaConfig := oauth1.NewConfig("jcLlFUfv3ENtacKegl6tC2OHw", "R5DiWaBrT0rdjFB7phPAbD90vcWeh020EHHJb16JHOUUv9q7O4")
	httpClient := oaConfig.Client(oauth1.NoContext, token)
	client := twitterClient.NewClient(httpClient)
	list, _, err := client.Users.Show(nil)
	if err != nil {
		//log.Errorv("List error", zap.Error(err))
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
