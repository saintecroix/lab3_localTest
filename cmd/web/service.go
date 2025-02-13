package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
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

func jsonRequest(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(target)
}

func (app *application) jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonResponse)
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

// Получение локальной новости из БД для отображения ее на отдельной странице.
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

		gotNew.Date = t.In(loc)
	}
	return &gotNew, nil
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

// Получение новостей с БД для объединения с новостями с "РИА Новости"
func (app *application) getLocalNews() ([]LocalNews, error) {
	getData, err := app.db.Query("SELECT * FROM `news`")
	if err != nil {
		return nil, err
	}
	defer getData.Close()
	news := make([]LocalNews, 0)
	for getData.Next() {
		var localNew LocalNews
		var createdAtString string
		err := getData.Scan(&localNew.Id, &localNew.Title, &localNew.Text, &localNew.User, &createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		t, err := ConvertToTime("2006-01-02 15:04:05", createdAtString)
		if err != nil {
			return nil, fmt.Errorf("error parsing datetime: %w", err)
		}

		localNew.Date = t
		localNew.ResDate, err = FromTimeToString(t)
		if err != nil {
			return nil, fmt.Errorf("error converting datetime: %w", err)
		}

		news = append(news, localNew)
	}
	if err := getData.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %w", err)
	}
	return news, nil
}

// layout определяет формат, который соответствует входной строке.
// Возвращает время по Москве.
func ConvertToTime(layout string, dateString string) (time.Time, error) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load location: %w", err)
	}
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime.In(loc), nil
}

func FromTimeToString(t time.Time) (string, error) {
	layout := "2006-01-02 15:04:05"
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", fmt.Errorf("failed to load location: %w", err)
	}
	t = t.In(loc)
	return t.Format(layout), nil
}

// Возвращает список подписок пользователя.
func (app *application) getSources(user string) (string, error) {
	getData, err := app.db.Query("SELECT `news_sources` FROM `user` WHERE login = ?", user)
	if err != nil {
		return "", fmt.Errorf("error bad db request: %w", err)
	}
	defer getData.Close() // Важно закрывать rows после использования

	var sources string
	for getData.Next() {
		var sourcesPtr *string          // Использовать указатель на string
		err = getData.Scan(&sourcesPtr) // Сканировать в указатель
		if err != nil {
			return "", fmt.Errorf("error bad response from db: %w", err)
		}

		if sourcesPtr != nil { // Проверить на NULL
			sources = *sourcesPtr // Разыменовать указатель, если не NULL
		} else {
			sources = "" // Если NULL, то sources = ""
		}
	}

	err = getData.Err() // Проверить на ошибки после итерации
	if err != nil {
		return "", fmt.Errorf("error after iteration: %w", err)
	}

	return sources, nil
}

