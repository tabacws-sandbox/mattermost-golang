// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package slashcommands

import (
	"testing"

	"github.com/mattermost/mattermost-server/v6/model"
)

func TestGetCustomStatus(t *testing.T) {
	for msg, expected := range map[string]model.CustomStatus{
		"":                         {Emoji: model.DefaultCustomStatusEmoji, Text: ""},
		"Hey":                      {Emoji: model.DefaultCustomStatusEmoji, Text: "Hey"},
		":cactus: Hurt":            {Emoji: "cactus", Text: "Hurt"},
		"π":                        {Emoji: "tongue", Text: ""},
		"π Eating":                 {Emoji: "tongue", Text: "Eating"},
		"πͺπ» Working out":           {Emoji: "muscle_light_skin_tone", Text: "Working out"},
		"π Swimming":               {Emoji: "bikini", Text: "Swimming"},
		"πSwimming":                {Emoji: model.DefaultCustomStatusEmoji, Text: "πSwimming"},
		"ππΏ Okay":                  {Emoji: "+1_dark_skin_tone", Text: "Okay"},
		"π€΄πΎ Dark king":             {Emoji: "prince_medium_dark_skin_tone", Text: "Dark king"},
		"βΉπΎββοΈ Playing basketball": {Emoji: "basketball_woman_medium_dark_skin_tone", Text: "Playing basketball"},
		"ππΏββοΈ Weightlifting":      {Emoji: "weight_lifting_woman_dark_skin_tone", Text: "Weightlifting"},
		"π Surfing":                {Emoji: "surfer", Text: "Surfing"},
		"π¨βπ¨βπ¦βπ¦ Family":           {Emoji: "family_man_man_boy_boy", Text: "Family"},
	} {
		actual := GetCustomStatus(msg)
		if actual.Emoji != expected.Emoji || actual.Text != expected.Text {
			t.Errorf("expected `%v`, got `%v`", expected, *actual)
		}
	}
}
