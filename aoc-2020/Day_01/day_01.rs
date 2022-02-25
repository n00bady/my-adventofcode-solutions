use std::io::{BufRead, BufReader};
use std::fs::File;

fn main() {
    let mut f = BufReader::new(File::open("Day_01_Input.txt").expect("Failed to open file!"));

    for line in f.lines() {
        for c in line.expect("Failed to read lines!").chars() {
            if (line.parse::<i32>().unwrap() + c.parse::<i32>().unwrap()) == 2020 {
                println!("FOUND!!!");
                println!("{} + {} = 2020 ", line, c);
                println!("{} * {} = {}", line, c, line*c);
            }
        }
    }
}
