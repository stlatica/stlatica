import { Sleep } from './utilities/Sleep';

/**
 * 大量の非同期処理を、同時実行数を制限しながら実行します
 */
export const AsyncQueue = <T, Return>(
	data: T[],
	callback: (data: T, index: number) => Promise<Return>,
	_limit: number = 10
): Promise<Return[]> => {
	const limit = Math.max(_limit, 1);

	let cnt = 0;

	return Promise.all(
		data.map(async (d, index) => {
			while (true) {
				// console.log({ window: cnt + limit, index });
				if (index < cnt + limit) {
					// console.log({ window: cnt + limit, cnt, index });
					const r = await callback(d, index);
					cnt += 1;
					return r;
				}
				await Sleep(10);
			}
		})
	);
};
