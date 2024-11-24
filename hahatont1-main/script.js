// Show modal
  function showTableModal() {
    document.getElementById("tableModal").style.display = "flex";
  }
  
  // Close modal
  function closeTableModal() {
    document.getElementById("tableModal").style.display = "none";
  }

  let currentChartType = '';

  // Show chart modal
  function showChartModal(type) {
    currentChartType = type;
    document.getElementById('chartModal').style.display = 'flex';
  }

  // Close chart modal
  function closeChartModal() {
    document.getElementById('chartModal').style.display = 'none';
    document.getElementById('chartLabels').value = '';
    document.getElementById('chartData').value = '';
  }

  // Create chart
  function createChart() {
    const labels = document.getElementById('chartLabels').value.split(',');
    const data = document.getElementById('chartData').value.split(',').map(Number);
    if (labels.length !== data.length) {
      alert('Labels and data must have the same length!');
      return;
    }

    const chartContainer = document.createElement('div');
    chartContainer.className = 'report-element';

    const canvas = document.createElement('canvas');
    chartContainer.appendChild(canvas);
    document.getElementById('report-container').appendChild(chartContainer);

    const ctx = canvas.getContext('2d');
    const chartConfig = {
      type: currentChartType === 'histogram' ? 'bar' : currentChartType,  // Treat histogram as bar chart
      data: {
        labels: labels,
        datasets: [{
          label: currentChartType.charAt(0).toUpperCase() + currentChartType.slice(1) + ' Chart',
          data: data,
          backgroundColor: currentChartType === 'pie' ? getPieColors(labels.length) : 'rgba(75, 192, 192, 0.2)',
          borderColor: currentChartType === 'pie' ? getPieColors(labels.length) : 'rgba(75, 192, 192, 1)',
          borderWidth: 1
        }]
      },
      options: {
        scales: currentChartType === 'pie' ? {} : {
          y: {
            beginAtZero: true
          }
        }
      }
    };

    new Chart(ctx, chartConfig);
    closeChartModal();
  }

  // Helper function to generate colors for pie chart
  function getPieColors(count) {
    const colors = [];
    for (let i = 0; i < count; i++) {
      const randomColor = `hsl(${Math.floor(Math.random() * 360)}, 100%, 75%)`;
      colors.push(randomColor);
    }
    return colors;
  }


  // Create table with specified rows and columns
  function createTable() {
    const rows = document.getElementById("tableRows").value;
    const columns = document.getElementById("tableColumns").value;
    addTable(parseInt(rows), parseInt(columns));
    closeTableModal();
  }
  
  // Function to add a table
  function addTable(rows, columns) {
    const reportContainer = document.getElementById("report-container");
    const tableElement = document.createElement("div");
    tableElement.className = "report-element";
  
    const table = document.createElement("table");
    table.style.width = "100%";
    table.style.borderCollapse = "collapse";
  
    for (let i = 0; i < rows; i++) {
      const row = table.insertRow();
      for (let j = 0; j < columns; j++) {
        const cell = row.insertCell();
        cell.contentEditable = "true";
        cell.style.border = "1px solid #ced4da";
        cell.style.padding = "8px";
        cell.textContent = `Cell ${i + 1}, ${j + 1}`;
      }
    }
  
    const addRowButton = document.createElement("button");
    addRowButton.textContent = "Add Row";
    addRowButton.onclick = () => {
      const row = table.insertRow();
      for (let i = 0; i < columns; i++) {
        const cell = row.insertCell();
        cell.contentEditable = "true";
        cell.style.border = "1px solid #ced4da";
        cell.style.padding = "8px";
        cell.textContent = `New Cell`;
      }
    };
  
    const addColumnButton = document.createElement("button");
    addColumnButton.textContent = "Add Column";
    addColumnButton.onclick = () => {
      for (let i = 0; i < table.rows.length; i++) {
        const cell = table.rows[i].insertCell();
        cell.contentEditable = "true";
        cell.style.border = "1px solid #ced4da";
        cell.style.padding = "8px";
        cell.textContent = `New Cell`;
      }
    };
  
    const removeButton = document.createElement("button");
    removeButton.className = "remove-button";
    removeButton.textContent = "✕";
    removeButton.onclick = () => tableElement.remove();
  
    tableElement.appendChild(table);
    tableElement.appendChild(addRowButton);
    tableElement.appendChild(addColumnButton);
    tableElement.appendChild(removeButton);
    reportContainer.appendChild(tableElement);
  }
  
  // Adding a title
  function addTitle() {
    const reportContainer = document.getElementById("report-container");
    const titleElement = document.createElement("div");
    titleElement.className = "report-element";
    const titleInput = document.createElement("input");
    titleInput.type = "text";
    titleInput.placeholder = "Enter report title...";
    const removeButton = document.createElement("button");
    removeButton.className = "remove-button";
    removeButton.textContent = "✕";
    removeButton.onclick = () => titleElement.remove();
    titleElement.appendChild(titleInput);
    titleElement.appendChild(removeButton);
    reportContainer.appendChild(titleElement);
  }
  
  // Adding a paragraph
  function addParagraph() {
    const reportContainer = document.getElementById("report-container");
    const paragraphElement = document.createElement("div");
    paragraphElement.className = "report-element";
    const paragraphTextarea = document.createElement("textarea");
    paragraphTextarea.rows = 3;
    paragraphTextarea.placeholder = "Enter report paragraph...";
    const removeButton = document.createElement("button");
    removeButton.className = "remove-button";
    removeButton.textContent = "✕";
    removeButton.onclick = () => paragraphElement.remove();
    paragraphElement.appendChild(paragraphTextarea);
    paragraphElement.appendChild(removeButton);
    reportContainer.appendChild(paragraphElement);
  }
  
