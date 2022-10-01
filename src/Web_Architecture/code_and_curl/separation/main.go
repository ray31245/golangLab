package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
)

// Action Layer
var ErrUserNotFound = errors.New("User not found")

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserStore interface {
	// Get may return ErrUserNotFound error
	Get(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

type MemoryUserStorage struct {
	store map[string]*User
}

func NewMemoryUserStorage() *MemoryUserStorage {
	return &MemoryUserStorage{
		store: map[string]*User{},
	}
}

func (ms *MemoryUserStorage) Get(ctx context.Context, email string) (*User, error) {
	if u, ok := ms.store[email]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}

func (ms *MemoryUserStorage) Save(ctx context.Context, user *User) error {
	ms.store[user.Email] = user
	return nil
}

// Business Logic
type RegistParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (rp *RegistParams) Validate() error {
	if rp.Email == "" {
		return errors.New("Email cannot be empty")
	}

	if !strings.ContainsRune(rp.Email, '@') {
		return errors.New("Email must include an '@' symbol")
	}

	if rp.Name == "" {
		return errors.New("Name cannot be empty")
	}

	return nil
}

type UserService interface {
	// Register may return an ErrEmailExists error
	Register(context.Context, *RegistParams) error
	// GetByEmail may return an ErrUserNotFound error
	GetByEmail(context.Context, string) (*User, error)
}

var ErrEmailExists = errors.New("Email is already in use")

type UserServiceImpl struct {
	userStorage UserStore
}

func NewUserServiceImpl(us UserStore) *UserServiceImpl {
	return &UserServiceImpl{
		userStorage: us,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, Params *RegistParams) error {
	_, err := us.userStorage.Get(ctx, Params.Email)
	if err == nil {
		return ErrEmailExists
	} else if err != ErrUserNotFound {
		return err
	}

	return us.userStorage.Save(ctx, &User{
		Email: Params.Email,
		Name:  Params.Name,
	})
}

func (us *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	return us.userStorage.Get(ctx, email)
}

// Access Layer
type JSONOverHTTP struct {
	router  *http.ServeMux
	usrServ UserService
}

func NewJsonOverHTTP(usrServ UserService) *JSONOverHTTP {
	r := http.NewServeMux()
	joh := &JSONOverHTTP{
		router:  r,
		usrServ: usrServ,
	}
	r.HandleFunc("/register", joh.Register)
	r.HandleFunc("/user", joh.GetUser)
	return joh
}

func (j *JSONOverHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	j.router.ServeHTTP(w, r)
}

func (j *JSONOverHTTP) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Register require a post request", http.StatusBadRequest)
		return
	}

	params := &RegistParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.Error(w, "Unable to read your request", http.StatusBadRequest)
		return
	}

	err = params.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = j.usrServ.Register(r.Context(), params)
	if err == ErrEmailExists {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (j *JSONOverHTTP) validateEmail(email string) error {
	if email == "" {
		return errors.New("Email must not be empty")
	}

	if !strings.ContainsRune(email, '@') {
		return errors.New("Email must include an '@' symbol")
	}

	return nil
}

func (j *JSONOverHTTP) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Error(w, "GetUser requires a get request", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	err := j.validateEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := j.usrServ.GetByEmail(r.Context(), email)
	if err == ErrUserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Write together
func main() {
	usrStore := NewMemoryUserStorage()
	usrServ := NewUserServiceImpl(usrStore)
	joh := NewJsonOverHTTP(usrServ)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, joh)
	if err != nil {
		panic(err)
	}
}
