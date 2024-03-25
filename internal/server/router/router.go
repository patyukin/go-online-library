package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patyukin/go-online-library/internal/handler"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	incomingTraffic = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "incoming_traffic",
		Help: "Incoming traffic to the application",
	})

	outgoingTraffic = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "outgoing_traffic",
		Help: "Outgoing traffic from the application",
	})
)

func Init(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	prometheus.MustRegister(incomingTraffic)
	prometheus.MustRegister(outgoingTraffic)

	r.Route("/v1/promotions", func(r chi.Router) {
		r.Get("/{id}", h.GetPromotionHandler)
		r.Delete("/{id}", h.DeletePromotionHandler)
		r.Post("/", h.CreatePromotionHandler)
		r.Put("/{id}", h.UpdatePromotionHandler)
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			incomingTraffic.Inc()
			outgoingTraffic.Inc()

			writer.Write([]byte("test"))
		})
	})

	r.Handle("/metrics", promhttp.Handler())

	return r
}
