use std::cmp;
use std::cmp::Ordering;
use std::collections::VecDeque;

use crate::files;

const RADIX: u32 = 10;


#[derive(Clone, Debug, Eq, PartialEq)]
enum Value {
    INT(u32),
    LIST(Vec<Value>),
}

impl Value {
    pub fn from(string: &str) -> Self {
        let mut parser = ValueParser::from(string);
        parser.parse()
    }

    fn compare(left: &Value, right: &Value) -> Option<bool> {
        match left {
            Value::INT(left) => match right {
                Value::INT(right) => Value::comp_int_int(left, right),
                Value::LIST(right) => Value::comp_int_list(left, right.clone())
            }
            Value::LIST(left) => match right {
                Value::INT(right) => Value::comp_list_int(left.clone(), right),
                Value::LIST(right) => Value::comp_list_list(left, right)
            }
        }
    }

    fn comp_int_int(left: &u32, right: &u32) -> Option<bool> {
        if left == right {
            None
        } else {
            Some(left < right)
        }
    }

    fn comp_int_list(left: &u32, right: Vec<Value>) -> Option<bool> {
        let left = Value::LIST(vec![Value::INT(*left)]);
        let right = Value::LIST(right);
        Value::compare(&left, &right)
    }

    fn comp_list_int(left: Vec<Value>, right: &u32) -> Option<bool> {
        let left = Value::LIST(left);
        let right = Value::LIST(vec![Value::INT(*right)]);
        Value::compare(&left, &right)
    }

    fn comp_list_list(left: &Vec<Value>, right: &Vec<Value>) -> Option<bool> {
        let longest = cmp::max(left.len(), right.len());
        for index in 0..longest {
            if let Some(right) = right.get(index) {
                if let Some(left) = left.get(index) {
                    if let Some(result) = Value::compare(left, right) {
                        return Some(result);
                    }
                } else {
                    return Some(true);
                }
            } else {
                return Some(false);
            }
        }
        None
    }
}

impl PartialOrd<Self> for Value {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Value {
    fn cmp(&self, other: &Self) -> Ordering {
        match Value::compare(self, other) {
            None => Ordering::Equal,
            Some(result) => match result {
                true => Ordering::Less,
                false => Ordering::Greater
            }
        }
    }
}


struct ValuePair {
    left: Value,
    right: Value,
}

impl ValuePair {
    pub fn from(string: &str) -> Self {
        let mut split = string.splitn(2, "\n");
        Self {
            left: Value::from(split.next().unwrap()),
            right: Value::from(split.next().unwrap()),
        }
    }

    pub fn is_ordered(&self) -> bool {
        Value::compare(&self.left, &self.right).unwrap()
    }
}


struct ValueParser {
    characters: VecDeque<char>,
    consumed: u32,
}

impl ValueParser {
    pub fn from(string: &str) -> Self {
        Self { characters: string.chars().collect::<VecDeque<char>>(), consumed: 0 }
    }

    fn new(chars: VecDeque<char>) -> Self {
        Self { characters: chars, consumed: 0 }
    }

    pub fn parse(&mut self) -> Value {
        let mut values: Vec<Value> = vec![];
        while let Some(next) = self.next() {
            if next.is_digit(RADIX) {
                values.push(self.next_number(next))
            } else if next == '[' {
                values.push(self.next_list())
            } else if next == ']' {
                break;
            }
        }
        return Value::LIST(values);
    }

    fn next_number(&mut self, beginning: char) -> Value {
        let mut number: Vec<char> = vec![beginning];
        while let Some(next_next) = self.peek() {
            if next_next.is_digit(RADIX) {
                number.push(self.next().unwrap());
            } else {
                break;
            }
        }
        let number = number.into_iter().collect::<String>().parse::<u32>().unwrap();
        Value::INT(number)
    }

    fn next_list(&mut self) -> Value {
        let mut parser = ValueParser::new(self.characters.clone());
        let value = parser.parse();
        self.skip(parser.consumed);
        value
    }

    fn skip(&mut self, characters: u32) {
        for _ in 0..characters { self.next(); }
    }

    fn next(&mut self) -> Option<char> {
        self.consumed += 1;
        self.characters.pop_front()
    }

    fn peek(&self) -> Option<&char> {
        self.characters.get(0)
    }
}


pub fn solve() {
    let file = "resources/day13.txt";
    let string = files::parse_string_from(file).unwrap();

    part_1(&string);
    part_2(&string);
}

fn part_1(string: &String) {
    let value_pairs = string
        .trim()
        .split("\n\n")
        .map(|double_line| ValuePair::from(double_line))
        .collect::<Vec<ValuePair>>();

    let mut index_sum = 0;

    for (index, value_pair) in value_pairs.iter().enumerate() {
        if value_pair.is_ordered() {
            index_sum += index + 1;
        }
    }

    println!("Sum of indices: {}", index_sum)
}

fn part_2(string: &String) {
    let mut values = string.trim()
        .split("\n")
        .filter(|line| !line.trim().is_empty())
        .map(|line| Value::from(line))
        .collect::<Vec<Value>>();

    let marker_1 = Value::from("[[2]]");
    let marker_2 = Value::from("[[6]]");
    let markers = vec![marker_1.clone(), marker_2.clone()];

    values.extend(markers);
    values.sort();

    let idx_marker_1 = values.iter().position(|r| r == &marker_1).unwrap() + 1;
    let idx_marker_2 = values.iter().position(|r| r == &marker_2).unwrap() + 1;

    println!("Decoder key: {}", idx_marker_1 * idx_marker_2)
}