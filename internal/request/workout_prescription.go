package request

type ExercisePrescription struct {
	ExerciseID int     `json:"exercise_id"`
	MinReps    int     `json:"min_reps"`
	MaxReps    int     `json:"max_reps"`
	Weight     float32 `json:"weight"`
	Duration   float32 `json:"duration"`
	MinRest    int     `json:"min_rest"`
	MaxRest    int     `json:"max_rest"`
}

type BlockPrescription struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Exercises   []ExercisePrescription `json:"exercises"`
}

type WorkoutPrescriptionRequest struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Blocks      []BlockPrescription `json:"blocks"`
}
