export async function part1(): Promise<number> {
	const filename = "input/04.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();
	return findSecretKey(input.trim());
}

export function findSecretKey(prefix: string): number {
	const hasher = new Bun.CryptoHasher("md5");
	let num = 10000;
	while (true) {
		const hash = hasher.update(`${prefix}${num}`).digest("hex");
		if (hash.startsWith("00000")) {
			break;
		}
		num += 1;
	}
	return num;
}

export async function part2(): Promise<number> {
	const filename = "input/04.txt";
	const foo = Bun.file(filename);
	const input = await foo.text();
	return findSecretKey2(input.trim());
}

export function findSecretKey2(prefix: string): number {
	const hasher = new Bun.CryptoHasher("md5");
	let num = 100000;
	while (true) {
		const hash = hasher.update(`${prefix}${num}`).digest("hex");
		if (hash.startsWith("000000")) {
			break;
		}
		num += 1;
	}
	return num;
}
