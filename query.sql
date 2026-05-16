-- name: GetIncome :one
SELECT * FROM income
WHERE income_id  = ? LIMIT 1;

-- name: AllIncomes :many
select * from income;

-- name: InsertIncome :execresult
insert into income (
        name, 
        money,
        timestamp
)values (?, ?, ?);
