package main

import (
	"goFinalTask/configure"
	"goFinalTask/handlers"
	modeldb "goFinalTask/modelDB"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	httpPort = "8089"
	//httpPort = "8080"
	//httpPort = "8081"

)

var log = logrus.New()

func init() {

	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	//log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.DebugLevel
	//log.Level = logrus.InfoLevel
	log.Out = os.Stdout
}

func main() {
	log.Println("try to start REST API server")
	configPath := "./configure/configs.yaml"
	main_config, err := configure.GetConfigs(configPath)
	if err != nil {
		log.Fatalf("Failed to get configs: %s", err)
	}

	os.Setenv("CURRENCY_API_KEY", main_config.ServiceAPIKey)
	log.Printf("Set ENV CURRENCY_API_KEY: %+v\n", main_config.ServiceAPIKey)

	os.Setenv("BASE_CURRENCY", main_config.BaseCurrency)
	log.Printf("Set ENV BASE_CURRENCY: %+v\n", main_config.BaseCurrency)

	// Подключение к БД
	log.Printf("Database DSN Config: %+v\n", main_config.DSN)
	modeldb.ConnectDB(main_config.DSN)

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/transactions", handlers.TransactionsHandler).Methods("GET", "POST")
	r.HandleFunc("/transactions/{id}", handlers.TransactionsHandler).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/transactions/{id}/{base}", handlers.TransactionsHandler).Methods("GET")

	r.HandleFunc("/users", handlers.HandleUsers).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", handlers.HandleUsers).Methods("DELETE") // Для PUT и DELETE

	http.Handle("/", r)

	log.Println("Starting server on port " + httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Request started")
		next.ServeHTTP(w, r)
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("Request finished")
	})
}
