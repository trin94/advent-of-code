use crate::files;

#[derive(Clone)]
struct Stacks {
    stacks: Vec<Vec<char>>,
}

impl Stacks {
    pub fn new(stacks: Vec<Vec<char>>) -> Self {
        Stacks { stacks }
    }

    pub fn execute_unordered(&mut self, instruction: &Instruction) {
        let _move = instruction._move;
        let _from = self.stacks.get_mut(instruction._from).unwrap();
        let mut pops: Vec<char> = Vec::new();
        for _ in 0.._move {
            pops.push(_from.pop().unwrap())
        }
        let _to = self.stacks.get_mut(instruction._to).unwrap();
        for _ in 0.._move {
            _to.push(pops.remove(0))
        }
    }

    pub fn execute_ordered(&mut self, instruction: &Instruction) {
        let _move = instruction._move;
        let _from = self.stacks.get_mut(instruction._from).unwrap();
        let pops = _from.split_off(_from.len() - _move);
        let _to = self.stacks.get_mut(instruction._to).unwrap();
        _to.extend(pops);
    }

    pub fn stack_message(self) -> String {
        self.stacks.iter()
            .map(|stack| stack.last().unwrap_or(&' ').to_string())
            .collect::<Vec<String>>()
            .join("")
    }
}


struct Instruction {
    _move: usize,
    _from: usize,
    _to: usize,
}

impl Instruction {
    pub fn new(string: &String) -> Self {
        let instruction = string.split(" ").collect::<Vec<&str>>();
        Instruction {
            _move: instruction.get(1).unwrap().parse::<usize>().unwrap(),
            _from: instruction.get(3).unwrap().parse::<usize>().unwrap() - 1,
            _to: instruction.get(5).unwrap().parse::<usize>().unwrap() - 1,
        }
    }
}


fn prepare(lines: &Vec<String>) -> (Stacks, Vec<Instruction>) {
    let (stacks, instructions): (Vec<_>, Vec<_>) = lines
        .into_iter()
        .partition(|line| !line.starts_with("move"));

    let mut stacks: Vec<&String> = stacks
        .into_iter()
        .filter(|line| !line.is_empty())
        .rev()
        .collect();

    let count = stacks.remove(0)
        .split_whitespace()
        .last()
        .unwrap()
        .parse::<usize>()
        .unwrap();

    let stacks: Vec<_> = stacks.iter()
        .map(|tower| tower.chars().collect::<Vec<char>>())
        .map(|line| line.chunks(4).map(|ch| ch[1]).collect::<Vec<char>>())
        .collect();

    let stacks: Vec<Vec<char>> = (0..count)
        .map(|index| stacks.iter()
            .map(move |vertical| vertical.get(index).unwrap())
            .filter(|c| **c != ' ')
            .map(|c| *c)
            .collect::<Vec<char>>())
        .collect();

    let playground = Stacks::new(stacks);

    let instructions: Vec<Instruction> = instructions
        .into_iter()
        .map(|line| Instruction::new(line))
        .collect();

    return (playground, instructions);
}


pub fn solve() {
    let file = "resources/day05.txt";
    let lines = files::parse_lines_from(file);

    let (stacks, instructions) = prepare(&lines);

    let mut stack_task_1 = stacks.clone();
    let mut stack_task_2 = stacks.clone();

    for instruction in instructions {
        stack_task_1.execute_unordered(&instruction);
        stack_task_2.execute_ordered(&instruction);
    }

    println!("Unordered  {}", stack_task_1.stack_message());
    println!("Ordered    {}", stack_task_2.stack_message());
}
