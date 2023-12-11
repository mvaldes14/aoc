def findNumbers(line):
    digits = (
        "zero",
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    )
    numbers = []
    for i, c in enumerate(line):
        if c.isdigit():
            numbers.append(int(c))
            continue
        for n, name in enumerate(digits):
            if line[i:].startswith(name):
                numbers.append(n)
                break
    return numbers[0] * 10 + numbers[-1]


if "__main__" == __name__:
    print(sum(findNumbers(line) for line in open("inputs.txt").readlines()))
