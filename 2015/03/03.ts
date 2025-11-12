interface Identifiable {
	identifier(): string;
}

class Point implements Identifiable {
	readonly x: number;
	readonly y: number;

	constructor(x: number, y: number) {
		this.x = x;
		this.y = y;
	}

	identifier(): string {
		return `(${this.x}|${this.y})`;
	}
}

class JavaScriptSucksSet<E extends Identifiable> {
	private _map = new Map<string, E>();

	add(element: E): this {
		this._map.set(element.identifier(), element);
		return this;
	}

	size(): number {
		return this._map.size;
	}
}

export async function part1(): Promise<number> {
	const filename = "input.txt";
	const foo = Bun.file(filename);
	const movements = await foo.text();
	return countDeliveries(movements);
}

export async function part2(): Promise<number> {
	const filename = "input.txt";
	const foo = Bun.file(filename);
	const movements = await foo.text();
	return countDeliveriesWithSupport(movements);
}

export function countDeliveries(arr: string): number {
	const visited = new JavaScriptSucksSet<Point>();

	let current = new Point(0, 0);
	visited.add(current);

	for (const movement of arr) {
		switch (movement) {
			case ">":
				current = new Point(current.x + 1, current.y);
				break;
			case "<":
				current = new Point(current.x - 1, current.y);
				break;
			case "^":
				current = new Point(current.x, current.y + 1);
				break;
			case "v":
				current = new Point(current.x, current.y - 1);
				break;
		}
		visited.add(current);
	}

	return visited.size();
}

export function countDeliveriesWithSupport(arr: string): number {
	const visited = new JavaScriptSucksSet<Point>();

	let santa = new Point(0, 0);
	let robot = new Point(0, 0);
	visited.add(santa);
	visited.add(robot);

	let index = 0;
	for (const movement of arr) {
		const isSanta = index % 2 === 0;
		let current = isSanta ? santa : robot;

		switch (movement) {
			case ">":
				current = new Point(current.x + 1, current.y);
				break;
			case "<":
				current = new Point(current.x - 1, current.y);
				break;
			case "^":
				current = new Point(current.x, current.y + 1);
				break;
			case "v":
				current = new Point(current.x, current.y - 1);
				break;
		}
		visited.add(current);

		if (isSanta) {
			santa = current;
		} else {
			robot = current;
		}

		index++;
	}

	return visited.size();
}
