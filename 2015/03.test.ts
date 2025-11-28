import { describe, expect, test } from "bun:test";
import {
	countDeliveries,
	countDeliveriesWithSupport,
	part1,
	part2,
} from "./03.ts";

describe("countDeliveries", () => {
	test(">", () => {
		expect(countDeliveries(">")).toBe(2);
	});
	test("^>v<", () => {
		expect(countDeliveries("^>v<")).toBe(4);
	});
	test("^v^v^v^v^v", () => {
		expect(countDeliveries("^v^v^v^v^v")).toBe(2);
	});
	test("input.txt", async () => {
		expect(await part1()).toBe(2565);
	});
});

describe("countDeliveriesWithSupport", () => {
	test("^v", () => {
		expect(countDeliveriesWithSupport("^v")).toBe(3);
	});
	test("^>v<", () => {
		expect(countDeliveriesWithSupport("^>v<")).toBe(3);
	});
	test("^v^v^v^v^v", () => {
		expect(countDeliveriesWithSupport("^v^v^v^v^v")).toBe(11);
	});
	test("input.txt", async () => {
		expect(await part2()).toBe(2639);
	});
});
