def matrix_vector_multiply(A, v):
    """Умножение матрицы на вектор."""
    n = len(A)
    result = [0] * n
    for i in range(n):
        result[i] = sum(A[i][j] * v[j] for j in range(n))
    return result

def generate_krylov_matrix(A, v):
    """Генерация крыловской матрицы на основе начального вектора."""
    n = len(A)
    K = [[0] * n for _ in range(n)]
    K[0] = v[:]  # Начальный вектор

    for i in range(1, n):
        K[i] = matrix_vector_multiply(A, K[i - 1])
    
    # Транспонируем матрицу для работы с характеристическим многочленом
    return [[K[j][i] for j in range(n)] for i in range(n)]

def solve_characteristic_polynomial(K):
    """Решение системы для нахождения коэффициентов характеристического многочлена."""
    n = len(K)
    coefficients = [0] * (n + 1)
    coefficients[n] = 1  # Коэффициент перед λ^n

    # Решаем систему уравнений K * c = 0, где c - вектор коэффициентов
    # Поскольку матрица K формирует систему вида [K | b] с последним столбцом b = -K[:,n]
    for i in range(n - 1, -1, -1):
        coefficients[i] = -sum(K[i][j] * coefficients[j] for j in range(i + 1, n)) / K[i][i]
    
    return coefficients

def krylov_method(A, v):
    """Реализация метода Крылова."""
    K = generate_krylov_matrix(A, v)
    print("Крыловская матрица:")
    for row in K:
        print(row)

    coefficients = solve_characteristic_polynomial(K)
    return coefficients

def main():
    A = [
        [18.736, -0.080167, -0.25998, 0.16584],
        [-29.757, 0.057914, 0.18781, -0.26340],
        [31.857, -0.12683, -0.41130, 0.28199],
        [-45.869, 0.089269, 0.28950, -0.40601]
    ]
    v = [1, 0, 0, 0]  # Начальный вектор

    coefficients = krylov_method(A, v)
    print("Характеристический многочлен:", coefficients)

    # Находим собственные значения как корни характеристического многочлена
    from sympy import symbols, solve
    λ = symbols('λ')
    char_poly = sum(c * λ**i for i, c in enumerate(coefficients))
    eigenvalues = solve(char_poly, λ)
    print("Собственные значения:", eigenvalues)

if __name__ == "__main__":
    main()
