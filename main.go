package main

import (
	"log"
	"net/url"
	"net/http"
	"fmt"
	"./config"
	"./reverseproxy"
	"./tokenbucket"
	"./jsonconv"
)

var inspector reverseproxy.RequestInspector
var tokenBucket tokenbucket.TokenBucket
var transporter reverseproxy.Transporter
var jsonQueryFactory jsonconv.JsonQueryFactory

func main() {
	log.Print("Init logger")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Load configuration")
	conf, err := config.LoadConfig("conf.json")
	if err != nil {
		log.Println("Unable to load config file")
		log.Println(err.Error())
		return
	}

	// setup dependencies in vars declared on top of this file
	setupDependencies(conf)

	log.Print("Init routes")
	err = initRoutes(conf.Routes, transporter)
	if err != nil {
		log.Println("Unable to initiate routes")
		log.Println(err.Error())
	}

	Address := fmt.Sprintf("%s:%s", conf.IP, conf.Port)
	log.Print("Start listening on " + Address)
	http.ListenAndServe(Address, nil)
}

func setupDependencies(conf *config.Config) {
	inspector = reverseproxy.NewAllowQueryTypesInspector(conf.AllowedQueryParameters)

	clock := tokenbucket.NewClock()
	metric := tokenbucket.NewTimeMetric()
	tokenBucket = tokenbucket.NewStandardTokenBucket(clock, metric, conf.TimePerToken, conf.BurstTokens)

	transporter = reverseproxy.NewHttpTransport()

	jsonQueryFactory = jsonconv.NewDefaultJsonConverterFactory()
}

func initRoutes(routes map[string]config.Target, transporter reverseproxy.Transporter) error {
	for hostPath, target := range routes {
		targetUrl, err := url.Parse(target.Host)
		if err != nil {
			return err
		}

		conv := jsonQueryFactory.Get(target.JQuery)
		if err != nil {
			return err
		}

		// setup a handler for specified target...
		responseModifier := reverseproxy.NewJsonConvResponseModifier(conv)
		director := reverseproxy.NewTargetDirector(targetUrl)
		p := reverseproxy.NewReverseProxy(director, responseModifier, transporter)
		handler := reverseproxy.NewHttpHandler(p, inspector, tokenBucket)

		http.HandleFunc(hostPath, handler.Get())

		log.Printf("%s > %s", hostPath, target)
	}

	return nil
}
