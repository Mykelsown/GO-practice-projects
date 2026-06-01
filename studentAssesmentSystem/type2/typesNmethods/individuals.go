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
	return fmt.Sprintf("[ %s ] Amara Osei · Go Programming · ▓▓▓▓▓▓░░ 78.5%", std.EnrollmentID, std.Name, )
}

