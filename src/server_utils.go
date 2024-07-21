package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	//"github.com/davecgh/go-spew/spew"
)

type Message struct {
	Data string `json:"data"`
}

// function to start the http server
func Run() error {
	mux := makeMuxRouter()
	httpAddr := 3000
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           fmt.Sprintf(": %d", httpAddr),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// registering routes
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bc := GetBlockchainInstance()
	fmt.Println("Length of GET blockchain", len(bc.Blocks))
	bytes, err := json.MarshalIndent(bc.Blocks, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		fmt.Println(err)
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	// release the resources
	defer r.Body.Close()

	bc := GetBlockchainInstance()
	oldBlock, err := bc.GetLastBlock()
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	newBlock := NewBlock(*oldBlock, m.Data)

	if isBlockValid(*newBlock, *oldBlock) {
		bc.AddBlock(newBlock)
	}
	fmt.Println("Adding a new block...")
	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
