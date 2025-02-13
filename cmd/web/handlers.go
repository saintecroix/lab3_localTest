package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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
	decoder := json.NewDecoder(r.Body)
	data := &Application{}
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		return
	}
	fmt.Println(data)

	db := app.db

	type Result struct {
		Success bool  `json:"success"`
		Error   error `json:"error,omitempty"`
	}
	var res Result

	good, err := strconv.Atoi(data.Goods)
	origin_state, err := strconv.Atoi(data.Origin_state)
	enter_station, err := strconv.Atoi(data.Enter_station)
	region_depart, err := strconv.Atoi(data.Region_depart)
	road_depart, err := strconv.Atoi(data.Road_depart)
	station_depart, err := strconv.Atoi(data.Station_depart)
	consigner, err := strconv.Atoi(data.Consigner)
	state_destination, err := strconv.Atoi(data.State_destination)
	exit_station, err := strconv.Atoi(data.Exit_station)
	region_destination, err := strconv.Atoi(data.Region_destination)
	road_destination, err := strconv.Atoi(data.Road_destination)
	station_destination, err := strconv.Atoi(data.Station_destination)
	consignee, err := strconv.Atoi(data.Consignee)
	wagon_type, err := strconv.Atoi(data.Wagon_type)

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `application` "+
		"(`Number`, `Reg_date`, `Status`, `Provide_date`, `Departure_type`, `Goods`, `Origin_state`, "+
		"`Enter_station`, `Region_depart`, `Road_depart`, `Station_depart`, `Consigner`, `State_destination`, "+
		"`Exit_station`, `Region_destination`, `Road_destination`, `Station_destination`, `Consignee`, "+
		"`Wagon_type`, `Property`, `Wagon_owner`, `Payer`, `Road_owner`, `Transport_manager`, `Tons_declared`, "+
		"`Tons_accepted`, `Wagon_declared`, `Wagon_accepted`, `Filing_date`, `Agreement_date`, `Approval_date`, "+
		"`Start_date`, `Stop_date`) VALUES('%d', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', "+
		"'%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', "+
		"'%s', '%s', '%s', '%s', '%s')", data.Number, data.Reg_date, data.Status, data.Provide_date,
		data.Departure_type, good, origin_state, enter_station, region_depart, road_depart,
		station_depart, consigner, state_destination, exit_station, region_destination,
		road_destination, station_destination, consignee, wagon_type, data.Property,
		data.Wagon_owner, data.Payer, data.Road_owner, data.Transport_manager, data.Tons_declared, data.Tons_accepted,
		data.Wagon_declared, data.Wagon_accepted, data.Filing_date, data.Agreement_date, data.Approval_date,
		data.Start_date, data.Stop_date))
	if err != nil {
		res = Result{Success: false, Error: fmt.Errorf("Ошибка при записи данных в бд: %v", err)}
	}
	defer insert.Close()

	res.Success = true
	jsonApps, err := json.Marshal(res)
	if err != nil {
		res = Result{Success: false, Error: fmt.Errorf("Ошибка при кодировании данных в JSON: %v", err)}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonApps)
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

/*----------------------------------------------------------------------------------------*/

func (app *application) getSecondPerSearch(w http.ResponseWriter, r *http.Request) {
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

func (app *application) regPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/regPage.page.tmpl",
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

func (app *application) registration(w http.ResponseWriter, r *http.Request) {
	type jsonData struct {
		User  string `json:"username"`
		Pass  string `json:"password"`
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	data := &jsonData{}
	err := decoder.Decode(data)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		app.serverError(w, err)
		return
	}
	type sendData struct {
		Success          bool   `json:"success"`
		ExistingUsername string `json:"existingUsername"`
		ExistingEmail    string `json:"existingEmail"`
	}
	res := sendData{}

	login, err := app.dbSearch("login", "user", fmt.Sprintf("where login = '%s'", data.User))
	if err != nil {
		app.serverError(w, err)
		return
	}
	if len(login) == 0 {
		res.ExistingUsername = ""
	} else {
		res.ExistingUsername = login[0]
	}

	email, err := app.dbSearch("mail", "user", fmt.Sprintf("where mail = '%s'", data.Email))
	if err != nil {
		app.serverError(w, err)
		return
	}
	if len(email) == 0 {
		res.ExistingEmail = ""
	} else {
		res.ExistingEmail = email[0]
	}

	if res.ExistingEmail == "" && res.ExistingUsername == "" {
		_, err = app.db.Query(fmt.Sprintf("INSERT INTO `user`(`login`, `password`, `mail`) VALUES ('%s','%s','%s')", data.User, data.Pass, data.Email))
		if err != nil {
			app.serverError(w, err)
			return
		}
		res.Success = true
	}

	answ, err := json.Marshal(res)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answ)
}

