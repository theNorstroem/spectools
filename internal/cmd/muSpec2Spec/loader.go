package muSpec2Spec

import (
	"github.com/theNorstroem/spectools/pkg/microservices"
	"github.com/theNorstroem/spectools/pkg/microtypes"
	"io/ioutil"
	"log"
)

func LoadServices(list []string, bigList *microservices.MicroServiceList) {

	for _, muSpecFile := range list {
		dataBytes, readError := ioutil.ReadFile(muSpecFile)
		if readError != nil {
			log.Fatal(readError)
		}
		microServicesList := microservices.MicroServiceList{}
		microServicesList.Unmarshal(dataBytes)
		// add types to global list
		for _, mt := range microServicesList.MicroServices {
			bigList.MicroServices = append(bigList.MicroServices, mt)
		}
	}

}

func LoadTypes(list []string, bigList *microtypes.MicroTypelist) {

	for _, muSpecFile := range list {
		dataBytes, readError := ioutil.ReadFile(muSpecFile)
		if readError != nil {
			log.Fatal(readError)
		}
		microList := microtypes.MicroTypelist{}
		microList.Unmarshal(dataBytes)
		// add types to global list
		for _, mt := range microList.MicroTypes {
			mt.SourceFile = muSpecFile
			bigList.MicroTypes = append(bigList.MicroTypes, mt)
		}
	}

}
