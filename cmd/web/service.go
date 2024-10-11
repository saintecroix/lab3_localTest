package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
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

func (app *application) dbSearch(resoultColomns, table, criterion string) ([]string, error) {
	sel := fmt.Sprintf("select %s from %s ", resoultColomns, table)
	if criterion != "" {
		sel = sel + criterion
	}
	result := make([]string, 0)
	req, err := app.db.Query(sel)
	if err != nil {
		return result, err
	}
	defer req.Close()
	for req.Next() {
		a := ""
		err := req.Scan(&a)
		if err != nil {
			return result, err
		}
		result = append(result, a)
	}
	return result, nil
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	type sendData struct {
		Result string `json:"result"`
	}
	result := &sendData{}
	mail := "z"
	res, err := app.dbSearch("mail", "user", fmt.Sprintf("where login = '%s'", mail))
	if err != nil {
		app.serverError(w, err)
		return
	}
	if len(res) != 0 {
		result.Result = res[0]
	}

	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Ошибка при кодировании данных в JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func createJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now().Unix(),                     // Время выдачи токена.
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Время истечения срока действия токена.
	})

	// Подписать токен с использованием секретного ключа.
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (app *application) verifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убедиться, что используется правильный алгоритм подписи.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Получить секретный ключ для проверки подписи.
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return nil, err
	}

	// Проверить целостность токена.
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Получить данные о пользователе из токена.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (app *application) protected(w http.ResponseWriter, r *http.Request) {
	result := struct {
		Success bool `json:"success"`
	}{Success: false}
	// Получить токен из заголовка запроса.
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Проверить токен.
	claims, _ := app.verifyJWT(tokenString)

	// Получить идентификатор пользователя из токена.
	userID := claims["user_id"].(string)
	if userID != "" {
		result.Success = true
	}

	// Преобразовать структуру в JSON и отправить клиенту.
	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
