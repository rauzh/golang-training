package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

// тут писать код тестов

type TestCase struct {
	TestName    string
	Result      *SearchResponse
	IsNegative  bool
	AccessToken string
	Request     SearchRequest
	FileName    string
}

func TestSearchOK(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "OK UP TO LIMIT NAME SORT 1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     15,
						Name:   "AllisonValdez",
						Age:    21,
						About:  "Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.",
						Gender: "male",
					},
					{
						ID:     16,
						Name:   "AnnieOsborn",
						Age:    35,
						About:  "Consequat fugiat veniam commodo nisi nostrud culpa pariatur. Aliquip velit adipisicing dolor et nostrud. Eu nostrud officia velit eiusmod ullamco duis eiusmod ad non do quis.",
						Gender: "female",
					},
					{
						ID:     19,
						Name:   "BellBauer",
						Age:    26,
						About:  "Nulla voluptate nostrud nostrud do ut tempor et quis non aliqua cillum in duis. Sit ipsum sit ut non proident exercitation. Quis consequat laboris deserunt adipisicing eiusmod non cillum magna.",
						Gender: "male",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      3,
				Offset:     0,
				Query:      "",
				OrderBy:    1,
				OrderField: "Name",
			},
		},
		{
			TestName: "OK UP TO LIMIT NAME SORT -1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     13,
						Name:   "WhitleyDavidson",
						Age:    40,
						About:  "Consectetur dolore anim veniam aliqua deserunt officia eu. Et ullamco commodo ad officia duis ex incididunt proident consequat nostrud proident quis tempor. Sunt magna ad excepteur eu sint aliqua eiusmod deserunt proident. Do labore est dolore voluptate ullamco est dolore excepteur magna duis quis. Quis laborum deserunt ipsum velit occaecat est laborum enim aute. Officia dolore sit voluptate quis mollit veniam. Laborum nisi ullamco nisi sit nulla cillum et id nisi.",
						Gender: "male",
					},
					{
						ID:     33,
						Name:   "TwilaSnow",
						Age:    36,
						About:  "Sint non sunt adipisicing sit laborum cillum magna nisi exercitation. Dolore officia esse dolore officia ea adipisicing amet ea nostrud elit cupidatat laboris. Proident culpa ullamco aute incididunt aute. Laboris et nulla incididunt consequat pariatur enim dolor incididunt adipisicing enim fugiat tempor ullamco. Amet est ullamco officia consectetur cupidatat non sunt laborum nisi in ex. Quis labore quis ipsum est nisi ex officia reprehenderit ad adipisicing fugiat. Labore fugiat ea dolore exercitation sint duis aliqua.",
						Gender: "female",
					},
					{
						ID:     18,
						Name:   "TerrellHall",
						Age:    27,
						About:  "Ut nostrud est est elit incididunt consequat sunt ut aliqua sunt sunt. Quis consectetur amet occaecat nostrud duis. Fugiat in irure consequat laborum ipsum tempor non deserunt laboris id ullamco cupidatat sit. Officia cupidatat aliqua veniam et ipsum labore eu do aliquip elit cillum. Labore culpa exercitation sint sint.",
						Gender: "male",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      3,
				Offset:     0,
				Query:      "",
				OrderBy:    -1,
				OrderField: "Name",
			},
		},
		{
			TestName: "OK UP TO LIMIT EMPTY SORT -1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     13,
						Name:   "WhitleyDavidson",
						Age:    40,
						About:  "Consectetur dolore anim veniam aliqua deserunt officia eu. Et ullamco commodo ad officia duis ex incididunt proident consequat nostrud proident quis tempor. Sunt magna ad excepteur eu sint aliqua eiusmod deserunt proident. Do labore est dolore voluptate ullamco est dolore excepteur magna duis quis. Quis laborum deserunt ipsum velit occaecat est laborum enim aute. Officia dolore sit voluptate quis mollit veniam. Laborum nisi ullamco nisi sit nulla cillum et id nisi.",
						Gender: "male",
					},
					{
						ID:     33,
						Name:   "TwilaSnow",
						Age:    36,
						About:  "Sint non sunt adipisicing sit laborum cillum magna nisi exercitation. Dolore officia esse dolore officia ea adipisicing amet ea nostrud elit cupidatat laboris. Proident culpa ullamco aute incididunt aute. Laboris et nulla incididunt consequat pariatur enim dolor incididunt adipisicing enim fugiat tempor ullamco. Amet est ullamco officia consectetur cupidatat non sunt laborum nisi in ex. Quis labore quis ipsum est nisi ex officia reprehenderit ad adipisicing fugiat. Labore fugiat ea dolore exercitation sint duis aliqua.",
						Gender: "female",
					},
					{
						ID:     18,
						Name:   "TerrellHall",
						Age:    27,
						About:  "Ut nostrud est est elit incididunt consequat sunt ut aliqua sunt sunt. Quis consectetur amet occaecat nostrud duis. Fugiat in irure consequat laborum ipsum tempor non deserunt laboris id ullamco cupidatat sit. Officia cupidatat aliqua veniam et ipsum labore eu do aliquip elit cillum. Labore culpa exercitation sint sint.",
						Gender: "male",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      3,
				Offset:     0,
				Query:      "",
				OrderBy:    -1,
				OrderField: "",
			},
		},
		{
			TestName: "OK UP TO LIMIT ID SORT -1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     16,
						Name:   "AnnieOsborn",
						Age:    35,
						About:  "Consequat fugiat veniam commodo nisi nostrud culpa pariatur. Aliquip velit adipisicing dolor et nostrud. Eu nostrud officia velit eiusmod ullamco duis eiusmod ad non do quis.",
						Gender: "female",
					},
					{
						ID:     15,
						Name:   "AllisonValdez",
						Age:    21,
						About:  "Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.",
						Gender: "male",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      2,
				Offset:     18,
				Query:      "",
				OrderBy:    -1,
				OrderField: "Id",
			},
		},
		{
			TestName: "OK UP TO LIMIT ID SORT 1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     15,
						Name:   "AllisonValdez",
						Age:    21,
						About:  "Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.",
						Gender: "male",
					},
					{
						ID:     16,
						Name:   "AnnieOsborn",
						Age:    35,
						About:  "Consequat fugiat veniam commodo nisi nostrud culpa pariatur. Aliquip velit adipisicing dolor et nostrud. Eu nostrud officia velit eiusmod ullamco duis eiusmod ad non do quis.",
						Gender: "female",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      2,
				Offset:     15,
				Query:      "",
				OrderBy:    1,
				OrderField: "Id",
			},
		},
		{
			TestName: "OK UP TO LIMIT AGE SORT 1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     1,
						Name:   "HildaMayer",
						Age:    21,
						About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.",
						Gender: "female",
					},
					{
						ID:     15,
						Name:   "AllisonValdez",
						Age:    21,
						About:  "Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.",
						Gender: "male",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      2,
				Offset:     0,
				Query:      "",
				OrderBy:    1,
				OrderField: "Age",
			},
		},
		{
			TestName: "OK UP TO LIMIT AGE SORT -1",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     15,
						Name:   "AllisonValdez",
						Age:    21,
						About:  "Labore excepteur voluptate velit occaecat est nisi minim. Laborum ea et irure nostrud enim sit incididunt reprehenderit id est nostrud eu. Ullamco sint nisi voluptate cillum nostrud aliquip et minim. Enim duis esse do aute qui officia ipsum ut occaecat deserunt. Pariatur pariatur nisi do ad dolore reprehenderit et et enim esse dolor qui. Excepteur ullamco adipisicing qui adipisicing tempor minim aliquip.",
						Gender: "male",
					},
					{
						ID:     1,
						Name:   "HildaMayer",
						Age:    21,
						About:  "Sit commodo consectetur minim amet ex. Elit aute mollit fugiat labore sint ipsum dolor cupidatat qui reprehenderit. Eu nisi in exercitation culpa sint aliqua nulla nulla proident eu. Nisi reprehenderit anim cupidatat dolor incididunt laboris mollit magna commodo ex. Cupidatat sit id aliqua amet nisi et voluptate voluptate commodo ex eiusmod et nulla velit.",
						Gender: "female",
					},
				},

				NextPage: true,
			},
			Request: SearchRequest{
				Limit:      2,
				Offset:     33,
				Query:      "",
				OrderBy:    -1,
				OrderField: "Age",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchFileOsError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "OS ERROR NO FILE",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "anatoliy_vasserman.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchFileUnMarshError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR BAD XML",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "ruben_buben.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchNoAccessTokenError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR NO ACCESS TOKEN",

			AccessToken: "",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchNoServerError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR NO SERVER",

			AccessToken: "",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	for _, testCase := range cases {
		testClient := &SearchClient{
			AccessToken: testCase.AccessToken,
			URL:         "",
		}

		FileName = testCase.FileName
		testResp, err := testClient.FindUsers(testCase.Request)

		if err == nil && testCase.IsNegative {
			t.Errorf("[%s] expected error, got nil", testCase.TestName)
		}
		if err != nil && !testCase.IsNegative {
			t.Errorf("[%s] unexpected error: %#v", testCase.TestName, err)
		}
		if !reflect.DeepEqual(testCase.Result, testResp) {
			t.Errorf("[%s] wrong result, expected %#v, got %#v", testCase.TestName, testCase.Result, testResp)
		}
	}
}

