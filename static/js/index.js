function clearInput() {
  document.getElementById('url').value = '';
}

document.body.addEventListener('htmx:afterSwap', function(event) {
  if (event.target.id === 'response') {
    document.getElementById('url').value = '';
  }
});
