package handlers

import "net/http"

// TODO: rewrite
func Users(w http.ResponseWriter, r *http.Request) {
	// sanity checks
	// var user models.User
	// result := db.Database.Where("id = ?", 0).Take(&user)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// var user models.User
	// result = db.Database.Where("sessions IN ?", dbSession.Token).Take(&user)
	// if result.Error != nil {
	// 	util.Log("api", fmt.Sprintf("/userInfo: %s", result.Error.Error()))
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("error finding user"))
	// 	return
	// }
	// var toMarshal models.UserResponse
	// marshalled, err := json.Marshal(toMarshal)
	// if err != nil {
	// 	util.Log("api", fmt.Sprintf("/userInfo: %s", err))
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("error marshalling user"))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(marshalled)
}
