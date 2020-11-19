package main

import (
	"encoding/json"
	"fmt"
	"github.com/TimothyStiles/poly"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to MVP optimization homepage")
	fmt.Println("Endpoint Hit: homePage")
}

var optimizationTable poly.CodonTable = poly.ReadGbk("data/ecoli.gb").GetOptimizationTable(poly.GetCodonTable(11))

type AminoAcidSequenceToOptimize struct {
	Sequence string
}

func createNewOptimization(w http.ResponseWriter, r *http.Request) {
	//optimizationTable := poly.ReadGbk("data/ecoli.gb").GetOptimizationTable(poly.GetCodonTable(11))
	reqBody, _ := ioutil.ReadAll(r.Body)
	var s AminoAcidSequenceToOptimize
	json.Unmarshal(reqBody, &s)
	newSeq := poly.Optimize(s.Sequence, optimizationTable)
	json.NewEncoder(w).Encode(newSeq)
	fmt.Println("Endpoint Hit: /optimize")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/optimize", createNewOptimization)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {
	fmt.Println("Starting up")
	handleRequests()

}
