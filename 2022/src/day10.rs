use crate::files;

enum Operation {
    NOOP,
    ADD(i32),
}

type Cycle = (i32, i32);  // (cycle, 💪 during)

fn parse() -> Vec<Operation> {
    files::parse_lines_from("resources/day10.txt")
        .iter()
        .map(|line| match line.split_at(4) {
            ("noop", _) => Operation::NOOP,
            ("addx", increment) => Operation::ADD(increment.trim().parse::<i32>().unwrap()),
            (_, _) => unreachable!(),
        })
        .collect()
}

fn parse_signal_history(operations: &Vec<Operation>) -> Vec<Cycle> {
    let mut strength: i32 = 1;
    let mut cycle: i32 = 1;
    let mut signal_history: Vec<Cycle> = Vec::new();
    for operation in operations {
        match operation {
            Operation::NOOP => {
                signal_history.push((cycle, strength));
                cycle += 1;
            }
            Operation::ADD(increment) => {
                signal_history.push((cycle, strength));
                cycle += 1;
                signal_history.push((cycle, strength));
                cycle += 1;
                strength += increment;
            }
        }
    }
    signal_history
}

struct Sprite {
    position: i32,
    size: i32,
}

impl Sprite {
    pub fn new(size: i32) -> Self {
        assert!(size > 0);
        Sprite { size, position: 1 }
    }

    pub fn move_to(&mut self, position: i32) {
        self.position = position;
    }

    pub fn covers(&self, position: i32) -> bool {
        let additional = (self.size - 1) / 2;
        position >= self.position - additional && position <= self.position + additional
    }
}


pub fn solve() {
    let operations = parse();
    let signal_history: Vec<_> = parse_signal_history(&operations);

    let mut relevant_cycles = [20, 60, 100, 140, 180, 220];
    let mut sum: i32 = 0;
    for cycle in relevant_cycles {
        let snapshot = signal_history[cycle - 1];
        let cycle = snapshot.0;
        let during = snapshot.1;
        sum += cycle * during;
    }
    println!("Sum of relevant cycles: {}", sum);

    println!("Screen output:");

    let mut sprite = Sprite::new(3);
    for cycle in signal_history {
        let index = cycle.0 - 1;
        let position = index % 40;
        let strength = cycle.1;

        if index % 40 == 0 { print!("\n") }

        sprite.move_to(strength);
        print!("{}", if sprite.covers(position) { '#' } else { '.' });
    }

    print!("\n")
}