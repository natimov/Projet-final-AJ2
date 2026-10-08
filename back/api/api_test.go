package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"meridian/back/api/config"
	"meridian/back/api/database"
	"meridian/back/api/models"
)

func newTestAPI(t *testing.T) (*gin.Engine, *gorm.DB, string, models.User) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(filepath.Join(t.TempDir(), "test.db")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.CategoriePrestation{}, &models.Prestation{}, &models.Evenement{}); err != nil {
		t.Fatal(err)
	}

	previousDB := database.DB
	previousSecret := config.JWTSecret
	database.DB = db
	config.JWTSecret = []byte("test-secret")
	t.Cleanup(func() {
		database.DB = previousDB
		config.JWTSecret = previousSecret
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	admin := models.User{Nom: "Admin", Email: "admin@example.com", Role: "admin", StatutCompte: "actif"}
	if err := db.Create(&admin).Error; err != nil {
		t.Fatal(err)
	}

	token := makeTestToken(t, admin.ID, "admin")
	router := gin.New()
	setupRoutes(router)
	return router, db, token, admin
}

func makeTestToken(t *testing.T, userID uint, role string) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	result, err := token.SignedString(config.JWTSecret)
	if err != nil {
		t.Fatal(err)
	}
	return result
}

func callAPI(router http.Handler, method, path, body, token string) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}

func TestRegisterAndLogin(t *testing.T) {
	router, _, _, _ := newTestAPI(t)

	register := callAPI(router, http.MethodPost, "/register", `{"nom":"Test","prenom":"User","email":"test@example.com","password":"secret"}`, "")
	if register.Code != http.StatusCreated {
		t.Fatalf("register status = %d, want %d", register.Code, http.StatusCreated)
	}

	invalidRegister := callAPI(router, http.MethodPost, "/register", `{"email":"incomplet@example.com"}`, "")
	if invalidRegister.Code != http.StatusBadRequest {
		t.Fatalf("invalid register status = %d, want %d", invalidRegister.Code, http.StatusBadRequest)
	}

	login := callAPI(router, http.MethodPost, "/login", `{"email":"test@example.com","password":"secret"}`, "")
	if login.Code != http.StatusOK {
		t.Fatalf("login status = %d, want %d", login.Code, http.StatusOK)
	}

	wrongPassword := callAPI(router, http.MethodPost, "/login", `{"email":"test@example.com","password":"wrong"}`, "")
	if wrongPassword.Code != http.StatusUnauthorized {
		t.Fatalf("wrong password status = %d, want %d", wrongPassword.Code, http.StatusUnauthorized)
	}
}

func TestAuthenticationAndAdminAccess(t *testing.T) {
	router, db, adminToken, _ := newTestAPI(t)

	noToken := callAPI(router, http.MethodGet, "/users", "", "")
	if noToken.Code != http.StatusUnauthorized {
		t.Fatalf("no token status = %d, want %d", noToken.Code, http.StatusUnauthorized)
	}
	invalidToken := callAPI(router, http.MethodGet, "/users", "", "not-a-jwt")
	if invalidToken.Code != http.StatusUnauthorized {
		t.Fatalf("invalid token status = %d, want %d", invalidToken.Code, http.StatusUnauthorized)
	}

	malformedClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "incorrect",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	malformedToken, err := malformedClaims.SignedString(config.JWTSecret)
	if err != nil {
		t.Fatal(err)
	}
	malformed := callAPI(router, http.MethodGet, "/users", "", malformedToken)
	if malformed.Code != http.StatusUnauthorized {
		t.Fatalf("malformed user_id status = %d, want %d", malformed.Code, http.StatusUnauthorized)
	}

	visitor := models.User{Nom: "Visitor", Email: "visitor@example.com", Role: "particulier", StatutCompte: "actif"}
	if err := db.Create(&visitor).Error; err != nil {
		t.Fatal(err)
	}
	visitorToken := makeTestToken(t, visitor.ID, visitor.Role)
	forbidden := callAPI(router, http.MethodGet, "/users", "", visitorToken)
	if forbidden.Code != http.StatusForbidden {
		t.Fatalf("non-admin status = %d, want %d", forbidden.Code, http.StatusForbidden)
	}

	allowed := callAPI(router, http.MethodGet, "/users", "", adminToken)
	if allowed.Code != http.StatusOK {
		t.Fatalf("admin status = %d, want %d", allowed.Code, http.StatusOK)
	}
}

