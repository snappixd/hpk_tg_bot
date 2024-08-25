package scraper

import (
	"errors"
	"hpk_tg_bot/models"

	"github.com/gocolly/colly"
)

func ScrapGroupReps(targerGroupName string) (models.GroupReplacement, error) {
	c := colly.NewCollector()

	var isTargerGroup bool

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	var groupRes models.GroupReplacement

	c.OnHTML("tr:nth-child(n+3)", func(e *colly.HTMLElement) {
		groupName := e.ChildText("td:nth-child(1)")

		if groupName != "" {
			isTargerGroup = groupName == targerGroupName
			if isTargerGroup {
				groupRes.GroupName = groupName
			}
		}

		if isTargerGroup {
			classroom := e.ChildText("td:nth-child(6)")
			if classroom == "" {
				classroom = "-"
			}

			replacement := models.Para{
				ParasNum:  e.ChildText("td:nth-child(2)"),
				Subject:   e.ChildText("td:nth-child(4)"),
				Teacher:   e.ChildText("td:nth-child(5)"),
				Classroom: classroom,
			}

			if replacement.ParasNum != "" || replacement.Subject != "" || replacement.Teacher != "" || replacement.Classroom != "" {
				groupRes.Replacements = append(groupRes.Replacements, replacement)
			}
		}
	})
	c.Visit("https://hpk.edu.ua/replacements")

	if len(groupRes.Replacements) == 0 {
		return models.GroupReplacement{}, errors.New("заміни у вас відсутні")
	}

	return groupRes, nil
}
