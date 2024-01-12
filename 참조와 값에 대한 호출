package main

import (
    "fmt"
)

var users []User

func main() {
    // jsonUsers := "hello world"
    users = []User{
        {
            Name:   "name1",
            Age:    1,
            Gender: "M",
        },
        {
            Name:   "name2",
            Age:    2,
            Gender: "F",
        },
        {
            Name:   "name3",
            Age:    3,
            Gender: "M",
        },
    }

    Delete("name2")
}

// Call by Reference -> 참조에 의한 호출
// Call by Value -> 값에 의한 호출

type User struct {
    Name   string
    Age    int
    Gender string
}

func Delete(name string) []User {
    var result []User
    for i, user := range users {
        if user.Name == name {
            continue
        }

        result = append(result, users[i:][0])
    }
    fmt.Println(result)

    return result
}
