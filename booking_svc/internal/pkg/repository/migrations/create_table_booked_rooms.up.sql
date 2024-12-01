CREATE TABLE booked_rooms (
      id VARCHAR(255) PRIMARY KEY,     -- Уникальный идентификатор комнаты
      hotel_id VARCHAR(255) NOT NULL, -- Идентификатор отеля
      entry TIMESTAMP NOT NULL,       -- Дата и время заезда
      exit TIMESTAMP NOT NULL,        -- Дата и время выезда
      email VARCHAR(255) NOT NULL,    -- Электронная почта клиента
      is_paid BOOLEAN DEFAULT FALSE   -- Флаг оплаты (по умолчанию FALSE)
);
