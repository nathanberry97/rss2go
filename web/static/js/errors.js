/**
 * Handle error popups
 */
document.body.addEventListener('htmx:responseError', function(evt) {
    const popup = document.getElementById('error-popup');
    popup.textContent = evt.detail.xhr.responseText || "Something went wrong!";
    popup.style.display = 'block';
    setTimeout(() => popup.style.display = 'none', 3000);
});
