package webhttp

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(h HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: &h,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/mine").Methods("POST").HandlerFunc(s.handlers.HandleHireMiner)
	router.Path("/mine").Methods("GET").HandlerFunc(s.handlers.HandleGetWorkingMiners)
	router.Path("/mine/info").Methods("GET").HandlerFunc(s.handlers.HandleGetMineInfo)
	router.Path("/mine/prices").Methods("GET").HandlerFunc(s.handlers.HandleGetMinerPrices)
	
	router.Path("/equip").Methods("GET").HandlerFunc(s.handlers.HandleGetBoughtEquip)
	router.Path("/equip/{tool}").Methods("POST").HandlerFunc(s.handlers.HandleBuyEquipment)
	router.Path("/equip/prices").Methods("GET").HandlerFunc(s.handlers.HandleGetEquipInfo)
	
	router.Path("/mine/done").Methods("PATCH").HandlerFunc(s.handlers.HandleStopGame)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
