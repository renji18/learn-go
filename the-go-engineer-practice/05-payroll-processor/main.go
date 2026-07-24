package main

import "fmt"

type Payable interface {
	CalculatePay() (float64, error)
}

type Employee interface {
	Payable
	fmt.Stringer
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

type HourlyEmployee struct {
	Name                string
	HourlyRate          float64
	HoursWorkedPerMonth float64
}

type CommissionEmployee struct {
	Name                string
	BaseSalaryPerMonth  float64
	CommissionRate      float64
	SalesAmountPerMonth float64
}

type Freelancer struct {
	Name                      string
	AmountPerProject          float64
	ProjectsCompletedPerMonth int
}

type Intern struct {
	Name    string
	Stipend float64
}

type PayrollResult struct {
	Employee Employee
	Pay      float64
}

func (s SalariedEmployee) CalculatePay() (float64, error) {
	if s.AnnualSalary < 0 {
		return 0, fmt.Errorf("Salary cannot be negative")
	}

	return s.AnnualSalary / 12, nil
}

func (h HourlyEmployee) CalculatePay() (float64, error) {
	if h.HourlyRate < 0 {
		return 0, fmt.Errorf("Hourly rate cannot be negative")
	}

	if h.HoursWorkedPerMonth < 0 {
		return 0, fmt.Errorf("Hours worked per month cannot be negative")
	}

	return h.HourlyRate * h.HoursWorkedPerMonth, nil
}

func (c CommissionEmployee) CalculatePay() (float64, error) {
	if c.BaseSalaryPerMonth < 0 {
		return 0, fmt.Errorf("Base Salary cannot be negative")
	}

	if c.SalesAmountPerMonth < 0 {
		return 0, fmt.Errorf("Sales Amount cannot be negative")
	}

	if c.CommissionRate > 1 {
		return 0, fmt.Errorf("Commission rate cannot be more than 1")
	}

	return c.BaseSalaryPerMonth + (c.CommissionRate * c.SalesAmountPerMonth), nil
}

func (f Freelancer) CalculatePay() (float64, error) {
	if f.AmountPerProject < 0 {
		return 0, fmt.Errorf("Amount per project cannot be negative")
	}

	if f.ProjectsCompletedPerMonth < 0 {
		return 0, fmt.Errorf("Projects completed per month cannot be negative")
	}

	return f.AmountPerProject * float64(f.ProjectsCompletedPerMonth), nil
}

func (i Intern) CalculatePay() (float64, error) {
	if i.Stipend < 0 {
		return 0, fmt.Errorf("Stipend cannot be negative")
	}

	return i.Stipend, nil
}

func (s SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", s.Name, s.AnnualSalary)
}

func (h HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", h.Name, h.HourlyRate, h.HoursWorkedPerMonth)
}

func (c CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, CommRate: %.2f%%, Sales: $%.2f)",
		c.Name, c.BaseSalaryPerMonth, c.CommissionRate*100, c.SalesAmountPerMonth)
}

func (f Freelancer) String() string {
	return fmt.Sprintf("Freelancer: %s (Per project: $%.2f) for (projects completed: %d)", f.Name, f.AmountPerProject, f.ProjectsCompletedPerMonth)
}

func (i Intern) String() string {
	return fmt.Sprintf("Intern %s (Stipend: %.2f)", i.Name, i.Stipend)
}

func PrintEmployeeSummary(employee Employee) {
	fmt.Printf("Processing... %s\n", employee)
}

func CalculatePayroll(employees []Employee) ([]PayrollResult, float64, error) {
	var payrollResults = make([]PayrollResult, 0)
	var totalPayroll float64

	for _, e := range employees {
		pay, err := e.CalculatePay()

		if err != nil {
			if s, ok := e.(fmt.Stringer); ok {
				return nil, 0, fmt.Errorf("Error calculating pay for employee %s: %w", s, err)
			}
			return nil, 0, fmt.Errorf("employee: %w", err)
		}

		payrollResults = append(payrollResults, PayrollResult{Employee: e, Pay: pay})
		totalPayroll += pay
	}

	return payrollResults, totalPayroll, nil
}

func PrintPayroll(payrollResults []PayrollResult, totalPayroll float64) {
	for _, p := range payrollResults {
		PrintEmployeeSummary(p.Employee)
		fmt.Printf("Monthly payroll %.2f\n", p.Pay)
	}

	fmt.Printf("Total Payroll %.2f\n", totalPayroll)
}

func FilterHighEarners(payrollResults []PayrollResult, threshold float64) ([]Employee, error) {
	highEarners := make([]Employee, 0)

	for _, p := range payrollResults {
		if p.Pay > threshold {
			highEarners = append(highEarners, p.Employee)
		}
	}

	return highEarners, nil
}

func main() {
	salaried := SalariedEmployee{Name: "Harry", AnnualSalary: 1200000}
	hourly := HourlyEmployee{Name: "James", HourlyRate: 600, HoursWorkedPerMonth: 400}
	commission := CommissionEmployee{Name: "Roger", BaseSalaryPerMonth: 120000, CommissionRate: 0.034, SalesAmountPerMonth: 45000}
	intern := Intern{Name: "Mukes", Stipend: 20}

	employees := []Employee{salaried, hourly, commission, Freelancer{Name: "Aadarsh", AmountPerProject: 45000, ProjectsCompletedPerMonth: 3}, intern}

	payrollResults, totalPayroll, err := CalculatePayroll(employees)

	if err != nil {
		fmt.Println(err)
		return
	}

	PrintPayroll(payrollResults, totalPayroll)

	highEarners, err := FilterHighEarners(payrollResults, 300)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, h := range highEarners {
		fmt.Printf("Employee %s is a High Earner\n", h)
	}
}
