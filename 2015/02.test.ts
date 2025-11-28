import { describe, expect, test } from "bun:test";
import { calculatePaper, calculateRibbon, part1, part2 } from "./02.ts";

describe("part 1", () => {
	test("calculatePaper 2x3x4", () => {
		expect(calculatePaper("2x3x4")).toBe(58);
	});
	test("calculatePaper 1x1x10", () => {
		expect(calculatePaper("1x1x10")).toBe(43);
	});
	test("input.txt", async () => {
		expect(await part1()).toBe(1598415);
	});
});

describe("part 2", () => {
	test("calculateRibbon 2x3x4", () => {
		expect(calculateRibbon("2x3x4")).toBe(34);
	});
	test("calculateRibbon 1x1x10", () => {
		expect(calculateRibbon("1x1x10")).toBe(14);
	});
	test("input.txt", async () => {
		expect(await part2()).toBe(3812909);
	});
});
