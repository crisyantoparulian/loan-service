-- ENUM definitions
CREATE TYPE loan_status AS ENUM ('proposed', 'approved', 'invested', 'disbursed');
CREATE TYPE user_role AS ENUM ('borrower', 'investor', 'admin', 'field_validator', 'field_officer');
CREATE TYPE borrower_type AS ENUM ('regular', 'premium', 'ultimate');
CREATE TYPE employee_role AS ENUM ('admin', 'field_officer');
CREATE TYPE document_reference AS ENUM ('table_visits', 'table_disbursements');
CREATE TYPE document_type AS ENUM ('proof_visit','loan_aggrement');

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name TEXT NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    nik CHAR(16) UNIQUE,
    user_role user_role[] NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Borrowers table
CREATE TABLE borrowers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    type borrower_type NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Investors table
CREATE TABLE investors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    company VARCHAR(255),
    investment_limit DECIMAL(12,2) NOT NULL CHECK (investment_limit >= 0),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Employees table
CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    nik VARCHAR(255) UNIQUE NOT NULL,
    role employee_role NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Loans table
CREATE TABLE loans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_code CHAR(14) UNIQUE NOT NULL,
    borrower_id UUID NOT NULL REFERENCES borrowers(id),
    field_validator_id UUID REFERENCES employees(id),
    principal_amount DECIMAL(12,2) NOT NULL CHECK (principal_amount > 0),
    total_investment_amount DECIMAL(12,2) NOT NULL DEFAULT 0 CHECK (total_investment_amount >= 0),
    total_investor INT NOT NULL DEFAULT 0 CHECK (total_investor >= 0),
    rate DECIMAL(5,2) CHECK (rate >= 0),
    roi DECIMAL(5,2) NOT NULL CHECK (roi >= 0),
    status loan_status NOT NULL DEFAULT 'proposed',
    approval_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Investments table
CREATE TABLE investments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_id UUID NOT NULL REFERENCES loans(id),
    investor_id UUID NOT NULL REFERENCES investors(id),
    amount DECIMAL(12,2) NOT NULL CHECK (amount >= 0),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Visits table
CREATE TABLE visits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_id UUID NOT NULL REFERENCES loans(id),
    employee_id UUID NOT NULL REFERENCES employees(id),
    notes TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Disbursements table
CREATE TABLE disbursements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    loan_id UUID NOT NULL REFERENCES loans(id),
    field_officer_id UUID NOT NULL REFERENCES employees(id),
    date_of_disbursement TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Documents table
CREATE TABLE documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reference_id UUID NOT NULL,
    reference_type document_reference NOT NULL,
    document_type document_type NOT NULL,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);