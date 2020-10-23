-- name: ListAnimals :many
SELECT * FROM animals
ORDER BY created_at;

-- name: CreateAnimal :exec
INSERT INTO animals (name, type) VALUES ($1, $2);