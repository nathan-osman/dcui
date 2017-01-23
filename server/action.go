package server

import (
	"net/http"
)

// action performs the specified action.
func (s *Server) action(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var (
		action  = r.Form.Get("action")
		service = r.Form.Get("service")
	)
	if err := s.docker.Do(action, service); err != nil {
		addAlert(alertDanger, err.Error())
	} else {
		addAlert(alertInfo, "operation completed successfully")
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
