package models

type DaySummary struct {
	Meals         []MealLog `json:"meals"`
	TotalCalories float64   `json:"total_calories"`
	TotalProteins float64   `json:"total_proteins"`
	TotalFats     float64   `json:"total_fats"`
	TotalCarbs    float64   `json:"total_carbs"`
	DailyNorm     float64   `json:"daily_norm"`
	Remaining     float64   `json:"remaining"`
}
