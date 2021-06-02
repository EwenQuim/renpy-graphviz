screen test(argument1, arg2=default_value):

    use nested_screen

    if condition:
        action Jump("label_continue")
    else:
        action Show("other_screen")

screen screen2():
    use nested_screen
        action Call("new_screen")


label game:
    "bla"
    narrator "Hello World"
    call screen new_screen()

    jump next

    call screen other_screen()
