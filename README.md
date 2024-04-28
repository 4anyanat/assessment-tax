# Tax Calculation (For GO Bootcamp)

## Description
Calculate level-based tax or tax refund regarding user inputs

### Inputs
1. Available inputs : Total income, With holding tax, Allowances (such as donation, k-receipt)
2. CSV format : Total Income, With holding tax, Donation allowance

### Outputs
1. Total tax that should be paid OR Total tax refund that should be paid back
2. Tax for each level tax rate

### Administration
The allowance amount of the personal deduction and k-receipt could be determined by authorized admins only

### Further Environment Variables
Database URL : `postgres://postgres:postgres@localhost:5432/ktaxes?sslmode=disable`