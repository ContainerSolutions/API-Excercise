package endpoint

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.com/DGuedes/API-Exercise/endpoint/apiv1"
	"gitlab.com/DGuedes/API-Exercise/storage/postgres"
)

// Server is an abstraction for the server that will host the titanic API.
type Server struct {
	Router     *mux.Router
	DAOFactory *postgres.DAOFactory
	APIImpl    *apiv1.APIImpl
	SvrCfg     *ServerConfig
}

// BuildSvr constructs a server that serve http requests using the given postgres configuration.
func BuildSvr(svrCfg ServerConfig) *Server {
	return &Server{
		Router:     mux.NewRouter(),
		DAOFactory: postgres.BuildDAOFactory(&svrCfg.PostgresConfig),
		SvrCfg:     &svrCfg,
	}
}

// Serve runs the server and listens on the configured port.
func (s *Server) Serve() {
	for {
		err := s.DAOFactory.Connect() // Connect to Postgres database.
		if err == nil {
			break
		}
		logrus.Errorf("Couldn't connect to Postgres, retrying in 10 seconds.")
		time.Sleep(10 * time.Second)
	}
	s.InitializeRoutes()
	http.ListenAndServe(":8080", s.Router)
}

// InitializeRoutes apply REST People routes to the server.
func (s *Server) InitializeRoutes() {
	personDAO := s.DAOFactory.NewPersonDAO()

	s.APIImpl = &apiv1.APIImpl{PersonDAO: personDAO, TemplateFolder: (*s.SvrCfg).TemplatePath}
	uuidRgx := "[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}"
	s.Router.HandleFunc("/people", s.APIImpl.GetPeople).Methods("GET")
	s.Router.HandleFunc("/people", s.APIImpl.CreatePerson).Methods("POST")
	s.Router.HandleFunc("/people/{uuid:"+uuidRgx+"}", s.APIImpl.GetPerson).Methods("GET")
	s.Router.HandleFunc("/people/{uuid:"+uuidRgx+"}", s.APIImpl.EditPerson).Methods("PUT")
	s.Router.HandleFunc("/people/{uuid:"+uuidRgx+"}", s.APIImpl.DeletePerson).Methods("DELETE")
}