func TestUserCRUD(t *testing.T) {
	router, _, token, _ := newTestAPI(t)

	created := callAPI(router, http.MethodPost, "/users", `{"nom":"Test","prenom":"User","email":"crud@example.com","password":"secret","role":"particulier"}`, token)
	if created.Code != http.StatusCreated {
		t.Fatalf("create user status = %d, want %d", created.Code, http.StatusCreated)
	}
	var user models.User
	if err := json.Unmarshal(created.Body.Bytes(), &user); err != nil {
		t.Fatal(err)
	}
	id := strconv.FormatUint(uint64(user.ID), 10)

	if response := callAPI(router, http.MethodGet, "/users/"+id, "", token); response.Code != http.StatusOK {
		t.Fatalf("get user status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodPut, "/users/"+id, `{"ville":"Paris"}`, token); response.Code != http.StatusOK {
		t.Fatalf("update user status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/users/99999", "", token); response.Code != http.StatusNotFound {
		t.Fatalf("missing user status = %d, want %d", response.Code, http.StatusNotFound)
	}
	if response := callAPI(router, http.MethodDelete, "/users/"+id, "", token); response.Code != http.StatusOK {
		t.Fatalf("delete user status = %d, want %d", response.Code, http.StatusOK)
	}
}

func TestCategoriesAndPrestations(t *testing.T) {
	router, _, token, _ := newTestAPI(t)

	categoryResponse := callAPI(router, http.MethodPost, "/categories-prestation", `{"libelle":"Réparation"}`, token)
	if categoryResponse.Code != http.StatusCreated {
		t.Fatalf("create category status = %d, want %d", categoryResponse.Code, http.StatusCreated)
	}
	var category models.CategoriePrestation
	if err := json.Unmarshal(categoryResponse.Body.Bytes(), &category); err != nil {
		t.Fatal(err)
	}
	categoryID := strconv.FormatUint(uint64(category.ID), 10)

	if response := callAPI(router, http.MethodGet, "/categories-prestation", "", token); response.Code != http.StatusOK {
		t.Fatalf("list categories status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/categories-prestation/"+categoryID, "", token); response.Code != http.StatusOK {
		t.Fatalf("get category status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodPut, "/categories-prestation/"+categoryID, `{"libelle":"Atelier"}`, token); response.Code != http.StatusOK {
		t.Fatalf("update category status = %d, want %d", response.Code, http.StatusOK)
	}

	prestationResponse := callAPI(router, http.MethodPost, "/prestations", `{"nom":"Atelier vélo","categorie_id":`+categoryID+`}`, token)
	if prestationResponse.Code != http.StatusCreated {
		t.Fatalf("create prestation status = %d, want %d: %s", prestationResponse.Code, http.StatusCreated, prestationResponse.Body.String())
	}
	var prestation models.Prestation
	if err := json.Unmarshal(prestationResponse.Body.Bytes(), &prestation); err != nil {
		t.Fatal(err)
	}
	if prestation.CategorieID != category.ID {
		t.Fatalf("prestation category id = %d, want %d", prestation.CategorieID, category.ID)
	}
	prestationID := strconv.FormatUint(uint64(prestation.ID), 10)

	if response := callAPI(router, http.MethodGet, "/prestations", "", token); response.Code != http.StatusOK {
		t.Fatalf("list prestations status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/prestations/"+prestationID, "", token); response.Code != http.StatusOK {
		t.Fatalf("get prestation status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/prestations/99999", "", token); response.Code != http.StatusNotFound {
		t.Fatalf("missing prestation status = %d, want %d", response.Code, http.StatusNotFound)
	}
	if response := callAPI(router, http.MethodPut, "/prestations/"+prestationID, `{"nom":"Atelier réparé"}`, token); response.Code != http.StatusOK {
		t.Fatalf("update prestation status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodDelete, "/prestations/"+prestationID, "", token); response.Code != http.StatusOK {
		t.Fatalf("delete prestation status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodDelete, "/categories-prestation/"+categoryID, "", token); response.Code != http.StatusOK {
		t.Fatalf("delete category status = %d, want %d", response.Code, http.StatusOK)
	}
}

func TestEvenementCRUD(t *testing.T) {
	router, _, token, _ := newTestAPI(t)

	created := callAPI(router, http.MethodPost, "/evenements", `{"type":"atelier","titre":"Atelier test","date_debut":"2026-10-01","date_fin":"2026-10-01","lieu":"Paris"}`, token)
	if created.Code != http.StatusCreated {
		t.Fatalf("create event status = %d, want %d", created.Code, http.StatusCreated)
	}
	var evenement models.Evenement
	if err := json.Unmarshal(created.Body.Bytes(), &evenement); err != nil {
		t.Fatal(err)
	}
	id := strconv.FormatUint(uint64(evenement.ID), 10)

	if response := callAPI(router, http.MethodGet, "/evenements", "", token); response.Code != http.StatusOK {
		t.Fatalf("list events status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/evenements/"+id, "", token); response.Code != http.StatusOK {
		t.Fatalf("get event status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodPut, "/evenements/"+id, `{"titre":"Atelier modifié"}`, token); response.Code != http.StatusOK {
		t.Fatalf("update event status = %d, want %d", response.Code, http.StatusOK)
	}
	if response := callAPI(router, http.MethodGet, "/evenements/99999", "", token); response.Code != http.StatusNotFound {
		t.Fatalf("missing event status = %d, want %d", response.Code, http.StatusNotFound)
	}
	if response := callAPI(router, http.MethodDelete, "/evenements/"+id, "", token); response.Code != http.StatusOK {
		t.Fatalf("delete event status = %d, want %d", response.Code, http.StatusOK)
	}
}
