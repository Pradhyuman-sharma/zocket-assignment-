package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type Person struct {
    Name    string
    Age     int
    Country string
}

func parse() {
    // Open the CSV file
    file, err := os.Open("people.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Parse the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }

    // Create a slice of Person structs
    var people []Person
    for _, record := range records {
        age := 0
        fmt.Sscanf(record[1], "%d", &age)
        person := Person{
            Name:    record[0],
            Age:     age,
            Country: record[2],
        }
        people = append(people, person)
    }

    // Output the data as a table
    fmt.Printf("%-20s%-10s%-20s\n", "Name", "Age", "Country")
    for _, person := range people {
        fmt.Printf("%-20s%-10d%-20s\n", person.Name, person.Age, person.Country)
    }
}
