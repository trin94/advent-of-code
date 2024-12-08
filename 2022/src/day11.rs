use std::collections::{BTreeMap, VecDeque};
use std::fmt;

use lazy_static::lazy_static;
use regex::Regex;

use crate::files;

struct Monkey {
    id: u64,
    items: VecDeque<u64>,
    operation: Box<dyn Fn(u64) -> u64>,
    divider: u64,
    test_passed_throw_to_id: u64,
    test_failed_throw_to_id: u64,

    inspect_counter: u64,
}

impl fmt::Display for Monkey {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Monkey(id={},items={:?},passed={},failed={},inspected={})",
               self.id, self.items, self.test_passed_throw_to_id, self.test_failed_throw_to_id, self.inspect_counter)
    }
}

impl Monkey {
    pub fn new_from(string: &str) -> Self {
        let string = string.trim();
        let id = {
            lazy_static! { static ref RE: Regex = Regex::new(r"Monkey (?P<id>\d+):").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let monkey_id: &str = &captures["id"];
            monkey_id.parse().unwrap()
        };
        let items = {
            lazy_static! { static ref RE: Regex = Regex::new(r"Starting items: (?P<items>.*?)\n").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let monkey_id: &str = &captures["items"];
            monkey_id.split(",")
                .map(|id| id.trim().parse::<u64>().unwrap())
                .collect::<VecDeque<u64>>()
        };
        let operation: Box<dyn Fn(u64) -> u64> = {
            lazy_static! { static ref RE: Regex = Regex::new(r"Operation: new = old (?P<operation>[*+]) (?P<other>.*?)\n").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let operation: &str = &captures["operation"];
            let other: &str = &captures["other"];
            match operation {
                "+" => {
                    let other = other.trim().parse::<u64>().unwrap();
                    Box::new(move |x: u64| x + other)
                }
                "*" => {
                    match other {
                        "old" => Box::new(|x: u64| x * x),
                        factor => {
                            let other = factor.trim().parse::<u64>().unwrap();
                            Box::new(move |old: u64| old * other)
                        }
                    }
                }
                _ => unreachable!()
            }
        };
        let divider = {
            lazy_static! { static ref RE: Regex = Regex::new(r"Test: divisible by (?P<divider>.*?)\n").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let divider: &str = &captures["divider"];
            divider.trim().parse::<u64>().unwrap()
        };
        let test_passed_throw_to_id = {
            lazy_static! { static ref RE: Regex = Regex::new(r"If true: throw to monkey (?P<id>.*?)\n").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let id: &str = &captures["id"];
            id.trim().parse::<u64>().unwrap()
        };
        let test_failed_throw_to_id = {
            lazy_static! { static ref RE: Regex = Regex::new(r"If false: throw to monkey (?P<id>.*?)$").unwrap(); }
            let captures = RE.captures(string).unwrap();
            let id: &str = &captures["id"];
            id.trim().parse::<u64>().unwrap()
        };
        let inspect_counter = 0;
        Monkey { id, items, operation, divider, test_passed_throw_to_id, test_failed_throw_to_id, inspect_counter }
    }

    /// Monkey inspects the item until he gets bored. Returns a tuple with new worry_level and receiver.
    pub fn inspect(&mut self, item: u64, relief_factor: u64, lcm: &u64) -> (u64, u64) {
        self.inspect_counter += 1;

        let worry_level: u64 = item % lcm;
        let worry_level: u64 = (self.operation)(worry_level);
        let worry_level: u64 = worry_level / relief_factor;

        if worry_level % self.divider == 0 {
            (worry_level, self.test_passed_throw_to_id)
        } else {
            (worry_level, self.test_failed_throw_to_id)
        }
    }
}

pub fn solve() {
    let file = "resources/day11.txt";
    let string = files::parse_string_from(file).unwrap();

    let relief_factor = 3;
    let rounds = 20;
    calculate_monkey_business(&string, relief_factor, rounds);

    let relief_factor = 1;
    let rounds = 10000;
    calculate_monkey_business(&string, relief_factor, rounds);
}

fn calculate_monkey_business(string: &String, relief_factor: u64, rounds: u64) {
    let monkeys = string.split("\n\n")
        .map(|line| Monkey::new_from(line))
        .map(|monkey| (monkey.id, monkey))
        .collect::<Vec<(u64, Monkey)>>();

    let prime_lcm: u64 = monkeys.iter()
        .map(|it| it.1.divider)
        .product();

    let mut circus: BTreeMap<u64, Monkey> = monkeys.into_iter().collect();
    let keys = circus.keys().cloned().collect::<Vec<u64>>();

    for _ in 0..rounds {
        for key in keys.iter() {
            let monkey = circus.get_mut(&key).unwrap();
            let mut items = monkey.items.clone();
            monkey.items.clear();
            while let Some(current_item) = items.pop_front() {
                let monkey = circus.get_mut(&key).unwrap();
                let (new_item, next_owner) = monkey.inspect(current_item, relief_factor, &prime_lcm);
                let next_owner = circus.get_mut(&next_owner).unwrap();
                next_owner.items.push_back(new_item);
            }
        }
    }

    let mut inspect_counters = circus.values()
        .map(|monkey| monkey.inspect_counter)
        .collect::<Vec<u64>>();

    inspect_counters.sort();

    let monkey_business: u64 = inspect_counters.iter()
        .rev()
        .take(2)
        .product();

    println!("Monkey Business {}", monkey_business);
}
