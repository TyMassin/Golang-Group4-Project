// Class: C490
// Name: Group 4
// Assignment: Project
// Date: 11/16/22
package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
	"bufio"
	"strings"
	// "io/ioutil"
	//"html/template"
)

//Struct is turning out to be more complicated than I thought it would be. Feel free to edit it.
type movie struct {
	Title string
	Runtime string //time in minutes rounded up
	ReleaseDate string //Just keep it simple and write the date
}

//Stringer - may come in handy so I implemented a basic one to remind myself it exists. Delete if unused.
func (m movie) String() string{
	return fmt.Sprintf("%v", m.Title)
}

//Globals
var db map[int]movie = make(map[int]movie) //Our database
var count int //used in helping create and manage the db keys

func main() {
	//Open the file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	defer fi.Close()
	//Make bufio scanner to read line by line later
	scanner := bufio.NewScanner(fi)

	count = 1
	for scanner.Scan() {
		//Grab every raw line
		raw := scanner.Text()
		//split into an array of data by the ; delliniator
		data := strings.Split(raw, ";")
		//add to db
		db[count] = movie{data[0], data[1], data[2]}
		//increment key
		count++
	}

	//MUX Handlers probably wont use but here it is
	/*
	mux := http.NewServeMux()
	mux.Handle("/createmovie", http.HandlerFunc(createMovie))
	mux.Handle("/savemovie", http.HandlerFunc(saveMovie))
	mux.Handle("/getmovies", http.HandlerFunc(getMovies))
	mux.Handle("/updatemovie", http.HandlerFunc(updateMovie))
	mux.Handle("/deletemovie", http.HandlerFunc(deleteMovie))
	log.Fatal(http.ListenAndServe(":8000", mux))
	*/
	//Handlers
	http.HandleFunc("/createmovie", createMovie)
	http.HandleFunc("/savemovie", saveMovie)
	http.HandleFunc("/getmovies", getMovies)
	http.HandleFunc("/updatemovie", updateMovie)
	http.HandleFunc("/deletemovie", deleteMovie)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func errorCheck(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	htmlForm := `
		<table border="10" align="left" style="margin-top:10px">
			<form action="/savemovie?" method="GET">
				<thead align="center">
					<tr>
						<th colspan="2">Create Movie</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>Title:</td>
						<td><input type="text" id="title" name="title"></td>
					</tr>
					<tr>
						<td>Runtime:</td>
						<td><input type="text" id="time" name="time"></td>
					</tr>
					<tr>
						<td>Date:</td>
						<td><input type="text" id="date" name="date"></td>
					</tr>
					<tr>
						<td colspan="2" align="right"><input type="submit"></td>
					</tr>
				</tbody>
			</form>
		</table>`
	fmt.Fprintf(w, htmlForm)
}

func saveMovie(w http.ResponseWriter, r *http.Request){	
	//grab keys and values
	data, err := url.ParseQuery(r.URL.RawQuery)
	errorCheck(err)
	//open file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	//defer closing file
	defer fi.Close()

	//add movie to db
	p := movie{data["title"][0],data["time"][0],data["date"][0]}
	db[count] = p
	count++

	//Format how to write the line
	s := data["title"][0] + ";" + data["time"][0] + ";" + data["date"][0] + ";\n"
	//Save that line to a file
	fi.Write([]byte(s))

	//Send them back to the create movie page
	//http://localhost:8000//createmovie
	http.Redirect(w, r, "/createmovie", http.StatusSeeOther)
}


func getMovies(w http.ResponseWriter, r *http.Request) {
	
	//Open the file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	defer fi.Close()

	htmlShowingMoviesHead := `<table border="10" align="left" style="margin-top:10px">
		<form action="/savemovie?" method="GET">
			<thead align="center">
				<tr>
					<th colspan="4">Movie List</th>
				</tr>
				<tr>
					<th>ID</th>
					<th>Title</th>
					<th>Runtime</th>
					<th>Release Date</th>
				</tr>
			</thead>
			<tbody>`


	//How the body is formatted (Each line from the db should be printed in this way)
	/*
	<tr>
		<td></td>
		<td></td>
		<td></td>
		<td></td>
	</tr>
	*/
	
	htmlShowingMoviesFoot := `
			</tbody>
		</form>
	</table>`

	/////////////Now to print/////////////////
	//Head
	fmt.Fprintf(w, htmlShowingMoviesHead)
	//Body
	//This prints them out mostly randomly. we could do different handlers to sort this differently if we need more.
	for i, v := range db {
		fmt.Fprintf(w, `<tr>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		</tr>`, i, v.Title, v.Runtime, v.ReleaseDate)
	}
	//Foot
	fmt.Fprintf(w, htmlShowingMoviesFoot)
}

func updateMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}