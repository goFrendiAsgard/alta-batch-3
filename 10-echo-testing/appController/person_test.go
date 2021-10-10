package appController

import (
	"bytes"
	"encoding/json"
	"gofrendi/structureExample/appModel"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPersonAddValid(t *testing.T) {
	e, pc := initTestEcho()
	// compose request
	newPerson, err := json.Marshal(map[string]string{
		"name":     "dono",
		"email":    "dono@warkop.id",
		"password": "rahasia",
	})
	if err != nil {
		t.Errorf("marshalling new person failed")
	}
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPerson))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/persons")
	// send request
	if err = pc.Add(c); err != nil {
		t.Errorf("should not get error, get error: %s", err)
		return
	}
	if rec.Code != 200 {
		t.Errorf("should return 200, get: %d", rec.Code)
	}
	// compare response
	var p appModel.Person
	if err = json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Errorf("unmarshalling returned person failed")
	}
	expectedName := "dono"
	if p.Name != expectedName {
		t.Errorf("person name should be %s, get: %s", expectedName, p.Name)
	}
	expectedEmail := "dono@warkop.id"
	if p.Email != expectedEmail {
		t.Errorf("person email should be %s, get: %s", expectedEmail, p.Email)
	}
	expectedPassword := "rahasia"
	if p.Password != expectedPassword {
		t.Errorf("person pasword should be %s, get: %s", expectedPassword, p.Password)
	}
}
