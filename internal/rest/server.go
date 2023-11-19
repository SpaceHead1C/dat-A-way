package rest

import (
	"dataway/internal/api"
	"dataway/internal/domain"
	"dataway/pkg/log"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	mw "github.com/go-chi/chi/v5/middleware"
)

const (
	defaultHTTPServerTimeout = time.Second * 5

	regexUUIDTemplate = `[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}`
)

type server struct {
	logger          *log.Logger
	srv             *http.Server
	errorHandler    func(error)
	timeout         time.Duration
	consumerManager *api.ConsumerManager
}

func (s *server) Serve() error {
	return s.srv.ListenAndServe()
}

type Config struct {
	Logger          *log.Logger
	Port            uint
	ErrorHandler    func(error)
	Timeout         time.Duration
	ConsumerManager *api.ConsumerManager
}

func NewServer(c Config) (domain.Server, error) {
	l := c.Logger
	if l == nil {
		l = log.GlobalLogger()
	}
	eh := c.ErrorHandler
	if eh == nil {
		eh = func(e error) {
			l.Errorln(e.Error())
		}
	}
	if c.ConsumerManager == nil {
		return nil, fmt.Errorf("consumers manager must be not nil")
	}
	if c.Timeout == 0 {
		c.Timeout = defaultHTTPServerTimeout
	}
	out := &server{
		logger:          l,
		errorHandler:    eh,
		timeout:         c.Timeout,
		consumerManager: c.ConsumerManager,
	}

	router := chi.NewRouter()
	router.Use(mw.StripSlashes)
	router.Use(mw.GetHead)
	router.Use(mw.Timeout(out.timeout))

	router.Mount("/health", healthRouter(out))
	router.Mount("/consumer", consumerRouter(out))
	out.srv = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		WriteTimeout: time.Second * 7,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}
	return out, nil
}

func healthRouter(s *server) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/ping", newPingHandler(s))
	return r
}

func consumerRouter(s *server) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", newAddConsumerHandler(s))
	r.Put(fmt.Sprintf("/{id:%s}", regexUUIDTemplate), newUpdConsumerHandler(s))
	r.Patch(fmt.Sprintf("/{id:%s}", regexUUIDTemplate), newPatchConsumerHandler(s))
	r.Get(fmt.Sprintf("/{id:%s}", regexUUIDTemplate), newGetConsumerHandler(s))
	return r
}

func (s *server) emptyResp(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (s *server) textResp(w http.ResponseWriter, status int, payload string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	if err := writeResp(w, status, []byte(payload)); err != nil {
		s.errorHandler(err)
	}
}

func (s *server) jsonResp(w http.ResponseWriter, status int, payload []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := writeResp(w, status, payload); err != nil {
		s.errorHandler(err)
	}
}

func writeResp(w http.ResponseWriter, status int, payload []byte) error {
	w.WriteHeader(status)
	if _, err := w.Write(payload); err != nil {
		return err
	}
	return nil
}
