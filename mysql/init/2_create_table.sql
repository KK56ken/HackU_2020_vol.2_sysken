USE raise_todo;

CREATE TABLE IF NOT EXISTS raise_todo.users(
  user_id INT NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  token VARCHAR(255),
  feed_num INT NOT NULL,
  register_date DATETIME NOT NULL default current_timestamp,
  PRIMARY KEY (`user_id`)
);

CREATE TABLE IF NOT EXISTS raise_todo.subjects(
  subject_id INT NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (`subject_id`)
);

CREATE TABLE IF NOT EXISTS raise_todo.tasks(
  task_id INT NOT NULL auto_increment,
  user_id INT NOT NULL,
  subject_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  time_limit VARCHAR(255) NOT NULL,
  end_flag INT NOT NULL,
  register_date DATETIME NOT NULL default current_timestamp,
  PRIMARY KEY (`task_id`)
);

CREATE TABLE IF NOT EXISTS raise_todo.characters(
  character_id INT NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL,
  image_path VARCHAR(255) NOT NULL,
  probability INT NOT NULL,
  PRIMARY KEY (`character_id`)
);

CREATE TABLE IF NOT EXISTS raise_todo.raisings(
  raising_id INT NOT NULL auto_increment,
  user_id INT NOT NULL,
  character_id INT NOT NULL,
  ate_num INT NOT NULL,
  degree INT NOT NULL,
  background_level INT NOT NULL,
  raising_level INT NOT NULL,
  register_date DATETIME NOT NULL default current_timestamp,
  PRIMARY KEY (`raising_id`)
);

CREATE TABLE IF NOT EXISTS raise_todo.collections(
  collection_id INT NOT NULL auto_increment,
  user_id INT NOT NULL,
  character_id INT NOT NULL,
  background_level INT NOT NULL,
  raising_level INT NOT NULL,
  PRIMARY KEY (`collection_id`)
);
