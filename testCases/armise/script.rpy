label Intro: #renpy-graphviz: SHAPE(invhouse)

    jump Noel
    jump Couloir
    jump Hermi_intro_1
    jump Sirene
    jump Quidditch
    jump Placard

label Quidditch:
label Ginny_intro_1: #renpy-graphviz: COLOR(orange)
label Ginny_intro_2: #renpy-graphviz: COLOR(orange)
#renpy-graphviz: BREAK

label Noel: #renpy-graphviz: SHAPE(doublecircle) COLOR(red)
    jump Noel_avec_Daphne
    jump Hermione_Daphne

label Couloir:
    jump Daphne_intro_1
    jump Luna_intro_1

label Daphne_intro_1: #renpy-graphviz: COLOR(green)
    jump Daphne_sexe
    jump Daphne_robe #renpy-graphviz: COLOR(green)
    jump Hermione_croise_Daphne

label Daphne_sexe: #renpy-graphviz: COLOR(green)
    jump Noel_avec_Daphne #renpy-graphviz: COLOR(green)

label Hermione_Daphne:
    jump Luna_degoutee

label Hermi_intro_1: #renpy-graphviz: COLOR(red)
label Hermi_intro_2: #renpy-graphviz: COLOR(red)
label Hermi_intro_3: #renpy-graphviz: COLOR(red)
label Hibou_pour_Luna:
    jump Luna_intro_1

label Luna_intro_1: #renpy-graphviz: COLOR(blue)
label Luna_intro_2: #renpy-graphviz: COLOR(blue)
label Hermione_hetero: #renpy-graphviz: COLOR(red)
    jump Hermione_croise_Daphne
    jump Rapport_a_Luna

label Hermione_croise_Daphne:
    jump Hermione_Daphne

label Rapport_a_Luna: #renpy-graphviz: COLOR(blue)
    jump Luna_degoutee

#renpy-graphviz: BREAK
label Placard:

label Creation_de_Pan: #renpy-graphviz: COLOR(yellow)

label Reve_de_Pan: #renpy-graphviz: COLOR(yellow)
    jump Lien_avec_Pan
    jump Reve_de_Pan

label Lien_avec_Pan: #renpy-graphviz: COLOR(yellow)
    jump Enquete_avec_pan #renpy-graphviz: COLOR(yellow)