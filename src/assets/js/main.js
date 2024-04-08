import "../scss/main.scss";

function removelementOnClose() {
    document.querySelectorAll(".delete").forEach((el) => {
        el.addEventListener("click", (e) => {
            e.preventDefault();
            el.closest("div").remove();
        });
    });
}

import("./quill").then(({ default: initializeEditor }) => {
    initializeEditor();
});

function main() {
        removelementOnClose();
    }

main();
