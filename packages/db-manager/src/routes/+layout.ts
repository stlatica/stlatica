// 時刻の bigint を無理やり表示するハック。行儀が良くない。暫定措置。

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
BigInt.prototype.toJSON = function () {
	const int = Number.parseInt(this.toString());
	return new Date(int).toISOString() ?? this.toString();
};
