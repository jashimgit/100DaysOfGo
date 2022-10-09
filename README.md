

### Let's hug Go
Date: 29-09-2022

### Topics
    1. syntax
    `
        package main
        func main(){
            
        }
    `
    2. comments
    `
        // single line comment
        /*
            multiline comments
            this is a comment
        */
    
    `
    3. variables
    `
        // declair a variable
        var variableName type
        
        // assign a value to a variable

        variableName = "value"

        // another way to declair a variable

        variableName := value

        note: variable short syntax must be declaired with a value.

        // variable declaration case

        we can use Camel Case, Pascal case or Snake Case

        myVariableName string = "John"
        MyVariableName string = "John"
        my_variable_name string = "John"

    `
    4. constants

    ` 
    syntax: const constantName type = value
    `

    problem 1: write a program to check whether a given possitive number integer
    is a power of 2

    ref: problem/problem1.go

    problem2: Write a PHP program to check whether a given positive integer is a power of three.

    ref: problem/problem2.go


    Write a PHP program to check whether an integer is the power of another integer. Go to the editor
    Input : 16, 2
    Output : 16 is power of 2
    Example: For x = 16 and y = 2 the answer is "true", and for x = 12 and y = 2 "false"