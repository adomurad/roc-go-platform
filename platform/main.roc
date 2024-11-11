platform ""
    requires {} { main! : {} => {} }
    exposes [
        Console,
    ]
    packages {}
    imports []
    provides [mainForHost!]

mainForHost! : {} => I32
mainForHost! = \{} ->
    main! {}

    0
