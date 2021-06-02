label first:
    "bla"
    label useless_should_be_removed:
        "some things"
    "continue"
label second:
    label selfLoop:
        "self"
        menu:
            "repeat":
                jump selfLoop
            "continue":
                pass

    jump third

label third:
    label inception:

        jump inception2

        label inception2:

            jump inception

        jump inception3

    jump otherLabelFromThird