package type2

import "fmt"

type Student struct {
	BaseProfile
	Scores []float64
	Course string
}

type TeachingAssistant struct {
	BaseProfile
	Scores     []float64
	Course     string
	SessionLed int
}

// Formatting types outputs
func (std Student) String() string {
	return fmt.Sprintf("[ %s ] %s · %s · ▓▓▓▓▓▓▓▓░░ %.1f%", std.EnrollmentID, std.Name, std.Course, std.Scores )
}

func barCreator(score []float64) string {
	
}