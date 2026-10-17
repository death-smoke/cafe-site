INSERT INTO categories (name, description) VALUES
 ('Салаты', 'Свежие и сытные салаты'),
 ('Супы', 'Горячие первые блюда'),
 ('Основное блюдо', 'Горячие основные блюда'),
 ('Напитки', 'Горячие и холодные напитки'),
 ('Десерты', 'Сладкое к чаю и кофе');

INSERT INTO public.images (id, path, attributes) VALUES
 (1, 'uploads/48b2b36639a6a36086c26737c238f470.png', '{"alt": "Греческий", "width": 800, "height": 800, "filename": "48b2b36639a6a36086c26737c238f470.png", "mime_type": "image/png", "size_bytes": 757126, "uploaded_at": "2026-10-17T13:41:41Z", "original_name": "salat_g.png"}');
 (2, 'uploads/285435fc9a0d8d5fdb097be1c1fda059.png', '{"alt": "Цезарь с курицей", "width": 800, "height": 800, "filename": "285435fc9a0d8d5fdb097be1c1fda059.png", "mime_type": "image/png", "size_bytes": 493603, "uploaded_at": "2026-10-17T13:41:52Z", "original_name": "salat_ces.png"}');
 (3, 'uploads/119e0fd97933cca72c21f9da84a8f9e7.png', '{"alt": "Борщ с говядиной", "width": 800, "height": 771, "filename": "119e0fd97933cca72c21f9da84a8f9e7.png", "mime_type": "image/png", "size_bytes": 1090295, "uploaded_at": "2026-10-17T13:42:02Z", "original_name": "borsh.png"}');
 (4, 'uploads/b269189dba8549a38241d4b383f3bd54.png', '{"alt": "Крем-суп грибной", "width": 800, "height": 800, "filename": "b269189dba8549a38241d4b383f3bd54.png", "mime_type": "image/png", "size_bytes": 824640, "uploaded_at": "2026-10-17T13:42:17Z", "original_name": "crem.png"}');
 (5, 'uploads/fff25baef1d8f7101fcda61639057a77.png', '{"alt": "Паста карбонара", "width": 800, "height": 571, "filename": "fff25baef1d8f7101fcda61639057a77.png", "mime_type": "image/png", "size_bytes": 508873, "uploaded_at": "2026-10-17T13:42:28Z", "original_name": "pasta.png"}');
 (6, 'uploads/d83976ce2e74b978eeebe0b71b75eedd.png', '{"alt": "Стейк из лосося", "width": 873, "height": 800, "filename": "d83976ce2e74b978eeebe0b71b75eedd.png", "mime_type": "image/png", "size_bytes": 987169, "uploaded_at": "2026-10-17T13:42:37Z", "original_name": "steik.png"}');
 (7, 'uploads/a1d70e16852bb5798bac091beb101bd9.png', '{"alt": "Домашний лимонад", "width": 800, "height": 1067, "filename": "a1d70e16852bb5798bac091beb101bd9.png", "mime_type": "image/png", "size_bytes": 1245527, "uploaded_at": "2026-10-17T13:42:47Z", "original_name": "limonade.png"}');
 (8, 'uploads/90882ecc38af8641bddb49e853b02a9d.png', '{"alt": "Капучино", "width": 800, "height": 688, "filename": "90882ecc38af8641bddb49e853b02a9d.png", "mime_type": "image/png", "size_bytes": 677171, "uploaded_at": "2026-10-17T13:42:59Z", "original_name": "coffe.png"}');
 (9, 'uploads/b068f91cdb383710a41cf8e722cefccb.png', '{"alt": "Тирамису", "width": 1086, "height": 800, "filename": "b068f91cdb383710a41cf8e722cefccb.png", "mime_type": "image/png", "size_bytes": 1187168, "uploaded_at": "2026-10-17T13:43:14Z", "original_name": "tiram.png"}');
 (10, 'uploads/0568b0359e00352f98269cb095fab6d5.png', '{"alt": "Чизкейк Нью-Йорк", "width": 800, "height": 533, "filename": "0568b0359e00352f98269cb095fab6d5.png", "mime_type": "image/png", "size_bytes": 462705, "uploaded_at": "2026-10-17T13:43:23Z", "original_name": "cake.png"}');

INSERT INTO menu (name, category_id, description, ingredients, price, image_id) VALUES
 ('Греческий', 1, 'Лёгкий овощной салат с сыром фета.', 'Огурцы, томаты, перец, красный лук, маслины, сыр фета, оливковое масло', 360, 1),
 ('Цезарь с курицей', 1, 'Классический салат с хрустящими гренками и соусом цезарь.', 'Куриное филе, салат романо, пармезан, гренки, соус цезарь, томаты черри', 420, 2),
 ('Борщ с говядиной', 2, 'Наваристый борщ со сметаной и чесночной пампушкой.', 'Говядина, свёкла, капуста, картофель, морковь, лук, томатная паста, сметана', 340, 3),
 ('Крем-суп грибной', 2, 'Нежный суп-пюре из шампиньонов.', 'Шампиньоны, сливки, лук, картофель, сливочное масло, гренки', 320, 4),
 ('Паста карбонара', 3, 'Спагетти в сливочном соусе с беконом.', 'Спагетти, бекон, яичный желток, пармезан, сливки, чёрный перец', 520, 5),
 ('Стейк из лосося', 3, 'Запечённый лосось с овощами гриль.', 'Филе лосося, кабачок, перец, лимон, оливковое масло, специи', 790, 6),
 ('Домашний лимонад', 4, 'Освежающий лимонад с мятой.', 'Лимон, лайм, мята, сахар, газированная вода, лёд', 220, 7),
 ('Капучино', 4, 'Эспрессо с нежной молочной пенкой.', 'Эспрессо, молоко', 210, 8),
 ('Тирамису', 5, 'Итальянский десерт с маскарпоне и кофе.', 'Маскарпоне, савоярди, эспрессо, яйца, какао', 380, 9);
 ('Чизкейк Нью-Йорк', 5, 'Классический чизкейк с ягодным соусом.', 'Сливочный сыр, печенье, сливочное масло, яйца, сахар, ягодный соус', 350, 10),

