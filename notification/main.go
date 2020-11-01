package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/personal_proj/notification/app"
	"github.com/personal_proj/notification/kafka"

	"github.com/rs/zerolog/log"
)

var (
	listenAddrApi  string
	kafkaBrokerUrl string
	kafkaVerbose   bool
	kafkaClientId  string
	kafkaTopic     string
	logger         = log.With().Str("pkg", "main").Logger()
)

func main() {
	configureProducer()
	app.StartApplication()
}

func configureProducer() {
	flag.StringVar(&listenAddrApi, "listen-address", "0.0.0.0:9000", "Listen address for api")
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", "localhost:9092", "Kafka brokers in comma separated value")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "my-kafka-client", "Kafka client id to connect")
	flag.StringVar(&kafkaTopic, "kafka-topic", "notification", "Kafka topic to push")

	flag.Parse()

	// connect to kafka
	kafkaProducer, err := kafka.Configure(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
	if err != nil {
		logger.Error().Str("error", err.Error()).Msg("unable to configure kafka")
		return
	}
	defer kafkaProducer.Close()
	var errChan = make(chan error, 1)

	go func() {
		log.Info().Msgf("starting server at %s", listenAddrApi)
		//	errChan <- server(listenAddrApi)
		app.StartApplication()
	}()

	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-signalChan:
		logger.Info().Msg("got an interrupt, exiting...")
	case err := <-errChan:
		if err != nil {
			logger.Error().Err(err).Msg("error while running api, exiting...")
		}
	}
}
