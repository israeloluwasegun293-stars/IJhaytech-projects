package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type pageData struct {
	Result       string
	ErrorMessage string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", 500)
		return
	}
	temp.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	r.ParseForm()

	textInput := r.FormValue("text")

	fontInput := r.FormValue("banner")

	textInput = strings.ReplaceAll(textInput, "\\n", "\n")

	if textInput == "" || fontInput == "" {
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "internal server error", 500)
			return
		}

		w.WriteHeader(400)
		temp.Execute(w, pageData{ErrorMessage: "The input cannot be empty, please input a valid text"})
		return

		// w.Write([]byte("pageData{ErrorMessage: Dear User, your input contains unprintable characters"))
		// return

	}

	// log.Printf("input: %q\n", textInput)
	// log.Println(fontInput)
	for _, r := range textInput {
		if r != '\n' && r != '\r' && (r < 32 || r > 126) {
			temp, err := template.ParseFiles("templates/index.html")
			if err != nil {
				http.Error(w, "internal server error", 500)
				return
			}

			w.WriteHeader(400)
			temp.Execute(w, pageData{ErrorMessage: "Dear User, your input contains unprintable characters"})
			return
		}
	}

	result := asciiGenerator(textInput, fontInput)

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("error parsing files")
		return
	}

	temp.Execute(w, pageData{Result: result})

}

func asciiGenerator(textInput string, fontInput string) string {

	file, err := os.ReadFile("banners/" + fontInput + ".txt")
	if err != nil {
		log.Println("error reading file")
		return ""
	}

	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", ""), "\n")

	textInput = strings.ReplaceAll(textInput, "\r\n", "\n")

	const letterHeight = 8
	var output strings.Builder

	linesInput := strings.Split(textInput, "\n")

	for _, line := range linesInput {
		if line == "" {
			output.WriteString("\n")
			continue
		}

		for i := 0; i < letterHeight; i++ {
			for _, char := range line {
				// if char < 32 || char > 126 {
				// 	output.WriteString("")
				// 	break
				// }

				index := (int(char) - 32) * 9

				output.WriteString(lines[index+i+1])
			}
			output.WriteString("\n")
		}
	}
	return output.String()
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ascii-art", asciiArtHandler)

	log.Println("starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
