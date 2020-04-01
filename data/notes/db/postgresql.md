Basic commands:

```
psql -U <username>     # connect to DB
\l                     -- list DBs
\dt                    -- list tables
\d <table>             -- table details
\x                     -- expanded display

select * from <table>;

SELECT
   last_name,
   first_name
FROM
   customer
WHERE
   first_name = 'Jamie'
AND last_name = 'Rice';
```
