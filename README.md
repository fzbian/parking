# Parking José Max León Bilingual School

Repository dedicated to the code with which I participated in the CBJML school programming olympiads.

## Project status

The project at this time has already been launched as the programming olympics have already been held, so development has been stopped, however any bugs found and optimizations needed will be received and done.

## Installation

To continue with the development of this application it is necessary to have **Go** installed in our machine, in the same way it is necessary to install **Fyne**.
It is recommended to follow the steps without skipping any in the installation of these two
programs, otherwise the deployment of the application will fail.

An `.env` file must be created which will store the following variables:

```env
DB_USER=YOUR_USER
DB_PASSWORD=YOUR_PASSWORD
DB_HOST=YOUR_HOST
DB_PORT=YOUR_PORT
DB_DATABASE=YOUR_DATABASE
```

Remember that the database must be uploaded, the database file is located in the main folder with the name [db.sql](db.sql), the database is recommended to be imported directly from the main folder.

After uploading the database it is necessary to create the spots, for this you must add the following line in the `main.go` file at the time of running the code

```go
// Create the spots.
utils.CreateSpots()
```

After this you must comment or delete the line, otherwise it will recreate the spots every time you run the code.

## Technologies

Remember to do a `go get` before running the code

- Go
- Fyne
- MySQL
- Gorm
- fogleman/gg

## Screenshots

## Credits
created by [Fabian Martin](https://github.com/fzbian/)