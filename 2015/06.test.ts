import { describe, expect, test } from "bun:test";
import { parseInstruction, part1, part2 } from "./06.ts";

describe("parse Instruction", async () => {
	test("turn on 0,0 through 999,999", async () => {
		const parsed = parseInstruction("turn on 0,0 through 999,999");
		expect(JSON.stringify(parsed)).toBe(
			JSON.stringify({
				command: "on",
				fromX: 0,
				fromY: 0,
				toX: 999,
				toY: 999,
			}),
		);
	});
	test("toggle 0,0 through 999,0", async () => {
		const parsed = parseInstruction("toggle 0,0 through 999,0");
		expect(JSON.stringify(parsed)).toBe(
			JSON.stringify({
				command: "toggle",
				fromX: 0,
				fromY: 0,
				toX: 999,
				toY: 0,
			}),
		);
	});
	test("turn off 499,499 through 500,500", async () => {
		const parsed = parseInstruction("turn off 499,499 through 500,500");
		expect(JSON.stringify(parsed)).toBe(
			JSON.stringify({
				command: "off",
				fromX: 499,
				fromY: 499,
				toX: 500,
				toY: 500,
			}),
		);
	});
});

test("part 1", async () => {
	expect(await part1()).toBe(569999);
});

test("part 2", async () => {
	expect(await part2()).toBe(17836115);
});
