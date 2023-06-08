# CPP Autograder Program

## Description

The CPP autograder program is a tool to test a CPP file program automatically using defined test cases by a test reviewer. This program is designed for college practical courses and built using Golang.

## Usage

To use the program, follow these steps:

1. Build the program and name it `grader.exe`.
2. Students have to log in to the system using the command `grader login`.
3. Generate the problem case file using the command `grader paket -p=<paket_code>`.
4. Test the CPP program file using the command `grader test -file=<filename> -p=<paket_code>`.
5. Submit the code and generate a report zip file using the command `grader submit -file=<filename> -p=<paket_code>`.

## Installation

To install the program, follow these steps:

1. Install Golang on your computer.
2. Clone the repository.
3. Build the program using the command `go build -o grader.exe .`

## Features

- Support for multiple test cases
- Ability to grade student submissions automatically
- Integration with Catch2 for testing

## How to Contribute

If you would like to contribute to the program, please follow these guidelines:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your changes and test them thoroughly.
4. Submit a pull request.

## Tests

To test the program, you can use Catch2 for testing, which is a popular C++ testing framework. Here is an example of how to run the tests:
