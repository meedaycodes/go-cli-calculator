# CLI Calculator (Go)

A simple Command-Line Interface (CLI) calculator built in Go.  
This calculator supports multiple mathematical operations, handles user input dynamically, and provides a friendly interactive experience.

---

## **Features**

- Addition of multiple numbers.
- Subtraction of two numbers.
- Multiplication of multiple numbers.
- Division of two numbers (with zero-division check).
- Exponentiation (power).
- Square root calculation.
- Continuous loop – you can perform multiple operations until you type `exit`.
- Clean and beginner-friendly Go code.

---

## **Prerequisites**

- [Go](https://go.dev/dl/) installed (Go 1.16+ recommended).
- A terminal or command prompt.

---

## **Project Structure**
cli-calculator/
│
├── main.go     # The main Go application code
└── README.md   # Project documentation


## **Installation and Setup**
---
1. **Clone the repository:**

```bash
git clone https://github.com/<your-username>/cli-calculator.git
cd cli-calculator

2.	**Run the program:**
go run main.go

## **Example session**
########################

Welcome to using Habeeb's CLI Calculator
The following Operations are supported: Addition, Subtraction, Multiplication, Division, Exponentiation and Squareroot

########################
State the Operation you want to perform:
Addition
########################
Provide the list of numbers to add:
5
10
20
<Press any letter or non-number to stop input>
You have provided this list: [5 10 20].
The sum of the numbers in the provided list is: 35

## ** To exit**
State the Operation you want to perform:
exit
