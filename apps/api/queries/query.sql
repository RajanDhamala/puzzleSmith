-- name: RegisterUser :one
INSERT INTO users (fullname,email,password)
VALUES ($1,$2,$3)
RETURNING _id;

-- name: CheckIfusrExists :one
SELECT _id FROM users WHERE email=$1;

-- name: LoginUser :one
SELECT _id, fullname,email,password FROM users WHERE email=$1;

-- name: ChangePassword :exec
UPDATE users SET password=$1 WHERE _id=$2;


