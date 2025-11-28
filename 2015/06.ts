interface Light {
	turnOn(): void;

	turnOff(): void;

	toggle(): void;
}

export class Grid {
	readonly _backend: Light[];

	constructor(create: () => Light, size = 1000 * 1000) {
		this._backend = new Array<Light>(size);
		for (let i = 0; i < size; i++) {
			this._backend[i] = create();
		}
	}

	turnOn(fromX: number, fromY: number, toX: number, toY: number) {
		for (let x = fromX; x <= toX; x++) {
			for (let y = fromY; y <= toY; y++) {
				this.elementAt(x, y).turnOn();
			}
		}
	}

	turnOff(fromX: number, fromY: number, toX: number, toY: number) {
		for (let x = fromX; x <= toX; x++) {
			for (let y = fromY; y <= toY; y++) {
				this.elementAt(x, y).turnOff();
			}
		}
	}

	toggle(fromX: number, fromY: number, toX: number, toY: number) {
		for (let x = fromX; x <= toX; x++) {
			for (let y = fromY; y <= toY; y++) {
				this.elementAt(x, y).toggle();
			}
		}
	}

	elementAt(x: number, y: number): Light {
		const position = 1000 * x + y;
		const element = this._backend[position];
		if (!element) {
			throw new Error(`Cannot get element at: (${x}|${y})`);
		}
		return element;
	}

	forEach(
		callbackfn: (value: Light, index: number, array: Light[]) => void,
	): void {
		this._backend.forEach(callbackfn);
	}
}

interface Instruction {
	command: "on" | "off" | "toggle";
	fromX: number;
	fromY: number;
	toX: number;
	toY: number;
}

const commandPattern =
	/^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)$/;

export function parseInstruction(instruction: string): Instruction {
	commandPattern.lastIndex = 0;

	const match = instruction.trim().match(commandPattern);
	if (match) {
		const command =
			match[1] === "turn on"
				? "on"
				: match[1] === "turn off"
					? "off"
					: "toggle";
		return {
			command,
			fromX: parseInt(match[2] as string, 10),
			fromY: parseInt(match[3] as string, 10),
			toX: parseInt(match[4] as string, 10),
			toY: parseInt(match[5] as string, 10),
		};
	}
	throw new Error(`Invalid instruction: ${instruction}`);
}

function followInstructions(grid: Grid, instructions: string[]) {
	for (const instruction of instructions) {
		const instruct = parseInstruction(instruction);
		const fromX = instruct.fromX;
		const fromY = instruct.fromY;
		const toX = instruct.toX;
		const toY = instruct.toY;

		if (instruct.command === "on") {
			grid.turnOn(fromX, fromY, toX, toY);
		} else if (instruct.command === "off") {
			grid.turnOff(fromX, fromY, toX, toY);
		} else {
			grid.toggle(fromX, fromY, toX, toY);
		}
	}
}

export async function part1(): Promise<number> {
	const filename = "input/06.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();

	class LightImpl implements Light {
		value: boolean = false;

		toggle(): void {
			this.value = !this.value;
		}

		turnOff(): void {
			this.value = false;
		}

		turnOn(): void {
			this.value = true;
		}
	}

	const grid = new Grid(() => new LightImpl());
	const instructions = input.split("\n").filter((line) => !!line);

	followInstructions(grid, instructions);

	let sum = 0;

	grid.forEach((light) => {
		const element = light as LightImpl;
		sum += element.value ? 1 : 0;
	});

	return sum;
}

export async function part2(): Promise<number> {
	const filename = "input/06.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();

	class LightImpl implements Light {
		brightness: number = 0;

		toggle(): void {
			this.brightness += 2;
		}

		turnOff(): void {
			this.brightness = Math.max(0, this.brightness - 1);
		}

		turnOn(): void {
			this.brightness += 1;
		}
	}

	const grid = new Grid(() => new LightImpl());
	const instructions = input.split("\n").filter((line) => !!line);

	followInstructions(grid, instructions);

	let sum = 0;

	grid.forEach((light) => {
		const element = light as LightImpl;
		sum += element.brightness;
	});

	return sum;
}