func TestSearchUnpackOKBodyError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR CANT UNPACK RESPONSE BODY JSON",

			AccessToken: "",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		jsonData, jsonDataErr := json.Marshal("err: buben")
		if jsonDataErr != nil {
			fmt.Print("error marshalling json")
		}
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			fmt.Print("error writing json")
		}
	})
	ts := httptest.NewServer(responseHandler)
	defer ts.Close()

	testClient := &SearchClient{
		AccessToken: cases[0].AccessToken,
		URL:         ts.URL,
	}

	FileName = cases[0].FileName
	testResp, err := testClient.FindUsers(cases[0].Request)

	if err == nil && cases[0].IsNegative {
		t.Errorf("[%s] expected error, got nil", cases[0].TestName)
	}
	if err != nil && !cases[0].IsNegative {
		t.Errorf("[%s] unexpected error: %#v", cases[0].TestName, err)
	}
	if !reflect.DeepEqual(cases[0].Result, testResp) {
		t.Errorf("[%s] wrong result, expected %#v, got %#v", cases[0].TestName, cases[0].Result, testResp)
	}
}

func TestSearchUnpackBadRequestBodyError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR CANT UNPACK BAD REQUEST RESPONSE BODY JSON",

			AccessToken: "",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		jsonData, jsonDataErr := json.Marshal("err: buben")
		if jsonDataErr != nil {
			fmt.Print("error marshalling json")
		}
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			fmt.Print("error writing json")
		}
	})
	ts := httptest.NewServer(responseHandler)
	defer ts.Close()

	testClient := &SearchClient{
		AccessToken: cases[0].AccessToken,
		URL:         ts.URL,
	}

	FileName = cases[0].FileName
	testResp, err := testClient.FindUsers(cases[0].Request)

	if err == nil && cases[0].IsNegative {
		t.Errorf("[%s] expected error, got nil", cases[0].TestName)
	}
	if err != nil && !cases[0].IsNegative {
		t.Errorf("[%s] unexpected error: %#v", cases[0].TestName, err)
	}
	if !reflect.DeepEqual(cases[0].Result, testResp) {
		t.Errorf("[%s] wrong result, expected %#v, got %#v", cases[0].TestName, cases[0].Result, testResp)
	}
}

