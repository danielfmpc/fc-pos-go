-- name: ListAllCategories :many
SELECT * FROM categories;

-- name: GetCategories :one
select * from categories
where id = ?;

-- name: CreateCategory :execresult
insert into categories (id, name, description) values (?, ?, ?);

-- name: UpdateCategory :exec
update categories set name = ?, description = ? where id = ?;

-- name: DeleteCategory :exec
delete from categories where id = ?;

-- name: CreateCourse :execresult
insert into courses (id, category_id, name, description, price) values (?, ?, ?, ?, ?);

-- name: ListCourses :many
select c.*, ca.name as category_name 
from courses c
inner join categories ca on c.category_id = ca.id;