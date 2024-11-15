export const Sleep = (timeout: number) => {
	return new Promise((resolve) => {
		setTimeout(() => resolve(timeout), timeout);
	});
};
