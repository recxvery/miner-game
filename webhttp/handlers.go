package webhttp

import (
	"encoding/json"
	"errors"
	"log"
	"miner-game/miner"
	"miner-game/miner/equipment"
	"miner-game/webhttp/input"
	"miner-game/webhttp/output"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	mine *miner.Mine
}

func NewHTTPHandlers(newMine *miner.Mine) *HTTPHandlers {
	return &HTTPHandlers{
		mine: newMine,
	}
}

/*
pattern: /mine
Method: GET
info: pattern & json body request
succeed:

	-status code 200(!OK)
	-response body:

failed:

	-status code: 404, 400, 500...
	-response body: json err message + time
*/
func (h *HTTPHandlers) HandleGetMineInfo(w http.ResponseWriter, r *http.Request) {
	mineInfo := h.mine.GetMineInfo()

	jsonData, err := json.MarshalIndent(mineInfo, "", "		")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Println(err)

		return
	}
}

/*
pattern: /mine
Method: POST
info: pattern & json body request
succeed:

	-status code 201(created)
	-response body: json body + status ok

failed:

	-status code: 500, 403
	-response body: json err message + time
*/
func (h *HTTPHandlers) HandleHireMiner(w http.ResponseWriter, r *http.Request) {
	var minerdto input.MinerDTO
	if err := json.NewDecoder(r.Body).Decode(&minerdto); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		errorHandle(w, err)
		return
	}

	if err := minerdto.ValidateToCreate(); err != nil {
		errorHandle(w, err)
		return
	}
	newminer, err := h.mine.HireMiner(minerdto.Class)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)

		errorHandle(w, err)
		return
	}
	data, err := json.MarshalIndent(newminer, "", "		")
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}

/*
pattern: /mine?working=&class=
Method: Get
info: pattern & json body request
succeed:

	-status code 201(created)
	-response body: json bodies list by filter

failed:

	-status code: 500, 400
	-response body: json err body + time
*/

func (h *HTTPHandlers) HandleGetWorkingMiners(w http.ResponseWriter, r *http.Request) {
	var (
		miners       map[int64]miner.MinerInfo
		outputMiners map[int64]miner.MinerInfo
	)

	query := r.URL.Query()
	classParamStr := query.Get("class")

	if classParamStr != "strong" && classParamStr != "normal" && classParamStr != "little" {
		classParamStr = ""
	}

	if classParamStr != "" {
		miners = h.mine.GetMinersWorkingClass(classParamStr)
		outputMiners = output.MinersDTObyType(miners)
	}

	jsonData, err := json.MarshalIndent(outputMiners, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Println(err)

		return
	}
}

/*
Pattern: /mine/{tool}
Method: PATCH
info: pattern + json body
succeed:

	-status code: 200(OK)
	-json changed body + time

failed:

	-status code: 500, 400, 404, 409
	-response: json err body + time
*/
func (h *HTTPHandlers) HandleBuyEquipment(w http.ResponseWriter, r *http.Request) {
	log.Println("handle working")
	tool := mux.Vars(r)["tool"]

	changedTool, err := h.mine.BuyEquipment(tool)

	if err != nil {
		errorHandle(w, err)
		return
	}

	jsonData, err := json.MarshalIndent(changedTool, "", "		")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Println(err)

		return
	}
}

/*
pattern: /mine?bought=
Method: Get
info: pattern & json body request
succeed:

	-status code 201(created)
	-response body: json bodies list by filter

failed:

	-status code: 500, 400
	-response body: json err body + time
*/

func (h *HTTPHandlers) HandleGetBoughtEquip(w http.ResponseWriter, r *http.Request) {
	var (
		equipment map[string]equipment.Equip
	)

	equipment = h.mine.GetBoughtEquipmentList()

	jsonData, err := json.MarshalIndent(equipment, "", "	")
	if err != nil {
		log.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonData); err != nil {
		log.Println(err)

		return
	}
}

/*
pattern: /mine/equip
Method: GET
info: pattern & json body request
succeed:

	-status code 200(OK)
	-response body: json bodies list by filter

failed:

	-status code: 500, 400
	-response body: json err body + time
*/
func (h *HTTPHandlers) HandleGetEquipInfo(w http.ResponseWriter, r *http.Request) {
	EquipInfo := h.mine.GetEquipmentList()

	data, err := json.MarshalIndent(EquipInfo, "", "	")
	if err != nil {
		errorHandle(w, err)

		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println(err)

		return
	}
}

/*
pattern: /mine
Method: PATCH
info: pattern & json body request
succeed:

	-status code 200(OK)
	-response body: json bodies list by filter

failed:

	-status code: 500, 400
	-response body: json stats body + time
*/
func (h *HTTPHandlers) HandleStopGame(w http.ResponseWriter, r *http.Request) {
	err := h.mine.StopGame()

	if err != nil {
		errorHandle(w, err)

		return
	}

	game_duration := time.Since(h.mine.RunningTime)
	mine := output.MineDTO{
		MinersCount:   h.mine.MinersCount,
		Miners:        output.NewAllMinerDTO(h.mine),
		RunningTime:   game_duration.String(),
		Equipmentlist: output.NewEquipmentDTO(h.mine),
		CoalSum:       h.mine.CoalSum,
	}

	data, err := json.MarshalIndent(mine, "", "		")
	if err != nil {
		errorHandle(w, err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println(err)

		return
	}
}

func errorHandle(w http.ResponseWriter, err error) {
	errDTO := output.ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	switch {
	case errors.Is(err, miner.ErrNotEnoughBalance):
		http.Error(w, errDTO.ErrToJson(), http.StatusForbidden)
	case errors.Is(err, strconv.ErrSyntax):
		http.Error(w, errDTO.ErrToJson(), http.StatusBadRequest)
	case errors.Is(err, strconv.ErrRange):
		http.Error(w, errDTO.ErrToJson(), http.StatusBadRequest)
	default:
		http.Error(w, errDTO.ErrToJson(), http.StatusBadRequest)
	}
}
