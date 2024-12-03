package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/", app.home).Methods(http.MethodGet)
	m.HandleFunc("/input", app.input).Methods(http.MethodGet)
	m.HandleFunc("/save_application", app.save_application).Methods(http.MethodPost)
	m.HandleFunc("/add_gruz", app.addGruz).Methods(http.MethodGet)
	m.HandleFunc("/save_gruz", app.save_gruz).Methods(http.MethodPost)
	m.HandleFunc("/add_state", app.add_state).Methods(http.MethodGet)
	m.HandleFunc("/save_state", app.save_state).Methods(http.MethodPost)
	m.HandleFunc("/add_station", app.add_station).Methods(http.MethodGet)
	m.HandleFunc("/save_station", app.save_station).Methods(http.MethodPost)
	m.HandleFunc("/add_region", app.add_region).Methods(http.MethodGet)
	m.HandleFunc("/save_region", app.save_region).Methods(http.MethodPost)
	m.HandleFunc("/add_road", app.add_road).Methods(http.MethodGet)
	m.HandleFunc("/save_road", app.save_road).Methods(http.MethodPost)
	m.HandleFunc("/add_consignee", app.add_consignee).Methods(http.MethodGet)
	m.HandleFunc("/save_consignee", app.save_consignee).Methods(http.MethodPost)
	m.HandleFunc("/add_wagon", app.add_wagon).Methods(http.MethodGet)
	m.HandleFunc("/save_wagon", app.save_wagon).Methods(http.MethodPost)
	m.HandleFunc("/stats", app.stats).Methods(http.MethodGet)
	m.HandleFunc("/soloSearch", app.soloSearch)
	m.HandleFunc("/duoSearch", app.duoSearch).Methods(http.MethodGet)
	m.HandleFunc("/getSecondPerSearch", app.getSecondPerSearch).Methods(http.MethodPost)
	m.HandleFunc("/letDuoSearch", app.letDuoSearch).Methods(http.MethodPost)
	m.HandleFunc("/about_db", app.aboutDB).Methods(http.MethodGet)
	m.HandleFunc("/xml", app.xmlPage).Methods(http.MethodGet)
	m.HandleFunc("/reg", app.regPage).Methods(http.MethodGet)
	m.HandleFunc("/registration", app.registration).Methods(http.MethodPost)
	m.HandleFunc("/authorization", app.authorisation).Methods(http.MethodPost)
	m.HandleFunc("/test", app.test)
	m.HandleFunc("/protected", app.protected).Methods(http.MethodPost)
	m.HandleFunc("/rss", app.rssPage).Methods(http.MethodGet)
	m.HandleFunc("/addNew", app.addNew).Methods(http.MethodPost)
	m.HandleFunc("/indexNews", app.indexNews).Methods(http.MethodGet)
	m.HandleFunc("/gallery", app.galleryPage).Methods(http.MethodGet)
	m.HandleFunc("/equipment", app.equipmentPage).Methods(http.MethodGet)
	m.HandleFunc("/platform1", app.platform1Page).Methods(http.MethodGet)
	m.HandleFunc("/platform2", app.closePlatformPage).Methods(http.MethodGet)
	m.HandleFunc("/platform3", app.semiPlatformPage).Methods(http.MethodGet)
	m.HandleFunc("/widgets", app.widgets).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	m.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return m
}
