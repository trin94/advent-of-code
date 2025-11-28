import { describe, expect, test } from "bun:test";
import { isCorrectNiceString, isNiceString, part1, part2 } from "./05.ts";

describe("isNiceString", () => {
	test("ugknbfddgicrmopn", () =>
		expect(isNiceString("ugknbfddgicrmopn")).toBe(true));
	test("aaa", () => expect(isNiceString("aaa")).toBe(true));
	test("jchzalrnumimnmhp", () =>
		expect(isNiceString("jchzalrnumimnmhp")).toBe(false));
	test("haegwjzuvuyypxyu", () =>
		expect(isNiceString("haegwjzuvuyypxyu")).toBe(false));
	test("dvszwmarrgswjxmb", () =>
		expect(isNiceString("dvszwmarrgswjxmb")).toBe(false));
});

test("part 1", async () => {
	expect(await part1()).toBe(238);
});

describe("isCorrectNiceString", () => {
	test("qjhvhtzxzqqjkmpb", () =>
		expect(isCorrectNiceString("qjhvhtzxzqqjkmpb")).toBe(true));
	test("xxyxx", () => expect(isCorrectNiceString("xxyxx")).toBe(true));
	test("uurcxstgmygtbstg", () =>
		expect(isCorrectNiceString("uurcxstgmygtbstg")).toBe(false));
	test("ieodomkazucvgmuy", () =>
		expect(isCorrectNiceString("ieodomkazucvgmuy")).toBe(false));
});

test("part 2", async () => {
	expect(await part2()).toBe(69);
});
