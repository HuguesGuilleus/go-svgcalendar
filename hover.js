const hover = Array.from(document.querySelectorAll("text")).reverse()[0],
	fontSize = window.getComputedStyle(hover).fontSize;

document.querySelectorAll("rect").forEach(r => r.addEventListener("mouseover", () => {
	hover.textContent =
		new Date(r.attributes.date.value).toLocaleDateString() +
		" : " +
		r.attributes.value.value;

	Array.from(r.children).forEach(t => {
		const tspan = document.createElementNS("http://www.w3.org/2000/svg", "tspan");
		tspan.textContent = t.textContent;
		tspan.setAttribute('x', '0');
		tspan.setAttribute('dy', fontSize);
		hover.appendChild(tspan);
	});
}));