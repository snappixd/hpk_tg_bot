package main

import (
	"encoding/json"
	"fmt"
	"hpk_tg_bot/models"
	"hpk_tg_bot/scraper"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("app start")

	//example
	todayName := "Wednesday"
	var groupName string = "КІ-213"

	groupReps, err := scraper.ScrapGroupReps(groupName)
	if err != nil {
		log.Println(err.Error())
	}

	day := ParseDaySchedule("schedule.json", groupName, todayName)
	updatedDay := UpdateDayWithReps(&day, groupReps)

	fmt.Printf("\t%s\n", updatedDay.DayName)
	fmt.Printf("\t%s\n", groupName)

	for _, para := range updatedDay.Paras {
		fmt.Printf("Пари: %s\nПредмет: %s\nВикладач: %s\nАудиторія: %s\n\n", para.ParasNum, para.Subject, para.Teacher, para.Classroom)
	}

	log.Println("app finish")
}

func UpdateDayWithReps(day *models.Day, reps models.GroupReplacement) *models.Day {
	for i := range day.Paras {
		for _, rep := range reps.Replacements {
			if day.Paras[i].ParasNum == rep.ParasNum {
				day.Paras[i].Classroom = rep.Classroom
				day.Paras[i].Teacher = rep.Teacher
				day.Paras[i].Subject = rep.Subject
			}
		}
	}

	return day
}

func ParseDaySchedule(fileName, groupName, dayName string) models.Day {
	var schedule models.Schedule

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}

	err = json.Unmarshal(byteValue, &schedule)
	if err != nil {
		log.Println(err.Error())
	}

	for _, group := range schedule.Groups {
		if group.GroupName == groupName {
			for _, day := range group.Days {
				if day.DayName == dayName {
					return day
				}
			}
		}
	}

	return models.Day{}
}
