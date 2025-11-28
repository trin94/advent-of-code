import { describe, expect, test } from "bun:test";
import { countCharacters, encodeString, part1, part2 } from "./08.ts";

describe("countCharacters", () => {
	test('""', () => {
		const actual = countCharacters(`""`);
		expect(actual.total).toBe(2);
		expect(actual.interpreted).toBe(0);
	});
	test('"abc"', () => {
		const actual = countCharacters(`"abc"`);
		expect(actual.total).toBe(5);
		expect(actual.interpreted).toBe(3);
	});
	test('"aaa"aaa"', () => {
		const actual = countCharacters(String.raw`"aaa\"aaa"`);
		expect(actual.total).toBe(10);
		expect(actual.interpreted).toBe(7);
	});
	test('"\\oc"', () => {
		const actual = countCharacters(String.raw`"\\oc"`);
		expect(actual.total).toBe(6);
		expect(actual.interpreted).toBe(3);
	});
	test('"\x27"', () => {
		const actual = countCharacters(String.raw`"\x27"`);
		expect(actual.total).toBe(6);
		expect(actual.interpreted).toBe(1);
	});
});

test("demo 1", async () => {
	const result = await part1("08.test.txt");
	expect(result).toBe(12);
});

test("part 1", async () => {
	const result = await part1("08.txt");
	expect(result).toBe(1342);
});

describe("encodeString", () => {
	const mapping = [
		{ input: '""', expected: String.raw`"\"\""` },
		{ input: '"abc"', expected: String.raw`"\"abc\""` },
		{ input: String.raw`"aaa\"aaa"`, expected: String.raw`"\"aaa\\\"aaa\""` },
		{ input: String.raw`"\x27"`, expected: String.raw`"\"\\x27\""` },
	];

	test.each(mapping)(
		"encodeString($input) -> $expected",
		({ input, expected }) => {
			expect(encodeString(input)).toBe(expected);
		},
	);
});

test("demo 2", async () => {
	const result = await part2("08.test.txt");
	expect(result).toBe(19);
});

test("part 2", async () => {
	const result = await part2("08.txt");
	expect(result).toBe(2074);
});
