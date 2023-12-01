use std::cmp;

use crate::files;

const RADIX: u32 = 10;

struct Grid {
    size: usize,
    elements: Vec<usize>,
}

impl Grid {
    pub fn new(string: String) -> Self {
        let size = string.find("\n").unwrap() as usize;
        let elements = string.chars()
            .filter(|c| c != &'\n')
            .map(|c| c.to_digit(RADIX).unwrap() as usize)
            .collect::<Vec<usize>>();
        Grid { size, elements }
    }

    pub fn can_view_border(&self, x: usize, y: usize) -> bool {
        self.can_view_left_border(x, y) ||
            self.can_view_right_border(x, y) ||
            self.can_view_upper_border(x, y) ||
            self.can_view_bottom_border(x, y)
    }

    fn item(&self, x: usize, y: usize) -> &usize {
        &self.elements[x + y * self.size]
    }

    fn can_view_left_border(&self, x: usize, y: usize) -> bool {
        let item = self.item(x, y);
        self.elements_left_of(x, y)
            .iter()
            .all(|e| e < &item)
    }

    fn can_view_right_border(&self, x: usize, y: usize) -> bool {
        let item = self.item(x, y);
        self.elements_right_of(x, y)
            .iter()
            .all(|e| e < &item)
    }

    fn can_view_upper_border(&self, x: usize, y: usize) -> bool {
        let item = self.item(x, y);
        self.elements_above_of(x, y)
            .iter()
            .all(|e| e < &item)
    }

    fn can_view_bottom_border(&self, x: usize, y: usize) -> bool {
        let item = self.item(x, y);
        self.elements_below_of(x, y)
            .iter()
            .all(|e| e < &item)
    }

    fn elements_left_of(&self, x: usize, y: usize) -> Vec<&usize> {
        let start = y * self.size;
        let end = x + y * self.size;
        self.elements[start..end]
            .iter()
            .collect::<Vec<&usize>>()
    }

    fn elements_right_of(&self, x: usize, y: usize) -> Vec<&usize> {
        let start = y * self.size + x + 1;
        let end = (y + 1) * self.size;
        self.elements[start..end]
            .iter()
            .collect::<Vec<&usize>>()
    }

    fn elements_above_of(&self, x: usize, y: usize) -> Vec<&usize> {
        let start = x;
        let end = y * self.size + x;
        self.elements[start..end]
            .iter()
            .step_by(self.size)
            .collect::<Vec<&usize>>()
    }

    fn elements_below_of(&self, x: usize, y: usize) -> Vec<&usize> {
        let end = self.size * self.size;
        let start = (y + 1) * self.size + x;
        let start = cmp::min(start, end);
        self.elements[start..end]
            .iter()
            .step_by(self.size)
            .collect::<Vec<&usize>>()
    }

    pub fn scenic_score_of(&self, x: usize, y: usize) -> usize {
        let left = self.viewing_distance_left_of(x, y);
        let right = self.viewing_distance_right_of(x, y);
        let above = self.viewing_distance_above_of(x, y);
        let below = self.viewing_distance_below_of(x, y);
        left * right * above * below
    }

    fn viewing_distance_left_of(&self, x: usize, y: usize) -> usize {
        let item = self.item(x, y);
        let elements = self.elements_left_of(x, y);
        let smaller = elements
            .iter()
            .rev()
            .take_while(|e| e < &&item)
            .count();
        cmp::min(smaller + 1, elements.len())
    }

    fn viewing_distance_right_of(&self, x: usize, y: usize) -> usize {
        let item = self.item(x, y);
        let elements = self.elements_right_of(x, y);
        let smaller = elements
            .iter()
            .take_while(|e| e < &&item)
            .count();
        cmp::min(smaller + 1, elements.len())
    }

    fn viewing_distance_above_of(&self, x: usize, y: usize) -> usize {
        let item = self.item(x, y);
        let elements = self.elements_above_of(x, y);
        let smaller = elements
            .iter()
            .rev()
            .take_while(|e| e < &&item)
            .count();
        cmp::min(smaller + 1, elements.len())
    }

    fn viewing_distance_below_of(&self, x: usize, y: usize) -> usize {
        let item = self.item(x, y);
        let elements = self.elements_below_of(x, y);
        let smaller = elements
            .iter()
            .take_while(|e| e < &&item)
            .count();
        cmp::min(smaller + 1, elements.len())
    }
}


pub fn solve() {
    let file = "resources/day08.txt";
    let lines = files::parse_string_from(file).unwrap();
    let grid = Grid::new(lines);

    let mut sum = 0;
    for x in 0..grid.size {
        for y in 0..grid.size {
            if grid.can_view_border(x, y) {
                sum += 1;
            }
        }
    }

    println!("Visible outside: {}", sum);

    let mut scenic = 0;
    for x in 0..grid.size {
        for y in 0..grid.size {
            let new_scenic = grid.scenic_score_of(x, y);
            if new_scenic > scenic {
                scenic = new_scenic;
            }
        }
    }

    println!("Scenic score: {}", scenic);
}
