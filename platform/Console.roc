module [log!]

import Effect

log!: Str => {}
log! = \str ->
    Effect.stdoutLine! str
