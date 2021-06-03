package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aborilov/tech-challenge-time/backend/v1/pkg/session/model"
)

// HTTPController holds the state and HTTP handlers for the ResourceQuery API
type HTTPController struct {
	svc model.Service
}

// NewHTTPController creates a new HTTPController.
func NewHTTPController(svc model.Service, router *mux.Router) {
	c := &HTTPController{svc: svc}
	// swagger:operation GET /session/ listSessins
	// ---
	// summary: Return list of session
	// description: return list of session
	// parameters:
	// - in: query
	//   name: start_after
	//   type: date-time
	// - name: end_before
	//   in: query
	//   type: date-time
	// responses:
	//   "200":
	//     "$ref": "#/responses/sessions"
	//   "503":
	//       description: unable to get list of sessions or marshal output
	//       "schema":
	//         "type": "string"
	router.Path("/").Methods("GET").HandlerFunc(c.handleList)

	// swagger:operation POST /session/start startSession
	// ---
	// summary: Start new tracking session
	// description: Start new tracking session
	// responses:
	//   "200":
	//     "$ref": "#/responses/session"
	//   "503":
	//       description: unable to start new session or marshal output
	//       "schema":
	//         "type": "string"
	//   "400":
	//       description: session already in progress
	//       "schema":
	//         "type": "string"
	router.Path("/start").Methods("POST").HandlerFunc(c.handleStart)

	// swagger:operation POST /session/stop stopSession
	// ---
	// summary: Stop in progress session
	// description: Start new tracking session
	// responses:
	//   "200":
	//     "$ref": "#/responses/session"
	//   "503":
	//       description: unable to stop session or marshal output
	//       "schema":
	//         "type": "string"
	//   "400":
	//       description: there is no session in progress
	//       "schema":
	//         "type": "string"
	router.Path("/stop").Methods("POST").HandlerFunc(c.handleStop)

	// swagger:operation PUT /session/{id}?{name}=work setSessionName
	// ---
	// summary: set name of the specified session
	// description: Start new tracking session
	// parameters:
	// - in: query
	//   name: name
	//   type: string
	//   description: name of the session
	//   required: true
	// - name: id
	//   in: path
	//   description: id of the session
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/session"
	//   "503":
	//       description: unable to update session or marshal output
	//       "schema":
	//         "type": "string"
	//   "400":
	//       "schema":
	//         "type": "string"
	router.Path("/{id}").Methods("PUT").HandlerFunc(c.handleSetName)
}

func (c *HTTPController) handleSetName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "session id required", http.StatusBadRequest)
		return
	}
	name, ok := r.URL.Query()["name"]
	if !ok || len(name) == 0 {
		http.Error(w, "name param required", http.StatusBadRequest)
		return
	}
	session, err := c.svc.Get(ctx, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get session: %s", err), http.StatusInternalServerError)
		return
	}
	session.Name = name[0]
	session, err = c.svc.Update(ctx, session)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to set session name: %s", err), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(fromSession(session))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal: %s", err), http.StatusInternalServerError)
		return
	}
}

func (c *HTTPController) handleStop(w http.ResponseWriter, r *http.Request) {
	session, err := c.svc.StopCurrentSession(r.Context())
	if err != nil {
		if errors.As(err, &model.ErrNoSessionInProgress{}) {
			http.Error(w, fmt.Sprintf("unable to stop session: %s", err), http.StatusBadRequest)
			return
		}
		http.Error(w, fmt.Sprintf("unable to stop session: %s", err), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(fromSession(session))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal: %s", err), http.StatusInternalServerError)
		return
	}
}

func (c *HTTPController) handleStart(w http.ResponseWriter, r *http.Request) {
	session, err := c.svc.StartNewSession(r.Context())
	if err != nil {
		if errors.As(err, &model.ErrSessionInProgress{}) {
			http.Error(w, fmt.Sprintf("unable to start new session: %s", err), http.StatusBadRequest)
			return
		}
		http.Error(w, fmt.Sprintf("unable to start new session: %s", err), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(fromSession(session))
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal: %s", err), http.StatusInternalServerError)
		return
	}
}

func (c *HTTPController) handleList(w http.ResponseWriter, r *http.Request) {
	f, err := parseFilter(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !f.StartAfter.IsZero() && !f.EndBefore.IsZero() {
		if f.StartAfter.After(f.EndBefore) {
			http.Error(w, "start_after should be before end_before", http.StatusBadRequest)
			return
		}
	}
	sessions, err := c.svc.List(r.Context(), f)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to get sessions: %s", err), http.StatusInternalServerError)
		return
	}
	ss := make([]*session, 0, len(sessions))
	for _, session := range sessions {
		ss = append(ss, fromSession(session))
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(ss)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to marshal: %s", err), http.StatusInternalServerError)
		return
	}
}
