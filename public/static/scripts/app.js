document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll("a.nav-link").forEach(link => {
        link.addEventListener("click", (event) => {
            event.preventDefault();
            const url = link.href;

            if (!document.startViewTransition) {
                window.location.href = url; // Fallback for unsupported browsers
                return;
            }

            document.startViewTransition(() => {
                return new Promise(resolve => {
                    fetch(url)
                        .then(response => response.text())
                        .then(html => {
                            const parser = new DOMParser();
                            const doc = parser.parseFromString(html, "text/html");
                            document.body.innerHTML = doc.body.innerHTML;
                            resolve();
                        });
                });
            });
        });
    });
});
