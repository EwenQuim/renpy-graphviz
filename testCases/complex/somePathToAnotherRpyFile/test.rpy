# Statrs here
label staAA_aa6rt     : # renpy-graphviz: TiTlE
    scene bg bespin
    show darkv sabre
    play music "imperial_march.mp3"
    d "Je suis ton père!"

menu:
    "Ah! D'accord, si vous le dites...":
         jump routeone

    "Vous êtes sûr?":
         jump route2

    "More":
        jump routeAlternative # comments do nothing