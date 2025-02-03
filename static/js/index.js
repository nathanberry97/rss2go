/**
 * The following code handles allowing users to get to the top of the screen
 */
window.onscroll = function () {
    const toTopButton = document.getElementById("toTop");
    if (document.documentElement.scrollTop > 100 || document.body.scrollTop > 100) {
        toTopButton.classList.add("show");
    } else {
        toTopButton.classList.remove("show");
    }
};

function toTop() {
    // For Safari
    document.body.scrollTo({
        top: 0,
        behavior: "smooth",
    });

    // For Chrome, Firefox, IE and Opera
    document.documentElement.scrollTo({
        top: 0,
        behavior: "smooth",
    });
}

/**
 * Function to clear forms
 */
function clearInput() {
    document.getElementById("url").value = "";
}
