CREATE DATABASE booking_room;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name                    VARCHAR(100),
    divisi                  VARCHAR(100),
    jabatan                 VARCHAR(100),
    email                   VARCHAR(100),
    password                VARCHAR(100),
    role                    VARCHAR(100),
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt               TIMESTAMP
);

CREATE TABLE facilities (
    Id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    RoomDescription         TEXT,
    Fwifi                   VARCHAR(100),
    FsoundSystem            VARCHAR(100),
    Fprojector              VARCHAR(100),
    FscreenProjector        VARCHAR(100),
    Fchairs                 VARCHAR(100),
    Ftables                 VARCHAR(100),
    FsoundProof             VARCHAR(100),
    FsmonkingArea           VARCHAR(100),
    Ftelevison              VARCHAR(100),
    FAc                     VARCHAR(100),
    Fbathroom               VARCHAR(100),
    FcoffeMaker             VARCHAR(100),
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt               TIMESTAMP
);

CREATE TABLE rooms (
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    RoomType                VARCHAR(100),
    Capacity                int,
    Facilities              UUID,
    Status                  VARCHAR(100),
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt               TIMESTAMP,
    CONSTRAINT FK_facility FOREIGN KEY(Facilities) REFERENCES facilities(id)
);

CREATE TABLE booking (
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userId                  UUID,
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt               TIMESTAMP,
    CONSTRAINT FK_userId FOREIGN KEY(userId) REFERENCES users(id)
);

CREATE TABLE booking_details(
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bookingId               UUID,
    roomId                  UUID,
    bookingDate             TIMESTAMP,
    bookingDateEnd          TIMESTAMP,
    status                  VARCHAR(100),
    description             TEXT,
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt               TIMESTAMP,
    CONSTRAINT FK_bookingId FOREIGN KEY(bookingId) REFERENCES booking(id),
    CONSTRAINT FK_roomId FOREIGN KEY(roomId) REFERENCES rooms(id),

);