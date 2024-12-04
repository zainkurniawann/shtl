-- Users Table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    user_uuid UUID DEFAULT gen_random_uuid() NOT NULL UNIQUE,  -- UUID for user
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    status VARCHAR(20),  -- User status (e.g., 'active', 'inactive')
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Roles Table
CREATE TABLE Roles (
    id SERIAL PRIMARY KEY,
    role_code VARCHAR(50) UNIQUE NOT NULL,  -- Role code (e.g., 'SA', 'AS', 'D', 'P')
    role_name VARCHAR(100) NOT NULL         -- Role name (e.g., 'Superadmin', 'School Admin', etc.)
);

-- Super Admin Details Table
CREATE TABLE Super_admin_details (
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,   -- Linking to Users table
    PRIMARY KEY (user_id)
);

-- School Table
CREATE TABLE School (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    contact VARCHAR(20) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- School Admin Details Table
CREATE TABLE School_admin_details (
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,  -- Linking to Users table
    school_id INT REFERENCES School(id) ON DELETE CASCADE,  -- Linking to School table
    PRIMARY KEY (user_id)
);

-- Vehicle Table
CREATE TABLE Vehicle (
    id SERIAL PRIMARY KEY,
    vehicle_name VARCHAR(100) NOT NULL,
    vehicle_number VARCHAR(50) UNIQUE NOT NULL,
    vehicle_type VARCHAR(50) NOT NULL,  -- Vehicle type (e.g., "Car", "Bus", etc.)
    colour VARCHAR(50) NOT NULL,
    seats INT NOT NULL,
    status VARCHAR(20),  -- Vehicle status (e.g., "active", "inactive")
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Driver Details Table
CREATE TABLE Driver_details (
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,  -- Linking to Users table
    vehicle_id INT REFERENCES Vehicle(id) ON DELETE CASCADE,  -- Linking to Vehicle table
    PRIMARY KEY (user_id)
);

-- Routes Table
CREATE TABLE Routes (
    id SERIAL PRIMARY KEY,
    route_name VARCHAR(255) NOT NULL,
    start_location VARCHAR(255) NOT NULL,
    end_location VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Route Points Table
CREATE TABLE Route_points (
    id SERIAL PRIMARY KEY,
    route_id INT REFERENCES Routes(id) ON DELETE CASCADE,  -- Linking to Routes table
    point_name VARCHAR(255) NOT NULL,  -- Name of the stop point
    latitude DECIMAL(9,6),  -- Latitude of the stop point
    longitude DECIMAL(9,6), -- Longitude of the stop point
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Parent Details Table
CREATE TABLE Parent_details (
    user_id INT REFERENCES Users(id) ON DELETE CASCADE,  -- Linking to Users table
    children INT[],  -- Array to store IDs of students
    PRIMARY KEY (user_id)
);

-- Student Table
CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    parent_id INT REFERENCES Users(id) ON DELETE SET NULL,  -- Linking to Users (parents)
    school_id INT REFERENCES School(id) ON DELETE CASCADE,  -- Linking to School table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
