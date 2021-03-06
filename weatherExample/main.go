package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type providerDirectory struct {
	temperatura        string
	humedad            string
	probabilidadLluvia string
}

func main() {
	start := time.Now()

	baseURL := "http://localhost:8082/"

	directory := providerDirectory{
		temperatura:        baseURL + "temp",
		humedad:            baseURL + "hum",
		probabilidadLluvia: baseURL + "rc",
	}

	respChan := make(chan string)

	go getData(directory.temperatura, respChan)
	go getData(directory.humedad, respChan)
	go getData(directory.probabilidadLluvia, respChan)

	reporteFinal := construirReporte(respChan)
	fmt.Println(reporteFinal)
	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Tiempo para elaborar el informe: %s", elapsed))
}

func construirReporte(c chan string) string {
	reporteFinal := <-c
	reporteFinal += <-c
	reporteFinal += <-c
	return reporteFinal
}

func getData(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		c <- link + ": Sin datos"
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	c <- bodyString
}
