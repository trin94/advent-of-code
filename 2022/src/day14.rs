use std::cmp;

use itertools::Itertools;

use crate::files;

#[derive(Copy, Clone, Debug)]
enum Element {
    ROCK,
    SAND,
    AIR,
}


#[derive(Copy, Clone, Debug)]
struct Coordinate {
    x: usize,
    y: usize,
}

impl Coordinate {
    pub fn new(x: usize, y: usize) -> Self {
        Self { x, y }
    }

    pub fn from_string(string: &str) -> Self {
        let mut split = string.split(",");
        let x = split.next().unwrap().parse::<usize>().unwrap();
        let y = split.next().unwrap().parse::<usize>().unwrap();
        Self { x, y }
    }

    pub fn between(c1: &Coordinate, c2: &Coordinate) -> Vec<Coordinate> {
        let mut coordinates: Vec<Coordinate> = Vec::new();
        if c1.is_horizontal_to(c2) {
            let min = cmp::min(c1.x, c2.x);
            let max = cmp::max(c1.x, c2.x);
            for x in min..=max {
                coordinates.push(Coordinate { x, y: c1.y })
            }
        } else {
            let min = cmp::min(c1.y, c2.y);
            let max = cmp::max(c1.y, c2.y);
            for y in min..=max {
                coordinates.push(Coordinate { x: c1.x, y })
            }
        }
        coordinates
    }

    fn is_horizontal_to(&self, other_coordinate: &Coordinate) -> bool {
        self.y == other_coordinate.y
    }
}


struct Cave {
    width: usize,
    height: usize,
    elements: Vec<Element>,
    sand_consumed: usize,
    deepest_rock: usize,
}

impl Cave {
    pub fn new(width: usize, height: usize) -> Self {
        let mut elements = Vec::new();
        elements.resize(width * height, Element::AIR);
        Self { width, height, elements, sand_consumed: 0, deepest_rock: 0 }
    }

    fn put_element_at(&mut self, x: usize, y: usize, element: Element) {
        match self.element_get(x, y) {
            Element::ROCK => match element {
                Element::ROCK => {}
                _ => panic!("Cannot put element where something already was put before"),
            },
            Element::SAND => panic!("Cannot put element where sand already was put before"),
            Element::AIR => self.element_set(x, y, element),
        }
    }

    fn element_get(&self, x: usize, y: usize) -> &Element {
        let index: usize = x + y * self.width;
        self.elements.get(index).unwrap()
    }

    fn element_set(&mut self, x: usize, y: usize, element: Element) {
        let index: usize = x + y * self.width;
        self.elements[index] = element;
    }

    pub fn put_rocks_between(&mut self, c1: Coordinate, c2: Coordinate) {
        let rock = Element::ROCK;
        for coordinate in Coordinate::between(&c1, &c2).iter() {
            self.put_element_at(coordinate.x, coordinate.y, rock);
            self.deepest_rock = cmp::max(self.deepest_rock, coordinate.y);
        }
    }

    pub fn drop_sand_at(&mut self, sand_entry: Coordinate) -> bool {
        let mut x = sand_entry.x;
        let mut y = sand_entry.y;
        loop {
            if self.dropped_out(y) {
                return false;
            }
            if self.element_below_is_air(x, y) {
                y += 1;
                continue;
            }
            if self.element_left_below_is_air(x, y) {
                y += 1;
                x -= 1;
                continue;
            }
            if self.element_right_below_is_air(x, y) {
                y += 1;
                x += 1;
                continue;
            }
            self.put_sand_at(x, y);
            return if x == sand_entry.x && y == sand_entry.y {
                false
            } else {
                true
            }
        }
    }

    fn dropped_out(&self, y: usize) -> bool {
        y >= self.height - 1
    }

    fn element_below_is_air(&self, x: usize, y: usize) -> bool {
        matches!(self.element_get(x, y + 1), Element::AIR)
    }

    fn element_left_below_is_air(&self, x: usize, y: usize) -> bool {
        matches!(self.element_get(x - 1, y + 1), Element::AIR)
    }

    fn element_right_below_is_air(&self, x: usize, y: usize) -> bool {
        matches!(self.element_get(x + 1, y + 1), Element::AIR)
    }

    fn put_sand_at(&mut self, x: usize, y: usize) {
        self.put_element_at(x, y, Element::SAND);
        self.sand_consumed += 1;
    }
}


fn construct_cave(width: usize, height: usize) -> Cave {
    let mut cave = Cave::new(width, height);
    let file = "resources/day14.txt";
    let lines = files::parse_lines_from(file);
    for line in lines {
        let coordinates = line.split(" -> ")
            .map(|s| Coordinate::from_string(s))
            .into_iter();
        for (c1, c2) in coordinates.tuple_windows() {
            cave.put_rocks_between(c1, c2)
        }
    }
    cave
}


pub fn solve() {
    solve_part1();
    solve_part2();
}

fn solve_part1() {
    let sand_entry = Coordinate::new(500, 0);
    let mut cave = construct_cave(1000, 1000);

    while cave.drop_sand_at(sand_entry) {};
    println!("Part 1: Consumed {} Sand", cave.sand_consumed)
}

fn solve_part2() {
    let sand_entry = Coordinate::new(500, 0);

    let width = 5000;
    let height = 5000;
    let mut cave = construct_cave(width, height);

    let deepest_rock = cave.deepest_rock;
    let c1 = Coordinate::new(0, deepest_rock + 2);
    let c2 = Coordinate::new(width, deepest_rock + 2);
    cave.put_rocks_between(c1, c2);

    while cave.drop_sand_at(sand_entry) {};
    println!("Part 2: Consumed {} Sand", cave.sand_consumed)
}
