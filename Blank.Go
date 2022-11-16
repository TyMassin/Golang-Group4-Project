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
	// "io/ioutil"
	"html/template"
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

//var db map[int]movie
var movies []movie

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

func createMovie(w http.ResponseWriter, r *http.Request){
	htmlForm := `<h1>Create Movie</h1>
	<form action="/savemovie?" method="GET">
		<div>Title <input type="text" id="title" name="title"</div>
		<div>Runtime <input type="text" id="time" name="time"></div>
		<div>Release Date <input type="text" id="date" name="date"></div>
		<div><input type="submit"></div>
	</form>`
	fmt.Fprintf(w, htmlForm)
}

func saveMovie(w http.ResponseWriter, r *http.Request){	
	//buttons for user
	htmlCreateMovieButton := `<h1>Movie Added!</h1>
	<form action="/createmovie?" method="GET">
		<div><input type="submit" value="Add Movie"></div>
	</form>`

	htmlGetMoviesButton := `<form action="/getmovies?" method="GET">
		<div><input type="submit" value="Show My Movies"></div>
	</form>`
	fmt.Fprintf(w, htmlCreateMovieButton)
	fmt.Fprintf(w, htmlGetMoviesButton)
	//fmt.Fprintf(w, "Attempting to write movie\n")
	//grab keys and values
	data, err := url.ParseQuery(r.URL.RawQuery)
	errorCheck(err)
	//open file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	//defer closing file
	defer fi.Close()

	//just get some data in it please. - Works now
	/*
	if _, err := fi.Write([]byte("appended some data\n")); err != nil { //well "Access is denied." so 
		fi.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	*/


	//Format how to write the line
	s := "Title:" + data["title"][0] + "Runtime:" + data["time"][0] + "Release Date:" + data["date"][0] + "\n"
	p:=movie{Title: data["title"][0], Runtime: data["time"][0], ReleaseDate: data["date"][0]}
	//Testing to see if I formated this right
	//fmt.Fprintf(w, s) //This works
	//Save that line to a file
	fi.Write([]byte(s))
	//Send them back to the create movie page
	//http.Redirect(w, r, "/createmovie", http.StatusSeeOther)
	t, err := template.ParseFiles("movies_template.html")
	fmt.Println(t.Execute(w, p))
}


func getMovies(w http.ResponseWriter, r *http.Request) {
	
	//reading and printing file
	fi, err := os.OpenFile("db.txt", os.O_RDWR|os.O_APPEND, 0660)
	errorCheck(err)
	defer fi.Close()

	htmlShowingMovies := `<table>
    <tr>
        <th>Title</th>
        <th>Runtime</th>
        <th>Release Date</th>
    </tr>

</table>

<style>
    th {
        padding-right: 1em;
    }
</style>`
fmt.Fprintf(w, htmlShowingMovies)
	//print only part of file

	

}

func updateMovie(w http.ResponseWriter, r *http.Request){

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

}