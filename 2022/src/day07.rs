use std::collections::{HashMap, VecDeque};

use crate::files;

const DISK_SPACE: u32 = 70000000;
const REQUIRED_SPACE: u32 = 30000000;

#[derive(Clone)]
struct File {
    size: usize,
}

pub fn solve() {
    let file = "resources/day07.txt";
    let lines = files::parse_string_from(file)
        .expect(&*(" Could not read file ".to_owned() + file));

    let lines = lines.trim();

    let commands = lines.split("$ ")
        .into_iter()
        .map(|command| command.trim())
        .map(|command| command.to_owned())
        .skip_while(|command| command.is_empty() || command == "cd /")
        .collect::<Vec<String>>();

    let mut map: HashMap<String, Vec<File>> = HashMap::new();
    let mut path = VecDeque::new();
    path.push_back(".".to_string());

    for command in commands {
        if command.starts_with("cd ..") {
            path.pop_back();
        } else if command.starts_with("cd ") {
            let option = command.split_whitespace().last().unwrap().to_string();
            path.push_back(option);
        } else if command.starts_with("ls") {
            let files: _ = command
                .split("\n")
                .skip(1)
                .filter(|line| !line.starts_with("dir"))
                .map(|line| {
                    let split: Vec<_> = line.split(' ').collect();
                    let size = split.first().unwrap().parse::<usize>().unwrap();
                    File { size }
                }).collect::<Vec<File>>();

            let mut copied = path.clone();
            while copied.len() > 0 {
                let directory = Vec::from_iter(&copied)
                    .iter()
                    .map(|s| s.to_string())
                    .collect::<Vec<String>>()
                    .join("/");
                if map.contains_key(&directory) {
                    let x = map.get_mut(&directory).unwrap();
                    x.extend(files.clone())
                } else {
                    let mut n = Vec::new();
                    n.extend(files.clone());
                    map.insert(directory, n);
                }
                copied.pop_back();
            }
        }
    }

    let sum: usize = map.iter()
        .map(|e| {
            let x2: usize = e.1.iter()
                .map(|file| file.size)
                .sum();
            x2
        })
        .filter(|s| s <= &(100_000 as usize))
        .sum();

    println!("-> Sum {}", sum);

    let used_space: usize = map.get(".")
        .unwrap().iter()
        .map(|f| f.size)
        .sum::<usize>();

    println!("Used space {}", used_space);

    let free_space: u32 = DISK_SPACE - used_space as u32;
    println!("Free space {}", free_space);

    let required_space: u32 = REQUIRED_SPACE - free_space as u32;
    println!("Required space {}", required_space);

    let size = map.iter()
        .map(|e| {
            let x2: usize = e.1.iter()
                .map(|file| file.size)
                .sum();
            x2
        })
        .filter(|s| s >= &(required_space as usize))
        .min().unwrap();

    println!("-> Size of file to be deleted {}", size)
}
