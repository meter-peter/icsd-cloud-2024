-- Connect to the database
\c userdb

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    identity_number VARCHAR(8) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    date_of_birth DATE NOT NULL,
    addresses JSONB NOT NULL,
    phone_numbers JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index
CREATE INDEX IF NOT EXISTS idx_identity_number ON users(identity_number);

-- Insert initial data
INSERT INTO users (identity_number, first_name, last_name, gender, date_of_birth, addresses, phone_numbers)
VALUES
('AB123456', 'Γιώργος', 'Παπαδόπουλος', 'male', '1990-05-15', '["Οδός Ερμού 10", "Λεωφόρος Αλεξάνδρας 50"]', '["2101234567", "6971234567"]'),
('CD789012', 'Μαρία', 'Γεωργίου', 'female', '1985-11-22', '["Πλατεία Συντάγματος 1"]', '["2109876543"]'),
('EF345678', 'Νίκος', 'Δημητρίου', 'male', '1995-03-30', '["Οδός Πανεπιστημίου 30", "Οδός Σόλωνος 15"]', '["2105555555", "6989876543"]')
ON CONFLICT (identity_number) DO NOTHING;