-- Таблица брендов (производителей)
CREATE TABLE brand (
    id SERIAL PRIMARY KEY,
    brand_name VARCHAR(50) UNIQUE NOT NULL  -- Название бренда
);

-- Таблица моделей устройств
CREATE TABLE model (
    id SERIAL PRIMARY KEY,
    brand_id INT REFERENCES brand(id), -- Внешний ключ к таблице брендов 
    model_name VARCHAR(50) NOT NULL    -- Название модели
);

-- Таблица типов аккумуляторов
CREATE TABLE battery_type (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL -- Название типа аккумулятора
);

-- Таблица основных характеристик накопителей
CREATE TABLE energy_storage_characteristics (
    id SERIAL PRIMARY KEY,
    specific_energy NUMERIC,     -- Удельная энергоемкость, Вт*ч/кг
    cycle_life INT,              -- Ресурс использования, циклов
    charge_time INTERVAL,        -- Продолжительность заряда
    discharge_time INTERVAL,     -- Продолжительность разряда 
    temperature_range INT4RANGE, -- Температурный диапазон, оС
    efficiency NUMERIC,          -- КПД, %
    self_discharge INTERVAL       -- Саморазряд, % в день
);

-- Таблица electrochemical_battery (Электрохимические аккумуляторы)
CREATE TABLE electrochemical_battery_history (
    id SERIAL PRIMARY KEY,
    type_id INT REFERENCES battery_type(id),                              -- Внешний ключ к таблице типов аккумуляторов
    characteristics_id INT REFERENCES energy_storage_characteristics(id), -- Внешний ключ к таблице характеристик 
    temperature_range INT4RANGE,                                          -- Температурный диапазон, оС
    input_voltage NUMERIC,                                                -- Входное напряжение
    output_voltage NUMERIC,                                               -- Выходное напряжение
    internal_resistance NUMERIC,                                          -- Внутреннее сопротивление
    operating_temperature NUMERIC,                                         -- Рабочая температура
    recorded_at TIMESTAMP DEFAULT now()                                   -- Дата и время записи мощности
);

-- Таблица thermal_storage_history (История тепловых накопителей)
CREATE TABLE thermal_storage_history (
    id SERIAL PRIMARY KEY,
    characteristics_id INT REFERENCES energy_storage_characteristics(id), -- Внешний ключ к таблице характеристик 
    thermal_power NUMERIC,                                                -- Тепловая мощность
    efficiency NUMERIC,                                                   -- КПД 
    recorded_at TIMESTAMP DEFAULT now()                                   -- Дата и время записи
);

-- Таблица hydrogen_storage (Водородные накопители)
CREATE TABLE hydrogen_storage (
    id SERIAL PRIMARY KEY,
    characteristics_id INT REFERENCES energy_storage_characteristics(id), -- Внешний ключ к таблице характеристик 
    temperature_range INT4RANGE,                                          -- Температурный диапазон, оС
    brand_id INT REFERENCES brand(id),                                    -- Внешний ключ к таблице брендов
    model_id INT REFERENCES model(id),                                    -- Внешний ключ к таблице моделей 
    nominal_voltage NUMERIC                                               -- Номинальное напряжение
);

-- Таблица solar_panel (Солнечные панели)
CREATE TABLE solar_panel (
    id SERIAL PRIMARY KEY,             -- Идентификатор панели 
    brand_id INT REFERENCES brand(id), -- Внешний ключ к таблице брендов
    model_id INT REFERENCES model(id), -- Внешний ключ к таблице моделей 
    length NUMERIC,                    -- Длина панели
    width NUMERIC,                     -- Ширина панели
    weight NUMERIC,                    -- Вес панели
    panel_count INTEGER,               -- Количество панелей
    battery_count INTEGER,             -- Количество аккумуляторов
    voltage NUMERIC                    -- Напряжение солнечной панели: В
);

-- Таблица solar_panel_power_history (История номинальной мощности солнечных панелей)
CREATE TABLE solar_panel_power_history (
    id SERIAL PRIMARY KEY,
    panel_id BIGINT REFERENCES solar_panel(id), -- Внешний ключ к таблице solar_panel
    nominal_power NUMERIC,                      -- Номинальная мощность: кВт
    recorded_at TIMESTAMP DEFAULT now()         -- Дата и время записи мощности
);

-- Таблица типов инверторов
CREATE TABLE inverter_type (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) UNIQUE NOT NULL, -- Название типа инвертора
    description TEXT                       -- Описание типа инвертора
);

-- Таблица inverter (Инверторы)
CREATE TABLE inverter (
    id SERIAL PRIMARY KEY,
    type_id INT REFERENCES inverter_type(id), -- Внешний ключ к таблице типов инверторов
    specific_energy NUMERIC                   -- Удельная энергоемкость 
);

-- Таблица wind_turbine_history (История ветроаккумулирующих источников)
CREATE TABLE wind_turbine_history (
    id SERIAL PRIMARY KEY,
    brand_id INT REFERENCES brand(id),  -- Внешний ключ к таблице брендов
    model_id INT REFERENCES model(id),  -- Внешний ключ к таблице моделей 
    power NUMERIC,                      -- Мощность
    voltage NUMERIC,                    -- Напряжение 
    recorded_at TIMESTAMP DEFAULT now() -- Дата и время записи
);