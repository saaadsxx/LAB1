document.addEventListener('DOMContentLoaded', function () {
    const burgerMenu = document.getElementById('burger-menu');
    const sidebar = document.getElementById('sidebar');
    const menuLinks = document.querySelectorAll('.menu-link');
  
    // Открытие/закрытие сайдбара при клике на бургер-меню
    burgerMenu.addEventListener('click', function() {
      sidebar.classList.toggle('active');
    });
  
    // Закрытие сайдбара при клике на ссылку в меню
    menuLinks.forEach(link => {
      link.addEventListener('click', function() {
        sidebar.classList.remove('active');
      });
    });
  });
  