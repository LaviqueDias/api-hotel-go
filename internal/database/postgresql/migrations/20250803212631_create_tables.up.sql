-- Enums
CREATE TYPE room_status AS ENUM ('available', 'maintenance', 'occupied');
CREATE TYPE booking_status AS ENUM ('booked', 'cancelled', 'completed');

-- Hotels
CREATE TABLE hotels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Guests
CREATE TABLE guests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Rooms
CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    number INTEGER NOT NULL,
    status room_status NOT NULL DEFAULT 'available',
    daily_price DECIMAL(10, 2) NOT NULL,
    hotel_id INTEGER NOT NULL REFERENCES hotels(id) ON DELETE CASCADE
);

-- Bookings
CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    checkin_date DATE NOT NULL,
    checkout_date DATE NOT NULL,
    status booking_status NOT NULL DEFAULT 'booked',
    total_price DECIMAL(10, 2),
    guest_id INTEGER NOT NULL REFERENCES guests(id) ON DELETE CASCADE
);

-- Booking_Room (join table N:N)
CREATE TABLE booking_room (
    booking_id INTEGER NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
    room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    PRIMARY KEY (booking_id, room_id)
);
