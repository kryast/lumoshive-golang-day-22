package handler

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("view/*.html"))

// func List(w http.ResponseWriter, r *http.Request) {
// 	db, err := database.InitDB()

// 	if err != nil {
// 		fmt.Println("err ", err)
// 	}

// 	repo := repository.NewUserRepository(db)
// 	serviceUser := service.NewUserService(repo)

// 	users, err := serviceUser.GetAllDataUser()
// 	if err != nil {
// 		fmt.Println("err ", err)
// 	}

// 	// fmt.Println("data :", *customers)
// 	templates.ExecuteTemplate(w, "list-user-view", users)
// }

func FormRegist(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register-view", nil)
}

// func Register(w http.ResponseWriter, r *http.Request) {
// 	var name, username, password string
// 	if r.Method == http.MethodPost {
// 		name = r.FormValue("name")
// 		username = r.FormValue("username")
// 		password = r.FormValue("password")
// 	}

// 	fmt.Println("data :", name, username, password)

// 	db, err := database.InitDB()
// 	if err != nil {
// 		fmt.Println("err ", err)
// 	}

// 	userRegister := model.User{
// 		Name:     name,
// 		Username: username,
// 		Password: password,
// 	}

// 	token := uuid.New()
// 	userRegister.Status = "active"
// 	userRegister.Token = token.String()

// 	repo := repository.NewUserRepository(db)
// 	serviceUser := service.NewUserService(repo)

// 	err = serviceUser.InputDataUser(userRegister)
// 	if err != nil {
// 		fmt.Println("Error :", err)
// 	}
// 	http.Redirect(w, r, "/all-customer", http.StatusSeeOther)
// }

// func UserDetail(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")
// 	id_int, _ := strconv.Atoi(id)

// 	db, err := database.InitDB()

// 	if err != nil {
// 		fmt.Println("err ", err)
// 	}

// 	repo := repository.NewUserRepository(db)
// 	serviceUser := service.NewUserService(repo)

// 	customers, err := serviceCustomer.CustomerByID(id_int)
// 	if err != nil {
// 		fmt.Println("Error :", err)
// 	}
// 	fmt.Println(customers)
// 	templates.ExecuteTemplate(w, "user-detail-view", *customers)
// }
