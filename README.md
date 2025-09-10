This is simply to demonstrate that the combobox will not function properly if the form is reloaded dynamically. It uses [Datastar](https://data-star.dev/examples/web_component) to dynamically load new HTML fragments.

This demo requires golang.


1. From the project root, run `go mod tidy` to install dependencies
2. Run the server `go run main.go`
2. Browse to 'http://localhost:8080
3. Click on the combobox to see that there are options that can be selected from. Optionally fill out any form fields (not required), then click 'Save"
4. Save will POST to the server, and the form's HTML is sent back and re-rendered.
5. At this point the combobox is unusable. Click in it and notice the errors in the console.
