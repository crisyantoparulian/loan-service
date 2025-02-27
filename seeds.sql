-- Insert Users
INSERT INTO users (id, full_name, email, phone, nik, user_role)
VALUES
    ('11111111-1111-1111-1111-111111111111', 'Alice Borrower', 'alice@example.com', '081234567890', '3201010101010101', ARRAY['borrower']),
    ('22222222-2222-2222-2222-222222222222', 'Bob Investor', 'bob@example.com', '081234567891', '3201010101010102', ARRAY['investor']),
    ('33333333-3333-3333-3333-333333333333', 'Charlie Admin', 'charlie@example.com', '081234567892', '3201010101010103', ARRAY['admin']),
    ('44444444-4444-4444-4444-444444444444', 'David Validator', 'david@example.com', '081234567893', '3201010101010104', ARRAY['field_validator']),
    ('55555555-5555-5555-5555-555555555555', 'Eve Officer', 'eve@example.com', '081234567894', '3201010101010105', ARRAY['field_officer']);

-- Insert Borrowers
INSERT INTO borrowers (id, user_id, type)
VALUES
    ('66666666-6666-6666-6666-666666666666', '11111111-1111-1111-1111-111111111111', 'regular');

-- Insert Investors
INSERT INTO investors (id, user_id, company, investment_limit)
VALUES
    ('77777777-7777-7777-7777-777777777777', '22222222-2222-2222-2222-222222222222', 'XYZ Capital', 1000000.00);

-- Insert Employees
INSERT INTO employees (id, user_id, nik, role)
VALUES
    ('88888888-8888-8888-8888-888888888888', '33333333-3333-3333-3333-333333333333', 'EMP001', 'admin'),
    ('99999999-9999-9999-9999-999999999999', '55555555-5555-5555-5555-555555555555', 'EMP002', 'field_officer');
