package main

import (
	"github.com/mmcdole/gofeed"
	"time"
)

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

type Station struct {
	Id   int
	Name string
	ESR  string
	Road string
}

type State struct {
	Id   int
	Name string
}

type Road struct {
	Id     int
	Name   string
	Length int
}

type Region struct {
	Id   int
	Name string
}

type Consignee struct {
	Id         int
	Name       string
	OKPO       string
	Was_sender int
}

type Gruz struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	ETSNG string `json:"etsng"`
	GNG   string `json:"gng"`
}

type LocalNews struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	User    string    `json:"user"`
	Date    time.Time `json:"date"`
	ResDate string    `json:"resDate"`
}

type ResultNews struct {
	IsLocal bool         `json:"isLocal"`
	Local   *LocalNews   `json:"local"`
	Global  *gofeed.Item `json:"global"`
}

type GlobalNews struct {
	Feed *gofeed.Feed
}
