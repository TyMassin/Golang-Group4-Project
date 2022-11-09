// Class: C490
// Name: Group 4
// Assignment: Project
// Date: 10/20/22
package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
)

//Struct is turning out to be more complicated than I thought it would be. Feel free to edit it.
type movie struct {
	Title string
	Runtime int //time in minutes rounded up
	ReleaseDate string //Just keep it simple and write the date
}

//Stringer - may come in handy so I implemented a basic one to remind myself it exists. Delete if unused.
func (m movie) String() string{
	return fmt.Sprintf("%v", m.Title)
}

var db map[int]movie

func main() {
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
//NEW INPUT
/*
<table border="10" align="center" style="margin-top:20px">
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
</table>
*/
//OLD INPUT
/*
<div>Title <input type="text" id="title" name="title"></div>
<div>Runtime <input type="text" id="time" name="time"></div>
<div>Release Date <input type="text" id="date" name="date"></div>
<div><input type="submit"></div>
*/
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
	//fmt.Fprintf(w, "Attempting to write movie\n")
	//grab keys and values
	data, err := url.ParseQuery(r.URL.RawQuery)
	errorCheck(err)
	//open file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	//defer closing file
	defer fi.Close()

	//Format how to write the line
	s := data["title"][0] + ";" + data["time"][0] + ";" + data["date"][0] + ";\n"

	//Save that line to a file
	fi.Write([]byte(s))
	//Send them back to the create movie page
	//http://localhost:8000//createmovie
	http.Redirect(w, r, "/createmovie", http.StatusSeeOther)
}


func getMovies(w http.ResponseWriter, r *http.Request){

}

func updateMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}
