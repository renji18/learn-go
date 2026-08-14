package main

import (
	"fmt"
)

// type course struct {
// 	Name     string   `json:"course_name"`
// 	Price    int      `json:"price"`
// 	Platform string   `json:"website"`
// 	Password string   `json:"-"`
// 	Tags     []string `json:"tags,omitempty"`
// }

func main() {
	fmt.Println("Welcome to a class of Golang")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	// myNum := 23
	// var ptr = &myNum
	// fmt.Println(myNum, ptr, *ptr)

	// *ptr = *ptr * 2
	// fmt.Println(myNum, ptr, *ptr)

	// var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	// var index int = 2
	// fmt.Println(courses)
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

	// 	for i := range courses {
	// 		fmt.Println(i)

	// 		if i == 2 {
	// 			goto lco
	// 		}
	// 	}

	// lco:
	// 	func() {
	// 		fmt.Println("In lco: ", courses)
	// 	}()

	// content := "This needs to go in a file"
	// fileName := "./mylcogofile.txt"

	// file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	// if err != nil {
	// 	fmt.Println("Error opening file", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// length, err := io.WriteString(file, content)
	// if err != nil {
	// 	fmt.Println("Error writing file", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("length is: ", length)

	// readFile(fileName)
	// func readFile(file string) {
	// 	fileByte, err := os.ReadFile(file)
	// 	if err != nil {
	// 		fmt.Println("Error reading file", err)
	// 	}
	// 	fmt.Println(string(fileByte))
	// }

	// const myurl string = "https://www.renjiriverstone.dev:6789/random?abcd=efgh&ijkl=mnop&qrst=uvwx&abcd=yz"
	// res, err := http.Get(myurl)
	// if err != nil {
	// 	panic(err)
	// }

	// defer res.Body.Close()

	// fmt.Printf("Response is of type: %T\n", res)

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(bodyBytes))

	// result, err := url.Parse(myurl)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Type of result is %T\n", result)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Query())

	// qparams := result.Query()

	// for _, val := range qparams {
	// 	fmt.Println("Param is ", val)
	// }

	// q := url.Values{}
	// q.Add("a", "b")
	// q.Add("a", "c")
	// q.Add("d", "e")

	// createUrl := url.URL{
	// 	Scheme: "https",
	// 	Host:   "renjiriverstone.dev:8000", // port does not have it's own specific field, it comes in the host itself
	// 	Path:   "/report",
	// 	// RawQuery: "user=renji",
	// 	RawQuery: q.Encode(),
	// }

	// fmt.Println(createUrl.String(), createUrl.Port())

	// const myUrl = "http://localhost:8000/"

	// performGetRequest(myUrl)
	// performPostJsonRequest(myUrl + "post")
	// performPostFormRequest(myUrl + "postform")
	// DecodeJson()

	// greeter()

	// r := mux.NewRouter()
	// r.HandleFunc("/", serveHome).Methods("GET")

	// log.Fatal(http.ListenAndServe(":4000", r))

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4
	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// random number
	// fmt.Println(rand.Intn(5) + 1)
	// num, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(num)
}

// func performGetRequest(url string) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer res.Body.Close()

// 	bodyBytes, err := io.ReadAll(res.Body)
// 	// fmt.Println(string(bodyBytes)) // Direct method
// 	if err != nil {
// 		panic(err)
// 	}

// 	if res.StatusCode != 200 {
// 		panic(fmt.Errorf("API not successful. StatusCode: %d\n", res.StatusCode))
// 	}

// 	var responseString strings.Builder
// 	_, err = responseString.Write(bodyBytes)

// 	if err != nil {
// 		panic(fmt.Errorf("Error writing bodyBytes to responseString: %w\n", err))
// 	}

// 	fmt.Println("API successful")
// 	fmt.Println(responseString.String())
// }