func (app *application) authorisation(w http.ResponseWriter, r *http.Request) {
	type authData struct {
		User string `json:"auth_username"`
		Pass string `json:"auth_password"`
	}
	decoder := json.NewDecoder(r.Body)
	getData := &authData{}
	err := decoder.Decode(getData)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusBadRequest)
		app.serverError(w, err)
		return
	}

	res := struct {
		User string `json:"existingUser"`
		Pass string `json:"existingPass"`
	}{User: "", Pass: ""}
	login, err := app.dbSearch("login", "user", fmt.Sprintf("where login = '%s'", getData.User))
	if err != nil {
		app.serverError(w, err)
		return
	}
	if len(login) > 0 {
		res.User = login[0]
		pass, err := app.dbSearch("password", "user", fmt.Sprintf("where login = '%s'", getData.User))
		if err != nil {
			app.serverError(w, err)
			return
		}
		if len(pass) > 0 && pass[0] == getData.Pass {
			res.Pass = pass[0]

			tokenString, err := app.createJWT(getData.User)
			if err != nil {
				app.serverError(w, err)
				return
			}
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", tokenString))
		}
	}

	answ, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Ошибка при кодировании запроса", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answ)
}

// Обработчик страницы новостной ленты
func (app *application) rssPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/rssPage.page.tmpl",
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

func (app *application) galleryPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/gallery.page.tmpl",
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

func (app *application) equipmentPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/equipment.page.tmpl",
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

func (app *application) platform1Page(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/platform1.page.tmpl",
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

func (app *application) closePlatformPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/platform2.page.tmpl",
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

func (app *application) semiPlatformPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/platform3.page.tmpl",
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

func (app *application) widgets(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/widgets.page.tmpl",
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

func (app *application) handleVisits(w http.ResponseWriter, r *http.Request) {
	cookie, err := app.handleVisit(r)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.SetCookie(w, cookie)
	raw, err := app.db.Query("SELECT COUNT(*) FROM user_visits")
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer raw.Close()
	var count int
	if raw.Next() {
		err = raw.Scan(&count)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	type Result struct {
		Visits int `json:"visits"`
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Result{Visits: count})
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) newPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/local_new.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	vars := mux.Vars(r)
	newsID, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.serverError(w, err)
		return
	}
	localNew, err := app.gotNew(newsID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	type Result struct {
		New *LocalNews
	}

	err = ts.Execute(w, Result{New: localNew})
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) addSource(w http.ResponseWriter, r *http.Request) {
	type JsonData struct {
		Login string `json:"login"`
		Url   string `json:"url"`
	}
	type Response struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error,omitempty"`
	}
	decoder := json.NewDecoder(r.Body)
	var data JsonData
	var response Response
	err := decoder.Decode(&data)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}
	if !app.sourceValid(data.Url) {
		response.Success = false
		response.Error = "url is invalid"
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	userExists, err := app.userExists(data.Login)
	if err != nil {
		response.Success = false
		response.Error = fmt.Errorf("error checking user existance: %w", err).Error()
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	if !userExists {
		response.Success = false
		response.Error = "user not found"
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	_, err = app.db.Exec( // Выполняем запрос через app.db.Exec()
		"UPDATE `user` SET `news_sources` = CASE WHEN `news_sources` IS NULL THEN ? ELSE CONCAT(`news_sources`, ',', ?) END WHERE `login` = ?", data.Url, data.Url, data.Login)
	if err != nil {
		response.Success = false
		response.Error = fmt.Errorf("error updating db: %w", err).Error()
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response.Success = true
	app.jsonResponse(w, http.StatusOK, response)
	return
}

func (app *application) sourcesPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/sources.page.tmpl",
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

func (app *application) sources(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		User string `json:"user"`
	}
	type Response struct {
		Sources []string    `json:"sources"`
		Error   interface{} `json:"error,omitempty"`
	}
	var response Response
	var request Request
	err := jsonRequest(r, &request)
	if err != nil {
		app.logError(err)
		response.Error = err
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	userExists, err := app.userExists(request.User)
	if err != nil {
		response.Error = fmt.Errorf("error checking user existance: %w", err).Error()
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	if !userExists {
		response.Error = "user not found"
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	sources, err := app.getSources(request.User)
	if err != nil {
		app.logError(err)
		response.Error = err
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}

	response.Sources = strings.Split(sources, ",")
	app.jsonResponse(w, http.StatusOK, response)
	return
}

func (app *application) deleteSource(w http.ResponseWriter, r *http.Request) {
	type UserRequest struct {
		Sources []string `json:"sources"`
		User    string   `json:"user"`
	}
	type ServerResponse struct {
		Error interface{} `json:"error,omitempty"`
	}
	var response ServerResponse
	var request UserRequest
	err := jsonRequest(r, &request)
	if err != nil {
		app.logError(err)
		response.Error = err
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}
	userExists, err := app.userExists(request.User)
	if err != nil {
		app.logError(err)
		response.Error = fmt.Errorf("error checking user existance: %w", err).Error()
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}
	if !userExists {
		response.Error = "user not found"
		app.jsonResponse(w, http.StatusBadRequest, response)
		return
	}
	err = app.deleteSources(request.User, request.Sources)
	if err != nil {
		app.logError(err)
		response.Error = fmt.Errorf("error deleting sources: %w", err).Error()
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}
	response.Error = nil
	app.jsonResponse(w, http.StatusOK, response)
	return
}
