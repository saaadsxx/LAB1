import math

def norm(v):
    return math.sqrt(sum(x ** 2 for x in v))

def normalize(v):
    n = norm(v)
    return [x / n for x in v]

def gram_schmidt(A):
    n = len(A)
    m = len(A[0])
    Q = [[0] * m for _ in range(n)]
    R = [[0] * m for _ in range(m)]
    
    for j in range(m):
        v = [A[i][j] for i in range(n)]
        
        for i in range(j):
            R[i][j] = sum(A[x][j] * Q[x][i] for x in range(n))
            v = [v[x] - R[i][j] * Q[x][i] for x in range(n)]
        
        R[j][j] = norm(v)
        qj = normalize(v)  # Нормализуем текущий вектор
        
        for i in range(n):
            Q[i][j] = qj[i]  # Заполняем столбец Q
    
    return Q, R

def eigenvalues_from_qr(A, max_iterations=100):
    """Метод QR для нахождения собственных значений"""
    n = len(A)
    for _ in range(max_iterations):
        Q, R = gram_schmidt(A)
        A = [[sum(R[i][k] * Q[k][j] for k in range(n)) for j in range(n)] for i in range(n)]
    
    # Собственные значения будут на диагонали матрицы A
    eigenvalues = [A[i][i] for i in range(n)]
    return eigenvalues

# Пример использования
A = [
    [18.736, -0.080167, -0.25998, 0.16584],
    [-29.757, 0.057914, 0.18781, -0.26340],
    [31.857, -0.12683, -0.41130, 0.28199],
    [-45.869, 0.089269, 0.28950, -0.40601]
]

Q, R = gram_schmidt(A)

print("Матрица Q:")
for row in Q:
    print(row)

print("\nМатрица R:")
for row in R:
    print(row)

# Вычисление собственных значений
eigenvalues = eigenvalues_from_qr(A)
print("\nСобственные значения матрицы A:")
for value in eigenvalues:
    print(value)
