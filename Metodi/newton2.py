import math

# Функции f1 и f2
def f1(x, y):
    return math.cos(x - 1) + y - 0.5

def f2(x, y):
    return x - math.cos(y) - 3

# Частные производные функций f1 и f2
def df1_dx(x, y):
    return -math.sin(x - 1)

def df1_dy(x, y):
    return 1

def df2_dx(x, y):
    return 1

def df2_dy(x, y):
    return math.sin(y)

# Начальные приближения
x = 1.0
y = 1.0
tolerance = 1e-6
max_iterations = 100

# Метод Ньютона
for i in range(max_iterations):
    # Вычисляем значения функций
    F1 = f1(x, y)
    F2 = f2(x, y)
    
    # Проверка на достижение заданной точности
    if abs(F1) < tolerance and abs(F2) < tolerance:
        print(f"итерицй: {i + 1}")
        print(f"x ≈ {x}, y ≈ {y}")
        break

    # Вычисляем элементы якобиана
    J11 = df1_dx(x, y)
    J12 = df1_dy(x, y)
    J21 = df2_dx(x, y)
    J22 = df2_dy(x, y)
    
    # Определитель якобиана
    det_J = J11 * J22 - J12 * J21
    if det_J == 0:
        print("Якобиан вырожден, метод не применим.")
        break

    # Вычисляем приращения Δx и Δy
    delta_x = (-F1 * J22 + F2 * J12) / det_J
    delta_y = (F1 * J21 - F2 * J11) / det_J
    
    # Обновляем значения x и y
    x += delta_x
    y += delta_y

else:
    print("Решение не найдено за заданное количество итераций")

# Вывод конечных приближений
print(f"Последние приближения: x ≈ {x}, y ≈ {y}")
