package handler

import (
	"github.com/DATA-DOG/go-sqlmock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// URL, по которому будет отправлен запрос
	url := "http://0.0.0.0:8087/v1/promotions"

	// Создаем новый запрос POST
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	// Устанавливаем необходимые заголовки
	req.Header.Set("Content-Type", "application/json")

	// Создаем клиент для отправки запроса
	client := &http.Client{}

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// now we execute our request
	if resp.StatusCode != 200 {
		t.Fatalf("expected status code to be 200, but got: %docker", resp.StatusCode)
	}

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create app with mocked db, request and response to test
	_, err = http.NewRequest("POST", "http://0.0.0.0:8087/v1/filters", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// before we actually execute our api function, we need to expect required DB actions
	rows := sqlmock.NewRows([]string{"id", "title", "body"}).
		AddRow(1, "post 1", "hello").
		AddRow(2, "post 2", "world")

	mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// now we execute our request
	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %docker", w.Code)
	}

	// we make sure that all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
