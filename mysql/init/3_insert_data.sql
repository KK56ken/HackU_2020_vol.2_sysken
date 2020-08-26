-- usersテーブルにインサートするもの
INSERT INTO raise_todo.users (name, password, feed_num) VALUES ("Nagoya","pass",1);

-- subjectsテーブルにインサートするもの
INSERT INTO raise_todo.subjects (name) VALUES ("国語");
INSERT INTO raise_todo.subjects (name) VALUES ("数学");
INSERT INTO raise_todo.subjects (name) VALUES ("英語");
INSERT INTO raise_todo.subjects (name) VALUES ("理科");
INSERT INTO raise_todo.subjects (name) VALUES ("社会");

-- tasksテーブルにインサートするもの
INSERT INTO raise_todo.tasks (user_id ,subject_id ,name, time_limit) VALUES (1,1,"書き取り　1ページ","9月1日");

-- charactersテーブルにインサートするもの
INSERT INTO raise_todo.characters (name, image_path,probability) VALUES ("beechild","./assets/beechild.png","0.2");
INSERT INTO raise_todo.characters (name, image_path,probability) VALUES ("beeadult","./assets/beeadult.png","0.2");
INSERT INTO raise_todo.characters (name, image_path,probability) VALUES ("beedie","./assets/beedie.png","0.2");
-- raisingsテーブルにインサートするもの
INSERT INTO raise_todo.raisings (user_id, character_id,ate_num,degree,background_level,raising_level) VALUES (1,1,0,7,1,1);

-- collectionsテーブルにインサートするもの
INSERT INTO raise_todo.collections (user_id,character_id,background_level,raising_level) VALUES (1,1,1,1);
