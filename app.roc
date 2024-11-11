app [main!] { r2e: platform "./platform/main.roc" }

import r2e.Console

main! = \{} ->
    Console.log! "wow"
