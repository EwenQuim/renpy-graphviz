image bg bespin = "bespin.png"

# renpy-graphviz: TITLE
image darkv sabre = "darth.png"

define d = Character('Dark Vador', color="#c8ffc8")


label routeone :  # renpy-graphviz: TiTlE
    d "We will conquer the galaxy!"
    if condition:
        jump bad_ending
    else:
        d "aussi"

    label notImportant_shouldbeIgnored: # renpy-graphviz : IGNORE

        d "this happens" 

label routeAlternative:
    d "The flow should be broken by the jump"
    jump good_ending

label route2:
    d "Puisque je te le dis..."

label bad_ending:  # renpy-graphviz : GAMEOVER 
    d "Bad ending"
    return

# renpy-graphviz : BREAK 
label good_ending: # renpy-graphviz : GAMEOVER TITLE
    d "Good ending"
