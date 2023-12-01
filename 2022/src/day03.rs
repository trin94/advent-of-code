use std::ffi::c_uint;

use crate::files;

const LOWER_CASE_OFFSET: u32 = 96;
const UPPER_CASE_OFFSET: u32 = 38;


pub fn priority_of(char: char) -> u32 {
    if char.is_lowercase() {
        char as c_uint - LOWER_CASE_OFFSET
    } else {
        char as c_uint - UPPER_CASE_OFFSET
    }
}

pub fn find_same_char_from(strings: Vec<String>) -> char {
    let mut vec = strings.clone();
    let (intersection, others) = vec.split_at_mut(1);
    let intersection = &mut intersection[0];
    for other in others {
        intersection.retain(|e| other.contains(e));
    }
    intersection.pop().unwrap()
}

pub fn solve() {
    let file = "resources/day03.txt";
    let lines = files::parse_lines_from(file);

    let sum: u32 = lines.iter()
        .map(|line| {
            let compartment = line.split_at(line.len() / 2);
            let mut strings = Vec::new();
            strings.push(compartment.0.to_string());
            strings.push(compartment.1.to_string());
            let same_char = find_same_char_from(strings);
            priority_of(same_char)
        })
        .sum();

    println!("Sum of priorities: {}", sum);

    let chunks: Vec<&[String]> = lines.chunks(3).collect();

    let sum: u32 = chunks.into_iter().map(|chunk| {
        let strings = chunk.iter().map(|element| element.clone()).collect();
        let same_char = find_same_char_from(strings);
        priority_of(same_char)
    }).sum();

    println!("Sum of priorities: {}", sum);
}