package repositories_test

import (
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != article.Title {
		t.Errorf("got %s but want %s", newArticle.Title, article.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
           delete from articles
           where article_id = ?;
       `
		_, err := testDB.Exec(sqlStr, newArticle.ID)
		if err != nil {
			t.Error(err)
		}
	})
}

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
			//if got.NiceNum != test.expected.NiceNum {
			//	t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			//}
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

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	article, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	updatedArticle, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}
	if article.NiceNum+1 != updatedArticle.NiceNum {
		t.Errorf("got %d but want %d", article.NiceNum, updatedArticle.NiceNum)
	}

}

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("got %d but want %d", num, expectedNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test comment",
	}

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.Message != comment.Message {
		t.Errorf("got %s but want %s", newComment.Message, comment.Message)
	}

	t.Cleanup(func() {
		const sqlStr = `
		   delete from comments
		   where comment_id = ?;
	   `
		_, err := testDB.Exec(sqlStr, newComment.CommentID)
		if err != nil {
			t.Error(err)
		}
	})
}
