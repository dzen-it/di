package node

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/dzen-it/di/app/balancer"
	"github.com/dzen-it/di/app/models"
	"github.com/dzen-it/di/app/nodeStorage"
	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	address := new(models.Address)
	name := params.ByName("id")
	nodeList, err := nodeStorage.Storage.Get(name)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.NewResult("", "Internal server error"))
		return
	}
	if len(nodeList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write(models.NewResult("", "Name not found"))
		return
	}

	address.Address = balancer.GetAddressNode(r.RemoteAddr, nodeList...)
	body, _ := json.Marshal(address)

	w.Write(body)
}

func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.NewResult("", "Internal server error"))
		return
	}

	node := new(models.Node)

	err = json.Unmarshal(body, node)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewResult("", "Invalid JSON schema"))
		return
	}

	ip, err := net.ResolveTCPAddr("tcp", node.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewResult("", "Error resolving of address"))
		return
	}

	nodeStorage.Storage.Add(node.Name, ip.String())

	w.WriteHeader(http.StatusCreated)
	w.Write(models.NewResult("Address is set", ""))
}

func Put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(models.NewResult("", "Internal server error"))
		return
	}

	node := new(models.Node)

	err = json.Unmarshal(body, node)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewResult("", "Invalid JSON schema"))
		return
	}

	ip, err := net.ResolveTCPAddr("tcp", node.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.NewResult("", "Error resolving of address"))
		return
	}

	nodeStorage.Storage.Set(node.Name, ip.String())

	w.WriteHeader(http.StatusCreated)
	w.Write(models.NewResult("Address is set", ""))
}
