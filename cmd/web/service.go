package main

import (
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"io"
	"net/http"
	"sort"
	"strings"
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

// Создание JWT ключа с данными пользователя(логин, почта, время создания и время истечения срока). Необходимо для
// авторизации пользователя на сайте.
func (app *application) createJWT(userID string) (string, error) {
	userEmail, err := app.dbSearch("mail", "user", fmt.Sprintf("where login = '%s'", userID))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"user_email": userEmail[0],
		"iat":        time.Now().Unix(),                     // Время выдачи токена.
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Время истечения срока действия токена.
	})

	// Подписать токен с использованием секретного ключа.
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Верификация JWT ключа.
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

// При запросе на любую страницу проверяется JWT ключ из localStorage на стороне пользователя для отображения
// пользователя на странице.
func (app *application) protected(w http.ResponseWriter, r *http.Request) {
	result := struct {
		User    string `json:"user"`
		Mail    string `json:"mail"`
		Success bool   `json:"success"`
	}{Success: false}
	// Получить токен из заголовка запроса.
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Проверить токен.
	claims, err := app.verifyJWT(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Получить идентификатор пользователя из токена.
	userID, ok := claims["user_id"].(string)
	userEmail, ok := claims["user_email"].(string)
	if ok && userID != "" {
		result.User = userID
		result.Mail = userEmail
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

// Обработка новостной ленты РИА Новости.
func (app *application) rssParse(w http.ResponseWriter) *RssNews {
	url1 := "https://ria.ru/export/rss2/archive/index.xml"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		app.serverError(w, err)
		return nil
	}

	req.Header.Set("X-Forwarded-For", "128.204.79.211")
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка при выполнении запроса к РИА новости", http.StatusInternalServerError)
		app.serverError(w, err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Ошибка при чтении тела ответа", http.StatusInternalServerError)
		app.serverError(w, err)
		return nil
	}

	var rss RssNews
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		http.Error(w, "Ошибка при декодировании xml", http.StatusInternalServerError)
		app.serverError(w, err)
		return nil
	}
	for i, r := range rss.Channel.Item {
		r.Id = (i + 1) * -1
	}
	return &rss
}

// Добавление новости в бд.
func (app *application) addNew(w http.ResponseWriter, r *http.Request) {
	type jsonData struct {
		Title string `json:"title"`
		Text  string `json:"text"`
		User  string `json:"user_id"`
	}

	var data jsonData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Ошибка при декодировании запроса", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	type successAnsw struct {
		Success bool `json:"success"`
	}

	var success successAnsw

	_, err = app.db.Query(fmt.Sprintf("INSERT INTO `news` "+
		"(`title`, `text`, `user_id`) VALUES('%s', '%s', '%s')", data.Title, data.Text, data.User))
	if err != nil {
		app.serverError(w, err)
		return
	} else {
		success = successAnsw{Success: true}
	}

	answ, err := json.Marshal(success)
	if err != nil {
		http.Error(w, "Ошибка при кодировании ответа JSON", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(answ)
}

// Получение новостей с БД с преображением в XML данные для объединения с новостями с "РИА Новости"
func (app *application) getLocalNews() ([]LocalNews, error) {
	getData, err := app.db.Query("SELECT * FROM `news`")
	if err != nil {
		return nil, err
	}
	defer getData.Close()
	news := make([]LocalNews, 0)
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %w", err)
	}
	for getData.Next() {
		var localNew LocalNews
		var createdAtString string
		err := getData.Scan(&localNew.Id, &localNew.Title, &localNew.Text, &localNew.User, &createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		t, err := time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error parsing datetime: %w", err)
		}

		localNew.Date = t.In(loc).Format("2006-01-02 15:04:05")

		news = append(news, localNew)
	}
	if err := getData.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %w", err)
	}
	return news, nil
}

// layout определяет формат, который соответствует входной строке
func ConvertToTime(dateString string, layout string) (time.Time, error) {
	// Парсим строку в тип time.Time
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err // Возвращаем пустое значение time.Time и ошибку
	}

	return parsedTime, nil // Возвращаем распарсенное время и nil для ошибки
}

