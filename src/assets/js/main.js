import "../scss/main.scss";
import "htmx.org"

document.querySelectorAll(".delete").forEach((el) => {
    el.addEventListener("click", (e) => {
        e.preventDefault();
        el.closest("div").remove();
    });
});
