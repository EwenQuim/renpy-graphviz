
# Just unit tests
#renpy-graphviz: FAKE_LABEL(a)
#renpy-graphviz: FAKE_JUMP(b, c)

# Links between labels
#renpy-graphviz: FAKE_LABEL(d)
#renpy-graphviz: FAKE_LABEL(e)
#renpy-graphviz: FAKE_JUMP(d, f)
#renpy-graphviz: FAKE_JUMP(g, e)
#renpy-graphviz: FAKE_JUMP(f, g)

# No links between fake labels and real ones

label real_one:
# There will be no 'indirect link'
#renpy-graphviz: FAKE_LABEL(fake_one)

# There will be an indirect jump from ``real_one to `real_two`
label real_two:

# no jump from `real_two` to whatever
#renpy-graphviz: FAKE_JUMP(fake_two, fake_two_destination)
