label one:
    menu:
        "Which will you choose ?"
        "first choice"
            jump two 
        "second choice"
            eileen "Go to the second choice !"
            jump three
        "third choice"
            if condition:
                jump four
            else:
                jump five
        "Some Choice":
            eileen "This one does not jump"
        "Last choice":
            eileen "This one neither"

label two:
# renpy-graphviz: BREAK
label three:
# renpy-graphviz: BREAK
label four:
# renpy-graphviz: BREAK
label five:
label six: