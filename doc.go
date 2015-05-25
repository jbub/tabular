/*
Package tabular implements a tabular data generator.

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
*/
package tabular
