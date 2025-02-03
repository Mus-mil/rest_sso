document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("loginForm");
    const errorMessage = document.getElementById("error-message");

    form.addEventListener("submit", async function (event) {
        event.preventDefault();

        const username = document.getElementById("username").value.trim();
        const password = document.getElementById("password").value.trim();

        if (!username || !password) {
            errorMessage.textContent = "Заполните все поля!";
            return;
        }

        try {
            const response = await fetch("/auth/signin", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username, password }),
                credentials: "include" // Для работы с куками
            });

            const data = await response.json();

            if (response.ok) {
                window.location.href = data.redirect; // Динамический редирект
            } else {
                errorMessage.textContent = data.message || "Ошибка входа";
            }
        } catch (error) {
            errorMessage.textContent = "Ошибка соединения с сервером";
        }
    });
});
