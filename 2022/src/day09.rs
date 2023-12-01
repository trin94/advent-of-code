use std::borrow::{Borrow, BorrowMut};
use std::cmp;
use std::collections::HashSet;

use crate::files;

#[derive(Copy, Clone, Hash, Eq, PartialEq, Debug)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Clone)]
struct Movement {
    dx: i32,
    dy: i32,
    times: u32,
}

impl Movement {
    pub fn new(dx: i32, dy: i32, times: &str) -> Self {
        Movement { dx, dy, times: times.trim().parse::<u32>().unwrap() }
    }

    pub fn flatten(&self) -> Vec<Movement> {
        let movement = Movement { dx: self.dx, dy: self.dy, times: 1 };
        vec![movement; self.times as usize]
    }
}

fn parse_movements(lines: Vec<String>) -> Vec<Movement> {
    lines.iter()
        .map(|movement| match movement.split_at(1) {
            ("L", n) => Movement::new(-1, 0, n),
            ("R", n) => Movement::new(1, 0, n),
            ("U", n) => Movement::new(0, 1, n),
            ("D", n) => Movement::new(0, -1, n),
            (_, _) => panic!("Not reachable"),
        })
        .flat_map(|movement| movement.flatten())
        .collect::<Vec<Movement>>()
}

struct Rope {
    size: u32,
    elements: Vec<Point>,
    visited: HashSet<Point>,
}

const MIN: i32 = -1;
const MAX: i32 = 1;

impl Rope {
    fn new(size: u32) -> Self {
        let start = Point { x: 0, y: 0 };
        let mut visited = HashSet::new();
        visited.insert(start);
        Rope { size, visited, elements: vec![start; size as usize] }
    }

    pub fn move_head(&mut self, movement: &Movement) {
        let mut point: &mut Point = self.elements.get_mut(0).unwrap();
        point.x += movement.dx;
        point.y += movement.dy;
    }

    pub fn follow_head(&mut self) {
        let i = self.size as usize;
        for index in 1..i {
            let head = self.elements.get(index - 1).unwrap();
            let tail = self.elements.get(index).unwrap();
            if let Some(delta) = self.calculate_position(head, tail) {
                let tail = self.elements.get_mut(index).unwrap();
                tail.x = delta.x;
                tail.y = delta.y;
            } else {
                break;
            }
        }
    }

    fn calculate_position(&self, head: &Point, tail: &Point) -> Option<Point> {
        let dx = tail.x - head.x;
        let dy = tail.y - head.y;

        if (dx == 2 || dx == -2) && (dy == 2 || dy == -2) {
            Some(Point { x: head.x + dx.clamp(MIN, MAX), y: head.y + dy.clamp(MIN, MAX) })
        } else if dx == 2 || dx == -2 {
            Some(Point { x: head.x + dx.clamp(MIN, MAX), y: head.y })
        } else if dy == 2 || dy == -2 {
            Some(Point { x: head.x, y: head.y + dy.clamp(MIN, MAX) })
        } else {
            None
        }
    }

    pub fn store_tail_position(&mut self) {
        let tail = self.elements[self.size as usize - 1];
        self.visited.insert(tail);
    }
}


pub fn solve() {
    let file = "resources/day09.txt";
    let lines = files::parse_lines_from(file);
    let movements = parse_movements(lines);

    let mut rope = Rope::new(2);
    calculate(&movements, &mut rope);

    let mut rope = Rope::new(10);
    calculate(&movements, &mut rope);
}


fn calculate(movements: &Vec<Movement>, rope: &mut Rope) {
    for movement in movements {
        rope.move_head(movement);
        rope.follow_head();
        rope.store_tail_position();
    }
    println!("{}", rope.visited.len());
}
