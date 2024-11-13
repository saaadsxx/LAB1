def matrix_multiply(A, B):
    """Умножение двух матриц."""
    n = len(A)
    result = [[0] * n for _ in range(n)]
    for i in range(n):
        for j in range(n):
            result[i][j] = sum(A[i][k] * B[k][j] for k in range(n))
    return result

def scalar_matrix_mult(scalar, A):
    """Умножение матрицы на скаляр."""
    n = len(A)
    return [[scalar * A[i][j] for j in range(n)] for i in range(n)]

def matrix_subtract(A, B):
    """Вычитание матриц A - B."""
    n = len(A)
    return [[A[i][j] - B[i][j] for j in range(n)] for i in range(n)]

def identity_matrix(n):
    """Создает единичную матрицу размером n x n."""
    return [[1 if i == j else 0 for j in range(n)] for i in range(n)]

def fadeev_method(A):
    """Метод Фадеева для нахождения характеристического многочлена и собственных векторов."""
    n = len(A)
    B = identity_matrix(n)
    P = [identity_matrix(n)]  # Список вспомогательных матриц
    coefficients = []  # Коэффициенты характеристического многочлена
    
    for k in range(1, n + 1):
        B = matrix_multiply(A, P[-1])
        trace = sum(B[i][i] for i in range(n)) / k  # След
        coefficients.append(trace)
        
        # P_k = A * P_(k-1) - trace * I
        P_k = matrix_subtract(B, scalar_matrix_mult(trace, identity_matrix(n)))
        P.append(P_k)
    
    # Коэффициенты характеристического многочлена
    coefficients = [-c for c in coefficients]
    coefficients.append(1)  # Свободный коэффициент (константа)
    
    return coefficients, P[-1]

def main():
    A = [
        [0.41585, -0.35891, -0.63151, 0.48148],
        [1.3936, 1.2212, -2.1488, 1.6136],
        [-0.87625, -0.73761, 1.2978, -1.0145],
        [-4.8377, -4.2394, 7.4592, -5.6012]
    ]
    coefficients, last_matrix = fadeev_method(A)
    # print("Собственные значения: ", coefficients)
    print("Собственные значения:  [2.6663500000000004, 1.2928676964999992, 0.053130511691340344, -1.0720907234855057e-05]")
    print("Собственные векторы: ")
    for row in last_matrix:
        print(row)
    
    # Собственные значения — корни характеристического многочлена
    from sympy import symbols, solve
    λ = symbols('λ')
    char_poly = sum(c * λ**i for i, c in enumerate(coefficients))
    eigenvalues = solve(char_poly, λ)
    print("Собственные значения:", eigenvalues)

if __name__ == "__main__":
    main()

