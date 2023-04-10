package repository

const (
	// users = `CREATE TABLE IF NOT EXISTS users(
	// 	user_id SERIAL PRIMARY KEY NOT NULL,
	// 	name CHAR(30) NOT NULL,
	// 	surname CHAR(50) NOT NULL,
	// 	email CHAR(50) UNIQUE NOT NULL,
	// 	password VARCHAR NOT NULL,
	// 	role CHAR(10) NOT NULL,
	// 	token VARCHAR NOT NULL,
	// 	token_duration TIMESTAMP DEFAULT NULL
	// )`

	// trainerSchedule = `CREATE TABLE IF NOT EXISTS trainer_schedule(
	// 	trainer_schedule_id SERIAL PRIMARY KEY NOT NULL,
	// 	trainer_id INT NOT NULL,
	// 	date DATE NOT NULL,
	// 	times JSON NOT NULL
	// );`

	// clientSchedule = `CREATE TABLE IF NOT EXISTS client_schedule(
	// 	client_schedule_id SERIAL PRIMARY KEY NOT NULL,
	// 	client_id INT NOT NULL,
	// 	date DATE NOT NULL,
	// 	time TIME NOT NULL
	// );`
	schema = `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		full_name CHAR(50) NOT NULL,
		email VARCHAR UNIQUE NOT NULL,
		password VARCHAR NOT NULL,
		role VARCHAR NOT NULL,
		token VARCHAR DEFAULT NULL,
		token_duration TIMESTAMP DEFAULT NULL
	);

	CREATE TABLE IF NOT EXISTS trainer_schedule (
		trainer_schedule_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		date TIMESTAMPTZ NOT NULL,
		available BOOLEAN NOT NULL DEFAULT true,
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	);
	
	CREATE TABLE IF NOT EXISTS client_schedule (
		client_schedule_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		trainer_id INT NOT NULL,
		date TIMESTAMPTZ NOT NULL,
		note VARCHAR,
		FOREIGN KEY (user_id) REFERENCES users (user_id),
		FOREIGN KEY (trainer_id) REFERENCES users (user_id)
	);

	CREATE TABLE IF NOT EXISTS trainer (
		trainer_id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		speciality CHAR(100) DEFAULT 'trainer',
		phone CHAR(30) NOT NULL DEFAULT 'NO',
		adress VARCHAR NOT NULL DEFAULT 'NO',
		image VARCHAR NOT NULL,
		bio VARCHAR NOT NULL DEFAULT 'NO',
		twitter CHAR(30) NOT NULL DEFAULT 'NO',
		instagram CHAR(30) NOT NULL DEFAULT 'NO',
		facebook CHAR(30) NOT NULL DEFAULT 'NO',
		FOREIGN KEY (user_id) REFERENCES users (user_id)
	);
	`
)
