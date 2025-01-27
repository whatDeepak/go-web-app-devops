# Go Web Application

This is a simple web application built using Go, utilizing the `net/http` package to serve static pages and handle HTTP requests.

## Features

- Serves multiple static HTML pages (`home.html`, `courses.html`, `about.html`, `contact.html`).
- Static assets (CSS, JavaScript, images) are served from the `static/` directory.
- Routes are defined to load different pages based on the URL.

## Folder Structure

Ensure your project folder has the following structure:

```
/project-folder
  /static
    /home.html
    /courses.html
    /about.html
    /contact.html
    /styles.css  (or any other CSS files)
    /scripts.js  (or any other JS files)
  main.go
```

## Running the Server

To run the server, execute the following command in your terminal:

```bash
go run main.go
```

This will start the Go server on `localhost` at port `8080`.

## Accessing Pages

Once the server is running, you can access the following pages:

- Home page: `http://localhost:8080/home`
- Courses page: `http://localhost:8080/courses`
- About page: `http://localhost:8080/about`
- Contact page: `http://localhost:8080/contact`

## Static Assets

The server will automatically serve static files (CSS, JS, images) from the `/static` folder. Ensure that all static assets (like `styles.css`, `scripts.js`) are placed inside the `static/` directory.

## License

This project is open-source and available under the MIT License.