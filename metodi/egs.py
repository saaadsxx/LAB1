def lagr(n, x, y, q):
    L = 0.0
    for i in range(n):
        s = 1.0
        for j in range(n):
            if j != i:
                s *= (q - x[j]) / (x[i] - x[j])  # Fixed the index and formula for Lagrange basis
        L += y[i] * s
    return round(L, 3)

# Initialize the known data points and the interpolation point
n = 10
x = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0]
y = [0.80, 0.29, 0.52, 0.77, 0.93, 1.20, 1.20, 1.35, 1.39, 1.48]

q = 1  # Start from the first point and iterate
for i in range(1, 20):  # Example loop for different q values
    print(f"Значение функции {lagr(n, x, y, q)} в точке {q}")
    q += 0.5  # Increment q by 0.5 for each iteration
