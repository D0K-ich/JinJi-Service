package models

import "errors"

func(u *User) NewAchievement(achievement *Achievement) (err error) {
	if u == nil {return errors.New("empty user")}
	if achievement == nil {return errors.New("empty achievement")}

	for _, arch := range u.Achievements.Achievements {if arch.Name == achievement.Name {return}}
	u.Achievements.Achievements = append(u.Achievements.Achievements, achievement)
	return
}
