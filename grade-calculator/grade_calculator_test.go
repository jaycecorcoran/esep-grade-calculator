package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 73, Assignment)
	gradeCalculator.AddGrade("exam 1", 76, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 71, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 64, Assignment)
	gradeCalculator.AddGrade("exam 1", 66, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 69, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}


func TestGetGradeF(t *testing.T) {
    expected_value := "F"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("open source assignment", 50, Assignment)
    gradeCalculator.AddGrade("exam 1", 55, Exam)
    gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}


func TestGradeTypeString(t *testing.T) {
    tests := []struct {
        gradeType GradeType
        expected  string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }

    for _, test := range tests {
        actual := test.gradeType.String()

        if actual != test.expected {
            t.Errorf("Expected '%s', got '%s'",
                test.expected, actual)
        }
    }
}

