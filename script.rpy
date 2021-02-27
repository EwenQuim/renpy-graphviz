# Vous pouvez placer le script de votre jeu dans ce fichier.

# Déclarez sous cette ligne les images, avec l'instruction 'image'
# ex: image eileen heureuse = "eileen_heureuse.png"

# Déclarez les personnages utilisés dans le jeu.
image bg bespin = "bespin.png"

image darkv sabre = "darth.png"

define d = Character('Dark Vador', color="#c8ffc8")


label routeone:
    d "Ensemble, nous règnerons sur la Galaxie!."
    jump final

label route2:
    d "Puisque je te le dis..."
    jump final

label final:
    d "Tu connaîtras la suite... quand tu l'auras écrite!"
    return