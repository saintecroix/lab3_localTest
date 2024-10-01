package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

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