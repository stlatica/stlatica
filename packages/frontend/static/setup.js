let userTheme = localStorage.theme;

// light, dark以外の値だったら検出を行う
if (!["light", "dark"].includes(userTheme)) {
	const systemIsDark = window.matchMedia("(prefers-color-scheme: dark)").matches;

	const detectedTheme = systemIsDark ? "dark" : "light";

	localStorage.setItem("theme", detectedTheme);
	userTheme = detectedTheme;
}

// 初回ダークモード適用
if (userTheme === "dark" && document) {
	document.documentElement.classList.add("dark");
}
