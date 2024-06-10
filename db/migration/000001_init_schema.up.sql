CREATE TABLE users (
  username VARCHAR(255) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  phone VARCHAR(11),
  password_changed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  profile VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accounts (
  id BIGSERIAL PRIMARY KEY,
  owner VARCHAR(255) NOT NULL,
  account_number VARCHAR(20) NOT NULL,
  bank_name VARCHAR(100) NOT NULL,
  account_holder_name VARCHAR(100) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (owner) REFERENCES users(username)
);

CREATE TABLE posts (
  id BIGSERIAL PRIMARY KEY,
  author VARCHAR(255) NOT NULL,  
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  medias VARCHAR[],  
  price BIGINT NOT NULL,
  stock BIGINT NOT NULL,
  views BIGINT DEFAULT 0,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (author) REFERENCES users(username)
);

CREATE TABLE transactions (
  transaction_id BIGSERIAL PRIMARY KEY,
  post_id BIGINT NOT NULL,
  buyer VARCHAR(255) NOT NULL,
  seller VARCHAR(255) NOT NULL,
  status VARCHAR(20) DEFAULT 'pending',
  total_amount DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (buyer) REFERENCES users(username),
  FOREIGN KEY (seller) REFERENCES users(username)
);

CREATE TABLE payments (
  payment_id BIGSERIAL PRIMARY KEY,
  transaction_id BIGINT NOT NULL,
  payment_status VARCHAR(20) DEFAULT 'Pending',
  payment_method VARCHAR(50) NOT NULL,
  payment_date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  payment_amount DECIMAL(10, 2) NOT NULL,
  FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id)
);

CREATE TABLE like_with_post (
  username VARCHAR(255) NOT NULL,
  post_id BIGINT NOT NULL,
  FOREIGN KEY (username) REFERENCES users(username),
  FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE wish_with_product (
  username VARCHAR(255) NOT NULL,
  post_id BIGINT NOT NULL,
  FOREIGN KEY (username) REFERENCES users(username),
  FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE TABLE reviews (
  id BIGSERIAL PRIMARY KEY,
  seller VARCHAR(255) NOT NULL,
  reviewer VARCHAR(255) NOT NULL,
  rating INT NOT NULL,  
  content TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (seller) REFERENCES users(username),
  FOREIGN KEY (reviewer) REFERENCES users(username)
);

CREATE TABLE comments (
  id BIGSERIAL PRIMARY KEY,
  post_id BIGINT NOT NULL,
  parent_comment_id BIGINT NULL,
  commentor VARCHAR(255) NOT NULL,
  comment_text TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (commentor) REFERENCES users(username),
  FOREIGN KEY (parent_comment_id) REFERENCES comments(id)
);

CREATE TABLE search_count (
  id BIGSERIAL PRIMARY KEY,  
  search_text TEXT NOT NULL UNIQUE,
  count BIGINT NOT NULL DEFAULT 1,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP    
);

CREATE TABLE Rooms (
  id BIGSERIAL AUTO_INCREMENT PRIMARY KEY,
  user_a BIGINT NOT NULL,
  user_b BIGINT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_a) REFERENCES Users(id),
  FOREIGN KEY (user_b) REFERENCES Users(id)
);

CREATE TABLE Messages (
  id BIGSERIAL AUTO_INCREMENT PRIMARY KEY,
  room_id BIGINT NOT NULL,
  sender_id BIGINT NOT NULL,
  message TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (room_id) REFERENCES Rooms(id),
  FOREIGN KEY (sender_id) REFERENCES Users(id)
);

-- Create indexes after table creation
CREATE INDEX ON users (username);
CREATE INDEX ON accounts (owner);
CREATE INDEX ON posts (author);
CREATE INDEX ON transactions (post_id);
CREATE INDEX ON transactions (buyer);
CREATE INDEX ON transactions (seller);
CREATE INDEX ON payments (transaction_id);
CREATE INDEX ON like_with_post (username);
CREATE INDEX ON like_with_post (post_id);
CREATE INDEX ON wish_with_product (username);
CREATE INDEX ON wish_with_product (post_id);
CREATE INDEX ON reviews (seller);
CREATE INDEX ON reviews (reviewer);
CREATE INDEX ON comments (post_id);
CREATE INDEX ON comments (commentor);
CREATE INDEX ON search_count (search_text);
CREATE INDEX ON Rooms (user_a);
CREATE INDEX ON Rooms (user_b);
CREATE INDEX ON Messages (room_id);
CREATE INDEX ON Messages (sender_id);

