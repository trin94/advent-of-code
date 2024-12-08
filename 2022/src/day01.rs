use crate::files;

mod problem {
    pub struct Elf {
        calories: Vec<u32>,
    }

    impl Elf {
        pub fn new() -> Self {
            Elf { calories: Vec::new() }
        }

        pub fn calories(&self) -> u32 {
            return self.calories.iter().sum();
        }
    }

    pub fn parse_elves_from(lines: Vec<String>) -> Vec<Elf> {
        let mut elves = Vec::new();
        let mut current = Elf::new();
        for line in lines {
            if line.is_empty() {
                elves.push(current);
                current = Elf::new();
            } else {
                let food = line.parse::<u32>().unwrap();
                current.calories.push(food)
            }
        }
        return elves;
    }
}


pub fn solve() {
    let file = "resources/day01.txt";

    let lines = files::parse_lines_from(file);
    let mut elves = problem::parse_elves_from(lines);

    elves.sort_by(|a, b| a.calories().cmp(&b.calories()));

    let calories = elves.last().unwrap().calories();
    println!("Fattest 1 elf:   {} calories", calories);

    let elf3 = elves.pop().unwrap();
    let elf2 = elves.pop().unwrap();
    let elf1 = elves.pop().unwrap();

    let sum = elf1.calories() + elf2.calories() + elf3.calories();

    println!("Fattest 3 elves: {} calories", sum)
}

