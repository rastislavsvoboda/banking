package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rastislavsvoboda/banking/domain"
	"github.com/rastislavsvoboda/banking/logger"
	"github.com/rastislavsvoboda/banking/service"
)

func Start() {
	router := mux.NewRouter()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = sanityCheck()
	if err != nil {
		panic(err)
	}

	// customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerHandlers := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	full_addr := fmt.Sprintf("%s:%s", address, port)
	logger.Info("starting http://" + full_addr)
	err = http.ListenAndServe(full_addr, router)
	if err != nil {
		logger.Error(err.Error())
	}
}

func sanityCheck() error {
	env_vars := [...]string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_ADDRESS",
		"DB_PORT",
		"DB_NAME"}
	missing_env_vars := make([]string, 0)
	for _, name := range env_vars {
		value := os.Getenv(name)
		if value == "" {
			missing_env_vars = append(missing_env_vars, name)
		} else {
			logger.Debug(fmt.Sprintf("%s=%s", name, value))
		}
	}

	if len(missing_env_vars) != 0 {
		return errors.New("Missing environment variable(s): " + strings.Join(missing_env_vars, ", "))
	}

	return nil
}
