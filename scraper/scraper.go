package scraper

import (
	"hpk_tg_bot/models"

	"github.com/gocolly/colly"
)

func ScrapGroupReps(groupName string) models.GroupReplacement {
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var groupRes models.GroupReplacement

}
