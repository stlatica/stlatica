export const FormdataToObject = (form: FormData) => {
	const formObject: { [key: string]: FormDataEntryValue } = {};
	for (const [key, value] of form.entries()) {
		formObject[key] = value;
	}

	return formObject;
};
