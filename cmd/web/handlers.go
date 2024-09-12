package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Gruz struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	ETSNG string `json:"etsng"`
	GNG   string `json:"gng"`
}

func getGruz(db *sql.DB) []Gruz {
	gruz, err := db.Query("Select * from gruz order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Gruz, 0)

	for gruz.Next() {
		var a Gruz
		err = gruz.Scan(&a.Id, &a.Name, &a.ETSNG, &a.GNG)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Consignee struct {
	Id         int
	Name       string
	OKPO       string
	Was_sender int
}

func getConsignee(db *sql.DB) []Consignee {
	gruz, err := db.Query("Select * from consignee order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Consignee, 0)

	for gruz.Next() {
		var a Consignee
		err = gruz.Scan(&a.Id, &a.Name, &a.OKPO, &a.Was_sender)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Region struct {
	Id   int
	Name string
}

func getRegion(db *sql.DB) []Region {
	gruz, err := db.Query("Select * from region order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Region, 0)

	for gruz.Next() {
		var a Region
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Road struct {
	Id     int
	Name   string
	Length int
}

func getRoad(db *sql.DB) []Road {
	gruz, err := db.Query("Select * from road order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Road, 0)

	for gruz.Next() {
		var a Road
		err = gruz.Scan(&a.Id, &a.Name, &a.Length)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type State struct {
	Id   int
	Name string
}

func getState(db *sql.DB) []State {
	gruz, err := db.Query("Select * from state order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]State, 0)

	for gruz.Next() {
		var a State
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Station struct {
	Id   int
	Name string
	ESR  string
	Road string
}

func getStation(db *sql.DB) []Station {
	gruz, err := db.Query("Select * from station order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Station, 0)

	for gruz.Next() {
		var a Station
		err = gruz.Scan(&a.Id, &a.Name, &a.ESR, &a.Road)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Wagon struct {
	Id   int
	Name string
}
type AllInfo struct {
	Gruz      []Gruz
	Consignee []Consignee
	Region    []Region
	Road      []Road
	State     []State
	Station   []Station
	Wagon     []Wagon
}
type Application struct {
	Id                  int    `json:"id_app"`
	Number              int    `json:"number_app"`
	Reg_date            string `json:"regDate_app"`
	Status              string `json:"status_app"`
	Provide_date        string `json:"provideDate_app"`
	Departure_type      string `json:"departureType_app"`
	Goods               string `json:"goods_app"`
	Origin_state        string `json:"originState_app"`
	Enter_station       string `json:"enterStation_app"`
	Region_depart       string `json:"regionDepart_app"`
	Road_depart         string `json:"roadDepart_app"`
	Station_depart      string `json:"stationDepart_app"`
	Consigner           string `json:"consigner_app"`
	State_destination   string `json:"stateDestination_app"`
	Exit_station        string `json:"exitStation_app"`
	Region_destination  string `json:"regionDestination_app"`
	Road_destination    string `json:"roadDestination_app"`
	Station_destination string `json:"stationDestination_app"`
	Consignee           string `json:"consignee_app"`
	Wagon_type          string `json:"wagonType_app"`
	Property            string `json:"property_app"`
	Wagon_owner         string `json:"wagonOwner_app"`
	Payer               string `json:"payer_app"`
	Road_owner          string `json:"roadOwner_app"`
	Transport_manager   string `json:"transportManager_app"`
	Tons_declared       int    `json:"tonsDeclared_app"`
	Tons_accepted       int    `json:"tonsAccepted_app"`
	Wagon_declared      int    `json:"wagonDeclared_app"`
	Wagon_accepted      int    `json:"wagonAccepted_app"`
	Filing_date         string `json:"filingDate_app"`
	Agreement_date      string `json:"agreementDate_app"`
	Approval_date       string `json:"approvalDate_app"`
	Start_date          string `json:"startDate_app"`
	Stop_date           string `json:"stopDate_app"`
}

func getWagon(db *sql.DB) []Wagon {
	gruz, err := db.Query("Select * from wagon order by Name")
	if err != nil {
		panic(err)
	}

	rez := make([]Wagon, 0)

	for gruz.Next() {
		var a Wagon
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) input(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/input.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	type Jopa struct {
		Gruz      []Gruz
		Consignee []Consignee
		Region    []Region
		Road      []Road
		State     []State
		Station   []Station
		Wagon     []Wagon
	}

	err = ts.Execute(w, Jopa{Gruz: getGruz(db), Consignee: getConsignee(db), Region: getRegion(db), Road: getRoad(db),
		State: getState(db), Station: getStation(db), Wagon: getWagon(db)})
	if err != nil {
		app.serverError(w, err)
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_application(w http.ResponseWriter, r *http.Request) {
	number, _ := strconv.Atoi(r.FormValue("Number"))
	reg_date := r.FormValue("Reg_date")
	status := r.FormValue("Status")
	provide_date := r.FormValue("Provide_date")
	departure_type := r.FormValue("Departure_type")
	property := r.FormValue("Property")
	wagon_owner := r.FormValue("Wagon_owner")
	payer := r.FormValue("Payer")
	road_owner := r.FormValue("Road_owner")
	transport_manager := r.FormValue("Transport_manager")
	tons_declared, _ := strconv.Atoi(r.FormValue("Tons_declared"))
	tons_accepted, _ := strconv.Atoi(r.FormValue("Tons_accepted"))
	wagon_declared, _ := strconv.Atoi(r.FormValue("Wagon_declared"))
	wagon_accepted, _ := strconv.Atoi(r.FormValue("Wagon_accepted"))
	filing_date := r.FormValue("Filing_date")
	agreement_date := r.FormValue("Agreement_date")
	approval_date := r.FormValue("Approval_date")
	start_date := r.FormValue("Start_date")
	stop_date := r.FormValue("Stop_date")

	db := app.db

	goods, _ := strconv.Atoi(r.FormValue("Goods"))

	origin_state, _ := strconv.Atoi(r.FormValue("Origin_state"))

	enter_station, _ := strconv.Atoi(r.FormValue("Enter_station"))

	region_depart, _ := strconv.Atoi(r.FormValue("Region_depart"))

	road_depart, _ := strconv.Atoi(r.FormValue("Road_depart"))

	station_depart, _ := strconv.Atoi(r.FormValue("Station_depart"))

	consigner, _ := strconv.Atoi(r.FormValue("Consigner"))

	state_destination, _ := strconv.Atoi(r.FormValue("State_destination"))

	exit_station, _ := strconv.Atoi(r.FormValue("Exit_station"))

	region_destination, _ := strconv.Atoi(r.FormValue("Region_destination"))

	road_destination, _ := strconv.Atoi(r.FormValue("Road_destination"))

	station_destination, _ := strconv.Atoi(r.FormValue("Station_destination"))

	consignees, _ := strconv.Atoi(r.FormValue("Consignee"))

	wagon_type, _ := strconv.Atoi(r.FormValue("Wagon_type"))

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `application` "+
		"(`Number`, `Reg_date`, `Status`, `Provide_date`, `Departure_type`, `Goods`, `Origin_state`, "+
		"`Enter_station`, `Region_depart`, `Road_depart`, `Station_depart`, `Consigner`, `State_destination`, "+
		"`Exit_station`, `Region_destination`, `Road_destination`, `Station_destination`, `Consignee`, "+
		"`Wagon_type`, `Property`, `Wagon_owner`, `Payer`, `Road_owner`, `Transport_manager`, `Tons_declared`, "+
		"`Tons_accepted`, `Wagon_declared`, `Wagon_accepted`, `Filing_date`, `Agreement_date`, `Approval_date`, "+
		"`Start_date`, `Stop_date`) VALUES('%d', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', "+
		"'%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', "+
		"'%s', '%s', '%s', '%s', '%s')", number, reg_date, status, provide_date, departure_type, goods, origin_state,
		enter_station, region_depart, road_depart, station_depart, consigner, state_destination, exit_station,
		region_destination, road_destination, station_destination, consignees, wagon_type, property, wagon_owner,
		payer, road_owner, transport_manager, tons_declared, tons_accepted, wagon_declared, wagon_accepted, filing_date,
		agreement_date, approval_date, start_date, stop_date))
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) addGruz(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_gruz.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Gruz []Gruz
	}

	err = ts.Execute(w, Jopa{Gruz: getGruz(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_gruz(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	etsng := r.FormValue("ETSNG")
	gng := r.FormValue("GNG")

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `gruz` "+
		"(`Name`, `ETSNG`, `GNG`) VALUES('%s', '%s', '%s')", name, etsng, gng))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_state(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_state.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		State []State
	}

	err = ts.Execute(w, Jopa{State: getState(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_state(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `state` "+
		"(`Name`) VALUES('%s')", name))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_station(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_station.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Station []Station
	}

	err = ts.Execute(w, Jopa{Station: getStation(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_station(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	esr := r.FormValue("ESR")
	road := r.FormValue("Road")

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `station` "+
		"(`Name`,`ESR`, `Road`) VALUES('%s', `%s`, `%s`)", name, esr, road))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_region(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_region.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Region []Region
	}

	err = ts.Execute(w, Jopa{Region: getRegion(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_region(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `region` "+
		"(`Name`) VALUES('%s')", name))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_road(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_road.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Road []Road
	}

	err = ts.Execute(w, Jopa{Road: getRoad(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_road(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	length, _ := strconv.Atoi(r.FormValue("Length"))

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `road` "+
		"(`Name`, `Length`) VALUES('%s', '%d')", name, length))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_consignee(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_consignee.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Consignee []Consignee
	}

	err = ts.Execute(w, Jopa{Consignee: getConsignee(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_consignee(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	okpo := r.FormValue("OKPO")
	var a int
	if r.FormValue("Was_sender") != "" {
		a = 1
	} else {
		a = 0
	}

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `consignee` "+
		"(`Name`, `OKPO`, `Was_sender`) VALUES('%s', '%s', '%d')", name, okpo, a))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_wagon(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_wagon.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Wagon []Wagon
	}

	err = ts.Execute(w, Jopa{Wagon: getWagon(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_wagon(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	db := app.db

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `wagon` "+
		"(`Name`) VALUES('%s')", name))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

type Average struct {
	Month_reg_date      int
	Avg_transport_level float32
}

type RegPop struct {
	FrstReg string
	SecReg  string
	Usages  int64
}

type PopRoad struct {
	FrstRoad string
	SecRoad  string
	Usage    int64
}

func (app *application) stats(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/stats.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db := app.db

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	avg, err := db.Query("SELECT month(reg_date) AS month_reg_date, avg(n_transportations) AS " +
		"avg_transport_level FROM (SELECT reg_date, count(*) AS n_transportations FROM application GROUP BY reg_date) " +
		"AS q GROUP BY month(reg_date);")
	if err != nil {
		app.serverError(w, err)
		return
	}

	avgRez := make([]Average, 0)

	for avg.Next() {
		var a Average
		err := avg.Scan(&a.Month_reg_date, &a.Avg_transport_level)
		if err != nil {
			panic(err)
		}
		avgRez = append(avgRez, a)
	}

	popReg, err := db.Query("SELECT o.Name, oo.Name, count(*) as n FROM (application AS h INNER JOIN region AS o" +
		" ON h.Region_depart=o.id) INNER JOIN region AS oo ON h.Region_destination=oo.ID GROUP BY o.Name, oo.Name ORDER BY count(*) DESC;")
	if err != nil {
		app.serverError(w, err)
		return
	}

	rezReg := make([]RegPop, 0)

	for popReg.Next() {
		var a RegPop
		err := popReg.Scan(&a.FrstReg, &a.SecReg, &a.Usages)
		if err != nil {
			panic(err)
		}
		rezReg = append(rezReg, a)
	}

	popRoad, err := db.Query("SELECT d.Name, do.Name, q.n_transportations FROM (SELECT Road_depart, Road_destination, " +
		"count(*) AS n_transportations FROM application GROUP BY Road_depart, Road_destination) AS q INNER JOIN road AS d ON " +
		"d.id = q.Road_depart INNER JOIN road AS do ON do.id=q.Road_destination ORDER BY q.n_transportations DESC LIMIT 10;")
	if err != nil {
		app.serverError(w, err)
		return
	}

	rezRoad := make([]PopRoad, 0)

	for popRoad.Next() {
		var a PopRoad
		err := popRoad.Scan(&a.FrstRoad, &a.SecRoad, &a.Usage)
		if err != nil {
			app.serverError(w, err)
			return
		}
		rezRoad = append(rezRoad, a)
	}

	type Jopa struct {
		RezReg  []RegPop
		AvgRez  []Average
		RezRoad []PopRoad
	}

	err = ts.Execute(w, Jopa{RezReg: rezReg, AvgRez: avgRez, RezRoad: rezRoad})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) soloSearch(w http.ResponseWriter, r *http.Request) {
	where := r.FormValue("Goods")
	what := r.FormValue("searchText")

	files := []string{
		"./ui/html/soloSearch.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	db := app.db

	Rez := make([]Application, 0)

	switch where {
	case "Goods":
		Rez = getApplication(app, db, w, " WHERE g.Name like '%%%s%%'", what)
	case "Origin_state":
		Rez = getApplication(app, db, w, " WHERE state.Name like '%%%s%%'", what)
	case "State_destination":
		Rez = getApplication(app, db, w, " WHERE state1.Name like '%%%s%%'", what)
	case "Region_depart":
		Rez = getApplication(app, db, w, " WHERE reg.Name like '%%%s%%'", what)
	case "Region_destination":
		Rez = getApplication(app, db, w, " WHERE reg1.Name like '%%%s%%'", what)
	case "Road_depart":
		Rez = getApplication(app, db, w, " WHERE r.Name like '%%%s%%'", what)
	case "Road_destination":
		Rez = getApplication(app, db, w, " WHERE r1.Name like '%%%s%%'", what)
	case "Station_depart":
		Rez = getApplication(app, db, w, " WHERE st1.Name like '%%%s%%'", what)
	case "Exit_station":
		Rez = getApplication(app, db, w, " WHERE st2.Name like '%%%s%%'", what)
	case "Enter_station":
		Rez = getApplication(app, db, w, " WHERE st.Name like '%%%s%%'", what)
	case "Station_destination":
		Rez = getApplication(app, db, w, " WHERE st3.Name like '%%%s%%'", what)
	case "Consigner":
		Rez = getApplication(app, db, w, " WHERE con.Name like '%%%s%%'", what)
	case "Consignee":
		Rez = getApplication(app, db, w, " WHERE con1.Name like '%%%s%%'", what)
	case "Wagon_type":
		Rez = getApplication(app, db, w, " WHERE w.Name like '%%%s%%'", what)
	case "Status":
		Rez = getApplication(app, db, w, " WHERE a.Status like '%%%s%%'", what)
	}

	type Jopa struct {
		Rez []Application
	}
	err = ts.Execute(w, Jopa{Rez: Rez})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

func (app *application) duoSearch(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/duoSearch.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	db := app.db

	gruz, err := db.Query("Select * from wagon where id in (select distinct Wagon_type from application);")
	if err != nil {
		panic(err)
	}

	rez := make([]Wagon, 0)

	for gruz.Next() {
		var a Wagon
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}

	err = ts.Execute(w, AllInfo{Wagon: rez})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

func getApplication(app *application, db *sql.DB, w http.ResponseWriter, whereString string, what string) []Application {
	sel := "select a.id, a.Number, a.Reg_date, a.Status, a.Provide_date, a.Departure_type, g.Name as Goods, state.Name " +
		"as Origin_state, st.Name as Enter_station, reg.Name as Region_depart, r.Name as Road_depart, st1.Name as " +
		"Station_depart, con.Name as Consigner, state1.Name as State_destination, st2.Name as Exit_station, reg1.Name " +
		"as Region_destination, r1.Name as Road_destination, st3.Name as Station_destination, con1.Name as Consignee, " +
		"w.Name as Wagon_type, a.Property, a.Wagon_owner, a.Payer, a.Road_owner, a.Transport_manager, a.Tons_declared, " +
		"a.Tons_accepted, a.Wagon_declared, a.Wagon_accepted, a.Filing_date, a.Agreement_date, a.Approval_date, " +
		"a.Start_date, a.Stop_date from application as a inner join gruz as g on a.Goods=g.id inner join state on " +
		"a.Origin_state=state.id inner join station as st on a.Enter_station=st.id inner join region as reg on " +
		"a.Region_depart=reg.id inner join road as r on a.Road_depart=r.id inner join (select * from station) " +
		"as st1 on a.Station_depart=st1.id inner join consignee as con on a.Consigner=con.id inner join (select * from " +
		"state) as state1 on a.State_destination=state1.id inner join (select * from station) as st2 on " +
		"a.Exit_station=st2.id inner join (select * from region) as reg1 on a.Region_destination=reg1.id inner join " +
		"(select * from road) as r1 on a.Road_destination=r1.id inner join (select * from station) as st3 on " +
		"a.Station_destination=st3.id inner join (select * from consignee) as con1 on a.Consignee=con1.id inner join " +
		"wagon as w on a.Wagon_type=w.id"

	rez := make([]Application, 0)
	get, err := db.Query(fmt.Sprintf(sel+whereString, what))
	if err != nil {
		app.serverError(w, err)
		return rez
	}
	defer get.Close()
	for get.Next() {
		var v Application
		err := get.Scan(&v.Id, &v.Number, &v.Reg_date, &v.Status, &v.Provide_date, &v.Departure_type, &v.Goods,
			&v.Origin_state, &v.Enter_station, &v.Region_depart, &v.Road_depart, &v.Station_depart, &v.Consigner,
			&v.State_destination, &v.Exit_station, &v.Region_destination, &v.Road_destination, &v.Station_destination,
			&v.Consignee, &v.Wagon_type, &v.Property, &v.Wagon_owner, &v.Payer, &v.Road_owner, &v.Transport_manager,
			&v.Tons_declared, &v.Tons_accepted, &v.Wagon_declared, &v.Wagon_accepted, &v.Filing_date, &v.Agreement_date,
			&v.Approval_date, &v.Start_date, &v.Stop_date)
		if err != nil {
			app.serverError(w, err)
			return rez
		}
		rez = append(rez, v)
	}
	return rez
}

/*----------------------------------------------------------------------------------------*/

func (app *application) getSecondPerSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
		return
	}
	type fromSite struct {
		CheckData string `json:"checkFirstEl"`
	}
	decoder := json.NewDecoder(r.Body)
	data := &fromSite{}
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}

	db := app.db

	getRaw, err := db.Query(fmt.Sprintf("SELECT * FROM `gruz` WHERE id in (SELECT Goods FROM `application` WHERE Wagon_type = %s);", data.CheckData))
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer getRaw.Close()
	rez := make([]Gruz, 0)
	for getRaw.Next() {
		var a Gruz
		err := getRaw.Scan(&a.Id, &a.Name, &a.ETSNG, &a.GNG)
		if err != nil {
			app.serverError(w, err)
		}
		rez = append(rez, a)
	}
	json, err := json.Marshal(rez)
	if err != nil {
		http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (app *application) letDuoSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
		return
	}
	type fromSite struct {
		Duos string `json:"checkDuos"`
		Gruz string `json:"checkGruz"`
	}
	decoder := json.NewDecoder(r.Body)
	data := &fromSite{}
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	fmt.Println(data.Duos, data.Gruz)
	db := app.db
	sel := "select a.id, a.Number, a.Reg_date, a.Status, a.Provide_date, a.Departure_type, g.Name as Goods, state.Name " +
		"as Origin_state, st.Name as Enter_station, reg.Name as Region_depart, r.Name as Road_depart, st1.Name as " +
		"Station_depart, con.Name as Consigner, state1.Name as State_destination, st2.Name as Exit_station, reg1.Name " +
		"as Region_destination, r1.Name as Road_destination, st3.Name as Station_destination, con1.Name as Consignee, " +
		"w.Name as Wagon_type, a.Property, a.Wagon_owner, a.Payer, a.Road_owner, a.Transport_manager, a.Tons_declared, " +
		"a.Tons_accepted, a.Wagon_declared, a.Wagon_accepted, a.Filing_date, a.Agreement_date, a.Approval_date, " +
		"a.Start_date, a.Stop_date from application as a inner join gruz as g on a.Goods=g.id inner join state on " +
		"a.Origin_state=state.id inner join station as st on a.Enter_station=st.id inner join region as reg on " +
		"a.Region_depart=reg.id inner join road as r on a.Road_depart=r.id inner join (select * from station) " +
		"as st1 on a.Station_depart=st1.id inner join consignee as con on a.Consigner=con.id inner join (select * from " +
		"state) as state1 on a.State_destination=state1.id inner join (select * from station) as st2 on " +
		"a.Exit_station=st2.id inner join (select * from region) as reg1 on a.Region_destination=reg1.id inner join " +
		"(select * from road) as r1 on a.Road_destination=r1.id inner join (select * from station) as st3 on " +
		"a.Station_destination=st3.id inner join (select * from consignee) as con1 on a.Consignee=con1.id inner join " +
		"wagon as w on a.Wagon_type=w.id"

	whereString := " WHERE w.id = %s and g.id = %s"
	rez := make([]Application, 0)
	get, err := db.Query(fmt.Sprintf(sel+whereString, data.Duos, data.Gruz))
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer get.Close()
	for get.Next() {
		var v Application
		err := get.Scan(&v.Id, &v.Number, &v.Reg_date, &v.Status, &v.Provide_date, &v.Departure_type, &v.Goods,
			&v.Origin_state, &v.Enter_station, &v.Region_depart, &v.Road_depart, &v.Station_depart, &v.Consigner,
			&v.State_destination, &v.Exit_station, &v.Region_destination, &v.Road_destination, &v.Station_destination,
			&v.Consignee, &v.Wagon_type, &v.Property, &v.Wagon_owner, &v.Payer, &v.Road_owner, &v.Transport_manager,
			&v.Tons_declared, &v.Tons_accepted, &v.Wagon_declared, &v.Wagon_accepted, &v.Filing_date, &v.Agreement_date,
			&v.Approval_date, &v.Start_date, &v.Stop_date)
		if err != nil {
			app.serverError(w, err)
			return
		}
		rez = append(rez, v)
	}
	jsonApps, err := json.Marshal(rez)
	if err != nil {
		http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonApps)
}

func (app *application) aboutDB(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/about_db.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) xmlPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/xml.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	db := app.db

	topTen, err1 := db.Query("select a.id, a.Number, a.Reg_date, a.Status, a.Provide_date, a.Departure_type, g.Name as Goods, state.Name " +
		"as Origin_state, st.Name as Enter_station, reg.Name as Region_depart, r.Name as Road_depart, st1.Name as " +
		"Station_depart, con.Name as Consigner, state1.Name as State_destination, st2.Name as Exit_station, reg1.Name " +
		"as Region_destination, r1.Name as Road_destination, st3.Name as Station_destination, con1.Name as Consignee, " +
		"w.Name as Wagon_type, a.Property, a.Wagon_owner, a.Payer, a.Road_owner, a.Transport_manager, a.Tons_declared, " +
		"a.Tons_accepted, a.Wagon_declared, a.Wagon_accepted, a.Filing_date, a.Agreement_date, a.Approval_date, " +
		"a.Start_date, a.Stop_date from application as a inner join gruz as g on a.Goods=g.id inner join state on " +
		"a.Origin_state=state.id inner join station as st on a.Enter_station=st.id inner join region as reg on " +
		"a.Region_depart=reg.id inner join road as r on a.Road_depart=r.id inner join (select * from station) " +
		"as st1 on a.Station_depart=st1.id inner join consignee as con on a.Consigner=con.id inner join (select * from " +
		"state) as state1 on a.State_destination=state1.id inner join (select * from station) as st2 on " +
		"a.Exit_station=st2.id inner join (select * from region) as reg1 on a.Region_destination=reg1.id inner join " +
		"(select * from road) as r1 on a.Road_destination=r1.id inner join (select * from station) as st3 on " +
		"a.Station_destination=st3.id inner join (select * from consignee) as con1 on a.Consignee=con1.id inner join " +
		"wagon as w on a.Wagon_type=w.id order by a.id DESC limit 10;")

	if err1 != nil {
		app.serverError(w, err1)
		return
	}

	resoult := make([]Application, 0)

	for topTen.Next() {
		var a Application
		err = topTen.Scan(&a.Id, &a.Number, &a.Reg_date, &a.Status, &a.Provide_date, &a.Departure_type, &a.Goods, &a.Origin_state,
			&a.Enter_station, &a.Region_depart, &a.Road_depart, &a.Station_depart, &a.Consigner, &a.State_destination,
			&a.Exit_station, &a.Region_destination, &a.Road_destination, &a.Station_destination, &a.Consignee, &a.Wagon_type,
			&a.Property, &a.Wagon_owner, &a.Payer, &a.Road_owner, &a.Transport_manager, &a.Tons_declared, &a.Tons_accepted,
			&a.Wagon_declared, &a.Wagon_accepted, &a.Filing_date, &a.Agreement_date, &a.Approval_date, &a.Start_date,
			&a.Stop_date)
		if err != nil {
			app.serverError(w, err)
			return
		}
		resoult = append(resoult, a)
	}

	xmlfile, err := os.OpenFile("./ui/static/file.xml", os.O_RDWR, 0644)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer xmlfile.Close()
	_, err = xmlfile.Seek(0, 0)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = xmlfile.Truncate(0)
	if err != nil {
		app.serverError(w, err)
		return
	}
	encoder := xml.NewEncoder(xmlfile)
	encoder.Indent(" ", "	")

	type DataXML struct {
		Application []Application `xml:"Application"`
	}

	if xmlErr := encoder.Encode(DataXML{Application: resoult}); xmlErr != nil {
		app.serverError(w, xmlErr)
		return
	}

	err = ts.Execute(w, DataXML{Application: resoult})
	if err != nil {
		app.serverError(w, err)
		return
	}
}
