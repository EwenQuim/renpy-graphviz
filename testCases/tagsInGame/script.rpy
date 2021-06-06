# Unit test
# renpy-graphviz: INGAME_LABEL(0, start)
# renpy-graphviz: INGAME_JUMP(4, option_one)

# Below: creates a link from `start` to `option_two` even if there was a jump before -like normal jumps
# renpy-graphviz: INGAME_JUMP(4, option_two) 

label option_one:
    "dialogue"

# should follow the previous label
# renpy-graphviz: INGAME_LABEL(0, indirect_label)

    jump option_two

# renpy-graphviz: INGAME_LABEL(0, option_two)
# renpy-graphviz: INGAME_JUMP(4, ending)

# renpy-graphviz: INGAME_LABEL(0, first) TITLE keywords combination

# should create an indirect link from `first` to `second`
# renpy-graphviz: INGAME_LABEL(0, second) 

# should avoid creating a link from `first_break` to `second_break`
# renpy-graphviz: BREAK 
# renpy-graphviz: INGAME_LABEL(0, not_indirect)

