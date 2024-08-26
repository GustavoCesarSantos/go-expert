package main

import (
    "log"

    checkouts "checkouts/cmd/server"
    payments "payments/cmd/server"
)

func main() {
    log.Println("running application")

    go func() {
        checkouts.Serve()
    }()

    payments.Serve()
}
