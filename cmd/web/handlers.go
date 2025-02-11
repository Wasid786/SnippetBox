package main

import (
	"errors"
	"fmt"
	"net/http"
	"pappu/internal/models"
	"pappu/internal/validator"
	"runtime/debug"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form: "content"`
	Expires             int    `form: "expires"`
	validator.Validator `form:"-"`
}

// type snippetCreateForm struct {
// 	Title       string
// 	Content     string
// 	Expires     int
// 	FieldErrors map[string]string
// }

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because httprouter matches the "/" path exactly, we can now remove the
	// manual check of r.URL.Path != "/" from this handler.
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Use the PopString() method to retrieve the value for the "flash" key.
	// PopString() also deletes the key and value from the session data, so it
	// acts like a one-time fetch. If there is no matching key in the session
	// data this will return the empty string.
	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl", data)
}

// Add a new snippetCreate handler, which for now returns a placeholder
// response. We'll update this shortly to show a HTML form.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = snippetCreateForm{
		Expires: 365,
	}
	app.render(w, http.StatusOK, "create.tmpl", data)
}

// Rename this handler to snippetCreatePost.
// func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
// 	// Checking if the request method is a POST is now superfluous and can be
// 	// removed, because this is done automatically by httprouter.
// 	// title := "O snail"
// 	// content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
// 	// expires := 7

// 	// id, err := app.snippets.Insert(title, content, expires)
// 	// if err != nil {
// 	// 	app.serverError(w, err)
// 	// 	return
// 	// }
// 	// Update the redirect path to use the new clean URL format.
// 	// http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
// 	err := r.ParseForm()
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	title := r.PostForm.Get("title")
// 	content := r.PostForm.Get("content")
// 	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
// 	if err != nil {
// 		app.clientError(w, http.StatusBadRequest)
// 		return
// 	}

// 	fieldErrors := make(map[string]string)
// 	if strings.TrimSpace(title) == "" {
// 		fieldErrors["title"] = "This filed cannot be blank"

// 	} else if utf8.RuneCountInString(title) > 100 {
// 		fieldErrors["title"] = "This field cannot be more than 100 characters long"
// 	}
// 	if strings.TrimSpace(content) == "" {
// 		fieldErrors["content"] = "This field cannot be blank"
// 	}

// 	if expires != 1 && expires != 7 && expires != 365 {
// 		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
// 	}
// 	if len(fieldErrors) > 0 {
// 		fmt.Fprint(w, fieldErrors)
// 		return
// 	}

// 	id, err := app.snippets.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
// }

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Use the Put() method to add a string value ("Snippet successfully
	// created!") and the corresponding key ("flash") to the session data.
	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

	fmt.Printf("Form Data: %+v\n", form)

}
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s \n %s ", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("The template %s does not exist ", page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
