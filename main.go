package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Persona struct {
	HogarId      string  `json:"hogarId"`
	Departamento string  `json:"departamento"`
	Provincia    string  `json:"provincia"`
	Distrito     string  `json:"distrito"`
	Area         string  `json:"area"`
	PersonaId    string  `json:"personaId"`
	Genero       string  `json:"genero"`
	Edad         int     `json:"edad"`
	Aprobado     bool    `json:"aprobado"`
	Distancia    float64 `json:"distancia"`
}

type Personas []Persona

var personas = Personas{
	{
		HogarId:      "30546846",
		Departamento: "LIMA",
		Provincia:    "LIMA",
		Distrito:     "LOS OLIVOS",
		Area:         "Urbano",
		PersonaId:    "75863340",
		Genero:       "Masculino",
		Edad:         15,
		Aprobado:     true,
	},
}

var dataset = Personas{}

var minList = Personas{}

func getAllPersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Endpoint Hit: GetAll Personas Endpoint")
	json.NewEncoder(w).Encode(personas)
}

func getByIdPersona(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	for _, persona := range personas {
		if persona.HogarId == vars["id"] {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(persona)
		}
	}

	fmt.Println("Endpoint Hit: GetById a Persona Endpoint")
}

func createPersona(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	var newPersona Persona
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Persona")
	}

	json.Unmarshal(reqBody, &newPersona)

	personas = append(personas, newPersona)

	json.NewEncoder(w).Encode(newPersona)

	fmt.Println("Endpoint Hit: Create a Personas Endpoint")
}

func deletePersona(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	for i, persona := range personas {
		if persona.HogarId == vars["id"] {
			personas = append(personas[:i], personas[i+1:]...)
			fmt.Fprintf(w, "El hogar con el ID %v ha sido removido exitosamente", vars["id"])
		}
	}

	fmt.Println("Endpoint Hit: Delete a Persona Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Runnig API GO...")
}

func getAllDataSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Endpoint Hit: GetAll Dataset Endpoint")
	json.NewEncoder(w).Encode(dataset)
}

func loadPersonas(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("dataset.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var dataSetPersona Persona

		line := scanner.Text()
		items := strings.Split(line, "|")

		dataSetPersona.HogarId = items[0]
		dataSetPersona.Departamento = items[1]
		dataSetPersona.Provincia = items[2]
		dataSetPersona.Distrito = items[3]
		if items[4] == "1" {
			dataSetPersona.Area = "Urbano"
		} else {
			dataSetPersona.Area = "Rural"
		}
		dataSetPersona.PersonaId = items[5]
		if items[6] == "1" {
			dataSetPersona.Genero = "Masculino"
		} else {
			dataSetPersona.Genero = "Femenino"
		}
		dataSetPersona.Edad, err = strconv.Atoi(items[7])
		if items[8] == "1" {
			dataSetPersona.Aprobado = true
		} else {
			dataSetPersona.Aprobado = false
		}

		dataset = append(dataset, dataSetPersona)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("Endpoint Hit: Load Personas DataSet Endpoint")
}

func loadKnnPersona(w http.ResponseWriter, r *http.Request) {

	dataset = nil
	minList = nil

	var newPersona Persona

	var areaP int
	var areaDS int

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Persona")
	}

	json.Unmarshal(reqBody, &newPersona)

	file, err := os.Open("dataset.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var dataSetPersona Persona

		line := scanner.Text()
		items := strings.Split(line, "|")

		dataSetPersona.HogarId = items[0]
		dataSetPersona.Departamento = items[1]
		dataSetPersona.Provincia = items[2]
		dataSetPersona.Distrito = items[3]
		if items[4] == "1" {
			dataSetPersona.Area = "Urbano"
			areaDS = 1
			areaP = 1
		} else {
			dataSetPersona.Area = "Rural"
			areaDS = 2
			areaP = 2
		}
		dataSetPersona.PersonaId = items[5]
		if items[6] == "1" {
			dataSetPersona.Genero = "Masculino"
		} else {
			dataSetPersona.Genero = "Femenino"
		}
		dataSetPersona.Edad, err = strconv.Atoi(items[7])
		if items[8] == "1" {
			dataSetPersona.Aprobado = true
		} else {
			dataSetPersona.Aprobado = false
		}

		dataSetPersona.Distancia = math.Sqrt(float64((newPersona.Edad-dataSetPersona.Edad)*(newPersona.Edad-dataSetPersona.Edad) + (areaP-areaDS)*(areaP-areaDS)))

		dataset = append(dataset, dataSetPersona)
	}

	var personaMin Persona

	for i := 0; i < 3; i++ {
		personaMin = dataset[0]
		go EncontrarMinimo(personaMin)
	}
	time.Sleep(2 * time.Second)

	aprobado := 0
	noAprobado := 0

	for _, min := range minList {
		if min.Aprobado == true {
			aprobado++
		} else {
			noAprobado++
		}
	}

	if aprobado > noAprobado {
		newPersona.Aprobado = true
	} else {
		newPersona.Aprobado = false
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	personas = append(personas, newPersona)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPersona)

	fmt.Println("Endpoint Hit: Load KNN Persona DataSet Endpoint")
}

func EncontrarMinimo(personaMin Persona) {
	personaMin = dataset[0]
	for _, v := range dataset {
		if v.Distancia < personaMin.Distancia {
			personaMin = v
		}
	}
	for j, persona := range dataset {
		if persona.HogarId == personaMin.HogarId {
			dataset = append(dataset[:j], dataset[j+1:]...)
		}
	}
	minList = append(minList, personaMin)
	fmt.Println(personaMin)
}

func getAllMin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Endpoint Hit: GetAll Min Endpoint")
	json.NewEncoder(w).Encode(minList)
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/personas", getAllPersona).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/personas/{id}", getByIdPersona).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/persona-crear", createPersona).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/personas/{id}", deletePersona).Methods("DELETE", "OPTIONS")
	myRouter.HandleFunc("/dataset", getAllDataSet).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/cargar", loadPersonas).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/knn-personas", loadKnnPersona).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/minimos", getAllMin).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func main() {
	handleRequest()
}
