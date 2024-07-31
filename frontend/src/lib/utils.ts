function formatDate(dateString: string): string {
	const parts = dateString.split(/[+. ]/); // Split on space or +
	const formattedDate = parts.slice(0, 2).join("T").concat("Z");
	const date = new Date(formattedDate); // Parse the date string
	const options: Intl.DateTimeFormatOptions = {
		year: "numeric",
		month: "2-digit",
		day: "2-digit",
	};
	return new Intl.DateTimeFormat("en-CA", options).format(date);
}
export { formatDate };
