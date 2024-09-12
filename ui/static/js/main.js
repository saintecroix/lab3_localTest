let select = document.getElementById("Goods");

if (select) {
	select.addEventListener("change", myFunction);
}


function myFunction() {
	let select = document.getElementById("Goods");
	let textInput = document.getElementById("text");
	let textLabel = document.getElementById("ltext");
	if (select.value !== "") {
		textInput.style = "display: content";
		textLabel.style = "display: content";
	}
}

/*----------------------------------------------------------------------------------------*/

// Получение элемента select HTML с идентификатором "Duos"
const duosSelect = document.getElementById('Duos');

// Добавление обработчика события изменения состояния
duosSelect?.addEventListener('change', (event) => {
	// Проверка значения value
	if (event.target.value !== '') {
		// Значение value не пустое, отправляем GET запрос на сервер
		fetch('/getSecondPerSearch', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ checkFirstEl: event.target.value })
		})
			.then(response => {
				if (response.ok) {
					// Успешный ответ, данные получены с сервера
					return response.json();
				} else {
					// Ошибка при отправке запроса
					console.error("Ошибка при получении данных с сервера.");
				}
			})
			.then(data => {
				// Данные получены и доступны в переменной data
				console.log("Полученные данные: ", data);

				// Получение списка HTML
				const gruzSelect = document.getElementById('Gruz-list');
				for (let i = gruzSelect.options.length - 1; i > 0; i--) {
					gruzSelect.remove(i);
				}

				// Перебор данных и создание элементов списка
				data.forEach(gruz => {
					const option = document.createElement('option');
					option.value = gruz.id;
					option.textContent = `${gruz.name} (ETSNG: ${gruz.etsng}, GNG: ${gruz.gng})`;
					gruzSelect.appendChild(option);
				});
			})
			.catch(error => {
				// Ошибка при выполнении запроса
				console.error("Ошибка при выполнении запроса на сервер: ", error);
			});
	}
});

/*----------------------------------------------------------------------------------------*/
const form = document.getElementById('my-form');
const gruzSelect = document.getElementById('Gruz-list');

form?.addEventListener('submit', (event) => {
	event.preventDefault();
	fetch('/letDuoSearch', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			checkDuos: duosSelect.value,
			checkGruz: gruzSelect.value
		})
	})
		.catch(error => {
			// Ошибка при выполнении запроса
			console.error("Ошибка при выполнении запроса на сервер: ", error);
		})
		.then(response => {
			if (response.ok) {
				// Успешный ответ, данные получены с сервера
				return response.json();
			} else {
				// Ошибка при отправке запроса
				console.error("Ошибка при получении данных с сервера.");
			}
		})
		.then(data => {
			// Данные получены и доступны в переменной data
			console.log("Полученные данные: ", data);

			const tableBody = document.getElementById("app-table-body");
			tableBody.innerHTML = "";

			data.forEach(app => {
				const row = tableBody.insertRow();

				row.insertCell().innerText = app.number_app;
				row.insertCell().innerText = app.regDate_app;
				row.insertCell().innerText = app.status_app;
				row.insertCell().innerText = app.provideDate_app;
				row.insertCell().innerText = app.departureType_app;
				row.insertCell().innerText = app.goods_app;
				row.insertCell().innerText = app.originState_app;
				row.insertCell().innerText = app.enterStation_app;
				row.insertCell().innerText = app.regionDepart_app;
				row.insertCell().innerText = app.roadDepart_app;
				row.insertCell().innerText = app.stationDepart_app;
				row.insertCell().innerText = app.consigner_app;
				row.insertCell().innerText = app.stateDestination_app;
				row.insertCell().innerText = app.exitStation_app;
				row.insertCell().innerText = app.regionDestination_app;
				row.insertCell().innerText = app.roadDestination_app;
				row.insertCell().innerText = app.stationDestination_app;
				row.insertCell().innerText = app.consignee_app;
				row.insertCell().innerText = app.wagonType_app;
				row.insertCell().innerText = app.property_app;
				row.insertCell().innerText = app.wagonOwner_app;
				row.insertCell().innerText = app.payer_app;
				row.insertCell().innerText = app.roadOwner_app;
				row.insertCell().innerText = app.transportManager_app;
				row.insertCell().innerText = app.tonsDeclared_app;
				row.insertCell().innerText = app.tonsAccepted_app;
				row.insertCell().innerText = app.wagonDeclared_app;
				row.insertCell().innerText = app.wagonAccepted_app;
				row.insertCell().innerText = app.filingDate_app;
				row.insertCell().innerText = app.agreementDate_app;
				row.insertCell().innerText = app.approvalDate_app;
				row.insertCell().innerText = app.startDate_app;
				row.insertCell().innerText = app.stopDate_app;
			});
		})
		.catch(error => {
			// Ошибка при выполнении запроса
			console.error("Ошибка при выполнении запроса на сервер: ", error);
		});
});

/*----------------------------------------------------------------------------------------*/

document.querySelectorAll('.menu__switcher > .menu__link').forEach(link => {
	link.addEventListener('click', (event) => {
		event.preventDefault();
		const parent = link.parentElement;
		parent.classList.toggle('active');
	});
});

/*----------------------------------------------------------------------------------------*/

// Получить ссылку на таблицу
const table = document.querySelector('table');

// Событие ввода для поля поиска
document.querySelector('#search-input').addEventListener('input', function() {
	// Получить значение поиска
	const searchValue = this.value.toLowerCase();

	// Скрыть все строки, которые не содержат значение поиска
	table.querySelectorAll('tbody tr').forEach(function(row) {
		const rowText = row.textContent.toLowerCase();
		if (!rowText.includes(searchValue)) {
			row.style.display = 'none';
		} else {
			row.style.display = '';
		}
	});
});

// Событие клика для заголовков столбцов для сортировки
table.querySelectorAll('th').forEach(function(header) {
	header.addEventListener('click', function() {
		// Получить индекс столбца
		const columnIndex = header.cellIndex;

		// Получить все строки в таблице
		const rows = Array.from(table.querySelectorAll('tbody tr'));

		// Отсортировать строки по значению указанного столбца
		rows.sort(function(a, b) {
			const aValue = a.children[columnIndex].textContent;
			const bValue = b.children[columnIndex].textContent;

			return aValue.localeCompare(bValue);
		});

		// Удалить все существующие строки из таблицы
		table.querySelectorAll('tbody tr').forEach(function(row) {
			row.remove();
		});

		// Добавить отсортированные строки обратно в таблицу
		rows.forEach(function(row) {
			table.querySelector('tbody').appendChild(row);
		});
	});
});
