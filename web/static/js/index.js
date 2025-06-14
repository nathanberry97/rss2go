/**
 * The following code handles allowing users to get to the top of the screen
 */
window.onscroll = function() {
    const toTopButton = document.getElementById("toTopBtn");

    if (document.documentElement.scrollTop > 100 || document.body.scrollTop > 100) {
        toTopButton.classList.add("show");
        toTopButton.style.visibility = "visible";
    } else {
        toTopButton.classList.remove("show");
        setTimeout(() => {
            if (!toTopButton.classList.contains("show")) {
                toTopButton.style.visibility = "hidden";
            }
        }, 300);
    }
};

function toTop() {
    document.body.scrollTo({ top: 0, behavior: "smooth" });
    document.documentElement.scrollTo({ top: 0, behavior: "smooth" });
}

/**
 * Function to clear forms
 */
function clearInput() {
    const urlInput = document.getElementById("url");
    const fileInput = document.getElementById("avatarInput");

    if (urlInput) urlInput.value = "";
    if (fileInput) fileInput.value = "";
}

/**
 * Handle hamburger menu
 */
document.addEventListener("DOMContentLoaded", () => {
    const hamburger = document.querySelector(".navbar__hamburger");
    const navigation = document.querySelector(".navbar__navigation");
    const navLinks = document.querySelectorAll(".navbar__link");

    if (hamburger && navigation) {
        hamburger.addEventListener("click", () => {
            navigation.classList.toggle("open");
        });
    }

    navLinks.forEach(link => {
        link.addEventListener("click", () => {
            navigation.classList.remove("open");
        });
    });
});
