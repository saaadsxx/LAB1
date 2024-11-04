import math

# Задаем начальные приближения
x = 1.0  # Начальное значение для x
y = 1.0  # Начальное значение для y
tolerance = 1e-6  # Заданная точность
max_iterations = 100  # Максимальное количество итераций

def iterate(x, y):
    """Одна итерация метода простых итераций."""
    x_new = math.cos(y) + 3
    y_new = 0.5 - math.cos(x - 1)
    return x_new, y_new

# Итерационный процесс
for i in range(max_iterations):
    x_new, y_new = iterate(x, y)
    # Проверка на достижение требуемой точности
    if abs(x_new - x) < tolerance and abs(y_new - y) < tolerance:
        print(f"итерций: {i + 1}")
        print(f"x ≈ {x_new}, y ≈ {y_new}")
        break
    x, y = x_new, y_new
else:
    print("Решение не найдено за заданное количество итераций")

# Вывод конечных приближений
print(f"Последние приближения: x ≈ {x}, y ≈ {y}")
