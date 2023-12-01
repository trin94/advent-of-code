use std::cmp;
use std::collections::HashSet;
use std::ops::RangeInclusive;

use itertools::Itertools;

use crate::files;

#[derive(Copy, Clone, Hash, Eq, PartialEq, Debug)]
struct Point {
    x: i32,
    y: i32,
}


#[derive(Debug)]
struct Deployment {
    sensor_location: Point,
    closest_beacon_location: Point,
}

impl Deployment {
    pub fn from_str(str: &String) -> Self {
        let mut locations = str.splitn(2, ":");
        let sensor_location = Deployment::extract_point(locations.next().unwrap());
        let closest_beacon_location = Deployment::extract_point(locations.next().unwrap());
        Self { sensor_location, closest_beacon_location }
    }

    fn extract_point(string: &str) -> Point {
        let splits: Vec<_> = string.splitn(4, |c| c == '=' || c == ',').collect();
        let x = splits[1].parse::<i32>().unwrap();
        let y = splits[3].parse::<i32>().unwrap();
        Point { x, y }
    }

    pub fn covers(&self, line: i32) -> bool {
        let sensor = self.sensor_location;
        let distance = self.sensor_beacon_distance();
        line >= sensor.y - distance && line <= sensor.y + distance
    }

    fn sensor_beacon_distance(&self) -> i32 {
        let p1 = self.sensor_location;
        let p2 = self.closest_beacon_location;
        (p1.x - p2.x).abs() + (p1.y - p2.y).abs()
    }

    pub fn covered_points_in(&self, line: &i32, min: &i32, max: &i32) -> RangeInclusive<i32> {
        let sensor = self.sensor_location;
        let sensor_beacon_distance = self.sensor_beacon_distance();
        let delta_vertical = (sensor.y - line).abs();
        let delta_horizontal = sensor_beacon_distance - delta_vertical;
        let start = *cmp::max(min, &(sensor.x - delta_horizontal));
        let end = *cmp::min(max, &(sensor.x + delta_horizontal));
        start..=end
    }
}


pub fn solve() {
    solve_part1();
    solve_part2();
}


fn solve_part1() {
    let file = "resources/day15.txt";
    let lines = files::parse_lines_from(file);

    let requested_line = 2_000_000;
    let min = i32::MIN;
    let max = i32::MAX;
    let deployments = lines
        .iter()
        .map(|line| Deployment::from_str(line))
        .collect::<Vec<Deployment>>();

    let mut covered_xs_in_requested_line = deployments.iter()
        .filter(|deployment| deployment.covers(requested_line))
        .flat_map(|deployment| deployment.covered_points_in(&requested_line, &min, &max).collect::<Vec<i32>>())
        .collect::<HashSet<i32>>();

    for deployment in deployments {
        if deployment.sensor_location.y == requested_line {
            covered_xs_in_requested_line.remove(&deployment.sensor_location.x);
        }
        if deployment.closest_beacon_location.y == requested_line {
            covered_xs_in_requested_line.remove(&deployment.closest_beacon_location.x);
        }
    }

    println!("Part 1: {}", covered_xs_in_requested_line.len())
}


fn solve_part2() {
    let file = "resources/day15.txt";
    let lines = files::parse_lines_from(file);

    let min = 0;
    let max = 4_000_000;

    let deployments = lines
        .iter()
        .map(|line| Deployment::from_str(line))
        .collect::<Vec<Deployment>>();

    let mut yy: i64 = 0;
    let mut xx: i64 = 0;

    'scan: for y in min..=max {
        let mut ranges: HashSet<RangeInclusive<i32>> = HashSet::new();
        for deployment in &deployments {
            if deployment.covers(y) {
                let range = deployment.covered_points_in(&y, &min, &max);
                ranges.insert(range);
            }
        }

        // Remove all ranges that are part of another range
        let ranges_copy = ranges.clone();
        for r1 in &ranges_copy {
            for r2 in &ranges_copy {
                if r1 != r2 && r1.start() >= r2.start() && r1.end() <= r2.end() {
                    ranges.remove(r1);
                }
            }
        }

        // sort
        let ranges: Vec<_> = ranges.iter()
            .sorted_by(|a, b| Ord::cmp(&a.start(), &b.start()))
            .collect();

        // Find the ranges that are not 'connected'
        for (r1, r2) in ranges.iter().tuple_windows() {
            if r2.start() - r1.end() > 1 {
                yy = y as i64;
                xx = (r1.end() + 1) as i64;
                break 'scan;
            }
        }
    }

    let frequency = xx * 4_000_000 + yy;
    println!("Part 2: {} at ({xx}|{yy})", frequency)
}