func TestSearchTimeoutError(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR CANT UNPACK RESPONSE BODY JSON",

			AccessToken: "",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      3,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	})
	ts := httptest.NewServer(responseHandler)
	defer ts.Close()

	testClient := &SearchClient{
		AccessToken: cases[0].AccessToken,
		URL:         ts.URL,
	}

	FileName = cases[0].FileName
	testResp, err := testClient.FindUsers(cases[0].Request)

	if err == nil && cases[0].IsNegative {
		t.Errorf("[%s] expected error, got nil", cases[0].TestName)
	}
	if err != nil && !cases[0].IsNegative {
		t.Errorf("[%s] unexpected error: %#v", cases[0].TestName, err)
	}
	if !reflect.DeepEqual(cases[0].Result, testResp) {
		t.Errorf("[%s] wrong result, expected %#v, got %#v", cases[0].TestName, cases[0].Result, testResp)
	}
}

func TestSearchParamErrors(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR LIMIT & OFFSET PARAM OUT OF RANGE",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "small_dataset.xml",
			Request: SearchRequest{
				Limit:      10,
				Offset:     8,
				Query:      "CHAPO",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
		{
			TestName: "ERROR ORDER_BY PARAM INVALID",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      10,
				Offset:     8,
				Query:      "",
				OrderBy:    9,
				OrderField: "Name",
			},
		},
		{
			TestName: "ERROR ORDER_FIELD PARAM INVALID",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      10,
				Offset:     8,
				Query:      "",
				OrderBy:    0,
				OrderField: "Buba",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchLimitOffsetErrors(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "ERROR LIMIT < 0",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "small_dataset.xml",
			Request: SearchRequest{
				Limit:      -10,
				Offset:     8,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
		{
			TestName: "ERROR OFFSET < 0",

			AccessToken: "123",
			IsNegative:  true,
			FileName:    "dataset.xml",
			Request: SearchRequest{
				Limit:      10,
				Offset:     -8,
				Query:      "",
				OrderBy:    9,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchBigLimitOk(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "OK LIMIT > 25",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{
					{
						ID:     10,
						Name:   "HendersonMaxwell",
						Age:    30,
						About:  "Ex et excepteur anim in eiusmod. Cupidatat sunt aliquip exercitation velit minim aliqua ad ipsum cillum dolor do sit dolore cillum. Exercitation eu in ex qui voluptate fugiat amet.",
						Gender: "male",
					},
				},

				NextPage: false,
			},
			Request: SearchRequest{
				Limit:      30,
				Offset:     0,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func TestSearchBigOffsetOk(t *testing.T) {
	cases := []TestCase{
		{
			TestName: "OK OFFSET > LEN USRS OUT",

			AccessToken: "123",
			IsNegative:  false,
			FileName:    "dataset.xml",
			Result: &SearchResponse{
				Users: []User{},

				NextPage: false,
			},
			Request: SearchRequest{
				Limit:      10,
				Offset:     4,
				Query:      "HendersonMaxwell",
				OrderBy:    0,
				OrderField: "Name",
			},
		},
	}

	runTestCases(t, cases)
}

func runTestCases(t *testing.T, cases []TestCase) {
	ts := httptest.NewServer(http.HandlerFunc(SearchServer))
	defer ts.Close()

	for _, testCase := range cases {
		testClient := &SearchClient{
			AccessToken: testCase.AccessToken,
			URL:         ts.URL,
		}

		FileName = testCase.FileName
		testResp, err := testClient.FindUsers(testCase.Request)

		if err == nil && testCase.IsNegative {
			t.Errorf("[%s] expected error, got nil", testCase.TestName)
		}
		if err != nil && !testCase.IsNegative {
			t.Errorf("[%s] unexpected error: %#v", testCase.TestName, err)
		}
		if !reflect.DeepEqual(testCase.Result, testResp) {
			t.Errorf("[%s] wrong result, expected %#v, got %#v", testCase.TestName, testCase.Result, testResp)
		}
	}
}
