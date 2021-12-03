# read input.txt and add them to a list


def read_input():
    with open('input.txt') as f:
        return [line.strip() for line in f]


def get_position():
    ''' get the position of the submarine, based on input
    '''
    instructions = read_input()

    instructions_map = {}

    for i in range(len(instructions)):
        instructions[i] = instructions[i].split()

        if instructions[i][0] not in instructions_map:
            instructions_map[instructions[i][0]] = 0
        
        instructions_map[instructions[i][0]] += int(instructions[i][1])
        
    depth = instructions_map.get('down', 0) - instructions_map.get('up', 0)
    horizonal_pos = instructions_map.get('forward', 0)

    solution = depth * horizonal_pos

    return solution

def get_advanced_position():
    ''' get the position of the submarine, keeping aim in mind
    '''
    aim = 0
    depth = 0
    horizontal_pos = 0
    instructions = read_input()

    # down adds to aim, up subtracts from aim
    # forward multiplies aim by horizontal_pos

    for i in range(len(instructions)):
        instructions[i] = instructions[i].split()

        if instructions[i][0] == 'down':
            aim += int(instructions[i][1])
        elif instructions[i][0] == 'up':
            aim -= int(instructions[i][1])
        elif instructions[i][0] == 'forward':
            horizontal_pos += int(instructions[i][1])
            if aim != 0:
                depth += aim * int(instructions[i][1])

    return depth * horizontal_pos


if '__main__' == __name__:
    print("First puzzle", get_position())
    print("Second puzzle", get_advanced_position())

