# AeroCast_API

A simple weather application with MongoDB as the database.

## Description

AeroCast_API is a weather application built using the Go programming language, MongoDB as the database, and the Echo framework. This application allows users to access and view weather data based on city and country information. Notably, the weather data is stored in MongoDB, and the application does not require a login system for user access.

## Features

- View weather data based on city and country.
- Add and searching cities.

## Technologies Used

- Programming Language: Go
- Database: MongoDB
- Framework: Echo

## Run Locally

Clone the project

```bash
git clone https://github.com/fauzibrillian/AeroCast_API.git
```

Go to the project directory

```bash
cd AeroCast_API
```

Install dependency

```bash
go mod tidy
```

## OpenWeatherMap API Key

Ensure you have an OpenWeatherMap API key for accessing weather data. Add the API key to your configuration file.

## KEY VALUE
![alt text](image.png)

## API

<summary>City</summary> 
<div>
  
| Feature City | Endpoint | Param | Function |
| --- | --- | --- | --- |
| GET | /city | NameCity or NameCountry | + | Get and Search all city. |
| POST | /city  | - | + | Post city data. |



</details>

<div>