// Проработать смешанную структуру либо функцию, которая будет сортировать новости по датам на фронте
func (app *application) mixNews(local []LocalNews, ria []Item) ([]LocalNews, []Item, error) {
	resultLocal := make([]LocalNews, 0)
	resultRia := make([]Item, 0)
	forSort := make([]ForSortNews, 0)
	for _, r := range local {
		var a ForSortNews
		localTime, err := ConvertToTime(r.Date, "2006-01-02 15:04:05")
		if err != nil {
			err = fmt.Errorf("ошибка при конвертации времени из дб: %w", err)
			return nil, nil, err
		}
		a.NewsId = r.Id
		a.Date = localTime
		forSort = append(forSort, a)
	}
	for _, r := range ria {
		var a ForSortNews
		riaTime, err := ConvertToTime(r.PubDate, "Mon, 02 Jan 2006 15:04:05 -0700")
		if err != nil {
			err = fmt.Errorf("ошибка при конвертации времени риа: %w", err)
			return nil, nil, err
		}

		a.NewsId = r.Id
		a.Date = riaTime
		forSort = append(forSort, a)
	}
	for _, r := range local {
		resultLocal = append(resultLocal, r)
	}
	for _, r := range ria {
		resultRia = append(resultRia, r)
	}
	sort.Slice(forSort, func(i, j int) bool {
		return forSort[i].Date.After(forSort[j].Date)
	})
	mp := make(map[int]int)
	for i, r := range forSort {
		mp[r.NewsId] = i
	}
	sort.Slice(resultLocal, func(i, j int) bool {
		return mp[resultLocal[i].Id] < mp[resultLocal[j].Id]
	})
	sort.Slice(resultRia, func(i, j int) bool {
		return mp[resultRia[i].Id] > mp[resultRia[j].Id]
	})
	return resultLocal, resultRia, nil
}

func (app *application) indexNews(w http.ResponseWriter, r *http.Request) {
	local, err := app.getLocalNews()
	if err != nil {
		app.serverError(w, err)
		return
	}
	ria := app.rssParse(w)
	items := make([]Item, 0)
	for _, r := range ria.Channel.Item {
		items = append(items, r)
	}
	resultLocal, resultRia, err := app.mixNews(local, items)
	if err != nil {
		app.serverError(w, err)
		return
	}
	type Response struct {
		Local []LocalNews
		Ria   []Item
	}
	res, err := json.Marshal(Response{Local: resultLocal, Ria: resultRia})
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (app *application) handleVisit(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("user_id")
	if errors.Is(err, http.ErrNoCookie) {
		userID := uuid.New().String()
		_, err = app.db.Exec("INSERT INTO user_visits (user_id) VALUES (?)", userID)
		if err != nil {
			return nil, fmt.Errorf("ошибка при добавлении нового cookie в БД: %w", err)
		}
		cookie = &http.Cookie{
			Name:     "user_id",
			Value:    userID,
			HttpOnly: true,
			Path:     "/",
		}
	} else if err != nil {
		return nil, fmt.Errorf("ошибка при чтении cookie: %w", err)
	}
	return cookie, nil
}

func (app *application) gotNew(newsID int) (*LocalNews, error) {
	raw, err := app.db.Query("SELECT * FROM news WHERE id=?", newsID)
	if err != nil {
		return nil, fmt.Errorf("error bad db request: %w", err)
	}
	var gotNew LocalNews
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %w", err)
	}
	var createdAtString string
	for raw.Next() {
		err = raw.Scan(&gotNew.Id, &gotNew.Title, &gotNew.Text, &gotNew.User, &createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error bad response from db: %w", err)
		}
		t, err := time.Parse("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error parsing datetime: %w", err)
		}

		gotNew.Date = t.In(loc).Format("2006-01-02 15:04:05")
	}
	return &gotNew, nil
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	local, err := app.getLocalNews()
	if err != nil {
		app.serverError(w, err)
		return
	}
	ria := app.rssParse(w)
	items := make([]Item, 0)
	for _, r := range ria.Channel.Item {
		items = append(items, r)
	}
	resultLocal, resultRia, err := app.mixNews(local, items)
	if err != nil {
		app.serverError(w, err)
		return
	}
	type Response struct {
		Local []LocalNews
		Ria   []Item
	}
	json, err := json.Marshal(Response{Local: resultLocal, Ria: resultRia})
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
