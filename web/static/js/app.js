// Функция для изменения иконки
function changeIcon() {
    const modal = document.getElementById('iconModal');
    modal.style.display = 'block';
}

// Закрытие модального окна
function closeModal() {
    const modal = document.getElementById('iconModal');
    modal.style.display = 'none';
}

// Обработчик загрузки нового изображения
function uploadIcon() {
    const fileInput = document.getElementById('iconUpload');
    const file = fileInput.files[0];

    if (file) {
        const reader = new FileReader();
        reader.onload = function(e) {
            const icon = document.getElementById('chatIcon');
            icon.src = e.target.result;
        };
        reader.readAsDataURL(file);
    }

    closeModal();
}

// Закрытие модального окна при клике вне его области
window.onclick = function(event) {
    const modal = document.getElementById('iconModal');
    if (event.target === modal) {
        closeModal();
    }
};
