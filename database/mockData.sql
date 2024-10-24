INSERT INTO products (product_id, user_id, name, url, price, description, title)
VALUES
(1, 1, 'Awesome Widget', 'https://images.unsplash.com/photo-1597740985671-2a8a3b80502e?q=80&w=1887&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D', 20, 'This widget will solve all your problems!', 'Widget Pro Max'),
(2, 1, 'Superb Gadget', 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8cHJvZHVjdHN8ZW58MHx8MHx8fDA%3D', 25, 'A perfect tool for your daily needs.', 'Gadget Elite'),
(3, 1, 'Amazing Tool', 'https://images.unsplash.com/photo-1585201731775-0597e1be4bfb?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fHRvb2xzfGVufDB8fDB8fHww', 30, 'Makes your work easier and faster!', 'Tool Pro Ultra'),
(4, 1, 'Incredible Device', 'https://images.unsplash.com/photo-1629131726692-1accd0c53ce0?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8ZGV2aWNlfGVufDB8fDB8fHww', 35, 'Experience technology like never before.', 'Device Supreme'),
(5, 1, 'Perfect Appliance', 'https://images.unsplash.com/photo-1626806787461-102c1bfaaea1?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OHx8YXBwbGlhbmNlfGVufDB8fDB8fHww', 40, 'Enhances your home and life.', 'Appliance Max'),
(6, 1, 'Exceptional Object', 'https://images.unsplash.com/photo-1459411552884-841db9b3cc2a?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8b2JqZWN0fGVufDB8fDB8fHww', 45, 'Unmatched in quality and functionality.', 'Object Deluxe'),
(7, 1, 'Ultimate Thing', 'https://plus.unsplash.com/premium_photo-1687558246422-e94f0864d467?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NXx8dGhpbmd8ZW58MHx8MHx8fDA%3D', 50, 'The last thing you will ever need!', 'Thing Ultimate'),
(8, 1, 'Prime Item', 'https://images.unsplash.com/photo-1494232410401-ad00d5433cfa?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTB8fHRoaW5nfGVufDB8fDB8fHww', 55, 'Top choice for professionals.', 'Item Prime'),
(9, 1, 'Super Item', 'https://plus.unsplash.com/premium_photo-1688678097910-706dbce46fa4?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTd8fHRoaW5nfGVufDB8fDB8fHww', 60, 'Super performance, super results.', 'Item Super'),
(10, 1, 'Mega Product', 'https://images.unsplash.com/photo-1565252503738-0d47de5d50e8?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjR8fHRoaW5nfGVufDB8fDB8fHww', 65, 'Bigger and better than ever before.', 'Product Mega'),
(11, 1, 'Hyper Gizmo', 'https://images.unsplash.com/photo-1506899689252-707a63f4c766?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzF8fHRoaW5nfGVufDB8fDB8fHww', 70, 'Hyper-efficient, hyper-useful.', 'Gizmo Hyper'),
(12, 1, 'Elite Contraption', 'https://images.unsplash.com/photo-1523132797263-747d5d0dbbb3?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MzZ8fHRoaW5nfGVufDB8fDB8fHww', 75, 'Crafted for the elite.', 'Contraption Elite');

INSERT INTO users (username, password_hash, role)
VALUES
    ('user1', '$2a$10$BwxIFdnfimG7K5PYzpggXeT4E8g6sSf8I1dudgjaMhzt0ZZLL0VEi', 'seller'),
    ('user2', '$2a$10$BwxIFdnfimG7K5PYzpggXeT4E8g6sSf8I1dudgjaMhzt0ZZLL0VEi', 'buyer'),
    ('user3', '$2a$10$BwxIFdnfimG7K5PYzpggXeT4E8g6sSf8I1dudgjaMhzt0ZZLL0VEi', 'seller'),
    ('user4', '$2a$10$BwxIFdnfimG7K5PYzpggXeT4E8g6sSf8I1dudgjaMhzt0ZZLL0VEi', 'buyer'),
    ('user5', '$2a$10$BwxIFdnfimG7K5PYzpggXeT4E8g6sSf8I1dudgjaMhzt0ZZLL0VEi', 'seller');
