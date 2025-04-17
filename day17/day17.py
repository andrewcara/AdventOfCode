register_a = 0
register_b = 0
register_c = 0


def get_combo(val):
    combo = {0: 0, 1: 1, 2: 2, 3: 3, 4: register_a, 5: register_b, 6: register_c}
    return combo[val]


def adv(val: int) -> None:
    global register_a
    adv_val = (register_a) // (2 ** get_combo(val))
    register_a = adv_val


def bxl(val: int) -> int:
    global register_b
    register_b = val ^ register_b


def bst(val: int) -> int:
    global register_b
    register_b = get_combo(val) % 8


def jnz(opcode: int) -> int:
    global register_a
    if register_a != 0:
        return opcode
    return -1


def bxc(val: int) -> int:
    global register_b, register_c
    register_b = register_b ^ register_c


def out(val: int) -> int:
    return get_combo(val) % 8


def bdv(val: int) -> None:
    global register_a, register_b
    adv_val = (register_a) // (2 ** get_combo(val))
    register_b = adv_val


def cdv(val: int) -> None:
    global register_a, register_c
    adv_val = (register_a) // (2 ** get_combo(val))
    register_c = adv_val


opcodes = {0: adv, 1: bxl, 2: bst, 3: jnz, 4: bxc, 5: out, 6: bdv, 7: cdv}


def get_register_value(filename: str, register_name: str) -> int | None:
    with open(filename, "r") as f:
        lines = f.readlines()
        for line in lines:
            if line.startswith(register_name + ":"):
                # Extract the integer value of the register
                return int(line.split(":")[1].strip())
    return None


# Function to read the program as a list of integers
def get_program_as_list(filename: str) -> list:
    with open(filename, "r") as f:
        lines = f.readlines()
        for line in lines:
            if line.startswith("Program:"):
                # Extract the program and convert it to a list of integers
                return list(map(int, line.split(":")[1].strip().split(",")))
    return []


def part1(program: list) -> str:
    instruction = 0
    output_str = ""
    loops = 0

    while instruction < len(program) - 1:
        func = opcodes[program[instruction]]
        result = func(program[instruction + 1])

        if func.__name__ == "jnz" and result != -1:
            instruction = result
        else:
            instruction += 2
            if result is not None and result != -1:
                output_str += "," + str(result) if output_str else str(result)
        loops += 1
    return output_str

def starts_with_bits(number, prefix):
    # Get the number of bits in the prefix
    prefix_bits = prefix.bit_length()
    shifted = number >> (number.bit_length() - prefix_bits)
    
    # Compare with prefix
    return shifted == prefix

def part2(program: list, greatest_current_bin_number=None) -> int:
    global register_a, register_b, register_c
    starting_a = register_a
    program_string = ','.join(map(str, program)) 
    
    while True:
        if greatest_current_bin_number is None or (starting_a, greatest_current_bin_number):
            output_str = part1(program)
            if output_str == program_string[0:len(output_str)]:
                greatest_current_bin_number = starting_a
                print(starting_a, output_str, greatest_current_bin_number)
        starting_a+=1
        register_a = starting_a
        
    

def main() -> None:

    global register_a, register_b, register_c
    filename = "input.txt"
    register_a = get_register_value(filename, "Register A")
    register_b = get_register_value(filename, "Register B")
    register_c = get_register_value(filename, "Register C")
    program = get_program_as_list(filename)
    part1_result = part1(program)
    print(part1_result)
    part2(program)


if __name__ == "__main__":
    main()
