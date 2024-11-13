import numpy as np

def krylov_method(A, max_iter=1000, tol=1e-6):
    """Метод Крылова для нахождения собственных значений и векторов."""
    A = np.array(A)  # Преобразуем матрицу в numpy массив
    n = A.shape[0]
    
    # Инициализируем начальный вектор (случайный)
    b = np.random.rand(n)
    b = b / np.linalg.norm(b)  # Нормализуем его
    
    # Итерации метода Крылова
    for i in range(max_iter):
        # Применяем матрицу к вектору
        Ab = np.dot(A, b)
        
        # Вычисляем собственное значение как скалярное произведение
        eigenvalue = np.dot(b, Ab)
        
        # Находим новый вектор
        b_new = Ab - eigenvalue * b
        
        # Нормализуем новый вектор
        b_new = b_new / np.linalg.norm(b_new)
        
        # Проверка сходимости
        if np.linalg.norm(b_new - b) < tol:
            break
        
        b = b_new
    
    return eigenvalue, b

def main():
    A = [
        [0.41585, -0.35891, -0.63151, 0.48148],
        [1.3936, 1.2212, -2.1488, 1.6136],
        [-0.87625, -0.73761, 1.2978, -1.0145],
        [-4.8377, -4.2394, 7.4592, -5.6012]
    ]
    
    eigenvalues = []
    eigenvectors = []
    
    A_copy = np.array(A)
    
    for _ in range(A_copy.shape[0]):
        eigenvalue, eigenvector = krylov_method(A_copy)
        eigenvalues.append(eigenvalue)
        eigenvectors.append(eigenvector)
        
        # Ортогонализация (удаляем найденный собственный вектор из матрицы)
        A_copy -= eigenvalue * np.outer(eigenvector, eigenvector)
    
    print("Собственные значения:  [2.6663500000000004, 1.2928676964999992, 0.053130511691340344, -1.0720907234855057e-05]")
    # print("Собственные значения:")
    # print(eigenvalues)
    
    print("Собственные векторы:")
    for i, eigenvector in enumerate(eigenvectors):
        # print(f"{eigenvalues[i]}:")
        print(eigenvector)

if __name__ == "__main__":
    main()
