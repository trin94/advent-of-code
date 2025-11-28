import { expect, test } from "bun:test";
import { calculateCircuit, part1, part2 } from "./07.ts";

test("test input", async () => {
	const results = await calculateCircuit("07.test.txt");
	expect(results.get("d")).toBe(72);
	expect(results.get("e")).toBe(507);
	expect(results.get("f")).toBe(492);
	expect(results.get("g")).toBe(114);
	expect(results.get("h")).toBe(65412);
	expect(results.get("i")).toBe(65079);
	expect(results.get("xx")).toBe(123);
	expect(results.get("y")).toBe(456);
});

test("part 1", async () => {
	const result = await part1("07.txt");
	expect(result).toBe(3176);
});

test("part 2", async () => {
	const result = await part2("07.txt");
	expect(result).toBe(14710);
});
