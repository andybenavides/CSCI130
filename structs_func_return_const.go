package main

import "fmt"

type Seasons struct {
	season_type string
	warning     string
}

func warning(season Seasons) {
	_, warning := CreateSeasonWarning(season.season_type, season.warning, "Practice precaution")

	fmt.Println(warning)
}

func CreateSeasonWarning(season_type string, warning ...string) (alert string, alternate string) {
	alert = warning[1] + " while in " + season_type
	alternate = "Caution: " + season_type + " brings cold temperatures."
	return
}

func main() {
	var season = Seasons{"Winter", "Prepare for cold winds."}

	warning(season)
}