// func performPostJsonRequest(url string) {
// 	// fake json payload
// 	requestBody := strings.NewReader(`
// 		{
// 			"courseName": "Let's go with golang",
// 			"price": 0,
// 			"platform": "renjiriverstone.dev"
// 		}
// 		`)

// 	res, err := http.Post(url, "application/json", requestBody)

// 	if err != nil {
// 		panic(fmt.Errorf("Error in POST request: %w\n", err))
// 	}
// 	defer res.Body.Close()

// 	bodyBytes, err := io.ReadAll(res.Body)
// 	var responseString strings.Builder

// 	_, err = responseString.Write(bodyBytes)
// 	if err != nil {
// 		panic(fmt.Errorf("Error writing bodyBytes to responseString: %w\n", err))
// 	}

// 	fmt.Println(responseString.String())
// }

// func performPostFormRequest(myUrl string) {
// 	q := url.Values{}
// 	q.Add("firstName", "Aadarsh")
// 	q.Add("lastName", "Jha")
// 	q.Add("email", "aadarshjha1401@gmail.com")

// 	res, err := http.PostForm(myUrl, q)
// 	if err != nil {
// 		panic(fmt.Errorf("Error in postform api: %w\n", err))
// 	}
// 	defer res.Body.Close()

// 	bodyBytes, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		panic(fmt.Errorf("Error reading body: %w\n", err))
// 	}

// 	var responseString strings.Builder
// 	_, err = responseString.Write(bodyBytes)
// 	if err != nil {
// 		panic(fmt.Errorf("Error writing bodyBytes to responseString: %w\n", err))
// 	}

// 	fmt.Println(responseString.String())
// }

// func EncodeJson() {
// 	lcoCourses := []course{
// 		{"ReactJs BootCamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
// 		{"MERN BootCamp", 199, "LearnCodeOnline.in", "def123", []string{"full-stack", "js"}},
// 		{"Angular BootCamp", 299, "LearnCodeOnline.in", "ghi123", nil},
// 	}

// 	// package this data as JSON data
// 	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
// 	if err != nil {
// 		panic(fmt.Errorf("Error converting to json: %w\n", err))
// 	}

// 	fmt.Printf("%s\n", finalJson)
// }

// func DecodeJson() {
// 	jsonDataFromWeb := []byte(`
// 		[
// 			{
// 				"course_name": "ReactJs BootCamp",
// 				"price": 299,
// 				"website": "LearnCodeOnline.in",
// 				"tags": [
// 					"web-dev",
// 					"js"
// 				]
// 			},
// 			{
// 				"course_name": "MERN BootCamp",
// 				"price": 199,
// 				"website": "LearnCodeOnline.in",
// 				"tags": [
// 					"full-stack",
// 					"js"
// 				]
// 			},
// 			{
// 				"course_name": "Angular BootCamp",
// 				"price": 299,
// 				"website": "LearnCodeOnline.in"
// 			}
// 		]
// 		`)

// 	// var lcoCourses []course

// 	// checkValid := json.Valid(jsonDataFromWeb)

// 	// if checkValid {
// 	// 	fmt.Println("JSON valid")
// 	// 	err := json.Unmarshal(jsonDataFromWeb, &lcoCourses)
// 	// 	if err != nil {
// 	// 		panic(fmt.Errorf("Error unmarshaling jsonDataFromWeb: %w", err))
// 	// 	}

// 	// 	fmt.Printf("%#v\n", lcoCourses)
// 	// } else {
// 	// 	panic(fmt.Errorf("INVALID JSON"))
// 	// }

// 	// some cases where you just want to add to key value
// 	var myOnlineData []map[string]any
// 	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

// 	for i, item := range myOnlineData {
// 		fmt.Println("Item:", i+1)
// 		for k, v := range item {
// 			fmt.Printf("key: %s, val:%#v\n", k, v)
// 		}
// 	}
// }

// func greeter() {
// 	fmt.Println("Hey there mod users")
// }

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Welcome to golang from YT</h1>"))
// }
