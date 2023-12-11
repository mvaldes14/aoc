file = open("inputs.txt")

red_max = 12
green_max = 13
blue_max = 14


def game_possible(line):
    for games in line.split(":")[1].split(";"):
        for play in games.split(","):
            num, color = play.strip().split(" ")
            if color == "green" and int(num) > green_max:
                return False
            if color == "red" and int(num) > red_max:
                return False
            if color == "blue" and int(num) > blue_max:
                return False
    return True


def max_cube_count(line):
    red = 0
    green = 0
    blue = 0
    for game in line.split(":")[1].split(";"):
        for play in game.split(","):
            num, color = play.strip().split(" ")
            if color == "green":
                green = max(int(num), green)
            if color == "red":
                red = max(int(num), red)
            if color == "blue":
                blue = max(int(num), blue)
    return red * green * blue


count_p2 = 0
count_p1 = 0
for id, line in enumerate(file, start=1):
    if game_possible(line):
        count_p1 += id
    count_p2 += max_cube_count(line)
print("p1",count_p1)
print("p2",count_p2)
