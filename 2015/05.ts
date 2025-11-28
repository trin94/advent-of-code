export function isNiceString(input: string): boolean {
	const inputLength = input.length;

	let vowelCount = 0;
	let doubleLetterFound = false;
	let previousLetter = "";

	for (let i = 0; i < input.length; i++) {
		const letter = input[i] as string;

		switch (letter) {
			case "a":
			case "e":
			case "i":
			case "o":
			case "u":
				vowelCount++;
				break;
		}

		if (!doubleLetterFound) {
			if (previousLetter === letter) {
				doubleLetterFound = true;
			} else {
				previousLetter = letter;
			}
		}

		if (i < inputLength) {
			const nextLetter = input[i + 1] as string;
			if (letter === "a" && nextLetter === "b") {
				return false;
			}
			if (letter === "c" && nextLetter === "d") {
				return false;
			}
			if (letter === "p" && nextLetter === "q") {
				return false;
			}
			if (letter === "x" && nextLetter === "y") {
				return false;
			}
		}
	}

	return vowelCount >= 3 && doubleLetterFound;
}

export async function part1(): Promise<number> {
	const filename = "input/05.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();
	return input
		.trim()
		.split(/\n/)
		.filter((line) => isNiceString(line)).length;
}

export function isCorrectNiceString(input: string): boolean {
	const inputLength = input.length;

	let twoLettersTwice = false;
	let oneLetterRepeats = false;

	for (let i = 0; i < input.length; i++) {
		const letter = input[i] as string;

		if (!twoLettersTwice && i < inputLength - 3) {
			const nextLetter = input[i + 1] as string;
			twoLettersTwice =
				input.substring(i + 2).indexOf(`${letter}${nextLetter}`) !== -1;
		}

		if (!oneLetterRepeats && i < inputLength - 2) {
			const possibleMirror = input[i + 2] as string;
			oneLetterRepeats = letter === possibleMirror;
		}
	}

	return twoLettersTwice && oneLetterRepeats;
}

export async function part2(): Promise<number> {
	const filename = "input/05.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();
	return input
		.trim()
		.split(/\n/)
		.filter((line) => isCorrectNiceString(line)).length;
}
