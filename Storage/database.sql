-- Create Customers Table
CREATE TABLE Customers
(
    customer_id  INT AUTO_INCREMENT PRIMARY KEY,
    first_name   VARCHAR(50),
    last_name    VARCHAR(50),
    email        VARCHAR(100) UNIQUE,
    password     VARCHAR(255), -- You should store hashed passwords
    address      TEXT,
    phone_number VARCHAR(20)
);

-- Create Categories Table
CREATE TABLE Categories
(
    category_id   INT AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(50)
);

-- Create Products Table
CREATE TABLE Products
(
    product_id     INT AUTO_INCREMENT PRIMARY KEY,
    product_name   VARCHAR(100),
    description    TEXT,
    price          DECIMAL(10, 2), -- Adjust precision and scale as needed
    stock_quantity INT,
    category_id    INT,
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES Categories (category_id)
);

-- Create Orders Table
CREATE TABLE Orders
(
    order_id    INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT,
    order_date  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status      VARCHAR(20),
    FOREIGN KEY (customer_id) REFERENCES Customers (customer_id)
);

-- Create Order_Items Table
CREATE TABLE Order_Items
(
    order_item_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id      INT,
    product_id    INT,
    quantity      INT,
    price         DECIMAL(10, 2), -- Adjust precision and scale as needed
    FOREIGN KEY (order_id) REFERENCES Orders (order_id),
    FOREIGN KEY (product_id) REFERENCES Products (product_id)
);

-- Create Payments Table
CREATE TABLE Payments
(
    payment_id     INT AUTO_INCREMENT PRIMARY KEY,
    order_id       INT,
    payment_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payment_method VARCHAR(50),
    amount         DECIMAL(10, 2), -- Adjust precision and scale as needed
    FOREIGN KEY (order_id) REFERENCES Orders (order_id)
);
