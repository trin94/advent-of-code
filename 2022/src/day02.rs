use crate::files;

enum Shape {
    ROCK,
    PAPER,
    SCISSOR,
}

impl Shape {
    pub fn new(shape: String) -> Self {
        match shape.as_str() {
            "A" | "X" => Shape::ROCK,
            "B" | "Y" => Shape::PAPER,
            "C" | "Z" => Shape::SCISSOR,
            _ => panic!("Don't know shape '{}'", shape)
        }
    }

    pub fn points(&self) -> u32 {
        match self {
            Shape::ROCK => 1,
            Shape::PAPER => 2,
            Shape::SCISSOR => 3,
        }
    }

    pub fn against(&self, other: &Shape) -> u32 {
        match self { // if u played self and other played other
            Shape::ROCK => match other {
                Shape::ROCK => 3,
                Shape::PAPER => 0,
                Shape::SCISSOR => 6,
            },
            Shape::PAPER => match other {
                Shape::ROCK => 6,
                Shape::PAPER => 3,
                Shape::SCISSOR => 0,
            },
            Shape::SCISSOR => match other {
                Shape::ROCK => 0,
                Shape::PAPER => 6,
                Shape::SCISSOR => 3,
            },
        }
    }

    pub fn to_get(&self, result: Result) -> Shape {
        match self {
            Shape::ROCK => match result {
                Result::LOSE => Shape::SCISSOR,
                Result::DRAW => Shape::ROCK,
                Result::WIN => Shape::PAPER,
            },
            Shape::PAPER => match result {
                Result::LOSE => Shape::ROCK,
                Result::DRAW => Shape::PAPER,
                Result::WIN => Shape::SCISSOR,
            },
            Shape::SCISSOR => match result {
                Result::LOSE => Shape::PAPER,
                Result::DRAW => Shape::SCISSOR,
                Result::WIN => Shape::ROCK,
            },
        }
    }
}

enum Result {
    LOSE,
    DRAW,
    WIN,
}

impl Result {
    pub fn new(result: String) -> Self {
        match result.as_str() {
            "X" => Result::LOSE,
            "Y" => Result::DRAW,
            "Z" => Result::WIN,
            _ => panic!("Don't know result '{}'", result)
        }
    }
}

struct Match {
    played: Shape,
    response: Shape,
}

impl Match {
    pub fn new(played: Shape, response: Shape) -> Self {
        return Match { played, response };
    }

    pub fn new_cheated(played: Shape, result: Result) -> Self {
        let i_play = played.to_get(result);
        return Match::new(played, i_play);
    }

    pub fn evaluate(&self) -> u32 {
        let base = self.response.points();
        let outcome = self.response.against(&self.played);
        return base + outcome;
    }
}

pub fn solve() {
    let file = "resources/day02.txt";
    let lines = files::parse_lines_from(file);

    let according_to_guess = |line: &String| {
        let separated: Vec<&str> = line.split(" ").collect();
        let played = separated.first().unwrap().to_string();
        let played = Shape::new(played);
        let response = separated.last().unwrap().to_string();
        let response = Shape::new(response);
        Match::new(played, response)
    };

    let sum: u32 = lines.iter()
        .map(according_to_guess)
        .map(|m| m.evaluate())
        .sum();

    println!("Puzzle 1 sum: {}", sum);

    let according_to_meaning = |line: &String| {
        let separated: Vec<&str> = line.split(" ").collect();
        let played = separated.first().unwrap().to_string();
        let played = Shape::new(played);
        let result = separated.last().unwrap().to_string();
        let result = Result::new(result);
        Match::new_cheated(played, result)
    };

    let sum: u32 = lines.iter()
        .map(according_to_meaning)
        .map(|m| m.evaluate())
        .sum();

    println!("Puzzle 2 sum: {}", sum);
}