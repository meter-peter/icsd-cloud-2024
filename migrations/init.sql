-- Δημιουργία της βάσης δεδομένων
CREATE DATABASE your_database_name;

-- Σύνδεση στη νέα βάση δεδομένων
\c your_database_name

-- Δημιουργία του πίνακα users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    identity_number VARCHAR(8) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    date_of_birth DATE NOT NULL,
    addresses TEXT[] NOT NULL,
    phone_numbers TEXT[] NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Δημιουργία ευρετηρίου για το identity_number
CREATE INDEX idx_identity_number ON users(identity_number);

-- Προσθήκη μερικών δοκιμαστικών δεδομένων
INSERT INTO users (identity_number, first_name, last_name, gender, date_of_birth, addresses, phone_numbers) VALUES
('AB123456', 'Γιώργος', 'Παπαδόπουλος', 'male', '1990-05-15', ARRAY['Οδός Ερμού 10', 'Λεωφόρος Αλεξάνδρας 50'], ARRAY['2101234567', '6971234567']),
('CD789012', 'Μαρία', 'Γεωργίου', 'female', '1985-11-22', ARRAY['Πλατεία Συντάγματος 1'], ARRAY['2109876543']),
('EF345678', 'Νίκος', 'Δημητρίου', 'male', '1995-03-30', ARRAY['Οδός Πανεπιστημίου 30', 'Οδός Σόλωνος 15'], ARRAY['2105555555', '6989876543']);

-- Δημιουργία χρήστη για τη διαχείριση της βάσης (προαιρετικό)
CREATE USER your_username WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_username;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO your_username;