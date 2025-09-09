package handlers

import "net/http"

// const getStageTag = "api:GetStage"

func GetStage(w http.ResponseWriter, r *http.Request) {
	// var stage models.Stage
	// result := db.Database.Where("id = ?", r.Context().Value(middleware.StageID)).Take(&stage)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	http.Error(w, l10n.StageNotFound, http.StatusNotFound)
	// 	return
	// }
	// // response := models.StageResponse{Stage: &stage, CreatedAt: stage.CreatedAt.String()}
	// // if err := jsonEncode(w, response, "public"); err != nil {
	// // 	util.Error(getStageTag, err)
	// // 	return
	// // }
	// w.WriteHeader(http.StatusOK)
	// util.Info(getStageTag, "id%d returned", stage.ID)
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateStage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
