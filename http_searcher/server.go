package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

var FileName string = "dataset.xml"

type UserWrapper struct {
	ID        int    `xml:"id"`
	FirstName string `xml:"first_name"`
	LastName  string `xml:"last_name"`
	Age       int    `xml:"age"`
	About     string `xml:"about"`
	Gender    string `xml:"gender"`
}

type UsersWrapper struct {
	Version string        `xml:"version,attr"`
	List    []UserWrapper `xml:"row"`
}

func readData(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("%#v", err)
		return "", err
	}

	var inData string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inData += line
	}

	return inData, nil
}

func unmarshalUsers(inData string) ([]User, error) {
	usersWrap := new(UsersWrapper)
	err := xml.Unmarshal([]byte(inData), &usersWrap)
	if err != nil {
		fmt.Printf("%#v", err)
		return nil, err
	}
	usrs := []User{}
	for _, usrWrapped := range usersWrap.List {
		usrs = append(usrs, User{Name: usrWrapped.FirstName + usrWrapped.LastName,
			ID: usrWrapped.ID, Age: usrWrapped.Age,
			About: usrWrapped.About, Gender: usrWrapped.Gender})
	}

	return usrs, nil
}

func parseParams(r *http.Request, usrs []User) (SearchRequest, map[string]error) {
	err := map[string]error{"limit": nil, "offset": nil, "order_by": nil, "order_field": nil}
	var searchReqParams SearchRequest

	searchReqParams.Limit, err["limit"] = strconv.Atoi(r.URL.Query().Get("limit"))

	searchReqParams.Offset, err["offset"] = strconv.Atoi(r.URL.Query().Get("offset"))
	if err["offset"] == nil && (searchReqParams.Offset >= len(usrs) || searchReqParams.Offset < 0) {
		err["offset"] = fmt.Errorf("invalid offset value")
	}
	searchReqParams.OrderBy, err["order_by"] = strconv.Atoi(r.URL.Query().Get("order_by"))
	if searchReqParams.OrderBy != 0 && searchReqParams.OrderBy != 1 && searchReqParams.OrderBy != -1 {
		err["order_by"] = fmt.Errorf("incorrect order_by data")
	} // order_by - обязательный параметр

	searchReqParams.OrderField = r.URL.Query().Get("order_field")
	if searchReqParams.OrderField != "" && searchReqParams.OrderField != "Id" && searchReqParams.OrderField != "Name" &&
		searchReqParams.OrderField != "Age" {
		err["order_field"] = fmt.Errorf("incorrect order_field data")
	}

	searchReqParams.Query = r.URL.Query().Get("query")

	return searchReqParams, err
}

func sortUsrs(orderBy int, orderField string, usrs []User) []User {
	switch orderField {
	case "":
		fallthrough
	case "Name":
		sort.Slice(usrs, func(i, j int) bool {
			if orderBy > 0 {
				return usrs[i].Name < usrs[j].Name
			} else {
				return usrs[i].Name > usrs[j].Name
			}
		})
	case "Id":
		sort.Slice(usrs, func(i, j int) bool {
			if orderBy > 0 {
				return usrs[i].ID < usrs[j].ID
			} else {
				return usrs[i].ID > usrs[j].ID
			}
		})
	case "Age":
		sort.Slice(usrs, func(i, j int) bool {
			if orderBy > 0 {
				return usrs[i].Age < usrs[j].Age
			} else {
				return usrs[i].Age > usrs[j].Age
			}
		})
	}

	return usrs
}

func usrsLimitOffsetHandle(usrsOut []User, limit, offset int) []User {
	if offset >= len(usrsOut) {
		return []User{}
	}
	if limit > len(usrsOut) {
		return usrsOut[offset:]
	} else {
		return usrsOut[offset : offset+limit]
	}
}

func SearchServer(w http.ResponseWriter, r *http.Request) {
	// Auth check
	if r.Header.Get("AccessToken") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	inData, errData := readData(FileName)
	// read data success check
	if errData != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usrs, errUnmarshal := unmarshalUsers(inData)
	// unmarshall data check
	if errUnmarshal != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	searchReqParams, errParams := parseParams(r, usrs)
	// params check
	errBody := SearchErrorResponse{}
	for errName, errParam := range errParams {
		if errParam != nil {
			errBody.Error = fmt.Sprintf("error parsing %s: %v", errName, errParam)
		}
	}

	if errParams["order_field"] != nil {
		errBody.Error = ErrorBadOrderField
	}
	if errBody.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
		jsonErrParams, errParamsJSON := json.Marshal(errBody)
		if errParamsJSON != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, errWrite := w.Write(jsonErrParams)
		if errWrite != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	usrsOut := []User{}
	for _, usr := range usrs {
		if searchReqParams.Query == "" || strings.Contains(usr.Name, searchReqParams.Query) ||
			strings.Contains(usr.About, searchReqParams.Query) {
			usrsOut = append(usrsOut, usr)
		}
	}

	if searchReqParams.OrderBy != 0 {
		usrsOut = sortUsrs(searchReqParams.OrderBy, searchReqParams.OrderField, usrsOut)
	}

	usrsOut = usrsLimitOffsetHandle(usrsOut, searchReqParams.Limit, searchReqParams.Offset)

	jsonUsrs, err := json.Marshal(usrsOut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK) // даже если никого не нашли из-за query, возвращаем 200 ок
	_, errWrite := w.Write(jsonUsrs)
	if errWrite != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
