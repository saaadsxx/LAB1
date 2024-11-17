import numpy as np
import matplotlib.pyplot as plt


x = np.array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0])
y = np.array([2.23, 2.29, 2.27, 2.62, 2.72, 2.82, 3.13, 3.49, 3.82, 3.95, 4.22, 4.48, 5.05, 5.50, 5.68, 6.19, 6.42, 7.04, 7.57, 8.10])

n = len(x) - 1  
h = np.diff(x)  


A = np.zeros((n + 1, n + 1))
b = np.zeros(n + 1)


A[0, 0] = 1
A[n, n] = 1


for i in range(1, n):
    A[i, i - 1] = h[i - 1]
    A[i, i] = 2 * (h[i - 1] + h[i])
    A[i, i + 1] = h[i]
    b[i] = (3 / h[i]) * (y[i + 1] - y[i]) - (3 / h[i - 1]) * (y[i] - y[i - 1])


def gauss_elimination(A, b):
    n = len(b)
    for i in range(n):
        
        for j in range(i + 1, n):
            factor = A[j, i] / A[i, i]
            A[j] -= factor * A[i]
            b[j] -= factor * b[i]
    
    
    x = np.zeros(n)
    for i in range(n - 1, -1, -1):
        x[i] = (b[i] - np.dot(A[i, i + 1:], x[i + 1:])) / A[i, i]
    return x


c = gauss_elimination(A.copy(), b)


a = y[:-1]
b_coeff = np.zeros(n)
d = np.zeros(n)

for i in range(n):
    b_coeff[i] = (y[i + 1] - y[i]) / h[i] - h[i] * (2 * c[i] + c[i + 1]) / 3
    d[i] = (c[i + 1] - c[i]) / (3 * h[i])


def cubic_spline(x_val):
    if x_val < x[0] or x_val > x[-1]:
        raise ValueError("x_val is out of the interpolation range.")
    
    
    for i in range(n):
        if x[i] <= x_val <= x[i + 1]:
            dx = x_val - x[i]
            return a[i] + b_coeff[i] * dx + c[i] * dx**2 + d[i] * dx**3


x_new = np.linspace(1, 20, 100)
y_new = [cubic_spline(xi) for xi in x_new]


plt.scatter(x, y, color='red', label="Nodes")
plt.plot(x_new, y_new, label='Cubic Spline', color='blue')
plt.title("Cubic Spline Interpolation")
plt.xlabel('x')
plt.ylabel('y')
plt.legend()
plt.grid()
plt.show()

# Вывод значений функции в точках от 1 до 10 с шагом 0.5
x_values = np.arange(1, 10.5, 0.5)
for x_val in x_values:
    try:
        y_val = cubic_spline(x_val)
        print(f"x: {x_val:.1f}, y: {y_val:.2f}")
    except ValueError as e:
        print(f"x: {x_val:.1f}, error: {e}")
