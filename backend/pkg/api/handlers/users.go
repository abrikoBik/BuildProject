package handlers

import (
	"errors"
	"mp1/pkg/api/middleware"
	"mp1/pkg/db"
	"mp1/pkg/l10n"
	"mp1/pkg/models"
	"mp1/pkg/util"
	"net/http"
	"sort"

	"gorm.io/gorm"
)

var (
	lGetUser      = util.CreateLogger("api:GetUser")
	lCreateUser   = util.CreateLogger("api:CreateUser")
	lPatchUser    = util.CreateLogger("api:PatchUser")
	lListProjects = util.CreateLogger("api:ListProjects")
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lGetUser.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	result := user.Get(db.Database, userID)
	if errors.Is(result, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	response := models.UserResponse{User: &user, CreatedAt: user.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		lGetUser.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	lGetUser.Info("id%d returned", user.ID)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var request models.UserRequest
	if !jsonDecode(w, r, &request) {
		return
	}
	if !errors.Is(db.Database.Where(
		"login = ? OR email = ? OR phone_number = ?",
		request.Name,
		request.Email,
		request.PhoneNumber,
	).Take(&models.User{}).Error, gorm.ErrRecordNotFound) {
		http.Error(w, l10n.FieldTaken, http.StatusBadRequest)
		return
	}
	var user models.User
	if err := user.Create(db.Database, request); err != nil {
		lCreateUser.Error(err)
		switch err {
		case l10n.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case l10n.InvalidEmailError:
			http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		default:
			http.Error(w, l10n.RegisterError, http.StatusInternalServerError)
		}
		return
	}
	lCreateUser.Info("%s (id%d) registered", user.Name, user.ID)
	response := models.UserResponse{ID: user.ID}
	if err := jsonEncode(w, response, "register"); err != nil {
		lCreateUser.Error(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lPatchUser.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	if user.Get(db.Database, id) != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	if !verifyUser(user, w, r) {
		return
	}
	request := models.UserRequest{User: &models.User{}}
	if !jsonDecode(w, r, &request) {
		return
	}
	if err := user.Patch(db.Database, request); err != nil {
		lPatchUser.Error(err)
		switch err {
		case l10n.InsecurePasswordError:
			http.Error(w, l10n.InsecurePassword, http.StatusBadRequest)
		case l10n.InvalidEmailError:
			http.Error(w, l10n.InvalidEmail, http.StatusBadRequest)
		case gorm.ErrDuplicatedKey:
			http.Error(w, l10n.FieldTaken, http.StatusBadRequest)
		default:
			http.Error(w, l10n.GenericUser, http.StatusInternalServerError)
		}
		return
	}
	lPatchUser.Info("%s (id%d) patched", user.Name, user.ID)
	response := models.UserResponse{User: &user, CreatedAt: user.CreatedAt.String()}
	if err := jsonEncode(w, response, "public"); err != nil {
		lPatchUser.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListProjects(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.GetID(middleware.UserID, r.Context())
	if err != nil {
		lListProjects.Error(err)
		http.Error(w, l10n.ContextID, http.StatusInternalServerError)
		return
	}
	var user models.User
	if user.Get(db.Database, id, "Projects") != nil {
		http.Error(w, l10n.UserNotFound, http.StatusNotFound)
		return
	}
	sort.Slice(user.Projects, func(i, j int) bool {
		return user.Projects[i].UpdatedAt.UnixMilli() > user.Projects[j].UpdatedAt.UnixMilli()
	})
	var response []models.ProjectResponse
	for _, p := range user.Projects {
		toAppend := models.ProjectResponse{
			Project:   &p,
			StartTime: util.MarshalTime(p.StartTime),
			EndTime:   util.MarshalTime(p.EndTime),
			ID:        p.ID,
		}
		response = append(response, toAppend)
	}
	if err := jsonEncode(w, response, "list"); err != nil {
		lListProjects.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
