package main

import (
	type2 "SAS/type2/typesNmethods"
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
		Scores: []float64{88.5, 95, 60, 56.4, 90.2},
		Course: "AI/ML Enginnering",
		SessionLed: 2,
	}

	TA2 := type2.TeachingAssistant{
		BaseProfile: type2.BaseProfile{
			Name:         "Itunu Emmanuel",
			Email:        "Eetunu@gmail.com",
			EnrollmentID: "EDU-2502",
		},
		Scores: []float64{65.5, 98.5, 87.3, 56.4, 76.8},
		Course: "AI/ML Enginnering",
		SessionLed: 4,
	}

	Students := []type2.Student{student1, student2}
	TAs := []type2.TeachingAssistant{TA1, TA2}
	
}

func displayLegitimate[T type2.Student|type2.TeachingAssistant](learners []T) {
	for _, learner := range learners {
		
	}
}