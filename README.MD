
# ESEP Grade Calculator

A Go library for computing final course letter grades from assignments,
exams, and essays. This project is used for an automated testing exercise.

## Requirements

- Go 1.24.4 or later: https://go.dev/dl/
- Git (for cloning the project)
- No external Go packages required

## Installation

```bash
git clone https://github.com/YOUR_USERNAME/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
go version
```

Replace `YOUR_USERNAME` with your GitHub username.

## Repository Structure

- `README.MD`: documentation and setup instructions.
- `grade-calculator/go.mod`: module name and Go version.
- `grade-calculator/grade_calculator.go`: grade calculator implementation.
- `grade-calculator/grade_calculator_test.go`: unit tests.

## Grading Rules

Each category's average is multiplied by its weight.

| Category | Weight |
| --- | ---: |
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

`Final = 0.50 * assignmentAvg + 0.35 * examAvg + 0.15 * essayAvg`

| Final score | Letter |
| --- | --- |
| 90 or higher | A |
| 80 to below 90 | B |
| 70 to below 80 | C |
| 60 to below 70 | D |
| Below 60 | F |

A grade of C or higher is considered passing.

## Main Functions

- `NewGradeCalculator()` creates a calculator.
- `AddGrade(name, grade, gradeType)` adds a score.
- `GetFinalGrade()` returns the final letter grade.
- Grade types: `Assignment`, `Exam`, and `Essay`.

Example (using import alias `calculator`):
```go
gc := calculator.NewGradeCalculator()
gc.AddGrade("Homework", 95, calculator.Assignment)
gc.AddGrade("Exam 1", 85, calculator.Exam)
gc.AddGrade("Essay 1", 90, calculator.Essay)
finalGrade := gc.GetFinalGrade()
```

## Testing

Run tests from `grade-calculator`:

```bash
go test ./...
go test -v ./...  # verbose output
```

Original tests: `TestGetGradeA`, `TestGetGradeB`, `TestGetGradeF`.

## Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

These commands generate a profile, summarize coverage, and open an
HTML report. Coverage alone does not prove tests are correct.

## Known Starter Issues

- The average function sums positions rather than grade values.
- The weighted formula mistakenly reuses exam scores for essays.
- Empty categories may cause division by zero.
- Grades and grade types are not validated.

## Resources

- https://go.dev/doc/
- https://pkg.go.dev/testing
- https://go.dev/blog/cover
