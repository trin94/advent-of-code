use crate::files;

struct Range {
    lower: u32,
    upper: u32,
}

impl Range {
    pub fn new(range: &str) -> Self {
        let mut split = range.splitn(2, "-");
        let lower = split.next().unwrap().parse::<u32>().unwrap();
        let upper = split.next().unwrap().parse::<u32>().unwrap();
        Range { lower, upper }
    }

    pub fn fully_includes(&self, other: &Range) -> bool {
        other.lower >= self.lower && other.upper <= self.upper
    }

    pub fn overlap(&self, other: &Range) -> bool {
        self.lower >= other.lower && self.lower <= other.upper
    }
}

pub fn solve() {
    let file = "resources/day04.txt";
    let lines = files::parse_lines_from(file);

    let ranges: Vec<(Range, Range)> = lines
        .iter()
        .map(|line| {
            let mut pairs = line.splitn(2, ",");
            let first = pairs.next().unwrap();
            let last = pairs.next().unwrap();
            let first = Range::new(first);
            let last = Range::new(last);
            (first, last)
        }).collect::<Vec<(Range, Range)>>();

    let contain = ranges
        .iter()
        .filter(|(first, last)| {
            first.fully_includes(last) || last.fully_includes(first)
        })
        .count();

    println!("Fully contain the other: {}", contain);

    let overlap = ranges.iter()
        .filter(|(first, last)| {
            first.overlap(last) || last.overlap(first)
        })
        .count();

    println!("Overlap each other: {}", overlap)
}