func (app *application) deleteSources(user string, sources []string) error {
	userExist, err := app.userExists(user)
	if err != nil {
		return fmt.Errorf("error bad db request: %w", err)
	}
	if !userExist {
		return fmt.Errorf("user %s does not exist", user)
	}

	oldSources, err := app.getSources(user)
	if err != nil {
		return fmt.Errorf("error bad db request: %w", err)
	}

	mp := make(map[string]bool)
	for _, s := range sources {
		mp[s] = true
	}

	old := strings.Split(oldSources, ",")
	result := make([]string, 0)
	for _, s := range old {
		s = strings.TrimSpace(s)
		if s != "" {
			if _, ok := mp[s]; !ok {
				result = append(result, s)
			}
		}
	}

	tx, err := app.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("Failed to rollback transaction: %v\n", rbErr)
			}
			return
		}
	}()

	_, err = tx.Exec("UPDATE user SET news_sources = NULL WHERE login = ?", user)
	if err != nil {
		return fmt.Errorf("error setting news_sources to NULL: %w", err)
	}

	for _, source := range result {
		source = strings.TrimSpace(source)
		if source != "" {
			_, err = tx.Exec(
				"UPDATE `user` SET `news_sources` = CASE WHEN `news_sources` IS NULL THEN ? ELSE CONCAT(`news_sources`, ',', ?) END WHERE `login` = ?", source, source, user)
			if err != nil {
				return fmt.Errorf("error updating news_sources with source %s: %w", source, err)
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func rssFeed(link string) (news *gofeed.Feed, err error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(link)
	if err != nil {
		return nil, err
	}
	return feed, nil
}

func (app *application) userExists(login string) (bool, error) {
	var count int
	err := app.db.QueryRow("SELECT COUNT(*) FROM `user` WHERE `login` = ?", login).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error querying user: %w", err)
	}
	return count > 0, nil
}

func (app *application) sourceValid(link string) bool {
	news, err := rssFeed(link)
	if err != nil {
		return false
	}
	var empty *gofeed.Feed
	if news == empty {
		return false
	}
	return true
}

func (app *application) sortNews(feed []ResultNews) []ResultNews {
	sort.Slice(feed, func(i, j int) bool {
		var date1 time.Time
		var date2 time.Time

		if feed[i].IsLocal {
			date1 = feed[i].Local.Date
		} else if feed[i].Global.PublishedParsed != nil {
			date1 = *feed[i].Global.PublishedParsed
		} else {
			date1 = time.Time{}
		}

		if feed[j].IsLocal {
			date2 = feed[j].Local.Date
		} else if feed[j].Global.PublishedParsed != nil {
			date2 = *feed[j].Global.PublishedParsed
		} else {
			date2 = time.Time{}
		}
		return date1.After(date2)
	})
	return feed
}

func (app *application) indexNews(w http.ResponseWriter, r *http.Request) {
	type JsonData struct {
		User string `json:"user,omitempty"`
	}
	type Response struct {
		Feed  []ResultNews `json:"feed"`
		Error interface{}  `json:"error,omitempty"`
	}

	var data JsonData
	var response Response
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		app.logError(err)
		response.Error = fmt.Errorf("failed to load location: %w", err)
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)

	// Ошибка декодирования не является критической, обрабатываем ее как отсутствие данных
	if err != nil {
		app.logError(err)
		fmt.Println("Error decoding JSON:", err)
	}

	local, err := app.getLocalNews()
	if err != nil {
		app.logError(err)
		response.Error = err
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	for _, l := range local {
		l.ResDate, err = FromTimeToString(l.Date)
		if err != nil {
			app.logError(err)
			response.Error = err
			app.jsonResponse(w, http.StatusInternalServerError, response)
			return
		}
		response.Feed = append(response.Feed, ResultNews{IsLocal: true, Local: &l})
	}

	ria, err := rssFeed("https://ria.ru/export/rss2/index.xml")
	if err != nil {
		app.logError(err)
		response.Error = err
		app.jsonResponse(w, http.StatusInternalServerError, response)
		return
	}
	for _, v := range ria.Items {
		if v.PublishedParsed != nil {
			*v.PublishedParsed = v.PublishedParsed.In(loc)
			v.Published, err = FromTimeToString(*v.PublishedParsed)
			v.GUID = ria.Title
			if err != nil {
				app.logError(err)
				response.Error = err
				app.jsonResponse(w, http.StatusInternalServerError, response)
				return
			}
		}
		response.Feed = append(response.Feed, ResultNews{IsLocal: false, Global: v})
	}
	if data.User != "" {
		getSources, err := app.getSources(data.User)
		if err != nil {
			app.logError(err)
			response.Error = err
			app.jsonResponse(w, http.StatusInternalServerError, response)
			return
		}
		if getSources != "" {
			sources := strings.Split(getSources, ",")
			for _, v := range sources {
				if !app.sourceValid(v) {
					app.logError(err)
					response.Error = fmt.Errorf("invalid source: %v", v)
					app.jsonResponse(w, http.StatusBadRequest, response)
					return
				}
				feed, err := rssFeed(v)
				if err != nil {
					app.logError(err)
					response.Error = err
					app.jsonResponse(w, http.StatusInternalServerError, response)
					return
				}
				for _, f := range feed.Items {
					if f.PublishedParsed != nil {
						*f.PublishedParsed = f.PublishedParsed.In(loc)
						f.Published, err = FromTimeToString(*f.PublishedParsed)
						f.GUID = feed.Title
						if err != nil {
							app.logError(err)
							response.Error = err
							app.jsonResponse(w, http.StatusBadRequest, response)
							return
						}
					}
					response.Feed = append(response.Feed, ResultNews{IsLocal: false, Global: f})
				}
			}
		}
	}

	response.Feed = app.sortNews(response.Feed)
	app.jsonResponse(w, http.StatusOK, response)
	return
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {

}
