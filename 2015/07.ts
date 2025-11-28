async function readLines(filename: string) {
	const foo = Bun.file(`input/${filename}`);
	const input = await foo.text();
	return input.split("\n").filter((line) => !!line);
}

function emulate(rules: Rule[], state: Map<string, number>) {
	let ruleCopy = [...rules];
	while (ruleCopy.length !== 0) {
		for (const rule of ruleCopy) {
			if (!rule.finished()) {
				rule.applyIfPossible(state);
			}
		}
		ruleCopy = ruleCopy.filter((rule) => !rule.finished());
	}
}

function calculate(
	lines: string[],
	state: Map<string, number>,
): Map<string, number> {
	const rules = lines.map((line) => parseRule(line));
	emulate(rules, state);
	return state;
}

export async function calculateCircuit(
	filename: string,
): Promise<Map<string, number>> {
	const lines = await readLines(filename);
	const state: Map<string, number> = new Map();
	calculate(lines, state);
	return state;
}

export async function part1(filename: string) {
	const lines = await readLines(filename);
	const state: Map<string, number> = new Map();
	calculate(lines, state);
	return state.get("a");
}

export async function part2(filename: string) {
	const lines = await readLines(filename);
	const state: Map<string, number> = new Map();
	calculate(lines, state);

	const a = state.get("a");
	if (a === undefined) throw new Error("Cannot find value a");

	const modifiedLines = lines.map((line) => {
		if (line.endsWith("-> b")) {
			return `${a} -> b`;
		}
		return line;
	});
	const newState: Map<string, number> = new Map();
	calculate(modifiedLines, newState);

	return newState.get("a");
}

interface Rule {
	name: string;
	finished: () => boolean;
	applyIfPossible: (state: Map<string, number>) => void;
}

class AssignmentRule implements Rule {
	static PATTERN = /^([a-zA-Z0-9]+) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const result = match[2];
		this.name = `AssignmentRule: ${line}`;
		if (!left || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const value = Number.isNaN(Number(this.left))
			? state.get(this.left)
			: parseInt(this.left, 10);

		if (value !== undefined) {
			state.set(this.result, value & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

class NotRule implements Rule {
	static PATTERN = /^NOT (.*?) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const result = match[2];
		this.name = `NotRule: ${line}`;
		if (!left || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const left = state.get(this.left);
		if (left !== undefined) {
			state.set(this.result, ~left & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

class AndRule implements Rule {
	static PATTERN = /^(.*?) AND (.*?) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly right: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const right = match[2];
		const result = match[3];
		this.name = `AndRule: ${line}`;
		if (!left || !right || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.right = right;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const left = Number.isNaN(Number(this.left))
			? state.get(this.left)
			: parseInt(this.left, 10);

		const right = state.get(this.right);
		if (left !== undefined && right !== undefined) {
			state.set(this.result, left & right & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

class OrRule implements Rule {
	static PATTERN = /^(.*?) OR (.*?) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly right: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const right = match[2];
		const result = match[3];
		this.name = `OrRule: ${line}`;
		if (!left || !right || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.right = right;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const left = state.get(this.left);
		const right = state.get(this.right);
		if (left !== undefined && right !== undefined) {
			state.set(this.result, (left | right) & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

class LeftShiftRule implements Rule {
	static PATTERN = /^(.*?) LSHIFT (.*?) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly right: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const right = match[2];
		const result = match[3];
		this.name = `LeftShiftRule: ${line}`;
		if (!left || !right || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.right = right;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const left = state.get(this.left);
		if (left !== undefined) {
			state.set(this.result, (left << parseInt(this.right, 10)) & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

class RightShiftRule implements Rule {
	static PATTERN = /^(.*?) RSHIFT (.*?) -> (.*?)$/;

	readonly name: string;
	private readonly left: string;
	private readonly right: string;
	private readonly result: string;
	private applied = false;

	constructor(line: string, match: RegExpMatchArray) {
		const left = match[1];
		const right = match[2];
		const result = match[3];
		this.name = `RightShiftRule: ${line}`;
		if (!left || !right || !result) {
			throw new Error(`Cannot create ${this.name}`);
		}
		this.left = left;
		this.right = right;
		this.result = result;
	}

	applyIfPossible(state: Map<string, number>): void {
		const left = state.get(this.left);
		if (left !== undefined) {
			state.set(this.result, (left >> parseInt(this.right, 10)) & 0xffff);
			this.applied = true;
		}
	}

	finished(): boolean {
		return this.applied;
	}
}

function parseRule(line: string): Rule {
	const assignmentMatch = line.match(AssignmentRule.PATTERN);
	if (assignmentMatch !== null) {
		return new AssignmentRule(line, assignmentMatch);
	}

	const notMatch = line.match(NotRule.PATTERN);
	if (notMatch !== null) {
		return new NotRule(line, notMatch);
	}

	const andMatch = line.match(AndRule.PATTERN);
	if (andMatch !== null) {
		return new AndRule(line, andMatch);
	}

	const orMatch = line.match(OrRule.PATTERN);
	if (orMatch !== null) {
		return new OrRule(line, orMatch);
	}

	const leftShiftMatch = line.match(LeftShiftRule.PATTERN);
	if (leftShiftMatch !== null) {
		return new LeftShiftRule(line, leftShiftMatch);
	}

	const rightShiftMatch = line.match(RightShiftRule.PATTERN);
	if (rightShiftMatch !== null) {
		return new RightShiftRule(line, rightShiftMatch);
	}

	throw new Error(`Cannot parse rule from line: ${line}`);
}
