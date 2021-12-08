# todo-app-go

Make sure to run `go mod vendor` first to download all needed libraries

To run this application, call `make run` on root
To build only, call `make build` on root
Calling one of above will generate executable file in `bin/` folder, that can be used to start the application

After app binary is running, the web can be opened in http://localhost:8080/

```
bin                         # Binary output folder
backend
 ├── datastore              # Main datastore implementation code, add necessary implementation here
 |    └──  interface.go     # Contains interface struct with all function that needs to be implemented
 ├── webstatic              # Frontend Codes, taken from https://github.com/themaxsandelin/todo
 └── main.go                # Main program of the app
model
 └── model.go               # Data model of main object
```
