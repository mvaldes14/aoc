file = open("input.txt").readlines()


for line in file:
    l = line.split(".")
    print(",".join(l))


