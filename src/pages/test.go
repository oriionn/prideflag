package pages

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strconv"
	"time"

	"gorm.io/gorm"
	"prideflag.fun/src/data"
	"prideflag.fun/src/database"
)

//go:embed templates/test.html
var testContent string

type TestPageData struct {
	Choices [4]database.Choices
	Flag    uint
	Test    database.Test
	Max     int
}

func Test(db *gorm.DB, ctx context.Context) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("current_test")
		query := r.URL.Query()
		hasTestCookie := err == nil

		if hasTestCookie {
			_, err := gorm.G[database.Test](db).Where("id = ?", cookie.Value).First(ctx)
			if err != nil {
				hasTestCookie = false
			}
		}

		if !hasTestCookie || query.Has("restart") {
			newTest := &database.Test{}

			err := gorm.G[database.Test](db).Create(ctx, newTest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			cookie = &http.Cookie{
				Name: "current_test",
				Value: strconv.FormatUint(uint64(newTest.ID), 10),
				Expires: time.Now().Add(48 * time.Hour),
				Path: "/",
				SameSite: http.SameSiteStrictMode,
			}

			http.SetCookie(w, cookie)
			RollChoices(newTest.ID, db, ctx)

			if query.Has("restart") {
				http.Redirect(w, r, "/test", 307)
				return
			}
		}

		test, err := gorm.G[database.Test](db).Where("id = ?", cookie.Value).First(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if test.Total >= len(data.DATASET) {
			http.Redirect(w, r, fmt.Sprintf("/results?t=%d", test.ID), 307)
			return
		}

		choices, trueChoice, err := GetChoices(test.ID, db, ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		validChoicesID := []uint{}
		for _, choice := range choices {
			validChoicesID = append(validChoicesID, choice.ID)
		}

		if query.Has("a") {
			answer := query.Get("a")
			answerInt, err := strconv.Atoi(answer)
			if err == nil && slices.Contains(validChoicesID, uint(answerInt)) {
				if int(trueChoice.ID) == answerInt {
					test.Note += 1
				}
				test.Total += 1

				gorm.G[database.Test](db).Where("id = ?", test.ID).Update(ctx, "Note", test.Note)
				gorm.G[database.Test](db).Where("id = ?", test.ID).Update(ctx, "Total", test.Total)

				RollChoices(test.ID, db, ctx)

				choices, trueChoice, err = GetChoices(test.ID, db, ctx)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		flag := database.Images{
			File: trueChoice.File,
		}

		gorm.G[database.Images](db).Create(ctx, &flag)

		// Templating
		t, err := template.New("test").Parse(testContent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TestPageData{
			Choices: [4]database.Choices(choices),
			Flag: flag.ID,
			Test: test,
			Max: len(data.DATASET),
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, data)
	}
}

func RollChoices(test_id uint, db *gorm.DB, ctx context.Context) {
	flagExists := []string{}
	flagInTheTest, _ := gorm.G[database.FlagInTheTest](db).Where("test_id = ?", test_id).Find(ctx)
	for _, flag := range flagInTheTest {
		flagExists = append(flagExists, flag.Flag)
	}

	choices, trueFlag, err := data.GetChoices(flagExists)
	if err != nil {
		RollChoices(test_id, db, ctx)
		return
	}

	db.Unscoped().Where("test_id = ?", test_id).Delete(&database.Choices{})
	gorm.G[database.FlagInTheTest](db).Create(ctx, &database.FlagInTheTest{
		TestID: test_id,
		Flag: choices[trueFlag].File,
	})

	for i, choice := range choices {
		choiceData := database.Choices{
			Name: choice.Name,
			File: choice.File,
			TestID: test_id,
			TrueFlag: i == int(trueFlag),
		}
		gorm.G[database.Choices](db).Create(ctx, &choiceData)
	}
}

func GetChoices(test_id uint, db *gorm.DB, ctx context.Context) ([]database.Choices, database.Choices, error) {
	choices, err := gorm.G[database.Choices](db).Where("test_id = ?", test_id).Find(ctx)
	if err != nil {
		return nil, database.Choices{}, err
	}

	var trueChoice database.Choices = database.Choices{Name: ""};
	for _, choice := range choices {
		if choice.TrueFlag == true {
			trueChoice = choice
		}
	}

	if trueChoice.Name == "" {
		return nil, database.Choices{}, errors.New("Internal Server Error")
	}

	return choices, trueChoice, nil
}
