def lagr(n, x, y, q):
    L = 0.0
    for i in range(n):
        s = 1.0
        for j in range(n):
            if j != i:
                s *= (q - x[j]) / (x[i] - x[j])  
        L += y[i] * s
    return round(L, 3)


n = 10
x = [1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0]
y = [2.23, 2.29, 2.27, 2.62, 2.72, 2.82, 3.13, 3.49, 3.82, 3.95, 4.22, 4.48, 5.05, 5.50, 5.68, 6.19, 6.42, 7.04, 7.57, 8.10]

q = 1  
for i in range(1, 20):  
    print(f" x: {q} y: {lagr(n, x, y, q)}")
    q += 0.5  
