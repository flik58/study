# SQL

## web SQL practice

https://www.databasestar.com/sql-practice/

## MEMO

### SQL: How to select distinct on some columns

https://stackoverflow.com/questions/41302062/sql-how-to-select-distinct-on-some-columns
https://stackoverflow.com/questions/5021693/distinct-for-only-one-column
https://stackoverflow.com/questions/966176/select-distinct-on-one-column
MySQLだとgruop byだけでできるがprestoでやるのにハマって見つけたやつ。

パーティション付けたけど日付カラムほしいと言われたとき
concat(year,'-',month,'-',day) AS filepath_date
