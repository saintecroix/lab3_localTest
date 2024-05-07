let select = document.getElementById("Goods");

if(select ){
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
console.log("test")

/*----------------------------------------------------------------------------------------*/

// Получение элемента select HTML с идентификатором "Duos"
const duosSelect = document.getElementById('Duos');

// Добавление обработчика события изменения состояния
duosSelect.addEventListener('change', (event) => {
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

form.addEventListener('submit', (event) => {
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
