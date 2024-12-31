package databse 

import ( 
	"log"
)

func RunMigrations() {
	log.Println("Running migrations")
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	categoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		category_id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
	);`

	postsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(user_id) REFERENCES users(user_id)
	);`

	postCategoriesTable := `
	CREATE TABLE IF NOT EXISTS post_categories (
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		PRIMARY KEY(post_id, category_id),
		FOREIGN KEY(post_id) REFERENCES posts(post_id),
		FOREIGN KEY(category_id) REFERENCES categories(category_id)
	);`

	execSQL(usersTable)
	execSQL(categoriesTable)
	execSQL(postsTable)
	execSQL(postCategoriesTable)
	log.Println("Migrations ran successfully")

}

func execSQL(sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("Migration failed: %v\nSQL: %s", err, statement)
	}
}


