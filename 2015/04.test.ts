import { describe, expect, test } from "bun:test";
import { findSecretKey, part1, part2 } from "./04.ts";

describe("findSecretKey", () => {
	test("abcdef", () => {
		expect(findSecretKey("abcdef")).toBe(609043);
	});
	test("pqrstuv", () => {
		expect(findSecretKey("pqrstuv")).toBe(1048970);
	});
});

test("part 1", async () => {
	expect(await part1()).toBe(254575);
});

test("part 2", async () => {
	expect(await part2()).toBe(1038736);
});
