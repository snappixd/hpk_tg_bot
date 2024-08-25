package main

import (
	"fmt"
	"hpk_tg_bot/scraper"
	"log"
)

func main() {
	log.Println("app start")

	groupReps, err := scraper.ScrapGroupReps("КІ-232")
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("\t%s\n", groupReps.GroupName)
	for _, rep := range groupReps.Replacements {
		fmt.Printf("Пари: %s\nПредмет: %s\nВикладач: %s\nАудиторія: %s\n\n", rep.ParasNum, rep.Subject, rep.Teacher, rep.Classroom)
	}

	log.Println("app finish")
}
