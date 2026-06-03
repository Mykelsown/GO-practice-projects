package main

import (
	type2 "SAS/type2/typesNmethods"
	"fmt"
)

func main() {
	student1 := type2.Student{
		BaseProfile: type2.BaseProfile{
			Name:         "Emmanuel Usang",
			Email:        "emmanuelsansai@gmail.com",
			EnrollmentID: "EDU-2512",
		},
		Scores: []float64{72.7, 90.2, 67, 83.3, 75},
		Course: "Go Programming",
	}

	student2 := type2.Student{
		BaseProfile: type2.BaseProfile{
			Name:         "Ogunrotimi Damilola",
			Email:        "complxDharmi@gmail.com",
			EnrollmentID: "EDU-2511",
		},
		Scores: []float64{85.7, 70.2, 90, 89, 72.3},
		Course: "Go Programming",
	}

	TA1 := type2.TeachingAssistant{
		BaseProfile: type2.BaseProfile{
			Name:         "Ezekiel Leke",
			Email:        "Lekezziel@gmail.com",
			EnrollmentID: "EDU-2500",
		},
		Scores:     []float64{88.5, 95, 80, 85.4, 90.2},
		Course:     "AI/ML Enginnering",
		SessionLed: 2,
	}

	TA2 := type2.TeachingAssistant{
		BaseProfile: type2.BaseProfile{
			Name:         "Itunu Emmanuel",
			Email:        "Eetunu@gmail.com",
			EnrollmentID: "EDU-2502",
		},
		Scores:     []float64{65.5, 98.5, 87.3, 56.4, 76.8},
		Course:     "AI/ML Enginnering",
		SessionLed: 5,
	}

	students := []type2.Student{student1, student2}
	TAs := []type2.TeachingAssistant{TA1, TA2}

	for _, student := range students {
		displayLegitimate(student)
	}
	for _, TA := range TAs {
		displayLegitimate(TA)
	}

	// ---------stuent1 Data Render Section------
	_, score := student1.Progress()
	st1Data := type2.DisplayData{
		Receipent: student1.Name,
		Course:    student1.Course,
		Score:     score,
		Grade:     student1.Grade(),
	}
	type2.Render(st1Data)

	fmt.Println()

	fmt.Printf("(*) Best Graduating Student: %s\n\n", type2.TopPerformer(students).Name)
	fmt.Printf("(*) Best Graduating TeachingAssistant: %s\n\n", type2.TopPerformer(TAs).Name)
}

func displayLegitimate(learner any) {
	switch j := learner.(type) {
	case type2.Student:
		real := learner.(type2.Student)
		if real.Validate() != nil {
			fmt.Printf("⚠  Skipping invalid profile: %s\n\n", real.Validate())
			break
		}
		fmt.Println(real, "\n")
	case type2.TeachingAssistant:
		real := learner.(type2.TeachingAssistant)
		if real.Validate() != nil {
			fmt.Printf("⚠  Skipping invalid profile: %s\n\n", real.Validate())
			break
		}
		fmt.Println(real, "\n")
	default:
		fmt.Printf("Wrong type passed: %T", j)
	}
}
