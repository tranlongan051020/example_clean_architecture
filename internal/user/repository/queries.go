package repository

const (
	createUser = `INSERT INTO users (id, name, email, created_at, updated_at, delete_at) VALUES ($1, $2, $3, $4, $5, $6)`
)
