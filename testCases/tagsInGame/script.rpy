# Unit test
# renpy-graphviz: INGAME_LABEL(start)
# renpy-graphviz: INGAME_JUMP(option_one)

# Below: creates a link from `start` to `option_two` even if there was a jump before -like normal jumps
# renpy-graphviz: INGAME_JUMP(option_two) 

label option_one:
    "dialogue"

# should follow the previous label
# renpy-graphviz: INGAME_LABEL(indirect_label)

    jump option_two

# renpy-graphviz: INGAME_LABEL(option_two)
# renpy-graphviz: INGAME_JUMP(ending)

# renpy-graphviz: INGAME_LABEL(first) TITLE keywords combination

# should create an indirect link from `first` to `second`
# renpy-graphviz: INGAME_LABEL(second) 

# should avoid creating a link from `first_break` to `second_break`
# renpy-graphviz: BREAK 
# renpy-graphviz: INGAME_LABEL(not_indirect)

