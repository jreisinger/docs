Basic commands:

```
psql -U <username>     # connect to DB
\l                     -- list DBs
\dt                    -- list tables
\d <table>             -- table details
\x                     -- expanded display

SELECT * FROM <table>;
SELECT * FROM customer WHERE last_name LIKE '%gen%';
SELECT
   last_name,
   first_name
FROM
   customer
WHERE
   first_name = 'Jamie'
AND last_name = 'Rice';
```
