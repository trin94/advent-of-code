export async function part1(): Promise<number> {
	const filename = "01.input.txt";
	const foo = Bun.file(filename);
	const text = await foo.text();

	let paper = 0;
	for (const dimension of text.split("\n")) {
		if (dimension) {
			paper += calculatePaper(dimension);
		}
	}
	return paper;
}

export function calculatePaper(dimensions: string): number {
	const [ll, ww, hh] = dimensions.split("x").map((part) => parseInt(part, 10));

	const l = ll as number;
	const w = ww as number;
	const h = hh as number;

	const surface = 2 * l * w + 2 * w * h + 2 * h * l;
	const smallestArea = Math.min(l * w, l * h, w * h);
	return surface + smallestArea;
}

export async function part2(): Promise<number> {
	const filename = "01.input.txt";
	const foo = Bun.file(filename);
	const text = await foo.text();

	let paper = 0;
	for (const dimension of text.split("\n")) {
		if (dimension) {
			paper += calculateRibbon(dimension);
		}
	}
	return paper;
}

export function calculateRibbon(dimensions: string): number {
	const [ll, ww, hh] = dimensions.split("x").map((part) => parseInt(part, 10));

	const l = ll as number;
	const w = ww as number;
	const h = hh as number;

	const sides = [l, w, h].sort((a, b) => a - b);

	const first = sides[0] as number;
	const second = sides[1] as number;

	const wrap = first + first + second + second;
	const bow = l * w * h;

	return wrap + bow;
}
