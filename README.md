# MySQL Base24

A small MySQL UDF library for making encoding/decoding base24 strings written in Golang.

## Usage

The base24 implementation used for this is the same that Microsoft used for their product keys, which uses the following characters, in this order:
```
2 3 4 6 7 8 9 b c d f g h j k m p q r t v w x y
```
The use of all lowercase letters means this is case insensitive. You can read more about it [here](https://news.ycombinator.com/item?id=22429809).

---

### `to_base24`

Converts the string argument to base24 encoded form and returns the result as a binary string. If the argument is not a string, it is converted to a string before conversion takes place. The result is `NULL` if the argument is `NULL`.

```sql
`to_base24` ( `string` )
```

 - `` `string` ``
   - The string to be encoded.

## Examples

```sql
select`to_base24`('abc');    -- 't8p76'
select`to_base24`(42);       -- 'y7r'
select`to_base24`('42');     -- 'y7r'
select`to_base24`(0x00abcd); -- '267cj'
select`to_base24`(null);     -- NULL
```
---

### `from_base24`

Takes a string encoded with the base24 encoded rules used by `to_base24` and returns the decoded result as a binary string. The result is NULL if the argument is NULL or not a valid base24 string.

```sql
`from_base24` ( `string` )
```

 - `` `string` ``
   - The string to be decoded.

## Examples

```sql
select`from_base24`('t8p76');       -- 'abc'
select`from_base24`('T8P76');       -- 'abc'
select`from_base24`('y7r');         -- '42'
select hex(`from_base24`('267cj')); -- '00ABCD'
select`from_base24`(null);          -- NULL
```
---

## Dependencies

You will need Golang, which you can get from here https://golang.org/doc/install.

Debian / Ubuntu

```shell
sudo apt update
sudo apt install libmysqlclient-dev
```

## Installing

You can find your MySQL plugin directory by running this MySQL query

```sql
select @@plugin_dir;
```

then replace `/usr/lib/mysql/plugin` below with your MySQL plugin directory.

```shell
cd ~ # or wherever you store your git projects
git clone https://github.com/StirlingMarketingGroup/mysql-base24.git
cd mysql-base24
go build -buildmode=c-shared -o base24.so
sudo cp base24.so /usr/lib/mysql/plugin/ # replace plugin dir here if needed
```

Enable the functions in MySQL by running this MySQL query

```sql
create function`to_base24`returns string soname'base24.so';
create function`from_base24`returns string soname'base24.so';
```