package repositories_test

import (
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "sub1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  3,
			},
		}, {
			testTitle: "sub2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %d but want %d", num, expectedNum)
	}
}

//func TestInsertArticle(t *testing.T) {
//	article := models.Article{
//		Title:    "insertTest",
//		Contents: "testest",
//		UserName: "saki",
//	}
//	expectedArticleNum := 3
//	newArticle, err := repositories.InsertArticle(testDB, article)
//	if err != nil {
//		t.Error(err)
//	}
//	if newArticle.ID != expectedArticleNum {
//		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
//	}
//
//	t.Cleanup(func() {
//		const sqlStr = `
//            delete from articles
//            where title = ? and contents = ? and username = ?;
//        `
//		_, err := testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
//		if err != nil {
//			t.Error(err)
//		}
//	})
//}
