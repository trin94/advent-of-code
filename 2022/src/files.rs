#![allow(dead_code)]

use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader, Lines, Read};
use std::path::Path;

pub fn parse_lines_from<P>(file: P) -> Vec<String> where P: AsRef<Path> {
    let mut lines: Vec<String> = Vec::new();
    if let Ok(file_lines) = read_lines(file) {
        for line in file_lines {
            if let Ok(ip) = line {
                lines.push(ip);
            }
        }
    }
    return lines;
}

fn read_lines<P>(filename: P) -> io::Result<Lines<BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(BufReader::new(file).lines())
}

pub fn parse_string_from(s: &str) -> io::Result<String> {
    let mut file = File::open(s)?;
    let mut s = String::new();
    file.read_to_string(&mut s)?;
    Ok(s)
}