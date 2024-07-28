CREATE TABLE users (
    username VARCHAR(255) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(15),
    password_changed_at TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    profile VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    medias TEXT[],
    price BIGINT NOT NULL,
    stock BIGINT NOT NULL,
    views BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author) REFERENCES users(username)
);

CREATE INDEX idx_posts_author_created_at ON posts(author, created_at);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    seller VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    price BIGINT NOT NULL,
    views BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (seller) REFERENCES users(username)
);

CREATE INDEX idx_products_seller ON products(seller);

CREATE TABLE transactions (
    transaction_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    buyer VARCHAR(255) NOT NULL,
    seller VARCHAR(255) NOT NULL,
    status VARCHAR(20) DEFAULT 'listing',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (buyer) REFERENCES users(username),
    FOREIGN KEY (seller) REFERENCES users(username)
);

CREATE INDEX idx_transactions_product_id ON transactions(product_id);
CREATE INDEX idx_transactions_buyer ON transactions(buyer);
CREATE INDEX idx_transactions_seller ON transactions(seller);

CREATE TABLE like_with_post (
    username VARCHAR(255) NOT NULL,
    post_id BIGINT NOT NULL,
    PRIMARY KEY (username, post_id),
    FOREIGN KEY (username) REFERENCES users(username),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);

CREATE INDEX idx_like_with_post_username ON like_with_post(username);
CREATE INDEX idx_like_with_post_post_id ON like_with_post(post_id);

CREATE TABLE wish_with_product (
    username VARCHAR(255) NOT NULL,
    product_id BIGINT NOT NULL,
    PRIMARY KEY (username, product_id),
    FOREIGN KEY (username) REFERENCES users(username),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE INDEX idx_wish_with_product_username ON wish_with_product(username);
CREATE INDEX idx_wish_with_product_product_id ON wish_with_product(product_id);

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

CREATE INDEX idx_reviews_seller ON reviews(seller);
CREATE INDEX idx_reviews_reviewer ON reviews(reviewer);

CREATE TABLE comments (
    id BIGSERIAL PRIMARY KEY,
    post_id BIGINT NOT NULL,
    parent_comment_id BIGINT,
    commentor VARCHAR(255) NOT NULL,
    comment_text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (commentor) REFERENCES users(username),
    FOREIGN KEY (parent_comment_id) REFERENCES comments(id) ON DELETE CASCADE
);

CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_commentor ON comments(commentor);
CREATE INDEX idx_comments_parent_comment_id ON comments(parent_comment_id);

CREATE TABLE search_count (
    id BIGSERIAL PRIMARY KEY,
    search_text TEXT NOT NULL UNIQUE,
    count BIGINT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_search_count_search_text ON search_count(search_text);

CREATE TABLE rooms (
    id BIGSERIAL PRIMARY KEY,
    user_a VARCHAR(255) NOT NULL,
    user_b VARCHAR(255) NOT NULL,
    type VARCHAR(255) DEFAULT 'chat',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_a) REFERENCES users(username),
    FOREIGN KEY (user_b) REFERENCES users(username),
    CONSTRAINT unique_users UNIQUE (user_a, user_b)
);

CREATE INDEX idx_rooms_user_a ON rooms(user_a);
CREATE INDEX idx_rooms_user_b ON rooms(user_b);

CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    room_id BIGINT NOT NULL,
    sender VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (sender) REFERENCES users(username)
);

CREATE INDEX idx_messages_room_id ON messages(room_id);
CREATE INDEX idx_messages_sender ON messages(sender);
