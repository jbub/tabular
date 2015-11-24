# tabular [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/jbub/tabular) [![Build Status](http://img.shields.io/travis/jbub/tabular.svg?style=flat-square)](https://travis-ci.org/jbub/tabular) [![Coverage Status](http://img.shields.io/coveralls/jbub/tabular.svg?style=flat-square)](https://coveralls.io/r/jbub/tabular)

Tabular data generator written in Go.

## Install

```bash
go get github.com/jbub/tabular
```

## Docs

http://godoc.org/github.com/jbub/tabular

## Example

```go
package main

import (
    "log"
    "os"

    "github.com/jbub/tabular"
)

func main() {
    set := tabular.NewDataSet()
    set.AddHeader("firstname", "First name")
    set.AddHeader("lastname", "Last name")

    r1 := tabular.NewRow("Julia", "Roberts")
    r2 := tabular.NewRow("John", "Malkovich")
    err := set.Append(r1, r2)
    if err != nil {
        log.Fatal(err)
    }

    opts := &tabular.CSVOpts{
        Comma: ';',
        UseCRLF: true,
    }
    csv := tabular.NewCSVWriter(opts)

    err = set.Write(csv, os.Stdout)
    if err != nil {
        log.Fatal(err)
    }
}
```

## CSV

```go
opts := &tabular.CSVOpts{
    Comma: ';',
    UseCRLF: true,
}
csvw := tabular.NewCSVWriter(opts)
```

### Output
```text
Firstname;Lastname;Age
Julia;Roberts;40
John;Malkovich;42
```

## HTML

```go
opts := &tabular.HTMLOpts{
    Caption:    "Popular people",
    Indent:     2,
    TableClass: "",
    HeadClass:  "",
    RowClass:   "",
    DataClass:  "",
}
htmlw := tabular.NewHTMLWriter(opts)
```

### Output

```html
<table>
  <caption>Popular people</caption>
  <thead>
    <tr>
      <th>Firstname</th>
      <th>Lastname</th>
      <th>Age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>Julia</td>
      <td>Roberts</td>
      <td>40</td>
    </tr>
    <tr>
      <td>John</td>
      <td>Malkovich</td>
      <td>42</td>
    </tr>
  </tbody>
</table>
```

## JSON

```go
opts := &tabular.JSONOpts{
    Indent: 4,
}
jsonw := tabular.NewJSONWriter(opts)
```

### Output

```json
[
    {
        "firstname": "Julia",
        "lastname": "Roberts",
        "age": "40"
    },
    {
        "firstname": "John",
        "lastname": "Malkovich",
        "age": "42"
    }
]
```

## LaTeX

```go
opts := &tabular.LatexOpts{
  Caption: "",
  Center: "",
  TabularX: false,
}
latexw := tabular.NewLatexWriter(opts)
```

### Output

```latex
\begin{table}[h]
\begin{tabular}{|l|l|l|}
\hline
Firstname & Lastname  & Age \\ \hline
Julia     & Roberts   & 40  \\ \hline
John      & Malkovich & 42  \\ \hline
\end{tabular}
\end{table}
```

## XML

```go
opts := &tabular.XMLOpts{
    Indent:     4,
    RowElem:    "row",
    ParentElem: "rows",
}
xmlw := tabular.NewXMLWriter(opts)
```

### Output

```xml
<rows>
    <row>
        <firstname>Julia</firstname>
        <lastname>Roberts</lastname>
        <age>40</age>
    </row>
    <row>
        <firstname>John</firstname>
        <lastname>Malkovich</lastname>
        <age>42</age>
    </row>
</rows>
```

## YAML

```go
opts := &tabular.YAMLOpts{}
yamlw := tabular.NewYAMLWriter(opts)
```

### Output

```yaml
- firstname: Julia
  lastname: Roberts
  age: 40
- firstname: John
  lastname: Malkovich
  age: 42
```

## SQL

```go
db, _ := sql.Open("postgres", "postgres://localhost/mydb?sslmode=disable")
opts := &tabular.SQLOpts{
    Driver: "postgres",
    DB:    db,
    Table: "my_table",
}
sqlw := tabular.NewSQLWriter(opts)
```

### Output

SQL queries are performed in transaction, these are example queries:

```sql
BEGIN
INSERT INTO my_table (firstname,lastname,age) VALUES ($1,$2,$3)
INSERT INTO my_table (firstname,lastname,age) VALUES ($1,$2,$3)
COMMIT
```