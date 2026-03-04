CREATE TABLE organizers (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    slug VARCHAR(150) UNIQUE,
    type VARCHAR(20) NOT NULL DEFAULT 'individual',
    company_name VARCHAR(200),
    document_type VARCHAR(50),
    document_number VARCHAR(50),
    contact_email VARCHAR(255),
    contact_whatsapp VARCHAR(20),
    logo_url VARCHAR(255),
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);