// Load and display CSV data with editable cells and add row/column functionality
    function loadCSV(event) {
        const files = event.target.files;
        if (!files.length) return;

        Array.from(files).forEach((file) => {
            const reader = new FileReader();
            reader.onload = function (e) {
                const csv = e.target.result;
                const delimiter = csv.includes(";") ? ";" : ",";
                const rows = csv.trim().split("\n").filter(row => row).map(row => row.split(delimiter));

                const tableElement = document.createElement("div");
                tableElement.className = "report-element";
                const table = document.createElement("table");

                // Treat the first row as headers
                const headerRow = table.insertRow();
                rows[0].forEach(headerData => {
                    const headerCell = document.createElement("th");
                    headerCell.contentEditable = "true"; // Make headers editable
                    headerCell.textContent = headerData.trim();
                    headerRow.appendChild(headerCell);
                });

                // Add the rest of the rows as data
                rows.slice(1).forEach(rowData => {
                    const row = table.insertRow();
                    rowData.forEach(cellData => {
                        const cell = document.createElement("td");
                        cell.contentEditable = "true"; // Make data cells editable
                        cell.textContent = cellData.trim();
                        row.appendChild(cell);
                    });
                });

                // Add Row button
                const addRowButton = document.createElement("button");
                addRowButton.textContent = "Add Row";
                addRowButton.onclick = () => {
                    const newRow = table.insertRow();
                    for (let i = 0; i < rows[0].length; i++) {
                        const newCell = newRow.insertCell();
                        newCell.contentEditable = "true";
                        newCell.textContent = "New Cell";
                    }
                };

                // Add Column button
                const addColumnButton = document.createElement("button");
                addColumnButton.textContent = "Add Column";
                addColumnButton.onclick = () => {
                    const headerCell = document.createElement("th");
                    headerCell.contentEditable = "true";
                    headerCell.textContent = "New Header";
                    headerRow.appendChild(headerCell); // Add new header

                    for (let i = 1; i < table.rows.length; i++) {
                        const newCell = table.rows[i].insertCell();
                        newCell.contentEditable = "true";
                        newCell.textContent = "New Cell";
                    }
                };

                tableElement.appendChild(table);
                tableElement.appendChild(addRowButton);
                tableElement.appendChild(addColumnButton);
                addRemoveButton(tableElement);
                document.getElementById("report-container").appendChild(tableElement);
            };
            reader.readAsText(file);
        });

        event.target.value = ""; // Clears the file input
    }


    // Load and display JSON data with editable inputs and add key-value functionality
    function loadJSON(event) {
        const file = event.target.files[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onload = function (e) {
            const data = JSON.parse(e.target.result);
            for (const [key, value] of Object.entries(data)) {
                const element = document.createElement("div");
                element.className = "report-element";

                if (typeof value === "string") {
                    const input = document.createElement("input");
                    input.type = "text";
                    input.value = value;
                    input.contentEditable = "true"; // Make the input field editable
                    element.innerHTML = `<strong>${key}:</strong> `;
                    element.appendChild(input);
                } else if (Array.isArray(value)) {
                    const list = document.createElement("ul");
                    value.forEach((item) => {
                        const listItem = document.createElement("li");
                        listItem.contentEditable = "true"; // Make list items editable
                        listItem.textContent = item;
                        list.appendChild(listItem);
                    });
                    element.innerHTML = `<strong>${key}:</strong>`;

                    // Add Item button for arrays
                    const addItemButton = document.createElement("button");
                    addItemButton.textContent = "Add Item";
                    addItemButton.onclick = () => {
                        const newItem = document.createElement("li");
                        newItem.contentEditable = "true";
                        newItem.textContent = "New Item";
                        list.appendChild(newItem);
                    };
                    element.appendChild(list);
                    element.appendChild(addItemButton);
                } else if (typeof value === "object") {
                    element.innerHTML = `<strong>${key}:</strong>`;
                    const innerTable = document.createElement("table");
                    for (const [subKey, subValue] of Object.entries(value)) {
                        const row = innerTable.insertRow();
                        row.insertCell(0).textContent = subKey;
                        const cell = row.insertCell(1);
                        cell.contentEditable = "true"; // Make cells editable
                        cell.textContent = subValue;
                    }
                    element.appendChild(innerTable);

                    // Add Key-Value button for objects
                    const addKeyValueButton = document.createElement("button");
                    addKeyValueButton.textContent = "Add Key-Value";
                    addKeyValueButton.onclick = () => {
                        const row = innerTable.insertRow();
                        const newKey = row.insertCell(0);
                        newKey.contentEditable = "true";
                        newKey.textContent = "New Key";
                        const newValue = row.insertCell(1);
                        newValue.contentEditable = "true";
                        newValue.textContent = "New Value";
                    };
                    element.appendChild(addKeyValueButton);
                }

                addRemoveButton(element);
                document.getElementById("report-container").appendChild(element);
            }
        };
        reader.readAsText(file);

        event.target.value = ""; // Clears the file input
    }

    
  // Function to create a button
function createButton(text, onClick) {
    const button = document.createElement("button");
    button.textContent = text;
    button.onclick = onClick;
    return button;
  }
  
  // Universal function to add a row to a table
  function addRowToTable(table) {
    const columns = table.rows[0].cells.length; // Количество колонок
    const newRow = table.insertRow();
    for (let i = 0; i < columns; i++) {
      const newCell = newRow.insertCell();
      newCell.contentEditable = "true";
      newCell.style.border = "1px solid #ced4da";
      newCell.style.padding = "8px";
      newCell.textContent = "New Cell";
    }
  }
  
  // Universal function to add a column to a table
  function addColumnToTable(table) {
    // Добавляем новый заголовок
    const headerRow = table.rows[0];
    const newHeaderCell = document.createElement("th");
    newHeaderCell.contentEditable = "true";
    newHeaderCell.style.border = "1px solid #ced4da";
    newHeaderCell.style.padding = "8px";
    newHeaderCell.textContent = "New Header";
    headerRow.appendChild(newHeaderCell);
  
    // Добавляем новые ячейки в каждую строку
    for (let i = 1; i < table.rows.length; i++) {
      const newCell = table.rows[i].insertCell();
      newCell.contentEditable = "true";
      newCell.style.border = "1px solid #ced4da";
      newCell.style.padding = "8px";
      newCell.textContent = "New Cell";
    }
  }
  
  // Universal function to remove the last row from a table
  function removeLastRowFromTable(table) {
    if (table.rows.length > 1) {
      table.deleteRow(-1); // Удаляем последнюю строку, но не трогаем заголовок
    }
  }
  
  // Universal function to remove the last column from a table
  function removeLastColumnFromTable(table) {
    const columnCount = table.rows[0].cells.length;
    if (columnCount > 1) {
      // Удаляем последнюю ячейку в каждой строке
      for (let i = 0; i < table.rows.length; i++) {
        table.rows[i].deleteCell(-1);
      }
    }
  }
  
  // Sort table by column
  function sortTableByColumn(table, columnIndex, ascending) {
    const rowsArray = Array.from(table.rows).slice(1); // Обрезаем первую строку с заголовками
    rowsArray.sort((rowA, rowB) => {
      const cellA = rowA.cells[columnIndex].textContent.trim();
      const cellB = rowB.cells[columnIndex].textContent.trim();
      
      if (!isNaN(cellA) && !isNaN(cellB)) { // Сортировка чисел
        return ascending ? cellA - cellB : cellB - cellA;
      } else { // Сортировка строк
        return ascending ? cellA.localeCompare(cellB) : cellB.localeCompare(cellA);
      }
    });
  
    // Обновляем порядок строк в таблице
    rowsArray.forEach(row => table.appendChild(row));
  }
  
  // Add sort functionality to header cells
  function addSortingToTable(table, columns) {
    const headerRow = table.rows[0]; // Первая строка — это заголовки
    for (let i = 0; i < columns; i++) {
      const headerCell = headerRow.cells[i];
      let ascending = true;
      headerCell.onclick = () => {
        sortTableByColumn(table, i, ascending);
        ascending = !ascending; // Переключаем порядок сортировки
      };
    }
  }
  
  // Add a table
  function addTable(rows, columns) {
    const reportContainer = document.getElementById("report-container");
    const tableElement = document.createElement("div");
    tableElement.className = "report-element";
  
    const table = document.createElement("table");
    table.style.width = "100%";
    table.style.borderCollapse = "collapse";
  
    const headerRow = table.insertRow(); // Создаем строку заголовков
  
    for (let i = 0; i < columns; i++) {
      const headerCell = document.createElement("th");
      headerCell.contentEditable = "true";
      headerCell.style.border = "1px solid #ced4da";
      headerCell.style.padding = "8px";
      headerCell.textContent = `Header ${i + 1}`;
      headerRow.appendChild(headerCell);
    }
  
    for (let i = 0; i < rows; i++) {
      const row = table.insertRow();
      for (let j = 0; j < columns; j++) {
        const cell = row.insertCell();
        cell.contentEditable = "true";
        cell.style.border = "1px solid #ced4da";
        cell.style.padding = "8px";
        cell.textContent = `Cell ${i + 1}, ${j + 1}`;
      }
    }
  
    addSortingToTable(table, columns); // Добавляем сортировку к таблице
  
    // Кнопки для добавления/удаления строк и колонок
    tableElement.appendChild(createButton("Add Row", () => addRowToTable(table)));
    tableElement.appendChild(createButton("Add Column", () => addColumnToTable(table)));
    tableElement.appendChild(createButton("Remove Last Row", () => removeLastRowFromTable(table)));
    tableElement.appendChild(createButton("Remove Last Column", () => removeLastColumnFromTable(table)));
  
    addRemoveButton(tableElement);
    tableElement.appendChild(table);
    reportContainer.appendChild(tableElement);
  }
  
  // Load and display CSV data with editable cells and add row/column functionality
  function loadCSV(event) {
    const files = event.target.files;
    if (!files.length) return;
  
    Array.from(files).forEach((file) => {
      const reader = new FileReader();
      reader.onload = function (e) {
        const csv = e.target.result;
        const delimiter = csv.includes(";") ? ";" : ",";
        const rows = csv.trim().split("\n").filter(row => row).map(row => row.split(delimiter));
  
        const tableElement = document.createElement("div");
        tableElement.className = "report-element";
        const table = document.createElement("table");
        table.style.width = "100%";
        table.style.borderCollapse = "collapse";
  
        const headerRow = table.insertRow(); // Создаем строку заголовков
        rows[0].forEach(headerData => {
          const headerCell = document.createElement("th");
          headerCell.contentEditable = "true"; // Заголовки можно редактировать
          headerCell.style.border = "1px solid #ced4da";
          headerCell.style.padding = "8px";
          headerCell.textContent = headerData.trim();
          headerRow.appendChild(headerCell);
        });
  
        rows.slice(1).forEach(rowData => {
          const row = table.insertRow();
          rowData.forEach(cellData => {
            const cell = row.insertCell();
            cell.contentEditable = "true"; // Данные можно редактировать
            cell.style.border = "1px solid #ced4da";
            cell.style.padding = "8px";
            cell.textContent = cellData.trim();
          });
        });
  
        // Добавляем сортировку и кнопки для добавления/удаления строк и колонок
        addSortingToTable(table, rows[0].length);
        tableElement.appendChild(createButton("Add Row", () => addRowToTable(table)));
        tableElement.appendChild(createButton("Add Column", () => addColumnToTable(table)));
        tableElement.appendChild(createButton("Remove Last Row", () => removeLastRowFromTable(table)));
        tableElement.appendChild(createButton("Remove Last Column", () => removeLastColumnFromTable(table)));
  
        addRemoveButton(tableElement);
        tableElement.appendChild(table);
        document.getElementById("report-container").appendChild(tableElement);
      };
      reader.readAsText(file);
    });
  
    event.target.value = ""; // Clears the file input
  }
  
  // Load and display JSON data with editable inputs and add key-value functionality
  function loadJSON(event) {
    const file = event.target.files[0];
    if (!file) return;
  
    const reader = new FileReader();
    reader.onload = function (e) {
      const data = JSON.parse(e.target.result);
      for (const [key, value] of Object.entries(data)) {
        const element = document.createElement("div");
        element.className = "report-element";
  
        if (typeof value === "string") {
          const input = document.createElement("input");
          input.type = "text";
          input.value = value;
          input.contentEditable = "true"; // Поле ввода можно редактировать
          element.innerHTML = `<strong>${key}:</strong> `;
          element.appendChild(input);
        } else if (Array.isArray(value)) {
          const list = document.createElement("ul");
          value.forEach((item) => {
            const listItem = document.createElement("li");
            listItem.contentEditable = "true"; // Элементы списка можно редактировать
            listItem.textContent = item;
            list.appendChild(listItem);
          });
          element.innerHTML = `<strong>${key}:</strong>`;
  
          // Кнопка для добавления элемента в список
          element.appendChild(createButton("Add Item", () => {
            const newItem = document.createElement("li");
            newItem.contentEditable = "true";
            newItem.textContent = "New Item";
            list.appendChild(newItem);
          }));
  
          element.appendChild(list);
        } else if (typeof value === "object") {
          element.innerHTML = `<strong>${key}:</strong>`;
          const innerTable = document.createElement("table");
          const headerRow = innerTable.insertRow(); // Для ключей и значений
          headerRow.insertCell(0).textContent = "Key";
          headerRow.insertCell(1).textContent = "Value";
  
          for (const [subKey, subValue] of Object.entries(value)) {
            const row = innerTable.insertRow();
            row.insertCell(0).textContent = subKey;
            const cell = row.insertCell(1);
            cell.contentEditable = "true"; // Ячейки можно редактировать
            cell.textContent = subValue;
          }
  
          // Добавляем сортировку и кнопки для добавления/удаления строк и колонок
          addSortingToTable(innerTable, 2);
          element.appendChild(createButton("Add Row", () => addRowToTable(innerTable)));
          element.appendChild(createButton("Add Column", () => addColumnToTable(innerTable)));
          element.appendChild(createButton("Remove Last Row", () => removeLastRowFromTable(innerTable)));
          element.appendChild(createButton("Remove Last Column", () => removeLastColumnFromTable(innerTable)));
  
          element.appendChild(innerTable);
        }
  
        addRemoveButton(element);
        document.getElementById("report-container").appendChild(element);
      }
    };
    reader.readAsText(file);
  
    event.target.value = ""; // Clears the file input
  }
  
  // Add a remove button to an element
  function addRemoveButton(element) {
    const removeButton = createButton("✕", () => element.remove());
    element.appendChild(removeButton);
  }
  
  // Clear the report
  function clearReport() {
    document.getElementById("report-container").innerHTML = "";
  }
  
function exportToPDF() {
    // Создаем новый контент для печати
    const reportElements = document.getElementById("report-container").children;
    let content = '<html><head><style>';
    content += 'body { font-family: Arial, sans-serif; margin: 20px; }';
    content += 'h1 { font-size: 24px; margin-bottom: 20px; }';
    content += 'p { font-size: 18px; margin-bottom: 15px; }';
    content += 'table { width: 100%; border-collapse: collapse; margin-bottom: 20px; }';
    content += 'th, td { border: 1px solid black; padding: 8px; text-align: left; }';
    content += '</style></head><body>';

    // Перебираем все элементы отчёта и добавляем их в HTML
    Array.from(reportElements).forEach((element) => {
        if (element.querySelector("input")) {
            const title = element.querySelector("input").value;
            content += `<h1>${title}</h1>`;
        } else if (element.querySelector("textarea")) {
            const paragraph = element.querySelector("textarea").value;
            content += `<p>${paragraph}</p>`;
        } else if (element.querySelector("table")) {
            const table = element.querySelector("table");
            content += '<table>';
            for (let row of table.rows) {
                content += '<tr>';
                for (let cell of row.cells) {
                    content += `<td>${cell.textContent}</td>`;
                }
                content += '</tr>';
            }
            content += '</table>';
        }
    });

    content += '</body></html>';
    
    // Создаем новое окно с контентом для печати
    const printWindow = window.open('', '_blank');
    printWindow.document.write(content);
    printWindow.document.close();
    
    // Запускаем диалоговое окно печати
    printWindow.focus();
    printWindow.print();
    printWindow.close();
}



function exportToPPTX() {
  const pptx = new PptxGenJS();
  const slide = pptx.addSlide();

  // Перебираем все элементы отчёта и добавляем их на слайд
  const reportElements = document.getElementById("report-container").children;
  for (let i = 0; i < reportElements.length; i++) {
      const element = reportElements[i];

      if (element.querySelector("input")) {
          // Элемент заголовка
          const title = element.querySelector("input").value;
          slide.addText(title, { x: 0.5, y: 0.5, fontSize: 18 });
      } else if (element.querySelector("textarea")) {
          // Элемент абзаца
          const paragraph = element.querySelector("textarea").value;
          slide.addText(paragraph, { x: 0.5, y: 1.5, fontSize: 14 });
      } else if (element.querySelector("table")) {
          // Элемент таблицы
          const table = element.querySelector("table");
          const tableData = [];
          for (let row of table.rows) {
              const rowData = Array.from(row.cells).map(cell => cell.textContent);
              tableData.push(rowData);
          }
          slide.addTable(tableData, { x: 0.5, y: 2.5, fontSize: 10 });
      }
  }

  pptx.writeFile("report.pptx").catch(err => console.error("Ошибка при экспорте в PPTX:", err));
}

// Create chart
function createChart() {
  const labels = document.getElementById('chartLabels').value.split(',');
  const data = document.getElementById('chartData').value.split(',').map(Number);
  if (labels.length !== data.length) {
    alert('Labels and data must have the same length!');
    return;
  }

  const chartContainer = document.createElement('div');
  chartContainer.className = 'report-element';

  const canvas = document.createElement('canvas');
  chartContainer.appendChild(canvas);
  document.getElementById('report-container').appendChild(chartContainer);

  const ctx = canvas.getContext('2d');
  const chartConfig = {
    type: currentChartType === 'histogram' ? 'bar' : currentChartType,  // Treat histogram as bar chart
    data: {
      labels: labels,
      datasets: [{
        label: currentChartType.charAt(0).toUpperCase() + currentChartType.slice(1) + ' Chart',
        data: data,
        backgroundColor: currentChartType === 'pie' ? getPieColors(labels.length) : 'rgba(75, 192, 192, 0.2)',
        borderColor: currentChartType === 'pie' ? getPieColors(labels.length) : 'rgba(75, 192, 192, 1)',
        borderWidth: 1
      }]
    },
    options: {
      scales: currentChartType === 'pie' ? {} : {
        y: {
          beginAtZero: true
        }
      }
    }
  };

  const chart = new Chart(ctx, chartConfig);

  // Add "Copy Chart" button
  const copyButton = document.createElement('button');
  copyButton.textContent = 'Copy Chart';
  copyButton.onclick = () => copyChartToClipboard(canvas);
  chartContainer.appendChild(copyButton);

  closeChartModal();
}

// Function to copy chart to clipboard
async function copyChartToClipboard(canvas) {
  try {
    const blob = await new Promise(resolve => canvas.toBlob(resolve));
    const item = new ClipboardItem({ 'image/png': blob });
    await navigator.clipboard.write([item]);
    alert('Chart copied to clipboard!');
  } catch (error) {
    console.error('Error copying chart:', error);
    alert('Failed to copy the chart.');
  }
}

