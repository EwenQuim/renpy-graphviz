label start:
    "Is my VN simple ?"
    menu:
        "yes it is":
            jump simple_ending
        "no it isn't":
            jump complexe

label complexe:
    eileen "There are a lot of unexpected things, but at the end..."

label simple_ending: # label complex virtually "jumps" into simple_ending
    eileen "The END !"