package main

import "net/http"

var landingTmpl = []byte(`
	<!doctype html>
	<html lang="en" class="h-100">
	<head>
	    <meta charset="utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	    <title></title>
	</head>
	<body class="d-flex flex-column h-100">
	    <!doctype html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <title>ESPC Images</title>
	    <style>
			@import url(//fonts.googleapis.com/css?family=Lato:700);

			body {
				margin:0;
				font-family:'Lato', sans-serif;
				text-align:center;
				color: #999;
			}

			.welcome {
				width: 300px;
				height: 200px;
				position: absolute;
				left: 50%;
				top: 50%;
				margin-left: -150px;
				margin-top: -100px;
			}

			a, a:visited {
				text-decoration:none;
			}

			h1 {
				font-size: 32px;
				margin: 16px 0 0 0;
			}
	    </style>
	</head>
	<body>
	    <div class="welcome">
	        <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
	        <svg width="100" version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	             viewBox="0 0 16 16" enable-background="new 0 0 16 16" xml:space="preserve">
	        <path fill="#81bb5e" d="M8,0C3.6,0,0,3.6,0,8s3.6,8,8,8s8-3.6,8-8S12.4,0,8,0z M7,11L4,8l1.5-1.5L7,8l3.5-3.5L12,6L7,11z" />

	    </svg>
	        <h1>Images</h1>

	    </div>
	</body>
	</html>



	</body>
	</html>

`)

func handleLanding(reqID string, rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.WriteHeader(200)
	rw.Write(landingTmpl)
}
