use std::collections::{HashSet, VecDeque};

use crate::files;

struct MarkerDetector {
    elements: VecDeque<char>,
    marker_size: u32,
    position: u32,
}

impl MarkerDetector {
    pub fn new(marker_size: u32) -> Self {
        MarkerDetector { marker_size, position: 0, elements: VecDeque::new() }
    }

    pub fn parse(&mut self, character: char) {
        self.elements.push_back(character);
        if self.elements.len() > self.marker_size as usize {
            self.elements.pop_front();
        }
        self.position += 1;
    }

    pub fn marker_found(&self) -> bool {
        let mut uniques = HashSet::new();
        let mut elements = self.elements.clone();
        elements.retain(|e| uniques.insert(*e));

        elements.len() == self.marker_size as usize && elements.len() == uniques.len()
    }
}


pub fn solve() {
    let file = "resources/day06.txt";
    let lines = files::parse_lines_from(file);

    let mut packet_detector = MarkerDetector::new(4);

    for c in lines.get(0).unwrap().chars() {
        packet_detector.parse(c);
        if packet_detector.marker_found() {
            println!("start-of-packet marker: {}", packet_detector.position);
            break;
        }
    }

    let mut message_detector = MarkerDetector::new(14);
    for c in lines.get(0).unwrap().chars() {
        message_detector.parse(c);
        if message_detector.marker_found() {
            println!("start-of-message marker: {}", message_detector.position);
            break;
        }
    }
}
