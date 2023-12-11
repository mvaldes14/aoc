data = [i.strip() for i in open("./input.txt")]

r = []
for i in ("\n".join(data)).split("\n\n"):
    c = 0
    for x in i.split("\n"):
        c += int(x)
    r.append(c)
R = sorted(r)

print(max(R))
print(R[-1] + R[-2] + R[-3])

