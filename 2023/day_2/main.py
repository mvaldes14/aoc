file = open("inputs.txt")

red_max = 12
green_max = 13
blue_max = 14
count = 0


def game_possible(line):
    for games in line.split(":")[1].split(";"):
        for play in games.split(","):
            num, color = play.strip().split(" ")
            if color == "green" and int(num) > green_max:
                print(id, num, color)
                return False
            if color == "red" and int(num) > red_max:
                print(id, num, color)
                return False
            if color == "blue" and int(num) > blue_max:
                print(id, num, color)
                return False
    return True


for id, line in enumerate(file, start=1):
    if game_possible(line):
        count += id

print(count)
